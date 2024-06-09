package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	switch kind {
	case "car", "truck":
		return true
	default:
		return false
	}
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var firstVehicle string
	if option1 < option2 {
		firstVehicle = option1
	} else {
		firstVehicle = option2
	}
	return firstVehicle + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var percentage float64
	if age < 3 {
		percentage = 0.8
	} else if age >= 3 && age < 10 {
		percentage = 0.7

	} else {
		percentage = 0.5
	}
	return percentage * originalPrice
}
