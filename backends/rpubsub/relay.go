package rpubsub

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/relistan/go-director"
	"github.com/sirupsen/logrus"

	rtypes "github.com/batchcorp/plumber/backends/rpubsub/types"
	"github.com/batchcorp/plumber/options"
	"github.com/batchcorp/plumber/stats"
	"github.com/batchcorp/plumber/types"
)

const (
	RetryReadInterval = 5 * time.Second
)

type Relayer struct {
	Client      *redis.Client
	Options     *options.Options
	RelayCh     chan interface{}
	log         *logrus.Entry
	Looper      *director.FreeLooper
	ShutdownCtx context.Context
}

var (
	ErrMissingChannel = errors.New("You must specify at least one channel")
)

// Relay sets up a new RedisPubSub relayer
func (r *Redis) Relay(ctx context.Context, relayCh chan interface{}, errorCh chan *types.ErrorMessage) error {
	if err := validateRelayOptions(r.Options); err != nil {
		return errors.Wrap(err, "unable to verify options")
	}

	// Relay reads messages from RedisPubSub and sends them to RelayCh which is then read by relay.Run()
	r.log.Infof("Relaying RedisPubSub messages from %d channel(s) (%s) -> '%s'",
		len(r.Options.RedisPubSub.Channels), r.Options.RedisPubSub.Channels, r.Options.Relay.GRPCAddress)

	r.log.Infof("HTTP server listening on '%s'", r.Options.Relay.HTTPListenAddress)

	sub := r.client.Subscribe(ctx, r.Options.RedisPubSub.Channels...)
	defer sub.Unsubscribe(ctx, r.Options.RedisPubSub.Channels...)

	msgCh := sub.Channel()

	for {
		// Redis library is not handling context cancellation, only timeouts. So we must use a timeout here
		// to ensure we eventually receive the context cancellation from ShutdownCtx
		select {
		case <-ctx.Done():
			r.log.Info("Received shutdown signal, exiting relayer")
			return nil
		case msg := <-msgCh:
			stats.Incr("redis-pubsub-relay-consumer", 1)

			r.log.Debugf("Relaying message received on channel '%s' to Batch (contents: %s)",
				msg.Channel, msg.Payload)

			relayCh <- &rtypes.RelayMessage{
				Value:   msg,
				Options: &rtypes.RelayMessageOptions{},
			}
		default:
			// NOOP
		}
	}

	r.log.Debug("relayer exiting")

	return nil
}

// validateRelayOptions ensures all required CLI options are present before initializing relay mode
func validateRelayOptions(opts *options.Options) error {
	if len(opts.RedisPubSub.Channels) == 0 {
		return ErrMissingChannel
	}

	// RedisPubSub either supports a password (v1+) OR a username+password (v6+)
	if opts.RedisPubSub.Username != "" && opts.RedisPubSub.Password == "" {
		return errors.New("missing password (either use only password or fill out both)")
	}

	return nil
}
