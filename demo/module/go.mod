module github.com/gowon-irc/go-gowon/demo/module

go 1.24.0

replace github.com/gowon-irc/go-gowon => ../..

require (
	github.com/eclipse/paho.mqtt.golang v1.5.1
	github.com/gowon-irc/go-gowon v0.0.0-00010101000000-000000000000
	github.com/gowon-irc/gowon v0.0.0-20211012014610-ece6c2510654
	github.com/jessevdk/go-flags v1.5.0
)

require (
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/net v0.44.0 // indirect
	golang.org/x/sync v0.17.0 // indirect
	golang.org/x/sys v0.36.0 // indirect
)
