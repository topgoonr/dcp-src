// Package unique provides the functionality to create unique versions of slices.
package unique

// Strings returns a unique subset of the string slice provided.
func Strings(input []string) []string {
	u := make([]string, 0, len(input))
	m := make(map[string]bool)

	for _, val := range input { // don't want the index
		if _, ok := m[val]; !ok { // if NOT ok (with an initializer section)
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}

// for changing the comments
