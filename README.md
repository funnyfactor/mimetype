# mimetype

The ultimate javascript content-type utility.

## Install

To start using mimetype, install Go and run `go get`:

```sh
$ go get -u github.com/funnyfactor/mimetype
```

## Note on MIME Type Data and Semver

This package considers the programmatic api as the semver compatibility. Additionally, the package which provides the MIME data
for this package (`mime-db`) _also_ considers it's programmatic api as the semver contract. This means the MIME type resolution is _not_ considered
in the semver bumps.

In the past the version of `mime-db` was pinned to give two decision points when adopting MIME data changes. This is no longer true. We still update the
`mime-db` package here as a `minor` release when necessary, but will use a `^` range going forward. This means that if you want to pin your `mime-db` data
you will need to do it in your application. While this expectation was not set in docs until now, it is how the pacakge operated, so we do not feel this is
a breaking change.

If you wish to pin your `mime-db` version you can do that with overrides via your package manager of choice. See their documentation for how to correctly configure that.

## Adding Types

All mime types are based on [mime-db](https://www.npmjs.com/package/mime-db),
so open a PR there if you'd like to add mime types.

## API

## License

[MIT](LICENSE)
