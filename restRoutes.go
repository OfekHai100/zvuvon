package main

import (
	"fmt"
	"github.com/labstack/echo"
	"strconv"
)

const bitSize = 64

func calculationHandler(c echo.Context) error {
	velocity, err := strconv.ParseFloat(c.Param("velocity"), bitSize)
	if err != nil {
		return fmt.Errorf("failed receiving velocity: %w", err)
	}

	height, err := strconv.ParseFloat(c.Param("height"), bitSize)
	if err != nil {
		return fmt.Errorf("failed receiving height: %w", err)
	}

	angle, err := strconv.ParseFloat(c.Param("angle"), bitSize)
	if err != nil {
		return fmt.Errorf("failed receiving angle: %w", err)
	}

	//return c.JSON(http.StatusOK, business.CalculateFinalLocation(velocity, height, angle))
}

func registerAPI(e *echo.Echo) {
	e.POST("/calculate", calculationHandler)
}
