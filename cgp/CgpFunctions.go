package cgp

import "math"

func evaluateFunction(leftOperand int, rightOperand int, functionNumber int, parameter int) int {
	switch functionNumber {
	case 0:
		return leftOperand | rightOperand
	case 1:
		return leftOperand & rightOperand
	case 2:
		return leftOperand / (1 + rightOperand + parameter)
	case 3:
		return (leftOperand * rightOperand) % 255
	case 4:
		if leftOperand > rightOperand {
			return leftOperand - rightOperand
		} else {
			return rightOperand - leftOperand
		}
	case 5:
		return 255 - leftOperand
	case 6:
		return int(math.Abs(math.Cos(float64(leftOperand)) * 255))
	case 7:
		return int(math.Abs(math.Sin(float64(leftOperand)) * 255))
	case 8:
		return int(math.Abs(math.Tan((((float64(leftOperand % 45)) * math.Pi) / 180)) * 255))
	case 9:
		return int(math.Abs(float64(int(math.Tan(float64(leftOperand))*255) % 255)))
	case 10:
		value := math.Pow(float64(leftOperand)-float64(parameter), 2) + math.Pow(float64(rightOperand)-float64(parameter), 2)
		if value <= 255 {
			return int(value)
		} else {
			return 255
		}
	case 11:
		return leftOperand%(parameter+1) + (255 - parameter)
	case 12:
		return (leftOperand + rightOperand) / 2
	case 13:
		if leftOperand > rightOperand {
			return 255 * ((rightOperand + 1) / (leftOperand + 1))
		} else {
			return 255 * ((leftOperand + 1) / (rightOperand + 1))
		}
	case 14:
		return int(math.Sqrt(math.Abs(math.Pow(float64(leftOperand), 2)+math.Pow(float64(rightOperand), 2)-2*math.Pow(float64(parameter), 2)))) % 255
	}

	return 0
}
