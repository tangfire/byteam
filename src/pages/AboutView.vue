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
  animation: fadeInDown 1s ease-out;
}

@keyframes fadeInDown {
  from {
    opacity: 0;
    transform: translateY(-30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.page-subtitle {
  font-size: 1.3rem;
  color: #7f8c8d;
  max-width: 600px;
  margin: 0 auto;
  line-height: 1.6;
  font-weight: 400;
  animation: fadeInUp 1s ease-out 0.2s both;
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* 主卡片样式 */
.content-card {
  max-width: 1100px;
  margin: 0 auto;
  border-radius: 20px;
  box-shadow: 0 12px 48px rgba(0, 0, 0, 0.1), inset 0 1px 0 rgba(255, 255, 255, 0.95), 0 4px 12px rgba(0, 0, 0, 0.06);
  border: 2px solid rgba(125, 18, 49, 0.12);
  overflow: hidden;
  transition: all 0.5s cubic-bezier(0.34, 1.56, 0.64, 1);
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.98) 0%, rgba(248, 249, 250, 0.95) 100%);
  backdrop-filter: blur(20px);
}

.content-card:hover {
  box-shadow: 0 20px 70px rgba(0, 0, 0, 0.15), inset 0 1px 0 rgba(255, 255, 255, 0.98);
  transform: translateY(-4px);
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
  position: relative;
  display: inline-block;
}

.section-title::after {
  content: '';
  position: absolute;
  bottom: -10px;
  left: 0;
  width: 60px;
  height: 5px;
  background: linear-gradient(90deg, #7d1231 0%, #a51c41 50%, #7d1231 100%);
  border-radius: 3px;
  box-shadow: 0 2px 8px rgba(125, 18, 49, 0.3);
  animation: shimmer 3s ease-in-out infinite;
}

@keyframes shimmer {
  0%, 100% {
    opacity: 1;
  }
  50% {
    opacity: 0.7;
  }
}



/* 导语部分 */
.overview-section {
  margin-bottom: 50px;
}

.lead-box {
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.98) 0%, rgba(248, 249, 250, 0.95) 100%);
  padding: 35px;
  border-radius: 16px;
  border-left: 6px solid #7d1231;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1), inset 0 1px 0 rgba(255, 255, 255, 0.95), 0 4px 12px rgba(0, 0, 0, 0.05);
  transition: all 0.4s ease;
}

.lead-box:hover {
  box-shadow: 0 12px 48px rgba(0, 0, 0, 0.15), inset 0 1px 0 rgba(255, 255, 255, 0.98);
  transform: translateX(4px);
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
  background: linear-gradient(120deg, rgba(125, 18, 49, 0.12), rgba(52, 152, 219, 0.1));
  padding: 3px 8px;
  border-radius: 6px;
  box-shadow: 0 2px 8px rgba(125, 18, 49, 0.1);
}

/* 研究领域网格 */
.research-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
  gap: 30px;
  margin-bottom: 40px;
}

.research-category {
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.98) 0%, rgba(248, 249, 250, 0.95) 100%);
  padding: 30px;
  border-radius: 16px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.08), inset 0 1px 0 rgba(255, 255, 255, 0.95), 0 4px 12px rgba(0, 0, 0, 0.05);
  border: 2px solid rgba(125, 18, 49, 0.12);
  transition: all 0.5s cubic-bezier(0.34, 1.56, 0.64, 1);
  position: relative;
  overflow: hidden;
}

.research-category::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 5px;
  background: linear-gradient(90deg, #7d1231, #3498db, #7d1231);
  background-size: 200% 100%;
  opacity: 0;
  transition: opacity 0.5s ease;
  animation: gradientFlow 3s ease-in-out infinite;
}

@keyframes gradientFlow {
  0%, 100% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
}

.research-category:hover::before {
  opacity: 1;
}

.research-category:hover {
  transform: translateY(-8px) scale(1.02);
  box-shadow: 0 16px 48px rgba(0, 0, 0, 0.12), inset 0 1px 0 rgba(255, 255, 255, 0.98);
  border-color: #7d1231;
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
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.98) 0%, rgba(248, 249, 250, 0.95) 100%);
  border-radius: 16px;
  transition: all 0.5s cubic-bezier(0.34, 1.56, 0.64, 1);
  border: 2px solid rgba(125, 18, 49, 0.12);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.08), inset 0 1px 0 rgba(255, 255, 255, 0.95);
}

.expertise-item:hover {
  background: linear-gradient(135deg, rgba(255, 255, 255, 1) 0%, rgba(255, 255, 255, 0.98) 100%);
  transform: translateY(-8px) scale(1.05);
  box-shadow: 0 16px 48px rgba(0, 0, 0, 0.12), inset 0 1px 0 rgba(255, 255, 255, 0.98);
  border-color: #7d1231;
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
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.98) 0%, rgba(248, 249, 250, 0.95) 100%);
  padding: 35px;
  border-radius: 16px;
  border: 2px solid rgba(125, 18, 49, 0.12);
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.08), inset 0 1px 0 rgba(255, 255, 255, 0.95);
  transition: all 0.4s ease;
}

.vision-content:hover {
  box-shadow: 0 12px 48px rgba(0, 0, 0, 0.12), inset 0 1px 0 rgba(255, 255, 255, 0.98);
  transform: translateY(-2px);
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
  background: linear-gradient(135deg, #7d1231 0%, #e74c3c 100%);
  color: white;
  padding: 10px 18px;
  border-radius: 24px;
  font-size: 0.9rem;
  font-weight: 600;
  box-shadow: 0 4px 12px rgba(125, 18, 49, 0.3), inset 0 1px 0 rgba(255, 255, 255, 0.3);
  transition: all 0.4s cubic-bezier(0.34, 1.56, 0.64, 1);
  border: 2px solid rgba(255, 255, 255, 0.2);
}

.domain-tag:hover {
  transform: translateY(-3px) scale(1.08);
  box-shadow: 0 8px 20px rgba(125, 18, 49, 0.4), inset 0 1px 0 rgba(255, 255, 255, 0.4);
}

.vision-quote {
  font-style: italic;
  color: #7d1231;
  font-size: 1.1rem;
  text-align: center;
  margin-top: 25px;
  padding: 25px;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.95) 0%, rgba(248, 249, 250, 0.9) 100%);
  border-radius: 12px;
  border-left: 6px solid #7d1231;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.08), inset 0 1px 0 rgba(255, 255, 255, 0.95);
  transition: all 0.4s ease;
}

.vision-quote:hover {
  box-shadow: 0 12px 36px rgba(0, 0, 0, 0.12), inset 0 1px 0 rgba(255, 255, 255, 0.98);
  transform: translateX(4px);
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