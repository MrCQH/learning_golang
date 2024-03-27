package rpc_objects

type Args struct {
	N, M int
}

/*
Golang RPC格式：
1. 它是导出的（Multiply的M是大写的）。
2. 它有两个参数，且第二个参数是一个指针。
3. 它的返回类型是error。
*/

func (a *Args) Multiply(args *Args, res *int) error {
	*res = args.N * args.M
	return nil
}
