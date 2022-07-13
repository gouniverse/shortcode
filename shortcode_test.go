package shortcode

import (
	"testing"
)

func TestReturnsErrorOnMisconfiguration(t *testing.T) {
	_, err := NewShortcode()

	if err == nil {
		t.Fatalf("Shortcode must throw one error")
	}

	expectedError := "shortcode: bracketOpening is required"
	if err.Error() != expectedError {
		t.Fatalf("Shortcode must return error of '%s', instead returned '%s'", expectedError, err.Error())
	}

	_, err2 := NewShortcode(WithBrackets("", ""))

	if err2 == nil {
		t.Fatalf("Shortcode must throw one error")
	}

	expectedError2 := "shortcode: bracketOpening is required"
	if err2.Error() != expectedError2 {
		t.Fatalf("Shortcode must return error of '%s', instead returned '%s'", expectedError2, err2.Error())
	}

	_, err3 := NewShortcode(WithBrackets("[", ""))

	if err3 == nil {
		t.Fatalf("Shortcode must throw one error")
	}

	expectedError3 := "shortcode: bracketClosing is required"
	if err3.Error() != expectedError3 {
		t.Fatalf("Shortcode must return error of '%s', instead returned '%s'", expectedError3, err3.Error())
	}
}

func TestWithSquareBrackets(t *testing.T) {
	sh, err := NewShortcode(WithBrackets("[", "]"))

	if err != nil {
		t.Fatalf("Shortcode must not throw an error, but '%s'", err.Error())
	}

	text := `TEXT BEFORE [x-text id="222"]
	TEST
	[/x-text] TEXT AFTER`

	parsed := sh.Render(text, "x-text", testShortcode)

	expected := "TEXT BEFORE SHORTCODE WITH ID 222 TEXT AFTER"

	if parsed != expected {
		t.Fatalf("Expected '%s', instead returned '%s'", expected, parsed)
	}
}

func TestWithBrackets(t *testing.T) {
	sh, err := NewShortcode(WithBrackets("(", ")"))

	if err != nil {
		t.Fatalf("Shortcode must not throw an error, but '%s'", err.Error())
	}

	text := `TEXT BEFORE (x-text id="222")
	TEST
	(/x-text) TEXT AFTER`

	parsed := sh.Render(text, "x-text", testShortcode)

	expected := "TEXT BEFORE SHORTCODE WITH ID 222 TEXT AFTER"

	if parsed != expected {
		t.Fatalf("Expected '%s', instead returned '%s'", expected, parsed)
	}
}

func TestWithAngleBrackets(t *testing.T) {
	sh, err := NewShortcode(WithBrackets("<", ">"))

	if err != nil {
		t.Fatalf("Shortcode must not throw an error, but '%s'", err.Error())
	}

	text := `TEXT BEFORE <x-text id="222">
	TEST
	</x-text> TEXT AFTER`

	parsed := sh.Render(text, "x-text", testShortcode)

	expected := "TEXT BEFORE SHORTCODE WITH ID 222 TEXT AFTER"

	if parsed != expected {
		t.Fatalf("Expected '%s', instead returned '%s'", expected, parsed)
	}
}

func TestWithMultipleSquareBrackets(t *testing.T) {
	sh, err := NewShortcode(WithBrackets("[", "]"))

	if err != nil {
		t.Fatalf("Shortcode must not throw an error, but '%s'", err.Error())
	}

	text := `TEXT BEFORE [x-text id="111"]TEST[/x-text] [x-text id="222"]TEST[/x-text] TEXT AFTER`

	parsed := sh.Render(text, "x-text", testShortcode)

	expected := "TEXT BEFORE SHORTCODE WITH ID 111 SHORTCODE WITH ID 222 TEXT AFTER"

	if parsed != expected {
		t.Fatalf("Expected '%s', instead returned '%s'", expected, parsed)
	}
}

func TestWithMultipleAngleBrackets(t *testing.T) {
	sh, err := NewShortcode(WithBrackets("<", ">"))

	if err != nil {
		t.Fatalf("Shortcode must not throw an error, but '%s'", err.Error())
	}

	text := `TEXT BEFORE <x-text id="111">TEST</x-text> <x-text id="222">TEST</x-text> TEXT AFTER`

	parsed := sh.Render(text, "x-text", testShortcode)

	expected := "TEXT BEFORE SHORTCODE WITH ID 111 SHORTCODE WITH ID 222 TEXT AFTER"

	if parsed != expected {
		t.Fatalf("Expected '%s', instead returned '%s'", expected, parsed)
	}
}

func TestInnerText(t *testing.T) {
	sh, err := NewShortcode(WithBrackets("[", "]"))

	if err != nil {
		t.Fatalf("Shortcode must not throw an error, but '%s'", err.Error())
	}

	text := `TEXT BEFORE [x-text id="111"]CONTENT[/x-text] TEXT AFTER`

	parsed := sh.Render(text, "x-text", testTextShortcode)

	expected := "TEXT BEFORE --- SHORTCODE START: CONTENT FROM ID 111 : SHORTCODE END --- TEXT AFTER"

	if parsed != expected {
		t.Fatalf("Expected '%s', instead returned '%s'", expected, parsed)
	}
}

func TestInnerTextMultiline(t *testing.T) {
	sh, err := NewShortcode(WithBrackets("[", "]"))

	if err != nil {
		t.Fatalf("Shortcode must not throw an error, but '%s'", err.Error())
	}

	text := `TEXT BEFORE [x-text id="111"]
	
	CONTENT
	
	[/x-text] TEXT AFTER`

	parsed := sh.Render(text, "x-text", testTextShortcode)

	expected := `TEXT BEFORE --- SHORTCODE START: 
	
	CONTENT
	
	 FROM ID 111 : SHORTCODE END --- TEXT AFTER`

	if parsed != expected {
		t.Fatalf("Expected '%s', instead returned '%s'", expected, parsed)
	}
}

func testShortcode(content string, args map[string]string) string {
	return "SHORTCODE WITH ID " + args["id"]
}

func testTextShortcode(content string, args map[string]string) string {
	return "--- SHORTCODE START: " + content + " FROM ID " + args["id"] + " : SHORTCODE END ---"
}
