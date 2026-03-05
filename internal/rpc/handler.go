package rpc

import "fmt"

type Calculator struct{}

func (c *Calculator) Add(params []interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, fmt.Errorf("Add expects 2 parameters")
	}

	a := int(params[0].(float64))
	b := int(params[1].(float64))

	return a + b, nil
}

func (c *Calculator) Subtract(params []interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, fmt.Errorf("Subtract expects 2 parameters")
	}

	a := int(params[0].(float64))
	b := int(params[1].(float64))

	return a - b, nil
}
func (c *Calculator) Mul(params []interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, fmt.Errorf("Mul expects 2 parameters")
	}
	a := int(params[0].(float64))
	b := int(params[1].(float64))
	return a * b, nil
}

func (c *Calculator) Div(params []interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, fmt.Errorf("Div expects 2 parameters")
	}
	a := int(params[0].(float64))
	b := int(params[1].(float64))
	if b == 0 {
		return nil, fmt.Errorf("division by zero is not possible")
	}
	return (float64)(a / b), nil
}
