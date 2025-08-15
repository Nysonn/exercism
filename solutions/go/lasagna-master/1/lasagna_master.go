package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, averagePreparationTimePerLayer int) int {
	numberOfLayers := len(layers)

	const defaultPreparationTimePerLayer = 2

	if averagePreparationTimePerLayer == 0 {
		averagePreparationTimePerLayer = defaultPreparationTimePerLayer
	}

	totalPreparationTime := numberOfLayers * averagePreparationTimePerLayer
	return totalPreparationTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	amountOfNoodles := 0
	amountOfSauce := 0.0

	for _, typeOfLayer := range layers {
		switch typeOfLayer {
		case "noodles":
			amountOfNoodles += 50
		case "sauce":
			amountOfSauce += 0.2
		}
	}

	return amountOfNoodles, amountOfSauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(myFriendRecipe, myRecipe []string) {

	secretIngredient := myFriendRecipe[len(myFriendRecipe)-1]

	myRecipe[len(myRecipe)-1] = secretIngredient
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amountsNeededForTwoPortions []float64, numberOfPortions int) []float64 {
	scalingFactor := float64(numberOfPortions) / 2.0

	scaledQuantities := make([]float64, len(amountsNeededForTwoPortions))

	for i, quantity := range amountsNeededForTwoPortions {
		scaledQuantities[i] = quantity * scalingFactor
	}

	return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
