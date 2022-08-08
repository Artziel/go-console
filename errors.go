package GoConsole

import "errors"

var ErrWrongFalgsType = errors.New("command flags must be a pointer to an structure")

var ErrSubCommandUnespecified = errors.New("you must pass a sub-command")
var ErrSubCommandNotFound = errors.New("sub-command not found")

var ErrTagNoFieldTag error = errors.New("no field tags")
var ErrTagEmptyFieldTag error = errors.New("empty field tag")
var ErrTagMissingType error = errors.New("missing field type")
var ErrTagInvalidType error = errors.New("invalid field tag type")
