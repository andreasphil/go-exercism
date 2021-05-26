// Constants and functions for working with gigasecpnds. A gigasecond is 10^9
// seconds.
package gigasecond

import "time"

// Gigasecond as duration
const Gigasecond = time.Second * 1_000_000_000

// Given a time, adds one gigasecond to the time.
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(Gigasecond)
	return t
}
