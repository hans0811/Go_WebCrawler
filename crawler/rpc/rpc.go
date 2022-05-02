// go also has package which called rpc
package rpcdemo

import "errors"

// Service.Method

type DemoService struct{}

type Args struct {
	A, B int
}

// rpc need 2 arguments, args and result
func (DemoService) Div(args Args, result *float64) error{

	if args.B == 0{
		return errors.New("division by zero")
	}

	*result = float64(args.A)/float64(args.B)
	return nil
}