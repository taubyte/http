module github.com/taubyte/http

go 1.18

// Taubyte Direct Imports
require github.com/taubyte/utils v0.1.5

replace github.com/taubyte/go-interfaces => /home/tafkhan/Documents/Work/Taubyte/github/go-interfaces

// Direct Imports
require (
	github.com/CAFxX/httpcompression v0.0.8
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.5.0
	github.com/ipfs/go-log/v2 v2.5.1
	github.com/rs/cors v1.8.2
	github.com/spf13/afero v1.9.5
	github.com/unrolled/secure v1.0.9

)

// Indirect Imports
require (
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/klauspost/compress v1.16.4 // indirect
	github.com/mattn/go-isatty v0.0.18 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.8.0 // indirect
)

require github.com/taubyte/go-interfaces v0.0.0-00010101000000-000000000000

require go.uber.org/goleak v1.1.12 // indirect
