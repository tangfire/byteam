<template>
  <div class="about-container">
    <!-- 页面标题区域 -->
    <div class="page-header">
      <h1 class="page-title">Beyond Machine Learning Group</h1>
      <p class="page-subtitle">Advancing Multimodal Intelligence for Real-World Impact</p>
    </div>

    <el-card class="content-card" shadow="hover">
      <div class="research-content">

        <!-- 导语部分 -->
        <div class="overview-section">
          <div class="overview-header">
            <h2 class="section-title">Our Mission</h2>
            <div class="accent-line"></div>
          </div>
          <div class="lead-box">
            <p class="lead-text">
              We pioneer <span class="highlight">foundational theories and scalable frameworks</span> to overcome
              challenges in multimodal data across its entire lifecycle. Our research bridges the gap between
              theoretical innovation and practical deployment in complex real-world scenarios.
            </p>
          </div>
        </div>

        <!-- 研究领域 -->
        <div class="research-section">
          <div class="section-header">
            <h2 class="section-title">Research Thrusts</h2>
            <div class="accent-line"></div>
          </div>

          <div class="research-grid">
            <div class="research-category" data-category="representation">
              <div class="category-header">
                <div class="category-icon">🔍</div>
                <h3 class="category-title">Multimodal Representation Learning</h3>
              </div>
              <ul class="styled-list">
                <li>Cross-modal alignment of medical time-series, imaging, and clinical text</li>
                <li>Contrastive learning for video understanding (visual, audio, text)</li>
                <li>Spatiotemporal fusion of heterogeneous IoT sensor data</li>
                <li>Self-supervised pretraining for multimodal foundation models</li>
              </ul>
            </div>

            <div class="research-category" data-category="decentralized">
              <div class="category-header">
                <div class="category-icon">🛡️</div>
                <h3 class="category-title">Decentralized & Privacy-Aware AI</h3>
              </div>
              <ul class="styled-list">
                <li>Federated learning for cross-institutional collaboration</li>
                <li>Differential privacy in multimodal data sharing</li>
                <li>Automated lesion annotation with attention mechanisms</li>
                <li>Weakly supervised learning under label scarcity</li>
              </ul>
            </div>

            <div class="research-category" data-category="infrastructure">
              <div class="category-header">
                <div class="category-icon">⚙️</div>
                <h3 class="category-title">AI Infrastructure & Scalability</h3>
              </div>
              <ul class="styled-list">
                <li>Unified architectures for seamless modality integration</li>
                <li>Reinforcement learning for industrial optimization</li>
                <li>Interpretable AI with cross-modal reasoning</li>
                <li>Memory-efficient model deployment</li>
              </ul>
            </div>
          </div>
        </div>

        <!-- 技术特色 -->
        <div class="expertise-section">
          <div class="section-header">
            <h2 class="section-title">Core Expertise</h2>
            <div class="accent-line"></div>
          </div>
          <div class="expertise-grid">
            <div class="expertise-item">
              <div class="expertise-icon">🌐</div>
              <h4 class="expertise-title">Heterogeneity-Aware Learning</h4>
              <p class="expertise-desc">Federated optimization for distributed multimodal systems</p>
            </div>
            <div class="expertise-item">
              <div class="expertise-icon">🎯</div>
              <h4 class="expertise-title">Noise-Robust Annotation</h4>
              <p class="expertise-desc">Weak supervision techniques for imperfect labeling</p>
            </div>
            <div class="expertise-item">
              <div class="expertise-icon">🔒</div>
              <h4 class="expertise-title">Secure Perception</h4>
              <p class="expertise-desc">Privacy-preserving multimodal embeddings</p>
            </div>
            <div class="expertise-item">
              <div class="expertise-icon">🔄</div>
              <h4 class="expertise-title">Cross-Modal Transfer</h4>
              <p class="expertise-desc">Knowledge sharing across different data modalities</p>
            </div>
          </div>
        </div>

        <!-- 愿景 -->
        <div class="vision-section">
          <div class="section-header">
            <h2 class="section-title">Strategic Vision</h2>
            <div class="accent-line"></div>
          </div>
          <div class="vision-content">
            <div class="vision-text">
              <p>
                We are pioneering the next generation of <span class="highlight">omni-modal foundation models</span>
                that seamlessly unify temporal, visual, textual, and sensor modalities. Our work directly supports
                national priorities in critical domains:
              </p>
              <div class="vision-domains">
                <span class="domain-tag">Smart Healthcare</span>
                <span class="domain-tag">Industry 4.0</span>
                <span class="domain-tag">Sustainable AI</span>
                <span class="domain-tag">Edge Intelligence</span>
              </div>
              <p class="vision-quote">
                "Bridging multimodal intelligence with real-world impact through continuous innovation
                and cross-domain knowledge transfer."
              </p>
            </div>
          </div>
        </div>

      </div>
    </el-card>

    <!-- 底部间隔 -->
    <div class="bottom-spacer"></div>
    <el-backtop class="mobile-backtop" :right="100" :bottom="100"/>
  </div>
</template>

<style scoped>
.about-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 30px 20px;
  min-height: 100vh;
}

/* 页面标题区域 */
.page-header {
  text-align: center;
  margin-bottom: 50px;
  padding: 0 20px;
}

.page-title {
  font-size: 2.8rem;
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
  max-width: 600px;
  margin: 0 auto;
  line-height: 1.6;
  font-weight: 400;
}

/* 主卡片样式 */
.content-card {
  max-width: 1100px;
  margin: 0 auto;
  border-radius: 16px;
  box-shadow: 0 8px 40px rgba(0, 0, 0, 0.08);
  border: none;
  overflow: hidden;
  transition: all 0.3s ease;
}

.content-card:hover {
  box-shadow: 0 12px 50px rgba(0, 0, 0, 0.12);
}

.research-content {
  padding: 40px;
  font-family: 'Segoe UI', system-ui, sans-serif;
  line-height: 1.7;
  color: #2c3e50;
}

/* 章节标题样式 */
.section-header {
  margin-bottom: 30px;
}

.section-title {
  color: #2c3e50;
  font-size: 2rem;
  font-weight: 600;
  margin-bottom: 10px;
  letter-spacing: -0.3px;
}



/* 导语部分 */
.overview-section {
  margin-bottom: 50px;
}

.lead-box {
  background: linear-gradient(135deg, #f8f9fa, #ffffff);
  padding: 30px;
  border-radius: 12px;
  border-left: 5px solid #7d1231;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.05);
}

.lead-text {
  font-size: 1.25rem;
  line-height: 1.8;
  color: #2c3e50;
  margin: 0;
  text-align: center;
  font-weight: 400;
}

.highlight {
  color: #7d1231;
  font-weight: 600;
  background: linear-gradient(120deg, rgba(125, 18, 49, 0.1), rgba(52, 152, 219, 0.1));
  padding: 2px 6px;
  border-radius: 4px;
}

/* 研究领域网格 */
.research-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
  gap: 30px;
  margin-bottom: 40px;
}

.research-category {
  background: #fff;
  padding: 30px;
  border-radius: 12px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.06);
  border: 1px solid #f0f0f0;
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
}

.research-category::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 4px;
  background: linear-gradient(90deg, #7d1231, #3498db);
  opacity: 0;
  transition: opacity 0.3s ease;
}

.research-category:hover::before {
  opacity: 1;
}

.research-category:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.1);
}

.category-header {
  display: flex;
  align-items: center;
  margin-bottom: 20px;
}

.category-icon {
  font-size: 2rem;
  margin-right: 15px;
  flex-shrink: 0;
}

.category-title {
  color: #2c3e50;
  font-size: 1.4rem;
  font-weight: 600;
  margin: 0;
  line-height: 1.3;
}

/* 列表样式 */
.styled-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.styled-list li {
  position: relative;
  padding-left: 24px;
  margin-bottom: 12px;
  line-height: 1.6;
  color: #5a6c7d;
}

.styled-list li::before {
  content: "▶";
  color: #7d1231;
  position: absolute;
  left: 0;
  font-size: 12px;
  top: 2px;
}

/* 技术特色网格 */
.expertise-section {
  margin-bottom: 50px;
}

.expertise-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 25px;
}

.expertise-item {
  text-align: center;
  padding: 25px 20px;
  background: #f8f9fa;
  border-radius: 12px;
  transition: all 0.3s ease;
  border: 1px solid #e9ecef;
}

.expertise-item:hover {
  background: #fff;
  transform: translateY(-3px);
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.08);
}

.expertise-icon {
  font-size: 2.5rem;
  margin-bottom: 15px;
}

.expertise-title {
  color: #2c3e50;
  font-size: 1.1rem;
  font-weight: 600;
  margin-bottom: 10px;
}

.expertise-desc {
  color: #7f8c8d;
  font-size: 0.95rem;
  line-height: 1.5;
  margin: 0;
}

/* 愿景部分 */
.vision-section {
  margin-bottom: 30px;
}

.vision-content {
  background: linear-gradient(135deg, #f8f9fa, #e9ecef);
  padding: 35px;
  border-radius: 12px;
  border: 1px solid #e9ecef;
}

.vision-text {
  font-size: 1.1rem;
  line-height: 1.7;
  color: #2c3e50;
}

.vision-domains {
  display: flex;
  flex-wrap: wrap;
  gap: 12px;
  margin: 20px 0;
}

.domain-tag {
  background: linear-gradient(135deg, #7d1231, #e74c3c);
  color: white;
  padding: 8px 16px;
  border-radius: 20px;
  font-size: 0.9rem;
  font-weight: 500;
  box-shadow: 0 2px 8px rgba(125, 18, 49, 0.2);
}

.vision-quote {
  font-style: italic;
  color: #7d1231;
  font-size: 1.1rem;
  text-align: center;
  margin-top: 25px;
  padding: 20px;
  background: rgba(255, 255, 255, 0.7);
  border-radius: 8px;
  border-left: 4px solid #7d1231;
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
  .about-container {
    padding: 20px 15px;
  }

  .page-title {
    font-size: 2.2rem;
  }

  .page-subtitle {
    font-size: 1.1rem;
    padding: 0 10px;
  }

  .research-content {
    padding: 25px 20px;
  }

  .section-title {
    font-size: 1.6rem;
  }

  .lead-box {
    padding: 20px;
  }

  .lead-text {
    font-size: 1.1rem;
  }

  .research-grid {
    grid-template-columns: 1fr;
    gap: 20px;
  }

  .research-category {
    padding: 20px;
  }

  .category-header {
    flex-direction: column;
    text-align: center;
  }

  .category-icon {
    margin-right: 0;
    margin-bottom: 10px;
  }

  .expertise-grid {
    grid-template-columns: repeat(2, 1fr);
    gap: 15px;
  }

  .expertise-item {
    padding: 20px 15px;
  }

  .vision-content {
    padding: 25px 20px;
  }

  .vision-domains {
    justify-content: center;
  }

  .domain-tag {
    font-size: 0.8rem;
    padding: 6px 12px;
  }

  .mobile-backtop {
    right: 20px !important;
    bottom: 80px !important;
  }
}

/* 小屏手机优化 */
@media (max-width: 480px) {
  .about-container {
    padding: 15px 10px;
  }

  .page-title {
    font-size: 1.8rem;
  }

  .research-content {
    padding: 20px 15px;
  }

  .expertise-grid {
    grid-template-columns: 1fr;
  }

  .vision-text {
    font-size: 1rem;
  }
}

/* 大屏幕优化 */
@media (min-width: 1400px) {
  .about-container {
    max-width: 1300px;
  }

  .content-card {
    max-width: 1200px;
  }
}

/* 确保图标颜色正确 */
::v-deep(.el-icon svg) {
  color: #7d1231 !important;
}
</style>