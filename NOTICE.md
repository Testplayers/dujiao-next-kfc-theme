# NOTICE — 来源与改动说明

## 上游项目

- 名称：dujiao-next（独角数卡 NEXT）
- 仓库：<https://github.com/dujiao-next/dujiao-next>
- 版本基线：**v1.5.2**（commit `c8144b2`）
- 许可：GNU GPL v3.0（仓库内 `LICENSE` 为 GPL-3.0 全文；GitHub 许可证识别结果为 `GPL-3.0`）

本仓库是上述项目 **前端（`frontend/user`）的衍生修改**，不含上游完整源码。
使用、分发、修改时请自行获取上游源码并遵守 GPL-3.0。

## 本仓库改了什么

核心主题仅涉及 5 个前台文件，均为视觉/结构层改动，不触碰业务逻辑、接口签名、数据库结构：

| 文件 | 改动量（相对 v1.5.2） | 内容 |
|---|---|---|
| `frontend/user/src/style.css` | +338 行（纯追加） | KFC 配色、导航、跑马灯、Hero、徽章、栏目标题贴纸等样式 |
| `frontend/user/src/App.vue` | +74 / −2 | 跑马灯、Hero 区、`VME50` 徽章、CTA 倒计时、在线人数、语录加载 |
| `frontend/user/src/views/Home.vue` | +5 / 0 | 首页适配 |
| `frontend/user/src/views/ProductDetail.vue` | +9 / −3 | 商品简介 URL 自动转链接（先转义防注入） |
| `frontend/user/src/templates/registry.ts` | +2 / −2 | 模板标识新增 `kfc` |

另有 1 项**可选**补丁（不打也能用，只是启用方式受限）：

| 可选内容 | 文件 | 作用 |
|---|---|---|
| 模板白名单补丁 | `patches/kfc-storefront-whitelist-v1.5.2.patch`（2 文件，各 1 段） | 让后端 `normalizeStorefrontTemplate()` 接受 `kfc`；不打则只能用 `?template=kfc` 预览 |

该补丁为 2 行级别的最小改动，同样以 GPL-3.0 发布。

> 本仓库不包含任何后端接口实现（语录 / 在线人数 / QQ 联系方式均未附带），主题对这两个接口只做
> 「有则用、无则降级」的调用；页脚保持官方样式，不引入自定义联系字段。

## 不含的内容

- 不含任何官方图片、Logo、字体、宣传素材
- 不含任何店铺私有配置、密钥、域名、联系方式
- 不含支付相关代码

## 侵权联系

如认为本仓库内容侵犯你的权益，请提 Issue 并提供：具体文件路径 + 你的权利依据。
核实后 48 小时内删除相关内容或下线整仓。
