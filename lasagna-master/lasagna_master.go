package lasagna

// TODO: define the 'PreparationTime()' function

// TODO: define the 'Quantities()' function

// TODO: define the 'AddSecretIngredient()' function

// TODO: define the 'ScaleRecipe()' function

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

const (
	defaultPreparationTimePerLayer = 2
	noodlesInGmPerNoodleLayer      = 50
	sauceInLPerSauceLayer          = 0.2
)

func PreparationTime(layers []string, avgPreparationTimePerLayer int) int {
	if avgPreparationTimePerLayer == 0 {
		avgPreparationTimePerLayer = defaultPreparationTimePerLayer
	}
	totalTime := len(layers) * avgPreparationTimePerLayer
	return totalTime
}
func Quantities(layers []string) (noodles int, sauce float64) {
	var noodleLayers, sauceLayers int
	for _, l := range layers {
		if l == "noodles" {
			noodleLayers++
		} else if l == "sauce" {
			sauceLayers++
		}
	}
	noodles = noodlesInGmPerNoodleLayer * noodleLayers
	sauce = sauceInLPerSauceLayer * float64(sauceLayers)
	return noodles, sauce
}
func AddSecretIngredient(friendsList []string, myList []string) {
	secretIngredient := friendsList[len(friendsList)-1]
	myList[len(myList)-1] = secretIngredient
}
func ScaleRecipe(quantitiesForTwoPortions []float64, portions int) (scaledQuantities []float64) {
	var scale float64 = float64(portions) / 2
	for _, q := range quantitiesForTwoPortions {
		scaledQuantities = append(scaledQuantities, q*scale)
	}
	return scaledQuantities
}
