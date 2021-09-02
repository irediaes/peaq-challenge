module github.com/ebikode/peaq-challenge/challenge-3/analytics

go 1.14

// replace github.com/ebikode/peaq-challenge/challenge3/exchange => ../exchange
// replace (
// 	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
// 	google.golang.org/grpc => google.golang.org/grpc v1.27.0
// )

require (
	github.com/ebikode/peaq-challenge/challenge-3/exchange v0.0.0-20210902180042-f969ab4aafd0 // indirect
	github.com/whiteshtef/clockwork v0.0.0-20200221012748-027e62affd84 // indirect
	// github.com/ebikode/peaq-challenge/challenge3/exchange v0.0.0-20201126173304-e61921fda5bd
	golang.org/x/sys v0.0.0-20201126233918-771906719818 // indirect
	google.golang.org/grpc v1.33.2
)
