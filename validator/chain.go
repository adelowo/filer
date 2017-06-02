package validator

import (
	"github.com/adelowo/filer"
)

//ChainedValidator is a validator that is composed of other validators
//It doesn't perform any validation itself
//instead it defers to the validators it has been composed of
type ChainedValidator struct {
	chain []Validator
}

//NewChainedValidator returns an instance of a ChainedValidator.
//It panics if no validator is added to the chain
func NewChainedValidator(v ...Validator) *ChainedValidator {

	chained := &ChainedValidator{chain: v}

	if len(chained.chain) == 0 {
		panic(`chain validator:
      Must have at least one validator present in the chain`)
	}

	return chained
}

func (c *ChainedValidator) Validate(f filer.File) (bool, error) {

	if len(c.chain) == 1 {
		return c.chain[0].Validate(f)
	}

	for _, validator := range c.chain {
		if valid, err := validator.Validate(f); err != nil {
			return valid, err
		}
	}

	return true, nil
}
