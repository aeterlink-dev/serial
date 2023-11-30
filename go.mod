module github.com/aeterlink-dev/go-serial-al

go 1.20

replace (
	github.com/aeterlink-dev/go-serial-al/enumrator => ./enumrator
	github.com/aeterlink-dev/go-serial-al/portlist => ./portlist
	github.com/aeterlink-dev/go-serial-al/serial => ./
	github.com/aeterlink-dev/go-serial-al/unixutils => ./unixutils
)

require (
	github.com/aeterlink-dev/go-serial-al/serial v0.0.0-00010101000000-000000000000
	github.com/creack/goselect v0.1.2
	github.com/stretchr/testify v1.8.4
	golang.org/x/sys v0.19.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
