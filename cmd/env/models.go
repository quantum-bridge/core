package env

// Tag is a struct that holds the key, required, and non-zero values of a tag.
type Tag struct {
	Key      string
	Required bool
	NonZero  bool
}
