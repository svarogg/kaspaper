module github.com/svarogg/kaspaper

go 1.17

require github.com/kaspanet/kaspad v0.11.6

require (
	github.com/btcsuite/go-socks v0.0.0-20170105172521-4720035b7bfd // indirect
	github.com/btcsuite/winsvc v1.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/jessevdk/go-flags v1.4.0 // indirect
	github.com/jrick/logrotate v1.0.0 // indirect
	github.com/kaspanet/go-muhash v0.0.4 // indirect
	github.com/kaspanet/go-secp256k1 v0.0.7 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/syndtr/goleveldb v1.0.1-0.20190923125748-758128399b1d // indirect
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a // indirect
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	golang.org/x/sys v0.0.0-20210525143221-35b2ab0089ea // indirect
	golang.org/x/text v0.3.5 // indirect
	google.golang.org/genproto v0.0.0-20210604141403-392c879c8b08 // indirect
	google.golang.org/grpc v1.38.0 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
)

replace (
	github.com/kaspanet/kaspad latest => ../../kaspanet/kaspad
)