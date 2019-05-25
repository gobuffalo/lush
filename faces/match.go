package faces

// Match defines an interface to support the
// "matches" of one type from an another.
type Match interface {
	Match(string) (bool, error)
}
