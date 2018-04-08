#!/usr/bin/env bash

[[ -d protodef ]] || mkdir protodef
(cd protodef;
    curl -sSL "https://github.com/ChromeDevTools/devtools-protocol/raw/master/json/browser_protocol.json" >js_protocol.json
    curl -sSL "https://github.com/ChromeDevTools/devtools-protocol/raw/master/json/js_protocol.json" >browser_protocol.json
)
