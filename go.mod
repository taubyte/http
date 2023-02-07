module github.com/taubyte/http

go 1.18

// Taubyte Direct Imports
require github.com/taubyte/utils v0.1.1

// Direct Imports
require (
	github.com/CAFxX/httpcompression v0.0.8
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.4.2
	github.com/ipfs/go-log/v2 v2.5.1
	github.com/rs/cors v1.8.2
	github.com/spf13/afero v1.9.2
	github.com/unrolled/secure v1.0.9

)

// Indirect Imports
require (
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/klauspost/compress v1.14.1 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.19.1 // indirect
	golang.org/x/sys v0.1.0 // indirect
	golang.org/x/text v0.4.0 // indirect
)

require (
	github.com/stretchr/testify v1.7.4 // indirect
	go.uber.org/goleak v1.1.12 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
