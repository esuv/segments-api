package app

import (
	"fmt"
	"segments-api/internal/config"
	"strings"
)

func Run(configDir string) {
	_ = config.MustLoad(configDir)

	str := "12 345 df"
	strings.TrimSuffix()
	fmt.Println(strings.ReplaceAll(str, " ", ""))

	arr := []int{124, 324, 345, -1, -234, 0}
	fmt.Println(Invert(arr))

}

func Invert(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)

	for i, elem := range arr {
		arr[i] = elem * -1
	}

	return arr
}
func CalculateYears(years int) (result [3]int) {

	var catAge, dogAge int

	for i := 1; i <= years; i++ {
		if i == 1 {
			catAge = 15
			dogAge = 15
		} else if i == 2 {
			catAge = catAge + 9
			dogAge = dogAge + 9
		} else {
			catAge = catAge + 4
			dogAge = dogAge + 5
		}
	}

	return [3]int{years, catAge, dogAge}

}
