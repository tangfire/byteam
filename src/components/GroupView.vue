<script setup lang="ts">
// 组件逻辑（如有需要可在此添加）
</script>

<template>
  <div class="members-container">
    <h2 class="section-title">Current Members</h2>
    <ul class="members-list">
      <li class="member-item">Dixin Chen</li>
      <li class="member-item">Huahong Deng</li>
      <li class="member-item">Canrong Du</li>
      <li class="member-item">Yuehui Fan</li>
      <li class="member-item">Jianxuan Huang</li>
      <li class="member-item">Yanhao Huang</li>
      <li class="member-item">Jingsong Lin</li>
      <li class="member-item">Yali Ma</li>
      <li class="member-item">Yanchao Tang</li>
      <li class="member-item">Kexin Xie</li>
      <li class="member-item">Yuebin Xie</li>
      <li class="member-item">Yufei Xu</li>
      <li class="member-item">Xianrun Xu</li>
      <li class="member-item">Xiaoqi Zheng</li>
      <li class="member-item">Sijia Zhou</li>
    </ul>
  </div>
</template>

<style scoped>
.members-container {
  max-width: 800px;
  margin: 2rem auto;
  padding: 0 1.5rem;
}

.section-title {
  color: #2c3e50;
  font-size: 1.8rem;
  border-bottom: 2px solid #7d1231;
  padding-bottom: 0.5rem;
  margin-bottom: 1.5rem;
}

.members-list {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 1rem;
  list-style: none;
  padding: 0;
}

.member-item {
  padding: 0.8rem 1.2rem;
  background-color: #f8f9fa;
  border-radius: 8px;
  transition: all 0.3s ease;
  color: #34495e;
  font-weight: 500;
}

.member-item:hover {
  background-color: #7d1231;
  color: white;
  transform: translateY(-2px);
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

@media (max-width: 600px) {
  .members-list {
    grid-template-columns: 1fr;
  }
}
</style>