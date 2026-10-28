package swissmap

import ("math")

type WGS84Pos struct {
	latitude float64
	longitude float64
}

type LV95Pos struct {
	easting float64
	northing float64
}

func WGS84ToLV95(pos WGS84Pos) LV95Pos {
	return LV95Pos{0,0}
}

func LV95ToWGS84(pos LV95Pos) WGS84Pos {
	e_prime := (pos.easting - 2600000)/1000000
	n_prime := (pos.northing - 1200000)/1000000

	lambda_prime := 2.6779094 +
	                4.728982 * e_prime +
				    0.791484 * e_prime * n_prime +
					0.1306 * e_prime * math.Pow(n_prime, 2) -
	                0.0436 * math.Pow(n_prime, 3)

	theta_prime := 16.9023892 +
				   3.238272 * n_prime -
				   0.270978 * math.Pow(e_prime, 2) -
				   0.002528 * math.Pow(n_prime, 2) -
				   0.0447 * math.Pow(e_prime, 2) * n_prime -
				   0.0140 * math.Pow(n_prime, 3)

	long := lambda_prime * 100 / 36
	lat := theta_prime * 100 / 36

	return WGS84Pos{lat, long}
}
