package words_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	. "github.com/bartdeboer/words"
)

func TestToSnakeCase(t *testing.T) {
	tests := map[string]string{
		"TestCase":          "test_case",
		"AnotherExample123": "another_example123",
		"IOSystem":          "io_system",
		"IOSystem-product":  "io_system_product",
		"A":                 "a",
		"-":                 "",
	}

	for input, expected := range tests {
		result := ToSnakeCase(input)
		if result != expected {
			t.Errorf("ToSnakeCase(%s) = %s; want %s", input, result, expected)
		}
	}
}

func TestToConstantCase(t *testing.T) {
	tests := map[string]string{
		"TestCase":          "TEST_CASE",
		"AnotherExample123": "ANOTHER_EXAMPLE123",
		"IOSystem":          "IO_SYSTEM",
		"IOSystem-product":  "IO_SYSTEM_PRODUCT",
		"A":                 "A",
		"-":                 "",
	}

	for input, expected := range tests {
		result := ToConstantCase(input)
		if result != expected {
			t.Errorf("ToSnakeCase(%s) = %s; want %s", input, result, expected)
		}
	}
}

func TestToKebabCase(t *testing.T) {
	tests := map[string]string{
		"TestCase":          "test-case",
		"HTTPServerError":   "http-server-error",
		"TestCaseB":         "test-case-b",
		"Test-CaseC":        "test-case-c",
		"AnotherExample123": "another-example123",
		"A":                 "a",
		"-":                 "",
	}

	for input, expected := range tests {
		result := ToKebabCase(input)
		if result != expected {
			t.Errorf("ToKebabCase(%s) = %s; want %s", input, result, expected)
		}
	}
}

func TestToCapWords(t *testing.T) {
	tests := map[string]string{
		"test_case":                   "TestCase",
		"another-example123":          "AnotherExample123",
		"__another-_-example---123--": "AnotherExample123",
		"io-system-123":               "IOSystem123",
		"xml-http-request":            "XMLHTTPRequest",
		"id":                          "ID",
		"app-id":                      "AppID",
	}

	for input, expected := range tests {
		result := ToCapWords(input)
		if result != expected {
			t.Errorf("ToCapWords(%s) = %s; want %s", input, result, expected)
		}
	}
}

func TestMixedCase(t *testing.T) {
	tests := map[string]string{
		"test_case":                   "testCase",
		"another-example123":          "anotherExample123",
		"__another-_-example---123--": "anotherExample123",
		"io-system-123":               "ioSystem123",
		"xml-http-request":            "xmlHTTPRequest",
		"id":                          "id",
		"app-id":                      "appID",
	}

	for input, expected := range tests {
		result := ToMixedCase(input)
		if result != expected {
			t.Errorf("ToMixedCase(%s) = %s; want %s", input, result, expected)
		}
	}
}

func TestSplitWords(t *testing.T) {
	tests := map[string][]string{
		"TestCase":                    {"Test", "Case"},
		"AaBbCc":                      {"Aa", "Bb", "Cc"},
		"Aa-B--b-CcC":                 {"Aa", "B", "b", "Cc", "C"},
		"helloWorld123":               {"hello", "World123"},
		"splitCamelCase":              {"split", "Camel", "Case"},
		"The quick brown fox":         {"The", "quick", "brown", "fox"},
		"another-example123":          {"another", "example123"},
		"__another-_-example---123--": {"another", "example", "123"},
		"io-system-123":               {"io", "system", "123"},
		"A":                           {"A"},
		"":                            {},
		"-":                           {},
		"-__---@#$*^  (*&$ ---)":      {},
		"handleXML":                   {"handle", "XML"},
		"HandleXMLSystem":             {"Handle", "XML", "System"},
		"HandleJSON-System":           {"Handle", "JSON", "System"},
		"IOSystemInterface":           {"IO", "System", "Interface"},
		"LegacyXMLSystem":             {"Legacy", "XML", "System"},
		"XMLToJSONConverter":          {"XML", "To", "JSON", "Converter"},
	}

	for input, expected := range tests {
		result := SplitWords(input)
		// t.Logf("SplitWords(%s) = %v", input, result)
		if !equalSlice(result, expected) {
			t.Errorf("SplitWords(%s) = %v; want %v", input, result, expected)
		}
	}
}

// Helper function to compare slices
func equalSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

var slugifyLetters = "(?:[a-zA-Z]+)"
var slugifyNumbers = "(?:[0-9]+)"
var slugifyAlphanumeric = "(?:[a-zA-Z0-9]+)"
var slugifyCamelCase = "(?:[A-Z]?[a-z]{2,})"
var slugifyUpperCase = "(?:[A-Z]+)"

var slugifyRe = regexp.MustCompile(fmt.Sprintf("(%s|%s|%s)", slugifyCamelCase, slugifyUpperCase, slugifyNumbers))

func SplitWordsRegExp(str string) (words []string) {
	// normalized := norm.NFKD.String(str)
	for _, match := range slugifyRe.FindAllString(str, -1) {
		words = append(words, strings.ToLower(match))
	}
	return
}

func BenchmarkSplitWordsRegExp(b *testing.B) {
	input := "ThisIsATestStringWithVariousCases123"
	for n := 0; n < b.N; n++ {
		SplitWordsRegExp(input)
	}
}

func BenchmarkSplitWords(b *testing.B) {
	input := "ThisIsATestStringWithVariousCases123"
	for n := 0; n < b.N; n++ {
		SplitWords(input)
	}
}
