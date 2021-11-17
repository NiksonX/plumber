package rstreams

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"

	"github.com/batchcorp/plumber-schemas/build/go/protos/opts"

	"github.com/batchcorp/plumber/dynamic"
)

func (r *RedisStreams) Dynamic(ctx context.Context, dynamicOpts *opts.DynamicOptions) error {
	if err := validateDynamicOptions(dynamicOpts); err != nil {
		return errors.Wrap(err, "unable to validate dynamic options")
	}

	// Start up dynamic connection
	grpc, err := dynamic.New(dynamicOpts, BackendName)
	if err != nil {
		return errors.Wrap(err, "could not establish connection to Batch")
	}

	go grpc.Start()

	for {
		select {
		case outbound := <-grpc.OutboundMessageCh:
			for _, streamName := range dynamicOpts.RedisStreams.Args.Streams {
				_, err := r.client.XAdd(ctx, &redis.XAddArgs{
					Stream: streamName,
					ID:     dynamicOpts.RedisStreams.Args.WriteId,
					Values: map[string]interface{}{
						dynamicOpts.RedisStreams.Args.Key: outbound.Blob,
					},
				}).Result()
				if err != nil {
					r.log.Errorf("unable to write message to '%s': %s", streamName, err)
					continue
				}

				r.log.Infof("Successfully wrote message to stream '%s' with key '%s' for replay '%s'",
					streamName, dynamicOpts.RedisStreams.Args.Key, outbound.ReplayId)
			}
		case <-ctx.Done():
			r.log.Warning("context cancelled")
			return nil
		}
	}

	return nil
}

func validateDynamicOptions(dynamicOpts *opts.DynamicOptions) error {
	if dynamicOpts == nil {
		return errors.New("write options cannot be nil")
	}

	if dynamicOpts.RedisStreams == nil {
		return errors.New("backend group options cannot be nil")
	}

	if dynamicOpts.RedisStreams.Args == nil {
		return errors.New("backend arg options cannot be nil")
	}

	if len(dynamicOpts.RedisStreams.Args.Streams) == 0 {
		return errors.New("you must specify at least one stream to write to")
	}

	return nil
}