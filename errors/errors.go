package errors

import "errors"

// MissingRequiredArgumentErr missing required argument error
var MissingRequiredArgumentErr = errors.New("missing required argument")

// FlagContentValidateErr flag content validate error
var FlagContentValidateErr = errors.New("flag content validate error")

// FlagDateValidateErr flag date validate error
var FlagDateValidateErr = errors.New("flag date validate error")

// FlagLevelValidateErr flag level validate error
var FlagLevelValidateErr = errors.New("flag level validate error")

// FlagStatusValidateErr flag status validate error
var FlagStatusValidateErr = errors.New("flag status validate error")

// TodoNotExistsErr todo not exists
var TodoNotExistsErr = errors.New("todo not exists")
