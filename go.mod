module github.com/hofstadter-io/mvs

go 1.14

require (
	cuelang.org/go v0.0.15
	github.com/bmatcuk/doublestar v1.2.2
	github.com/go-git/go-billy/v5 v5.0.0
	github.com/go-git/go-git/v5 v5.0.0
	github.com/google/go-github/v30 v30.1.0
	github.com/parnurzeal/gorequest v0.2.16
	github.com/spf13/cobra v0.0.6
	github.com/spf13/viper v1.4.0 // indirect
	golang.org/x/mod v0.2.0
	golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be
	golang.org/x/xerrors v0.0.0-20191011141410-1b5146add898
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c
	moul.io/http2curl v1.0.0 // indirect
)

// replace github.com/hofstadter-io/go-utils => ../go-utils
