module github.com/calqs/gopkg

go 1.25.0

replace github.com/calqs/gopkg/router => ./router

replace github.com/calqs/gopkg/crypt => ./crypt

replace github.com/calqs/gopkg/dt => ./dt

require github.com/calqs/gopkg/dt v0.0.0-20250813215151-bc3b58d35fe9 // indirect
