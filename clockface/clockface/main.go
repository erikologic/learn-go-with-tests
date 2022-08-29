package main

import (
	"github.com/mrenrich84/learn-go-with-tests/clockface"
	"os"
	"time"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}