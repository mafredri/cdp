#!/usr/bin/env bash

base64_opts=(-D)
# Feature detect if base64 uses a -D or -d flag for decoding, macOS normally uses -D.
if ! base64 -D <<<"" 1>/dev/null 2>&1; then
	base64_opts=(-d)
fi

[[ -d protodef ]] || mkdir protodef
(cd protodef;
    curl -s "https://chromium.googlesource.com/chromium/src/+/master/third_party/WebKit/Source/core/inspector/browser_protocol.json?format=TEXT" \
        | base64 ${base64_opts[@]} >js_protocol.json
    curl -s "https://chromium.googlesource.com/v8/v8.git/+/master/src/inspector/js_protocol.json?format=TEXT" \
        | base64 ${base64_opts[@]} >browser_protocol.json
)
