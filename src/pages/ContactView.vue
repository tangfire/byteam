<script setup lang="ts">
// 可在此处添加交互逻辑（如复制邮箱功能）
</script>

<template>
  <div class="contact-page">
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
    position: static; /* 移除fixed定位 */
    width: 100%;
    height: auto;
    min-height: calc(100vh - 180px);
    margin-top: 60px; /* 为固定菜单留出空间 */
    background-attachment: scroll; /* 移除固定背景 */
  }
}

.contact-page {
  position: fixed;
  top: 140px;
  left: 10px;
  width: 99%;
  height: calc(100vh - 130px - 190px); /* 使用视口高度减去顶部偏移 */
  background:
      linear-gradient(rgba(255,255,255,0.7), rgba(255,255,255,0.7)),
      url('/background/ContactBackground.png');
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
