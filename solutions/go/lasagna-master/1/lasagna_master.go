package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepTimePerLayer int) int {
    if prepTimePerLayer == 0 {
        prepTimePerLayer = 2
    }
    return len(layers) * prepTimePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    noodles := 0
    sauce := 0.0
    for i := 0; i < len(layers); i++ {
        if layers[i] == "noodles" {
            noodles += 50
        } else if layers[i] == "sauce" {
            sauce += 0.2
        }
    }
    return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendLayers, myLayers []string) {
    friendLastItem := friendLayers[len(friendLayers) - 1]
    myLayers[len(myLayers) - 1] = friendLastItem
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(twoPortionsSlice []float64, numPortions int) []float64 {
    output := make([]float64, len(twoPortionsSlice))
    copy(output, twoPortionsSlice)

    for i := 0; i < len(twoPortionsSlice); i++ {
        onePortionSize := twoPortionsSlice[i] / 2.0
        output[i] = onePortionSize * float64(numPortions)
    }

    return output
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
