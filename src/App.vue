<script setup lang="ts">


import { ref } from 'vue'

const activeIndex = ref('1')

const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
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
          <el-menu-item>
            <img
                style="width: 100px"
                src="/images/logo.svg"
                alt="Element logo"
            />
          </el-menu-item>

          <el-menu-item index="/"><p style="font-size: 18px">Home</p></el-menu-item>
          <el-sub-menu  index="/research">
            <template #title ><p style="font-size: 18px">Research</p></template>
            <el-menu-item index="/current-projects">Current Projects</el-menu-item>
          </el-sub-menu>
          <el-sub-menu index="/our-team">
            <template #title><p style="font-size: 18px">Our Team</p></template>
            <el-menu-item index="/dr-Baoyao-Yang">Yang Bao Yao</el-menu-item>
            <el-menu-item index="3-2" >Our Group</el-menu-item>
          </el-sub-menu>
          <el-sub-menu index="4">
            <template #title><p style="font-size: 18px">Publications</p></template>
            <el-menu-item index="4-1">
              <div style="line-height: 1.2; margin: 0;">
                International <br> Journals/Conferences
              </div>
            </el-menu-item>
          </el-sub-menu>

          <el-menu-item index="5"><p style="font-size: 18px">News</p></el-menu-item>


          <el-menu-item index="6"><p style="font-size: 18px">Contact</p></el-menu-item>
        </el-menu>
      </el-header>

      <div style="height: 100px"></div>

      <el-main>
        <router-view />
      </el-main>

      <!-- 页脚 -->
      <el-footer :height="120" class="custom-footer">
        <div class="footer-content" style="height: 100px">
          <p>© 2025 Your Company Name. All Rights Reserved.</p>
          <div class="social-links">
            <a href="#" class="social-icon">Facebook</a>
            <a href="#" class="social-icon">Twitter</a>
            <a href="#" class="social-icon">Instagram</a>
          </div>
        </div>
      </el-footer>

    </el-container>
  </div>
</template>

<style scoped>


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
  margin-right: auto;
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
</style>