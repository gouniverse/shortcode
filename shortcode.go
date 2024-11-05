package shortcode

import (
	"errors"
	"net/http"
	"regexp"
	"strings"
)

// Shortcode defines a shortcode engine
type Shortcode struct {
	bracketOpening string
	bracketClosing string
}

// ShortcodeOption options for the shortcode
type ShortcodeOption func(*Shortcode)

// WithBracketType sets the type of bracket used for the shortcode
func WithBrackets(bracketOpening string, bracketClosing string) ShortcodeOption {
	return func(s *Shortcode) {
		s.bracketOpening = bracketOpening
		s.bracketClosing = bracketClosing
	}
}

// NewShortcode creates a new shortcode engine
func NewShortcode(opts ...ShortcodeOption) (*Shortcode, error) {
	shortcode := &Shortcode{}
	for _, opt := range opts {
		opt(shortcode)
	}

	if shortcode.bracketOpening == "" {
		return nil, errors.New("shortcode: bracketOpening is required")
	}

	if shortcode.bracketClosing == "" {
		return nil, errors.New("shortcode: bracketClosing is required")
	}

	return shortcode, nil
}

func (sh Shortcode) RenderWithRequest(req *http.Request, str string, shortcode string, fn func(*http.Request, string, map[string]string) string) string {
	escapedBracketOpening := strings.ReplaceAll(sh.bracketOpening, "[", "\\[")
	escapedBracketOpening = strings.ReplaceAll(escapedBracketOpening, "(", "\\(")
	escapedBracketClosing := strings.ReplaceAll(sh.bracketClosing, "]", "\\]")
	escapedBracketClosing = strings.ReplaceAll(escapedBracketClosing, ")", "\\)")
	attr := `(\s+[^` + escapedBracketClosing + `]+)?`
	start := escapedBracketOpening + shortcode + attr + escapedBracketClosing
	end := escapedBracketOpening + `/` + shortcode + escapedBracketClosing
	//content := `([^` + escapedBracketClosing + `]*)`
	//content := `([\S\s]+?.*?|\s?)`
	content := `([^~]*?)`

	// DEBUG: log.Println(start + content + end)

	regex := regexp.MustCompile(start + content + end)
	for _, match := range regex.FindAllStringSubmatch(str, -1) {
		if match[0] == "" {
			continue
		}
		attrs, content := match[1], match[2]
		// DEBUG: log.Println(attrs, content)
		shortcodeResult := fn(req, content, attrsToArgs(attrs))
		str = strings.Replace(str, match[0], shortcodeResult, 1)
	}

	return str
}

func (sh Shortcode) Render(str string, shortcode string, fn func(string, map[string]string) string) string {
	escapedBracketOpening := strings.ReplaceAll(sh.bracketOpening, "[", "\\[")
	escapedBracketOpening = strings.ReplaceAll(escapedBracketOpening, "(", "\\(")
	escapedBracketClosing := strings.ReplaceAll(sh.bracketClosing, "]", "\\]")
	escapedBracketClosing = strings.ReplaceAll(escapedBracketClosing, ")", "\\)")
	attr := `(\s+[^` + escapedBracketClosing + `]+)?`
	start := escapedBracketOpening + shortcode + attr + escapedBracketClosing
	end := escapedBracketOpening + `/` + shortcode + escapedBracketClosing
	// content := `([^` + escapedBracketClosing + `]*)`
	// content := `([\S\s]+?.*?|\s?)`
	content := `([^~]*?)`

	regex := regexp.MustCompile(start + content + end)
	for _, match := range regex.FindAllStringSubmatch(str, -1) {
		if match[0] == "" {
			continue
		}
		attrs, content := match[1], match[2]
		// DEBUG: log.Println(attrs, content)
		shortcodeResult := fn(content, attrsToArgs(attrs))
		str = strings.Replace(str, match[0], shortcodeResult, 1)
	}

	return str
}

func attrsToArgs(attrs string) map[string]string {
	args := map[string]string{}
	argsRegex := regexp.MustCompile(`\s*([^=]+)="([^"]+)"`)
	for _, argMatch := range argsRegex.FindAllStringSubmatch(attrs, -1) {
		args[argMatch[1]] = argMatch[2]
	}
	return args
}
