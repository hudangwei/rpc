package main

import (
	"github.com/hudangwei/rpc"
	"github.com/hudangwei/rpc/example/data"
)

func main() {
	srv := rpc.NewServer()
	srv.Register(new(data.TestRpc))
	srv.Run()
}
