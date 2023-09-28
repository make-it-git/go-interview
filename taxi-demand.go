package main

type Point struct {
	Lat float32
	Lng float32
}

type Location struct {
	UpperLeft  Point
	LowerRight Point
}
