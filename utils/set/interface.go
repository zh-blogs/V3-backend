package set

// Set ...
type Set interface {
	// Contains return true if item is in set
	Contains(interface{}) bool
	// ForEach will callback each items. If returns error without nil, will stop enumerate
	ForEach(func(interface{}) error) error
	// Len returns the count of items
	Len() int

	// Add item to set
	Add(...interface{})
	// Remove item from set
	Remove(...interface{})
	// Clear the set
	Clear()

	// Equal returns true if two set all equal
	Equal(Set) bool
	// Copy returns new set with same items(not deep copy)
	Copy() Set

	// Union returns new set which is union of two set
	Union(Set) Set
	// Intersect returns new set which is intersect of two set
	Intersect(Set) Set
	// Difference returns new set which is difference of two set
	Difference(Set) Set
}
