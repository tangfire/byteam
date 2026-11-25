<script setup lang="ts">
import { Collection, Link, Promotion, Tickets, VideoCamera } from '@element-plus/icons-vue'
import { computed, ref } from 'vue'

const asset = (name: string) => `/KineticGS/${name}`
const content = {
  title: 'KineticGS: Momentum-driven Coherent 4D Gaussian Splatting\nfor Monocular Dynamic Scene Reconstruction',
  eyebrow: 'ECCV 2026 · Project Page',
  summary: 'KineticGS reconstructs dynamic scenes from monocular videos through a physically grounded hierarchy of Gaussians, momentum-carrying particles, and dynamic objects — modeling how objects move, rather than only how they appear.',
  stats: [
    { value: '+0.91 dB', label: 'Avg. PSNR gain on dynamic objects' },
    { value: '2 Datasets', label: 'DyCheck-iPhone · HyperNeRF' },
    { value: 'ECCV 2026', label: 'European Conf. on Computer Vision' },
  ],
  introduction: {
    title: 'What is KineticGS?',
    text: 'KineticGS reconstructs dynamic scenes from monocular videos through a physically grounded hierarchy of Gaussians, momentum-carrying particles, and dynamic objects. It models how objects move, rather than only how they appear, preserving temporal coherence through occlusion and reappearance.',
    items: [
      'Momentum-guided particle sampling allocates denser representation to high-energy regions and fine-grained non-rigid motion.',
      'Learnable energy-flow activation predicts a smooth temporal window for each particle, enabling continuous motion propagation.',
      'Holistic kinetic synchronization aligns correlated particles without explicit segmentation, reducing drift and fragmentation.',
    ],
    image: asset('figures/teaser_ver2.6.png'),
  },
  video: {
    src: asset('kineticgs-results.mp4'),
  },
  motivation: {
    title: 'Motivation',
    text: 'Monocular videos provide limited observations for dynamic scene reconstruction. Complex motion, occlusion, and non-rigid deformation make depth and optical-flow guidance incomplete and error-prone. Existing 4D Gaussian Splatting methods lack explicit motion modeling, which can cause fragmented trajectories and temporal drift.',
    image: asset('figures/methods_compare4.png'),
    insight: 'Key Insight: Model how objects move, rather than only how they appear, through momentum-aware motion modeling.',
  },
  statistics: {
    title: 'Three-stage momentum hierarchy',
    text: 'KineticGS combines momentum-guided initialization, learnable energy-flow activation, and holistic kinetic synchronization to bridge pixel observations and object-level coherence.',
    image: asset('figures/main_ver9.png'),
  },
  results: {
    title: 'Results',
    text: 'On DyCheck-iPhone and HyperNeRF-interp, KineticGS improves fidelity and temporal coherence, achieving a +0.91 dB average PSNR gain in dynamic-object regions over the second-best method.',
    figures: [
      { title: 'Qualitative comparison on dynamic scenes', image: asset('figures/quali.png') },
      { title: 'Per-frame PSNR comparison', image: asset('figures/per-frame_vis.png') },
      { title: 'Ablation study of the three components', image: asset('figures/ablation_vis.png') },
      { title: 'Comparison of motion modeling', image: asset('figures/methods_compare4.png') },
    ],
    tables: [
      { title: 'Table 1. Quantitative comparison on dynamic object regions of the DyCheck-iPhone dataset.', image: asset('figures/table1.png') },
      { title: 'Table 2. Overall quantitative comparison on the DyCheck-iPhone and HyperNeRF-interp datasets.', image: asset('figures/table2.png') },
    ],
  },
  links: [
    { type: 'code', label: 'Code', url: 'https://github.com/wyunqing/KineticGS.git' },
    { type: 'paper', label: 'Paper PDF', url: asset('main.pdf') },
    { type: 'video', label: 'Video', url: asset('kineticgs-results.mp4') },
  ],
  citation: `@inproceedings{wang2026kineticgs,
  title={KineticGS: Momentum-driven Coherent 4D Gaussian Splatting for Monocular Dynamic Scene Reconstruction},
  author={Wang, Yunqing and Yang, Baoyao and Liu, Si-Qi and Yin, Chong and Shao, Yihua and Tang, Hao},
  booktitle={European Conference on Computer Vision},
  year={2026}
}`,
}

const lines = computed(() => content.title.split('\n'))
const figures = computed(() => content.results.figures)
const links = computed(() => content.links)
const iconForLink = (type: string) => type === 'code' ? Promotion : type === 'paper' ? Tickets : type === 'video' ? VideoCamera : type === 'data' ? Collection : Link

const previewSrc = ref('')
const previewAlt = ref('')
const dialogVisible = ref(false)
const openPreview = (src: string, alt: string) => {
  previewSrc.value = src
  previewAlt.value = alt
  dialogVisible.value = true
}
</script>

<template>
  <div class="kinetic-page">
    <header class="kinetic-hero">
      <div class="kinetic-shell hero-center">
        <p class="eyebrow">{{ content.eyebrow }}</p>
        <h1 class="hero-title">
          <template v-for="(line, index) in lines" :key="index">
            {{ line }}<br v-if="index < lines.length - 1">
          </template>
        </h1>
        <p class="hero-summary">{{ content.summary }}</p>

        <div class="hero-actions">
          <a v-for="item in links" :key="item.label" :href="item.url"
             class="kinetic-button" :class="item.type === 'code' ? 'kinetic-button--primary' : 'kinetic-button--light'"
             target="_blank" rel="noopener noreferrer">
            <el-icon size="18" style="margin-right: 8px;"><component :is="iconForLink(item.type)" /></el-icon>
            {{ item.label }}
          </a>
        </div>

        <figure class="hero-video-frame">
          <video class="hero-video" controls preload="metadata" :src="content.video.src">
            Your browser does not support video playback.
          </video>
          <figcaption>Demonstration video — KineticGS dynamic scene reconstruction</figcaption>
        </figure>

        <div class="hero-stats">
          <div v-for="stat in content.stats" :key="stat.value">
            <strong>{{ stat.value }}</strong>
            <span>{{ stat.label }}</span>
          </div>
        </div>
      </div>
    </header>

    <section class="kinetic-section">
      <div class="kinetic-shell intro-grid">
        <div>
          <p class="section-kicker">Overview</p>
          <h2>{{ content.introduction.title }}</h2>
          <p class="lead">{{ content.introduction.text }}</p>
          <ul class="feature-list">
            <li v-for="item in content.introduction.items" :key="item">{{ item }}</li>
          </ul>
        </div>
        <figure class="framed-figure" @click="openPreview(content.introduction.image, 'KineticGS teaser')">
          <img :src="content.introduction.image" alt="KineticGS teaser">
          <figcaption>Momentum-aware dynamic scene reconstruction from a single monocular video</figcaption>
        </figure>
      </div>
    </section>

    <section class="kinetic-section kinetic-section--soft">
      <div class="kinetic-shell">
        <div class="section-heading">
          <div>
            <p class="section-kicker">Why</p>
            <h2>{{ content.motivation.title }}</h2>
          </div>
          <p>{{ content.motivation.text }}</p>
        </div>
        <figure class="framed-figure" @click="openPreview(content.motivation.image, 'Appearance-based vs momentum-aware modeling')">
          <img :src="content.motivation.image" alt="Comparison between appearance-based and momentum-aware Gaussian modeling">
          <figcaption>Appearance-based modeling loses object identity during occlusion; momentum-driven modeling keeps it</figcaption>
        </figure>
        <p class="key-insight">{{ content.motivation.insight }}</p>
      </div>
    </section>

    <section class="kinetic-section">
      <div class="kinetic-shell">
        <div class="section-heading">
          <div>
            <p class="section-kicker">Method</p>
            <h2>{{ content.statistics.title }}</h2>
          </div>
          <p>{{ content.statistics.text }}</p>
        </div>
        <figure class="framed-figure" @click="openPreview(content.statistics.image, 'KineticGS method overview')">
          <img :src="content.statistics.image" alt="KineticGS method overview">
          <figcaption>Three-stage pipeline: momentum-guided initialization → energy-flow activation → kinetic synchronization</figcaption>
        </figure>
      </div>
    </section>

    <section class="kinetic-section kinetic-section--soft">
      <div class="kinetic-shell">
        <div class="section-heading">
          <div>
            <p class="section-kicker">Experiments</p>
            <h2>{{ content.results.title }}</h2>
          </div>
          <p>{{ content.results.text }}</p>
        </div>

        <div class="result-grid">
          <figure v-for="figure in figures" :key="figure.image" class="result-card"
                  @click="openPreview(figure.image, figure.title)">
            <img :src="figure.image" :alt="figure.title">
            <figcaption>{{ figure.title }}</figcaption>
          </figure>
        </div>

        <div class="table-results">
          <figure v-for="table in content.results.tables" :key="table.image" class="framed-figure"
                  @click="openPreview(table.image, table.title)">
            <img :src="table.image" :alt="table.title">
            <figcaption>{{ table.title }}</figcaption>
          </figure>
        </div>
      </div>
    </section>

    <section class="kinetic-section kinetic-section--dark">
      <div class="kinetic-shell citation-grid">
        <div>
          <p class="section-kicker">Resources</p>
          <h2>Download</h2>
          <div class="resource-links">
            <a v-for="item in links" :key="item.label" :href="item.url" class="resource-link"
               target="_blank" rel="noopener noreferrer">
              <el-icon size="20"><component :is="iconForLink(item.type)" /></el-icon>
              <span>{{ item.label }}</span>
              <small>↗</small>
            </a>
          </div>
        </div>
        <div>
          <p class="citation-lead">If you find this work useful in your research, please cite the following paper:</p>
          <pre class="citation-block">{{ content.citation }}</pre>
        </div>
      </div>
    </section>

    <el-backtop class="mobile-backtop" :right="100" :bottom="100" />

    <el-dialog v-model="dialogVisible" width="90%" top="5vh" :title="previewAlt" destroy-on-close>
      <img class="dialog-image" :src="previewSrc" :alt="previewAlt">
    </el-dialog>
  </div>
</template>

<style scoped>
.kinetic-page {
  --ink: #172033;
  --muted: #617087;
  --line: #dfe6ef;
  --brand: #7d1231;
  --brand-dark: #4f0b20;
  --pink: #d95f80;
  color: var(--ink);
  background: #fff;
  min-height: 100vh;
}
.kinetic-page :deep(h1), .kinetic-page :deep(h2), .kinetic-page :deep(h3) { font-family: 'Poppins', 'Segoe UI', system-ui, sans-serif; }
.kinetic-shell { width: min(1180px, calc(100% - 40px)); margin: 0 auto; }

/* Hero */
.kinetic-hero {
  background: radial-gradient(circle at 82% 18%, rgba(217, 95, 128, .35), transparent 40%),
              radial-gradient(circle at 10% 90%, rgba(23, 105, 170, .25), transparent 45%),
              linear-gradient(135deg, #280b1a 0%, #111b2f 100%);
  color: #fff;
  padding: 64px 0 70px;
}
.hero-center { display: flex; flex-direction: column; align-items: center; text-align: center; }
.eyebrow {
  margin: 0 0 18px;
  font-size: .78rem;
  font-weight: 800;
  letter-spacing: .16em;
  text-transform: uppercase;
  color: #e5a7b9;
}
.hero-title {
  margin: 0;
  max-width: 980px;
  font-size: clamp(1.9rem, 3.6vw, 3.3rem);
  line-height: 1.14;
  letter-spacing: -.03em;
  font-weight: 700;
}
.hero-summary {
  max-width: 780px;
  margin: 22px 0 0;
  color: #d7deea;
  font-size: 1.08rem;
  line-height: 1.75;
}
.hero-actions { display: flex; flex-wrap: wrap; justify-content: center; gap: 14px; margin-top: 30px; }
.kinetic-button {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  min-height: 46px;
  padding: 0 24px;
  border-radius: 999px;
  font-weight: 700;
  font-size: .98rem;
  text-decoration: none;
  transition: transform .2s ease, box-shadow .2s ease, background .2s ease;
}
.kinetic-button:hover { transform: translateY(-2px); box-shadow: 0 10px 22px rgba(0, 0, 0, .28); }
.kinetic-button--primary { background: var(--pink); color: #fff; }
.kinetic-button--light { background: rgba(255, 255, 255, .1); border: 1px solid rgba(255, 255, 255, .35); color: #fff; }
.kinetic-button--light:hover { background: rgba(255, 255, 255, .18); }

.hero-video-frame {
  margin: 44px 0 0;
  padding: 14px;
  width: min(960px, 100%);
  background: rgba(255, 255, 255, .07);
  border: 1px solid rgba(255, 255, 255, .18);
  border-radius: 18px;
  box-shadow: 0 24px 70px rgba(0, 0, 0, .35);
}
.hero-video { display: block; width: 100%; border-radius: 10px; background: #000; }
.hero-video-frame figcaption { margin: 12px 4px 2px; color: #c4cedd; font-size: .84rem; }

.hero-stats {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 18px;
  width: min(760px, 100%);
  margin-top: 44px;
  padding-top: 24px;
  border-top: 1px solid rgba(255, 255, 255, .22);
}
.hero-stats div { display: flex; flex-direction: column; gap: 5px; }
.hero-stats strong { font-size: 1.5rem; color: #fff; }
.hero-stats span { color: #b8c2d2; font-size: .8rem; line-height: 1.4; }

/* Sections */
.kinetic-section { padding: 78px 0; }
.kinetic-section--soft { background: #f5f7fa; border-top: 1px solid #edf0f4; border-bottom: 1px solid #edf0f4; }
.kinetic-section--dark { background: #131b2b; color: #fff; }
.section-kicker {
  margin: 0 0 12px;
  font-size: .76rem;
  font-weight: 800;
  letter-spacing: .14em;
  text-transform: uppercase;
  color: var(--brand);
}
.kinetic-section--dark .section-kicker { color: #e5a7b9; }
.kinetic-section h2 {
  margin: 0 0 18px;
  font-size: clamp(1.7rem, 2.8vw, 2.4rem);
  letter-spacing: -.02em;
  color: var(--ink);
}
.kinetic-section--dark h2 { color: #fff; }
.lead { margin: 0; color: var(--muted); font-size: 1.06rem; line-height: 1.8; }
.section-heading { display: flex; justify-content: space-between; gap: 56px; align-items: flex-end; margin-bottom: 32px; }
.section-heading > p { max-width: 520px; margin: 0; color: var(--muted); line-height: 1.75; }

.intro-grid { display: grid; grid-template-columns: 1.15fr .85fr; gap: 64px; align-items: center; }
.feature-list { list-style: none; padding: 0; margin: 26px 0 0; display: grid; gap: 12px; }
.feature-list li {
  position: relative;
  padding: 14px 16px 14px 44px;
  border-left: 3px solid var(--brand);
  background: #fafbfc;
  border-radius: 0 8px 8px 0;
  color: #344258;
  line-height: 1.55;
}
.feature-list li::before {
  content: '✓';
  position: absolute;
  left: 16px;
  top: 14px;
  color: var(--brand);
  font-weight: 800;
}

/* Figures */
.framed-figure {
  margin: 0;
  padding: 18px;
  background: #fff;
  border: 1px solid var(--line);
  border-radius: 16px;
  box-shadow: 0 12px 34px rgba(29, 46, 74, .08);
  cursor: zoom-in;
  transition: transform .2s ease, box-shadow .2s ease;
}
.framed-figure:hover { transform: translateY(-3px); box-shadow: 0 18px 40px rgba(29, 46, 74, .14); }
.framed-figure img { display: block; width: 100%; height: auto; border-radius: 8px; }
.framed-figure figcaption { margin-top: 12px; color: var(--muted); font-size: .88rem; line-height: 1.55; text-align: center; }
.key-insight {
  margin: 26px 0 0;
  padding: 18px 22px;
  border-left: 4px solid var(--brand);
  background: #fff;
  border-radius: 0 10px 10px 0;
  color: var(--brand);
  font-weight: 600;
  font-size: 1.02rem;
  box-shadow: 0 6px 18px rgba(29, 46, 74, .06);
}

/* Results */
.result-grid { display: grid; grid-template-columns: repeat(2, 1fr); gap: 24px; }
.result-card {
  margin: 0;
  overflow: hidden;
  background: #fff;
  border: 1px solid var(--line);
  border-radius: 16px;
  cursor: zoom-in;
  transition: transform .2s ease, box-shadow .2s ease;
}
.result-card:hover { transform: translateY(-4px); box-shadow: 0 16px 34px rgba(29, 46, 74, .14); }
.result-card img { display: block; width: 100%; aspect-ratio: 16 / 9; object-fit: contain; background: #f7f8fa; }
.result-card figcaption { padding: 15px 18px 17px; color: var(--ink); font-weight: 600; font-size: .95rem; line-height: 1.5; }
.table-results { display: grid; gap: 24px; margin-top: 24px; }

/* Dark resources / citation */
.citation-grid { display: grid; grid-template-columns: .9fr 1.1fr; gap: 64px; align-items: start; }
.resource-links { display: grid; gap: 12px; margin-top: 8px; }
.resource-link {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 15px 18px;
  border: 1px solid rgba(255, 255, 255, .22);
  border-radius: 12px;
  background: rgba(255, 255, 255, .06);
  color: #fff;
  text-decoration: none;
  font-weight: 600;
  transition: border-color .2s ease, background .2s ease, transform .2s ease;
}
.resource-link:hover { border-color: var(--pink); background: rgba(255, 255, 255, .12); transform: translateX(4px); }
.resource-link small { margin-left: auto; color: #9fb0c8; }
.citation-lead { margin: 0 0 14px; color: #bac5d6; line-height: 1.7; }
.citation-block {
  overflow-x: auto;
  margin: 0;
  padding: 20px;
  color: #e7d2d9;
  background: #0d1421;
  border: 1px solid #2d3b53;
  border-radius: 12px;
  font: .84rem/1.65 'SFMono-Regular', Consolas, monospace;
  white-space: pre-wrap;
  word-break: break-word;
}

.dialog-image { display: block; width: 100%; max-height: 78vh; object-fit: contain; }

/* Responsive */
@media (max-width: 900px) {
  .intro-grid, .citation-grid { grid-template-columns: 1fr; gap: 36px; }
  .section-heading { display: block; }
  .section-heading > p { margin-top: 14px; }
  .result-grid { grid-template-columns: 1fr; }
}
@media (max-width: 640px) {
  .kinetic-hero { padding: 42px 0 48px; }
  .hero-title { font-size: 1.75rem; }
  .hero-summary { font-size: .98rem; }
  .hero-stats { gap: 10px; }
  .hero-stats strong { font-size: 1.15rem; }
  .kinetic-section { padding: 50px 0; }
  .hero-video-frame { padding: 8px; border-radius: 12px; }
  .framed-figure { padding: 10px; }
  .mobile-backtop { right: 20px !important; bottom: 80px !important; }
}
</style>
