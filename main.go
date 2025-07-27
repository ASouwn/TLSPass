package main

import (
	"os"

	"github.com/ASouwn/TLSPass/cmds"
)

func main() {
	if len(os.Args) < 2 {
		println("No command provided.")
		return
	}
	cmd := os.Args[1] // 获取子命令

	switch cmd {
	case "start":
		cmds.Start()
	case "help":
		println(
			"TLSPass is to proxy TLS requests to a local server.\n",
			"Server will listen on port 443 with TLS. Please ensure you have the correct certificate and key files.\n",
			"certPath: /etc/TLSPass/tlspass.pem\n",
			"keyPath: /etc/TLSPass/tlspass.key\n",
			"Usage:\n",
			"  start - Start the TLSPass server with TLS.\n",
		)
	default:
		println("Unknown command:", cmd)
	}
}
