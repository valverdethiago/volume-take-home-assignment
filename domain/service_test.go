package domain

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFlight(t *testing.T) {
	t.Run("single flight", func(t *testing.T) {
		paths := []*Flight{NewFlight("SFO", "EWR")}
		result, err := NewFlightPathServiceImpl().CalculateFlightPath(&paths)
		require.NotNil(t, result)
		require.Nil(t, err)
		require.Equalf(t, result, paths, "input and output params should be equal")
	})
}

func TestEmptyFlight(t *testing.T) {
	t.Run("empty flight", func(t *testing.T) {
		result, err := NewFlightPathServiceImpl().CalculateFlightPath(nil)
		require.NotNil(t, err)
		require.Nil(t, result)
		require.Truef(t, errors.Is(err, EmptyFlightPathMapError), "Wrong error message")
	})
}

func TestDoublePathScenario(t *testing.T) {
	t.Run("double path flight", func(t *testing.T) {
		paths := []*Flight{
			NewFlight("ATL", "EWR"),
			NewFlight("SFO", "ATL"),
		}
		result, err := NewFlightPathServiceImpl().CalculateFlightPath(&paths)
		require.NotNil(t, result)
		require.Nil(t, err)
		require.Equalf(t, (*result)[0].Origin, "SFO", "First origin should be SFO")
		require.Equalf(t, (*result)[1].Destination, "EWR", "Last destination should be EWR")
	})
}

func TestCompletePathScenario(t *testing.T) {
	t.Run("complete path flight", func(t *testing.T) {
		paths := []*Flight{
			NewFlight("IND", "EWR"),
			NewFlight("SFO", "ATL"),
			NewFlight("GSO", "IND"),
			NewFlight("ATL", "GSO"),
		}
		result, err := NewFlightPathServiceImpl().CalculateFlightPath(&paths)
		require.NotNil(t, result)
		require.Nil(t, err)
		require.Equalf(t, (*result)[0].Origin, "SFO", "First origin should be SFO")
		require.Equalf(t, (*result)[3].Destination, "EWR", "Last destination should be EWR")
	})
}
