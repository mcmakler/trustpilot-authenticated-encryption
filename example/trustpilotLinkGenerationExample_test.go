package example

import "testing"

func TestExampleLinkGeneration(t *testing.T) {
	t.Run("Should: no panic", func(t *testing.T) {
		ExampleLinkGeneration()
	})
}
