package app

import "context"

// Usecase ...
type Usecase interface {
	HelloWorld(ctx context.Context)
}
