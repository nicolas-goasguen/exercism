package lasagnamaster

const (
	noodles            = "noodles"
	sauce              = "sauce"
	noodleBaseQuantity = 50
	sauceBaseQuantity  = 0.2
)

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		timePerLayer = 2
	}
	return len(layers) * timePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodleQuantity := 0
	sauceQuantity := 0.0
	for _, l := range layers {
		switch l {
		case noodles:
			noodleQuantity += noodleBaseQuantity
		case sauce:
			sauceQuantity += sauceBaseQuantity
		default:
			continue
		}
	}
	return noodleQuantity, sauceQuantity
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, mylist []string) []string {
	return append(mylist[:len(mylist)-1], friendsList[len(friendsList)-1])
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaled := make([]float64, len(quantities))
	scaleFloat := float64(portions)
	for i := range quantities {
		scaled[i] = (quantities[i] / 2) * scaleFloat
	}
	return scaled
}
