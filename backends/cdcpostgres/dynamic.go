package cdcpostgres

import (
	"context"

	"github.com/batchcorp/plumber-schemas/build/go/protos/opts"
	"github.com/batchcorp/plumber/dynamic"
	"github.com/batchcorp/plumber/types"
)

func (c *CDCPostgres) Dynamic(ctx context.Context, dynamicOpts *opts.DynamicOptions, dynamicSvc dynamic.IDynamic) error {
	return types.NotImplementedErr
}
