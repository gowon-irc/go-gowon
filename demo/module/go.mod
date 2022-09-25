module github.com/gowon-irc/go-gowon/demo/module

go 1.19

replace github.com/gowon-irc/go-gowon => ../..

require (
	github.com/eclipse/paho.mqtt.golang v1.3.5
	github.com/gowon-irc/go-gowon v0.0.0-00010101000000-000000000000
	github.com/gowon-irc/gowon v0.0.0-20211012014610-ece6c2510654
	github.com/jessevdk/go-flags v1.5.0
)

require (
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/net v0.0.0-20200425230154-ff2c4b7c35a0 // indirect
	golang.org/x/sys v0.0.0-20210320140829-1e4c9ba3b0c4 // indirect
)
