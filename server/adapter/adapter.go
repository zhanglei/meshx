package adapter

import (
	"github.com/smallnest/meshx/util"
)

// Adapter is the server agent interface of sidecar which implements protocol of rpc frameworks.
type Adapter interface {
	Start(op util.Option) error
	Close() error
}
