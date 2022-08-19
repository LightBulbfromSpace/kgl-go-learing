package main

import (
	clock "github.com/LightBulbfromSpace/kld-go-learning/math/v3"
	"os"
	"time"
)

func main() {
	t := time.Now()
	clock.SVGWriter(os.Stdout, t)
}
