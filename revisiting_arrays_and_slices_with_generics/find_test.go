package revisiting_arrays_and_slices

import (
	"strings"
	"testing"
)

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 8, 9, 10}
		firstEvenNumber, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})

		AssertEqual(t, firstEvenNumber, 2)
		AssertTrue(t, found)
	})
	t.Run("no number suitable to criteria", func(t *testing.T) {
		numbers := []int{1, 3, 5, 9}
		firstEvenNumber, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})

		AssertEqual(t, firstEvenNumber, 0)
		AssertFalse(t, found)
	})
	type Person struct {
		Name string
	}
	t.Run("Find the person", func(t *testing.T) {
		people := []Person{
			Person{Name: "Kent Beck"},
			Person{Name: "Martin Fowler"},
			Person{Name: "Chris James"},
		}
		person, found := Find(people, func(p Person) bool {
			return strings.Contains(p.Name, "Chris")
		})
		AssertEqual(t, person, Person{Name: "Chris James"})
		AssertTrue(t, found)
	})
}
