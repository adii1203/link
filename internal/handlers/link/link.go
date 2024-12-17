package link

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/adii1203/link/internal/initializers"
	"github.com/adii1203/link/internal/response"
	"github.com/adii1203/link/internal/types"
	"github.com/adii1203/link/internal/utils"
	"io"
	"log/slog"
	"net/http"
	"strconv"
)

func New(storage initializers.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Creating link")

		var link types.Link

		err := json.NewDecoder(r.Body).Decode(&link)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GenericError(fmt.Errorf("empty body")))
			return
		}
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GenericError(err))
			return
		}

		// validation
		if err := utils.ValidateStruct(link); err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.ValidationErrors(err))
			return
		}

		key := utils.GenerateKey(6)

		// check if key is already exist
		isKeyExist := storage.GetKey(key)
		for isKeyExist {
			key = utils.GenerateKey(6)
			isKeyExist = storage.GetKey(key)
		}

		id, err := storage.CreateLink(link.DestinationUrl, key)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GenericError(err))
			return
		}

		responseData := map[string]string{"key": key, "id": strconv.Itoa(int(id)), "destination_url": link.DestinationUrl}

		response.WriteJson(w, http.StatusCreated, response.SuccessResponse(responseData))
	}
}
