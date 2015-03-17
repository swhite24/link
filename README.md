# link

Parses, generates `"Link"` headers used for pagination, as defined in
[rfc5988](http://tools.ietf.org/html/rfc5988).

## Installation

```
go get github.com/swhite24/link
```

## Usage

```go

r, _ := http.NewRequest("GET", "", nil)
r.Header.Add("Link", "<https://api.github.com/search/code?q=addClass+user%3Amozilla&page=2>; rel=\"next\""+
	", <https://api.github.com/search/code?q=addClass+user%3Amozilla&page=34>; rel=\"last\"")

group := Parse(r)

for rel, link := range group {
    fmt.Println("rel:", rel)
    fmt.Println("uri:", link.URI)
    fmt.Println("link title", link.Extra["title"])
}
```

## License

See [LICENSE](LICENSE).
