package data

import "testing"

func TestGetProducts(t *testing.T) {

	p := &Product{
		Name:  "babu",
		Price: 90.00,
		SKU:   "abc-def-ghi",
	}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
