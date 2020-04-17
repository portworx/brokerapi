package req_auth_header

import (
	"context"
	"net/http"
	"strings"
)

const (
	iamToken = "IBM_IAM_TOKEN"
)

func AddToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		bearerTokens := strings.SplitN(req.Header.Get("Authorization"), " ", 2)
		newCtx := context.WithValue(req.Context(), iamToken, bearerTokens)
		next.ServeHTTP(w, req.WithContext(newCtx))
	})
}
