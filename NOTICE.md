# NOTICE — 来源与改动说明

## 上游项目

- 名称：dujiao-next（独角数卡 NEXT）
- 仓库：<https://github.com/dujiao-next/dujiao-next>
- 版本基线：**v1.5.2**（commit `c8144b2`）
- 许可：GNU AGPL-3.0（见本仓库 `LICENSE`，版权归上游作者所有）

本仓库是上述项目 **前端（`frontend/user`）的衍生修改**，不含上游完整源码。
使用、分发、修改时请自行获取上游源码并遵守 AGPL-3.0。

## 本仓库改了什么

核心主题仅涉及 6 个前台文件，均为视觉/结构层改动，不触碰业务逻辑、接口签名、数据库结构：

| 文件 | 改动量（相对 v1.5.2） | 内容 |
|---|---|---|
| `frontend/user/src/style.css` | +338 行（纯追加） | KFC 配色、导航、跑马灯、Hero、徽章、栏目标题贴纸、页脚等样式 |
| `frontend/user/src/App.vue` | +74 / −2 | 跑马灯、Hero 区、`VME50` 徽章、CTA 倒计时、在线人数、语录加载 |
| `frontend/user/src/components/Footer.vue` | +34 / −16 | QQ / QQ 群联系方式展示 |
| `frontend/user/src/views/Home.vue` | +5 / 0 | 首页适配 |
| `frontend/user/src/views/ProductDetail.vue` | +9 / −3 | 商品简介 URL 自动转链接（先转义防注入） |
| `frontend/user/src/templates/registry.ts` | +2 / −2 | 模板标识新增 `kfc` |

另有 3 项**可选**内容，不安装时主题照常工作（自动降级）：

| 可选内容 | 文件 | 作用 |
|---|---|---|
| 模板白名单补丁 | `patches/kfc-storefront-whitelist-v1.5.2.patch`（2 文件，各 1 段） | 让后端 `normalizeStorefrontTemplate()` 接受 `kfc`；不打则只能用 `?template=kfc` 预览 |
| 语录接口 | `optional-backend/quote_handler.go` | `GET /api/v1/public/daily-quote`；缺失时副标题显示内置文案 |
| 在线人数接口 | `optional-backend/presence_handler.go` | `POST /api/v1/public/presence`；缺失时显示固定数字 |

以上均为本仓库作者原创（不含上游代码的复制粘贴式改写），同样以 AGPL-3.0 发布。

## 不含的内容

- 不含任何官方图片、Logo、字体、宣传素材
- 不含任何店铺私有配置、密钥、域名、联系方式
- 不含支付相关代码

## 侵权联系

如认为本仓库内容侵犯你的权益，请提 Issue 并提供：具体文件路径 + 你的权利依据。
核实后 48 小时内删除相关内容或下线整仓。
