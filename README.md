# go-import-path
This is to be used if you have a custom domain you want to use as your base import path for your Go projects.

# usage
Figure out how you want to store the key/value pairs that serve as your router. For simplicity, this example just uses a map. The keys in the map correspond to the request uri (r.RequestURI) in the incoming request, and the vaue represents the target of where that repo should live.

For example, with a custom domain of `go.example.com` and the router used here:
```go
router = map[string]string{
    "/package1": "https://github.com/jrkt/package1",
    "/package2": "https://github.com/jrkt/package2",
    "/package3": "https://gitlab.com/jrkt/package3",
}
```

- `go get go.example.com/{package1,package2}` would both resolve to github
- `go get go.example.com/package3` would resolve to gitlab.

