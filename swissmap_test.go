package swissmap

import (
	"testing"
	"fmt"
)


func TestBern(t *testing.T) {
	// 46.9481°, 7.4474°
    bern := LV95Pos{2600000,1200000}
    fmt.Println(LV95ToWGS84(bern))
    fmt.Println(WGS84Pos{7.4474, 46.9481})
}

// Numerical example
// given: E = 2 700 000 m N = 1 100 000 m hCH = 600 m
// → y' = 0.1 x' = -0.1
// → λ' = 3.14297976 φ' = 16.57588564 hWGS = 650.55 m
// → λ = 8° 43' 49.80" φ = 46° 02' 38.86"
// Reference: λ = 8° 43' 49.79" φ = 46° 02' 38.87" h = 650.60 m
// The precision of the approximate formulas is better than 0.12" in longitude, 0.08" in latitude and 0.5
// metres in height everywhere in Switzerland.
