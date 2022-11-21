package domain

import (
	"errors"
	"math"
)

var (
	EmptyFlightPathMapError        = errors.New("it's not possible to calculate with an empty map")
	DisconnectedFlightPathMapError = errors.New("it's not possible to calculate with an empty map")
)

type FlightPathService interface {
	CalculateFlightPath(paths *[]*Flight) (orderedPaths *[]*Flight, err error)
}

type FlightPathServiceImpl struct {
}

func NewFlightPathServiceImpl() FlightPathService {
	return &FlightPathServiceImpl{}
}

func (f FlightPathServiceImpl) CalculateFlightPath(paths *[]*Flight) (orderedPaths *[]*Flight, err error) {
	if paths == nil || len(*paths) == 0 {
		return orderedPaths, EmptyFlightPathMapError
	}
	if len(*paths) == 1 {
		return paths, nil
	}
	var result []*Flight
	var next *Flight
	for hasElement := true; hasElement; hasElement = len(*paths) > 0 {
		index := f.findNextOrFirstFlight(paths, next)
		if index == math.MaxInt {
			return nil, DisconnectedFlightPathMapError
		}
		next, paths = getAndRemove(paths, index)
		result = append(result, next)
	}
	return &result, nil
}

func (f FlightPathServiceImpl) findNextOrFirstFlight(paths *[]*Flight, previous *Flight) (result int) {
	result = math.MaxInt
out:
	for i, path := range *paths {
		if previous == nil {
			for j := i + 1; j < len(*paths); j++ {
				if path.Origin == (*paths)[j].Destination {
					continue out
				}
			}
			result = i
			break out
		} else if path.Origin == previous.Destination {
			result = i
			break out
		}
	}
	return result
}

func getAndRemove(paths *[]*Flight, index int) (*Flight, *[]*Flight) {
	flight := (*paths)[index]
	result := append((*paths)[:index], (*paths)[index+1:]...)
	return flight, &result
}
