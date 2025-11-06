<script setup lang="ts">
defineProps<{ msg: string }>()

// 项目数据
const researchProjects = [
  {
    id: 1,
    title: "弹性医学联邦学习与攻击防御关键技术研究",
    fund: "国家自然科学基金,面上项目",
    number: "62472105",
    period: "2025-01 至 2028-12",
    amount: "50万元",
    status: "在研",
    role: "主持"
  },
  {
    id: 2,
    title: "基于多模态时间序列表征学习的肝癌早期风险预测算法研究",
    fund: "国家自然科学基金,青年科学基金项目",
    number: "62102098",
    period: "2022-01 至 2024-12",
    amount: "30万元",
    status: "在研",
    role: "主持"
  },
  {
    id: 3,
    title: "产业聚集区域业务资源服务工业软件平台",
    fund: "国家重点研发计划，工业软件专项",
    number: "2023YFB3308700",
    period: "2023-12-01 至 2026-11-31",
    amount: "3500万（50万）",
    status: "在研",
    role: "子课题负责人"
  },
  {
    id: 4,
    title: "非侵入式冠状动脉病变的智能定位与评估关键技术研究",
    fund: "广东省自然科学基金委,广东省基础与应用基础研究基金面上项目",
    period: "2024-01 至 2026-12",
    amount: "15万",
    status: "在研",
    role: "主持"
  },
  {
    id: 5,
    title: "多模态医学大模型设计与医学影像自动标注技术研究",
    fund: "广东省自然科学基金委,广东省基础与应用基础研究基金面上项目",
    period: "2025-01 至 2027-12",
    amount: "10万",
    status: "在研",
    role: "主持"
  },
  {
    id: 6,
    title: "多源异构热处理数据智能联网及其保护的关键技术研究",
    fund: "广东省自然科学基金委,广东省基础与应用基础研究基金区域联合基金项目",
    number: "2022A1515140096",
    period: "2022-10 至 2025-09",
    amount: "30 万元",
    status: "在研",
    role: "校内主持"
  },
  {
    id: 7,
    title: "肝活检图像的多类病变细胞弱监督自动检测算法研究",
    fund: "广州市科技局,广州市基础与应用基础研究项目",
    number: "202201010266",
    period: "2022-04 至 2024-03",
    amount: "5 万元",
    status: "已结题",
    role: "主持"
  },
  {
    id: 8,
    title: "面向工业物联网的数据保护与协同学习关键技术研究",
    fund: "广东工业大学,交叉学科培育项目",
    period: "2022-07 至 2024-07",
    amount: "5 万元",
    status: "已结题",
    role: "主持"
  },
  {
    id: 9,
    title: "合成生物基因数据库的建立与查询工具的研发",
    fund: "广东省科技厅，重点领域研发计划'绿色生物制造'重点专项",
    period: "2024-01 至 2026-12",
    amount: "1000 万元",
    status: "在研",
    role: "参与"
  },
  {
    id: 10,
    title: "基于复杂异构特征医学数据挖掘的脓毒症智能预测方法研究",
    fund: "广东省自然科学基金委,广东省基础与应用基础研究面上项目",
    number: "2022A1515011592",
    period: "2022-01 至 2024-12",
    amount: "10 万元",
    status: "在研",
    role: "参与"
  }
]
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
            :class="{ 'status-completed': project.status === '已结题' }"
        >
          <div class="project-header">
            <h3 class="project-title">{{ project.title }}</h3>
            <div class="project-meta">
              <span class="project-role">{{ project.role }}</span>
              <span class="project-status" :class="{ 'status-completed': project.status === '已结题' }">
                {{ project.status }}
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
</style>