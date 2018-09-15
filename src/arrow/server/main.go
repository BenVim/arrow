package server

import (
	"arrow/log"
	"arrow/util"
	"math/rand"
	"crypto/tls"
	"net"
	"arrow/conn"
	"time"
)

// GLOBALS
var (
	opts *Options
)

func Main() {

	opts = parseArgs()
	log.LogTo(opts.logTo, opts.logLevel)

	seed, err := util.RandomSeed()

	if err != nil {
		panic(err)
	}
	rand.Seed(seed)

	tlsConfig, err := LoadTLSConfig(opts.tlsCrt, opts.tlsKey)

	if err != nil {
		panic(err)
	}

	createControllerConnection(opts.tunnelAddr, tlsConfig)
}


//创建控制连接
func createControllerConnection(addr string, tlsCfg *tls.Config) {

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}

	go func() {
		for {
			conn, err := listen.Accept()

			if err!=nil {
				log.Error("error conn", err)
				continue
			}
			//每个接入的客户端
			go handleTunnelController(conn)
		}
	}()
}

//处理每个接入的用户
//负责验证身份
func handleTunnelController(conn net.Conn)  {

	defer func() {
		conn.Close()
	}()

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))






}
