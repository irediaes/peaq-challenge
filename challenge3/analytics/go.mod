module github.com/ebikode/peaq-challenge/challenge3/analytics

go 1.14

replace github.com/ebikode/peaq-challenge/challenge3/exchange => ../exchange

require (
	github.com/ebikode/peaq-challenge/challenge3/exchange v0.0.0-20201126173304-e61921fda5bd
	github.com/go-chi/render v1.0.1
	golang.org/x/sys v0.0.0-20201126144705-a4b67b81d3d2 // indirect
	google.golang.org/grpc v1.33.2
)
