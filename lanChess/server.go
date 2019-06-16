package main

import (
	"SchorChessServer/player"
	"SchorChessServer/util"
	"net"
)

func main() {

	ln, err := net.Listen("tcp", "10.0.0.106:8080")
	util.Check(err, "Server is ready.")

	for {
		conn, err := ln.Accept()
		util.Check(err, "Accepted connection.")

		// Spawn new player thread
		go player.Player(conn)
	}

}
