package application

import port "travel.go/ports/input"

type CalculatorService struct {
	CalculatorInput port.Calculator
}

func NewCalculatorService(calculatorInput port.Calculator) *CalculatorService {
	return &CalculatorService{CalculatorInput: calculatorInput}
}

func (s *CalculatorService) Add(a, b int) (int, error) {
	return s.CalculatorInput.Add(a, b)
}
