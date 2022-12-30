package hello

import (
	"context"
	"fmt"
)

// HelloWorld says hello with the given name
// encore:api public path=/hello/:name
func HelloWorld(ctx context.Context, name string) (Response, error) {
	return Response{fmt.Sprintf("Hello, %s!", name)}, nil
}

type Response struct {
	Greetings string
}
