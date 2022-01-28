package data

import (
	"testing"
)

func Test_ChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Tea",
		Price: 1.50,
		SKU:   "abs-endic-wdn",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}

}
