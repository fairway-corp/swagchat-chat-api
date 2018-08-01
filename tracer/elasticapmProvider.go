package tracer

import (
	"context"
	"fmt"

	elasticapm "github.com/elastic/apm-agent-go"
	"github.com/swagchat/chat-api/logger"
	"github.com/swagchat/chat-api/utils"
)

var elasticapmTracer *elasticapm.Tracer

type elasticapmProvider struct {
	ctx context.Context
}

func (ep *elasticapmProvider) NewTracer() error {
	tracer, err := elasticapm.NewTracer(utils.AppName, utils.BuildVersion)
	if err != nil {
		logger.Error(err.Error())
		return nil
	}
	tracer.SetLogger(logger.Logger())
	elasticapmTracer = tracer
	return nil
}

func (ep *elasticapmProvider) StartTransaction(name, transactionType string) context.Context {
	if elasticapmTracer == nil {
		return ep.ctx
	}

	transaction := elasticapmTracer.StartTransaction(name, transactionType)
	ctx := elasticapm.ContextWithTransaction(ep.ctx, transaction)
	ctx = context.WithValue(ctx, utils.CtxTracerTransaction, transaction)
	return ctx
}

func (ep *elasticapmProvider) StartSpan(name, spanType string) interface{} {
	span, _ := elasticapm.StartSpan(ep.ctx, name, spanType)
	return span
}

func (ep *elasticapmProvider) SetTag(key string, value interface{}) {
	transaction := ep.ctx.Value(utils.CtxTracerTransaction)
	if transaction != nil {
		txCtx := transaction.(*elasticapm.Transaction).Context
		txCtx.SetTag(key, fmt.Sprintf("%v", value))
	}
}

func (ep *elasticapmProvider) Finish(span interface{}) {
	if span != nil {
		span.(*elasticapm.Span).End()
	}
}

func (ep *elasticapmProvider) Close() {
	transaction := ep.ctx.Value(utils.CtxTracerTransaction)
	if transaction != nil {
		transaction.(*elasticapm.Transaction).End()
	}

	elasticapmTracer.Flush(nil)
}
