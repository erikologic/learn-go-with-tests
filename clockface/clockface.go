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
	bySeconds = 60
	byMinutes = 60 * 60
	byHours = 12 * 60 * 60
	fullCircle = 2 * math.Pi
)

func secondHandPoint(t time.Time) Point {
	return angleToPoint(toRadians(bySeconds, secsSinceMidnight(t)))
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(toRadians(byMinutes, secsSinceMidnight(t)))

}

func hourHandPoint(t time.Time) Point {
	return angleToPoint(toRadians(byHours, secsSinceMidnight(t)))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}

func toRadians(base, secsSinceMidnight int) float64 {
	return fullCircle / float64(base) * float64(secsSinceMidnight)
}

func secsSinceMidnight(t time.Time) int {
    year, month, day := t.Date()
    t2 := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
    return int(t.Sub(t2).Seconds())
}