package units_test

import (
	"fmt"
	"testing"

	"github.com/neilwhitlow/rpgcore/pkg/units"
)

type float64TestDefinition struct {
	Input    float64
	Expected string
}

type intTestDefinition struct {
	Input    int
	Expected string
}

type meterTestDefinition struct {
	Feet     int
	Inches   int
	Expected string
}

func TestGetKilograms(t *testing.T) {
	// expected results from calculator conversions
	tests := map[string]float64TestDefinition{
		"positive5": {
			Input:    5,
			Expected: "2.26796185",
		},
		"positive1": {
			Input:    1,
			Expected: "0.45359237",
		},
		"positive2": {
			Input:    2,
			Expected: "0.90718474",
		},
		"positive99": {
			Input:    99,
			Expected: "44.90564463",
		},
		"negative": {
			Input:    -3,
			Expected: "-1.36077711",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := fmt.Sprintf("%.8f", units.GetKilograms(test.Input))
			if result != test.Expected {
				t.Errorf("GetKilograms(%f); want %v - got %v", test.Input, test.Expected, result)
			}
		})
	}
}

func TestGetPounds(t *testing.T) {
	// expected results from calculator conversions
	tests := map[string]float64TestDefinition{
		"positive5": {
			Input:    5,
			Expected: "11.02311311",
		},
		"positive1": {
			Input:    1,
			Expected: "2.20462262",
		},
		"positive2": {
			Input:    2,
			Expected: "4.40924524",
		},
		"positive99": {
			Input:    99,
			Expected: "218.25763956",
		},
		"negative": {
			Input:    -3,
			Expected: "-6.61386787",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := fmt.Sprintf("%.8f", units.GetPounds(test.Input))
			if result != test.Expected {
				t.Errorf("GetPounds(%f); want %v - got %v", test.Input, test.Expected, result)
			}
		})
	}
}

func TestGetFeet(t *testing.T) {
	// expected results from calculator conversions
	tests := map[string]float64TestDefinition{
		"positive5": {
			Input:    5,
			Expected: "16.40419948",
		},
		"positive1": {
			Input:    1,
			Expected: "3.28083990",
		},
		"positive2": {
			Input:    2,
			Expected: "6.56167979",
		},
		"positive99": {
			Input:    99,
			Expected: "324.80314961",
		},
		"negative": {
			Input:    -3,
			Expected: "-9.84251969",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := fmt.Sprintf("%.8f", units.GetFeet(test.Input))
			if result != test.Expected {
				t.Errorf("GetFeet(%f); want %v - got %v", test.Input, test.Expected, result)
			}
		})
	}
}

func TestGetMeters(t *testing.T) {
	// expected results from calculator conversions
	tests := map[string]meterTestDefinition{
		"positive5": {
			Feet:     5,
			Inches:   0,
			Expected: "1.52400000",
		},
		"positive13inches": {
			Feet:     1,
			Inches:   1,
			Expected: "0.33020000",
		},
		"positive69inches": {
			Feet:     5,
			Inches:   9,
			Expected: "1.75260000",
		},
		"positive1195inches": {
			Feet:     99,
			Inches:   7,
			Expected: "30.35300000",
		},
		"negative37": {
			Feet:     -3,
			Inches:   -1,
			Expected: "-0.93980000",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := fmt.Sprintf("%.8f", units.GetMeters(test.Feet, test.Inches))
			if result != test.Expected {
				t.Errorf("GetMeters(%d, %d); want %v - got %v", test.Feet, test.Inches, test.Expected, result)
			}
		})
	}
}

func TestGetCelsius(t *testing.T) {
	// expected results from calculator conversions
	tests := map[string]intTestDefinition{
		"positive5": {
			Input:    5,
			Expected: "-15.00",
		},
		"positive32": {
			Input:    32,
			Expected: "0.00",
		},
		"positive100": {
			Input:    100,
			Expected: "37.78",
		},
		"negative12": {
			Input:    -12,
			Expected: "-24.44",
		},
		"positive99": {
			Input:    99,
			Expected: "37.22",
		},
		"negative3": {
			Input:    -3,
			Expected: "-19.44",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := fmt.Sprintf("%.2f", units.GetCelsius(test.Input))
			if result != test.Expected {
				t.Errorf("GetCelsius(%d); want %v - got %v", test.Input, test.Expected, result)
			}
		})
	}
}

func TestGetFahrenheit(t *testing.T) {
	// expected results from calculator conversions
	tests := map[string]intTestDefinition{
		"zero": {
			Input:    0,
			Expected: "32.00",
		},
		"positive32": {
			Input:    32,
			Expected: "89.60",
		},
		"positive100": {
			Input:    100,
			Expected: "212.00",
		},
		"negative32": {
			Input:    -32,
			Expected: "-25.60",
		},
		"positive99": {
			Input:    99,
			Expected: "210.20",
		},
		"positive1": {
			Input:    1,
			Expected: "33.80",
		},
		"negative3": {
			Input:    -3,
			Expected: "26.60",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := fmt.Sprintf("%.2f", units.GetFahrenheit(test.Input))
			if result != test.Expected {
				t.Errorf("GetFahrenheit(%d); want %v - got %v", test.Input, test.Expected, result)
			}
		})
	}
}

func TestGetKilometers(t *testing.T) {
	// expected results from calculator conversions
	tests := map[string]float64TestDefinition{
		"positive5": {
			Input:    5,
			Expected: "8.04672000",
		},
		"positive1": {
			Input:    1,
			Expected: "1.60934400",
		},
		"positive2": {
			Input:    2,
			Expected: "3.21868800",
		},
		"positive99": {
			Input:    99,
			Expected: "159.32505600",
		},
		"negative": {
			Input:    -3,
			Expected: "-4.82803200",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := fmt.Sprintf("%.8f", units.GetKilometers(test.Input))
			if result != test.Expected {
				t.Errorf("GetKilometers(%f); want %v - got %v", test.Input, test.Expected, result)
			}
		})
	}
}

func TestGetMiles(t *testing.T) {
	// expected results from calculator conversions
	tests := map[string]float64TestDefinition{
		"positive5": {
			Input:    5,
			Expected: "3.10685596",
		},
		"positive1": {
			Input:    1,
			Expected: "0.62137119",
		},
		"positive2": {
			Input:    2,
			Expected: "1.24274238",
		},
		"positive99": {
			Input:    99,
			Expected: "61.51574803",
		},
		"negative": {
			Input:    -7,
			Expected: "-4.34959835",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := fmt.Sprintf("%.8f", units.GetMiles(test.Input))
			if result != test.Expected {
				t.Errorf("GetMiles(%f); want %v - got %v", test.Input, test.Expected, result)
			}
		})
	}
}
