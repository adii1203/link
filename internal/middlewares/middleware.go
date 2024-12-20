package middlewares

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/adii1203/link/internal/response"
	"github.com/adii1203/link/internal/types"
)

func ValidatePayload(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var link types.Link

		if r.ContentLength == 0 {
			response.WriteJson(w, http.StatusBadRequest, response.GenericError(fmt.Errorf("invalid request payload")))
		}

		err := json.NewDecoder(r.Body).Decode(&link)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GenericError(fmt.Errorf("invalid request payload")))
			return
		}
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GenericError(fmt.Errorf("failed to create link, please try again later")))
			return
		}

		ctx := context.WithValue(r.Context(), "validatedPayload", link)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func IsCrawler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userAgent := r.Header.Get("User-Agent")
		crawlers := []string{
			"twitterbot",
			"facebookexternalhit",
			"linkedinbot",
			"googlebot",
			"postmanruntime",
		}

		for _, crawler := range crawlers {
			if strings.Contains(strings.ToLower(userAgent), crawler) {
				// retunr html with metadata
				ctx := context.WithValue(r.Context(), "isCrawler", "true")
				next.ServeHTTP(w, r.WithContext(ctx))
			}
		}
		next.ServeHTTP(w, r)

	})
}
