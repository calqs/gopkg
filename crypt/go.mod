module github.com/calqs/crypt

go 1.24.5

require golang.org/x/crypto v0.41.0

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require (
	github.com/calqs/http v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.10.0
	golang.org/x/sys v0.35.0 // indirect
)

replace github.com/calqs/http => ../http
