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
        >
          <el-menu-item index="/">
            <img
                style="width: 100px"
                src="/images/logo.svg"
                alt="Element logo"
            />
          </el-menu-item>

          <el-menu-item index="/">Home</el-menu-item>
          <el-sub-menu  index="/research">
            <template #title >Research</template>
            <el-menu-item index="/current-projects">Current Projects</el-menu-item>
          </el-sub-menu>
          <el-sub-menu index="/our-team">
            <template #title>Our Team</template>
            <el-menu-item index="/dr-Baoyao-Yang">Yang Bao Yao</el-menu-item>
            <el-menu-item index="3-2" >Our Group</el-menu-item>
          </el-sub-menu>
          <el-sub-menu index="4">
            <template #title>Publications</template>
            <el-menu-item index="4-1">
              <div style="line-height: 1.2; margin: 0;">
                International <br> Journals/Conferences
              </div>
            </el-menu-item>
          </el-sub-menu>

          <el-menu-item index="5">News</el-menu-item>

        </el-menu>
      </el-header>

      <div style="height: 100px"></div>

      <el-main>
        <router-view />
      </el-main>

      <!-- 页脚 -->
      <el-footer :height="120" class="custom-footer">
        <div class="footer-content">
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
.el-menu--horizontal > .el-menu-item:nth-child(1) {
  margin-right: auto;
}

.el-menu--horizontal {
  --el-menu-horizontal-height: 100px;
}

.demonstration {
  color: var(--el-text-color-secondary);
}

.el-carousel__item h3 {
  color: #475669;
  opacity: 0.75;
  line-height: 200px;
  margin: 0;
  text-align: center;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n + 1) {
  background-color: #d3dce6;
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