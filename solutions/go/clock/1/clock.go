package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	Hours   int
	Minutes int
}

func New(h, m int) Clock {
	minutes := 60*h + m
	if minutes < 0 {
		minutes = 1440 - (-minutes)%1440
	}
	return Clock{(minutes / 60) % 24, minutes % 60}
}

func (c Clock) Add(m int) Clock {
	return New(c.Hours, c.Minutes+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.Hours, c.Minutes-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hours, c.Minutes)
}
