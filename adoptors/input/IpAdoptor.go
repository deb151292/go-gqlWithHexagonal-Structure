package adoptor

type CalculatorAdapter struct{}

func NewCalculatorAdapter() *CalculatorAdapter {
	return &CalculatorAdapter{}
}

func (c *CalculatorAdapter) Add(a, b int) (int, error) {
	// Implement addition logic here
	return a + b, nil
}