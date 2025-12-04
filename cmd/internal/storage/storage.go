package storage

import "github.com/lib/pq"

const (
	CodeUniqueViolation pq.ErrorCode = "23505"
)

type ErrURLExists struct {}

func (ErrURLExists) Error() string  {
	return "url alias already exists"
}

type ErrURLNotFound struct {}

func (ErrURLNotFound) Error() string  {
	return "url alias not found"
}