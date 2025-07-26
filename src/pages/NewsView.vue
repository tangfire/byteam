<script lang="ts" setup>

import type { TimelineItemProps } from 'element-plus'

interface ActivityType extends Partial<TimelineItemProps> {
  content: string
}

const activities2025: ActivityType[] = [
  {
    content:'Our paper \"Multi-Agent Reinforcement Learning Algorithm Using Dynamic OW-QMIX in Complex Supply Chain Scenarios\" has been accepted by IEEE SMC.Congrats,Liu,Zhu and co-authors!',
    timestamp:'2025-07-20',
    color:'#7d1231',
  },

  {
    content:'Two of our papers were accepted by ECAI: \"FairFed++: Closing the Fairness Gap in Federated Learning through Self-Evolving Clustered Optimization\" and \"Unlocking the Potential of mLLMs: Enhancing Video-Text Retrieval through Caption Supplementation and Conical Embedding Optimization\".',
    timestamp:'2025-07-11',
    color:'#7d1231',
  },
  {
    content:'Our paper \"Simple but Effective: Sub-Volume Contrastive Learning for Class-Imbalanced Semi-Supervised 3D Medical Image Segmentation\" has been accepted by ACM Multimedia.Congrats,Xu,Yang and co-authors!',
    timestamp:'2025-07-06',
    color:'#7d1231',
  },
  {
    content:'Our paper \"Harnessing Feature Distribution Consistency for Federated Learning with Noisy\" has been accepted by IEEE International Conference on Image Processing (ICIP).Congrats,Ma,Yang and co-authors!',
    timestamp:'2025-05-20',
    color:'#7d1231',
  },

  {
    content:'Warm welcome to new group members AoqiYan,GuangyangLin,HaifengLin,JiahaoLian,XiaojieChen,XiWang & YuhaoChen',
    timestamp:'2025-05-06',
    color:'#7d1231',
  },
  {
    content:'Beyond Machine Learning team introduction official website officially launched!',
    timestamp:'2025-04-08',
    color:'#7d1231',
  },

  {
    content:'Two of our papers were accepted by ICME: \"Unifying Spatio-Temporal Contexts for Advanced Text-Video Retrieval\" and \"Action Decomposition-based Actor-Critic for Supply Chain Optimization\".',
    timestamp:'2025-03-21',
    color:'#7d1231',
  },

  {
    content: 'Our paper \"Image-assisted Label Connective Completion for Vessel Segmentation with Insufficient Annotations\" has been accepted by ICASSP.Congrats,Zheng,Yang and co-authors!',
    timestamp: '2025-03-08',
    color: '#7d1231',
  },
  {
    content: 'Happy Chinese New Year!',
    timestamp: '2025-01-29',
    color: '#7d1231'
  }


]
</script>

<template>
  <!-- 外层容器改为响应式宽度 -->
  <div class="main-container">
    <p class="year-title">2025</p>

    <!-- 时间线容器增加响应式 -->
    <el-timeline class="timeline-container">
      <el-timeline-item
          v-for="(activity, index) in activities2025"
          :key="index"
          :timestamp="activity.timestamp"
          :color="activity.color"
          class="news-item"
      >
        <div class="content-box">{{ activity.content }}</div>
      </el-timeline-item>
    </el-timeline>
  </div>

  <el-backtop class="mobile-backtop" :right="100" :bottom="100"/>
</template>



<style scoped>


/* 保证 el-icon 的图标颜色在初始时是正确的 */
::v-deep(.el-icon svg) {
  color: #7d1231 !important;
}

.news-item{
  color: #333; /* 正文文字颜色 */
  font-size: 16px; /* 增加字体大小 */
  white-space: pre-line;       /* 保留换行符并允许自动换行 */
  word-break: normal;          /* 保持单词完整性 */
  overflow-wrap: break-word;   /* 优先在单词间换行 */
  line-height: 1.6;            /* 增加行间距提升可读性 */
  max-width: 580px;            /* 根据容器宽度优化显示 */
}

::v-deep(.el-timeline-item) .el-timeline-item__content {
  white-space: pre-line;
  word-break: keep-all;
  overflow-wrap: normal;
  hyphens: none;
  max-width: 780px;
  line-height: 1.5;
  font-size: 16px;
  letter-spacing: 0.03em;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;  /* 增加边距防止贴边 */
}


/* 新增：主容器响应式 */
.main-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
  width: 100%;
  box-sizing: border-box;
}

.year-title {
  color: #7d1231;
  font-size: 26px;
  margin: 0 0 20px 100px;
}

/* 时间线容器响应式 */
.timeline-container {
  max-width: 800px;
  margin: 0 auto;
  width: 100%;
}

/* 移动端适配 */
@media (max-width: 768px) {
  .main-container {
    padding: 0 15px;
  }

  .timeline-container {
    padding-left: 10px;
    padding-right: 0;
  }

  .year-title {
    font-size: 22px;
    margin-left: 5px;
  }

  /* 时间戳移动端优化 */
  :deep(.el-timeline-item__timestamp) {
    display: block;
    margin-bottom: 8px;
    font-size: 14px;
    padding-left: 0 !important;
  }

  /* 内容框移动端优化 */
  .content-box {
    padding: 10px 15px;
    background: #f9f9f9;
    border-radius: 8px;
    margin-top: 5px;
  }
}

/* 内容区域优化 */
.news-item {
  font-size: 16px;
  line-height: 1.6;
  padding-bottom: 25px;
}

/* 内容换行优化 */
:deep(.el-timeline-item__content) {
  white-space: normal !important;
  word-break: break-word !important;
  overflow-wrap: anywhere !important;
}

/* 返回顶部按钮移动端适配 */
.mobile-backtop {
  right: 100px;
  bottom: 100px;
}
@media (max-width: 768px) {
  .mobile-backtop {
    right: 20px !important;
    bottom: 80px !important;
  }
}

/* 时间线图标颜色 */
:deep(.el-timeline-item__node) {
  background-color: #7d1231 !important;
}

</style>