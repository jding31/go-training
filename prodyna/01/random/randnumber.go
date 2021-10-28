package random

import (
	ff "fmt"
	"math/rand"
)

func init() {
	ff.Printf("Initial 3 Random Numbers: %d %d %d\n", rand.Intn(1000), rand.Intn(1000), rand.Intn(1000))
}
