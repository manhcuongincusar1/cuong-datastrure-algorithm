package hashmultimaps

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashMultiMap(t *testing.T) {
	multiMaps := New[string, int]()
	// Put
	multiMaps.Put("x", 1)

	// Info
	assert.Equal(t, multiMaps.Size(), 1)
	assert.True(t, multiMaps.Contains("x"))
	assert.True(t, reflect.DeepEqual(multiMaps.GetValues("x"), []int{1}))
	assert.True(t, reflect.DeepEqual(multiMaps.GetValues("unkown"), []int{}))

	// Put Multi
	multiMaps.PutAll("y", 1, 3, 4, 56)
	assert.Equal(t, multiMaps.Size(), 2)
	assert.True(t, multiMaps.ContainsAll("x", "y"))
	assert.True(t, multiMaps.ContainsAny("x"))
	// Remove
	multiMaps.Remove("y", 1)
	assert.Equal(t, multiMaps.Size(), 2)
	assert.True(t, multiMaps.ContainsAll("x", "y"))
	assert.True(t, multiMaps.ContainsAny("x"))
	assert.True(t, reflect.DeepEqual(multiMaps.GetValues("y"), []int{1, 3, 4, 56}))

	multiMaps.Remove("y", 3)
	assert.True(t, reflect.DeepEqual(multiMaps.GetValues("y"), []int{1, 4, 56}))
	// Clear
	multiMaps.Clear()
	assert.Equal(t, multiMaps.Size(), 0)
}
func TestHashMultiMapMerge(t *testing.T) {
	// Merge
	root := New[string, int]()
	root.PutAll("x", 1, 2, 3)

	ob := New[string, int]()
	ob.PutAll("y", 1, 2, 3)
	_map := ob.(*HashMultiMap[string, int])

	root.Merge(_map)

	assert.Equal(t, root.Size(), 2)
	assert.True(t, reflect.DeepEqual(root.GetValues("x"), []int{1, 2, 3}))
	assert.True(t, reflect.DeepEqual(root.GetValues("y"), []int{1, 2, 3}))

}
