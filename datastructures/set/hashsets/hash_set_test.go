package hashsets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashSet(t *testing.T) {
	set := New[string]()

	assert.Equal(t, set.IsEmpty(), true)

	// Test Add
	set.Add("cuong")

	assert.Equal(t, set.Contains("cuong"), true)
	assert.Equal(t, set.Contains("unknown"), false)
	assert.Equal(t, set.ContainsAll("cuong", "unknown"), false)
	assert.Equal(t, set.ContainsAny("cuong", "unknown"), true)
	assert.Equal(t, set.IsEmpty(), false)
	assert.Equal(t, set.Size(), 1)
	assert.Equal(t, set.GetValues()[0], "cuong")

	// Test Clear
	set.Clear()

	assert.Equal(t, set.IsEmpty(), true)
	assert.Equal(t, set.Size(), 0)

	// Test Remove
	set.Add("hello", "cool")

	assert.Equal(t, set.Size(), 2)
	assert.Equal(t, set.ContainsAll("hello", "cool"), true)
	assert.Equal(t, set.ContainsAny("hello"), true)

	set.Remove("hello")
	assert.Equal(t, set.Contains("hello"), false)
	assert.Equal(t, set.ContainsAll("hello"), false)
	assert.Equal(t, set.ContainsAny("hello"), false)
}

func TestHashSetMerge(t *testing.T) {
	set1 := New[string]().(*HashSet[string])
	set1.Add("hello", "cool")

	set2 := New("hello", "sweet").(*HashSet[string])

	set3 := New("cool", "cold")
	set3.Merge(set1, set2)

	assert.Equal(t, set3.Size(), 4)
	assert.Equal(t, set3.ContainsAll("hello", "cool", "sweet", "cold"), true)
}
func TestHashSetCopy(t *testing.T) {
	set1 := New[string]()
	set1.Add("a", "b", "c")

	set2 := New[string]()
	set2.Copy()

	assert.Equal(t, set2.Size(), 3)
	assert.Equal(t, set2.ContainsAll("a", "b", "c"), true)
}
func TestHashSetUnion(t *testing.T) {
	set1 := New("hello", "cool", "cold").(*HashSet[string])

	set2 := New("cool", "cold", "bye").(*HashSet[string])

	set3 := set1.Union(set2)
	assert.Equal(t, set3.Size(), 4)
	assert.Equal(t, set3.ContainsAll("cool", "cold", "hello", "bye"), true)
}
func TestHashSetIntersection(t *testing.T) {
	set1 := New("hello", "cool", "cold").(*HashSet[string])

	set2 := New("cool", "cold", "bye").(*HashSet[string])

	set3 := set1.Intersection(set2)
	assert.Equal(t, set3.Size(), 4)
	assert.Equal(t, set3.ContainsAll("cool", "cold", "hello", "bye"), true)
}
func TestHashSetSymmetricDifference(t *testing.T) {
	set1 := New("hello", "cool", "cold").(*HashSet[string])

	set2 := New("cool", "cold", "bye").(*HashSet[string])

	set3 := set1.SymmetricDifference(set2)
	assert.Equal(t, set3.Size(), 2)
	assert.Equal(t, set3.ContainsAll("hello", "bye"), true)
}
func TestHashSetSubtraction(t *testing.T) {
	set1 := New("hello", "cool", "cold").(*HashSet[string])

	set2 := New("cool", "cold", "bye").(*HashSet[string])

	set3 := set1.Subtraction(set2)
	assert.Equal(t, set3.Size(), 1)
	assert.Equal(t, set3.ContainsAll("hello"), true)
}
