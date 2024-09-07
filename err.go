package main

import (
	"fmt"
)

var (
	ErrPlayerExists = &PlayerExistsError{}
	ErrNoSuchUser   = &NoSuchUserError{}
	ErrNoSuchPlayer = &NoSuchPlayerError{}
)

type PlayerExistsError struct {
	Login string
}

func (e *PlayerExistsError) Error() string {
	return fmt.Sprintf("player with login %s already exists", e.Login)
}

type NoSuchUserError struct {
	Login string
}

func (e *NoSuchUserError) Error() string {
	return fmt.Sprintf("cannot register player with login %s. No such user on course", e.Login)
}

type NoSuchPlayerError struct {
	Login string
}

func (e *NoSuchPlayerError) Error() string {
	return fmt.Sprintf("player not found: %s", e.Login)
}
