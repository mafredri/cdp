# cdpgen

The cdpgen tool is used to generate the Golang API for the Chrome Debugging Protocol from the protocol definitions (JSON).

Beware, this tool is not a feat of engineering, it's only purpose is to generate the Golang API. It has gone thourgh many revisions while prototyping the API and might contain both messy and dead code.

## Installing

```console
go get -u github.com/mafredri/cdp/cmd/cdpgen
```

## Usage

The current protocol definitions are committed in this repository under the `protodef` directory.

### Generating the cdp package

```console
$ cdpgen -dest-pkg github.com/mafredri/cdp \
    -browser-proto $GOPATH/src/github.com/mafredri/cdp/cmd/cdpgen/protodef/browser_protocol.json \
    -js-proto $GOPATH/src/github.com/mafredri/cdp/cmd/cdpgen/protodef/js_protocol.json
```

### Updating protocol definitions

```console
$ ./update.sh
```

## Future improvements

- Better formatting for comments, consider sentence construction, proper casing and line length.
