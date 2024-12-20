package link

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"

	"github.com/adii1203/link/internal/initializers"
	"github.com/adii1203/link/internal/models"
	"github.com/adii1203/link/internal/response"
	"github.com/adii1203/link/internal/types"
	"github.com/adii1203/link/internal/utils"
)

func generateUniqueKey(storage initializers.Storage, slug string, length int) (string, error) {
	key := slug
	if key == "" {
		key = utils.GenerateKey(length)
	}

	keyExist, _, _ := storage.GetKey(key)

	for keyExist {
		key = utils.GenerateKey(length)
		keyExist, _, _ = storage.GetKey(key)
	}

	return key, nil
}

func New(storage initializers.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Creating link")

		link := r.Context().Value("validatedPayload").(types.Link)

		// validation
		if err := utils.ValidateStruct(link); err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.ValidationErrors(err))
			return
		}

		key, err := generateUniqueKey(storage, link.Slug, 6)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GenericError(fmt.Errorf("failed to create link, please try again later")))
		}

		metadata := utils.GetMetadata(link.DestinationUrl)

		newLink, err := storage.CreateLink(&models.Link{
			DestinationUrl: link.DestinationUrl,
			Slug:           key,
			Metadata: models.Metadata{
				Title:       metadata.Title,
				Description: metadata.Description,
				Image:       metadata.Image,
			},
		})
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GenericError(err))
			return
		}

		response.WriteJson(w, http.StatusCreated, response.SuccessResponse(newLink))
	}
}

func Redirect(storage initializers.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Redirecting to destination")

		slug := r.PathValue("slug")
		isCrawler := r.Context().Value("isCrawler")

		if slug == "" {
			response.WriteJson(w, http.StatusBadRequest, response.GenericError(fmt.Errorf("invalid slug")))
			return
		}

		_, link, err := storage.GetKey(slug)

		if err != nil {
			response.WriteJson(w, http.StatusNotFound, response.GenericError(fmt.Errorf("link not found")))
			return
		}

		if isCrawler == "true" {
			// check for crewler and return metadata
			tmpl, err := template.ParseFiles("internal/templates/meta_redirect.html")
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "text/html")
			tmpl.Execute(w, link)
		}

		http.Redirect(w, r, link.DestinationUrl, http.StatusFound)
	}
}

func Metadata() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")
		if url == "" {
			response.WriteJson(w, http.StatusBadRequest, response.GenericError(fmt.Errorf("url query param is required")))
			return
		}

		responseData := utils.GetMetadata(url)

		response.WriteJson(w, http.StatusOK, response.SuccessResponse(responseData))

	}
}
