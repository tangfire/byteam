<script setup lang="ts">

import {Tickets} from "@element-plus/icons-vue";

defineProps<{ msg: string }>()


</script>

<template>

  <el-space direction="vertical" :size="30" style="width: 100%">
   <div>
      <div style=" align-items: center;">

        <p style="margin-left: 8px; color: #7d1231; font-size: 26px;">2024</p> <!-- 增加文字大小 -->
        <p style="margin-left: 8px; color: #7d1231; font-size: 26px;">Journal</p> <!-- 增加文字大小 -->

        <el-card
            style="max-width: 1000px; margin: 0 auto; display: flex; align-items: center; justify-content: flex-start;">
          <div style="display: flex; align-items: center; justify-content: flex-start;">
            <!-- 图片左边 -->
            <img
                src="/publications/002.png"
                style="width: 350px; height: 200px; margin-right: 20px;object-fit: contain;"/>
            <!-- 文字右边 -->
            <div style="display: flex; flex-direction: column; justify-content: flex-start; align-items: flex-start;">
              <!-- 使文字稍微往上移 -->
              <p style="margin-top: -10px;font-size: 18px">Huang J, Yang B, Yin K, et al. Dna-t: Deformable neighborhood attention transformer for irregular medical time series[J]. IEEE Journal of Biomedical and Health Informatics, 2024.</p>

              <!-- 水平排列的超链接 -->
              <div style="display: flex; gap: 70px; align-items: center;">
<!--                <div style="display: flex; align-items: center;">-->
<!--                  <a href="" style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">-->
<!--                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">-->
<!--                      <Promotion />-->
<!--                    </el-icon>-->
<!--                    <span style="font-size: 18px; font-weight: 520;">Code</span>-->
<!--                  </a>-->
<!--                </div>-->

                <!-- 第二个链接（Paper） -->
                <div style="display: flex; align-items: center;">
                  <a href="https://ieeexplore.ieee.org/document/10510586" style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                      <Tickets />
                    </el-icon>
                    <span style="font-size: 18px; font-weight: 520;">Paper</span>
                  </a>
                </div>
              </div>
            </div>
          </div>
        </el-card>





        <p style="margin-left: 8px; color: #7d1231; font-size: 26px;"> Conference</p> <!-- 增加文字大小 -->
        <!-- 每个卡片 -->
        <el-card
            style="max-width: 1000px; margin: 0 auto; display: flex; align-items: center; justify-content: flex-start;">
          <div style="display: flex; align-items: center; justify-content: flex-start;">
            <!-- 图片左边 -->
            <img
                src="/publications/001.png"
                style="width: 350px; height: 200px; margin-right: 20px;object-fit: contain;"/>
            <!-- 文字右边 -->
            <div style="display: flex; flex-direction: column; justify-content: flex-start; align-items: flex-start;">
              <!-- 使文字稍微往上移 -->
              <p style="margin-top: -10px;font-size: 18px">Xie Y, He X, Yang B, et al. CAM-Guided Translation for Unpaired Weakly-Supervised Medical Image Segmentation[C]//2024 IEEE International Conference on Multimedia and Expo (ICME). IEEE, 2024: 1-6.</p>

              <!-- 水平排列的超链接 -->
              <div style="display: flex; gap: 70px; align-items: center;">
                <!-- 第一个链接（Code） -->
<!--                <div style="display: flex; align-items: center;">-->
<!--                  <a href="" style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">-->
<!--                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">-->
<!--                      <Promotion />-->
<!--                    </el-icon>-->
<!--                    <span style="font-size: 18px; font-weight: 520;">Code</span>-->
<!--                  </a>-->
<!--                </div>-->

                <!-- 第二个链接（Paper） -->
                <div style="display: flex; align-items: center;">
                  <a href="https://ieeexplore.ieee.org/abstract/document/10687752" style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                      <Tickets />
                    </el-icon>
                    <span style="font-size: 18px; font-weight: 520;">Paper</span>
                  </a>
                </div>
              </div>
            </div>
          </div>
        </el-card>

        <div style="height: 35px"></div>

        <el-card
            style="max-width: 1000px; margin: 0 auto; display: flex; align-items: center; justify-content: flex-start;">
          <div style="display: flex; align-items: center; justify-content: flex-start;">
            <!-- 图片左边 -->
            <img
                src="/publications/003.png"
                style="width: 350px; height: 200px; margin-right: 20px;object-fit: contain;"/>
            <!-- 文字右边 -->
            <div style="display: flex; flex-direction: column; justify-content: flex-start; align-items: flex-start;">
              <!-- 使文字稍微往上移 -->
              <p style="margin-top: -10px;font-size: 18px">Fan Y, Yang B, Shen M, et al. Domain Dilation for Single Domain Generalization[C]//2024 IEEE International Conference on Image Processing (ICIP). IEEE, 2024: 3931-3937.</p>

              <!-- 水平排列的超链接 -->
              <div style="display: flex; gap: 70px; align-items: center;">
                <!-- 第一个链接（Code） -->
                <!--                <div style="display: flex; align-items: center;">-->
                <!--                  <a href="" style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">-->
                <!--                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">-->
                <!--                      <Promotion />-->
                <!--                    </el-icon>-->
                <!--                    <span style="font-size: 18px; font-weight: 520;">Code</span>-->
                <!--                  </a>-->
                <!--                </div>-->

                <!-- 第二个链接（Paper） -->
                <div style="display: flex; align-items: center;">
                  <a href="https://ieeexplore.ieee.org/abstract/document/10648093" style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                      <Tickets />
                    </el-icon>
                    <span style="font-size: 18px; font-weight: 520;">Paper</span>
                  </a>
                </div>
              </div>
            </div>
          </div>
        </el-card>




      </div>




   </div>
  </el-space>

<!--  底部间隔-->
  <div style="height: 50px"></div>

  <el-backtop :right="100" :bottom="100"  />

</template>

<style scoped>


/* 保证 el-icon 的图标颜色在初始时是正确的 */
::v-deep(.el-icon svg)  {
  color: #7d1231 !important;
}

/* 鼠标移到超链接或者图标上时变色 */
a:hover {
  color: #13393e !important; /* 设置悬停时的颜色 */
}

a:hover ::v-deep(.el-icon svg)  {
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
