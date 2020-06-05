package goatcounter

import (
	"math"

	"github.com/mattn/go-sqlite3"
)

var SQLiteHook = func(c *sqlite3.SQLiteConn) error {
	//return c.RegisterFunc("percent_diff", func(start, final float64) float64 {
	return c.RegisterFunc("percent_diff", func(start, final int) float64 {
		if start == 0 {
			return math.Inf(0)
		}
		return float64(float64((final-start)/start) * 100)
	}, true)
}
