# mimetype

The ultimate golang content-type utility.

## Install

To start using mimetype, install Go and run `go get`:

```sh
$ go get -u github.com/funnyfactor/mimetype
```

## Adding Types

All mime types are based on [mime-db](https://www.npmjs.com/package/mime-db),
so open a PR there if you'd like to add mime types.

## API Usage Examples

### Get File Extension from MIME Type or Content Type

```go
package main

import (
    "fmt"
    "github.com/funnyfactor/mimetype"
)

func main() {
    // Get the default file extension for a MIME type
    ext := mimetype.ExtensionByType("image/jpeg")
    fmt.Println(ext) // Output: jpg

    // Handle MIME type with charset
    ext = mimetype.ExtensionByType("text/html; charset=utf-8")
    fmt.Println(ext) // Output: html

    // Handle non-existent MIME type
    ext = mimetype.ExtensionByType("unknown/type")
    fmt.Println(ext) // Output: ""
}
```

### Get MIME Type from File Extension

```go
package main

import (
    "fmt"
    "github.com/funnyfactor/mimetype"
)

func main() {
    // Get MIME type for a file extension
    mimeType := mimetype.TypeByExtension("jpg")
    fmt.Println(mimeType) // Output: image/jpeg

    // Support extension with leading dot
    mimeType = mimetype.TypeByExtension(".png")
    fmt.Println(mimeType) // Output: image/png

    // Handle non-existent extension
    mimeType = mimetype.TypeByExtension("unknown")
    fmt.Println(mimeType) // Output: ""
}
```

## License

[MIT](LICENSE)
