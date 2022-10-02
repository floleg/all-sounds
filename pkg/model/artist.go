package model

type Artist struct {
	Entity
	Name   string
	Tracks []Track
}
