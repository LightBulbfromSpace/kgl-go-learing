package generics

import (
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		myStackOfInts := new(Stack[int])

		// check stack is empty
		AssertTrue(t, myStackOfInts.IsEmpty())

		// add a thing, then check it's not empty
		myStackOfInts.Push(142)
		AssertFalse(t, myStackOfInts.IsEmpty())

		// add another thing, pop it back again
		myStackOfInts.Push(456)
		value, _ := myStackOfInts.Pop()
		AssertEqual(t, value, 456)

		value, _ = myStackOfInts.Pop()
		AssertEqual(t, value, 142)
		AssertTrue(t, myStackOfInts.IsEmpty())
	})
	t.Run("string stack", func(t *testing.T) {
		myStackOfStrings := new(Stack[string])

		// check stack is empty
		AssertTrue(t, myStackOfStrings.IsEmpty())

		// add a thing, then check it's not empty
		myStackOfStrings.Push("142")
		AssertFalse(t, myStackOfStrings.IsEmpty())

		// add another thing, pop it back again
		myStackOfStrings.Push("456")
		value, _ := myStackOfStrings.Pop()
		AssertEqual(t, value, "456")

		value, _ = myStackOfStrings.Pop()
		AssertEqual(t, value, "142")
		AssertTrue(t, myStackOfStrings.IsEmpty())
	})
	t.Run("interface stack dx is horrid", func(t *testing.T) {
		myStackOfInts := new(Stack[interface{}])

		myStackOfInts.Push(1)
		myStackOfInts.Push(2)
		firstNum, _ := myStackOfInts.Pop()
		secondNum, _ := myStackOfInts.Pop()

		// get our ints from out interface{}
		reallyFirstNum, ok := firstNum.(int)
		AssertTrue(t, ok)

		reallySecondNum, ok := secondNum.(int)

		AssertEqual(t, reallyFirstNum+reallySecondNum, 3)
	})
}

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})
	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "Grace")
	})
	t.Run("asserting on bool", func(t *testing.T) {
		AssertTrue(t, true)
		AssertFalse(t, false)
	})

	// AssertEqual(t, 1, "1") // uncomment to see the error
}
