module github.com/gowon-irc/go-gowon/demo/module

go 1.21.4

toolchain go1.22.0

replace github.com/gowon-irc/go-gowon => ../..

require (
	github.com/eclipse/paho.mqtt.golang v1.4.3
	github.com/gowon-irc/go-gowon v0.0.0-20220719115350-ec869e1addf7
	github.com/gowon-irc/gowon v0.0.0-20240225235519-41763101318b
	github.com/jessevdk/go-flags v1.5.0
)

require (
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sync v0.5.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
)
