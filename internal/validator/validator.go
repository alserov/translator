package validator

import (
	"fmt"
	"sync"
)

type validator struct {
	wg *sync.WaitGroup
}

type Validator interface {
	Validate(str interface{}, validations ...Fn) error
}

func NewValidator() Validator {
	return &validator{
		wg: &sync.WaitGroup{},
	}
}

type Fn func(interface{}) error

func (v *validator) Validate(str interface{}, validations ...Fn) error {
	v.wg.Add(len(validations))
	chErr := make(chan error, len(validations))

	for _, validation := range validations {
		validation := validation
		go func() {
			defer v.wg.Done()
			if err := validation(str); err != nil {
				chErr <- fmt.Errorf("validation fail: %v", err)
			}
		}()
	}

	go func() {
		v.wg.Wait()
		close(chErr)
	}()

	if err := <-chErr; err != nil {
		return err
	}

	return nil
}
