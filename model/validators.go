package model

import (
	"errors"
)

// Validate validates a redirect entry
func (r *Redirect) Validate() error {
	if r.Code < 301 || r.Code > 308 || r.Code == 306 || r.Code == 303 {
		return errors.New(ErrInvalidHTTPCode)
	}
	return nil

}
