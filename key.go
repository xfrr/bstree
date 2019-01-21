package btree

// Key ...
type Key interface {
	GreaterThan(Key) bool
	LessThan(Key) bool
	EqualTo(Key) bool
}

// IKey is a int key type
type IKey int

// GreaterThan ...
func (i IKey) GreaterThan(vk Key) bool {
	return i < vk.(IKey)
}

// LessThan ...
func (i IKey) LessThan(vk Key) bool {
	return i > vk.(IKey)
}

// EqualTo ...
func (i IKey) EqualTo(vk Key) bool {
	return i == vk.(IKey)
}

// SKey is a string key type
type SKey string

// GreaterThan ...
func (s SKey) GreaterThan(vk Key) bool {
	return s < vk.(SKey)
}

// LessThan ...
func (s SKey) LessThan(vk Key) bool {
	return s > vk.(SKey)
}

// EqualTo ...
func (s SKey) EqualTo(vk Key) bool {
	return s == vk.(SKey)
}
