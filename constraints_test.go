package g

import "testing"

// Named types are the norm in domain code, and Verifiable is the constraint
// behind the helpers domain code reaches for. Every one of these calls used to
// fail to compile when the union said string instead of ~string.
func TestVerifiableAcceptsNamedTypes(t *testing.T) {
	type Category string
	type Level int
	type Price float64

	if !Contains(Category("books"), []Category{"books", "music"}) {
		t.Error("Contains lost a named string type")
	}
	if !In(Level(2), Level(1), Level(2), Level(3)) {
		t.Error("In lost a named int type")
	}
	if got := Max(Price(1.5), Price(2.5)); got != 2.5 {
		t.Errorf("Max on a named float type = %v, want 2.5", got)
	}
	if got := Rank(Level(2), []Level{1, 2, 3}); got != 2 {
		t.Errorf("Rank on a named int type = %d, want 2", got)
	}
}

// rune is an alias for int32, and Integer already contains ~int32, so dropping
// the explicit rune member changed nothing: the type set is the same.
func TestVerifiableStillCoversRunes(t *testing.T) {
	if !Contains('b', []rune{'a', 'b'}) {
		t.Error("Contains lost plain runes")
	}
	type Glyph rune
	if !In(Glyph('x'), Glyph('x'), Glyph('y')) {
		t.Error("In lost a named rune type")
	}
}
