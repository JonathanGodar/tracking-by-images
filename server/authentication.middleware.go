package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	s "strings"

	"github.com/golang-jwt/jwt"
)

const bearerTokenPrefix = "Bearer "

func authenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		authToken, authTokenIsSet := getAuthToken(r)

		if !authTokenIsSet {
			ctx := context.WithValue(r.Context(), SignedInUserKey, nil)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		token, err := jwt.ParseWithClaims(authToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected token signing method")
			}
			return secret, nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

		claims, ok := token.Claims.(*UserClaims)
		if !ok {
			log.Println("Could not decode jwt")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized error"))
		}

		log.Println(claims.User)
		if err != nil {
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, SignedInUserKey, claims.User)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func getAuthToken(r *http.Request) (string, bool) {
	authHeaders, authHeaderExists := r.Header["Authorization"]

	if !authHeaderExists {
		return "", false
	}

	if len(authHeaders) != 1 {
		log.Println("No headers")
		return "", false
	}

	authHeader := authHeaders[0]
	if !authHeaderExists || !s.HasPrefix(authHeader, bearerTokenPrefix) {
		return "", false
	}

	authHeader = authHeader[len(bearerTokenPrefix):]

	return authHeader, true
}
