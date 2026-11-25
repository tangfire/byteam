<script setup lang="ts">
import { computed, ref } from 'vue'

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
    caption: 'Demonstration video — KineticGS dynamic scene reconstruction',
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
  citation: `@inproceedings{wang2026kineticgs,
  title={KineticGS: Momentum-driven Coherent 4D Gaussian Splatting for Monocular Dynamic Scene Reconstruction},
  author={Wang, Yunqing and Yang, Baoyao and Liu, Si-Qi and Yin, Chong and Shao, Yihua and Tang, Hao},
  booktitle={European Conference on Computer Vision},
  year={2026}
}`,
}

const lines = computed(() => content.title.split('\n'))
const figures = computed(() => content.results.figures)

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
        <p class="video-caption">{{ content.video.caption }}</p>
      </div>

      <el-card class="content-card">
        <div class="project-content">
          <div class="introduction">
            <h3 class="subsection-title first-subtitle">{{ content.introduction.title }}</h3>
            <p class="vision-text">{{ content.introduction.text }}</p>
            <ul class="styled-list">
              <li v-for="item in content.introduction.items" :key="item">{{ item }}</li>
            </ul>
            <img :src="content.introduction.image" class="example2-image content-image" alt="KineticGS teaser"
                 @click="openPreview(content.introduction.image, 'KineticGS teaser')">
          </div>

          <div class="motivation">
            <h3 class="subsection-title">{{ content.motivation.title }}</h3>
            <p class="vision-text">{{ content.motivation.text }}</p>
            <ul class="styled-list">
              <li v-for="item in content.motivation.items" :key="item">{{ item }}</li>
            </ul>
            <img :src="content.motivation.image" class="caterogy-image content-image"
                 alt="Comparison between appearance-based and momentum-aware Gaussian modeling"
                 @click="openPreview(content.motivation.image, 'Appearance-based vs momentum-aware modeling')">
            <p class="key-insight">{{ content.motivation.insight }}</p>
          </div>

          <div class="Dataset-statistics">
            <h3 class="subsection-title">{{ content.statistics.title }}</h3>
            <p class="vision-text">{{ content.statistics.text }}</p>
            <img :src="content.statistics.image" class="caterogy-image content-image" alt="KineticGS method overview"
                 @click="openPreview(content.statistics.image, 'KineticGS method overview')">
          </div>

          <div class="Results">
            <h3 class="subsection-title">{{ content.results.title }}</h3>
            <p class="vision-text">{{ content.results.text }}</p>
            <div class="results-container">
              <template v-for="figure in figures" :key="figure.image">
                <h4 class="results-title">{{ figure.title }}</h4>
                <img :src="figure.image" :alt="figure.title" class="results-image content-image"
                     @click="openPreview(figure.image, figure.title)">
              </template>
            </div>
            <div class="tables-container">
              <figure v-for="table in content.results.tables" :key="table.image" class="table-figure">
                <img :src="table.image" :alt="table.title" class="table-image content-image"
                     @click="openPreview(table.image, table.title)">
                <figcaption>{{ table.title }}</figcaption>
              </figure>
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
  <el-backtop class="mobile-backtop" :right="100" :bottom="100" />

  <el-dialog v-model="dialogVisible" width="90%" top="5vh" :title="previewAlt" destroy-on-close>
    <img class="dialog-image" :src="previewSrc" :alt="previewAlt">
  </el-dialog>
</template>

<style scoped>
/* KineticGS follows the VideoMind project-page visual system used across the site. */
@media (max-width: 768px) {
  .section-title { font-size: 0.85rem !important; }
  .subsection-title { font-size: 1rem !important; }
  .results-title { font-size: 0.9rem !important; }
  .first-subtitle { margin-top: -1rem !important; }
  .content-card, .project-content, .introduction, .Dataset-statistics, .Citation { width: 98%; max-width: 100%; margin: 0 auto; box-sizing: border-box; padding: 0.8rem 0.5rem; }
  .pillar { max-width: 350px; margin: 0 auto; }
  .example2-image, .caterogy-image, .results-image { max-width: 360px; margin: 0 auto; }
  .citation-block { font-size: 0.9rem !important; }
  .mobile-backtop { right: 20px !important; bottom: 80px !important; }
}

.project-content { font-family: 'Segoe UI', system-ui, sans-serif; line-height: 1.7; color: #2c3e50; }
.section-title { color: #7d1231; font-size: 2.2rem; margin: 1rem 0 2rem; font-weight: 600; letter-spacing: -0.5px; text-align: center; }
.content-card { max-width: 1200px; width: 90%; margin: 0 auto; border-radius: 12px; box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08); }
.subsection-title { color: #7d1231; font-size: 1.5rem; margin: 2rem 0 1.5rem; padding-bottom: 0.5rem; border-bottom: 2px solid #eee; }
.styled-list { list-style: none; padding-left: 1.2rem; margin: 0.8rem 0; }
.styled-list li { position: relative; padding-left: 1.2rem; margin-bottom: 0.8rem; }
.styled-list li::before { content: "•"; color: #7d1231; position: absolute; left: 0; font-weight: bold; }
.pillar { background: #f8f9fa; padding: 1.5rem; border-radius: 8px; border-left: 4px solid #7d1231; }
.results-title { margin-top: 1.4rem; margin-bottom: 0.6rem; text-align: center; color: #172033; font-weight: 600; }
::v-deep(.el-icon svg) { color: #7d1231 !important; }
.citation-block { color: #2c3e50; font-family: 'Fira Mono', 'Consolas', 'Menlo', monospace; font-size: 1rem; box-shadow: 0 2px 8px rgba(44,62,80,0.06); overflow-x: auto; white-space: pre-wrap; word-break: break-word; }
.example2-image, .caterogy-image, .results-image, .dataset-images { width: 100%; height: auto; object-fit: contain; display: block; }
.key-insight { margin: 1.2rem 0 0; padding: 1rem 1.2rem; border-left: 4px solid #7d1231; background: #f8f9fa; color: #7d1231; font-weight: 600; border-radius: 0 6px 6px 0; }
.section-title, .subsection-title { font-weight: 600; }
.tables-container { margin-top: 2rem; }
.table-figure { margin: 2rem 0; }
.table-image { width: 100%; height: auto; }
.table-figure figcaption { margin-top: .6rem; color: #5d6d7e; font-size: .95rem; text-align: center; line-height: 1.5; }

/* Subtle polish within the site's visual language */
.demo-video { width: 100%; display: flex; flex-direction: column; align-items: center; }
.demo-video-player { width: 100%; max-width: 900px; display: block; margin: 0 auto; border-radius: 10px; background: #000; border: 1px solid #e5e8ed; box-shadow: 0 8px 26px rgba(29, 46, 74, .12); }
.video-caption { margin: .7rem auto 0; max-width: 900px; text-align: center; color: #5d6d7e; font-size: .88rem; }
.content-image { border: 1px solid #eceff3; border-radius: 8px; cursor: zoom-in; transition: box-shadow .2s ease; }
.content-image:hover { box-shadow: 0 6px 18px rgba(29, 46, 74, .12); }
.dialog-image { display: block; width: 100%; max-height: 75vh; object-fit: contain; }
</style>
