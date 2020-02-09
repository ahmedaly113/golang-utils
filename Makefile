test:
	go test -v -race github.com/ahmedaly113/golang-utils/maps
	go test -v -race github.com/ahmedaly113/golang-utils/sets
	go test -v -race github.com/ahmedaly113/golang-utils/netutil
	go test -v -race github.com/ahmedaly113/golang-utils/worker
	go test -v -race github.com/ahmedaly113/golang-utils/sync

test-only:
	go test -v -race github.com/ahmedaly113/golang-utils/${name}

setup:
	glide install
