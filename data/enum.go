package data

// Enum is the data out to render Enums in a file
// Enums that nested inside messages will be pulled out to the top level
// Because the way it works in typescript
type Enum struct {
	// Name will be the package level unique name
	// Nested names will concat with their parent messages so that it will remain unique
	// This also means nested type might be a bit ugly in type script but whatever
	Name   string
	Values []int32
}

// NewEnum creates an enum instance.
func NewEnum() *Enum {
	return &Enum{
		Name:   "",
		Values: make([]int32, 0),
	}
}
