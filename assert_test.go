package assert

import (
	"testing"
)

func TestCore(t *testing.T) {

	t.Run("assert with testing", func(t *testing.T) {
		AssertT(t, 3 == 3, "they must be equal")
	})

	t.Run("raw assert ", func(t *testing.T) {
		Assert(3 != 4, "they must be equal")
	})
}
