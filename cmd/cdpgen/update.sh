[[ -d protodef ]] || mkdir protodef
(cd protodef;
    curl -s "https://chromium.googlesource.com/chromium/src/+/master/third_party/WebKit/Source/core/inspector/browser_protocol.json?format=TEXT" \
        | base64 -D >js_protocol.json
    curl -s "https://chromium.googlesource.com/v8/v8.git/+/master/src/inspector/js_protocol.json?format=TEXT" \
        | base64 -D >browser_protocol.json
)