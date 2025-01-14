package words

import (
	"strings"
)

var initialisms = map[string]bool{
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IO":    true,
	"IP":    true,
	"JS":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SQL":   true,
	"SSH":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF":   true,
	"VM":    true,
	"XML":   true,
	"XMPP":  true,
	"XSRF":  true,
	"XSS":   true,
}

// ToSnakeCase converts a given string to snake_case.
// It splits words based on transitions between alphanumeric and non-alphanumeric characters,
// as well as transitions between lowercase and uppercase characters.
// Example: "ThisIsATest" -> "this_is_a_test"
func ToSnakeCase(s string) string {
	words := SplitWords(s)
	return strings.ToLower(strings.Join(words, "_"))
}

// ToConstantCase converts a given string to CONSTANT_CASE.
// It splits words based on transitions between alphanumeric and non-alphanumeric characters,
// as well as transitions between lowercase and uppercase characters.
// Example: "ThisIsATest" -> "THIS_IS_A_TEST"
func ToConstantCase(s string) string {
	words := SplitWords(s)
	return strings.ToUpper(strings.Join(words, "_"))
}

// ToKebabCase converts a given string to kebab-case.
// It splits words based on transitions between alphanumeric and non-alphanumeric characters,
// as well as transitions between lowercase and uppercase characters.
// Example: "ThisIsATest" -> "this-is-a-test"
func ToKebabCase(s string) string {
	words := SplitWords(s)
	return strings.ToLower(strings.Join(words, "-"))
}

// Capitalize capitalizes the first letter of a given string and converts the rest to lowercase.
// Example: "example" -> "Example"
func Capitalize(s string) string {
	if len(s) == 0 {
		return ""
	}
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}

// ToCapWords converts a given string to capitalized words, handling initialisms correctly.
// Example: "thisIsATest" -> "ThisIsATest"
func ToCapWords(s string) string {
	in := SplitWords(s)
	var out []string
	for _, word := range in {
		if initialisms[strings.ToUpper(word)] {
			out = append(out, strings.ToUpper(word))
		} else {
			out = append(out, Capitalize(word))
		}
	}
	return strings.Join(out, "")
}

// ToMixedCase converts a given string to mixedCase, handling initialisms correctly.
// Example: "thisIsATest" -> "thisIsATest"
func ToMixedCase(s string) string {
	in := SplitWords(s)
	var out []string
	for i, word := range in {
		upperWord := strings.ToUpper(word)
		if i == 0 {
			out = append(out, strings.ToLower(word))
		} else {
			if initialisms[upperWord] {
				out = append(out, upperWord)
			} else {
				out = append(out, Capitalize(word))
			}
		}
	}
	return strings.Join(out, "")
}

const (
	nonAlphaNumMask = 0 // 000
	alphaNumMask    = 1 // 001
	upperMask       = 2 // 010
)

// This function uses bitwise operations for efficiency, leveraging the ASCII
// structure of characters:
//
// 1. 'A-Z' and 'a-z' differ only in bit 5, so `(r & 0xDF)` normalizes both to 'A-Z'.
// 2. Digits ('0-9') are checked directly using range comparisons.
// 3. The uppercase detection `(r & 0x20) == 0` identifies uppercase letters.
func getMask(r byte) int {
	// Check if the character is alphanumeric (letters or digits).
	if (('A' <= (r & 0xDF)) && ((r & 0xDF) <= 'Z')) || ('0' <= r && r <= '9') {
		// Check if the character is uppercase using bitwise operations.
		if (r & 0x20) == 0 { // Uppercase letters have bit 5 cleared.
			return alphaNumMask | upperMask
		}
		return alphaNumMask
	}
	// Non-alphanumeric characters.
	return nonAlphaNumMask
}

// SplitWords splits a string into words based on transitions between alphanumeric
// and non-alphanumeric characters, as well as transitions between lowercase and uppercase characters.
// Example: "ThisIsATest" -> ["This", "Is", "A", "Test"]
func SplitWords(s string) []string {
	if len(s) == 0 {
		return nil
	}

	var words = make([]string, 0)
	var wordStart int = 0

	prevMask := nonAlphaNumMask
	curMask := nonAlphaNumMask
	nextMask := nonAlphaNumMask
	if len(s) > 0 {
		nextMask = getMask(s[0])
	}

	for i := 0; i < len(s); i++ {
		prevMask = curMask
		curMask = nextMask

		// Determine the mask for the next character, or set to non-alphanumeric if at the end of the string.
		if i+1 < len(s) {
			nextMask = getMask(s[i+1])
		} else {
			nextMask = nonAlphaNumMask
		}

		// Handle word boundary: Split when transitioning from an alphanumeric character to a non-alphanumeric character.
		if (curMask & alphaNumMask) == 0 {
			// previous character was alphanumeric. Append the last word.
			if ((prevMask & alphaNumMask) != 0) && i > wordStart {
				words = append(words, s[wordStart:i])
			}
			wordStart = i + 1
			continue
		}

		// Handle CamelCase: Split when current char is uppercase and either the next or previous is not.
		if ((curMask & upperMask) != 0) &&
			// Previous is not uppercase so boundary has occured
			(((prevMask & upperMask) == 0) ||
				// Next is not uppercase and alphanumeric so boundary has occured
				(((nextMask & upperMask) == 0) && ((nextMask & alphaNumMask) != 0))) {
			if i > wordStart {
				words = append(words, s[wordStart:i])
				wordStart = i
			}
		}
	}

	// Append the last word if the string ends with an alphanumeric character.
	if (curMask & alphaNumMask) != 0 {
		words = append(words, s[wordStart:])
	}

	return words
}
