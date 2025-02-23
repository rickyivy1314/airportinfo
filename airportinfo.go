package airportinfo

// IsValidAirport checks if the given airport code is valid
func IsValidAirport(code string) bool {
	return NorthAmericaAirports[code] ||
		EuropeAirports[code] ||
		AsiaAirports[code]
}
