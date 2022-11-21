package domain

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

type TestCase struct {
	name      string
	paths     *[]*Flight
	assertion func(t *testing.T, input *[]*Flight, result *[]*Flight, err error)
}

func (testCase *TestCase) run(t *testing.T) {
	target := NewFlightPathServiceImpl()
	t.Run(testCase.name, func(t *testing.T) {
		result, err := target.CalculateFlightPath(testCase.paths)
		testCase.assertion(t, testCase.paths, result, err)
	})
}

func TestSingleFlight(t *testing.T) {
	singleFlight := TestCase{
		name:  "single flight",
		paths: &[]*Flight{NewFlight("SFO", "EWR")},
		assertion: func(t *testing.T, input *[]*Flight, result *[]*Flight, err error) {
			require.NotNil(t, result)
			require.Nil(t, err)
			require.Equalf(t, result, input, "input and output params should be equal")
		},
	}
	singleFlight.run(t)
}

func TestEmptyFlight(t *testing.T) {
	emptyFlightList := TestCase{
		name:  "empty flight list",
		paths: nil,
		assertion: func(t *testing.T, input *[]*Flight, result *[]*Flight, err error) {
			require.NotNil(t, err)
			require.Nil(t, result)
			require.Truef(t, errors.Is(err, EmptyFlightPathMapError), "Wrong error message")
		},
	}
	emptyFlightList.run(t)
}

func TestDoublePathScenario(t *testing.T) {
	doublePathFlight := TestCase{
		name: "double path flight",
		paths: &[]*Flight{
			NewFlight("ATL", "EWR"),
			NewFlight("SFO", "ATL"),
		},
		assertion: func(t *testing.T, input *[]*Flight, result *[]*Flight, err error) {
			require.NotNil(t, result)
			require.Nil(t, err)
			require.Equalf(t, (*result)[0].Origin, "SFO", "First origin should be SFO")
			require.Equalf(t, (*result)[1].Destination, "EWR", "Last destination should be EWR")
		},
	}
	doublePathFlight.run(t)
}

func TestCompletePathScenario(t *testing.T) {
	completeScenario := TestCase{
		name: "complete path flight",
		paths: &[]*Flight{
			NewFlight("IND", "EWR"),
			NewFlight("SFO", "ATL"),
			NewFlight("GSO", "IND"),
			NewFlight("ATL", "GSO"),
		},
		assertion: func(t *testing.T, input *[]*Flight, result *[]*Flight, err error) {
			require.NotNil(t, result)
			require.Nil(t, err)
			require.Equalf(t, (*result)[0].Origin, "SFO", "First origin should be SFO")
			require.Equalf(t, (*result)[3].Destination, "EWR", "Last destination should be EWR")
		},
	}
	completeScenario.run(t)
}
