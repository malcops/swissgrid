package swissmap

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBern(t *testing.T) {

    bern := LV95Pos{2600000, 1200000}
    expected_lat := 46.9481
    expected_lon := 7.4474

    p := LV95ToWGS84(bern)

    assert.InDelta(t, expected_lat, p.latitude, 0.08)
    assert.InDelta(t, expected_lon, p.longitude, 0.12)
}

func TestGeneva(t *testing.T) {

    geneva := LV95Pos{2499944, 1117674}
    expected_lat := 46.204391
    expected_lon := 6.143158

    p := LV95ToWGS84(geneva)

    assert.InDelta(t, expected_lat, p.latitude, 0.08)
    assert.InDelta(t, expected_lon, p.longitude, 0.12)
}

func TestPoint1(t *testing.T){

    e := 2610735.0
    n := 1183359.0
    expected_lat := 46.801306998
    expected_lon := 7.579259163

    p := LV95ToWGS84(LV95Pos{e, n})

    assert.InDelta(t, expected_lat, p.latitude, 0.08)
    assert.InDelta(t, expected_lon, p.longitude, 0.12)
}

func TestPoint2(t  *testing.T) {

    e := 2700000.0
    n := 1100000.0
    expected_lat := 46.044130339
    expected_lon := 8.730497076

    p := LV95ToWGS84(LV95Pos{e, n})

    assert.InDelta(t, expected_lat, p.latitude, 0.08)
    assert.InDelta(t, expected_lon, p.longitude, 0.12)
}

func TestGivenExample(t *testing.T) {
	given := WGS84Pos{46.044131, 8.730497}
	p := WGS84ToLV95(given)
	assert.InDelta(t, 2700000, p.easting, 1)
    assert.InDelta(t, 1100000, p.northing, 1)
}
