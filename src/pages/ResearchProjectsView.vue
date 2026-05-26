<script setup lang="ts">
import { getPublicResearchProjects } from '../api/public'
import type { ResearchProject } from '../api/client'
import { usePublicList } from '../composables/usePublicList'
import { fallbackResearchProjects } from '../data/fallbacks/publicContent'

const researchProjects = usePublicList<ResearchProject>({
  fallback: fallbackResearchProjects,
  load: async () => (await getPublicResearchProjects()).items,
  fallbackMessage: 'Using local research project fallback data',
})

const displayProjectStatus = (project: ResearchProject) => project.projectStatus || project.status
</script>

<template>
  <div class="research-projects-container">
    <div class="header-section">
      <h1 class="page-title">Research Projects</h1>
    </div>

    <div class="projects-content">
      <div class="projects-list">
        <el-card
            v-for="project in researchProjects"
            :key="project.id"
            class="project-card"
            :class="{ 'status-completed': displayProjectStatus(project) === '已结题' }"
        >
          <div class="project-header">
            <h3 class="project-title">{{ project.title }}</h3>
            <div class="project-meta">
              <span class="project-role">{{ project.role }}</span>
              <span class="project-status" :class="{ 'status-completed': displayProjectStatus(project) === '已结题' }">
                {{ displayProjectStatus(project) }}
              </span>
            </div>
          </div>

          <div class="project-details">
            <div class="detail-item">
              <span class="detail-label">基金来源:</span>
              <span class="detail-content">{{ project.fund }}</span>
            </div>

            <div v-if="project.number" class="detail-item">
              <span class="detail-label">项目编号:</span>
              <span class="detail-content project-number">{{ project.number }}</span>
            </div>

            <div class="detail-item">
              <span class="detail-label">执行期限:</span>
              <span class="detail-content">{{ project.period }}</span>
            </div>

            <div class="detail-item">
              <span class="detail-label">项目经费:</span>
              <span class="detail-content project-amount">{{ project.amount }}</span>
            </div>
          </div>
        </el-card>
      </div>
    </div>
  </div>

  <div class="bottom-spacer"></div>
  <el-backtop class="mobile-backtop" :right="100" :bottom="100"/>
</template>

<style scoped>
/* 统一变量管理 */
:root {
  --primary-color: #7d1231;
  --primary-light: rgba(125, 18, 49, 0.1);
  --hover-color: #13393e;
  --ongoing-color: #27ae60;
  --completed-color: #95a5a6;
  --card-max-width: 1000px;
  --mobile-breakpoint: 768px;
}

/* 基础容器 */
.research-projects-container {
  width: 100%;
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
  box-sizing: border-box;
}

/* 页面标题区域 */
.header-section {
  text-align: center;
  margin-bottom: 30px;
  padding-top: 20px;
}

.page-title {
  color: #7d1231;
  font-size: 26px;
  font-weight: 600;
  margin: 0;
}

/* 项目列表 */
.projects-list {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

/* 项目卡片 */
.project-card {
  max-width: 1000px;
  margin: 0 auto;
  border-radius: 8px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
  border: 1px solid #eaeaea;
  width: 100%;
}

.project-card:hover {
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.15);
  transform: translateY(-2px);
}

.project-card.status-completed {
  opacity: 0.8;
}

/* 项目头部 */
.project-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 16px;
  padding: 20px 20px 0;
}

.project-title {
  font-size: 18px;
  font-weight: 600;
  color: #7d1231;
  line-height: 1.4;
  flex: 1;
  margin-right: 16px;
  margin-top: 0;
  margin-bottom: 0;
}

.project-meta {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  gap: 6px;
  flex-shrink: 0;
}

.project-role {
  background: rgba(125, 18, 49, 0.1);
  color: #7d1231;
  padding: 4px 12px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 500;
}

.project-status {
  padding: 4px 12px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 500;
  color: white;
  background: #27ae60;
}

.project-status.status-completed {
  background: #95a5a6;
}

/* 项目详情 */
.project-details {
  display: flex;
  flex-direction: column;
  gap: 12px;
  margin-bottom: 16px;
  padding: 0 20px 20px;
}

.detail-item {
  display: flex;
  align-items: flex-start;
  gap: 8px;
}

.detail-label {
  font-size: 14px;
  color: #666;
  font-weight: 500;
  min-width: 80px;
  flex-shrink: 0;
}

.detail-content {
  font-size: 14px;
  color: #333;
  line-height: 1.5;
}

.project-number {
  font-family: 'Courier New', monospace;
  background: #f8f9fa;
  padding: 2px 6px;
  border-radius: 4px;
}

.project-amount {
  font-weight: 600;
  color: #7d1231;
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
  .research-projects-container {
    padding: 0 15px;
  }

  .page-title {
    font-size: 22px;
    text-align: left;
    margin-left: 8px;
  }

  .project-header {
    flex-direction: column;
    gap: 12px;
    padding: 15px 15px 0;
  }

  .project-meta {
    flex-direction: row;
    align-items: center;
    width: 100%;
    justify-content: space-between;
  }

  .project-title {
    margin-right: 0;
    font-size: 16px;
  }

  .detail-item {
    flex-direction: column;
    gap: 4px;
  }

  .detail-label {
    min-width: auto;
  }

  .mobile-backtop {
    right: 20px !important;
    bottom: 80px !important;
  }

  .project-details {
    padding: 0 15px 15px;
  }

  .projects-list {
    gap: 15px;
  }
}

/* 小屏幕手机优化 */
@media (max-width: 480px) {
  .research-projects-container {
    padding: 0 10px;
  }

  .project-header,
  .project-details {
    padding-left: 10px;
    padding-right: 10px;
  }

  .project-title {
    font-size: 15px;
  }

  .detail-label,
  .detail-content {
    font-size: 13px;
  }
}

/* 确保图标颜色正确 */
::v-deep(.el-icon svg) {
  color: #7d1231 !important;
}
</style>
