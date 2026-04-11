package middleware

import (
	"context"
	"log"
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
			log.Printf("[AuthMiddleware] Token validation failed: %v", err)
			response.Error(w, http.StatusUnauthorized, "invalid or expired token")
			return
		}

		log.Printf("[AuthMiddleware] User ID: %d, Role: %s", claims.UserID, claims.Role)
		ctx := context.WithValue(r.Context(), UserContextKey, claims.UserID)
		ctx = context.WithValue(ctx, RoleContextKey, claims.Role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		roleVal := r.Context().Value(RoleContextKey)
		if roleVal == nil {
			log.Printf("[AdminOnly] Role not found in context")
			response.Error(w, http.StatusForbidden, "forbidden")
			return
		}
		
		role := roleVal.(string)
		if role != "admin" {
			log.Printf("[AdminOnly] Denied access for role: %s", role)
			response.Error(w, http.StatusForbidden, "admin access required")
			return
		}
		log.Printf("[AdminOnly] Access granted for admin")
		next.ServeHTTP(w, r)
	})
}
