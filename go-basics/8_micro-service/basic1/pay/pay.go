package pay

import "net/rpc"

type Payment struct{}

func main() {
	rpc.Register("Payment")
}
