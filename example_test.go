package sshsig_test

import (
	"bytes"
	"fmt"
	"net"
	"os"

	"github.com/42wim/sshsig"
	"golang.org/x/crypto/ssh/agent"
)

func ExampleSignWithAgent() {
	// This example will panic when you don't have a ssh-agent running.
	conn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	if err != nil {
		panic(err)
	}

	ag := agent.NewClient(conn)

	// This public key must match in your agent (use `ssh-add -L` to get the public key)
	pubkey := []byte(`ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIAo3D7CGN01tTYY/dLKXEv8RxRyxa32c51X0uKMhnMab wim@localhost`)
	//
	data := []byte("hello world")

	res, err := sshsig.SignWithAgent(pubkey, ag, bytes.NewBuffer(data), "file")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))
}

func ExampleSign() {
	privkey := []byte(`-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
QyNTUxOQAAACCOjP6i4Pm/pYAAmpAMNZ6xrbHl9RW8xdul6kzIWuKMMAAAAIhoQm34aEJt
+AAAAAtzc2gtZWQyNTUxOQAAACCOjP6i4Pm/pYAAmpAMNZ6xrbHl9RW8xdul6kzIWuKMMA
AAAEBfIl93TLj6qHeg37GnPuZ00h8OVv1mzlhy0rhuO4Y0do6M/qLg+b+lgACakAw1nrGt
seX1FbzF26XqTMha4owwAAAAAAECAwQF
-----END OPENSSH PRIVATE KEY-----`)

	data := []byte("hello world")

	res, err := sshsig.Sign(privkey, bytes.NewBuffer(data), "file")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))

	// Output:
	// -----BEGIN SSH SIGNATURE-----
	// U1NIU0lHAAAAAQAAADMAAAALc3NoLWVkMjU1MTkAAAAgjoz+ouD5v6WAAJqQDDWesa2x5f
	// UVvMXbpepMyFrijDAAAAAEZmlsZQAAAAAAAAAGc2hhNTEyAAAAUwAAAAtzc2gtZWQyNTUx
	// OQAAAEBeu9Z+vLxBORysiqEbTzJP0EZKG0/aE5HpTtvimjQS6mHZCAGFg+kimNatBE0Y1j
	// gS4pfD73TlML1SyB5lb/YO
	// -----END SSH SIGNATURE-----
}
