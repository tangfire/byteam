<script setup lang="ts">


import {ref, onMounted, onBeforeUnmount} from 'vue'
import {Menu} from "@element-plus/icons-vue";


const activeIndex = ref('1')
const isMobile = ref(false)
const isMenuCollapsed = ref(true) // 移动端菜单折叠状态

const checkScreenSize = () => {
  // 增加touch事件检测作为辅助判断
  const isTouchDevice = 'ontouchstart' in window || navigator.maxTouchPoints > 0
  const viewportWidth = Math.max(document.documentElement.clientWidth || 0, window.innerWidth || 0)

  isMobile.value = viewportWidth <= 768 && isTouchDevice
}


// App.vue script部分修改
// 修改handleSelect方法
const handleSelect = () => {
  isMenuCollapsed.value = true
  // 添加滚动复位逻辑
  if (isMobile.value) {
    window.scrollTo({
      top: 0,
      behavior: 'smooth'
    })
  }
}

// 添加点击外部关闭菜单的逻辑
const clickOutsideHandler = (e: MouseEvent) => {
  const menu = document.querySelector('.mobile-menu-content')
  const button = document.querySelector('.hamburger-btn')

  if (
      menu &&
      !menu.contains(e.target as Node) &&
      !button?.contains(e.target as Node)
  ) {
    isMenuCollapsed.value = true
  }
}

onMounted(() => {
  // 添加全局点击监听
  document.addEventListener('click', clickOutsideHandler)
})

onBeforeUnmount(() => {
  document.removeEventListener('click', clickOutsideHandler)
})

// 添加窗口大小监听
// 修改后的onMounted逻辑
onMounted(() => {
  // 添加延迟确保DOM渲染完成
  setTimeout(() => {
    checkScreenSize()
    // 强制触发resize事件
    window.dispatchEvent(new Event('resize'))
  }, 100)
  window.addEventListener('resize', checkScreenSize)
})

onBeforeUnmount(() => {
  window.removeEventListener('resize', checkScreenSize)
})

// 在代码中直接判断
if (import.meta.env.MODE === 'production') {
  console.log('生产环境');
} else if (import.meta.env.MODE === 'test') {
  console.log('测试环境');
} else {
  console.log('开发环境');
}


</script>

<template>
  <div class="common-layout">


    <el-container>

      <!-- 移动端导航 -->
      <div v-if="isMobile" class="mobile-nav">
        <div class="mobile-nav-header">
          <img class="logo-img" src="/logo/001.png" alt="BYML Logo"/>
          <el-button
              @click="isMenuCollapsed = !isMenuCollapsed"
              class="hamburger-btn"
              :style="{
        background: isMenuCollapsed ? 'transparent' : '#7d1231',
        color: isMenuCollapsed ? '#7d1231' : 'white'
      }"
          >
            <el-icon :size="24">
              <Menu/>
            </el-icon>
          </el-button>
        </div>

        <el-collapse-transition>
          <div v-show="!isMenuCollapsed" class="mobile-menu-content" :style="{top: isMobile ? '60px' : 'auto'}">
            <el-menu
                :default-active="activeIndex"
                active-text-color="#7d1231"
                @select="handleSelect"
                :router="true"
                class="vertical-menu"
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
                  <a href="https://github.com/BaoyaoGroup" target="_blank">Github-Repositories</a>
                </el-menu-item>
              </el-sub-menu>

              <el-sub-menu index="our-team">
                <template #title>Our Team</template>
                <el-menu-item index="/dr-Baoyao-Yang">Baoyao Yang</el-menu-item>
                <el-menu-item index="/our-group">Our Group</el-menu-item>
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
                 target="_blank">Github-Repositories</a>
            </el-menu-item>
          </el-sub-menu>
          
          <el-sub-menu index="/our-team">
            <template #title>Our Team</template>
            <el-menu-item index="/dr-Baoyao-Yang">Baoyao Yang</el-menu-item>
            <el-menu-item index="/our-group">Our Group</el-menu-item>
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
      <el-footer :height="120" class="custom-footer">
        <div class="footer-content" style="height: 100px">
          <p>Welcome to BYML @ <a class="gdutlink" href="https://www.gdut.edu.cn/"
                                  style="text-decoration: none;color: white" target="_blank">Guangdong University of
            Technology</a></p>

          <p>Email: ybaoyao@gdut.edu.cn</p>

          <a class="BaoyaoGroupLink" href="https://github.com/BaoyaoGroup"
             style="text-decoration: none;color: white" target="_blank">Github-BYML</a>


          <p style="margin-top: 20px">© 2025 By Baoyao Yang.</p>
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

/* 新增移动端样式 */
@media (max-width: 768px) {
  .mobile-nav {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    z-index: 1002; /* 提高 z-index 确保高于首页背景 */
    background: white;
    box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
    height: auto;
  }

  .el-menu-item a {
    color: #333b49 !important;  /* 与PC端保持一致 */
    text-decoration: none;

  }

  .el-menu-item a:hover {
    color: white !important;
  }

  .mobile-nav-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0 15px;
    height: 60px;
  }

  .logo-img {
    width: 120px !important; /* 移动端缩小logo */
    margin: 0 !important;
  }

  .hamburger-btn {
    padding: 4px;
    margin-left: auto;
  }

  .hamburger-btn:hover {
    background: #f5f5f5;
  }

  /* 调整子菜单项间距 */
  .el-sub-menu .el-menu-item {
    padding-left: 30px !important;
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

  .vertical-menu {
    border-right: none;
  }

  .common-layout {
    padding-bottom: env(safe-area-inset-bottom);
  }

  .vertical-menu .el-menu-item,
  .vertical-menu .el-sub-menu__title {
    height: 48px;
    line-height: 48px;
    font-size: 16px;
  }

  .el-main {
    margin-top: 0 !important; /* 固定顶部间距 */
    padding-top: 0 !important;
    transition: none !important;
    overflow-x: hidden !important; /* 禁用水平滚动 */
    overflow-y: hidden !important; /* 禁用垂直滚动 */
  }


  /* 新增穿透选择器 */
  .vertical-menu :deep(.el-menu-item) {
    background-color: rgba(0, 0, 0, 0) !important;
  }

  .vertical-menu :deep(.el-menu-item.is-active) {
    color: #7d1231 !important;
    background-color: transparent !important;
  }

  .vertical-menu :deep(.el-sub-menu.is-active) > .el-sub-menu__title {
    color: #7d1231 !important;
  }

  /* 子菜单激活项 */
  .vertical-menu :deep(.el-menu--inline) .el-menu-item.is-active {
    color: #7d1231 !important;
    background: transparent !important;
  }

  /* 悬停同步 */
  .vertical-menu :deep(.el-menu-item:hover),
  .vertical-menu :deep(.el-sub-menu__title:hover) {
    background-color: #7d1231 !important;
    color: #f2f2f2 !important;
  }

  .mobile-menu-content {
    position: fixed;
    top: 60px; /* 固定在导航栏下方 */
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(255, 255, 255, 0.98);
    backdrop-filter: blur(5px);
    z-index: 1000;
    overflow-y: auto;
    box-shadow: 0 8px 20px rgba(0, 0, 0, 0.1);
    /* 新增以下属性 */
    height: calc(100vh - 60px); /* 占据剩余屏幕高度 */
    max-height: calc(100vh - 60px); /* 新增最大高度限制 */
    overscroll-behavior: contain; /* 防止滚动链 */
    transition:
        opacity 0.3s cubic-bezier(0.4, 0, 0.2, 1),
        transform 0.3s cubic-bezier(0.4, 0, 0.2, 1),
        max-height 0.3s cubic-bezier(0.4, 0, 0.2, 1); /* 同步过渡属性 */
  }




  .common-layout {
    padding-top: 60px;  /* 将导航栏高度转为padding */
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


.el-menu {
  height: 70px;
  position: relative;
  z-index: 1002;
  border-bottom: none !important;
}

.el-menu::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 3px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.is-opened /deep/ .el-submenu__title {
  background: #F2F2F2 !important;
}

.el-submenu.is-active:not(.is-opened) /deep/ .el-submenu__title,
/deep/ .el-submenu__title:hover,
/deep/ .el-menu-item:hover,
/deep/ .el-menu-item.is-active {
  i {
    color: #F2F2F2 !important;
  }

  color: #F2F2F2 !important;
  background: #7d1231 !important;

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

/* 菜单项样式优化 */
.el-menu--horizontal > .el-menu-item,
.el-menu--horizontal > .el-sub-menu {
  height: 70px;
  line-height: 70px;
  font-size: 15px;
  font-weight: 500;
  padding: 0 18px;
  margin: 0;
  border-bottom: none !important;
  transition: all 0.3s ease;
}

.el-menu--horizontal > .el-menu-item:hover,
.el-menu--horizontal > .el-sub-menu:hover {
  background: transparent !important;
  color: var(--primary-color) !important;
}

.el-menu--horizontal > .el-menu-item.is-active {
  border-bottom: 2px solid var(--primary-color);
}

/* 子菜单标题特殊处理 - 使用最强选择器 */
.el-menu--horizontal .el-sub-menu > .el-sub-menu__title,
.el-menu--horizontal .el-sub-menu .el-sub-menu__title * {
  color: #2f3542 !important; /* 默认颜色 */
  transition: all 0.3s ease;
}

.el-menu--horizontal .el-sub-menu > .el-sub-menu__title:hover,
.el-menu--horizontal .el-sub-menu .el-sub-menu__title:hover,
.el-menu--horizontal .el-sub-menu .el-sub-menu__title:hover * {
  color: var(--primary-color) !important; /* hover 时变红 */
}

/* 子菜单弹出层优化 - 使用:deep() 穿透 */
.el-menu--horizontal .el-menu--popup {
  min-width: 200px;
  padding: 8px 0;
  border-radius: 8px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.15);
  border: none;
  background: white !important;
}

/* 子菜单项基础样式 */
.el-menu--horizontal :deep(.el-menu--popup .el-menu-item) {
  height: 45px !important;
  line-height: 45px !important;
  font-size: 14px;
  padding: 0 20px;
  margin: 0 !important;
  background: transparent !important;
  color: #2f3542 !important;
  transition: all 0.3s ease;
}

/* 子菜单项悬停效果 - 使用:deep() 穿透 */
.el-menu--horizontal :deep(.el-menu--popup .el-menu-item:hover) {
  background-color: var(--primary-light) !important;
  color: var(--primary-color) !important;
  font-weight: 600 !important;
}

/* 确保激活状态也生效 */
.el-menu--horizontal :deep(.el-menu--popup .el-menu-item.is-active) {
  color: var(--primary-color) !important;
  background: var(--primary-light) !important;
  font-weight: 600 !important;
}

/* 子菜单标题悬停效果 */
.el-menu--horizontal .el-sub-menu__title:hover {
  color: var(--primary-color) !important;
}

/* GitHub 链接样式 */
.githublink {
  color: inherit;
  text-decoration: none;
  transition: color 0.3s ease;
}

.githublink:hover {
  color: var(--primary-color) !important;
}


.el-carousel__item h3 {
  color: #475669;
  opacity: 0.75;
  line-height: 200px;
  margin: 0;
  text-align: center;
}


.footer-content {
  text-align: center;
  padding: 20px;
  color: #fff; /* 设置文字为白色 */
}

.footer-content p {
  margin: 5px 0;
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
  background-color: #2c2c2c;
  color: white;
  padding: 20px 0;
}


/* 设置整个页面的布局为 flexbox */
.common-layout {
  display: flex;
  flex-direction: column;
  height: 100vh; /* 设置页面的总高度为视口高度 */
}

.el-container {
  flex: 1; /* 让容器占据剩余的空间 */
  display: flex;
  flex-direction: column;
}

/* 页脚保持在底部 */
.custom-footer {
  background-color: #2c2c2c;
  color: white;
  padding: 20px 0;
  position: relative;
  bottom: 0;
  width: 100%;
}

.footer-content {
  text-align: center;
  padding: 20px;
  color: #fff;
}

.footer-content p {
  margin: 5px 0;
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