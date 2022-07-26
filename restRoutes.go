package main

import (
	"example.com/m/business"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"strconv"
)

const bitSize = 64

type calcParams struct {
	velocity float64
	height   float64
	angle    float64
}

func validateInput(input calcParams) error {
	if input.velocity > 0 && input.height >= 0 && input.angle <= 90 && input.angle >= 0 {
		if input.angle == 0 && input.height == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "bad input: angle cannot be 0 degrees while height is 0 meters (and vice versa)")
		}
		return nil
	}

	return echo.NewHTTPError(http.StatusBadRequest, "bad input: velocity should be positive, height should be 0 or above, and angle should be min 0 and max 90 degrees")
}

func calculationHandler(c echo.Context) error {
	var input calcParams

	err := parseInput(c, &input)
	if err != nil {
		return err
	}

	err = validateInput(input)
	if err != nil {
		return err
	}

	finalLocation := fmt.Sprintf("%f", business.CalculateFinalLocation(input.velocity, input.height, input.angle))
	finalVelocity := fmt.Sprintf("%f", business.CalculateFinalVelocity(input.velocity, input.height, input.angle))
	finalAngle := fmt.Sprintf("%f", business.CalculateFinalAngle(input.velocity, input.height, input.angle))

	return c.String(http.StatusOK, fmt.Sprintf("Final location is: %s\nFinal velocity is: %s\nFinal angle is: %s", finalLocation, finalVelocity, finalAngle))
}

func parseInput(c echo.Context, inputs *calcParams) error {
	velocity, err := strconv.ParseFloat(c.QueryParam("velocity"), bitSize)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("failed receiving velocity: %w", err).Error())
	}

	inputs.velocity = velocity

	height, err := strconv.ParseFloat(c.QueryParam("height"), bitSize)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("failed receiving height: %w", err).Error())
	}

	inputs.height = height

	angle, err := strconv.ParseFloat(c.QueryParam("angle"), bitSize)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("failed receiving angle: %w", err).Error())
	}

	inputs.angle = angle

	return nil
}

func registerAPI(e *echo.Echo) {
	e.POST("/calculate", calculationHandler)
}
