package clockface

import (
	"math"
	"time"
)

// A Point represents a two dimensional Cartesian coordinate.
type Point struct {
	X float64
	Y float64
}

const (
	secsInMin = 60
	secsInHours = 60 * 60
	secsInHalfDay = 12 * 60 * 60
)

var fullCircle = 2 * math.Pi
func secondsInRadians(t time.Time) float64 {
	secs := float64(t.Second())
	return fullCircle / secsInMin * secs
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func minutesInRadians(t time.Time) float64 {
	secs := float64(60 * t.Minute() + t.Second())
	return fullCircle / secsInHours * secs
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}

func hoursInRadians(t time.Time) float64 {
	secs := float64(t.Hour() % 12 * 60 * 60 + 60 * t.Minute() + t.Second())
	return fullCircle / secsInHalfDay * secs
}

func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}