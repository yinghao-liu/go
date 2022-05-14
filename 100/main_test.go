package main

import (
	"fmt"
	"net/url"
	"testing"
)

func TestURL(t *testing.T) {
	var u url.URL
	fmt.Printf("path is [%s]\n", u.Path)
	fmt.Printf("host is [%s]\n", u.Host)
}
