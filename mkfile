E: cmd/src/E/E.go
	go install cmd/src/E/E.go

a: cmd/src/a+/main.go cmd/src/a-/main.go
	go install cmd/src/a+/main.go
	go install cmd/src/a-/main.go

cmd:
	cp cmd/+acme cmd/gg cmd/F cmd/lf cmd/rssh-agent cmd/surr $HOME/bin

