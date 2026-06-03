package thefarm

import (
    "fmt"
	"errors"
)

// TODO: define the 'DivideFood' function
func DivideFood (fc FodderCalculator, cows int) (float64, error) {
    am, err := fc.FodderAmount(cows)
    am = am / float64(cows)
    if err != nil {
        return 0, err
    }
    factor , err1 := fc.FatteningFactor()
    if err1 != nil {
        return 0, err1
    }

    return am*factor, nil
}
// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood (fc FodderCalculator, cows int) (float64, error) {
    if cows <= 0 {
        return 0, errors.New("invalid number of cows")
    }
    res, err := DivideFood(fc, cows);
    if err != nil {
        return 0, err
    }
    return res, nil
}

// TODO: define the 'ValidateNumberOfCows' function

type CustomError struct {
    numCows int
    message string
}

func (e *CustomError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", e.numCows, e.message)
}
func ValidateNumberOfCows(cows int) error {
    if cows < 0 {
        return &CustomError{
            numCows : cows,
            message: "there are no negative cows",
        }
    }
    if cows == 0{
        return &CustomError{
            numCows : cows,
            message: "no cows don't need food",
        }
    }
    return nil
}


// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
