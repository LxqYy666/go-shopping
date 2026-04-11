# go-shopping
这是一个go语言后端练习项目
# TodoList
- [ ] 使用Gin、GORM、mysql实现基本的商城后端功能
- [x] 加入Redis组件（购物车/订单缓存）
- [ ] 加入消息队列组件
- [x] 订单超时自动取消（默认30分钟，自动回补库存）

# Redis配置
在 `config/config.go` 中可配置：
- `RedisAddr` / `RedisPassword` / `RedisDB`
- `CacheTTLSeconds`：缓存过期秒数
- `OrderAutoCancelMinutes`：订单超时取消分钟数
- `OrderCancelScanSeconds`：超时扫描间隔秒数
