<script lang="ts" setup>

import type { TimelineItemProps } from 'element-plus'

interface ActivityType extends Partial<TimelineItemProps> {
  content: string
}

const activities2025: ActivityType[] = [
  {
    content:'Our research lab has achieved a significant milestone with two papers accepted to ICME 2025: \"Action Decomposition-based Actor-Critic for Supply Chain Optimization\" and \"Unifying Spatio-Temporal Contexts for Advanced Text-Video Retrieval\". We are proud to showcase our cutting-edge research at this prestigious conference.',
    timestamp:'2025-03-21',
    color:'#7d1231',
  },

  {
    content: 'Our paper \"Image-assisted Label Connective Completion for Vessel Segmentation with Insufficient Annotations\" is online now on ICASSP.Congrats,Zheng,Yang and co-authors!',
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
  <el-space direction="vertical" style="width: 80%">
  <div class="container">
    <p style="margin-left: 8px; color: #7d1231; font-size: 26px;">2025</p> <!-- 增加文字大小 -->

  <el-timeline style="max-width: 800px; margin: 0 auto">
    <el-timeline-item
        v-for="(activity, index) in activities2025"
        :key="index"
        :icon="activity.icon"
        :type="activity.type"
        :color="activity.color"
        :size="activity.size"
        :hollow="activity.hollow"
        :timestamp="activity.timestamp"
        class="news-item"
    >
      {{ activity.content }}
    </el-timeline-item>
  </el-timeline>
  </div>
  </el-space>
</template>



<style scoped>



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

</style>