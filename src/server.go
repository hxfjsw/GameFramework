package main

import (
	"net"
	"strconv"
)

type Server struct {
	ln net.Listener
}

func (this *Server)  Run(ip string, port uint32) {
	if ln, err := net.Listen("tcp", ip + ":" + strconv.Itoa(int(port))); err != nil {
		panic(err)
	} else {
		this.ln = ln
	}

	for {
		if conn, err := this.ln.Accept(); err != nil {
			conn.Close()
			panic(err)
		} else {
			vm.Run("onConnect(" + conn.LocalAddr().String() + "); ")
			go this.handle(conn)
		}

	}

}

func (this *Server) handle(conn net.Conn) {
	sesssion := hall.Join(conn)
	buffer := make([]byte, 4069)
	for {
		if n, err := conn.Read(buffer); err != nil {
			hall.Leave(sesssion)
			vm.Run("OnClose(" + strconv.Itoa(sesssion.Fd) + ");")
		} else {
			if (n == 0) {
				hall.Leave(sesssion)
				vm.Run("OnClose(" + strconv.Itoa(sesssion.Fd) + ");")
			} else {
				msg := string(buffer[:n])
				vm.Run("OnMessage(" + strconv.Itoa(sesssion.Fd) + "," + msg + ")")
			}
		}
	}
}