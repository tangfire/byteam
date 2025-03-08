<script setup lang="ts">


import {Tickets} from "@element-plus/icons-vue";

defineProps<{ msg: string }>()


</script>

<template>
  <div class="block text-center" style="width: 80%;margin: 0 auto">

<!--    <el-carousel height="500px"  motion-blur >-->
<!--      <el-carousel-item>-->
<!--        <div style="display: flex">-->
<!--          <img src="/currentprojects/006.png" style="object-fit: contain;height: 500px;width: auto" alt="">-->
<!--          <div style="background: #f1f2f6;display: block;width: 100%">-->
<!--            Our Research-->


<!--          </div>-->
<!--        </div>-->

<!--      </el-carousel-item>-->
<!--      <el-carousel-item>-->
<!--        <div>-->
<!--          <button>dasd</button>-->
<!--        </div>-->
<!--      </el-carousel-item>-->
<!--      <el-carousel-item>-->

<!--      </el-carousel-item>-->
<!--      <el-carousel-item>-->

<!--      </el-carousel-item>-->

<!--    </el-carousel>-->

    <div class="content">

      <p>Welcome to BY Group</p>

      <p>Our team focuses on medical time series representation learning and multimodal data fusion, privacy-preserving federated learning with heterogeneous collaboration, automated lesion annotation with weakly supervised models, short video multimodal understanding and contrastive learning, intelligent perception and privacy security in industrial IoT, theoretical research on algorithm generalization and interpretability, as well as reinforcement learning and intelligent scheduling in industrial software.</p>
    </div>

    <el-divider />



    <div>



      <el-space direction="vertical" :size="30" style="width: 100%">

        <p style="font-size: 32px;color: #7d1231;">Latest Research</p>
        <div>
          <div style=" align-items: center;">

            <el-card
                style="max-width: 1000px; margin: 0 auto; display: flex; align-items: center; justify-content: flex-start;">
              <div style="display: flex; align-items: center; justify-content: flex-start;">
                <!-- 图片左边 -->
                <img
                    src="/publications/004.png"
                    style="width: 350px; height: 200px; margin-right: 20px;object-fit: contain;"/>
                <!-- 文字右边 -->
                <div style="display: flex; flex-direction: column; justify-content: flex-start; align-items: flex-start;">
                  <!-- 使文字稍微往上移 -->
                  <p style="margin-top: -10px;font-size: 18px">X. Zheng, B. Yang, X. Fang, W. Yao and M. Ye, "Image-assisted Label Connective Completion for Vessel Segmentation with Insufficient Annotations," ICASSP 2025 - 2025 IEEE International Conference on Acoustics, Speech and Signal Processing (ICASSP), Hyderabad, India, 2025, pp. 1-5, doi: 10.1109/ICASSP49660.2025.10888997.</p>

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
                      <a href="https://ieeexplore.ieee.org/document/10888997" style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
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



            <div style="height: 35px"></div>


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

    </div>

    <!--  底部间隔-->
    <div style="height: 50px"></div>

    <el-backtop :right="100" :bottom="100"  />
  </div>
</template>

<style scoped>

::v-deep(.el-icon svg) {
  color: #7d1231 !important;
}

/* 设置整体样式 */
body {
  font-family: 'Arial', sans-serif; /* 设置字体 */
  line-height: 1.6; /* 行间距 */
  background-color: #f4f4f4; /* 背景色 */
  color: #333; /* 默认文字颜色 */
  margin: 0;
  padding: 0;
}

/* 设置div容器的样式 */
.content {
  max-width: 800px;
  margin: 20px auto;
  padding: 20px;
  /* 背景为白色 */
  /*background-color: #fff;*/
  /* 阴影效果 */
  /*box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);*/
  /*border-radius: 8px; /* 圆角 */
}

/* 设置标题样式 */
.content p:first-of-type {
  font-size: 36px;
  color: #7d1231; /* Welcome标题颜色 */
  margin-bottom: 20px;
  font-weight: bold; /* 加粗 */
  text-align: center; /* 居中 */
}

/* 设置正文段落样式 */
.content p:last-of-type {
  font-size: 18px;
  color: #2c2f3b; /* 正文颜色 */
  line-height: 1.8; /* 行间距 */
  text-align: justify; /* 两端对齐 */
}

.demonstration {
  color: var(--el-text-color-secondary);
}

.el-carousel__item h3 {
  color: #475669;
  opacity: 0.75;
  line-height: 200px;
  margin: 0;
  text-align: center;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n + 1) {
  background-color: #d3dce6;
}
</style>
