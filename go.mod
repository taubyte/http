module github.com/taubyte/http

go 1.18

replace (
	bitbucket.org/taubyte/auth => ../auth
	bitbucket.org/taubyte/billing => ../billing
	bitbucket.org/taubyte/config-compiler => ../config-compiler
	bitbucket.org/taubyte/dreamland => ../dreamland
	bitbucket.org/taubyte/go-node-counters => ../go-node-counters
	bitbucket.org/taubyte/go-node-database => ../go-node-database
	bitbucket.org/taubyte/go-node-http => ../go-node-http
	bitbucket.org/taubyte/go-node-ipfs => ../go-node-ipfs
	bitbucket.org/taubyte/go-node-p2p => ../go-node-p2p
	bitbucket.org/taubyte/go-node-pubsub => ../go-node-pubsub
	bitbucket.org/taubyte/go-node-smartops => ../go-node-smartops
	bitbucket.org/taubyte/go-node-storage => ../go-node-storage
	bitbucket.org/taubyte/go-node-tvm => ../go-node-tvm
	bitbucket.org/taubyte/hoarder => ../hoarder
	bitbucket.org/taubyte/http-auto => ../http-auto
	bitbucket.org/taubyte/kvdb => ../kvdb
	bitbucket.org/taubyte/monkey => ../monkey
	bitbucket.org/taubyte/node => ../node
	bitbucket.org/taubyte/p2p => ../p2p
	bitbucket.org/taubyte/patrick => ../patrick
	bitbucket.org/taubyte/seer => ../seer
	bitbucket.org/taubyte/seer-p2p-client => ../seer-p2p-client
	bitbucket.org/taubyte/tns => ../tns
	bitbucket.org/taubyte/tns-p2p-client => ../tns-p2p-client
	bitbucket.org/taubyte/vm-test-examples => ../vm-test-examples
	github.com/taubyte/go-interfaces => ../go-interfaces
	github.com/taubyte/go-sdk => ../go-sdk
	github.com/taubyte/go-sdk-symbols => ../go-sdk-symbols
	github.com/taubyte/go-specs => ../go-specs
	github.com/taubyte/http => ../http
	github.com/taubyte/utils => ../utils
	github.com/taubyte/vm => ../vm
	github.com/taubyte/vm-plugins => ../vm-plugins
	github.com/taubyte/vm-wasm-utils => ../vm-wasm-utils
)

// Taubyte Direct Imports
require (
	github.com/taubyte/go-interfaces v0.1.3
	github.com/taubyte/utils v0.1.5
)

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
	go.uber.org/goleak v1.1.12 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.8.0 // indirect
)
