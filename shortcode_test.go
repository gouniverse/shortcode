package shortcode

import (
	"testing"
)

func TestTrue(t *testing.T) {
	isOk := true

	if isOk == false {
		t.Fatalf("Cache could not be created")
	}
}
