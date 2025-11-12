<script lang="ts" setup>
import type { TimelineItemProps } from 'element-plus'

interface ActivityType extends Partial<TimelineItemProps> {
  content: string
  type?: 'publication' | 'team' | 'award' | 'event' | 'general'
}

const activities2025: ActivityType[] = [
  {
    content:'Our paper "Multi-modal Brain Tumor Segmentation via Multi-category Interaction and Graph Co-reasoning" has been accepted by IEEE Transactions on Multimedia. Congrats, Li, Yang and co-authors!',
    timestamp:'2025-11-06',
    color:'#7d1231',
    type: 'publication'
  },
  {
    content:'Our paper "FedCD: A Hybrid Federated Learning Framework for Adaptive Training under Data Heterogeneity" has been accepted by PRCV. Congrats, Zhan, Yang and co-authors!',
    timestamp:'2025-10-22',
    color:'#7d1231',
    type: 'publication'
  },
  {
    content:'Our paper "CAM-interacted Vision GNN for Multi-label Medical Images" to be published in IEEE Journal of Biomedical and Health Informatics. Congrats, Wang, Yang and co-authors!',
    timestamp:'2025-10-16',
    color:'#7d1231',
    type: 'publication'
  },
  {
    content:'Our paper "Multi-Agent Reinforcement Learning Algorithm Using Dynamic OW-QMIX in Complex Supply Chain Scenarios" has been accepted by IEEE SMC. Congrats, Liu, Zhu and co-authors!',
    timestamp:'2025-07-20',
    color:'#7d1231',
    type: 'publication'
  },
  {
    content:'Two of our papers were accepted by ECAI: "FairFed++: Closing the Fairness Gap in Federated Learning through Self-Evolving Clustered Optimization" and "Unlocking the Potential of mLLMs: Enhancing Video-Text Retrieval through Caption Supplementation and Conical Embedding Optimization".',
    timestamp:'2025-07-11',
    color:'#7d1231',
    type: 'publication'
  },
  {
    content:'Our paper "Simple but Effective: Sub-Volume Contrastive Learning for Class-Imbalanced Semi-Supervised 3D Medical Image Segmentation" has been accepted by ACM Multimedia. Congrats, Xu, Yang and co-authors!',
    timestamp:'2025-07-06',
    color:'#7d1231',
    type: 'publication'
  },
  {
    content:'Our paper "Harnessing Feature Distribution Consistency for Federated Learning with Noisy" has been accepted by IEEE International Conference on Image Processing (ICIP). Congrats, Ma, Yang and co-authors!',
    timestamp:'2025-05-20',
    color:'#7d1231',
    type: 'publication'
  },
  {
    content:'Warm welcome to new group members Aoqi Yan, Guangyang Lin, Haifeng Lin, Jiahao Lian, Xiaojie Chen, Xi Wang & Yuhao Chen',
    timestamp:'2025-05-06',
    color:'#7d1231',
    type: 'team'
  },
  {
    content:'Beyond Machine Learning team introduction official website officially launched!',
    timestamp:'2025-04-08',
    color:'#7d1231',
    type: 'event'
  },
  {
    content:'Two of our papers were accepted by ICME: "Unifying Spatio-Temporal Contexts for Advanced Text-Video Retrieval" and "Action Decomposition-based Actor-Critic for Supply Chain Optimization".',
    timestamp:'2025-03-21',
    color:'#7d1231',
    type: 'publication'
  },
  {
    content: 'Our paper "Image-assisted Label Connective Completion for Vessel Segmentation with Insufficient Annotations" has been accepted by ICASSP. Congrats, Zheng, Yang and co-authors!',
    timestamp: '2025-03-08',
    color: '#7d1231',
    type: 'publication'
  },
  {
    content: 'Happy Chinese New Year!',
    timestamp: '2025-01-29',
    color: '#7d1231',
    type: 'event'
  }
]

const getNewsIcon = (type: string = 'general') => {
  const icons = {
    publication: '📄',
    team: '👥',
    award: '🏆',
    event: '🎉',
    general: '📢'
  }
  return icons[type as keyof typeof icons] || icons.general
}
</script>

<template>
  <div class="news-container">
    <!-- 页面标题区域 -->
    <div class="page-header">
      <h1 class="page-title">News & Announcements</h1>
      <p class="page-subtitle">Latest updates and achievements from our research team</p>
    </div>

    <!-- 年份标签 -->
    <div class="year-section">
      <div class="year-badge">
        <span class="year-icon">📅</span>
        2025
      </div>

      <!-- 时间线 -->
      <el-timeline class="timeline-container">
        <el-timeline-item
            v-for="(activity, index) in activities2025"
            :key="index"
            :timestamp="activity.timestamp"
            :color="activity.color"
            class="news-item"
        >
          <div class="news-content-wrapper">
            <div class="news-icon">{{ getNewsIcon(activity.type) }}</div>
            <div class="news-content">
              {{ activity.content }}
            </div>
          </div>
        </el-timeline-item>
      </el-timeline>
    </div>

    <el-backtop class="mobile-backtop" :right="100" :bottom="100"/>
  </div>
</template>

<style scoped>
.news-container {
  max-width: 1000px;
  margin: 0 auto;
  padding: 40px 20px;
  min-height: 100vh;
}

/* 页面标题样式 */
.page-header {
  text-align: center;
  margin-bottom: 60px;
  padding: 0 20px;
}

.page-title {
  font-size: 3rem;
  color: #2c3e50;
  margin-bottom: 15px;
  font-weight: 700;
  background: linear-gradient(135deg, #7d1231, #3498db);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  letter-spacing: -0.5px;
}

.page-subtitle {
  font-size: 1.3rem;
  color: #7f8c8d;
  max-width: 500px;
  margin: 0 auto;
  line-height: 1.6;
  font-weight: 400;
}

/* 年份区域 */
.year-section {
  position: relative;
}

.year-badge {
  display: inline-flex;
  align-items: center;
  gap: 10px;
  background: linear-gradient(135deg, #7d1231, #e74c3c);
  color: white;
  padding: 12px 25px;
  border-radius: 25px;
  font-size: 1.2rem;
  font-weight: 600;
  margin-bottom: 40px;
  margin-left: 30px;
  box-shadow: 0 6px 20px rgba(125, 18, 49, 0.25);
}

.year-icon {
  font-size: 1.4rem;
}

/* 时间线容器 */
.timeline-container {
  max-width: 900px;
  margin: 0 auto;
  padding: 0 30px;
}

/* 新闻项样式 */
.news-item {
  margin-bottom: 25px;
  transition: all 0.3s ease;
}

.news-item:hover {
  transform: translateX(5px);
}

/* 简化后的新闻内容包装 */
.news-content-wrapper {
  display: flex;
  align-items: flex-start;
  gap: 15px;
  padding: 20px 0;
  border-bottom: 1px solid #f0f0f0;
  transition: all 0.3s ease;
}

.news-item:hover .news-content-wrapper {
  border-bottom-color: #e0e0e0;
  padding-left: 5px;
}

.news-icon {
  font-size: 1.5rem;
  margin-top: 2px;
  flex-shrink: 0;
  transition: transform 0.3s ease;
}

.news-item:hover .news-icon {
  transform: scale(1.1);
}

.news-content {
  color: #2c3e50;
  font-size: 16px;
  line-height: 1.7;
  margin: 0;
  word-break: break-word;
  white-space: normal;
  flex: 1;
}

/* 时间线节点样式 */
:deep(.el-timeline-item__node) {
  background-color: #7d1231 !important;
  border: 3px solid white;
  box-shadow: 0 0 0 2px #7d1231, 0 2px 8px rgba(125, 18, 49, 0.2);
  width: 16px !important;
  height: 16px !important;
  left: -1px;
  transition: all 0.3s ease;
}

:deep(.el-timeline-item:hover .el-timeline-item__node) {
  transform: scale(1.1);
  box-shadow: 0 0 0 2px #7d1231, 0 4px 12px rgba(125, 18, 49, 0.3);
}

:deep(.el-timeline-item__tail) {
  border-left-color: #e8e8e8 !important;
  left: 7px !important;
}

/* 时间戳样式 */
:deep(.el-timeline-item__timestamp) {
  color: #7d1231 !important;
  font-size: 14px !important;
  font-weight: 600;
  margin-bottom: 8px !important;
  padding-left: 10px !important;
}

/* 返回顶部按钮 */
.mobile-backtop {
  right: 100px;
  bottom: 100px;
}

/* 移动端适配 */
@media (max-width: 768px) {
  .news-container {
    padding: 25px 15px;
  }

  .page-title {
    font-size: 2.3rem;
  }

  .page-subtitle {
    font-size: 1.1rem;
    padding: 0 10px;
  }

  .year-badge {
    margin-left: 15px;
    font-size: 1.1rem;
    padding: 10px 20px;
  }

  .timeline-container {
    padding: 0 15px;
  }

  .news-content-wrapper {
    padding: 15px 0;
    gap: 12px;
  }

  .news-icon {
    font-size: 1.3rem;
  }

  .news-content {
    font-size: 15px;
    line-height: 1.6;
  }

  /* 时间线移动端调整 */
  :deep(.el-timeline) {
    padding-left: 10px !important;
  }

  :deep(.el-timeline-item__node) {
    width: 14px !important;
    height: 14px !important;
    left: 0 !important;
  }

  :deep(.el-timeline-item__tail) {
    left: 6px !important;
  }

  :deep(.el-timeline-item__timestamp) {
    font-size: 13px !important;
    padding-left: 5px !important;
  }

  .mobile-backtop {
    right: 20px !important;
    bottom: 80px !important;
  }
}

/* 小屏手机优化 */
@media (max-width: 480px) {
  .news-container {
    padding: 20px 10px;
  }

  .page-title {
    font-size: 2rem;
  }

  .page-header {
    margin-bottom: 40px;
  }

  .news-content-wrapper {
    padding: 12px 0;
  }

  .news-content {
    font-size: 14px;
  }
}

/* 大屏幕优化 */
@media (min-width: 1400px) {
  .news-container {
    max-width: 1100px;
  }

  .timeline-container {
    max-width: 950px;
  }
}

/* 确保图标颜色正确 */
::v-deep(.el-icon svg) {
  color: #7d1231 !important;
}
</style>