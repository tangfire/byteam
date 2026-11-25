<script setup lang="ts">
import { Collection, Link, Promotion, Tickets, VideoCamera } from '@element-plus/icons-vue'
import { computed } from 'vue'

const asset = (name: string) => `/KineticGS/${name}`
const content = {
  title: 'KineticGS: Momentum-driven Coherent 4D Gaussian Splatting\nfor Monocular Dynamic Scene Reconstruction',
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
    items: [
      'Limited observations: monocular videos struggle with complex motion, occlusion, and non-rigid deformation.',
      'Motion inconsistency: existing 4DGS lacks explicit motion modeling, causing trajectory fragmentation and temporal drift.',
      'Noisy guidance: depth and optical-flow cues become unreliable under occlusion and motion blur.',
    ],
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
</script>

<template>
  <el-space direction="vertical" :size="30" style="width: 100%">
    <div style="display: flex; flex-direction: column; align-items: center;">
      <h1 class="section-title">
        <template v-for="(line, index) in lines" :key="index">
          {{ line }}<br v-if="index < lines.length - 1">
        </template>
      </h1>

      <div class="demo-video">
        <video class="demo-video-player" controls preload="metadata" :src="content.video.src">
          Your browser does not support video playback.
        </video>
      </div>

      <el-card class="content-card">
        <div class="project-content">
          <div class="introduction">
            <h3 class="subsection-title first-subtitle">{{ content.introduction.title }}</h3>
            <p class="vision-text">{{ content.introduction.text }}</p>
            <ul class="styled-list">
              <li v-for="item in content.introduction.items" :key="item">{{ item }}</li>
            </ul>
            <img :src="content.introduction.image" class="example2-image" alt="KineticGS teaser">
          </div>

          <div class="motivation">
            <h3 class="subsection-title">{{ content.motivation.title }}</h3>
            <p class="vision-text">{{ content.motivation.text }}</p>
            <ul class="styled-list">
              <li v-for="item in content.motivation.items" :key="item">{{ item }}</li>
            </ul>
            <img :src="content.motivation.image" class="caterogy-image" alt="Comparison between appearance-based and momentum-aware Gaussian modeling">
            <p class="key-insight">{{ content.motivation.insight }}</p>
          </div>

          <div class="Dataset-statistics">
            <h3 class="subsection-title">{{ content.statistics.title }}</h3>
            <p class="vision-text">{{ content.statistics.text }}</p>
            <img :src="content.statistics.image" class="caterogy-image" alt="KineticGS method overview">
          </div>

          <div class="Results">
            <h3 class="subsection-title">{{ content.results.title }}</h3>
            <p class="vision-text">{{ content.results.text }}</p>
            <div class="results-container">
              <template v-for="figure in figures" :key="figure.image">
                <h4 class="results-title">{{ figure.title }}</h4>
                <img :src="figure.image" :alt="figure.title" class="results-image">
              </template>
            </div>
            <div class="tables-container">
              <figure v-for="table in content.results.tables" :key="table.image" class="table-figure">
                <img :src="table.image" :alt="table.title" class="table-image">
                <figcaption>{{ table.title }}</figcaption>
              </figure>
            </div>
          </div>

          <div class="relevant-links">
            <h3 class="subsection-title">Download</h3>
            <div class="relevant-links-container" style="display: flex; gap: 70px; align-items: center;margin-left: 20px;">
              <div v-for="item in links" :key="item.label" style="display: flex; align-items: center;">
                <a :href="item.url" style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank" rel="noopener noreferrer">
                  <el-icon size="25" style="margin-right: 8px; vertical-align: middle;"><component :is="iconForLink(item.type)" /></el-icon>
                  <span style="font-size: 18px; font-weight: 500;">{{ item.label }}</span>
                </a>
              </div>
            </div>
          </div>

          <div class="Citation">
            <h3 class="subsection-title">Citation</h3>
            <p class="vision-text">If you find this work useful in your research, please cite the following paper:</p>
            <pre class="pillar citation-block">{{ content.citation }}</pre>
          </div>
        </div>
      </el-card>
    </div>
  </el-space>
  <div style="height: 50px"></div>
  <el-backtop class="mobile-backtop" :right="100" :bottom="100"/>
</template>

<style scoped>
.kinetic-page { --ink: #172033; --muted: #617087; --line: #dfe6ef; --brand: #7d1231; --brand-dark: #4f0b20; --blue: #1769aa; color: var(--ink); background: #fff; min-height: 100vh; }
.kinetic-shell { width: min(1180px, calc(100% - 40px)); margin: 0 auto; }
.kinetic-hero { background: radial-gradient(circle at 82% 20%, rgba(125,18,49,.45), transparent 35%), linear-gradient(135deg, #280b1a 0%, #111b2f 100%); color: #fff; padding: 66px 0 58px; }
.hero-grid { display: grid; grid-template-columns: minmax(0, 1fr) minmax(360px, .9fr); gap: 60px; align-items: center; }
.eyebrow, .section-kicker, .stage-label { margin: 0 0 12px; font-size: .76rem; font-weight: 800; letter-spacing: .12em; text-transform: uppercase; color: #e5a7b9; }
.hero-copy h1 { margin: 0; max-width: 700px; font-size: clamp(2.25rem, 4.5vw, 4.2rem); line-height: 1.06; letter-spacing: -.045em; }
.title-line { display: block; }
.authors { margin: 24px 0 0; color: #f0d7df; font-size: .98rem; }
.hero-summary { max-width: 600px; color: #d7deea; font-size: 1.12rem; line-height: 1.75; }
.hero-actions { display: flex; flex-wrap: wrap; gap: 12px; margin-top: 28px; }
.kinetic-button { display: inline-flex; align-items: center; justify-content: center; min-height: 44px; padding: 0 18px; border-radius: 999px; font-weight: 700; transition: transform .2s ease, box-shadow .2s ease; }
.kinetic-button:hover { transform: translateY(-2px); box-shadow: 0 10px 22px rgba(0,0,0,.22); }
.kinetic-button--primary { background: #d95f80; color: #fff; }.kinetic-button--light { background: rgba(255,255,255,.1); border: 1px solid rgba(255,255,255,.35); color: #fff; }
.hero-stats { display: grid; grid-template-columns: repeat(3, 1fr); gap: 16px; max-width: 620px; margin-top: 42px; padding-top: 20px; border-top: 1px solid rgba(255,255,255,.22); }
.hero-stats div { display: flex; flex-direction: column; gap: 4px; }.hero-stats strong { font-size: 1.42rem; color: #fff; }.hero-stats span { color: #b8c2d2; font-size: .76rem; }
.hero-figure { margin: 0; padding: 14px; background: rgba(255,255,255,.08); border: 1px solid rgba(255,255,255,.18); border-radius: 18px; box-shadow: 0 24px 70px rgba(0,0,0,.3); }.hero-figure img { width: 100%; border-radius: 10px; background: #fff; }.hero-figure figcaption { margin: 13px 4px 0; color: #c4cedd; font-size: .82rem; line-height: 1.5; }
.kinetic-section { padding: 76px 0; }.kinetic-section--soft { background: #f5f7fa; border-top: 1px solid #edf0f4; border-bottom: 1px solid #edf0f4; }.kinetic-section--dark { background: #131b2b; color: #fff; }
.intro-grid, .download-grid, .citation-grid { display: grid; grid-template-columns: 1.3fr .7fr; gap: 70px; align-items: center; }.section-kicker { color: var(--brand); }.kinetic-section--dark .section-kicker { color: #e5a7b9; }.kinetic-section h2 { margin: 0 0 18px; color: var(--ink); font-size: clamp(1.75rem, 3vw, 2.55rem); letter-spacing: -.025em; }.kinetic-section--dark h2 { color: #fff; }.lead { margin: 0; color: var(--muted); font-size: 1.07rem; line-height: 1.8; }
.feature-list { list-style: none; padding: 0; margin: 25px 0 0; display: grid; gap: 12px; }.feature-list li { position: relative; padding: 14px 16px 14px 42px; border-left: 3px solid var(--brand); background: #fafbfc; color: #344258; line-height: 1.55; }.feature-list li::before { content: '✓'; position: absolute; left: 15px; top: 14px; color: var(--brand); font-weight: 800; }
.principle-card { padding: 30px; min-height: 240px; display: flex; flex-direction: column; justify-content: center; border-radius: 18px; background: linear-gradient(145deg, #f8e8ed, #edf4fb); border: 1px solid #ead6dc; }.principle-number { color: var(--brand); font-size: 4rem; font-weight: 850; line-height: 1; }.principle-card p { margin: 14px 0 8px; font-size: 1.15rem; font-weight: 800; }.principle-card small { color: var(--muted); line-height: 1.55; }
.section-heading { display: flex; justify-content: space-between; gap: 50px; align-items: end; margin-bottom: 28px; }.section-heading > p { max-width: 490px; margin: 0; color: var(--muted); line-height: 1.7; }.method-figure { margin: 0; padding: 20px; background: #fff; border: 1px solid var(--line); border-radius: 14px; box-shadow: 0 10px 30px rgba(29,46,74,.06); }.method-figure img { width: 100%; }.stage-grid { display: grid; grid-template-columns: repeat(3, 1fr); gap: 16px; margin-top: 20px; }.stage-card { padding: 24px; background: #fff; border: 1px solid var(--line); border-radius: 14px; }.stage-icon { color: var(--brand); }.stage-label { margin: 14px 0 8px; color: var(--brand); font-size: .7rem; }.stage-card h3 { margin: 0 0 8px; font-size: 1.1rem; }.stage-card p:last-child { margin: 0; color: var(--muted); line-height: 1.6; font-size: .92rem; }
.metric-grid { display: grid; grid-template-columns: repeat(4, 1fr); gap: 15px; }.metric-card { position: relative; padding: 24px; background: #fff; border: 1px solid var(--line); border-radius: 14px; }.metric-card strong { display: block; color: var(--brand); font-size: 2rem; }.metric-card span { display: block; margin-top: 8px; color: var(--muted); font-size: .83rem; line-height: 1.4; }.metric-card em { position: absolute; right: 13px; top: 13px; padding: 4px 7px; background: #f8e8ed; border-radius: 999px; color: var(--brand); font-size: .65rem; font-style: normal; font-weight: 700; }
.result-grid { display: grid; grid-template-columns: repeat(2, 1fr); gap: 20px; }.result-card { overflow: hidden; padding: 0; text-align: left; background: #fff; border: 1px solid var(--line); border-radius: 14px; cursor: pointer; transition: transform .2s ease, box-shadow .2s ease; }.result-card:hover { transform: translateY(-3px); box-shadow: 0 14px 28px rgba(29,46,74,.12); }.result-card img { width: 100%; aspect-ratio: 16 / 9; object-fit: contain; background: #f7f8fa; }.result-card-copy { display: flex; flex-direction: column; gap: 7px; padding: 16px 18px 18px; }.result-card-copy strong { color: var(--ink); }.result-card-copy small { color: var(--muted); line-height: 1.45; }.view-figure { color: var(--brand); font-size: .78rem; font-weight: 700; }
.resource-links { display: grid; grid-template-columns: 1fr 1fr; gap: 12px; }.resource-link { display: flex; align-items: center; gap: 10px; padding: 15px 16px; border: 1px solid var(--line); border-radius: 10px; color: var(--ink); background: #fff; }.resource-link:hover { border-color: var(--brand); color: var(--brand); }.resource-link small { margin-left: auto; color: var(--muted); }.citation-grid { align-items: start; }.citation-grid > div > p:last-child { color: #bac5d6; }.citation-grid pre { overflow-x: auto; margin: 0; padding: 20px; color: #e7d2d9; background: #0d1421; border: 1px solid #2d3b53; border-radius: 12px; font: .82rem/1.6 'SFMono-Regular', Consolas, monospace; white-space: pre-wrap; }
.dialog-image { display: block; width: 100%; max-height: 75vh; object-fit: contain; }
@media (max-width: 900px) { .hero-grid, .intro-grid, .download-grid, .citation-grid { grid-template-columns: 1fr; gap: 34px; }.hero-figure { max-width: 760px; margin: 0 auto; }.section-heading { display: block; }.section-heading > p { margin-top: 12px; }.metric-grid { grid-template-columns: repeat(2, 1fr); } }
@media (max-width: 640px) { .kinetic-shell { width: min(100% - 26px, 1180px); }.kinetic-hero { padding: 38px 0 42px; }.hero-copy h1 { font-size: 2.15rem; }.hero-stats { gap: 8px; }.hero-stats strong { font-size: 1.05rem; }.hero-stats span { font-size: .68rem; }.kinetic-section { padding: 48px 0; }.stage-grid, .result-grid, .resource-links { grid-template-columns: 1fr; }.metric-card { padding: 18px; }.metric-card strong { font-size: 1.6rem; }.feature-list li { padding-left: 35px; } }
</style>

<style scoped>
.key-insight { margin: 1.2rem 0 0; padding: 1rem 1.2rem; border-left: 4px solid #7d1231; background: #f8f9fa; color: #7d1231; font-weight: 600; }
.section-title, .subsection-title { font-weight: 600; }
.tables-container { margin-top: 2rem; }
.table-figure { margin: 2rem 0; }
.table-image { width: 100%; height: auto; }
.table-figure figcaption { margin-top: .6rem; color: #5d6d7e; font-size: .95rem; text-align: center; line-height: 1.5; }
</style>

<style scoped>
/* KineticGS intentionally reuses the VideoMind project-page visual system. */
@media (max-width: 768px) {
  .section-title { font-size: 0.85rem !important; }
  .subsection-title { font-size: 1rem !important; }
  .results-title { font-size: 0.9rem !important; }
  .first-subtitle { margin-top: -1rem !important; }
  .content-card, .project-content, .introduction, .Dataset-statistics, .relevant-links, .Citation { width: 98%; max-width: 100%; margin: 0 auto; box-sizing: border-box; padding: 0.8rem 0.5rem; }
  .pillar { max-width: 350px; margin: 0 auto; }
  .example2-image, .caterogy-image, .results-image { max-width: 360px; margin: 0 auto; }
  .dataset-images { max-width: 300px; margin: 0 auto; }
  .citation-block { font-size: 0.9rem !important; }
  .mobile-backtop { right: 20px !important; bottom: 80px !important; }
  .relevant-links-container { gap: 30px !important; }
}
.project-content { font-family: 'Segoe UI', system-ui, sans-serif; line-height: 1.7; color: #2c3e50; }
.section-title { color: #7d1231; font-size: 2.2rem; margin: 1rem 0 2rem; font-weight: 600; letter-spacing: -0.5px; text-align: center; }
.content-card { max-width: 1200px; width: 90%; margin: 0 auto; border-radius: 12px; box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08); }
.subsection-title { color: #7d1231; font-size: 1.5rem; margin: 2rem 0 1.5rem; padding-bottom: 0.5rem; border-bottom: 2px solid #eee; }
.styled-list { list-style: none; padding-left: 1.2rem; margin: 0.8rem 0; }
.styled-list li { position: relative; padding-left: 1.2rem; margin-bottom: 0.8rem; }
.styled-list li::before { content: "•"; color: #7d1231; position: absolute; left: 0; font-weight: bold; }
.pillar { background: #f8f9fa; padding: 1.5rem; border-radius: 8px; border-left: 4px solid #7d1231; }
.results-title { margin-top: 0.5rem; margin-bottom: 0.5rem; text-align: center; }
::v-deep(.el-icon svg) { color: #7d1231 !important; }
.citation-block { color: #2c3e50; font-family: 'Fira Mono', 'Consolas', 'Menlo', monospace; font-size: 1rem; box-shadow: 0 2px 8px rgba(44,62,80,0.06); overflow-x: auto; white-space: pre; }
.example2-image, .caterogy-image, .results-image, .dataset-images { width: 100%; height: auto; object-fit: contain; display: block; }
.demo-video-player { width: 100%; max-width: 900px; display: block; margin: 0 auto; border-radius: 8px; background: #000; border: 1px solid #e5e8ed; }
</style>
