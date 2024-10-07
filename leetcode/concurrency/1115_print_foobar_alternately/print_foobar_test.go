package print_foobar

import (
	"bytes"
	"sync"
	"testing"
)

func TestCase1(t *testing.T) {
	cases := []int{1, 2, 3, 5, 8, 13, 21}

	for _, n := range cases {
		// Setup buffer to cupter the output
		var buf bytes.Buffer
		printFoo := func() {
			buf.WriteString("foo")
		}
		printBar := func() {
			buf.WriteString("bar")
		}

		fb := NewFooBar(n)
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			fb.Foo(printFoo)
		}()

		go func() {
			defer wg.Done()
			fb.Bar(printBar)
		}()

		wg.Wait()

		expected := ""
		for i := 0; i < n; i++ {
			expected += "foobar"
		}

		if result := buf.String(); result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	}
}

