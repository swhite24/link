/**
 * link.go
 *
 * (C) Steven White 2015
 */

package link

import (
	"net/http"
	"regexp"
)

type (
	// Group is delivered by Parse, contains multiple links indexed by
	// "rel"
	Group map[string]*Link

	// Link contains an individual Link item, with all non-URI components
	// placed into Extra.
	Link struct {
		URI   string
		Extra map[string]string
	}
)

var (
	linkSplit  = regexp.MustCompile("^<(.*)>;(.*)$")
	commaSplit = regexp.MustCompile(", *")
	semiSplit  = regexp.MustCompile("; *")
	equalSplit = regexp.MustCompile(" *= *")
	keySplit   = regexp.MustCompile("[a-z*]+")
	valSplit   = regexp.MustCompile("\"*([^\"]*)\"*")
)

// Parse inspects the provided *http.Request for a Link header and delivers
// a parsed Group
func Parse(r *http.Request) Group {
	header := r.Header.Get("Link")
	if len(header) == 0 {
		return nil
	}
	group := Group{}
	links := commaSplit.Split(header, -1)
	for _, link := range links {
		l := &Link{Extra: map[string]string{}}
		pieces := linkSplit.FindAllStringSubmatch(link, -1)[0]
		// for _, pieces := range m {
		l.URI = pieces[1]
		extras := pieces[2]
		for _, extra := range semiSplit.Split(extras, -1) {
			vals := equalSplit.Split(extra, -1)
			key := keySplit.FindString(vals[0])
			val := valSplit.FindStringSubmatch(vals[1])[1]

			l.Extra[key] = val
			if key == "rel" {
				group[val] = l
			}
		}
	}
	return group
}
