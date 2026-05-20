package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPrepTime int) int {
	if avgPrepTime == 0 {
		avgPrepTime = 2
	}
	return len(layers)*avgPrepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64){
	noodleCount := 0
	sauceCount := 0
	for _, layer := range layers {
		if layer == "noodles" {
			noodleCount++
		}
		if layer == "sauce"{
			sauceCount++
		}
	}
	return int(noodleCount*50), float64(sauceCount)*0.2
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string){
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portion int) []float64 {
	var scaledQty []float64
	for _, val := range quantities {
		scaledQty = append(scaledQty, (val/2)*float64(portion))
	}
	return scaledQty
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
