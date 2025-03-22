<script setup lang="ts">


defineProps<{ msg: string }>()


</script>

<template>

  <el-space direction="vertical" :size="30" style="width: 100%">
    <div style=" align-items: center;">
      <p style="margin-left: 8px; color: #7d1231; font-size: 26px;">Research Projects</p> <!-- 增加文字大小 -->

      <el-card
          style="max-width: 1000px; margin: 0 auto; display: flex; align-items: center; justify-content: flex-start;">
        <div style="display: flex; align-items: center; justify-content: flex-start;">

          <ul class="patent-list">


            <li>国家自然科学基金,面上项目,62472105, 弹性医学联邦学习与攻击防御关键技术研究,2025-01 至 2028-12,50万元,
              在研, 主持
            </li>
            <li>国家自然科学基金,青年科学基金项目,62102098,基于多模态时间序列表征学习的肝癌早期风险预测算法研究,2022-01
              至 2024-12,30万元, 在研, 主持
            </li>
            <li>国家重点研发计划，工业软件专项，2023YFB3308700，产业聚集区域业务资源服务工业软件平台,2023-12-01 至
              2026-11-31, 3500万（50万），在研，子课题负责人
            </li>
            <li>广东省自然科学基金委,广东省基础与应用基础研究基金面上项目，非侵入式冠状动脉病变的智能定位与评估关键技术研究,2024-01
              至 2026-12,15万,在研，主持
            </li>
            <li>广东省自然科学基金委,广东省基础与应用基础研究基金面上项目，多模态医学大模型设计与医学影像自动标注技术研究,2025-01至2027-12,10万，主持
            </li>
            <li>广东省自然科学基金委,广东省基础与应用基础研究基金区域联合基金项目, 2022A1515140096,
              多源异构热处理数据智能联网及其保护的关键技术研究, 2022-10 至 2025-09,30 万元,在研,校内主持
            </li>
            <li>广州市科技局,广州市基础与应用基础研究项目,202201010266,肝活检图像的多类病变细胞弱监督自动检测算法研究,2022-04
              至 2024-03,5 万元,已结题,主持
            </li>
            <li>广东工业大学,交叉学科培育项目,面向工业物联网的数据保护与协同学习关键技术研究,2022-07
              至 2024-07,5 万元,已结题,主持
            </li>
            <li>广东省科技厅，重点领域研发计划“绿色生物制造”重点专项,
              合成生物基因数据库的建立与查询工具的研发，2024-01 至 2026-12, 1000 万元,在研,参与
            </li>
            <li>广东省自然科学基金委,广东省基础与应用基础研究面上项目,
              2022A1515011592,基于复杂异构特征医学数据挖掘的脓毒症智能预测方法研究,
              2022-01 至 2024-12,10 万元,在研,参与
            </li>


          </ul>


        </div>
      </el-card>
    </div>
  </el-space>


  <!--  底部间隔-->
  <div style="height: 50px"></div>

  <el-backtop :right="100" :bottom="100"/>

</template>

<style scoped>

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
