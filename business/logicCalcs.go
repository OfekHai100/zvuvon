package business

import (
	"math"
)

func CalculateFinalLocation(velocity float64, height float64, angle float64) float64 {
	if angle == 0 {
		return velocity * math.Sqrt(2*height/gravityAcceleration*(-1))
	}

	duration := getQuadraticEquationResults(gravityAcceleration/2, velocity*math.Sin(degreesToRadians(angle)), height)

	return velocity * math.Cos(degreesToRadians(angle)) * duration
}

func CalculateFinalVelocity(velocity float64, height float64, angle float64) float64 {
	fallDuration := getFallDuration(getMaxHeight(velocity, height, angle))

	return math.Sqrt(math.Pow(velocity*math.Cos(degreesToRadians(angle)), 2) + math.Pow(gravityAcceleration*fallDuration, 2))
}

//bonus
func CalculateFinalAngle(velocity float64, height float64, angle float64) float64 {
	if angle == 90 {
		return 270
	}

	fallDuration := getFallDuration(getMaxHeight(velocity, height, angle))

	return math.Abs(radiansToDegrees(math.Atan((gravityAcceleration * fallDuration) / (velocity * math.Cos(degreesToRadians(angle))))))
}
