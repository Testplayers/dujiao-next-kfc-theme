<template>
  <div id="app" class="min-h-screen bg-background text-foreground flex flex-col" :class="{ 'template-kfc': isKfc }">
    <!-- 首次访问还不知道站点用哪套模板时先不渲染外壳，等配置返回（见 shellReady） -->
    <template v-if="shellReady">
      <!-- vault 模板：自带顶栏/页脚的外壳包裹页面（控制台仍走下方分支） -->
      <VaultLayout v-if="isVault && !isResellerConsole">
        <ErrorBoundary>
          <RouterView v-slot="{ Component }">
            <Transition name="page-fade" mode="out-in">
              <component :is="Component" />
            </Transition>
          </RouterView>
        </ErrorBoundary>
      </VaultLayout>

      <!-- classic 模板 / 分销控制台（保持原有结构不变） -->
      <template v-else>
        <Navbar v-if="!isResellerConsole" />
        <div v-if="isKfc && !isResellerConsole" class="kfc-marquee">
          <div class="kfc-marquee-track">
            <span class="px-4">精神状态遥遥领先</span><span class="px-2">✦</span><span class="px-4">一周只疯一天</span><span class="px-2">✦</span><span class="px-4">今天你 V50 了吗</span><span class="px-2">✦</span><span class="px-4">不疯狂不星期四</span><span class="px-2">✦</span><span class="px-4">疯四文字，代代相传</span><span class="px-2">✦</span><span class="px-4">V ME 50</span><span class="px-2">✦</span><span class="px-4">精神状态遥遥领先</span><span class="px-2">✦</span><span class="px-4">一周只疯一天</span><span class="px-2">✦</span><span class="px-4">今天你 V50 了吗</span><span class="px-2">✦</span><span class="px-4">不疯狂不星期四</span><span class="px-2">✦</span><span class="px-4">疯四文字，代代相传</span><span class="px-2">✦</span><span class="px-4">V ME 50</span><span class="px-2">✦</span>
          </div>
        </div>
        <div v-if="isKfc && !isResellerConsole" class="kfc-hero">
          <div class="kfc-hero-title">
            <span>疯狂星期四</span>
            <span class="kfc-hero-badge">VME50</span>
          </div>
          <p class="kfc-hero-sub" :style="{ fontSize: heroQuoteSize }">{{ heroQuote }}</p>
          <div class="kfc-hero-row">
            <a href="/products" class="kfc-hero-cta">{{ kfcCta }}</a>
            <span class="kfc-hero-online"><span class="kfc-hero-online-dot"></span>{{ kfcOnline }} 位文案鬼才在线</span>
          </div>
        </div>
        <main class="flex-1" :class="isResellerConsole ? '' : 'pb-14 lg:pb-0'">
          <ErrorBoundary>
            <RouterView v-slot="{ Component }">
              <Transition name="page-fade" mode="out-in">
                <component :is="Component" />
              </Transition>
            </RouterView>
          </ErrorBoundary>
        </main>
        <Footer v-if="!isResellerConsole" />
        <BackToTop v-if="!isResellerConsole" />
        <MobileBottomNav v-if="!isResellerConsole" />
      </template>
    </template>

    <Loading :loading="appStore.loading" />
    <Toast />
    <ConfirmDialog />
  </div>
</template>

<script setup lang="ts">
import { computed, defineAsyncComponent, ref, watch, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAppStore } from './stores/app'
import { getActiveTemplate, hasTemplateHint, rememberTemplate } from './templates/registry'
import Navbar from './components/Navbar.vue'
import Footer from './components/Footer.vue'
import Loading from './components/Loading.vue'
import Toast from './components/Toast.vue'
import ConfirmDialog from './components/ConfirmDialog.vue'
import ErrorBoundary from './components/ErrorBoundary.vue'
import BackToTop from './components/BackToTop.vue'
import MobileBottomNav from './components/MobileBottomNav.vue'

// vault 外壳按需加载，classic 用户不会拉取其 chunk/样式
const VaultLayout = defineAsyncComponent(() => import('./templates/vault/layout/VaultLayout.vue'))

// config 由 router.beforeEach 统一加载，无需在此重复调用
const appStore = useAppStore()
const route = useRoute()
const isResellerConsole = computed(() => route.meta.resellerConsole === true)
// getActiveTemplate 读取 appStore.config（响应式），config 加载后会重新计算
const isVault = computed(() => getActiveTemplate() === 'vault')
const isKfc = computed(() => getActiveTemplate() === 'kfc')

// 模板来自异步的站点配置。已有线索（预览覆盖 / 上次访问记住的模板）时立即渲染外壳；
// 首次访问则等配置返回再渲染，避免先闪一下 classic 外壳。配置加载失败时由首个路由就绪兜底
// （守卫里的加载无论成败都会结束），不会一直停在空白。
const shellReady = ref(hasTemplateHint())
if (!shellReady.value) {
  watch(() => appStore.config, (config) => { if (config) shellReady.value = true }, { immediate: true })
  void useRouter().isReady().catch(() => {}).finally(() => { shellReady.value = true })
}
watch(() => appStore.config?.storefront_template, (value) => rememberTemplate(value), { immediate: true })

const heroQuote = ref('不疯狂不星期四')
const kfcCta = ref('今天就是疯狂星期四！V我50')
const kfcOnline = ref(38)

const heroQuoteSize = computed(() => {
  const len = heroQuote.value.length
  if (len <= 30) return '16px'
  if (len <= 60) return '15px'
  if (len <= 100) return '14px'
  return '13px'
})

const loadKfcBanner = async () => {
  try {
    const res = await fetch('/api/v1/public/daily-quote')
    const json = await res.json()
    const text = (json && (json.quote || (json.data && json.data.quote))) || ''
    if (text) heroQuote.value = text
  } catch {
    /* 保底文案 */
  }
  const day = new Date().getDay()
  if (day === 4) kfcCta.value = '今天就是疯狂星期四！V我50'
  else if (day === 5) kfcCta.value = '明天就是疯狂星期四！'
  else kfcCta.value = '距下个疯狂星期四还有 ' + ((11 - day) % 7) + ' 天'
}

const loadPresence = async () => {
  try {
    let vid = localStorage.getItem('dj-visitor-id')
    if (!vid) {
      vid = crypto.randomUUID ? crypto.randomUUID() : String(Date.now()) + Math.random()
      localStorage.setItem('dj-visitor-id', vid)
    }
    const res = await fetch('/api/v1/public/presence', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ visitor_id: vid })
    })
    const j = await res.json()
    if (j && typeof j.count === 'number' && j.count > 0) kfcOnline.value = j.count
  } catch {
    /* 网络异常保留当前值 */
  }
}

onMounted(() => {
  if (isKfc.value) { loadKfcBanner(); loadPresence() }
  setInterval(() => { if (isKfc.value) loadPresence() }, 30000)
})

watch(isKfc, (v) => {
  if (v) { loadKfcBanner(); loadPresence() }
})
</script>

<style>
.page-fade-enter-active,
.page-fade-leave-active {
  transition: opacity 200ms ease;
}

.page-fade-enter-from,
.page-fade-leave-to {
  opacity: 0;
}
</style>
