<script setup lang="ts">
import { Promotion, Tickets, Collection} from '@element-plus/icons-vue';
</script>

<template>
  <el-space direction="vertical" :size="30" style="width: 100%">
    <div style="display: flex; flex-direction: column; align-items: center;">
      <h1 class="section-title">VideoMind: An Omni-Modal Video Dataset with <br>Intent Grounding for Deep-Cognitive Video Understanding</h1>

      <el-card class="content-card">
        <div class="project-content">
          <!-- 项目介绍 -->
          <div class="introduction">
            <h3 class="subsection-title first-subtitle">What is VideoMind?</h3>

            <p class="vision-text">
              VideoMind is <span class="highlight">a large-scale video-centric multimodal dataset </span>that can be used to
              learn powerful and transferable text-video representations for video understanding tasks such as video question answering and video retrieval.
              The corresponding description of each video corresponds to three layers, namely factual layer, abstract layer, intentional layer.
            </p>
            <ul class="styled-list">
              <li>The factual layer includes visual description, OCR of video frames, ASR of videos, description of audio, and the title of the original video</li>
              <li>The intentional layer consists of two parts: the intent of the video uploader and the intent of the video main character </li>
            </ul>
            <img src="/VideoMind/Examples-v2.jpg" class="example2-image">
          </div>

          <!-- 数据集统计 -->
          <div class="Dataset-statistics">
            <h3 class="subsection-title">Dataset statistics</h3>
            <p class="vision-text">
              The dataset contains <span class="highlight">110K</span> video samples, each of which is <span class="highlight">ac-companied by audio, as well as systematic and detailed textual descriptions.</span>
            </p>
            <img src="/VideoMind/caterogy.png" class="caterogy-image">
            <div class="pillar-container">
              <div class="pillar">
                <img src="/VideoMind/ASR_length.png" class="dataset-images">
              </div>
              <div class="pillar">
                <img src="/VideoMind/duration.png" class="dataset-images">
              </div>
              <div class="pillar">
                <img src="/VideoMind/OCR_length.png" class="dataset-images">
              </div>
            </div>
          </div>

          <!-- 相关链接 -->
          <div class="relevant-links">
            <h3 class="subsection-title">Download</h3>
            <!-- 水平排列的超链接 -->
            <div class="relevant-links-container" style="display: flex; gap: 70px; align-items: center;margin-left: 20px;">
              <!--                 第一个链接（Code）-->
              <div style="display: flex; align-items: center; ">
                <a href="https://github.com/cdx-cindy/VideoMind" style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                  <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                    <Promotion />
                  </el-icon>
                  <span style="font-size: 18px; font-weight: 500;">Code</span>
                </a>
              </div>

              <!-- 第二个链接（Paper） -->
              <div style="display: flex; align-items: center;">
                <a href="https://arxiv.org/abs/2507.18552"
                   style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                  <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                    <Tickets/>
                  </el-icon>
                  <span style="font-size: 18px; font-weight: 500;">Paper</span>
                </a>
              </div>

              <!-- 第三个链接（Data） -->
              <div style="display: flex; align-items: center;">
                <a href="https://opendatalab.com/Dixin/VideoMind"
                   style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                  <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                    <Collection/>
                  </el-icon>
                  <span style="font-size: 18px; font-weight: 500;">Data</span>
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
            <pre class="pillar citation-block">
@misc{yang2025videomindomnimodalvideodataset,
  title={VideoMind: An Omni-Modal Video Dataset with Intent Grounding for Deep-Cognitive Video Understanding},
  author={Baoyao Yang and Wanyun Li and Dixin Chen and Junxiang Chen and Wenbin Yao and Haifeng Lin},
  year={2025},
  eprint={2507.18552},
  archivePrefix={arXiv},
  primaryClass={cs.CV},
  url={https://arxiv.org/abs/2507.18552},
}</pre>
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
  .caterogy-image {
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

.example2-image,
.caterogy-image,
.dataset-images {
  width: 100%;
  height: auto;
  object-fit: contain;
  display: block;
}
</style>