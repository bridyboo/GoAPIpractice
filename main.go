package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CalculatorResult struct {
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
}

func main() {
	server := gin.Default()

	// Correcting the route paths
	server.POST("/add/:a/:b", additionFunction)
	server.POST("/multiply/:a/:b", multiplyFunction)
	server.POST("/divide/:a/:b", divisionFunction)
	server.POST("/subtract/:a/:b", subtractFunction)

	// Start the server on port 8080
	server.Run(":8080")
}

func subtractFunction(c *gin.Context) {
	a, b, err := stringToFloat(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, CalculatorResult{Error: err.Error()})
		return
	}

	result := a - b

	c.JSON(http.StatusOK, CalculatorResult{Result: result})
}

func divisionFunction(c *gin.Context) {
	a, b, err := stringToFloat(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, CalculatorResult{Error: err.Error()})
	}

	result := a / b

	c.JSON(http.StatusOK, CalculatorResult{Result: result})
}

func additionFunction(c *gin.Context) {
	a, b, err := stringToFloat(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, CalculatorResult{Error: err.Error()})
		return
	}
	result := a + b
	c.JSON(http.StatusOK, CalculatorResult{Result: result})
}

func multiplyFunction(c *gin.Context) {
	a, b, err := stringToFloat(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, CalculatorResult{Error: err.Error()})
		return
	}
	result := a * b
	c.JSON(http.StatusOK, CalculatorResult{Result: result})
}

func stringToFloat(c *gin.Context) (float64, float64, error) {
	aString := c.Param("a")
	bString := c.Param("b")

	// Convert string to float64
	a, err := strconv.ParseFloat(aString, 64)
	if err != nil {
		return 0, 0, err
	}

	b, err := strconv.ParseFloat(bString, 64)
	if err != nil {
		return 0, 0, err
	}

	return a, b, nil
}
