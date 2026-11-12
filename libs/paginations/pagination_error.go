package paginations

import "errors"

// ErrMaxPage is returned when requested page exceeds total pages.
var ErrMaxPage = errors.New("chosen page more than total page")

// ErrPageInvalid is returned when page is less than 1.
var ErrPageInvalid = errors.New("page must be greater than 0")