package domain

import (
	"errors"
	"math"
)

var EmptyFlightPathMapError = errors.New("it's not possible to calculate with an empty map")

type FlightPathService interface {
	CalculateFlightPath(paths *[]*Flight) (orderedPaths *[]*Flight, err error)
}

type FlightPathServiceImpl struct {
}

func NewFlightPathServiceImpl() FlightPathService {
	return &FlightPathServiceImpl{}
}

func (f FlightPathServiceImpl) CalculateFlightPath(paths *[]*Flight) (orderedPaths *[]*Flight, err error) {
	if len(*paths) == 0 {
		return orderedPaths, EmptyFlightPathMapError
	}
	if len(*paths) == 1 {
		return paths, nil
	}
	var result []*Flight
	var next *Flight
	for hasElement := true; hasElement; hasElement = len(*paths) > 0 {
		index := f.findNextOrFirstFlight(paths, next)
		//*next = *((*paths)[index])
		next, paths = getAndRemove(paths, index)
		result = append(result, next)
	}
	return &result, nil
}

func (f FlightPathServiceImpl) findNextOrFirstFlight(paths *[]*Flight, previous *Flight) int {
out:
	for i, path := range *paths {
		if previous == nil {
			for j := i + 1; j < len(*paths); j++ {
				if path.Origin == (*paths)[j].Destination {
					continue out
				}
				if j == len(*paths)-1 {
					return i
				}
			}
		} else if path.Origin == previous.Destination {
			return i
		}
	}
	return math.MaxInt
}

func getAndRemove(paths *[]*Flight, index int) (*Flight, *[]*Flight) {
	flight := (*paths)[index]
	result := append((*paths)[:index], (*paths)[index+1:]...)
	return flight, &result
}

func (f FlightPathServiceImpl) findFirstFlight(paths []Flight) (flight Flight, index int) {
out:
	for i, path := range paths {
		for j := i; j < len(paths); j++ {
			if path.Origin == paths[j].Destination {
				continue
			}
			if j == len(paths)-1 {
				flight = paths[j]
				index = j
				break out
			}
		}
	}
	return flight, index
}

func (f FlightPathServiceImpl) findNextFlight(paths []Flight, previous *Flight) (flight Flight, index int) {
out:
	for i, path := range paths {
		if path.Origin == previous.Destination {
			flight = path
			index = i
			break out
		}
	}
	return flight, index
}
