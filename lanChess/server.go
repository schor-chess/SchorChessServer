package main

import (
	"net"
	"wds/SchorChessServer/player"
	"wds/SchorChessServer/util"
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
