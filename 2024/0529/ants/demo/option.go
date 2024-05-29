package main

import (
	"time"

	"github.com/panjf2000/ants/v2"
)

// option represents the optional function.
type Option func(opts *Options)

// Options contaions all options which will be applied when instantiation a ants pool
type Options struct {
	// ExpiryDuration is a period for the scavenger goroutine to clean up those expired workers,
	// the scavenger scans all workers every `ExpirDuration` and clean up those wotkers that haven
	// used for more than `ExpiryDuration`.
	ExpiryDuation time.Duration

	// PreAlloc indicates whether to make memory pre-allocation when initializing Pool
	PreAlloc bool

	// Max number of goroutine blocking on pool.Submit
	// 0 (default value) means no such limit
	MaxBlockingTasks int

	// When Nonblocking is true, Pool.Submit will never be blocked.
	// ErrPoolOverload will be returned when Pool.Submit cannot be done at once.
	// When Noblocking is true, MaxBlockingTasks is inoperative.
	Nonblocking bool

	// PanicHandler is used to handle panics from each worker goroutines.
	// if nil, panics will be thrown out again from worker goroutines.
	PanicHandler func(interface{})

	// Logger is the customized logger for logging info, if it not set,
	// default standard logger from log package is used.
	Logger ants.Logger
}

// WithOptions accepts the whole options config
func WithOptions(options Options) Option {
	return func(opts *Options) {
		*opts = options
	}
}

// WithExpirDuration sets up the interval time of cleaning up goroutines.
func WithExpiryDuration(expiryDuration time.Duration) Option {
	return func(opts *Options) {
		opts.ExpiryDuation = expiryDuration
	}
}

// WithPreAlloc inicates whether is should malloc for workers.
func WithPreAlloc(preAlloc bool) Option {
	return func(opts *Options) {
		opts.PreAlloc = preAlloc
	}
}

// WithMaxBlockingsTasks sets up the maximum number of goroutines that are blocked when is reaches the capacity of pool
func WithMaxBlockingTasks(maxBlockingTasks int) Option {
	return func(opts *Options) {
		opts.MaxBlockingTasks = maxBlockingTasks
	}
}

// WithNoblocking indicates that pool will return nil when there is no available workers.
func WithNoblicking(nonblocking bool) Option {
	return func(opts *Options) {
		opts.Nonblocking = nonblocking
	}
}

// WithPanicHandler sets up panic handler
func WithPanicHandler(panicHandler func(interface{})) Option {
	return func(opts *Options) {
		opts.PanicHandler = panicHandler
	}
}

func WithLogger(logger ants.Logger) Option {
	return func(opts *Options) {
		opts.Logger = logger
	}
}
