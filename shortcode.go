package shortcode

import (
	"errors"
	"log"
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

func (sh Shortcode) Render(str string, shortcode string, fn func(string, map[string]string) string) string {
	escapedBracketOpening := strings.ReplaceAll(sh.bracketOpening, "[", "\\[")
	escapedBracketOpening = strings.ReplaceAll(escapedBracketOpening, "(", "\\(")
	escapedBracketClosing := strings.ReplaceAll(sh.bracketClosing, "]", "\\]")
	escapedBracketClosing = strings.ReplaceAll(escapedBracketClosing, ")", "\\)")
	attr := `(\s+[^` + escapedBracketClosing + `]+)?`
	start := escapedBracketOpening + shortcode + attr + escapedBracketClosing
	end := escapedBracketOpening + `/` + shortcode + escapedBracketClosing
	// content := `([\S\s]+)`
	// content := `([^\[]*)`
	content := `([^` + escapedBracketClosing + `]*)`

	regex := regexp.MustCompile(start + content + end)
	for _, match := range regex.FindAllStringSubmatch(str, -1) {
		if match[0] == "" {
			continue
		}
		attrs, content := match[1], match[2]
		log.Println(attrs, content)
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

// type Args map[string]string

// func (sh Shortcode) Render(str string, shortcode string, fn func(map[string]string) string) string {
// 	escapedBracketOpening := strings.ReplaceAll(sh.bracketOpening, "[", "\\[")
// 	escapedBracketClosing := strings.ReplaceAll(sh.bracketClosing, "]", "\\]")
// 	start := escapedBracketOpening + shortcode + `(.*)` + escapedBracketClosing
// 	//end := escapedBracketOpening + `/` + shortcode + escapedBracketClosing
// 	content := `([\S\s]+)`
// 	regex := regexp.MustCompile(start + content)
// 	//matches := re.FindAllStringSubmatch(str, -1)
// 	//log.Println(matches)
// 	log.Println("=============")
// 	// log.Println(content)

// 	for {
// 		match := regex.FindStringSubmatchIndex(str)
// 		if match == nil {
// 			break
// 		}
// 		// args := Args{}
// 		log.Println("MATCH 0:")
// 		log.Println(match[0])

// 		log.Println("MATCH 1:")
// 		log.Println(match[1])

// 		log.Println("MATCH 2:")
// 		log.Println(match[2])

// 		log.Println("MATCH 3:")
// 		log.Println(match[3])

// 		log.Println("MATCH 4:")
// 		log.Println(match[4])

// 		//openingTagStart, openingTagClose, tagNameStart, tagNameEnd := match[0], match[1], match[2], match[3]
// 		startStr, endStr, tagNameStart, tagNameEnd := match[0], match[1], match[2], match[3]
// 		fullMatch := str[startStr:endStr]
// 		tagName := str[tagNameStart:tagNameEnd]
// 		log.Println("FULL MATCH: " + fullMatch)
// 		log.Println("TAG NAME: " + tagName)

// 		// Parse the arguments
// 		// if match[4] != -1 {
// 		// 	argsString := str[match[4]:match[5]]
// 		// 	argsRegex := regexp.MustCompile(`\s*([^=]+)="([^"]+)"`)
// 		// 	for _, argMatch := range argsRegex.FindAllStringSubmatch(argsString, -1) {
// 		// 		args[argMatch[1]] = argMatch[2]
// 		// 	}
// 		// }
// 		// var closingTagEnd = openingTagClose
// 		// var textToReplace = str[tagNameStart:tagNameEnd]

// 		// replaced := fn(args)

// 		// str = strings.Replace(str, textToReplace, replaced, 1)
// 	}

// 	// for _, match := range matches {
// 	// 	log.Println("==== MATCH ====")
// 	// 	if match[0] == "" {
// 	// 		continue
// 	// 	}
// 	// 	//log.Println(match[0])
// 	// 	// attrs := match[1]
// 	// 	// text := (match[2])
// 	// 	match := regex.FindStringSubmatchIndex(text)
// 	// 	if match == nil {
// 	// 		break
// 	// 	}

// 	// 	str = strings.Replace(str, textToReplace, replaced, 1)
// 	// }

// 	return str
// }
