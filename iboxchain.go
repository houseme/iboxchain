package iboxchain

import (
	"context"

	"github.com/houseme/iboxchain/utility/cache"
	"github.com/houseme/iboxchain/utility/logger"
	"github.com/houseme/iboxchain/utility/request"
)

// IBoxChain 字节系开放平台
type IBoxChain struct {
	cache   cache.Cache
	request request.Request
	logger  logger.ILogger
	version string
}

// NewIBoxChain 创建一个新的 IBoxChain
func NewIBoxChain(ctx context.Context) *IBoxChain {
	return &IBoxChain{
		cache:   cache.NewRedis(ctx, cache.NewDefaultRedisOpts()),
		request: request.NewDefaultRequest(""),
		logger:  logger.NewDefaultLogger(),
		version: version,
	}
}

// SetCache 设置缓存
func (i *IBoxChain) SetCache(c cache.Cache) {
	i.cache = c
}

// SetRequest 设置请求
func (i *IBoxChain) SetRequest(r request.Request) {
	i.request = r
}

// SetLogger 设置日志
func (i *IBoxChain) SetLogger(l logger.ILogger) {
	i.logger = l
}

// Version 获取版本
func (i *IBoxChain) Version() string {
	return i.version
}
