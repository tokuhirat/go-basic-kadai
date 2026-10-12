package main

type ValidationError struct {
	Message string
	Field   string
	Value   float64
}
