<script setup lang="ts">

import {ArrowRightBold} from "@element-plus/icons-vue";

defineProps<{ msg: string }>()


</script>

<template>

  <el-space direction="vertical" :size="30" style="width: 100%">
    <div>
      <div style="display: flex; align-items: center;">
        <el-icon color="#7d1231" style="font-size: 24px;">
          <ArrowRightBold/>
        </el-icon> <!-- 增加图标大小 -->
        <p style="margin-left: 8px; color: #7d1231; font-size: 26px;">Medical Time-Series Representation Learning</p>
        <!-- 增加文字大小 -->
      </div>


      <!-- 每个卡片 -->
      <el-card
          style="max-width: 800px; margin: 0 auto; display: flex; align-items: center; justify-content: flex-start;">
        <div style="display: flex; align-items: center; justify-content: flex-start;">
          <!-- 图片左边 -->
          <img
              src="/currentprojects/001.png"
              style="width: 400px; height: auto; margin-right: 20px;object-fit: contain;"/>
          <!-- 文字右边 -->
          <div>

            <p><strong>Challenges:</strong> Irregular sampling intervals, missing/contaminated data, patient-specific
              variability.</p>
            <p><strong>Solutions:</strong> Multimodal alignment networks for fusing EHRs, imaging, and sensor data;
              contrastive learning frameworks to reconstruct regularized time-series patterns.</p>
          </div>
        </div>
      </el-card>

    </div>

    <div>
      <div style="display: flex; align-items: center;">
        <el-icon color="#7d1231" style="font-size: 24px;">
          <ArrowRightBold/>
        </el-icon>
        <p style="margin-left: 8px;color: #7d1231 ;font-size: 26px;">Federated Learning for Medical Applications</p>
      </div>

      <!-- 其他卡片 -->
      <el-card
          style="max-width: 800px; margin: 0 auto; display: flex; align-items: center; justify-content: flex-start;">
        <div style="display: flex; align-items: center; justify-content: flex-start;">
          <img
              src="/currentprojects/002.png"
              style="width: 400px; height: auto; margin-right: 20px;object-fit: contain;"/>
          <div>

            <p><strong>Challenges:</strong> Data privacy constraints, siloed datasets, heterogeneous model
              architectures.
            </p>
            <p><strong>Innovations:</strong> Privacy-preserving collaborative training protocols (e.g., differential
              autoencoders) with adaptive aggregation mechanisms for cross-institutional knowledge sharing.</p>
          </div>
        </div>
      </el-card>
    </div>


    <div>
      <div style="display: flex; align-items: center;">
        <el-icon color="#7d1231" style="font-size: 24px;">
          <ArrowRightBold/>
        </el-icon>
        <p style="margin-left: 8px;color: #7d1231;font-size: 26px;">Automated Lesion Annotation</p>
      </div>
      <el-card
          style="max-width: 800px; margin: 0 auto; display: flex; align-items: center; justify-content: flex-start;">
        <div style="display: flex; align-items: center; justify-content: flex-start;">
          <img
              src="/currentprojects/003.png"
              style="width: 400px; height: auto; margin-right: 20px;object-fit: contain;"/>
          <div>

            <p><strong>Challenges:</strong> Labor-intensive manual labeling, limited annotated samples, cross-modal
              discrepancies.</p>
            <p><strong>Approaches:</strong> Weakly supervised and prompt-based learning to localize lesions in
              multimodal
              MRI/CT scans while minimizing clinician workload.</p>
          </div>
        </div>
      </el-card>
    </div>


    <div>
      <div style="display: flex; align-items: center;">
        <el-icon color="#7d1231" style="font-size: 24px;">
          <ArrowRightBold/>
        </el-icon>
        <p style="margin-left: 8px;color: #7d1231;font-size: 26px;">Short-Form Video Content Intelligence</p>
      </div>

      <!-- 其他卡片 -->
      <el-card
          style="max-width: 800px; margin: 0 auto; display: flex; align-items: center; justify-content: flex-start;">
        <div style="display: flex; align-items: center; justify-content: flex-start;">
          <img
              src="/currentprojects/004.jpg"
              style="width: 400px; height: auto; margin-right: 20px;object-fit: contain;"/>
          <div>

            <p><strong>Challenges:</strong> Multimodal data fusion (audio, text, video), real-time interpretation,
              copyright/ethics compliance.</p>
            <p><strong>Tools:</strong> Large-scale pretrained models with attention mechanisms for contextual reasoning;
              AI-driven metadata tagging for content authorization.</p>
          </div>
        </div>
      </el-card>
    </div>


    <div>
      <div style="display: flex; align-items: center;">
        <el-icon color="#7d1231" style="font-size: 24px;">
          <ArrowRightBold/>
        </el-icon>
        <p style="margin-left: 8px;color: #7d1231;font-size: 26px;">Industrial IoT & Supply Chain Optimization</p>
      </div>

      <el-card
          style="max-width: 800px; margin: 0 auto; display: flex; align-items: center; justify-content: flex-start;">
        <div style="display: flex; align-items: center; justify-content: flex-start;">
          <img
              src="/currentprojects/005.jpg"
              style="width: 400px; height: auto; margin-right: 20px;object-fit: contain;"/>
          <div>

            <p><strong>Challenges:</strong> Multi-tier communication latency, dynamic resource allocation, cybersecurity
              risks.</p>
            <p><strong>Solutions:</strong> Reinforcement learning for adaptive scheduling; federated edge computing to
              balance efficiency and data sovereignty.</p>
          </div>
        </div>
      </el-card>
    </div>

    <div>
      <div style="display: flex; align-items: center;">
        <el-icon color="#7d1231" style="font-size: 24px;">
          <ArrowRightBold/>
        </el-icon>
        <p style="margin-left: 8px;color: #7d1231;font-size: 26px;">Foundational Algorithm Research</p>
      </div>

      <el-card
          style="max-width: 800px; margin: 0 auto; display: flex; align-items: center; justify-content: flex-start;">
        <div style="display: flex; align-items: center; justify-content: flex-start;">
          <img
              src="/currentprojects/006.png"
              style="width: 400px; height: auto; margin-right: 20px;object-fit: contain;"/>
          <div>

            <p><strong>Focus Areas:</strong> Model explainability in black-box AI, memory-efficient architectures,
              robust
              generalization under limited labels.</p>
            <p><strong>Breakthroughs:</strong> Hybrid frameworks combining causal reasoning, prompt engineering, and
              adversarial robustness testing.</p>
          </div>
        </div>
      </el-card>
    </div>

    <div>
      <div style="display: flex; align-items: center;">
        <el-icon color="#7d1231" style="font-size: 24px;">
          <ArrowRightBold/>
        </el-icon>
        <p style="margin-left: 8px;color: #7d1231;font-size: 26px;">Weak supervised instance segmentation</p>
      </div>

      <el-card
          style="max-width: 800px; margin: 0 auto; display: flex; align-items: center; justify-content: flex-start;">
        <div style="display: flex; align-items: center; justify-content: flex-start;">
          <img
              src="/currentprojects/007.png"
              style="width: 400px; height: auto; margin-right: 20px;object-fit: contain;"/>
          <div>

            <p><strong>Challenges:</strong> Mask merging often occurs for densely packed objects of the same class.</p>
            <p><strong>Solutions:</strong> Contrastive learning to enhance instance discrimination; leveraging motion
              cues
              (in video data) or geometric priors.</p>
          </div>
        </div>
      </el-card>
    </div>


  </el-space>

  <!--  底部间隔-->
  <div style="height: 50px"></div>

  <el-backtop class="mobile-backtop" :right="100" :bottom="100"/>

</template>

<style scoped>

@media (max-width: 768px) {

  .mobile-backtop {
    right: 20px !important;
    bottom: 80px !important;
  }

  /* 调整标题布局 */
  .el-icon + p {
    font-size: 1.2rem !important;
    margin-left: 8px !important;
  }

  .el-card > div {
    display: flex;
    flex-direction: column; /* 移动端垂直排列 */
    align-items: center; /* 居中对齐 */
  }

  .el-card img {
    width: 100% !important; /* 图片宽度100% */
    max-width: 350px; /* 最大宽度保持不变 */
    height: auto !important;
    margin-right: 0 !important;
    margin-bottom: 15px; /* 图片和文字之间添加间距 */
    object-fit: contain;
  }

  /* 文字部分样式 */
  .el-card > div > div:last-child {
    width: 100%; /* 文字区域占满宽度 */
    display: flex;
    flex-direction: column;
    align-items: center; /* 文字和链接居中 */
    text-align: center; /* 文字居中 */
  }

  /* 文字调整 */
  p {
    font-size: 14px !important;
  }
}


::v-deep(.el-icon svg) {
  color: #7d1231 !important;
}

.module p {
  font-size: 20px;
  color: #2980b9;
  margin-bottom: 10px;
}

.module p {
  margin: 8px 0;
  font-size: 16px;
  color: #34495e;
}

.module strong {
  font-weight: bold;
}

.module p:first-child {
  margin-top: 0;
}

.module p:last-child {
  margin-bottom: 0;
}

.research-overview p {
  font-size: 16px;
  color: #34495e;
  margin-bottom: 20px;
  line-height: 1.7;
}
</style>
