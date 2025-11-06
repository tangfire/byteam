<script setup lang="ts">
defineProps<{ msg: string }>()
</script>

<template>
  <el-space direction="vertical" :size="30" style="width: 100%">
    <div class="patent-container">
      <h1 class="patent-title">Patents</h1>

      <el-card class="patent-card">
        <div class="patent-content">
          <ul class="patent-list">
            <li>
              <div class="patent-item">
                <div class="patent-authors">杨宝瑶,黄彦浩,陈涤新</div>
                <div class="patent-title-desc">一种基于时空信息聚合的视频特征提取模型训练方法、系统及特征提取方法</div>
                <div class="patent-details">
                  <span class="patent-date">2025-09-09</span>
                  <span class="patent-country">中国</span>
                  <span class="patent-number">ZL202510359255.5</span>
                </div>
              </div>
            </li>
            <li>
              <div class="patent-item">
                <div class="patent-authors">杨宝瑶,麻亚利,詹伟德,唐彦超,卢泽坚</div>
                <div class="patent-title-desc">一种联邦学习场景下的检测噪声标注的方法及系统</div>
                <div class="patent-details">
                  <span class="patent-date">2025-09-05</span>
                  <span class="patent-country">中国</span>
                  <span class="patent-number">ZL202510375049.3</span>
                </div>
              </div>
            </li>
            <li>
              <div class="patent-item">
                <div class="patent-authors">甘孟坤;潘丹;曾安;杨宝瑶;张逸群</div>
                <div class="patent-title-desc">一种主动脉弓再缩窄的预后方法及系统</div>
                <div class="patent-details">
                  <span class="patent-date">2025-04-01</span>
                  <span class="patent-country">中国</span>
                  <span class="patent-number">ZL202210655677.3</span>
                </div>
              </div>
            </li>
            <li>
              <div class="patent-item">
                <div class="patent-authors">杨宝瑶; 陈俊祥; 黄彦浩; 姚文彬</div>
                <div class="patent-title-desc">一种基于多模态大模型的视频-文本检索方法</div>
                <div class="patent-details">
                  <span class="patent-date">2025-06-13</span>
                  <span class="patent-country">中国</span>
                  <span class="patent-number">ZL202411271756.X</span>
                </div>
              </div>
            </li>
            <li>
              <div class="patent-item">
                <div class="patent-authors">杨宝瑶；郑晓琦</div>
                <div class="patent-title-desc">一种噪声标注的血管图像分割方法及系统</div>
                <div class="patent-details">
                  <span class="patent-date">2025-02-11</span>
                  <span class="patent-country">中国</span>
                  <span class="patent-number">ZL 202311792631.7</span>
                </div>
              </div>
            </li>
            <li>
              <div class="patent-item">
                <div class="patent-authors">Dan Pan, An Zeng, Baoyao Yang</div>
                <div class="patent-title-desc">Method for extracting neuroimaging biomarker based on interpretable ensemble 3dcnn</div>
                <div class="patent-details">
                  <span class="patent-date">2024</span>
                  <span class="patent-country">US Patent</span>
                  <span class="patent-number">App. 18/450,338</span>
                </div>
              </div>
            </li>
            <li>
              <div class="patent-item">
                <div class="patent-authors">詹伟德; 杨宝瑶; 方志祥; 李东哲; 胡子逸; 吴恺盈</div>
                <div class="patent-title-desc">一种基于多级智联的数据异构性联邦学习方法及存储介质</div>
                <div class="patent-details">
                  <span class="patent-date">2024-06-21</span>
                  <span class="patent-country">中国</span>
                  <span class="patent-number">ZL 202410159303.1</span>
                </div>
              </div>
            </li>
            <li>
              <div class="patent-item">
                <div class="patent-authors">潘丹; 曾安; 杨宝瑶</div>
                <div class="patent-title-desc">基于可解释集成 3DCNN 的神经影像学生物标志物的提取方法</div>
                <div class="patent-details">
                  <span class="patent-date">2022-11-22</span>
                  <span class="patent-country">中国</span>
                  <span class="patent-number">ZL202210987102.5</span>
                </div>
              </div>
            </li>
            <li>
              <div class="patent-item">
                <div class="patent-authors">曾安; 谢锐伟; 潘丹; 杨宝瑶; 张逸群</div>
                <div class="patent-title-desc">一种心脏图像分割方法及系统</div>
                <div class="patent-details">
                  <span class="patent-date">2022-04-22</span>
                  <span class="patent-country">中国</span>
                  <span class="patent-number">ZL202210030012.3</span>
                </div>
              </div>
            </li>
            <li>
              <div class="patent-item">
                <div class="patent-authors">潘丹; 罗琳; 曾安; 廖清青; 杨宝瑶; 张逸群</div>
                <div class="patent-title-desc">一种基于多头两级注意力的三维点云语义分割方法</div>
                <div class="patent-details">
                  <span class="patent-date">2022-11-04</span>
                  <span class="patent-country">中国</span>
                  <span class="patent-number">ZL202210709918.8</span>
                </div>
              </div>
            </li>
          </ul>
        </div>
      </el-card>
    </div>
  </el-space>

  <!-- 底部间隔 -->
  <div class="bottom-spacing"></div>

  <el-backtop class="mobile-backtop" :right="100" :bottom="100"/>
</template>

<style scoped>
.patent-container {
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px 0;
}

.patent-title {
  color: #7d1231;
  font-size: 32px;
  font-weight: 600;
  margin-bottom: 20px;
  text-align: center;
  position: relative;
  padding-bottom: 10px;
}

.patent-title::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 50%;
  transform: translateX(-50%);
  width: 80px;
  height: 3px;
  background: linear-gradient(90deg, transparent, #7d1231, transparent);
}

.patent-card {
  max-width: 1000px;
  width: 90%;
  margin: 0 auto;
  border-radius: 12px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
  border: none;
  transition: all 0.3s ease;
}

.patent-card:hover {
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.12);
}

.patent-content {
  padding: 10px 5px;
}

.patent-list {
  list-style-type: none;
  padding: 0;
  margin: 0;
}

.patent-list li {
  padding: 20px 0;
  border-bottom: 1px solid #f0f0f0;
  transition: all 0.2s ease;
  position: relative;
  padding-left: 20px;
}

.patent-list li:last-child {
  border-bottom: none;
}

.patent-list li:hover {
  background-color: #f9f9f9;
  border-radius: 8px;
  padding-left: 25px;
}

.patent-list li::before {
  content: '';
  position: absolute;
  left: 0;
  top: 25px;
  width: 8px;
  height: 8px;
  background-color: #7d1231;
  border-radius: 50%;
  transition: all 0.2s ease;
}

.patent-list li:hover::before {
  width: 10px;
  height: 10px;
  background-color: #a51c41;
}

.patent-item {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.patent-authors {
  font-size: 14px;
  color: #666;
  font-weight: 500;
}

.patent-title-desc {
  font-size: 18px;
  color: #333;
  font-weight: 600;
  line-height: 1.4;
}

.patent-details {
  display: flex;
  flex-wrap: wrap;
  gap: 15px;
  margin-top: 5px;
}

.patent-date, .patent-country, .patent-number {
  font-size: 14px;
  padding: 4px 10px;
  border-radius: 4px;
  font-weight: 500;
}

.patent-date {
  background-color: #e8f4fd;
  color: #1a73e8;
}

.patent-country {
  background-color: #e6f4ea;
  color: #34a853;
}

.patent-number {
  background-color: #fef7e0;
  color: #f9ab00;
  font-family: 'Courier New', monospace;
}

.bottom-spacing {
  height: 50px;
}

@media (max-width: 768px) {
  .mobile-backtop {
    right: 20px !important;
    bottom: 80px !important;
  }

  .patent-title {
    font-size: 26px;
  }

  .patent-card {
    width: 95%;
    border-radius: 8px;
  }

  .patent-list li {
    padding: 15px 0;
    padding-left: 15px;
  }

  .patent-list li:hover {
    padding-left: 20px;
  }

  .patent-list li::before {
    top: 20px;
  }

  .patent-title-desc {
    font-size: 16px;
  }

  .patent-authors, .patent-date, .patent-country, .patent-number {
    font-size: 13px;
  }

  .patent-details {
    gap: 8px;
  }
}

/* 保证 el-icon 的图标颜色在初始时是正确的 */
::v-deep(.el-icon svg) {
  color: #7d1231 !important;
}

/* 鼠标移到超链接或者图标上时变色 */
a:hover {
  color: #13393e !important; /* 设置悬停时的颜色 */
}

a:hover ::v-deep(.el-icon svg) {
  color: #13393e !important; /* 设置悬停时图标的颜色 */
}
</style>