package main

import "math/rand"

type RandomNumbersGenerator interface {
	Random() float64
}

type RandomNumbersGeneratorImpl struct {
}

type MockRandomNumbersGenerator struct {
	value float64
}

func (r *RandomNumbersGeneratorImpl) Random() float64 {
	return rand.Float64()
}

func (r *MockRandomNumbersGenerator) Random() float64 {
	return r.value
}
