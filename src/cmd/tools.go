// +build tools
package main

import (
	_ "github.com/FiloSottile/mkcert"
	_ "github.com/FiloSottile/age/cmd/age"
	_ "9fans.net/go/acme/acmego"
	_ "9fans.net/go/acme/Dict"
)
