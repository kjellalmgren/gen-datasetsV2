package randoms

import (
	"math/rand"
	"time"
)

// v4randomNumber
func RandomNumberv4(min float64, max float64) float64 {

	var rn float64
	rn = 0
	for rn == 0 || (rn < min && rn > max) {
		rand.Seed(time.Now().UnixNano())
		rn = min + rand.Float64()*(max-min) // rand.int63(max)
	}
	return rn
}
