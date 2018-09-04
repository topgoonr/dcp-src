// Package unique provides the functionality to create unique versions of slices.
package unique

// Ints returns a unique subset of the int slice provided.
func Strings(input []string) []string {
	u := make([]string, 0, len(input))
	m := make(map[string]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}

// for changing the comments
