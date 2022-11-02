package model

// Artist data entity
type Artist struct {
	Entity
	Name   string
	Tracks []Track
}
