package array

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("From inputs %v I Wanted %d but recieved %d", numbers, want, got)
		}

		fmt.Printf("From inputs %v I got %d ", numbers, got)
		fmt.Println("")
	})
}

func TestSumAll(t *testing.T) {

	t.Run("2 slices 1-5", func(t *testing.T) {
		numbers1 := []int{1, 2, 3, 4, 5}
		numbers2 := []int{1, 2, 3, 4, 5}
		got := SumAll(numbers1, numbers2)
		want := 30

		if got != want {
			t.Errorf("From inputs %v and %v I wanted %d but recieved %d", numbers1, numbers2, want, got)
		}

		fmt.Printf("From inputs %v and %v I got %d ", numbers1, numbers2, got)
		fmt.Println("")
	})
}

func TestSumReflect(t *testing.T) {

	t.Run("2 slices, 1,2 and 1,5", func(t *testing.T) {
		got := SumReflect([]int{1, 2}, []int{1, 5})
		want := []int{3, 6}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Wanted %v, got %v", got, want)
		}

		if reflect.DeepEqual(got, want) {
			fmt.Printf("it worked! I wanted %v and got %v", got, want)
			fmt.Println()
		}
	})
}
