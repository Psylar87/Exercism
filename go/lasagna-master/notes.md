# Lasagna-Master Notes

- `PreparationTime`: Calculates the total preparation time for making lasagna based on the number of layers and the time taken per layer. It takes in two parameters: layers, a slice of strings representing each layer of the lasagna, and TimePerLayer, an integer representing the time taken to prepare each layer. It returns an integer representing the total preparation time.

```go
func PreparationTime(layers []string, TimePerLayer int) int {
	if TimePerLayer == 0 {
		return len(layers) * 2
	}
	return len(layers) * TimePerLayer
}
```

- `Quantities`: Calculates the quantities of noodles and sauce required for making the lasagna. It takes in a slice of strings layers representing each layer of the lasagna. It iterates through each layer, counting the number of "noodles" and "sauce" layers and calculates the total grams of noodles and liters of sauce required. It returns two values: an integer representing the total grams of noodles and a float64 representing the total liters of sauce.

```go
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
```

- `AddSecretIngredient`: This function adds a secret ingredient to the lasagna recipe. It takes in two parameters: friendsIngredients, a slice of strings representing the friend's ingredients, and myIngredients, a slice of strings representing the user's ingredients. It checks if both slices are non-empty and then replaces the last ingredient in myIngredients with the last ingredient in friendsIngredients.

```go
func AddSecretIngredient(friendsIngredients []string, myIngredients []string) {
	if len(friendsIngredients) > 0 && len(myIngredients) > 0 {
		(myIngredients)[len(myIngredients)-1] = (friendsIngredients)[len(friendsIngredients)-1]
	}
}
```

- `ScaleRecipe`: This function scales the quantities of ingredients in the recipe based on the number of portions. It takes in a slice of float64 quantities, representing the quantities of ingredients, and an integer portions representing the desired number of portions. It iterates through each quantity, scales it based on the number of portions, and returns a slice of scaled quantities.

```go
func ScaleRecipe(quantities []float64, portions int) (scaled []float64) {
	for _, portion := range quantities {
		scaled = append(scaled, float64(portions)/2.0*portion)
	}
	return scaled
}
```

The function `ScaleRecipe` iterates over each quantity in the `quantities` slice using a `range` loop. For each iteration, it assigns the current quantity to the variable `portion`. Inside the loop, the function calculates the scaled quantity for the current portion by dividing the desired number of portions (`portions`) by 2.0 and then multiplying it by the current portion quantity. This is the core mathematical operation. The scaled quantity calculated in the previous step is then appended to the `scaled` slice. Finally, once all quantities have been scaled and added to the `scaled` slice, the function returns this slice.

# Math breakdown

First, you count how many servings of lasagna you want to make. Let's call this number "portions". For example, if you want to make lasagna for 8 people, "portions" would be 8. Then, you look at the original recipe, which tells you how much of each ingredient you need for one serving. This is what we call a "portion". Now, we need to figure out how much of each ingredient we need for all the servings we want to make. To do this, we take the original amount of each ingredient for one serving and multiply it by the number of servings we want to make. But wait! Lasagna is a big meal, so we only need half of what we calculated. We divide the total by 2 because we already have one portion in the recipe. So, for each ingredient, we take the original amount, multiply it by the number of servings we want, and then divide by 2. That's how we get the amount of each ingredient we need for the bigger batch of lasagna.
