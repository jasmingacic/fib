package controllers

import (
	"fib/fibonacci"
	"fmt"
	"net/http"
	"strconv"

	"fib/utils"

	"github.com/gin-gonic/gin"
)

// FibResponse represents the JSON response containing Fibonacci numbers
type FibResponse struct {
	Numbers []int `json:"numbers"`
}

// ErrorResponse represents the JSON error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// getFibonacci handles requests for Fibonacci numbers
func GetFibonacci(c *gin.Context) {
	nStr := c.Param("n")
	if nStr == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Missing parameter 'n'"})
		return
	}

	n, err := strconv.Atoi(nStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Parameter 'n' must be an integer"})
		return
	}

	if n < 0 {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Parameter 'n' cannot be negative"})
		return
	}

	maxPossibleFibonacci := utils.ArchitectureBitSizeMaxSequence()
	// Check if 'n' is too large for the current architecture
	if n > maxPossibleFibonacci {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: fmt.Sprintf("Parameter 'n' exceeds maximum Fibonacci sequence length (%d) for the current architecture", maxPossibleFibonacci)})
		return
	}

	// Generate Fibonacci numbers
	fibNumbers := fibonacci.GenerateFibonacci(n)

	// Return the result as JSON
	c.JSON(http.StatusOK, FibResponse{Numbers: fibNumbers})
}
