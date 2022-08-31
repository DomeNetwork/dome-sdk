#!/bin/bash

# Setup the environment default to `dev`.
env=$2
if [ -z "$env" ]; then
    env="dev"
fi

function build {
    docker build \
        -f config/Dockerfile.dev \
        -t domenetwork/$1:$env \
        --target $1 \
        .
}

case $1 in
    "docs")
        # Generate the documentation using go doc.
        # Note: godoc can be used to server the documentation.
        # go doc -all ./...
        [[ $OSTYPE == 'darwin'* ]] && \
            open -n -a "Google Chrome" --args "--new-window" "http://localhost:6060/pkg/github.com/domenetwork/dome-sdk/"
        godoc -http=:6060
        ;;
    "wasm")
        # Build the WASM module.
        [[ -d "build/" ]] && rm -fR build
        [[ ! -d "build/" ]] && mkdir -p build

        CGO_ENABLED=0 GOARCH=wasm GOOS=js go build -o "build/dome.wasm" wasm/main.go
        [[ ! -f "build/wasm_exec.js" ]] && \
            cp -f "$(go env GOROOT)/misc/wasm/wasm_exec.js" "build/wasm_exec.js"

        # JS testing
        cp -f "build/dome.wasm" "ui/test/dome.wasm"
        cp -f "build/wasm_exec.js" "ui/test/wasm_exec.js"

        # DOME wallet
        cp -f "build/dome.wasm" "ui/wallet/public/dome.wasm"
        cp -f "build/wasm_exec.js" "ui/wallet/src/assets/js/wasm_exec.js"
        ;;
    *)
        echo "Unknown command, must be one of: docs or nym."
        echo "    docs  - build the docs"
        echo "    wasm  - build the WASM web module"
        ;;
esac
