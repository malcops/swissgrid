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

	lat := pos.latitude * 3600
	lon := pos.longitude * 3600

	theta_prime := (lat - 169028.66)/10000
	lambda_prime :=  (lon - 26782.5)/10000

	easting := 2600072.37 +
	    211455.93 * lambda_prime -
	    10938.51 * lambda_prime * theta_prime -
	    0.36 * lambda_prime * math.Pow(theta_prime, 2) -
	    44.54 * math.Pow(lambda_prime, 3)

	northing := 1200147.07 +
		308807.95 * theta_prime +
		3745.25 * math.Pow(lambda_prime, 2) +
		76.63 * math.Pow(theta_prime, 2) -
		194.56 * math.Pow(lambda_prime, 2) * theta_prime +
		119.79 * math.Pow(theta_prime, 3)

	return LV95Pos{easting, northing}
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
