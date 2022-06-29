package shortcode

import (
	"log"
	"regexp"
	"strings"
)

type Args map[string]string

func parse(str string, shortcode string, fn func(map[string]string) string) string {
	regex := regexp.MustCompile(`<` + shortcode + `(.*)>([\S\s]+)</` + shortcode + `>`)
	//matches := re.FindAllStringSubmatch(str, -1)
	//log.Println(matches)
	log.Println("=============")
	// log.Println(content)

	for {
		match := regex.FindStringSubmatchIndex(str)
		if match == nil {
			break
		}
		args := Args{}
		openingTagStart, openingTagClose, tagNameStart, tagNameEnd := match[0], match[1], match[2], match[3]
		//fullMatch := str[openingTagStart:openingTagClose]
		tagName := str[tagNameStart:tagNameEnd]
		log.Println("TAG NAME: " + tagName)

		// Parse the arguments
		if match[4] != -1 {
			argsString := str[match[4]:match[5]]
			argsRegex := regexp.MustCompile(`\s*([^=]+)="([^"]+)"`)
			for _, argMatch := range argsRegex.FindAllStringSubmatch(argsString, -1) {
				args[argMatch[1]] = argMatch[2]
			}
		}

		//var closingTagEnd = openingTagClose
		var textToReplace = str[openingTagStart:openingTagClose]

		//replaced := s.registered[tagName](args)
		replaced := fn(args)

		str = strings.Replace(str, textToReplace, replaced, 1)
	}

	// for _, match := range matches {
	// 	log.Println("==== MATCH ====")
	// 	if match[0] == "" {
	// 		continue
	// 	}
	// 	//log.Println(match[0])
	// 	// attrs := match[1]
	// 	// text := (match[2])
	// 	match := regex.FindStringSubmatchIndex(text)
	// 	if match == nil {
	// 		break
	// 	}

	// 	str = strings.Replace(str, textToReplace, replaced, 1)
	// }

	return str
}
