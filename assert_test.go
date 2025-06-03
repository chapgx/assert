package assert

import (
	"fmt"
	"testing"
)

func TestCore(t *testing.T) {

	t.Run("assert with testing", func(t *testing.T) {
		AssertT(t, 3 == 3, "they must be equal")
	})

	t.Run("raw assert ", func(t *testing.T) {
		Assert(3 != 4, "they must be equal")
	})

	output := Must(helper())
	fmt.Println(output...)
}

func helper() (string, error) {
	return "hello", nil
}
