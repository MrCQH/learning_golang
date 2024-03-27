package main

import "net"

type Args struct {
	N, M int
}

func (a *Args) Multiply(args *Args, reply *int) net.Error {
	*reply = args.N * args.M
	return nil
}
