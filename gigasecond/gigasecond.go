package gigasecond

// import path for the time package from the standard library
import "time"

const testVersion = 4

const giga = 1000000000 * time.Second

// AddGigasecond : Calculate the moment when someone has lived for 10^9 seconds
func AddGigasecond(date time.Time) time.Time {
	return date.Add(giga)
}
