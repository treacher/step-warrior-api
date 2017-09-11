package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/treacher/step-warrior-api/models"
	"gopkg.in/pg.v5"
)

type Response struct {
	Name  string     `json:'name'`
	Id    string     `json:'id'`
	Error OAuthError `json:'error'`
}

type OAuthError struct {
	Message string `json:'message'`
	Type    string `json:'type'`
}

type OAuthHandler struct {
	db *pg.DB
}

func NewOAuthHandler(db *pg.DB) *OAuthHandler {
	return &OAuthHandler{
		db: db,
	}
}

func (handler *OAuthHandler) OAuthHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		var responseBody Response

		request, err := http.Get(fmt.Sprintf("https://graph.facebook.com/me?access_token=%s", token))

		if err != nil {
			panic(err)
		}

		decoder := json.NewDecoder(request.Body)

		decoder.Decode(&responseBody)

		if responseBody.Error.Type == "OAuthException" {
			Error(r.Context(), w, http.StatusUnauthorized, "unauthorized", "invalid OAuth request")
			return
		}

		queryString := fmt.Sprintf("SELECT * FROM users WHERE identifier = '%s' LIMIT 1;", responseBody.Id)

		var user models.User

		result, err := handler.db.Query(&user, queryString)

		if !(result != nil && result.RowsReturned() == 1) {
			user := models.User{Identifier: responseBody.Id}
			user.Persist(handler.db)
		}

		ctx := context.WithValue(r.Context(), "User", user)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
