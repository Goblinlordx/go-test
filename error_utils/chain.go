package error_utils

type errorChain struct {
	err *error
}

func NewErrorChain() errorChain {
	return errorChain{}
}

func (ec errorChain) Next(fn func() error) errorChain {
	if ec.err == nil {
		err := fn()
		if err != nil {
			ec.err = &err
		}
	}

	return ec
}

func (ec errorChain) Check() error {
	if ec.err == nil {
		return nil
	}

	return *ec.err
}
