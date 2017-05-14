package main

import (
	"fmt"
)

type state uint8

const (
	// stateText is parsed character data. An HTML parser is in
	// this state when its parse position is outside an HTML tag,
	// directive, comment, and special element body.
	stateText state = iota
	// stateTag occurs before an HTML attribute or the end of a tag.
	stateTag
	// stateAttrName occurs inside an attribute name.
	// It occurs between the ^'s in ` ^name^ = value`.
	stateAttrName
	// stateAfterName occurs after an attr name has ended but before any
	// equals sign. It occurs between the ^'s in ` name^ ^= value`.
	stateAfterName
	// stateBeforeValue occurs after the equals sign but before the value.
	// It occurs between the ^'s in ` name =^ ^value`.
	stateBeforeValue
	// stateHTMLCmt occurs inside an <!-- HTML comment -->.
	stateHTMLCmt
	// stateRCDATA occurs inside an RCDATA element (<textarea> or <title>)
	// as described at http://www.w3.org/TR/html5/syntax.html#elements-0
	stateRCDATA
	// stateAttr occurs inside an HTML attribute whose content is text.
	stateAttr
	// stateURL occurs inside an HTML attribute whose content is a URL.
	stateURL
	// stateJS occurs inside an event handler or script element.
	stateJS
	// stateJSDqStr occurs inside a JavaScript double quoted string.
	stateJSDqStr
	// stateJSSqStr occurs inside a JavaScript single quoted string.
	stateJSSqStr
	// stateJSRegexp occurs inside a JavaScript regexp literal.
	stateJSRegexp
	// stateJSBlockCmt occurs inside a JavaScript /* block comment */.
	stateJSBlockCmt
	// stateJSLineCmt occurs inside a JavaScript // line comment.
	stateJSLineCmt
	// stateCSS occurs inside a <style> element or style attribute.
	stateCSS
	// stateCSSDqStr occurs inside a CSS double quoted string.
	stateCSSDqStr
	// stateCSSSqStr occurs inside a CSS single quoted string.
	stateCSSSqStr
	// stateCSSDqURL occurs inside a CSS double quoted url("...").
	stateCSSDqURL
	// stateCSSSqURL occurs inside a CSS single quoted url('...').
	stateCSSSqURL
	// stateCSSURL occurs inside a CSS unquoted url(...).
	stateCSSURL
	// stateCSSBlockCmt occurs inside a CSS /* block comment */.
	stateCSSBlockCmt
	// stateCSSLineCmt occurs inside a CSS // line comment.
	stateCSSLineCmt
	// stateError is an infectious error state outside any valid
	// HTML/CSS/JS construct.
	stateError
)

var stateNames = [...]string{
	stateText:        "stateText",
	stateTag:         "stateTag",
	stateAttrName:    "stateAttrName",
	stateAfterName:   "stateAfterName",
	stateBeforeValue: "stateBeforeValue",
	stateHTMLCmt:     "stateHTMLCmt",
	stateRCDATA:      "stateRCDATA",
	stateAttr:        "stateAttr",
	stateURL:         "stateURL",
	stateJS:          "stateJS",
	stateJSDqStr:     "stateJSDqStr",
	stateJSSqStr:     "stateJSSqStr",
	stateJSRegexp:    "stateJSRegexp",
	stateJSBlockCmt:  "stateJSBlockCmt",
	stateJSLineCmt:   "stateJSLineCmt",
	stateCSS:         "stateCSS",
	stateCSSDqStr:    "stateCSSDqStr",
	stateCSSSqStr:    "stateCSSSqStr",
	stateCSSDqURL:    "stateCSSDqURL",
	stateCSSSqURL:    "stateCSSSqURL",
	stateCSSURL:      "stateCSSURL",
	stateCSSBlockCmt: "stateCSSBlockCmt",
	stateCSSLineCmt:  "stateCSSLineCmt",
	stateError:       "stateError2",
}

func (s state) String() string {
	if int(s) < len(stateNames) {
		return stateNames[s]
	}
	return fmt.Sprintf("illegal state %d", int(s))
}

func main() {
	fmt.Println(stateNames)
	fmt.Println(stateNames[0])
	fmt.Println(stateNames[1])
}
