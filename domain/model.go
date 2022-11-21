package domain

type Flight struct {
	Origin      string
	Destination string
}

func NewFlight(origin, destination string) *Flight {
	return &Flight{
		Origin:      origin,
		Destination: destination,
	}
}
