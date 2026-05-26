<script setup lang="ts">
import { Collection, Link, Promotion, Tickets } from '@element-plus/icons-vue'
import { computed } from 'vue'

const content = {
  title: 'VideoMind: An Omni-Modal Video Dataset with\nIntent Grounding for Deep-Cognitive Video Understanding',
  introduction: {
    title: 'What is VideoMind?',
    text: 'VideoMind is a large-scale video-centric multimodal dataset that can be used to learn powerful and transferable text-video representations for video understanding tasks such as video question answering and video retrieval. The corresponding description of each video corresponds to three layers, namely factual layer, abstract layer, intentional layer.',
    items: [
      'The factual layer includes visual description, OCR of video frames, ASR of videos, description of audio, and the title of the original video',
      'The intentional layer consists of two parts: the intent of the video uploader and the intent of the video main character',
    ],
    image: '/VideoMind/Examples-v2.jpg',
  },
  statistics: {
    title: 'Dataset statistics',
    text: 'The dataset contains 110K video samples, each of which is ac-companied by audio, as well as systematic and detailed textual descriptions.',
    image: '/VideoMind/caterogy.png',
    charts: [
      { image: '/VideoMind/ASR_length.png', alt: 'ASR length' },
      { image: '/VideoMind/duration.png', alt: 'Duration' },
      { image: '/VideoMind/OCR_length.png', alt: 'OCR length' },
    ],
  },
  results: {
    title: 'Results',
    text: 'we present the cross-modal retrieval results of several standard video-centric foundation models, including InternVideo, UMT-L, CLIP-VIP, mPLUG-2, and VAST.',
    figures: [
      { title: 'Results of hybrid-cognitive text-to-video retrieval on VideoMind-3K', image: '/VideoMind/results001.jpg' },
      { title: 'Results of hybrid-cognitive video-to-text retrieval on VideoMind-3K', image: '/VideoMind/results002.jpg' },
    ],
  },
  links: [
    { type: 'code', label: 'Code', url: 'https://github.com/cdx-cindy/VideoMind' },
    { type: 'paper', label: 'Paper', url: 'https://arxiv.org/abs/2507.18552' },
    { type: 'data', label: 'Data', url: 'https://opendatalab.com/Dixin/VideoMind' },
  ],
  citation: `@misc{yang2025videomindomnimodalvideodataset,
  title={VideoMind: An Omni-Modal Video Dataset with Intent Grounding for Deep-Cognitive Video Understanding},
  author={Baoyao Yang and Wanyun Li and Dixin Chen and Junxiang Chen and Wenbin Yao and Haifeng Lin},
  year={2025},
  eprint={2507.18552},
  archivePrefix={arXiv},
  primaryClass={cs.CV},
  url={https://arxiv.org/abs/2507.18552},
}`,
}

const lines = computed(() => content.title.split('\n'))
const charts = computed(() => content.statistics.charts)
const figures = computed(() => content.results.figures)
const links = computed(() => content.links)
const iconForLink = (type: string) => {
  if (type === 'code') return Promotion
  if (type === 'paper') return Tickets
  if (type === 'data') return Collection
  return Link
}
</script>

<template>
  <el-space direction="vertical" :size="30" style="width: 100%">
    <div style="display: flex; flex-direction: column; align-items: center;">
      <h1 class="section-title">
        <template v-for="(line, index) in lines" :key="index">
          {{ line }}<br v-if="index < lines.length - 1">
        </template>
      </h1>

      <el-card class="content-card">
        <div class="project-content">
          <!-- 项目介绍 -->
          <div class="introduction">
            <h3 class="subsection-title first-subtitle">{{ content.introduction?.title }}</h3>

            <p class="vision-text">
              {{ content.introduction?.text }}
            </p>
            <ul class="styled-list">
              <li v-for="item in content.introduction?.items || []" :key="item">{{ item }}</li>
            </ul>
            <img :src="content.introduction?.image" class="example2-image">
          </div>

          <!-- 数据集统计 -->
          <div class="Dataset-statistics">
            <h3 class="subsection-title">{{ content.statistics?.title }}</h3>
            <p class="vision-text">
              {{ content.statistics?.text }}
            </p>
            <img :src="content.statistics?.image" class="caterogy-image">
            <div class="pillar-container">
              <div v-for="chart in charts" :key="chart.image" class="pillar">
                <img :src="chart.image" :alt="chart.alt" class="dataset-images">
              </div>
            </div>
          </div>

          <!-- 结果展示 -->
          <div class="Results">
            <h3 class="subsection-title">{{ content.results?.title }}</h3>
            <p class="vision-text">
              {{ content.results?.text }}
            </p>
            <div class="results-container">
              <template v-for="figure in figures" :key="figure.image">
                <h4 class="results-title">{{ figure.title }}</h4>
                <img :src="figure.image" class="results-image">
              </template>
            </div>
          </div>
          <!-- 相关链接 -->
          <div class="relevant-links">
            <h3 class="subsection-title">Download</h3>
            <!-- 水平排列的超链接 -->
            <div class="relevant-links-container" style="display: flex; gap: 70px; align-items: center;margin-left: 20px;">
              <div v-for="item in links" :key="item.label" style="display: flex; align-items: center; ">
                <a :href="item.url" style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                  <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                    <component :is="iconForLink(item.type)" />
                  </el-icon>
                  <span style="font-size: 18px; font-weight: 500;">{{ item.label }}</span>
                </a>
              </div>
            </div>
          </div>

          <!-- 引用 -->
          <div class="Citation">
            <h3 class="subsection-title">Citation</h3>
            <p class="vision-text">
              If you find this work useful in your research, please cite the following paper:
            </p>
            <pre class="pillar citation-block">{{ content.citation }}</pre>
          </div>
        </div>
      </el-card>
    </div>
  </el-space>

  <!-- 底部间隔 -->
  <div style="height: 50px"></div>
  <el-backtop class="mobile-backtop" :right="100" :bottom="100"/>
</template>

<style scoped>

/* 响应式设计 */
@media (max-width: 768px) {
  .section-title {
    font-size: 0.85rem !important;
  }

  .subsection-title{
    font-size: 1rem !important;
  }
  .results-title{
    font-size: 0.9rem !important;
  }
  .first-subtitle{
    margin-top: -1rem !important;
  }
  .content-card,
  .project-content,
  .introduction,
  .Dataset-statistics,
  .relevant-links,
  .Citation {
    width: 98%;
    max-width: 100%;
    margin: 0 auto;
    box-sizing: border-box;
    padding: 0.8rem 0.5rem;
  }

  .pillar-container {
    display: flex !important;
    flex-direction: column !important;
    gap: 1rem !important;
  }
  .pillar {
    max-width: 350px;
    margin: 0 auto;
  }

  .example2-image,
  .caterogy-image,
  .results-image {
    max-width: 360px;
    margin: 0 auto;
  }

  .dataset-images{
    max-width: 300px;
    margin: 0 auto;
  }

  .citation-block {
    font-size: 0.9rem !important;
  }

  .mobile-backtop {
    right: 20px !important;
    bottom: 80px !important;
  }

  .relevant-links-container {
    gap: 30px !important;
  }
}

/* 统一字体设置 */
.project-content {
  font-family: 'Segoe UI', system-ui, sans-serif;
  line-height: 1.7;
  color: #2c3e50;
}

/* 主标题 */
.section-title {
  color: #7d1231;
  font-size: 2.2rem;
  margin: 1rem 0 2rem;
  font-weight: 600;
  letter-spacing: -0.5px;
  text-align: center;
}

/* 卡片样式 */
.content-card {
  max-width: 1200px;
  width: 90%;
  margin: 0 auto;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
}

/* 小标题 */
.subsection-title {
  color: #7d1231;
  font-size: 1.5rem;
  margin: 2rem 0 1.5rem;
  padding-bottom: 0.5rem;
  border-bottom: 2px solid #eee;
}

/* 列表样式 */
.styled-list {
  list-style: none;
  padding-left: 1.2rem;
  margin: 0.8rem 0;
}

.styled-list li {
  position: relative;
  padding-left: 1.2rem;
  margin-bottom: 0.8rem;
}

.styled-list li::before {
  content: "•";
  color: #7d1231;
  position: absolute;
  left: 0;
  font-weight: bold;
}
/* 高亮文字 */
.highlight {
  color: #7d1231;
  font-weight: 600;
}

/* 特色模块 */
.pillar-container {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1.5rem;
  margin: 1.5rem 0;
}

.pillar {
  background: #f8f9fa;
  padding: 1.5rem;
  border-radius: 8px;
  border-left: 4px solid #7d1231;
}

.pillar-title {
  color: #7d1231;
  margin: 0 0 0.5rem;
  font-size: 1.1rem;
}
/* 结果展示标题 */
.results-title {
  margin-top: 0.5rem;
  margin-bottom: 0.5rem;
  text-align: center;
}


/* 保证 el-icon 的图标颜色在初始时是正确的 */
::v-deep(.el-icon svg) {
  color: #7d1231 !important;
}

/* 引用代码块 */
.citation-block {
  color: #2c3e50;
  font-family: 'Fira Mono', 'Consolas', 'Menlo', monospace;
  font-size: 1rem;
  box-shadow: 0 2px 8px rgba(44,62,80,0.06);
  overflow-x: auto;
  white-space: pre;
}
/* 图片样式 */
.example2-image,
.caterogy-image,
.results-image,
.dataset-images {
  width: 100%;
  height: auto;
  object-fit: contain;
  display: block;
}
</style>
