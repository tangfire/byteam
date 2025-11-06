<script setup lang="ts">


defineProps<{ msg: string }>()


</script>

<template>

  <el-space direction="vertical" :size="30" style="width: 100%">
    <div style=" align-items: center;">
      <p style="margin-left: 8px; color: #7d1231; font-size: 26px;">Patents</p> <!-- 增加文字大小 -->

      <el-card
          style="max-width: 1000px; margin: 0 auto; display: flex; align-items: center; justify-content: flex-start;">
        <div style="display: flex; align-items: center; justify-content: flex-start;">

          <ul class="patent-list">
            <li>杨宝瑶,黄彦浩,陈涤新，一种基于时空信息聚合的视频特征提取模型训练方法、系统及特征提取方法，2025-09-09,中国, ZL202510359255.5</li>
            <li>杨宝瑶,麻亚利,詹伟德,唐彦超,卢泽坚，一种联邦学习场景下的检测噪声标注的方法及系统，2025-09-05,中国, ZL202510375049.3</li>
            <li>甘孟坤;潘丹;曾安;杨宝瑶;张逸群，一种主动脉弓再缩窄的预后方法及系统，2025-04-01,中国, ZL202210655677.3</li>
            <li>杨宝瑶; 陈俊祥; 黄彦浩; 姚文彬，一种基于多模态大模型的视频-文本检索方法, 2025-06-13,中国, ZL202411271756.X</li>
            <li>杨宝瑶；郑晓琦，一种噪声标注的血管图像分割方法及系统2025-02-11, 中国, ZL 202311792631.7</li>
            <li>Dan Pan, An Zeng, Baoyao Yang, Method for extracting neuroimaging biomarker based on interpretable ensemble 3dcnn,US Patent App. 18/450,338, 2024.</li>
            <li>詹伟德; 杨宝瑶; 方志祥; 李东哲; 胡子逸; 吴恺盈; 一种基于多级智联的数据异构性联邦学习方法及存储介质, 2024-06-21,中国, ZL 202410159303.1</li>
            <li>潘丹; 曾安; 杨宝瑶; 基于可解释集成 3DCNN 的神经影像学生物标志物的提取方法, 2022-11-22,中国, ZL202210987102.5</li>
            <li>曾安; 谢锐伟; 潘丹; 杨宝瑶; 张逸群; 一种心脏图像分割方法及系统, 2022-04-22, 中国,ZL202210030012.3</li>
            <li>潘丹; 罗琳; 曾安; 廖清青; 杨宝瑶; 张逸群; 一种基于多头两级注意力的三维点云语义分割方法, 2022-11-04, 中国, ZL202210709918.8</li>

          </ul>


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
}


.patent-list {
  list-style-type: disc;
  padding-left: 20px;
  line-height: 1.8; /* 增加行间距 */
}

.patent-list li {
  color: #333; /* 正文文字颜色 */
  font-size: 16px; /* 增加字体大小 */
  margin-bottom: 15px; /* 增加每个列表项之间的间距 */
}

.patent-list li::marker {
  color: #7d1231; /* 小圆点颜色 */
  font-size: 15px; /* 可以调整小圆点大小 */
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

.module h3 {
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
