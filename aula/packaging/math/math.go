package math

import "github.com/google/uuid"

type Math struct {
	A int
	B int
}

func (m Math) Add() int {
	return m.A + m.B
}

func Generate() string {
	return uuid.New().String()
}
