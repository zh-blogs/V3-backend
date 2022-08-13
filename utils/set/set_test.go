package set

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapSet(t *testing.T) {
	testSet(t, NewMapSet)
}

func testSet(t *testing.T, newSet func(...interface{}) Set) {
	assert := assert.New(t)

	// { a1 a2 }
	set := newSet("a1", "a2")
	assert.Equal(2, set.Len())
	assert.True(set.Contains("a1"))
	assert.True(set.Contains("a2"))
	assert.False(set.Contains("a3"))

	// { a1 a2 }
	set.Add("a1")
	assert.Equal(2, set.Len())
	assert.True(set.Contains("a1"))
	assert.True(set.Contains("a2"))
	assert.False(set.Contains("a3"))

	// { a1 a2 a3 }
	set.Add("a3")
	assert.Equal(3, set.Len())
	assert.True(set.Contains("a1"))
	assert.True(set.Contains("a2"))
	assert.True(set.Contains("a3"))

	// { a1 a2 a3 }
	set.Remove("a4")
	assert.Equal(3, set.Len())
	assert.True(set.Contains("a1"))
	assert.True(set.Contains("a2"))
	assert.True(set.Contains("a3"))

	// { a2 a3 }
	set.Remove("a1")
	assert.Equal(2, set.Len())
	assert.False(set.Contains("a1"))
	assert.True(set.Contains("a2"))
	assert.True(set.Contains("a3"))

	// { }
	set.Clear()
	assert.Equal(0, set.Len())
	assert.False(set.Contains("a1"))
	assert.False(set.Contains("a2"))
	assert.False(set.Contains("a3"))

	// { } { a3 a4 }
	set2 := set.Copy()
	set2.Add("a3")
	set2.Add("a4")
	assert.Equal(0, set.Len())
	assert.False(set.Contains("a1"))
	assert.False(set.Contains("a2"))
	assert.False(set.Contains("a3"))
	assert.False(set.Contains("a4"))
	assert.Equal(2, set2.Len())
	assert.False(set2.Contains("a1"))
	assert.False(set2.Contains("a2"))
	assert.True(set2.Contains("a3"))
	assert.True(set2.Contains("a4"))

	set.Add("a2")
	set.Add("a3")

	// { a2 a3 } union { a3 a4 }
	set3 := set.Union(set2)
	assert.Equal(2, set.Len())
	assert.False(set.Contains("a1"))
	assert.True(set.Contains("a2"))
	assert.True(set.Contains("a3"))
	assert.False(set.Contains("a4"))
	assert.Equal(2, set2.Len())
	assert.False(set2.Contains("a1"))
	assert.False(set2.Contains("a2"))
	assert.True(set2.Contains("a3"))
	assert.True(set2.Contains("a4"))
	assert.Equal(3, set3.Len())
	assert.False(set3.Contains("a1"))
	assert.True(set3.Contains("a2"))
	assert.True(set3.Contains("a3"))
	assert.True(set3.Contains("a4"))

	// { a2 a3 } intersect { a3 a4 }
	set3 = set.Intersect(set2)
	assert.Equal(2, set.Len())
	assert.False(set.Contains("a1"))
	assert.True(set.Contains("a2"))
	assert.True(set.Contains("a3"))
	assert.False(set.Contains("a4"))
	assert.Equal(2, set2.Len())
	assert.False(set2.Contains("a1"))
	assert.False(set2.Contains("a2"))
	assert.True(set2.Contains("a3"))
	assert.True(set2.Contains("a4"))
	assert.Equal(1, set3.Len())
	assert.False(set3.Contains("a1"))
	assert.False(set3.Contains("a2"))
	assert.True(set3.Contains("a3"))
	assert.False(set3.Contains("a4"))

	// { a2 a3 } difference { a3 a4 }
	set3 = set.Difference(set2)
	assert.Equal(2, set.Len())
	assert.False(set.Contains("a1"))
	assert.True(set.Contains("a2"))
	assert.True(set.Contains("a3"))
	assert.False(set.Contains("a4"))
	assert.Equal(2, set2.Len())
	assert.False(set2.Contains("a1"))
	assert.False(set2.Contains("a2"))
	assert.True(set2.Contains("a3"))
	assert.True(set2.Contains("a4"))
	assert.Equal(1, set3.Len())
	assert.False(set3.Contains("a1"))
	assert.True(set3.Contains("a2"))
	assert.False(set3.Contains("a3"))
	assert.False(set3.Contains("a4"))

	assert.False(set.Equal(set2)) // { a2 a3 } == { a3 a4 }
	assert.False(set.Equal(set3)) // { a2 a3 } == { a2 }
	set3.Add("a3")
	assert.True(set.Equal(set3)) // { a2 a3 } == { a2 a3 }

	count := 0
	err := set.ForEach(func(item interface{}) error {
		count++
		return fmt.Errorf("TESTERROR")
	})
	assert.Equal(1, count)
	assert.Equal("TESTERROR", err.Error())
}
