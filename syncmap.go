package syncmap

import "sync"

type Map[K comparable, V any] struct {
	m sync.Map
}

// Delete deletes the value for a key.
func (m *Map[K, V]) Delete(k K) { m.m.Delete(k) }

// Load returns the value stored in the map for a key, or nil if no value is present.
// The ok result indicates whether value was found in the map.
func (m *Map[K, V]) Load(k K) (value V, ok bool) {
	v, ok := m.m.Load(k)
	if v == nil {
		return value, ok
	}
	return v.(V), ok
}

// LoadAndDelete deletes the value for a key, returning the previous value if any.
// The loaded result reports whether the key was present.
func (m *Map[K, V]) LoadAndDelete(k K) (value V, loaded bool) {
	v, loaded := m.m.LoadAndDelete(k)
	if v == nil {
		return value, loaded
	}
	return v.(V), loaded
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (m *Map[K, V]) LoadOrStore(k K, v V) (actual V, loaded bool) {
	a, loaded := m.m.LoadOrStore(k, v)
	return a.(V), loaded
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
//
// Range does not necessarily correspond to any consistent snapshot of the Map's contents: no key will be visited more than once, but if the value for any key is stored or deleted concurrently (including by f), Range may reflect any mapping for that key from any point during the Range call.
// Range does not block other methods on the receiver; even f itself may call any method on m.
//
// Range may be O(N) with the number of elements in the map even if f returns false after a constant number of calls.
func (m *Map[K, V]) Range(f func(k K, v V) bool) {
	m.m.Range(func(k, v any) bool { return f(k.(K), v.(V)) })
}

// Store sets the value for a key.
func (m *Map[K, V]) Store(k K, v V) { m.m.Store(k, v) }
