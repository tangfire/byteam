<template>
  <div class="home-container">
    <!-- 背景图层 -->
    <div class="hero-background">
      <div class="gradient-overlay"></div>
    </div>

    <!-- 内容区块 -->
    <main class="hero-content">
      <h1 class="lab-title">
        <span class="lab-acronym">Beyond Machine Learning Group</span>
<!--        <span class="lab-fullname">Beyond Machine Learning Lab</span>-->
      </h1>

      <div class="slogan-container">
<!--        <p class="main-slogan">Where Machines Learn to Transcend</p>-->
        <p class="sub-slogan">Beyond Algorithms, Towards Omni-Intelligence</p>
      </div>

      <!-- 装饰性科技元素 -->
      <div class="tech-ornament">
        <div class="particle-dot"></div>
        <div class="neon-line"></div>
      </div>
    </main>

    <el-backtop :right="100" :bottom="100" />
  </div>
</template>

<style scoped>
@media (max-width: 768px) {
  .home-container {
    position: static;
    width: 100%;
    height: auto;
    min-height: calc(100vh - 180px);
    margin-top: 60px;
  }





  .hero-content {
    padding: 4rem 1rem;
  }

  .lab-acronym {
    font-size: 1.6rem !important;
  }

  .sub-slogan {
    font-size: 1.2rem !important;

  }

  .tech-ornament {
    display: none;


    .hero-background {
      position: relative; /* 改为相对定位 */
      background-attachment: scroll; /* 移除固定背景 */
    }
  }
}


/* 基础重置 */
.home-container {
  --primary-color: #7d1231;
  --secondary-color: #2c3e50;
  --accent-color: #13393e;
  --light-bg: #f8f9fa;
  --text-primary: #2c3e50;
  --text-secondary: #2c3e50;;

  position: fixed;
  top: 140px;
  left: 10px;
  width: 99%;
  height: calc(100vh - 130px - 190px); /* 使用视口高度减去顶部偏移 */
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
      url('/background/ContactBackground.png') center/cover fixed;
  z-index: 0;
}

.gradient-overlay {
  position: absolute;
  width: 100%;
  height: 100%;
  background: radial-gradient(circle at 50% 30%, rgba(248,249,250,0.6) 20%, rgba(248,249,250,0.9) 80%);
}

/* 主内容 */
.hero-content {
  position: relative;
  z-index: 1;
  padding: 8rem 2rem 4rem;
  max-width: 1200px;
  margin: 0 auto;
  text-align: center;
}

.lab-title {
  font-family: 'Poppins', sans-serif;
  margin-bottom: 2.5rem;
  line-height: 1.1;
}

.lab-acronym {
  display: block;
  font-size: 3.2rem;
  font-weight: 700;
  color: var(--primary-color);
  letter-spacing: -1.5px;
  text-shadow: 2px 2px 4px rgba(0,0,0,0.08);
  animation: float 3s ease-in-out infinite;
}

.lab-fullname {
  display: block;
  font-size: 1.8rem;
  font-weight: 500;
  color: var(--text-secondary);
  margin-top: 1rem;
}

/* 口号样式 */
.slogan-container {
  max-width: 800px;
  margin: 0 auto;
}

.main-slogan {
  font-family: 'Roboto Condensed', sans-serif;
  font-size: 2.2rem;
  font-weight: 600;
  color: var(--secondary-color);
  line-height: 1.3;
  margin: 2rem 0;
  padding: 1.5rem;
  background: rgba(255,255,255,0.9);
  border-radius: 12px;
  box-shadow: 0 4px 20px rgba(0,0,0,0.06);
  backdrop-filter: blur(8px);
}

.sub-slogan {
  font-size: 1.8rem;
  color: var(--text-secondary);
  letter-spacing: 0.5px;
  margin-top: 1.5rem;
  opacity: 0.9;
}

/* 动态装饰元素 */
.tech-ornament {
  position: relative;
  height: 100px;
  margin-top: 4rem;
}

.particle-dot {
  position: absolute;
  width: 6px;
  height: 6px;
  background: var(--primary-color);
  border-radius: 50%;
  animation: particleMove 3s infinite;
}

.neon-line {
  position: absolute;
  height: 2px;
  width: 50%;
  background: linear-gradient(90deg, transparent, var(--primary-color), transparent);
  filter: drop-shadow(0 0 4px rgba(125,18,49,0.3));
  animation: lineScan 4s linear infinite;
}

/* 动画 */
@keyframes float {
  0%, 100% { transform: translateY(0); }
  50% { transform: translateY(-10px); }
}

@keyframes particleMove {
  0% { transform: translate(0, 0); opacity: 0; }
  30% { opacity: 1; }
  100% { transform: translate(200px, -150px); opacity: 0; }
}

@keyframes lineScan {
  0% { left: -50%; }
  100% { left: 150%; }
}

/* 响应式设计 */
@media (max-width: 768px) {
  .lab-acronym {
    font-size: 3.5rem;
  }

  .main-slogan {
    font-size: 1.8rem;
    padding: 1rem;
  }

  .sub-slogan {
    font-size: 1.1rem;
  }
}
</style>