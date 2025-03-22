<template>
  <el-space direction="vertical" :size="30" style="width: 100%">
    <div style="display: flex; flex-direction: column; align-items: center;">
      <h1 class="section-title">About Beyond Machine Learning Group</h1>

      <el-card class="content-card">
        <div class="research-content">

          <!-- 导语部分 -->
          <div class="overview-box">
            <p class="lead-text">
              Our team specializes in <span class="highlight">multimodal intelligence</span> and its real-world applications,
              with a unified focus on addressing challenges inherent to multimodal data across its lifecycle.
              We pioneer foundational theories and scalable frameworks to advance cross-modal synergy,
              privacy-aware collaboration, and robust learning in complex scenarios.
            </p>
          </div>

          <!-- 研究领域 -->
          <div class="research-section">
            <h3 class="subsection-title">Research Thrusts</h3>

            <div class="research-category">
              <h4 class="category-title">1. Multimodal Representation Learning & Fusion</h4>
              <ul class="styled-list">
                <li>Cross-modal alignment of medical time-series signals (e.g., ECG), imaging data, and clinical text</li>
                <li>Contrastive learning frameworks for short video understanding (integrating visual, audio, and textual modalities)</li>
                <li>Spatiotemporal fusion of heterogeneous sensor data in industrial IoT systems</li>
              </ul>
            </div>

            <div class="research-category">
              <h4 class="category-title">2. Decentralized Multimodal Computing</h4>
              <ul class="styled-list">
                <li>Privacy-preserving federated learning for cross-institutional multimodal collaboration</li>
                <li>Weakly/semi-supervised learning under label noise and scarcity</li>
                <li>Automated lesion annotation via multimodal attention mechanisms</li>
              </ul>
            </div>

            <div class="research-category">
              <h4 class="category-title">3. Omni-modal AI Infrastructure</h4>
              <ul class="styled-list">
                <li>Unified architectures for seamless modality integration</li>
                <li>Reinforcement learning-driven industrial scheduling</li>
                <li>Interpretable algorithm design with cross-modal reasoning</li>
              </ul>
            </div>
          </div>

          <!-- 技术特色 -->
          <div class="highlight-section">
            <h3 class="subsection-title">Technical Distinctiveness</h3>
            <div class="pillar-container">
              <div class="pillar">
                <h5 class="pillar-title">Heterogeneity-Aware Learning</h5>
                <p class="pillar-desc">Federated optimization for distributed systems</p>
              </div>
              <div class="pillar">
                <h5 class="pillar-title">Noise-Robust Annotation</h5>
                <p class="pillar-desc">Weak supervision for imperfect labels</p>
              </div>
              <div class="pillar">
                <h5 class="pillar-title">Secure Perception</h5>
                <p class="pillar-desc">Privacy-preserving embeddings</p>
              </div>
            </div>
          </div>

          <!-- 愿景 -->
          <div class="vision-section">
            <h3 class="subsection-title">Strategic Vision</h3>
            <p class="vision-text">
              Pioneering <span class="highlight">omni-modal foundation models</span> unifying temporal, visual,
              textual, and sensor modalities. Supporting national priorities in
              <span class="tag">smart healthcare</span> and <span class="tag">Industry 4.0</span> through
              continuous cross-modal knowledge transfer.
            </p>
          </div>

        </div>
      </el-card>
    </div>
  </el-space>

  <!-- 底部间隔 -->
  <div style="height: 50px"></div>
  <el-backtop :right="100" :bottom="100"/>
</template>

<style scoped>
/* 统一字体设置 */
.research-content {
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

/* 分类标题 */
.category-title {
  color: #2c3e50;
  font-size: 1.2rem;
  margin: 1.5rem 0 1rem;
  font-weight: 600;
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

/* 响应式设计 */
@media (max-width: 768px) {
  .section-title {
    font-size: 1.8rem;
  }

  .content-card {
    width: 95%;
    margin: 0 auto;
  }
}
</style>