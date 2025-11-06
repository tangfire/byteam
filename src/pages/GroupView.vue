<script setup lang="ts">
// 组件逻辑（如有需要可在此添加）
</script>

<template>
  <div class="group-container">
    <h1 class="section-title">Graduate Students</h1>

    <div class="members-grid">
      <!-- 第一行 -->
      <div class="member-card">
        <div class="avatar-container">
          <el-avatar class="DixinChen" shape="square" :size="220" src="/avatar/DixinChen.jpg" />
        </div>
        <div class="member-info">
          <p class="name">Dixin Chen</p>
          <!-- <p class="work">Short video understanding</p> -->
        </div>
      </div>

      <div class="member-card">
        <div class="avatar-container">
          <el-avatar class="HuahongDeng" shape="square" :size="220" src="/avatar/HuahongDeng.jpg" />
        </div>
        <div class="member-info">
          <p class="name">Huahong Deng</p>
          <!-- <p class="work">Industrial software</p> -->
        </div>
      </div>

      <div class="member-card">
        <div class="avatar-container">
          <el-avatar class="CanrongDu" shape="square" :size="220" src="/avatar/CanrongDu.jpg" />
        </div>
        <div class="member-info">
          <p class="name">Canrong Du</p>
          <!-- <p class="work">Short video understanding</p> -->
        </div>
      </div>

      <div class="member-card">
        <div class="avatar-container">
          <el-avatar shape="square" :size="220" src="/avatar/AoqiYan.jpg" />
        </div>
        <div class="member-info">
          <p class="name">Aoqi Yan</p>
          <!-- <p class="work">Short video understanding</p> -->
        </div>
      </div>

      <!-- 第二行 -->
      <div class="member-card">
        <div class="avatar-container">
          <el-avatar class="HaifengLin" shape="square" :size="220" src="/avatar/HaifengLin.jpg" />
        </div>
        <div class="member-info">
          <p class="name">Haifeng Lin</p>
        </div>
      </div>

      <div class="member-card">
        <div class="avatar-container">
          <el-avatar class="YanhaoHuang" shape="square" :size="220" src="/avatar/YanhaoHuang.jpg" />
        </div>
        <div class="member-info">
          <p class="name">Yanhao Huang</p>
        </div>
      </div>

      <div class="member-card">
        <div class="avatar-container">
          <el-avatar class="JingsongLin" shape="square" :size="220" src="/avatar/JingsongLin.jpg" />
        </div>
        <div class="member-info">
          <p class="name">Jingsong Lin</p>
        </div>
      </div>

      <div class="member-card">
        <div class="avatar-container">
          <el-avatar class="YaliMa" shape="square" :size="220" src="/avatar/YaliMa.jpg" />
        </div>
        <div class="member-info">
          <p class="name">Yali Ma</p>
        </div>
      </div>

      <!-- 第三行 -->
      <div class="member-card">
        <div class="avatar-container">
          <el-avatar class="YanchaoTang" shape="square" :size="220" src="/avatar/YanchaoTang.jpg" />
        </div>
        <div class="member-info">
          <p class="name">Yanchao Tang</p>
        </div>
      </div>

      <div class="member-card">
        <div class="avatar-container">
          <el-avatar class="KexinXie" shape="square" :size="220" src="/avatar/KexinXie.jpg" />
        </div>
        <div class="member-info">
          <p class="name">Kexin Xie</p>
        </div>
      </div>

      <div class="member-card">
        <div class="avatar-container">
          <el-avatar shape="square" :size="220" src="/avatar/YuhaoChen.jpg" />
        </div>
        <div class="member-info">
          <p class="name">Yuhao Chen</p>
        </div>
      </div>

      <div class="member-card">
        <div class="avatar-container">
          <el-avatar shape="square" :size="220" src="/avatar/YufeiXu.jpg" />
        </div>
        <div class="member-info">
          <p class="name">Yufei Xu</p>
        </div>
      </div>

      <!-- 第四行 -->
      <div class="member-card">
        <div class="avatar-container">
          <el-avatar class="XianrunXu" shape="square" :size="220" src="/avatar/XianrunXu.jpg" />
        </div>
        <div class="member-info">
          <p class="name">Xianrun Xu</p>
        </div>
      </div>

      <div class="member-card">
        <div class="avatar-container">
          <el-avatar class="XiaoqiZheng" shape="square" :size="220" src="/avatar/XiaoqiZheng.jpg" />
        </div>
        <div class="member-info">
          <p class="name">Xiaoqi Zheng</p>
        </div>
      </div>

      <div class="member-card">
        <div class="avatar-container">
          <el-avatar class="SijiaZhou" shape="square" :size="220" src="/avatar/SijiaZhou.jpg" />
        </div>
        <div class="member-info">
          <p class="name">Sijia Zhou</p>
        </div>
      </div>

      <div class="member-card">
        <div class="avatar-container">
          <el-avatar shape="square" :size="220" src="/avatar/XiaojieChen.jpg" />
        </div>
        <div class="member-info">
          <p class="name">Xiaojie Chen</p>
        </div>
      </div>

      <!-- 第五行 -->
      <div class="member-card">
        <div class="avatar-container">
          <el-avatar class="XiWang" shape="square" :size="220" src="/avatar/XiWang.jpg" />
        </div>
        <div class="member-info">
          <p class="name">Xi Wang</p>
        </div>
      </div>

      <div class="member-card">
        <div class="avatar-container">
          <el-avatar class="JiahaoLian" shape="square" :size="220" src="/avatar/JiahaoLian.jpg" />
        </div>
        <div class="member-info">
          <p class="name">Jiahao Lian</p>
        </div>
      </div>

      <div class="member-card">
        <div class="avatar-container">
          <el-avatar class="SijiaZhou" shape="square" :size="220" src="/avatar/GuangyangLin.jpg" />
        </div>
        <div class="member-info">
          <p class="name">Guangyang Lin</p>
        </div>
      </div>
    </div>
  </div>

  <el-backtop class="mobile-backtop" :right="100" :bottom="100"/>
</template>

<style scoped>
.group-container {
  width: 100%;
  padding: 20px 0;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.section-title {
  color: #7d1231;
  font-size: 32px;
  font-weight: 600;
  margin-bottom: 40px;
  text-align: center;
  position: relative;
  padding-bottom: 10px;
}

.section-title::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 50%;
  transform: translateX(-50%);
  width: 80px;
  height: 3px;
  background: linear-gradient(90deg, transparent, #7d1231, transparent);
}

.members-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 40px 30px;
  max-width: 1200px;
  width: 90%;
  margin: 0 auto;
}

.member-card {
  display: flex;
  flex-direction: column;
  align-items: center;
  transition: all 0.3s ease;
  padding: 15px;
  border-radius: 12px;
}

.member-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
  background-color: #f9f9f9;
}

.avatar-container {
  margin-bottom: 15px;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
}

.member-card:hover .avatar-container {
  box-shadow: 0 6px 18px rgba(0, 0, 0, 0.15);
}

.member-info {
  text-align: center;
}

.name {
  color: #2c2f3b;
  font-size: 1.5rem;
  font-weight: 600;
  margin: 0;
  transition: color 0.3s ease;
}

.member-card:hover .name {
  color: #7d1231;
}

.work {
  font-size: 1rem;
  color: #7f8c8d;
  margin: 5px 0 0 0;
}

.bottom-spacing {
  height: 50px;
}

/* 响应式设计 */
@media (max-width: 1200px) {
  .members-grid {
    grid-template-columns: repeat(3, 1fr);
    gap: 35px 25px;
  }
}

@media (max-width: 900px) {
  .members-grid {
    grid-template-columns: repeat(2, 1fr);
    gap: 30px 20px;
  }

  .section-title {
    font-size: 28px;
  }
}

@media (max-width: 768px) {
  .mobile-backtop {
    right: 20px !important;
    bottom: 80px !important;
  }

  .members-grid {
    grid-template-columns: repeat(2, 1fr);
    gap: 25px 15px;
    width: 95%;
  }

  .section-title {
    font-size: 26px;
    margin-bottom: 30px;
  }

  .member-card {
    padding: 10px;
  }

  .name {
    font-size: 1.3rem;
  }
}

@media (max-width: 480px) {
  .members-grid {
    grid-template-columns: 1fr;
    gap: 20px;
  }

  .section-title {
    font-size: 24px;
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

.member-card:hover ::v-deep(.el-avatar) {
  transform: scale(1.05);
}

/* 特定头像调整 */
::v-deep(.DixinChen img) {
  object-position: center 5%;
  object-fit: cover;
}

::v-deep(.HuahongDeng img) {
  object-position: center 10%;
  object-fit: cover;
}

::v-deep(.CanrongDu img) {
  object-position: center 5%;
  object-fit: cover;
}

::v-deep(.YanhaoHuang img) {
  object-position: center 35%;
  object-fit: cover;
}

::v-deep(.JingsongLin img) {
  object-position: center 20%;
  object-fit: cover;
}

::v-deep(.YaliMa img) {
  object-position: center 10%;
  object-fit: cover;
}

::v-deep(.YanchaoTang img) {
  object-position: center 15%;
  object-fit: cover;
}

::v-deep(.KexinXie img) {
  object-position: center 100%;
  object-fit: cover;
}

::v-deep(.XianrunXu img) {
  object-position: center 15%;
  object-fit: cover;
}

::v-deep(.XiaoqiZheng img) {
  object-position: center 100%;
  object-fit: cover;
}

::v-deep(.SijiaZhou img) {
  object-position: center 15%;
  object-fit: cover;
}

::v-deep(.XiWang img) {
  object-position: center 30%;
  object-fit: cover;
}

::v-deep(.JiahaoLian img) {
  object-position: center 20%;
  object-fit: cover;
}

::v-deep(.HaifengLin img) {
  object-position: center 15%;
  object-fit: cover;
}
</style>