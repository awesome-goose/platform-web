module github.com/awesome-goose/platform-web

go 1.25.3

require (
	github.com/awesome-goose/contracts v0.0.0
)

replace (
	github.com/awesome-goose/contracts => ../contracts
)
