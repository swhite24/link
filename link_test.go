/**
 * link_test.go
 *
 * (C) Steven White 2015
 */

package link

import (
	"net/http"
	"testing"
)

func TestReadSingle(t *testing.T) {
	r, _ := http.NewRequest("GET", "", nil)
	r.Header.Add("Link", "<https://api.github.com/search/code?q=addClass+user%3Amozilla&page=2>; rel=\"next\"")

	g := Parse(r)

	if len(g) != 1 {
		t.Errorf("Incorrent number of links parsed")
	}
	if g["next"] == nil {
		t.Errorf("Unable to parse next link")
	}
}

func TestReadMultiple(t *testing.T) {
	r, _ := http.NewRequest("GET", "", nil)
	r.Header.Add("Link", "<https://api.github.com/search/code?q=addClass+user%3Amozilla&page=2>; rel=\"next\""+
		", <https://api.github.com/search/code?q=addClass+user%3Amozilla&page=34>; rel=\"last\"")

	g := Parse(r)

	if len(g) != 2 {
		t.Errorf("Incorrent number of links parsed")
	}
	if g["next"] == nil {
		t.Errorf("Unable to parse next link")
	}
	if g["last"] == nil {
		t.Errorf("Unable to parse last link")
	}
}

func TestNoLink(t *testing.T) {
	r, _ := http.NewRequest("GET", "", nil)
	g := Parse(r)

	if g != nil {
		t.Error("Failed to deliver nil for empty Link header")
	}
}
