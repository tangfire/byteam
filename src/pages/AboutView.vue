<script setup lang="ts">
const pageTitle = 'Beyond Machine Learning Group'

const content = {
  subtitle: 'Multimodal learning for healthcare, video understanding, and trustworthy distributed intelligence.',
  intro: [
    'The Beyond Machine Learning Group at Guangdong University of Technology studies machine learning methods that learn from multiple sources of evidence, including images, videos, signals, text, and distributed data.',
    'Our research is driven by real data conditions: limited labels, noisy annotations, privacy constraints, heterogeneous clients, and changing deployment environments. We build models and systems for medical image analysis, multimodal representation learning, federated learning, and robust training.',
  ],
  focusAreas: [
    {
      index: '01',
      title: 'Multimodal Medical Intelligence',
      text: 'Combining imaging, clinical signals, and structured records for diagnosis-oriented tasks such as segmentation, disease analysis, and decision support.',
    },
    {
      index: '02',
      title: 'Federated and Privacy-Aware Learning',
      text: 'Designing learning frameworks for cross-site collaboration when data cannot be centralized, with attention to heterogeneity, reliability, and privacy.',
    },
    {
      index: '03',
      title: 'Video and Cross-Modal Understanding',
      text: 'Learning representations that connect visual, audio, and textual cues for retrieval, reasoning, summarization, and content understanding.',
    },
    {
      index: '04',
      title: 'Robust Learning from Imperfect Data',
      text: 'Developing methods for weak supervision, label noise, missing modalities, and data imbalance in practical machine learning pipelines.',
    },
  ],
  principles: [
    'Start from concrete scientific or clinical problems, then choose the modeling tools that fit the data.',
    'Keep experiments reproducible, with clear baselines, ablations, and implementation details.',
    'Build collaborations across computer science, healthcare, and industrial scenarios.',
    'Value models that are useful, reliable, and understandable in real deployment settings.',
  ],
  keywords: [
    'Multimodal learning',
    'Medical image analysis',
    'Federated learning',
    'Weak supervision',
    'Video understanding',
    'Representation learning',
    'Privacy-preserving AI',
    'Robust optimization',
  ],
  vision: {
    title: 'What We Aim For',
    text: 'We aim to develop learning systems that connect modalities, institutions, and application domains while remaining practical enough to support real research and deployment needs.',
    domains: ['Healthcare AI', 'Distributed Learning', 'Multimodal Systems', 'Applied Machine Learning'],
  },
}
</script>

<template>
  <div class="about-container">
    <header class="about-hero">
      <p class="eyebrow">About BYML</p>
      <h1 class="page-title">{{ pageTitle }}</h1>
      <p class="page-subtitle">{{ content.subtitle }}</p>
    </header>

    <main class="about-main">
      <section class="intro-section" aria-labelledby="about-overview">
        <div class="section-meta">Overview</div>
        <div class="intro-copy">
          <h2 id="about-overview">Research grounded in real multimodal data</h2>
          <p v-for="paragraph in content.intro" :key="paragraph">{{ paragraph }}</p>
        </div>
      </section>

      <section class="focus-section" aria-labelledby="about-focus">
        <div class="section-heading">
          <span class="section-meta">Focus</span>
          <h2 id="about-focus">Main Research Directions</h2>
        </div>

        <div class="focus-list">
          <article v-for="item in content.focusAreas" :key="item.title" class="focus-row">
            <span class="focus-index">{{ item.index }}</span>
            <div class="focus-copy">
              <h3>{{ item.title }}</h3>
              <p>{{ item.text }}</p>
            </div>
          </article>
        </div>
      </section>

      <section class="work-section" aria-labelledby="about-work-style">
        <div class="section-heading work-heading">
          <span class="section-meta">Approach</span>
          <h2 id="about-work-style">How We Work</h2>
        </div>

        <div class="work-layout">
          <ol class="principle-list">
            <li v-for="(principle, index) in content.principles" :key="principle">
              <span>{{ String(index + 1).padStart(2, '0') }}</span>
              <p>{{ principle }}</p>
            </li>
          </ol>

          <aside class="keyword-panel" aria-label="Research keywords">
            <h2>Research Keywords</h2>
            <div class="keyword-list">
              <span v-for="keyword in content.keywords" :key="keyword">{{ keyword }}</span>
            </div>
          </aside>
        </div>
      </section>

      <section class="vision-section" aria-labelledby="about-vision">
        <div>
          <span class="section-meta">Vision</span>
          <h2 id="about-vision">{{ content.vision.title }}</h2>
        </div>
        <div class="vision-copy">
          <p>{{ content.vision.text }}</p>
          <div class="domain-list">
            <span v-for="domain in content.vision.domains" :key="domain">{{ domain }}</span>
          </div>
        </div>
      </section>
    </main>

    <el-backtop class="mobile-backtop" :right="100" :bottom="100"/>
  </div>
</template>

<style scoped>
.about-container {
  color: #273445;
  margin: 0 auto;
  max-width: 1180px;
  min-height: 100vh;
  padding: 42px 24px 72px;
}

.about-hero {
  border-bottom: 1px solid #e7d9df;
  margin-bottom: 46px;
  padding-bottom: 34px;
}

.eyebrow,
.section-meta {
  color: #7d1231;
  display: block;
  font-size: 0.95rem;
  font-weight: 700;
  letter-spacing: 0;
  margin-bottom: 10px;
}

.page-title {
  color: #7d1231;
  font-size: 3.2rem;
  font-weight: 700;
  letter-spacing: 0;
  line-height: 1.12;
  margin: 0;
}

.page-subtitle {
  color: #566273;
  font-size: 1.25rem;
  line-height: 1.7;
  margin: 18px 0 0;
  max-width: 760px;
}

.about-main {
  display: flex;
  flex-direction: column;
  gap: 58px;
}

.intro-section,
.vision-section {
  display: grid;
  gap: 42px;
  grid-template-columns: 220px minmax(0, 1fr);
}

.intro-copy {
  max-width: 820px;
}

.intro-copy h2,
.section-heading h2,
.keyword-panel h2,
.vision-section h2 {
  color: #7d1231;
  font-size: 1.85rem;
  font-weight: 650;
  letter-spacing: 0;
  line-height: 1.25;
  margin: 0;
}

.intro-copy p,
.focus-copy p,
.principle-list p,
.vision-copy p {
  color: #4f5b6a;
  font-size: 1.04rem;
  line-height: 1.8;
  margin: 16px 0 0;
}

.focus-section {
  border-top: 1px solid #eef0f2;
  padding-top: 44px;
}

.section-heading {
  margin-bottom: 20px;
}

.focus-list {
  border-top: 1px solid #e7d9df;
}

.focus-row {
  display: grid;
  gap: 30px;
  grid-template-columns: 90px minmax(0, 1fr);
  padding: 26px 0;
  border-bottom: 1px solid #edf0f2;
}

.focus-index {
  color: #7d1231;
  font-size: 1rem;
  font-variant-numeric: tabular-nums;
  font-weight: 700;
}

.focus-copy h3 {
  color: #273445;
  font-size: 1.25rem;
  font-weight: 650;
  letter-spacing: 0;
  line-height: 1.3;
  margin: 0;
}

.focus-copy p {
  margin-top: 8px;
}

.work-section {
  border-top: 1px solid #eef0f2;
  padding-top: 44px;
}

.work-heading {
  margin-bottom: 28px;
}

.work-layout {
  display: grid;
  gap: 52px;
  grid-template-columns: minmax(0, 1.15fr) minmax(300px, 0.85fr);
  align-items: start;
}

.principle-list {
  list-style: none;
  margin: 0;
  padding: 0;
  border-top: 1px solid #e7d9df;
}

.principle-list li {
  display: grid;
  gap: 22px;
  grid-template-columns: 46px minmax(0, 1fr);
  padding: 20px 0;
  border-bottom: 1px solid #edf0f2;
}

.principle-list span {
  color: #7d1231;
  font-size: 0.95rem;
  font-variant-numeric: tabular-nums;
  font-weight: 700;
}

.principle-list p {
  margin: 0;
}

.keyword-panel {
  background: #faf7f8;
  border-top: 4px solid #7d1231;
  padding: 24px 26px 28px;
}

.keyword-list {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
  margin-top: 22px;
}

.keyword-list span,
.domain-list span {
  border: 1px solid #e2cdd5;
  color: #4f2435;
  font-size: 0.94rem;
  line-height: 1.2;
  padding: 8px 12px;
}

.vision-section {
  align-items: start;
  background: #faf7f8;
  border-top: 4px solid #7d1231;
  padding: 30px;
}

.vision-copy p {
  margin-top: 0;
}

.domain-list {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
  margin-top: 22px;
}

.mobile-backtop {
  right: 100px;
  bottom: 100px;
}

@media (max-width: 768px) {
  .about-container {
    padding: 30px 16px 56px;
  }

  .about-hero {
    margin-bottom: 36px;
    padding-bottom: 28px;
  }

  .page-subtitle {
    font-size: 1.1rem;
  }

  .page-title {
    font-size: 2.4rem;
  }

  .about-main {
    gap: 44px;
  }

  .intro-section,
  .vision-section {
    grid-template-columns: 1fr;
    gap: 18px;
  }

  .work-layout {
    grid-template-columns: 1fr;
    gap: 32px;
  }

  .focus-row {
    grid-template-columns: 56px minmax(0, 1fr);
    gap: 18px;
  }

  .keyword-panel {
    padding: 22px 18px;
  }

  .vision-section {
    padding: 24px 18px;
  }

  .mobile-backtop {
    right: 20px !important;
    bottom: 80px !important;
  }
}

@media (max-width: 480px) {
  .about-container {
    padding: 24px 12px 48px;
  }

  .intro-copy h2,
  .section-heading h2,
  .keyword-panel h2,
  .vision-section h2 {
    font-size: 1.55rem;
  }

  .page-title {
    font-size: 2rem;
  }

  .focus-row {
    grid-template-columns: 1fr;
    gap: 8px;
  }

  .principle-list li {
    grid-template-columns: 36px minmax(0, 1fr);
  }
}
</style>
