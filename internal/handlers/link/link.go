package link

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/adii1203/link/internal/response"
	"github.com/adii1203/link/internal/types"
	"github.com/adii1203/link/internal/utils"
	"io"
	"log/slog"
	"net/http"
)

func New() http.HandlerFunc {
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

		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"})
	}
}
