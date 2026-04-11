package middleware

import (
	"context"
	"net/http"
	"strings"
	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/pkg/response"
	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/pkg/utils"
)

type contextKey string

const (
	UserContextKey contextKey = "user"
	RoleContextKey contextKey = "role"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.Error(w, http.StatusUnauthorized, "missing authorization header")
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Error(w, http.StatusUnauthorized, "invalid authorization format")
			return
		}

		claims, err := utils.ValidateToken(parts[1])
		if err != nil {
			response.Error(w, http.StatusUnauthorized, "invalid or expired token")
			return
		}

		ctx := context.WithValue(r.Context(), UserContextKey, claims.UserID)
		ctx = context.WithValue(ctx, RoleContextKey, claims.Role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role := r.Context().Value(RoleContextKey).(string)
		if role != "admin" {
			response.Error(w, http.StatusForbidden, "admin access required")
			return
		}
		next.ServeHTTP(w, r)
	})
}
