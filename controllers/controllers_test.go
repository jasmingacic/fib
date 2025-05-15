package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/fibonacci/:n", GetFibonacci)
	return r
}

func TestGetFibonacci(t *testing.T) {
	router := setupRouter()

	tests := []struct {
		name           string
		param          string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Missing parameter",
			param:          "",                  // no value for :n
			expectedStatus: http.StatusNotFound, // Gin returns 404 for missing path param
		},
		{
			name:           "Non-integer parameter",
			param:          "abc",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Parameter 'n' must be an integer"}`,
		},
		{
			name:           "Negative parameter",
			param:          "-5",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Parameter 'n' cannot be negative"}`,
		},
		{
			name:           "Too large parameter",
			param:          "200",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Parameter 'n' exceeds maximum Fibonacci sequence length (93) for the current architecture"}`,
		},
		{
			name:           "Valid parameter",
			param:          "5",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"numbers":[0,1,1,2,3]}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/fibonacci/"+tt.param, nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedBody != "" {
				assert.JSONEq(t, tt.expectedBody, w.Body.String())
			}
		})
	}
}
