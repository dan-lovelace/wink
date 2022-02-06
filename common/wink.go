package common

import (
	"context"

	"github.com/dan-lovelace/wink/configs"
)

type Wink struct {
	Config  configs.Config
	Context context.Context
}
