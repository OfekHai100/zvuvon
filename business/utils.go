package business

import (
	"errors"
	"math"
)

const gravityAcceleration = -9.81

func radiansToDegrees(radAngle float64) float64 {
	return radAngle * (180 / math.Pi)
}

func degreesToRadians(regAngle float64) float64 {
	return regAngle * (math.Pi / 180)
}

func getQuadraticEquationResults(a float64, b float64, c float64) float64 {
	var root1 float64
	var root2 float64

	discriminant := (b * b) - (4 * a * c)

	if discriminant > 0 {
		root1 = (-b + (math.Sqrt(discriminant))) / (2 * a)
		root2 = (-b - (math.Sqrt(discriminant))) / (2 * a)
	} else if discriminant == 0 {
		root1 = -b / (2 * a)
		root2 = -b / (2 * a)
	}

	if root1 >= 0 {
		return root1
	}

	return root2
}

func getMaxHeight(velocity float64, height float64, angle float64) float64 {
	if angle == 0 {
		return height
	} else if angle == 90 {
		return math.Pow(velocity, 2) / 2 * (gravityAcceleration * (-1))
	}

	return math.Pow(velocity, 2)*math.Pow(math.Sin(degreesToRadians(angle)), 2)/2*gravityAcceleration*(-1) + height
}

func getFallDuration(maxHeight float64) float64 {
	return getQuadraticEquationResults(gravityAcceleration/2, 0, maxHeight)
}

func validateInput(velocity float64, height float64, angle float64) error {
	if velocity > 0 && height >= 0 && angle <= 90 && angle >= 0 {
		if angle == 0 && height == 0 {
			return errors.New("bad input: angle cannot be 0 degrees while height is 0 meters (and vice versa)")
		}
		return nil
	}

	return errors.New("bad input: velocity should be positive, height should be 0 or above, and angle should be min 0 and max 90 degrees")
}
