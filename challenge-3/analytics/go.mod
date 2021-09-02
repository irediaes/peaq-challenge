module github.com/ebikode/peaq-challenge/challenge3/analytics

go 1.14

// replace github.com/ebikode/peaq-challenge/challenge3/exchange => ../exchange
// replace (
// 	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
// 	google.golang.org/grpc => google.golang.org/grpc v1.27.0
// )

require (
	github.com/ebikode/peaq-challenge/challenge3/exchange v0.0.0-20201127184601-61eb53bfcb3a
	// github.com/ebikode/peaq-challenge/challenge3/exchange v0.0.0-20201126173304-e61921fda5bd
	golang.org/x/sys v0.0.0-20201126233918-771906719818 // indirect
	google.golang.org/grpc v1.33.2
)
