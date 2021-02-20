package color

import "errors"

type MalformedColor error

var ErrMalformedColor MalformedColor = errors.New("color: malformed value")
