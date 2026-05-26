<script setup lang="ts">
import { computed } from 'vue'
import { getPublicPeople } from '../api/public'
import type { Person } from '../api/client'
import { usePublicList } from '../composables/usePublicList'
import { fallbackAlumni } from '../data/fallbacks/publicContent'

const alumni = usePublicList<Person>({
  fallback: fallbackAlumni,
  load: async () => (await getPublicPeople()).items.filter((item) => item.category.includes('alumni')),
  fallbackMessage: 'Using local alumni fallback data',
})

const graduateAlumni = computed(() => alumni.value.filter((item) => item.category === 'graduate_alumni'))
const undergraduateAlumni = computed(() => alumni.value.filter((item) => item.category === 'undergraduate_alumni'))
</script>

<template>
  <div class="alumni-container">
    <section class="alumni-section">
      <h2 class="section-title">Graduate Alumni</h2>
      <div class="alumni-grid">
        <div v-for="person in graduateAlumni" :key="person.name" class="alumni-card">
          <div class="avatar-container">
            <el-avatar shape="square" :size="200" :src="person.avatarUrl" />
          </div>
          <div class="alumni-info">
            <h3 class="name">{{ person.name }}</h3>
            <div class="graduation-info">
              <span class="graduation-icon">🎓</span>
              <span class="graduation-date">{{ person.graduationDate }}</span>
            </div>
          </div>
        </div>
      </div>
    </section>

    <section class="alumni-section">
      <h2 class="section-title">Undergraduate Research Alumni</h2>
      <div class="alumni-grid">
        <div v-for="person in undergraduateAlumni" :key="person.name" class="alumni-card">
          <div class="avatar-container">
            <el-avatar shape="square" :size="200" :src="person.avatarUrl" />
          </div>
          <div class="alumni-info">
            <h3 class="name">{{ person.name }}</h3>
            <div class="graduation-info">
              <span class="graduation-icon">🎓</span>
              <span class="graduation-date">{{ person.graduationDate }}</span>
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
}

.alumni-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
}

.avatar-container {
  margin-bottom: 18px;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.alumni-info {
  text-align: center;
}

.name {
  color: #2c2f3b;
  font-size: 1.3rem;
  font-weight: 600;
  margin: 0 0 10px;
}

.graduation-info {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 6px;
  color: #7f8c8d;
  font-size: 0.95rem;
}

@media (max-width: 768px) {
  .mobile-backtop {
    right: 20px !important;
    bottom: 80px !important;
  }

  .alumni-grid {
    grid-template-columns: repeat(auto-fill, minmax(180px, 1fr));
    gap: 25px 18px;
  }
}
</style>
