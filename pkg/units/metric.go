// Package units supports a limited amount of metric<->U.S. conversions
// to support common sci-fi RPG rules that use the metric system.
// Based on https://www.nist.gov/pml/owm/metric-si/unit-conversion
// ans section 4.4 of the publication (NIST SP 1038 - 2006)

package units

// GetKilograms converts pounds to kilograms
func GetKilograms(pounds float64) float64 {
	return pounds * 0.45359237
}

// GetPounds converst kilograms to pounds
func GetPounds(kilograms float64) float64 {
	return kilograms / 0.45359237
}

// GetMeters converts a measurement expressed in terms of
// feet and inches (6ft 1in, 3ft 0in, etc) together into Meters
func GetMeters(feet int, inches int) float64 {
	return float64(inches+(feet*12)) * .0254
}

// GetFeet converts meters into feet
func GetFeet(meters float64) float64 {
	return (meters / .0254) / 12
}

// GetCelsius converts fahrenheit degrees to celsius degrees
func GetCelsius(fahrenheit int) float64 {
	return float64(fahrenheit-32) * 0.5555555555555556
}

// GetFahrenheit converts celsius degrees into fahrenheit degrees
func GetFahrenheit(celsius int) float64 {
	return (float64(celsius) / 0.5555555555555556) + float64(32)
}

// GetKilometers converst miles into kilometers
func GetKilometers(miles float64) float64 {
	return miles * 1.609344
}

// GetMiles converts kilometers into miles
func GetMiles(kilometers float64) float64 {
	return kilometers / 1.609344
}
