package common

import (
	"context"
	"io"

	"github.com/dan-lovelace/wink/configs"
)

type Wink struct {
	Config  *configs.Config
	Context context.Context
	Out     io.Writer
}
