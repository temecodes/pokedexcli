
package main	

import(
	"testing"
)

func testCleanInput(t *testing.T){
cases := []struct {
	input    string
	expected []string
}{
	{
		input:    "  hello  world  ",
		expected: []string{"hello", "world"},
	},
	// add more cases here
}
for _, c := range cases {
	actual := cleanInput(c.input)
	// Check the length of the actual slice against the expected slice
	if len(actual) != len(c.expected){
		t.Errorf("expected length %d, got %d", len(c.expected), len(actual))
		continue
	}
	// if they don't match, use t.Errorf to print an error message
	// and fail the test
	for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]
		
		if word != expectedWord{
			t.Errorf("expected word %q, got %q", expectedWord, word)
		}
		// Check each word in the slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
	}
}
}
