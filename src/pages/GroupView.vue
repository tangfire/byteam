<script setup lang="ts">
import { getPublicPeople } from '../api/public'
import type { Person } from '../api/client'
import { usePublicList } from '../composables/usePublicList'
import { fallbackGraduatePeople } from '../data/fallbacks/publicContent'

const members = usePublicList<Person>({
  fallback: fallbackGraduatePeople,
  load: async () => (await getPublicPeople('graduate')).items,
  fallbackMessage: 'Using local group fallback data',
})
</script>

<template>
  <div class="group-container">
    <h1 class="section-title">Graduate Students</h1>

    <div class="members-grid">
      <div v-for="member in members" :key="member.name" class="member-card">
        <div class="avatar-container">
          <el-avatar shape="square" :size="220" :src="member.avatarUrl" />
        </div>
        <div class="member-info">
          <p class="name">{{ member.name }}</p>
          <p v-if="member.research" class="work">{{ member.research }}</p>
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

::v-deep(.el-avatar) {
  transition: transform 0.3s ease;
}

.member-card:hover ::v-deep(.el-avatar) {
  transform: scale(1.05);
}
</style>
