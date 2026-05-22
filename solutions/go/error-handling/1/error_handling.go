package erratum

import "errors"

func Use(opener ResourceOpener, input string) (err error) {
	var resource Resource
	var transientErr TransientError

	for {
		resource, err = opener()
		if err == nil {
			break
		}
		if errors.As(err, &transientErr) {
			continue
		}
		return err
	}
	defer resource.Close()

	defer func() {
		r := recover()
		if r == nil {
			return
		}
		var frobErr FrobError
		rErr, ok := r.(error)
		if !ok {
			err = errors.New("unknown error")
			return
		}
		if errors.As(rErr, &frobErr) {
			resource.Defrob(frobErr.defrobTag)
		}
		err = rErr
	}()

	resource.Frob(input)
	return nil
}
