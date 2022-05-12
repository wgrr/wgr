//go:build tools
package tools

import (
	_ "9fans.net/go/acme/acmego"
	_ "9fans.net/go/acme/editinacme"
	_ "9fans.net/go/acme/Watch"

	_ "filippo.io/age/cmd/age"
	_ "filippo.io/age/cmd/age-keygen"

	_ "golang.org/x/tools/cmd/goimports"
	_ "golang.org/x/tools/cmd/stringer"
	_ "golang.org/x/tools/godoc"
	_ "golang.org/x/tools/gopls"

	_ "honnef.co/go/tools/cmd/staticcheck"
	_ "mvdan.cc/gofumpt"

	_ "robpike.io/ivy"
)

