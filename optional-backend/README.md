# 可选后端接口（不装也能用）

主题在前台会请求两个接口。**这两个接口不是必需依赖**：请求失败时主题自动降级，页面照常显示。

| 接口 | 用途 | 缺失时的降级表现 |
|---|---|---|
| `GET /api/v1/public/daily-quote` | Hero 副标题的每日随机语录 | 固定显示「不疯狂不星期四 · 精神状态遥遥领先」 |
| `POST /api/v1/public/presence` | 实时在线人数（访问者心跳） | 固定显示 `38` |

---

## 一、安装位置

两个文件都是 `publicconfighttp` 包内的 `*Handler` 方法，必须放到与 `handler.go` 同级目录：

```
internal/modules/settings/transport/http/public/
├── handler.go            # 上游已有
├── routes.go             # 上游已有（需加两行路由）
├── quote_handler.go      # 本目录文件
└── presence_handler.go   # 本目录文件
```

复制：

```bash
cp optional-backend/quote_handler.go    internal/modules/settings/transport/http/public/
cp optional-backend/presence_handler.go internal/modules/settings/transport/http/public/
```

## 二、注册路由

编辑 `internal/modules/settings/transport/http/public/routes.go`，在 `RegisterPublicRoutes` 内加两条：

```go
public.GET("/config", handler.GetConfig)
public.GET("/daily-quote", handler.GetDailyQuote)   // 新增
public.POST("/presence", handler.PostPresence)      // 新增
```

最终前缀由上游挂载点决定，通常是 `/api/v1/public`。

## 三、语录文件（可选）

`quote_handler.go` 读取一个 JSON 数组文件，10 分钟缓存：

```json
[
  { "text": "V我50，我去买两个汉堡，一个给你一个给我" },
  { "text": "不疯狂不星期四 · 精神状态遥遥领先" }
]
```

- 默认路径：`data/kfc_quotes.json`（相对进程工作目录）
- 自定义：设置环境变量 `KFC_QUOTE_FILE=/绝对路径/quotes.json`
- 文件不存在或格式错误 → 自动用内置兜底文案，不影响服务

## 四、编译与验证

```bash
go build -tags fullstack -o dujiao-next ./cmd/server

# 语录
curl -s http://127.0.0.1:8080/api/v1/public/daily-quote
# {"quote":"..."}

# 在线人数（先发一次心跳，再发一次应显示 1）
curl -s -X POST http://127.0.0.1:8080/api/v1/public/presence \
  -H 'Content-Type: application/json' -d '{"visitor_id":"test-1"}'
# {"count":1}
```

在线人数存在**进程内存**里：重启清零，多实例部署时各实例独立计数。个人站单实例足够，不需要 Redis。

---

## 五、相关：把 `kfc` 加进后端模板白名单（建议做）

上游 `internal/modules/settings/application/site_normalize.go` 的 `normalizeStorefrontTemplate()`
只放行 `classic` / `vault`，其余一律回落 `classic`。也就是说：**只装前台主题时，数据库里写 `kfc`
不会生效**，只有 `?template=kfc`（写 localStorage）能看到效果。

要正式启用，需应用本仓库的补丁：

```bash
git apply /path/to/dujiao-next-kfc-theme/patches/kfc-storefront-whitelist-v1.5.2.patch
```

补丁只改两个文件（`internal/constants/constants.go` 加常量、`site_normalize.go` 放行 `kfc`），不动业务逻辑。

## 六、相关：页脚 QQ / QQ 群字段（可选）

`Footer.vue` 会读 `config.contact.qq`、`config.contact.qq_group`、`config.contact.qq_mode`。
上游 `normalizeSiteContact()` 只保留 `telegram` / `whatsapp`，**未识别的字段会被丢弃**，
所以页脚 QQ 区块在不打后端补丁时永远不渲染（不影响其他内容）。

要让 QQ 联系方式生效，需要自行在后端补：

1. `normalizeSiteContact()` 增加 `qq` / `qq_group` / `qq_mode`（`qq_mode` 取 `qq` 或 `qq_group`）
2. `validateSiteSetting()` 增加对应校验（QQ 号纯数字、群链接为 URL）
3. 后台「站点设置 → 联系方式」增加三个输入框，保存到 `site_config.contact`

本仓库**不包含**这部分后台 UI 改动（`frontend/admin` 属另一套构建），仅作说明。
