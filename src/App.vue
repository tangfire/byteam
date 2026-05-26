<script setup lang="ts">
import { computed, onBeforeUnmount, onMounted, ref, watch } from 'vue'
import { useRoute } from 'vue-router'
import { Close, Fold } from '@element-plus/icons-vue'


const route = useRoute()
const activeIndex = computed(() => route.path as string)
const isAdminRoute = computed(() => route.path.startsWith('/admin'))
const isMobile = ref(false)
const isMenuCollapsed = ref(true)
const mobileMenuRef = ref<HTMLElement | null>(null)
const hamburgerButtonRef = ref<HTMLButtonElement | null>(null)
const mobileBreakpoint = 860
let bodyOverflowBeforeLock: string | null = null

const mobileDefaultOpeneds = computed(() => {
  const path = route.path
  if (path === '/research-direction' || path === '/research-projects') return ['/research']
  if (path === '/dr-Baoyao-Yang' || path === '/our-group' || path === '/undergraduate' || path === '/Alumni') return ['/our-team']
  if (path === '/international-journals-conferences' || path === '/patents') return ['/publications']
  if (path === '/vKnow' || path === '/VideoMind') return ['/Project']
  return []
})

const checkScreenSize = () => {
  const viewportWidth = Math.max(document.documentElement.clientWidth || 0, window.innerWidth || 0)
  isMobile.value = viewportWidth <= mobileBreakpoint

  if (!isMobile.value) {
    isMenuCollapsed.value = true
  }
}

const closeMobileMenu = () => {
  isMenuCollapsed.value = true
}

const toggleMobileMenu = () => {
  isMenuCollapsed.value = !isMenuCollapsed.value
}

const handleSelect = () => {
  closeMobileMenu()
  if (isMobile.value) {
    window.scrollTo({
      top: 0,
      behavior: 'smooth'
    })
  }
}

const clickOutsideHandler = (e: MouseEvent) => {
  if (isMenuCollapsed.value || !isMobile.value) return
  const target = e.target as Node | null
  if (!target) return

  if (
      !mobileMenuRef.value?.contains(target) &&
      !hamburgerButtonRef.value?.contains(target)
  ) {
    closeMobileMenu()
  }
}

const keydownHandler = (event: KeyboardEvent) => {
  if (event.key === 'Escape') {
    closeMobileMenu()
  }
}

const syncBodyScrollLock = () => {
  const shouldLock = isMobile.value && !isMenuCollapsed.value && !isAdminRoute.value
  if (shouldLock && bodyOverflowBeforeLock === null) {
    bodyOverflowBeforeLock = document.body.style.overflow
    document.body.style.overflow = 'hidden'
  } else if (!shouldLock && bodyOverflowBeforeLock !== null) {
    document.body.style.overflow = bodyOverflowBeforeLock
    bodyOverflowBeforeLock = null
  }
}

onMounted(() => {
  document.addEventListener('click', clickOutsideHandler)
  document.addEventListener('keydown', keydownHandler)

  setTimeout(() => {
    checkScreenSize()
    window.dispatchEvent(new Event('resize'))
  }, 100)
  window.addEventListener('resize', checkScreenSize)
})

onBeforeUnmount(() => {
  document.removeEventListener('click', clickOutsideHandler)
  document.removeEventListener('keydown', keydownHandler)
  window.removeEventListener('resize', checkScreenSize)
  if (bodyOverflowBeforeLock !== null) {
    document.body.style.overflow = bodyOverflowBeforeLock
  }
})

watch([isMobile, isMenuCollapsed, isAdminRoute], syncBodyScrollLock)
watch(() => route.fullPath, closeMobileMenu)
</script>

<template>
  <router-view v-if="isAdminRoute"/>
  <div v-else class="common-layout">


    <el-container class="site-container">

      <!-- 移动端导航 -->
      <div v-if="isMobile" class="mobile-nav">
        <div class="mobile-nav-header">
          <router-link class="mobile-brand" to="/" @click="closeMobileMenu">
            <img class="mobile-logo-img" src="/logo/001.png" alt="BYML Logo"/>
            <span>BYML</span>
          </router-link>
          <el-button
              ref="hamburgerButtonRef"
              @click="toggleMobileMenu"
              class="hamburger-btn"
              :aria-expanded="String(!isMenuCollapsed)"
              aria-label="切换导航菜单"
          >
            <el-icon :size="24">
              <Fold v-if="isMenuCollapsed"/>
              <Close v-else/>
            </el-icon>
          </el-button>
        </div>

        <el-collapse-transition>
          <div
              v-show="!isMenuCollapsed"
              ref="mobileMenuRef"
              class="mobile-menu-content"
          >
            <el-menu
                :key="activeIndex"
                :default-active="activeIndex"
                :default-openeds="mobileDefaultOpeneds"
                active-text-color="#7d1231"
                @select="handleSelect"
                :router="true"
                class="vertical-menu"
                unique-opened
                @click.stop
            >
              <el-menu-item index="/">Home</el-menu-item>
              <el-menu-item index="/about">About</el-menu-item>
              <el-menu-item index="/news">News</el-menu-item>

              <el-sub-menu index="research">
                <template #title>Research</template>
                <el-menu-item index="/research-direction">Research Direction</el-menu-item>
                <el-menu-item index="/research-projects">Research Projects</el-menu-item>
                <el-menu-item>
                  <a href="https://github.com/BaoyaoGroup" target="_blank" rel="noopener noreferrer">Github-Repositories</a>
                </el-menu-item>
              </el-sub-menu>

              <el-sub-menu index="our-team">
                <template #title>Our Team</template>
                <el-menu-item index="/dr-Baoyao-Yang">Baoyao Yang</el-menu-item>
                <el-menu-item index="/our-group">Our Group</el-menu-item>
                <el-menu-item index="/undergraduate">Undergraduate Education</el-menu-item>
                <el-menu-item index="/Alumni">Alumni</el-menu-item>
              </el-sub-menu>

              <el-sub-menu index="publications">
                <template #title>Publications</template>
                <el-menu-item index="/international-journals-conferences">
                  International Journals/Conferences
                </el-menu-item>
                <el-menu-item index="/patents">Patents</el-menu-item>
              </el-sub-menu>

              <!-- Project -->
              <el-sub-menu index="/Project">
                <template #title>Project</template>
                <el-menu-item index="/vKnow">vKnow</el-menu-item>
                <el-menu-item index="/VideoMind">Dataset</el-menu-item>
              </el-sub-menu>

              <el-menu-item index="/contact">Contact</el-menu-item>
            </el-menu>
          </div>
        </el-collapse-transition>
      </div>

      <!-- PC端导航（优化版） -->
      <el-header v-if="!isMobile" class="pc-header">
        <el-menu
            :default-active="activeIndex"
            class="el-menu-demo"
            mode="horizontal"
            :ellipsis="false"
            active-text-color="#7d1231"
            @select="handleSelect"
            :router="true"
            :popper-offset="16"
            popper-class="site-nav-popper"
            text-color="#2f3542"
        >
          <!-- Logo 区域 -->
          <div class="logo-container">
            <img
                class="logo-img"
                src="/logo/001.png"
                alt="BYML Logo"
            />
          </div>

          <!-- 菜单项 -->
          <el-menu-item index="/">Home</el-menu-item>
          <el-menu-item index="/about">About</el-menu-item>
          <el-menu-item index="/news">News</el-menu-item>

          <el-sub-menu index="/research">
            <template #title>Research</template>
            <el-menu-item index="/research-direction">Research Direction</el-menu-item>
            <el-menu-item index="/research-projects">Research Projects</el-menu-item>
            <el-menu-item>
              <a class="githublink" href="https://github.com/BaoyaoGroup"
                 target="_blank" rel="noopener noreferrer">Github-Repositories</a>
            </el-menu-item>
          </el-sub-menu>
          
          <el-sub-menu index="/our-team">
            <template #title>Our Team</template>
            <el-menu-item index="/dr-Baoyao-Yang">Baoyao Yang</el-menu-item>
            <el-menu-item index="/our-group">Our Group</el-menu-item>
            <el-menu-item index="/undergraduate">Undergraduate Education</el-menu-item>
            <el-menu-item index="/Alumni">Alumni</el-menu-item>
          </el-sub-menu>
          
          <el-sub-menu index="/publications">
            <template #title>Publications</template>
            <el-menu-item index="/international-journals-conferences">
              International Journals/Conferences
            </el-menu-item>
            <el-menu-item index="/patents">Patents</el-menu-item>
          </el-sub-menu>

          <!-- Project -->
          <el-sub-menu index="/Project">
            <template #title>Project</template>
            <el-menu-item index="/vKnow">vKnow</el-menu-item>
            <el-menu-item index="/VideoMind">Dataset</el-menu-item>
          </el-sub-menu>

          <el-menu-item index="/contact">Contact</el-menu-item>
        </el-menu>
      </el-header>

      <div v-if="!isMobile" style="height: 70px"></div>



      <el-main>
        <router-view/>
      </el-main>


      <!-- 页脚 -->
      <el-footer class="custom-footer">
        <div class="footer-content">
          <p>Welcome to BYML @ <a class="gdutlink" href="https://www.gdut.edu.cn/"
                                  style="text-decoration: none;color: white" target="_blank" rel="noopener noreferrer">Guangdong University of
            Technology</a></p>

          <p>Email: ybaoyao@gdut.edu.cn</p>

          <a class="BaoyaoGroupLink" href="https://github.com/BaoyaoGroup"
             style="text-decoration: none;color: white" target="_blank" rel="noopener noreferrer">Github-BYML</a>


          <p class="copyright">© 2025 By Baoyao Yang.</p>
          <!--          <div class="social-links" style="line-height: 100px">-->
          <!--            <a href="#" class="social-icon">Facebook</a>-->
          <!--            <a href="#" class="social-icon">Twitter</a>-->
          <!--            <a href="#" class="social-icon">Instagram</a>-->
          <!--          </div>-->
        </div>
      </el-footer>

    </el-container>
  </div>
</template>

<style scoped>
:global(body) {
  margin: 0;
}

:global(#app) {
  min-height: 100vh;
}


/* 新增移动端样式 */
@media (max-width: 860px) {
  .mobile-nav {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    z-index: 1002;
    background: rgba(255, 255, 255, 0.98);
    box-shadow: 0 1px 0 rgba(17, 24, 39, 0.08);
    height: 60px;
    padding-top: env(safe-area-inset-top);
  }

  .vertical-menu {
    border-right: none;
    padding: 8px 10px 24px;
    -webkit-tap-highlight-color: rgba(125, 18, 49, 0.12);
    --el-color-primary: #7d1231;
    --el-color-primary-light-3: #9f435c;
    --el-color-primary-light-5: #be7b8d;
    --el-color-primary-light-7: #deb9c3;
    --el-color-primary-light-8: #ead0d7;
    --el-color-primary-light-9: #f6e7eb;
    --el-color-primary-dark-2: #640e27;
    --el-menu-active-color: #7d1231;
    --el-menu-hover-text-color: #7d1231;
    --el-menu-hover-bg-color: rgba(125, 18, 49, 0.08);
    --el-menu-text-color: #2f3542;
  }

  .vertical-menu :deep(.el-menu-item),
  .vertical-menu :deep(.el-sub-menu__title) {
    height: 46px;
    line-height: 46px;
    margin: 2px 0;
    border-radius: 8px;
    font-size: 15px;
    font-weight: 500;
    color: #2f3542 !important;
    background: transparent !important;
    transition: background-color 0.18s ease, color 0.18s ease;
  }

  .vertical-menu :deep(.el-menu-item:hover),
  .vertical-menu :deep(.el-menu-item:focus),
  .vertical-menu :deep(.el-menu-item:active),
  .vertical-menu :deep(.el-sub-menu__title:hover),
  .vertical-menu :deep(.el-sub-menu__title:focus),
  .vertical-menu :deep(.el-sub-menu__title:active) {
    background: rgba(125, 18, 49, 0.08) !important;
    color: #7d1231 !important;
  }

  .vertical-menu :deep(.el-menu-item:hover a),
  .vertical-menu :deep(.el-menu-item:focus a),
  .vertical-menu :deep(.el-menu-item:active a) {
    color: #7d1231 !important;
  }

  .vertical-menu :deep(.el-sub-menu.is-opened > .el-sub-menu__title),
  .vertical-menu :deep(.el-sub-menu.is-opened > .el-sub-menu__title:hover),
  .vertical-menu :deep(.el-sub-menu.is-opened > .el-sub-menu__title:focus),
  .vertical-menu :deep(.el-sub-menu.is-opened > .el-sub-menu__title:active) {
    color: #7d1231 !important;
    background: rgba(125, 18, 49, 0.08) !important;
  }

  .vertical-menu :deep(.el-menu-item.is-active) {
    color: #fff !important;
    background: #7d1231 !important;
    font-weight: 600;
  }

  .vertical-menu :deep(.el-menu-item.is-active a) {
    color: #fff !important;
  }

  .vertical-menu :deep(.el-sub-menu.is-active > .el-sub-menu__title) {
    color: #7d1231 !important;
    background: rgba(125, 18, 49, 0.08) !important;
    font-weight: 600;
  }

  .vertical-menu :deep(.el-sub-menu__icon-arrow) {
    color: currentColor !important;
  }

  .vertical-menu :deep(.el-menu--inline) {
    padding: 2px 0 6px;
    background: #f8f9fb !important;
    border-radius: 8px;
  }

  .el-menu-item a {
    color: inherit !important;
    text-decoration: none;
    width: 100%;
  }

  .vertical-menu :deep(.el-sub-menu .el-menu-item) {
    height: 42px;
    line-height: 42px;
    padding-left: 40px !important;
    font-size: 14px;
  }

  .vertical-menu :deep(.el-sub-menu .el-menu-item.is-active) {
    background: #7d1231 !important;
    color: #fff !important;
  }

  .mobile-nav-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    box-sizing: border-box;
    padding: 0 16px;
    height: 60px;
  }

  .mobile-brand {
    display: inline-flex;
    align-items: center;
    gap: 9px;
    min-width: 0;
    color: #2f3542;
    text-decoration: none;
    font-weight: 700;
    letter-spacing: 0;
  }

  .mobile-logo-img {
    width: 96px;
    height: auto;
    object-fit: contain;
  }

  .hamburger-btn {
    width: 40px;
    height: 40px;
    padding: 0;
    margin-left: auto;
    border: 1px solid rgba(125, 18, 49, 0.28);
    border-radius: 8px;
    color: #7d1231;
    background: #fff;
    transition: background-color 0.18s ease, color 0.18s ease, border-color 0.18s ease;
  }

  .hamburger-btn:hover,
  .hamburger-btn:focus {
    color: #fff;
    border-color: #7d1231;
    background: #7d1231;
  }

  .hamburger-btn :deep(.el-icon) {
    color: currentColor;
  }

  .vertical-menu .el-sub-menu .el-menu {
    border-left: none !important;
  }

  .el-menu--vertical::after {
    display: none !important;
  }




  .custom-footer {
    padding: 15px 0;
    height: auto !important;
  }

  .common-layout {
    padding-bottom: env(safe-area-inset-bottom);
  }

  .el-main {
    margin-top: 0 !important; /* 固定顶部间距 */
    padding-top: 0 !important;
    transition: none !important;
    overflow-x: hidden !important; /* 禁用水平滚动 */
    overflow-y: hidden !important; /* 禁用垂直滚动 */
  }


  .mobile-menu-content {
    position: fixed;
    top: calc(60px + env(safe-area-inset-top));
    left: 0;
    right: 0;
    bottom: 0;
    background: #fff;
    z-index: 1000;
    overflow-y: auto;
    box-shadow: 0 18px 36px rgba(17, 24, 39, 0.14);
    height: calc(100dvh - 60px - env(safe-area-inset-top));
    max-height: calc(100dvh - 60px - env(safe-area-inset-top));
    overscroll-behavior: contain;
    transition:
        opacity 0.22s ease,
        transform 0.22s ease;
    will-change: transform;
    -webkit-transform: translateZ(0);
    transform: translateZ(0);
  }




  .common-layout {
    padding-top: calc(60px + env(safe-area-inset-top));
  }



  .custom-footer {
    position: relative;
    margin-top: auto;  /* 确保页脚贴底 */
  }

}

.logo-img {
  width: 180px; /* 根据需求调整宽度 */
  height: auto; /* 高度自适应，保持原始比例 */
  margin: 0 20px; /* 左右留白，避免贴边 */
}

.gdutlink {
  text-decoration: none;
  color: white;
  transition: color 0.3s ease; /* 让颜色变化时有 0.3 秒的过渡效果 */
}

.gdutlink:hover {
  color: #747d8c !important;
}

.BaoyaoGroupLink {
  text-decoration: none;
  color: white;
  transition: color 0.3s ease; /* 让颜色变化时有 0.3 秒的过渡效果 */
}

.BaoyaoGroupLink:hover {
  color: #747d8c !important;
}

.githublink:hover {
  color: white !important;
}

.logo {
  font-size: 36px;
}

.logo span {
  position: relative;
  font-weight: 500;
  text-transform: uppercase;
}

.logo span::before {
  content: "";
  position: absolute;
  bottom: -8px;
  left: 0;
  width: 100%;
  height: 3px;
  background: #7d1231;
  transform: scaleX(0);
  transition: transform 0.3s ease;
}

.logo:hover span::before {
  transform: scaleX(1);
}


.el-menu-demo {
  height: 70px;
  position: relative;
  z-index: 1002;
  border-bottom: none !important;
  --el-menu-active-color: #7d1231;
  --el-menu-hover-text-color: #7d1231;
  --el-menu-hover-bg-color: transparent;
  --el-menu-text-color: #2f3542;
}

.el-menu-demo::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 3px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.el-menu--horizontal {
  display: flex;
  align-items: center;
  gap: 2px; /* 控制菜单项之间的间距 */
}

/* PC 端导航容器样式 - 优化版 */
.el-header.pc-header {
  padding: 0;
  background: white;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
  position: fixed;
  top: 0;
  width: 100%;
  z-index: 1002;
  height: 70px;
  transition: all 0.3s ease;
  /* 性能优化：启用硬件加速 */
  will-change: transform;
  -webkit-transform: translateZ(0);
  transform: translateZ(0);
  /* 移除 contain 以避免影响 hover 效果 */
}

/* Logo 容器 */
.logo-container {
  display: flex;
  align-items: center;
  margin-right: auto;
  order: -1;
  margin-left: 20px;
  padding: 8px 0;
  transition: transform 0.3s ease;
}

.logo-container:hover {
  transform: scale(1.02);
}

/* Logo 图片优化 */
.logo-img {
  width: 150px;
  height: auto;
  object-fit: contain;
  transition: all 0.3s ease;
}

.el-menu-demo > .el-menu-item,
.el-menu-demo > .el-sub-menu {
  height: 70px;
  font-size: 15px;
  font-weight: 500;
  padding: 0 18px;
  margin: 0;
  border: none !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
  cursor: pointer;
  color: #2f3542 !important;
  background: transparent !important;
}

.el-menu-demo :deep(.el-sub-menu__title) {
  height: 70px;
  padding: 0;
  border: none !important;
  color: inherit !important;
  background: transparent !important;
}

.el-menu-demo > .el-menu-item:hover,
.el-menu-demo > .el-sub-menu:hover,
.el-menu-demo > .el-sub-menu.is-opened {
  color: #fff !important;
  background: #7d1231 !important;
}

.el-menu-demo > .el-menu-item.is-active,
.el-menu-demo > .el-sub-menu.is-active {
  color: #7d1231 !important;
  background: #fff !important;
  border-bottom: 3px solid #7d1231 !important;
  font-weight: 600;
}

.el-menu-demo > .el-menu-item.is-active:hover,
.el-menu-demo > .el-sub-menu.is-active:hover,
.el-menu-demo > .el-sub-menu.is-active.is-opened {
  color: #fff !important;
  background: #7d1231 !important;
}

.el-menu-demo > .el-menu-item::after,
.el-menu-demo > .el-sub-menu :deep(.el-sub-menu__title)::after {
  display: none !important;
}

.githublink {
  color: inherit;
  text-decoration: none;
  transition: color 0.3s ease;
}

.githublink:hover {
  color: inherit !important;
}


.el-carousel__item h3 {
  color: #475669;
  opacity: 0.75;
  line-height: 200px;
  margin: 0;
  text-align: center;
}


.footer-content {
  color: #fff; /* 设置文字为白色 */
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 8px;
  line-height: 1.4;
  min-height: 0;
  text-align: center;
}

.footer-content p {
  margin: 0;
}

.footer-content .copyright {
  margin-top: 4px;
}

.social-links {
  margin-top: 10px;
}

.social-icon {
  margin: 0 10px;
  text-decoration: none;
  color: #fff; /* 设置图标文字颜色为白色 */
}

.social-icon:hover {
  color: #35495e;
}

/* 自定义页脚的背景色 */
.custom-footer {
  align-items: center;
  background-color: #2c2c2c;
  box-sizing: border-box;
  color: white;
  display: flex;
  height: auto !important;
  justify-content: center;
  min-height: 120px;
  padding: 24px 20px;
}


/* 设置整个页面的布局为 flexbox */
.common-layout {
  display: flex;
  flex-direction: column;
  height: 100vh; /* 设置页面的总高度为视口高度 */
}

.site-container {
  flex: 1; /* 让容器占据剩余的空间 */
  display: flex;
  flex-direction: column;
}

/* 页脚保持在底部 */
.custom-footer {
  position: relative;
  width: 100%;
}

.social-links {
  margin-top: 10px;
}

.social-icon {
  margin: 0 10px;
  text-decoration: none;
  color: #fff;
}

.social-icon:hover {
  color: #35495e;
}
</style>

<style>
/* Element Plus popper is mounted outside App.vue scoped DOM. */
.site-nav-popper {
  border: none !important;
  border-radius: 8px !important;
  box-shadow: 0 10px 30px rgba(20, 24, 31, 0.16) !important;
  overflow: hidden;
}

.site-nav-popper .el-menu--popup {
  min-width: 210px;
  padding: 8px;
  border: none !important;
  background: #fff !important;
  --el-menu-active-color: #7d1231;
  --el-menu-hover-text-color: #fff;
  --el-menu-hover-bg-color: #7d1231;
  --el-menu-text-color: #2f3542;
}

.site-nav-popper .el-menu-item {
  height: 42px !important;
  line-height: 42px !important;
  margin: 2px 0 !important;
  padding: 0 14px !important;
  border-radius: 6px;
  color: #2f3542 !important;
  background: transparent !important;
  font-size: 14px;
  transition: background-color 0.18s ease, color 0.18s ease, transform 0.18s ease;
}

.site-nav-popper .el-menu-item:hover,
.site-nav-popper .el-menu-item:focus {
  color: #fff !important;
  background: #7d1231 !important;
  transform: translateX(3px);
}

.site-nav-popper .el-menu-item:hover a,
.site-nav-popper .el-menu-item:focus a {
  color: #fff !important;
}

.site-nav-popper .el-menu-item.is-active {
  color: #fff !important;
  background: #7d1231 !important;
  font-weight: 600;
}

.site-nav-popper .el-menu-item a {
  width: 100%;
  color: inherit !important;
  text-decoration: none;
}

.el-backtop {
  --el-backtop-text-color: #7d1231;
}

.el-backtop .el-icon,
.el-backtop .el-icon svg {
  color: #7d1231 !important;
}
</style>
