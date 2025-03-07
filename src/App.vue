<script setup lang="ts">


import { ref } from 'vue'

const activeIndex = ref('1')

const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}

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
      <el-header>


        <el-menu
            :default-active="activeIndex"
            class="el-menu-demo"
            mode="horizontal"
            :ellipsis="false"
            active-text-color="#7d1231"
            @select="handleSelect"
            :router="true"

            text-color="#2f3542"


        >

<!--          <img-->
<!--              style="width: 100px"-->
<!--              src="/currentprojects/logo.svg"-->
<!--              alt="Element logo"-->
<!--          />-->

          <div class="logo">

            <span>BY Group</span>
          </div>


          <el-menu-item index="/"><p style="font-size: 18px">Home</p></el-menu-item>
          <el-sub-menu  index="/research">
            <template #title ><p style="font-size: 18px">Research</p></template>
            <el-menu-item index="/current-projects">Current Projects</el-menu-item>
          </el-sub-menu>
          <el-sub-menu index="/our-team">
            <template #title><p style="font-size: 18px">Our Team</p></template>
            <el-menu-item index="/dr-Baoyao-Yang">Baoyao Yang</el-menu-item>
            <el-menu-item index="3-2" >Our Group</el-menu-item>
          </el-sub-menu>
          <el-sub-menu index="/publications">
            <template #title><p style="font-size: 18px">Publications</p></template>
            <el-menu-item index="/international-journals-conferences">
              <div style="line-height: 1.2; margin: 0;">
                International <br> Journals/Conferences
              </div>
            </el-menu-item>
          </el-sub-menu>

<!--          <el-menu-item index="5"><p style="font-size: 18px">News</p></el-menu-item>-->


          <el-menu-item index="/contact"><p style="font-size: 18px">Contact</p></el-menu-item>
        </el-menu>
      </el-header>

      <div style="height: 100px"></div>

      <el-main>
        <router-view />
      </el-main>


      <!-- 页脚 -->
      <el-footer :height="120" class="custom-footer">
        <div class="footer-content" style="height: 100px">
          <p>Welcome to BY Group @ Guangdong University of Technology</p>

          <p >Email: ybaoyao@gdut.edu.cn</p>
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

.logo{
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

  height: 120px;


}




.is-opened /deep/.el-submenu__title {
  background: #f4f4f4 !important;
}

.el-submenu.is-active:not(.is-opened) /deep/.el-submenu__title,
/deep/.el-submenu__title:hover,
/deep/.el-menu-item:hover,
/deep/.el-menu-item.is-active {
  i {
    color: #ffffff !important;
  }

  color: #ffffff !important;
  background: #7d1231 !important;
}


.el-menu--horizontal > .el-menu-item:nth-child(1) {
  /*margin-right: auto;*/
}

.el-menu--horizontal {
  display: flex;
  align-items: center;
  gap: 2px; /* 控制菜单项之间的间距 */
}

/* Logo保持靠左 */
.el-menu--horizontal > .logo {
  margin-right: auto;
  order: -1; /* 确保logo在最左侧 */
}

/* 菜单项靠右 */
.el-menu--horizontal > .el-menu-item,
.el-menu--horizontal > .el-sub-menu {
  margin-left: 0; /* 取消自动边距 */
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
  color: #fff;  /* 设置文字为白色 */
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
  color: #fff;  /* 设置图标文字颜色为白色 */
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
  height: 100vh;  /* 设置页面的总高度为视口高度 */
}

.el-container {
  flex: 1;  /* 让容器占据剩余的空间 */
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