package rueidis

import (
	"bytes"
	"context"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/util"
	"github.com/redis/rueidis"
	"github.com/redis/rueidis/rueidishook"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"math"
	"net"
	"strconv"
	"strings"
	"time"
)

type datadogHook struct {
	serviceName          string
	additionalTagOptions []ddtrace.StartSpanOption
	skipRaw              bool
	analyticsRate        float64
}

func newDatadogHook(serviceName string, dsn string, db int, skipRaw bool) rueidishook.Hook {
	var analyticsRate float64
	if util.BoolEnv("DD_TRACE_REDIS_ANALYTICS_ENABLED", false) {
		analyticsRate = 1.0
	} else {
		analyticsRate = math.NaN()
	}

	return &datadogHook{
		serviceName:          serviceName,
		additionalTagOptions: newAdditionalTagOptions(dsn, db),
		skipRaw:              skipRaw,
		analyticsRate:        analyticsRate,
	}
}

//nolint:revive
func (h *datadogHook) Do(
	client rueidis.Client, ctx context.Context, cmd rueidis.Completed,
) rueidis.RedisResult {
	ctx = h.beforeProcess(ctx, cmd.Commands())
	resp := client.Do(ctx, cmd)
	h.afterProcess(ctx, resp)

	return resp
}

//nolint:revive
func (h *datadogHook) DoMulti(
	client rueidis.Client, ctx context.Context, multi ...rueidis.Completed,
) []rueidis.RedisResult {
	ctx = h.beforeProcess(ctx, h.completedToArgs(multi...)...)
	resps := client.DoMulti(ctx, multi...)
	h.afterProcess(ctx, resps...)

	return resps
}

//nolint:revive
func (h *datadogHook) DoCache(
	client rueidis.Client, ctx context.Context, cmd rueidis.Cacheable, ttl time.Duration,
) rueidis.RedisResult {
	ctx = h.beforeProcess(ctx, cmd.Commands())
	resp := client.DoCache(ctx, cmd, ttl)
	h.afterProcess(ctx, resp)

	return resp
}

//nolint:revive
func (h *datadogHook) DoMultiCache(
	client rueidis.Client, ctx context.Context, multi ...rueidis.CacheableTTL,
) []rueidis.RedisResult {
	ctx = h.beforeProcess(ctx, h.cacheableTTLToArgs(multi...)...)
	resps := client.DoMultiCache(ctx, multi...)
	h.afterProcess(ctx, resps...)

	return resps
}

// nolint: revive
func (h *datadogHook) DoStream(
	client rueidis.Client, ctx context.Context, cmd rueidis.Completed,
) rueidis.RedisResultStream {
	ctx = h.beforeProcess(ctx, cmd.Commands())
	resps := client.DoStream(ctx, cmd)
	h.afterProcessStream(ctx, resps)

	return resps
}

// nolint: revive
func (h *datadogHook) DoMultiStream(
	client rueidis.Client, ctx context.Context, multi ...rueidis.Completed,
) rueidis.MultiRedisResultStream {
	ctx = h.beforeProcess(ctx, h.completedToArgs(multi...)...)
	resps := client.DoMultiStream(ctx, multi...)
	h.afterProcessStream(ctx, resps)

	return resps
}

//nolint:revive
func (h *datadogHook) Receive(
	client rueidis.Client, ctx context.Context, subscribe rueidis.Completed, fn func(msg rueidis.PubSubMessage),
) error {
	// do whatever you want before client.Receive
	err := client.Receive(ctx, subscribe, fn)
	// do whatever you want after client.Receive
	return err
}

func (h *datadogHook) commandsToString(cmds ...[]string) string {
	var b bytes.Buffer
	for _, cmd := range cmds {
		b.WriteString(strings.Join(cmd, ""))
		b.WriteString("\n")
	}

	return b.String()
}

func (h *datadogHook) completedToArgs(commands ...rueidis.Completed) [][]string {
	var args [][]string
	for _, c := range commands {
		args = append(args, c.Commands())
	}

	return args
}

func (h *datadogHook) cacheableTTLToArgs(commands ...rueidis.CacheableTTL) [][]string {
	var args [][]string
	for _, c := range commands {
		args = append(args, c.Cmd.Commands())
	}

	return args
}

func (h *datadogHook) beforeProcess(ctx context.Context, args ...[]string) context.Context {
	resourceName := ""
	nArg := len(args)
	pipeline := false
	if len(args) > 0 {
		resourceName = args[0][0]
		pipeline = true
	}

	var opts []ddtrace.StartSpanOption
	opts = append(opts,
		tracer.SpanType(ext.SpanTypeRedis),
		tracer.ServiceName(h.serviceName),
		tracer.ResourceName(resourceName),
		tracer.Tag("redis.args_length", strconv.Itoa(nArg)),
	)

	if pipeline {
		opts = append(opts,
			tracer.Tag("redis.pipeline_length", strconv.Itoa(len(args))),
		)
	}
	if h.skipRaw {
		opts = append(opts, tracer.Tag("redis.raw_command", h.commandsToString(args...)))
	}
	opts = append(opts, h.additionalTagOptions...)

	analyticsRate := h.analyticsRate
	if !math.IsNaN(analyticsRate) {
		opts = append(opts, tracer.Tag(ext.EventSampleRate, analyticsRate))
	}

	_, ctx = tracer.StartSpanFromContext(ctx, "redis.command", opts...)

	return ctx
}

func (h *datadogHook) afterProcess(ctx context.Context, results ...rueidis.RedisResult) {
	span, _ := tracer.SpanFromContext(ctx)
	var finishOpts []ddtrace.FinishOption
	for _, result := range results {
		if err := result.Error(); err != nil && !rueidis.IsRedisNil(err) {
			finishOpts = append(finishOpts, tracer.WithError(err))
		}
	}

	span.Finish(finishOpts...)
}

func (h *datadogHook) afterProcessStream(ctx context.Context, resp rueidis.RedisResultStream) {
	span, _ := tracer.SpanFromContext(ctx)
	var finishOpts []ddtrace.FinishOption

	for resp.HasNext() {
		if err := resp.Error(); err != nil {
			finishOpts = append(finishOpts, tracer.WithError(err))
		}
	}

	span.Finish(finishOpts...)
}

// TODO Support tags for cluster options.
func newAdditionalTagOptions(dsn string, db int) []ddtrace.StartSpanOption {
	var additionalTags []ddtrace.StartSpanOption

	host, port, err := net.SplitHostPort(dsn)
	if err != nil {
		host = dsn
		port = "6379"
	}
	additionalTags = []ddtrace.StartSpanOption{
		tracer.Tag(ext.TargetHost, host),
		tracer.Tag(ext.TargetPort, port),
		tracer.Tag("out.db", strconv.Itoa(db)),
	}

	return additionalTags
}
