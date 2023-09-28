package main

import (
	"context"
)

type Point struct {
	Lat float32
	Lng float32
}

type Location struct {
	UpperLeft  Point
	LowerRight Point
}

var locations = []Location{
	{
		UpperLeft: Point{
			Lat: 70,
			Lng: 70,
		},
		LowerRight: Point{
			Lat: 69,
			Lng: 69,
		},
	},
	{
		UpperLeft: Point{
			Lat: 69,
			Lng: 69,
		},
		LowerRight: Point{
			Lat: 68,
			Lng: 68,
		},
	},
}

func (p Point) WithinBounds(upperLeft Point, lowerRight Point) bool {
	// naive (not correct) realization
	if p.Lat > upperLeft.Lat {
		return false
	}
	if p.Lng > upperLeft.Lng {
		return false
	}
	if p.Lat < lowerRight.Lat {
		return false
	}
	if p.Lng < lowerRight.Lng {
		return false
	}
	return true
}

type Consumer struct {
}

func (c *Consumer) ConsumeRider(ctx context.Context, points []Point) {
	for _, p := range points {
		for _, l := range locations {
			if p.WithinBounds(l.UpperLeft, l.LowerRight) {
				// point belongs to this location
			}
		}
		// what if we can not find location
	}
}

func (c *Consumer) ConsumeDriver(ctx context.Context, points []Point) {
	for _, p := range points {
		for _, l := range locations {
			if p.WithinBounds(l.UpperLeft, l.LowerRight) {
				// point belongs to this location
			}
		}
		// what if we can not find location
	}
}

func (c *Consumer) GetDemand(ctx context.Context, point Point) float64 {
	riders := 123
	drivers := 50
	return float64(riders) / float64(drivers)
}
