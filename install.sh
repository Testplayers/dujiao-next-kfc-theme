#!/bin/bash
# ============================================================
# 疯狂星期四主题包 —— 覆盖式安装脚本
# 用法: ./install.sh /path/to/dujiao-next
# 作用: 把 files/ 下的 5 个前台文件复制到源码树对应位置（会先备份 .bak-kfc）
# ============================================================
set -euo pipefail

TARGET="${1:-}"
if [ -z "$TARGET" ] || [ ! -d "$TARGET/frontend/user/src" ]; then
  echo "用法: $0 /path/to/dujiao-next"
  echo "（目标目录下必须存在 frontend/user/src）"
  exit 1
fi

HERE="$(cd "$(dirname "$0")" && pwd)"
SRC="$HERE/files/frontend/user/src"
DST="$TARGET/frontend/user/src"

echo "主题包目录: $HERE"
echo "目标源码树: $TARGET"
echo

for f in App.vue style.css views/Home.vue views/ProductDetail.vue templates/registry.ts; do
  if [ -f "$SRC/$f" ]; then
    if [ -f "$DST/$f" ]; then
      cp -f "$DST/$f" "$DST/$f.bak-kfc"
    fi
    mkdir -p "$(dirname "$DST/$f")"
    cp -f "$SRC/$f" "$DST/$f"
    echo "  覆盖 $f（原文件已备份为 .bak-kfc）"
  else
    echo "  跳过 $f（主题包内缺失）"
  fi
done

echo
echo "完成。下一步："
echo "  cd $TARGET/frontend/user && pnpm install && pnpm build"
echo "  rm -rf $TARGET/internal/web/dist/user && cp -r $TARGET/frontend/user/dist $TARGET/internal/web/dist/user"
echo "  cd $TARGET && go build -tags fullstack -o dujiao-next ./cmd/server"
echo
echo "启用主题：先 git apply patches/kfc-storefront-whitelist-v1.5.2.patch（否则后端会把 kfc 回落成 classic），"
echo "          再把 settings.site_config 的 storefront_template 改成 kfc；也可用 ?template=kfc 预览。"
echo "回滚：把 .bak-kfc 备份覆盖回去，或 git checkout -- frontend/user/src"
