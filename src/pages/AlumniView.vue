<script setup lang="ts">
// 组件逻辑（如有需要可在此添加）
</script>

<template>
  <div class="alumni-container">
    <!-- 研究生校友部分 -->
    <section class="alumni-section">
      <h2 class="section-title">Graduate Alumni</h2>
      <div class="alumni-grid">
        <div class="alumni-card">
          <div class="avatar-container">
            <el-avatar shape="square" :size="200" src="/avatar/YuehuiFan.jpg" />
          </div>
          <div class="alumni-info">
            <h3 class="name">Yuehui Fan</h3>
            <div class="graduation-info">
              <span class="graduation-icon">🎓</span>
              <span class="graduation-date">Graduated: 2025.07</span>
            </div>
          </div>
        </div>

        <div class="alumni-card">
          <div class="avatar-container">
            <el-avatar class="jianxuanHuang" shape="square" :size="200" src="/avatar/JianxuanHuang.jpg" />
          </div>
          <div class="alumni-info">
            <h3 class="name">Jianxuan Huang</h3>
            <div class="graduation-info">
              <span class="graduation-icon">🎓</span>
              <span class="graduation-date">Graduated: 2025.07</span>
            </div>
          </div>
        </div>

        <div class="alumni-card">
          <div class="avatar-container">
            <el-avatar shape="square" :size="200" src="/avatar/YuebinXie.jpg" />
          </div>
          <div class="alumni-info">
            <h3 class="name">Yuebin Xie</h3>
            <div class="graduation-info">
              <span class="graduation-icon">🎓</span>
              <span class="graduation-date">Graduated: 2025.07</span>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- 本科生校友部分 -->
    <section class="alumni-section">
      <h2 class="section-title">Undergraduate Research Alumni</h2>
      <div class="alumni-grid">
        <div class="alumni-card">
          <div class="avatar-container">
            <el-avatar class="JingchaoWang" shape="square" :size="200" src="/avatar/JingchaoWang.jpg" />
          </div>
          <div class="alumni-info">
            <h3 class="name">Jingchao Wang</h3>
            <div class="graduation-info">
              <span class="graduation-icon">🎓</span>
              <span class="graduation-date">Graduated: 2025.07</span>
            </div>
          </div>
        </div>

        <div class="alumni-card">
          <div class="avatar-container">
            <el-avatar class="DongzheLi" shape="square" :size="200" src="/avatar/DongzheLi.jpg" />
          </div>
          <div class="alumni-info">
            <h3 class="name">Dongzhe Li</h3>
            <div class="graduation-info">
              <span class="graduation-icon">🎓</span>
              <span class="graduation-date">Graduated: 2025.07</span>
            </div>
          </div>
        </div>

        <div class="alumni-card">
          <div class="avatar-container">
            <el-avatar class="XiaochenHe" shape="square" :size="200" src="/avatar/XiaochenHe.jpg" />
          </div>
          <div class="alumni-info">
            <h3 class="name">Xiaochen He</h3>
            <div class="graduation-info">
              <span class="graduation-icon">🎓</span>
              <span class="graduation-date">Graduated: 2025.07</span>
            </div>
          </div>
        </div>

        <div class="alumni-card">
          <div class="avatar-container">
            <el-avatar class="WeideZhan" shape="square" :size="200" src="/avatar/WeideZhan.jpg" />
          </div>
          <div class="alumni-info">
            <h3 class="name">Weide Zhan</h3>
            <div class="graduation-info">
              <span class="graduation-icon">🎓</span>
              <span class="graduation-date">Graduated: 2025.07</span>
            </div>
          </div>
        </div>

        <div class="alumni-card">
          <div class="avatar-container">
            <el-avatar class="ZhixiangFang" shape="square" :size="200" src="/avatar/ZhixiangFang.jpg" />
          </div>
          <div class="alumni-info">
            <h3 class="name">Zhixiang Fang</h3>
            <div class="graduation-info">
              <span class="graduation-icon">🎓</span>
              <span class="graduation-date">Graduated: 2025.07</span>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>

  <el-backtop class="mobile-backtop" :right="100" :bottom="100"/>
</template>

<style scoped>
.alumni-container {
  width: 100%;
  padding: 30px 0;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.alumni-section {
  width: 100%;
  max-width: 1200px;
  margin-bottom: 60px;
  padding: 0 20px;
}

.section-title {
  color: #7d1231;
  font-size: 28px;
  font-weight: 600;
  margin-bottom: 40px;
  text-align: center;
  position: relative;
  padding-bottom: 15px;
}

.section-title::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 50%;
  transform: translateX(-50%);
  width: 100px;
  height: 3px;
  background: #7d1231;
  border-radius: 2px;
}

.alumni-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(220px, 1fr));
  gap: 40px 30px;
  justify-content: center;
}

.alumni-card {
  display: flex;
  flex-direction: column;
  align-items: center;
  transition: all 0.3s ease;
  padding: 25px 20px;
  border-radius: 16px;
  background: white;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.05);
  position: relative;
  overflow: hidden;
}

.alumni-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, #7d1231, #a51c41);
  opacity: 0;
  transition: opacity 0.3s ease;
}

.alumni-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 12px 30px rgba(125, 18, 49, 0.15);
}

.alumni-card:hover::before {
  opacity: 1;
}

.avatar-container {
  margin-bottom: 20px;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
  position: relative;
}

.alumni-card:hover .avatar-container {
  box-shadow: 0 6px 18px rgba(0, 0, 0, 0.15);
  transform: scale(1.03);
}

.alumni-info {
  text-align: center;
  width: 100%;
}

.name {
  color: #2c2f3b;
  font-size: 1.4rem;
  font-weight: 600;
  margin: 0 0 10px 0;
  transition: color 0.3s ease;
  position: relative;
  display: inline-block;
}

.name::after {
  content: '';
  position: absolute;
  bottom: -5px;
  left: 0;
  width: 0;
  height: 2px;
  background: #7d1231;
  transition: width 0.3s ease;
}

.alumni-card:hover .name::after {
  width: 100%;
}

.graduation-info {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  margin-top: 8px;
}

.graduation-icon {
  font-size: 1.2rem;
}

.graduation-date {
  font-size: 0.9rem;
  color: #7f8c8d;
  font-weight: 500;
}

/* 响应式设计 */
@media (max-width: 1200px) {
  .alumni-grid {
    grid-template-columns: repeat(3, 1fr);
    gap: 35px 25px;
  }
}

@media (max-width: 900px) {
  .alumni-grid {
    grid-template-columns: repeat(2, 1fr);
    gap: 30px 20px;
  }

  .section-title {
    font-size: 26px;
  }
}

@media (max-width: 768px) {
  .mobile-backtop {
    right: 20px !important;
    bottom: 80px !important;
  }

  .alumni-grid {
    grid-template-columns: repeat(2, 1fr);
    gap: 25px 15px;
  }

  .section-title {
    font-size: 24px;
    margin-bottom: 30px;
  }

  .alumni-card {
    padding: 20px 15px;
  }

  .name {
    font-size: 1.3rem;
  }

  .alumni-container {
    padding: 20px 0;
  }
}

@media (max-width: 480px) {
  .alumni-grid {
    grid-template-columns: 1fr;
    gap: 20px;
  }

  .section-title {
    font-size: 22px;
  }

  .alumni-section {
    padding: 0 15px;
  }
}

/* 头像样式调整 */
::v-deep(.el-icon svg) {
  color: #7d1231 !important;
}

/* 通用头像样式 */
::v-deep(.el-avatar) {
  transition: transform 0.3s ease;
}

.alumni-card:hover ::v-deep(.el-avatar) {
  transform: scale(1.05);
}

/* 特定头像调整 */
::v-deep(.jianxuanHuang img) {
  object-position: center 5%;
  object-fit: cover;
}

::v-deep(.DongzheLi img) {
  object-position: center 50%;
  object-fit: cover;
}

::v-deep(.XiaochenHe img) {
  object-position: center 50%;
  object-fit: cover;
}

::v-deep(.JingchaoWang img) {
  object-position: center 60%;
  object-fit: cover;
}

::v-deep(.WeideZhan img) {
  object-position: center 40%;
  object-fit: cover;
}

::v-deep(.ZhixiangFang img) {
  object-position: center 5%;
  object-fit: cover;
}
</style>