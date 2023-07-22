package syncmap

import (
	"testing"
)

type nillable struct {
	value string
}

func TestDelete(t *testing.T) {
	var m Map[string, *nillable]

	// Delete a key that was never present.
	m.Delete("test")

	// Store a key, Delete it and then try to Load it.
	a := &nillable{"value"}
	m.Store("test", a)
	m.Delete("test")
	g, ok := m.Load("test")
	if ok || g == a {
		t.Error("Key was present after Delete")
	}
}

func TestLoad(t *testing.T) {
	var m Map[string, *nillable]

	// Load a key that was never present.
	g, ok := m.Load("test")
	if ok || g != nil {
		t.Error("Key was present before Store")
	}

	// Store a key and then try to Load it.
	a := &nillable{"value"}
	m.Store("test", a)
	g, ok = m.Load("test")
	if !ok || g != a {
		t.Error("Loaded value does not match Stored value")
	}

	// Store and Load a nil value.
	m.Store("testnil", nil)
	g, ok = m.Load("testnil")
	if !ok || g != nil {
		t.Errorf("Loaded value does not match Stored value (nil)")
	}
}

func TestLoadAndDelete(t *testing.T) {
	var m Map[string, *nillable]

	// LoadAndDelete a key that was never present.
	g, ok := m.LoadAndDelete("test")
	if ok || g != nil {
		t.Error("Key was present before Store")
	}

	// Store a key and then try to LoadAndDelete it.
	a := &nillable{"value"}
	m.Store("test", a)
	g, ok = m.LoadAndDelete("test")
	if !ok || g != a {
		t.Error("Loaded value does not match Stored value")
	}
	g, ok = m.LoadAndDelete("test")
	if ok || g == a {
		t.Error("Loaded value still present after LoadAndDelete")
	}

	// Store and LoadAndDelete a nil value.
	m.Store("testnil", nil)
	g, ok = m.LoadAndDelete("testnil")
	if !ok || g != nil {
		t.Errorf("Loaded value does not match Stored value (nil)")
	}
	_, ok = m.LoadAndDelete("testnil")
	if ok {
		t.Errorf("Loaded value still present after LoadAndDelete (nil)")
	}
}

func TestLoadOrStore(t *testing.T) {
	var m Map[string, *nillable]

	a := &nillable{"value"}

	// LoadOrStore a key that was never present.
	g, ok := m.LoadOrStore("test", a)
	if ok || g != a {
		t.Error("Key was present before LoadOrStore")
	}

	a2 := &nillable{"value2"}

	// LoadOrStore a different value.
	g, ok = m.LoadOrStore("test", a2)
	if !ok || g == a2 {
		t.Error("Stored value was not Loaded by LoadOrStore")
	}

	// LoadOrStore a nil value.
	g, ok = m.LoadOrStore("testnil", nil)
	if ok || g != nil {
		t.Errorf("Key was present before LoadOrStore (nil)")
	}
	g, ok = m.LoadOrStore("testnil", nil)
	if !ok || g != nil {
		t.Error("Stored value was not Loaded by LoadOrStore (nil)")
	}
}

func TestRange(t *testing.T) {
	var m Map[string, *nillable]

	a := &nillable{"value"}

	m.Store("test", a)
	m.Store("testnil", nil)

	var test, testnil bool

	m.Range(func(k string, v *nillable) bool {
		if k == "test" && v == a {
			test = true
		}
		if k == "testnil" && v == nil {
			testnil = true
		}
		return true
	})

	if !test || !testnil {
		t.Error("Not all keys were visited")
	}
}
