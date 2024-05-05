package lasagna

func PreparationTime(layers []string, TimePerLayer int) int {
	if TimePerLayer == 0 {
		return len(layers) * 2
	}
	return len(layers) * TimePerLayer
}

func Quantities(layers []string) (noodles int, sauce float64) {
	gramCount := 0
	litersCount := 0
	for _, layer := range layers {
		switch layer {
		case "noodles":
			gramCount++
			noodles = gramCount * 50
		case "sauce":
			litersCount++
			sauce = float64(litersCount) * 0.2
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendsIngredients []string, myIngredients []string) {
	if len(friendsIngredients) > 0 && len(myIngredients) > 0 {
		(myIngredients)[len(myIngredients)-1] = (friendsIngredients)[len(friendsIngredients)-1]
	}
}

func ScaleRecipe(quantities []float64, portions int) (scaled []float64) {
	for _, portion := range quantities {
		scaled = append(scaled, float64(portions)/2.0*portion)
	}
	return scaled
}
