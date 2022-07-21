// Package godoc is a test of go doc.
//
//
// The godoc package should only be used for test go doc, this is
// package comment
package godoc

// ddd
//
// # ddd
//
// ddd
const DDD = "ddd"

// An Op is a single regular expression operator.
//
//
// kkkk
// # ddd
//go:generate stringer -type Op -trimprefix Op
type Op uint8

// Search uses binary search...
//
// As a more whimsical example, this program guesses your number:
//
// func GuessingGame() {
//var s string
//fmt.Printf("Pick an integer from 0 to 100.\n")
//answer := sort.Search(100, func(i int) bool {
//fmt.Printf("Is your number <= %d? ", i)
//fmt.Scanf("%s", &s)
//return s != "" && s[0] == 'y'
//})
//fmt.Printf("Your number is %d.\n", answer)
//}
//
func Search(n int, f func(int) bool) int {
	return 0
}

// Match reports whether name matches the shell pattern.
// The pattern syntax is:
//  pattern:
//      { term }
//  term:
//      '*'         matches any sequence of non-/ characters
//      '?'         matches any single non-/ character
//      '[' [ '^' ] { character-range } ']'
//                  character class (must be non-empty)
//      c           matches character c (c != '*', '?', '\\', '[')
//      '\\' c      matches character c
//
//  character-range:
//      c           matches character c (c != '\\', '-', ']')
//      '\\' c      matches character c
//      lo '-' hi   matches character c for lo <= c <= hi
//
// Match requires pattern to match all of name, not just a substring.
// The only possible returned error is [ErrBadPattern], when pattern
// is malformed.
func Match(pattern, name string) (matched bool, err error) {
	return
}

// PublicSuffixList provides the public suffix of a domain. For example:
//
//   - the public suffix of "example.com" is "com",
//   - the public suffix of "foo1.foo2.foo3.co.uk" is "co.uk", and
//   - the public suffix of "bar.pvt.k12.ma.us" is "pvt.k12.ma.us".
// Implementations of PublicSuffixList must be safe for concurrent use by
// multiple goroutines.
//
// An implementation that always returns "" is valid and may be useful for
// testing but it is not secure: it means that the HTTP server for foo.com can
// set a cookie for bar.com.
//
// A public suffix list implementation is in the package
// golang.org/x/net/publicsuffix.
type PublicSuffixList interface {
}
