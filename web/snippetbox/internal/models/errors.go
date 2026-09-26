package models

import "errors"

var (
	ErrNoRecord            = errors.New("models: No Matching Record Found")
	ErrInvalidCredientials = errors.New("models: invalid credientials")
	ErrDuplicateEmail      = errors.New("models: duplicate email")
)
