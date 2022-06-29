package shortcode

import (
	"log"
	"testing"
)

func TestTrue(t *testing.T) {
	text := `<x-text id="222"> 
	TEST    
	</x-text>`

	parsed := parse(text, "x-text", testShortcode)

	isOk := parsed == "SHORTCODE"

	if isOk == false {
		log.Println(parsed)
		t.Fatalf("Shortcode could not be found")
	}
}

func testShortcode(args map[string]string) string {
	log.Println(args)
	return "SHORTCODE"
}
