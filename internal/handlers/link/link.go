package link

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/adii1203/link/internal/initializers"
	"github.com/adii1203/link/internal/response"
	"github.com/adii1203/link/internal/types"
	"github.com/adii1203/link/internal/utils"
)

func generateUniqueKey(storage initializers.Storage, slug string, length int) (string, error) {
	key := slug
	if key == "" {
		key = utils.GenerateKey(length)
	}

	for storage.GetKey(key) {
		key = utils.GenerateKey(length)
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

		id, err := storage.CreateLink(link.DestinationUrl, key)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GenericError(err))
			return
		}

		responseData := map[string]string{
			"key":             key,
			"id":              strconv.Itoa(int(id)),
			"destination_url": link.DestinationUrl,
		}

		response.WriteJson(w, http.StatusCreated, response.SuccessResponse(responseData))
	}
}
