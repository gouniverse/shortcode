# Shortcode <a href="https://gitpod.io/#https://github.com/gouniverse/shortcode" style="float:right:"><img src="https://gitpod.io/button/open-in-gitpod.svg" alt="Open in Gitpod" loading="lazy"></a>

[![Tests Status](https://github.com/gouniverse/shortcode/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/gouniverse/shortcode/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/gouniverse/shortcode)](https://goreportcard.com/report/github.com/gouniverse/shortcode)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/gouniverse/shortcode)](https://pkg.go.dev/github.com/gouniverse/shortcode)

Shortcode implementation. Convenient for CMS content where you want to add additional functionality. Similar to the short codes, as known from WordPress.

## Installation
```
go get -u github.com/gouniverse/shortcode
```

## Example

Source file:

```html
[myshortcode id="111"][/myshortcode]
[myshortcode id="222"][/myshortcode]
```

Go code
```golang
func myShortcode(content string, args map[string]string) string {
	return "MY SHORTCODE WITH ID " + args["id"]
}
sh, err := NewShortcode(WithBrackets("[", "]"))
parsed := sh.Render(text, "myshortcode", myShortcode)
```

Result
```html
MY SHORTCODE WITH ID 111
MY SHORTCODE WITH ID 222
```
