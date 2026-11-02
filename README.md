# 疯狂星期四主题包（dujiao-next KFC Theme）

一套给 [dujiao-next](https://github.com/dujiao-next/dujiao-next) 官方前端使用的「疯狂星期四」风格皮肤：
黑黄红撞色、黄色跑马灯、`VME50` 徽章、每日语录副标题、实时在线人数、栏目标题贴纸。

> 基于 **dujiao-next v1.5.2** 的官方前端二次修改，仅改前台（`frontend/user`），不动后台与后端业务逻辑。

---

## ⚠️ 免责声明 / 来源说明（请先读完再使用）

1. **来源**：本仓库所有代码均衍生自 [dujiao-next](https://github.com/dujiao-next/dujiao-next) 官方开源项目的前端源码，
   原始版权归原作者所有。本仓库仅在其基础上做视觉风格修改，**不是官方项目**，与官方无任何隶属或背书关系。
2. **许可**：官方项目使用 **AGPL-3.0**。作为衍生作品，本仓库同样以 **AGPL-3.0** 发布（见 `LICENSE`），
   保留原版权声明，并随仓库提供完整可构建的源码。你若分发修改版，需同样遵守 AGPL-3.0（开源、保留声明、提供源码）。
3. **用途**：仅供学习交流与个人自建站点参考，**禁止用于任何违法违规用途**。
   使用本主题造成的任何直接或间接后果（含店铺经营、合规、数据损失），由使用者自行承担。
4. **商标与素材**：`疯狂星期四`、`KFC`、`VME50`、`V我50` 等为网络流行语或第三方商标/品牌。
   本仓库**不包含任何官方图片、Logo、字体或宣传素材**，仅以纯 CSS/文字做风格致敬，与相关品牌方无任何关联或授权关系。
5. **侵权处理**：若本仓库内容（代码、文案、样式）侵犯了你的合法权益，请在本仓库提 Issue 说明具体文件与权利依据，
   核实后会在 **48 小时内删除相关内容或整仓下线**。

---

## 一、效果概览

| 位置 | 效果 |
|---|---|
| 顶部 | 黑色导航条（文档流，不悬浮），白字 / hover 变黄 / 激活项红底白字 |
| 导航下方 | 黄色跑马灯，循环滚动疯四梗文案 |
| Hero 区 | 超大标题「疯狂星期四」+ 倾斜 `VME50` 徽章（3° 顺时针、黑描边硬阴影） |
| 副标题 | 每日随机语录（超长自动降字号，完整换行不截断） |
| CTA | 白底黑字按钮，按星期动态显示「距下个疯狂星期四还有 X 天」 |
| CTA 右侧 | 真实在线人数「N 位文案鬼才在线」（每 30 秒心跳刷新） |
| 栏目标题 | 黄底黑字黑边、倾斜 2° 的贴纸标签 |
| 商品页 | 商品简介里的 URL 自动转成可点击链接（防注入转义） |
| 页脚 | 支持 QQ / QQ 群联系方式展示 |

> 仓库本身不含截图。建议你自己部署后截图补充，或参考样式预览。

---

## 二、原理（重要，别期待「一个文件夹丢进去就能用」）

dujiao-next 的模板机制是「classic / vault 两套内置模板 + `storefront_template` 开关」，
**不支持第三方模板目录自动发现**。所以这套 KFC 主题的实现方式是：

```
classic 模板（官方原有结构）
   + style.css 追加 338 行 KFC 专属样式（作用域都在 .template-kfc / .kfc-* 下）
   + App.vue 增加跑马灯 / Hero / 徽章 / CTA / 在线数的结构块
   + Home.vue / ProductDetail.vue / Footer.vue 少量适配
   + registry.ts 增加 'kfc' 这个模板标识
```

激活条件：`storefront_template === 'kfc'`（或 URL 带 `?template=kfc` 做本地预览）。
未激活时，这些结构块不渲染、样式不生效，站点仍是官方 classic 外观——**可以安全地「装上但不启用」**。

---

## 三、目录结构

```
.
├── README.md                  # 本文档
├── LICENSE                    # AGPL-3.0（继承自上游）
├── NOTICE.md                  # 来源与改动清单
├── install.sh                 # 把 files/ 覆盖到指定源码树（可选）
├── files/                     # 完整文件版（覆盖式安装）
│   └── frontend/user/src/
│       ├── App.vue
│       ├── style.css
│       ├── components/Footer.vue
│       ├── views/Home.vue
│       ├── views/ProductDetail.vue
│       └── templates/registry.ts
├── patches/
│   ├── kfc-theme-v1.5.2.patch              # 6 个前台文件，相对官方 v1.5.2 的 diff（推荐）
│   └── kfc-storefront-whitelist-v1.5.2.patch # 可选：让后端接受 storefront_template=kfc（2 文件）
└── optional-backend/          # 可选：语录接口 + 在线人数接口
    ├── README.md
    ├── quote_handler.go
    └── presence_handler.go
```

---

## 四、环境要求

- dujiao-next **v1.5.2** 源码（其他 v1.5.x 大概率可用，patch 失败时改用覆盖法并手工比对）
- Node.js **20.x**、pnpm（官方 `packageManager: pnpm@10.34.3`）
- Go **1.22+**（构建 `-tags fullstack`）

---

## 五、安装

### 方式一：打补丁（推荐）

```bash
git clone https://github.com/dujiao-next/dujiao-next.git
cd dujiao-next
git checkout v1.5.2

# 应用主题补丁
git apply /path/to/dujiao-next-kfc-theme/patches/kfc-theme-v1.5.2.patch

# 若有冲突（上游已改同名文件），改用方式二
```

### 方式二：覆盖文件

```bash
chmod +x /path/to/dujiao-next-kfc-theme/install.sh
# 在 dujiao-next 源码根目录执行
/path/to/dujiao-next-kfc-theme/install.sh /path/to/dujiao-next
```

或手工把 `files/frontend/user/src/` 下 6 个文件复制到源码树对应位置。

---

## 六、构建与嵌入

```bash
# 1) 前台构建
cd frontend/user
pnpm install
pnpm build                 # = vue-tsc -b && vite build（想跳过类型检查可只跑 npx vite build）
cd ../..

# 2) 把产物嵌入 Go 二进制（embed 目录）
rm -rf internal/web/dist/user
mkdir -p internal/web/dist
cp -r frontend/user/dist internal/web/dist/user

# 3) 编译（必须带 fullstack tag，否则前端走外部文件、不嵌入）
go build -tags fullstack -o dujiao-next ./cmd/server
```

> ⚠️ 官方后台 `frontend/admin` 不需要动，除非你还想把 KFC 加进后台模板下拉（见下节）。

---

## 七、启用主题

三种方式，任选其一：

**A. 改数据库（正式启用）**

```sql
-- settings 表，key = site_config 的 value_json 里 storefront_template 改成 "kfc"
-- 建议用后台「站点设置」保存一次，或以 SQL 精确替换后重启服务
UPDATE settings
SET value_json = REPLACE(value_json, '"storefront_template":"classic"', '"storefront_template":"kfc"')
WHERE key = 'site_config';
```

改完重启 dujiao-next，或等缓存过期。前端读取的是 `/api/v1/public/config` 返回的 `storefront_template`。

> ⚠️ **v1.5.2 必做**：`site_normalize.go` 的 `normalizeStorefrontTemplate()` 只允许 `classic` / `vault`，
> 传 `kfc` 会被回落成 `classic`，即「只装前台主题 → 数据库里写 kfc 不生效」。
> 执行 `git apply patches/kfc-storefront-whitelist-v1.5.2.patch` 放行 `kfc` 后再重新编译即可。

**B. URL 预览（不落库，适合先看效果）**

```
https://你的域名/?template=kfc      # 写入 localStorage 覆盖
https://你的域名/?template=reset    # 清除覆盖
```

**C. 加进后台下拉（可选）**

需要三处小改 + 重建 admin：
- `frontend/admin/src/views/admin/Settings.vue`：模板选项里加一项 `kfc`
- `frontend/admin/src/i18n/index.ts`：加 `kfcMode` / `kfcModeDesc` 文案（三语）
- 重建：`cd frontend/admin && pnpm build:fullstack`，产物复制到 `internal/web/dist/admin`

---

## 八、可选：语录接口与在线人数接口

主题会请求两个接口，**没有它们也能跑**（自动降级）：

| 接口 | 缺失时的表现 |
|---|---|
| `GET /api/v1/public/daily-quote` | 副标题固定显示默认文案「不疯狂不星期四」 |
| `POST /api/v1/public/presence` | 在线人数固定显示 `38` |

想启用真实数据，把 `optional-backend/` 里两个 handler 放进源码并注册两条路由即可，详见该目录的 `README.md`。

---

## 九、卸载 / 回滚

```bash
git checkout -- frontend/user/src
# 或
git apply -R /path/to/dujiao-next-kfc-theme/patches/kfc-theme-v1.5.2.patch
```

再把 `storefront_template` 改回 `classic`，重新构建部署即可。

---

## 十、常见问题（踩过的坑）

1. **改完样式没生效**
   Vite 构建有哈希文件名，`dist` 必须重新 `cp` 到 `internal/web/dist/user` 并重新 `go build -tags fullstack`，
   否则跑的还是旧产物。
2. **`<script type="module" src="...index-xxx.js?v=时间戳">` 加了 `?v=` 会导致整站白屏**
   懒加载 chunk 内部仍以无参数形式 `import('./index-xxx.js')`，浏览器视为两个模块 → 入口执行两次 →
   两个 Vue 实例抢 `#app` → 报 `Cannot read properties of null (reading 'nextSibling')`。
   **入口 script 的 URL 绝不要加查询参数**；缓存失效靠 `Cache-Control: no-cache`（官方 index.html 已设）。
3. **Tailwind 任意值类（如 `bg-[#c41200]`）不生效**
   新增任意值类后必须重新 `pnpm build`，热更新不一定编译出来。
4. **样式被 Tailwind 的 `!important` 覆盖**
   KFC 的覆盖规则写在 `@layer utilities` 里，才能赢过官方 utilities。
5. **首屏闪一下 classic 外壳**
   主题用 `hasTemplateHint()` 判断是否已知模板；首次访问者会等配置返回再渲染，属正常设计。
6. **`?template=kfc` 只在当前浏览器生效**
   因为它写的是 `localStorage`，换设备/无痕窗口无效，正式启用请走数据库配置。
7. **移动端跑马灯/徽章错位**
   优先用内联 `style` 固定关键尺寸（字体、行高、旋转），Tailwind 的响应式前缀在部分断点下会被覆盖。

---

## 十一、许可

[AGPL-3.0](LICENSE)，继承自上游 dujiao-next。衍生于 <https://github.com/dujiao-next/dujiao-next>。
