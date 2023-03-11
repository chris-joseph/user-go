package utility

import "testing"

func TestUtility(t *testing.T) {
	type test[T comparable] struct {
		name       string
		collection []T
		element    T
		expected   bool
	}
	tests := []test[string]{
		{
			name:       "simple string test",
			collection: []string{"something", "somethingElse"},
			element:    "something",
			expected:   true,
		},
		{
			name:       "simple string test",
			collection: []string{"something", "somethingElse"},
			element:    "somethingOther",
			expected:   false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsInArray(test.collection, test.element)
			if result && !test.expected || !result && test.expected {
				t.Errorf("`%v` is present in the `%v` collection", test.element, test.collection)
			}
		})
	}
}
