<script setup lang="ts">
import { getPublicUndergraduates } from '../api/public'
import type { UndergraduateEducation } from '../api/client'
import { usePublicList } from '../composables/usePublicList'
import { fallbackUndergraduates } from '../data/fallbacks/publicContent'

const undergraduateStudents = usePublicList<UndergraduateEducation>({
  fallback: fallbackUndergraduates,
  load: async () => (await getPublicUndergraduates()).items,
  fallbackMessage: 'Using local undergraduate fallback data',
})
</script>

<template>
  <div class="undergraduate-container">
    <div class="header-section">
      <h1 class="page-title">Undergraduate Education</h1>
      <p class="page-subtitle">本科生培养成果展示</p>
    </div>

    <div class="students-content">
      <div class="students-grid">
        <el-card
            v-for="student in undergraduateStudents"
            :key="student.id"
            class="student-card"
        >
          <div class="student-header">
            <div class="student-info">
              <h3 class="student-name">{{ student.name }}</h3>
              <div class="student-meta">
                <span class="student-grade">{{ student.grade }}</span>
                <span class="student-major">{{ student.major }}</span>
              </div>
            </div>
            <div class="student-direction">
              <el-icon><Research /></el-icon>
              <span>{{ student.direction }}</span>
            </div>
          </div>

          <div class="achievements-section">
            <h4 class="achievements-title">培养成果</h4>
            <ul class="achievements-list">
              <li v-for="(achievement, index) in student.achievements" :key="index">
                <el-icon><Star /></el-icon>
                <span>{{ achievement }}</span>
              </li>
            </ul>
          </div>
        </el-card>
      </div>
    </div>
  </div>

  <div class="bottom-spacer"></div>
  <el-backtop class="mobile-backtop" :right="100" :bottom="100"/>
</template>

<style scoped>
.undergraduate-container {
  width: 100%;
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
  box-sizing: border-box;
}

.header-section {
  text-align: center;
  margin-bottom: 40px;
  padding-top: 20px;
}

.page-title {
  color: #7d1231;
  font-size: 32px;
  font-weight: 600;
  margin: 0 0 10px 0;
}

.page-subtitle {
  color: #666;
  font-size: 16px;
  margin: 0;
}

.students-content {
  padding: 20px 0;
}

.students-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(380px, 1fr));
  gap: 24px;
}

.student-card {
  border-radius: 12px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
  transition: all 0.3s ease;
  border: 1px solid #eaeaea;
  background: white;
}

.student-card:hover {
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.12);
  transform: translateY(-4px);
  border-color: #7d1231;
}

.student-header {
  padding: 20px 20px 15px;
  border-bottom: 2px solid #f5f5f5;
}

.student-info {
  margin-bottom: 12px;
}

.student-name {
  font-size: 22px;
  font-weight: 600;
  color: #7d1231;
  margin: 0 0 8px 0;
}

.student-meta {
  display: flex;
  gap: 12px;
  align-items: center;
}

.student-grade,
.student-major {
  font-size: 14px;
  color: #666;
  padding: 4px 10px;
  background: #f8f9fa;
  border-radius: 6px;
  font-weight: 500;
}

.student-direction {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 14px;
  color: #13393e;
  font-weight: 500;
}

.student-direction .el-icon {
  color: #7d1231;
  font-size: 16px;
}

.achievements-section {
  padding: 20px;
}

.achievements-title {
  font-size: 16px;
  font-weight: 600;
  color: #7d1231;
  margin: 0 0 15px 0;
  padding-bottom: 10px;
  border-bottom: 1px solid #eee;
}

.achievements-list {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.achievements-list li {
  display: flex;
  align-items: flex-start;
  gap: 10px;
  font-size: 14px;
  color: #333;
  line-height: 1.6;
}

.achievements-list li .el-icon {
  color: #f9ab00;
  font-size: 16px;
  flex-shrink: 0;
  margin-top: 2px;
}

.bottom-spacer {
  height: 50px;
}

.mobile-backtop {
  right: 100px;
  bottom: 100px;
}

/* 移动端适配 */
@media (max-width: 768px) {
  .undergraduate-container {
    padding: 0 15px;
  }

  .page-title {
    font-size: 26px;
  }

  .page-subtitle {
    font-size: 14px;
  }

  .students-grid {
    grid-template-columns: 1fr;
    gap: 16px;
  }

  .student-name {
    font-size: 20px;
  }

  .student-meta {
    flex-direction: column;
    gap: 6px;
    align-items: flex-start;
  }

  .achievements-list li {
    font-size: 13px;
  }

  .mobile-backtop {
    right: 20px !important;
    bottom: 80px !important;
  }
}

/* 小屏幕优化 */
@media (max-width: 480px) {
  .student-header,
  .achievements-section {
    padding: 15px;
  }

  .student-name {
    font-size: 18px;
  }

  .student-grade,
  .student-major,
  .student-direction {
    font-size: 13px;
  }
}

/* 确保图标颜色正确 */
::v-deep(.el-icon svg) {
  color: #7d1231 !important;
}
</style>
