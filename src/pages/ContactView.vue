<script setup lang="ts">
// 可在此处添加交互逻辑（如复制邮箱功能）
</script>

<template>
  <div class="contact-page">

    <!-- 背景图层 -->
    <div class="hero-background">
      <div class="gradient-overlay"></div>
    </div>

  <div class="contact-container">
    <h1 class="title">Contact Me</h1>

    <div class="availability">
      <p>At any time</p>
      <div class="divider"></div>
    </div>

    <p class="email-link">
      Email: ybaoyao@gdut.edu.cn
    </p>
  </div>
  </div>
</template>

<style scoped>

/* ContactView.vue 样式修改 */
@media (max-width: 768px) {
  .contact-page {
    position: static !important;
    width: 100%;
    height: auto;
    min-height: calc(100vh - 180px);
    margin-top: 0 !important;
    overflow-x: hidden !important; /* 禁用水平滚动 */
    overflow-y: hidden !important; /* 禁用垂直滚动 */
  }



  /* 标题样式 */
  .title {
    font-size: 2.6rem !important; /* 更大的字体 */
    font-weight: 800 !important; /* 更加加粗 */

  }

  /* 可用时间段落 */
  .availability {

    font-size: 1.4rem !important; /* 更大字体 */

  }

  .email-link{
    font-size: 1.2rem !important;
    margin-bottom: 200px;
  }

  .hero-background {
    position: relative; /* 改为相对定位 */
    background-attachment: scroll; /* 移除固定背景 */

  }

}

.contact-page {
  position: fixed;
  top: 140px;
  left: 10px;
  width: 99%;
  height: calc(100vh - 130px - 190px); /* 使用视口高度减去顶部偏移 */
  background-position: center center;
  background-size: cover; /* 保持cover属性 */
  background-attachment: scroll; /* 改为scroll让背景随滚动移动 */
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: auto;
}

.contact-container {
  max-width: 600px;
  width: 90%;
  margin: 0 auto;
  padding: 2rem;
  text-align: center;

}

/* 背景层 */
.hero-background {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background:
      linear-gradient(135deg, rgba(125,18,49,0.08) 0%, rgba(19,57,62,0.12) 100%),
      url('/background/ContactBackground.jpg') center/cover fixed;
  z-index: 0;
}

.gradient-overlay {
  position: absolute;
  width: 100%;
  height: 100%;
  background: radial-gradient(circle at 50% 30%, rgba(248,249,250,0.6) 20%, rgba(248,249,250,0.9) 80%);
}

/* 标题样式 */
.title {
  font-size: 4rem; /* 更大的字体 */
  font-weight: 800; /* 更加加粗 */
  color: #2c3e50; /* 设置正常颜色 */
  opacity: 0;
  animation: fadeInTitle 1.5s forwards; /* 标题渐显 */
  animation-delay: 0.5s; /* 延迟出现 */
}

/* 可用时间段落 */
.availability {
  margin: 2rem 0;
  font-size: 1.8rem; /* 更大字体 */
  color: #7d1231; /* 设置正常颜色 */
  font-weight: 600; /* 更加加粗 */
  opacity: 0;
  animation: fadeInText 2s forwards; /* 渐显效果 */
  animation-delay: 1s; /* 延迟出现 */
}

/* 分割线实现 */
.divider {
  width: 200px;
  height: 2px;
  background: #7d1231; /* 设置正常颜色 */
  margin: 1.5rem auto;
  position: relative;
  transition: all 0.3s ease;
}

/* 邮箱链接 */
.email-link {
  display: inline-block;
  font-size: 1.6rem; /* 略增字体 */
  color: #2c2f3b; /* 设置正常颜色 */
  text-decoration: none;
  padding: 1.2rem 2.5rem;
  border-radius: 25px;
  font-weight: 700; /* 加粗 */
  transition: background 0.3s ease, color 0.3s ease, transform 0.3s ease;
  opacity: 0;
  animation: fadeInText 2s forwards; /* 渐显效果 */
  animation-delay: 1.5s; /* 延迟出现 */
}


/* 渐变出现的动画 */
@keyframes fadeIn {
  to {
    opacity: 1;
  }
}

@keyframes fadeInTitle {
  to {
    opacity: 1;
  }
}

@keyframes fadeInText {
  to {
    opacity: 1;
  }
}
</style>
