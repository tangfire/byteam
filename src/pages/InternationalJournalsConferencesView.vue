<script setup lang="ts">

import {Files, Postcard, Promotion, Tickets, VideoPlay} from "@element-plus/icons-vue";

defineProps<{ msg: string }>()


const download_MMS_PPT = () => {
  const pptPath = '/files/MMS.pptx'

  const link = document.createElement('a')
  link.href = pptPath
  link.download = 'Team_Presentation.pptx'

  document.body.appendChild(link)
  link.click()
  document.body.removeChild(link)
}

const download_HBAI_PPT = () => {
  const pptPath = '/files/HBAI.pptx'

  const link = document.createElement('a')
  link.href = pptPath
  link.download = 'Team_Presentation.pptx'

  document.body.appendChild(link)
  link.click()
  document.body.removeChild(link)
}

const download_HBAI_Poster = () => {
  const pptPath = '/files/ldzposter001.pdf'

  const link = document.createElement('a')
  link.href = pptPath
  link.download = 'Team_Poster.pdf'

  document.body.appendChild(link)
  link.click()
  document.body.removeChild(link)
}


const download_WJC001_Poster = () => {
  const pptPath = '/files/wjcposter001.pdf'

  const link = document.createElement('a')
  link.href = pptPath
  link.download = 'Team_Poster.pdf'

  document.body.appendChild(link)
  link.click()
  document.body.removeChild(link)
}


const download_WJC002_Poster = () => {
  const pptPath = '/files/wjcposter002.pptx'

  const link = document.createElement('a')
  link.href = pptPath
  link.download = 'Team_Poster.pptx'

  document.body.appendChild(link)
  link.click()
  document.body.removeChild(link)
}







</script>

<template>

  <el-space direction="vertical" :size="30" style="width: 100%">

<!--    2025-->
    <div>
      <div style=" align-items: center;">

        <p style="margin-left: 8px; color: #7d1231; font-size: 26px;">2025</p> <!-- 增加文字大小 -->
        <!--        <p style="margin-left: 8px; color: #7d1231; font-size: 26px;">Journal</p> &lt;!&ndash; 增加文字大小 &ndash;&gt;-->


        <p style="margin-left: 8px; color: #7d1231; font-size: 26px;"> Conference</p> <!-- 增加文字大小 -->
        <!-- 每个卡片 -->
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
              <p style="margin-top: -10px;font-size: 18px">X. Zheng, B. Yang, X. Fang, W. Yao and M. Ye, "Image-assisted
                Label Connective Completion for Vessel Segmentation with Insufficient Annotations," ICASSP 2025 - 2025
                IEEE International Conference on Acoustics, Speech and Signal Processing (ICASSP), Hyderabad, India,
                2025, pp. 1-5, doi: 10.1109/ICASSP49660.2025.10888997.</p>

              <!-- 水平排列的超链接 -->
              <div style="display: flex; gap: 70px; align-items: center;">
<!--                 第一个链接（Code）-->
                                <div style="display: flex; align-items: center;">
                                  <a href="https://github.com/BaoyaoGroup/LabelCompletion" style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                                      <Promotion />
                                    </el-icon>
                                    <span style="font-size: 18px; font-weight: 500;">Code</span>
                                  </a>
                                </div>

                <!-- 第二个链接（Paper） -->
                <div style="display: flex; align-items: center;">
                  <a href="https://ieeexplore.ieee.org/document/10888997"
                     style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                      <Tickets/>
                    </el-icon>
                    <span style="font-size: 18px; font-weight: 500;">Paper</span>
                  </a>
                </div>

                <div style="display: flex; align-items: center;">
                  <router-link
                      :to="{ name: 'video-player-XiaoqiZheng01' }"
                      style="color: #7d1231; text-decoration: none; display: flex; align-items: center;"
                  >
                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                      <VideoPlay />
                    </el-icon>
                    <span style="font-size: 18px; font-weight: 500;">Video</span>
                  </router-link>
                </div>
              </div>
            </div>
          </div>
        </el-card>


      </div>


    </div>

<!--    2024-->
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
              <p style="margin-top: -10px;font-size: 18px">Huang J, Yang B, Yin K, et al. Dna-t: Deformable neighborhood
                attention transformer for irregular medical time series[J]. IEEE Journal of Biomedical and Health
                Informatics, 2024.</p>

              <!-- 水平排列的超链接 -->
              <div style="display: flex; gap: 70px; align-items: center;">
                                <div style="display: flex; align-items: center;">
                                  <a href="https://github.com/BaoyaoGroup/DNA-T" style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                                      <Promotion />
                                    </el-icon>
                                    <span style="font-size: 18px; font-weight: 500;">Code</span>
                                  </a>
                                </div>

                <!-- 第二个链接（Paper） -->
                <div style="display: flex; align-items: center;">
                  <a href="https://ieeexplore.ieee.org/document/10510586"
                     style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                      <Tickets/>
                    </el-icon>
                    <span style="font-size: 18px; font-weight: 500;">Paper</span>
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
                src="/publications/005.png"
                style="width: 350px; height: 200px; margin-right: 20px;object-fit: contain;"/>
            <!-- 文字右边 -->
            <div style="display: flex; flex-direction: column; justify-content: flex-start; align-items: flex-start;">
              <!-- 使文字稍微往上移 -->
              <p style="margin-top: -10px;font-size: 18px">Baoyao Yang*, PC Yuen, Yiqun Zhang, An Zeng, "Allosteric Feature Collaboration for Model-Heterogeneous Federated Learning," IEEE Transactions on Neural Networks and Learning Systems (TNNLS), 2024.</p>

              <!-- 水平排列的超链接 -->
              <div style="display: flex; gap: 70px; align-items: center;">
                <!--                <div style="display: flex; align-items: center;">-->
                <!--                  <a href="" style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">-->
                <!--                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">-->
                <!--                      <Promotion />-->
                <!--                    </el-icon>-->
                <!--                    <span style="font-size: 18px; font-weight: 500;">Code</span>-->
                <!--                  </a>-->
                <!--                </div>-->

                <!-- 第二个链接（Paper） -->
                <div style="display: flex; align-items: center;">
                  <a href="https://ieeexplore.ieee.org/abstract/document/10373104"
                     style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                      <Tickets/>
                    </el-icon>
                    <span style="font-size: 18px; font-weight: 500;">Paper</span>
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
              <p style="margin-top: -10px;font-size: 18px">Xie Y, He X, Yang B, et al. CAM-Guided Translation for
                Unpaired Weakly-Supervised Medical Image Segmentation[C]//2024 IEEE International Conference on
                Multimedia and Expo (ICME). IEEE, 2024: 1-6.</p>

              <!-- 水平排列的超链接 -->
              <div style="display: flex; gap: 70px; align-items: center;">
                <!-- 第一个链接（Code） -->
                                <div style="display: flex; align-items: center;">
                                  <a href="https://github.com/BaoyaoGroup/CAM_Guided" style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                                      <Promotion />
                                    </el-icon>
                                    <span style="font-size: 18px; font-weight: 500;">Code</span>
                                  </a>
                                </div>

                <!-- 第二个链接（Paper） -->
                <div style="display: flex; align-items: center;">
                  <a href="https://ieeexplore.ieee.org/abstract/document/10687752"
                     style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                      <Tickets/>
                    </el-icon>
                    <span style="font-size: 18px; font-weight: 500;">Paper</span>
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
              <p style="margin-top: -10px;font-size: 18px">Fan Y, Yang B, Shen M, et al. Domain Dilation for Single
                Domain Generalization[C]//2024 IEEE International Conference on Image Processing (ICIP). IEEE, 2024:
                3931-3937.</p>

              <!-- 水平排列的超链接 -->
              <div style="display: flex; gap: 70px; align-items: center;">
                <!-- 第一个链接（Code） -->
                                <div style="display: flex; align-items: center;">
                                  <a href="https://github.com/BaoyaoGroup/Domain_Dilation_for_Single_Domain_Generalization" style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                                      <Promotion />
                                    </el-icon>
                                    <span style="font-size: 18px; font-weight: 500;">Code</span>
                                  </a>
                                </div>

                <!-- 第二个链接（Paper） -->
                <div style="display: flex; align-items: center;">
                  <a href="https://ieeexplore.ieee.org/abstract/document/10648093"
                     style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                      <Tickets/>
                    </el-icon>
                    <span style="font-size: 18px; font-weight: 500;">Paper</span>
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
                src="/publications/006.png"
                style="width: 350px; height: 200px; margin-right: 20px;object-fit: contain;"/>
            <!-- 文字右边 -->
            <div style="display: flex; flex-direction: column; justify-content: flex-start; align-items: flex-start;">
              <!-- 使文字稍微往上移 -->


              <p style="margin-top: -10px;font-size: 18px">Xiaochen He, Baoyao Yang*, Fei Lyu, "MMS: Morphology-mixup Stylized Data Generation for Single Domain Generalization in Medical Image Segmentation," IEEE International Conference on Acoustics, Speech, and Signal Processing (ICASSP), 2024.</p>

              <!-- 水平排列的超链接 -->
              <div style="display: flex; gap: 70px; align-items: center;">
                <!-- 第一个链接（Code） -->
                                <div style="display: flex; align-items: center;">
                                  <a href="https://github.com/BaoyaoGroup/MMS" style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                                      <Promotion />
                                    </el-icon>
                                    <span style="font-size: 18px; font-weight: 500;">Code</span>
                                  </a>
                                </div>

                <!-- 第二个链接（Paper） -->
                <div style="display: flex; align-items: center;">
                  <a href="https://ieeexplore.ieee.org/abstract/document/10448305"
                     style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                      <Tickets/>
                    </el-icon>
                    <span style="font-size: 18px; font-weight: 500;">Paper</span>
                  </a>
                </div>
                <div>
                  <el-button class="custom-button"
                    type="primary"
                    @click="download_MMS_PPT"
                    style="background-color: white; border-color: white;"
                >

                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;"><Files /></el-icon>

                  <span style="font-size: 18px; font-weight: 500; color: #7d1231;">PPT</span>
                </el-button></div>
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
                src="/publications/007.png"
                style="width: 350px; height: 200px; margin-right: 20px;object-fit: contain;"/>
            <!-- 文字右边 -->
            <div style="display: flex; flex-direction: column; justify-content: flex-start; align-items: flex-start;">
              <!-- 使文字稍微往上移 -->


              <p style="margin-top: -10px;font-size: 18px">Dongzhe Li, Baoyao Yang*, Weide Zhan, Xiaochen He, "Multi-category Graph Reasoning for Multi-modal Brain Tumor Segmentation," Medical Image Computing and Computer Assisted Intervention (MICCAI), 2024.</p>

              <!-- 水平排列的超链接 -->
              <div style="display: flex; gap: 70px; align-items: center;">
                <!-- 第一个链接（Code） -->
                                <div style="display: flex; align-items: center;">
                                  <a href="https://github.com/BaoyaoGroup/Graph-Co-reasoning" style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                                      <Promotion />
                                    </el-icon>
                                    <span style="font-size: 18px; font-weight: 500;">Code</span>
                                  </a>
                                </div>

                <!-- 第二个链接（Paper） -->
                <div style="display: flex; align-items: center;">
                  <a href="https://link.springer.com/chapter/10.1007/978-3-031-72111-3_42"
                     style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                      <Tickets/>
                    </el-icon>
                    <span style="font-size: 18px; font-weight: 500;">Paper</span>
                  </a>
                </div>
                <div>
                  <el-button
                      class="custom-button"
                      type="primary"
                      @click="download_HBAI_PPT"
                      style="background-color: white; border-color: white;"
                  >

                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;"><Files /></el-icon>

                    <span style="font-size: 18px; font-weight: 500; color: #7d1231;">PPT</span>
                  </el-button></div>
                <div>
                  <el-button
                      class="custom-button"
                      type="primary"
                      @click="download_HBAI_Poster"
                      style="background-color: white; border-color: white;"
                  >

                    <el-icon  size="25" style="margin-right: 8px; vertical-align: middle;"><Postcard /></el-icon>

                    <span style="font-size: 18px; font-weight: 500; color: #7d1231;">Poster</span>
                  </el-button></div>
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
                src="/publications/008.png"
                style="width: 350px; height: 200px; margin-right: 20px;object-fit: contain;"/>
            <!-- 文字右边 -->
            <div style="display: flex; flex-direction: column; justify-content: flex-start; align-items: flex-start;">
              <!-- 使文字稍微往上移 -->


              <p style="margin-top: -10px;font-size: 18px">Wang J, Deng Z, Lin T, et al. Beyond Direct Relationships: Exploring Multi-Order Label Pair Dependencies for Knowledge Distillation[C]//Proceedings of the 32nd ACM International Conference on Multimedia. 2024: 8527-8535.</p>

              <!-- 水平排列的超链接 -->
              <div style="display: flex; gap: 70px; align-items: center;">
                <!-- 第一个链接（Code） -->
                                <div style="display: flex; align-items: center;">
                                  <a href="https://github.com/BaoyaoGroup/Beyond-Direct-Relationships--Exploring-Multi-Order-Distillation-main" style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                                      <Promotion />
                                    </el-icon>
                                    <span style="font-size: 18px; font-weight: 500;">Code</span>
                                  </a>
                                </div>

                <!-- 第二个链接（Paper） -->
                <div style="display: flex; align-items: center;">
                  <a href="https://dl.acm.org/doi/abs/10.1145/3664647.3681029"
                     style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                      <Tickets/>
                    </el-icon>
                    <span style="font-size: 18px; font-weight: 500;">Paper</span>
                  </a>
                </div>
<!--                <div>-->
<!--                  <el-button-->
<!--                      class="custom-button"-->
<!--                      type="primary"-->
<!--                      @click="download_HBAI_PPT"-->
<!--                      style="background-color: white; border-color: white;"-->
<!--                  >-->

<!--                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;"><Files /></el-icon>-->

<!--                    <span style="font-size: 18px; font-weight: 500; color: #7d1231;">PPT</span>-->
<!--                  </el-button></div>-->
                <div>
                  <el-button
                      class="custom-button"
                      type="primary"
                      @click="download_WJC001_Poster"
                      style="background-color: white; border-color: white;"
                  >

                    <el-icon  size="25" style="margin-right: 8px; vertical-align: middle;"><Postcard /></el-icon>

                    <span style="font-size: 18px; font-weight: 500; color: #7d1231;">Poster</span>
                  </el-button></div>
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
                src="/publications/009.png"
                style="width: 350px; height: 200px; margin-right: 20px;object-fit: contain;"/>
            <!-- 文字右边 -->
            <div style="display: flex; flex-direction: column; justify-content: flex-start; align-items: flex-start;">
              <!-- 使文字稍微往上移 -->


              <p style="margin-top: -10px;font-size: 18px">Wang J, Deng Z, Lin T, et al. A Novel Prompt Tuning for Graph Transformers: Tailoring Prompts to Graph Topologies[C]//Proceedings of the 30th ACM SIGKDD Conference on Knowledge Discovery and Data Mining. 2024: 3116-3127.</p>

              <!-- 水平排列的超链接 -->
              <div style="display: flex; gap: 70px; align-items: center;">
                <!-- 第一个链接（Code） -->
                                <div style="display: flex; align-items: center;">
                                  <a href="https://github.com/BaoyaoGroup/TGPT_main-main" style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                                      <Promotion />
                                    </el-icon>
                                    <span style="font-size: 18px; font-weight: 500;">Code</span>
                                  </a>
                                </div>

                <!-- 第二个链接（Paper） -->
                <div style="display: flex; align-items: center;">
                  <a href="https://dl.acm.org/doi/10.1145/3637528.3671804"
                     style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                      <Tickets/>
                    </el-icon>
                    <span style="font-size: 18px; font-weight: 500;">Paper</span>
                  </a>
                </div>
                <!--                <div>-->
                <!--                  <el-button-->
                <!--                      class="custom-button"-->
                <!--                      type="primary"-->
                <!--                      @click="download_HBAI_PPT"-->
                <!--                      style="background-color: white; border-color: white;"-->
                <!--                  >-->

                <!--                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;"><Files /></el-icon>-->

                <!--                    <span style="font-size: 18px; font-weight: 500; color: #7d1231;">PPT</span>-->
                <!--                  </el-button></div>-->
                <div>
                  <el-button
                      class="custom-button"
                      type="primary"
                      @click="download_WJC002_Poster"
                      style="background-color: white; border-color: white;"
                  >

                    <el-icon  size="25" style="margin-right: 8px; vertical-align: middle;"><Postcard /></el-icon>

                    <span style="font-size: 18px; font-weight: 500; color: #7d1231;">Poster</span>
                  </el-button></div>
              </div>
            </div>
          </div>
        </el-card>



      </div>


    </div>

<!--    2023-->
    <div>
      <div style=" align-items: center;">

        <p style="margin-left: 8px; color: #7d1231; font-size: 26px;">2023</p> <!-- 增加文字大小 -->
        <!--        <p style="margin-left: 8px; color: #7d1231; font-size: 26px;">Journal</p> &lt;!&ndash; 增加文字大小 &ndash;&gt;-->


        <p style="margin-left: 8px; color: #7d1231; font-size: 26px;"> Conference</p> <!-- 增加文字大小 -->
        <!-- 每个卡片 -->


        <el-card
            style="max-width: 1000px; margin: 0 auto; display: flex; align-items: center; justify-content: flex-start;">
          <div style="display: flex; align-items: center; justify-content: flex-start;">
            <!-- 图片左边 -->
            <img
                src="/publications/010.png"
                style="width: 350px; height: 200px; margin-right: 20px;object-fit: contain;"/>
            <!-- 文字右边 -->
            <div style="display: flex; flex-direction: column; justify-content: flex-start; align-items: flex-start;">
              <!-- 使文字稍微往上移 -->
              <p style="margin-top: -10px;font-size: 18px">Yi Li, Baoyao Yang, Dan Pan, An  Zeng, Yang Yang, "Early diagnosis of Alzheimer's disease based on multimodal hypergraph attention network", International Conference on Multimedia and Expo (ICME), 2023.</p>

              <!-- 水平排列的超链接 -->
              <div style="display: flex; gap: 70px; align-items: center;">
                <!--                <div style="display: flex; align-items: center;">-->
                <!--                  <a href="" style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">-->
                <!--                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">-->
                <!--                      <Promotion />-->
                <!--                    </el-icon>-->
                <!--                    <span style="font-size: 18px; font-weight: 500;">Code</span>-->
                <!--                  </a>-->
                <!--                </div>-->

                <!-- 第二个链接（Paper） -->
                <div style="display: flex; align-items: center;">
                  <a href="https://ieeexplore.ieee.org/abstract/document/10219893"
                     style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                      <Tickets/>
                    </el-icon>
                    <span style="font-size: 18px; font-weight: 500;">Paper</span>
                  </a>
                </div>
              </div>
            </div>
          </div>
        </el-card>


      </div>


    </div>

<!--    2022-->
    <div>
      <div style=" align-items: center;">

        <p style="margin-left: 8px; color: #7d1231; font-size: 26px;">2022</p> <!-- 增加文字大小 -->
        <!--        <p style="margin-left: 8px; color: #7d1231; font-size: 26px;">Journal</p> &lt;!&ndash; 增加文字大小 &ndash;&gt;-->


        <p style="margin-left: 8px; color: #7d1231; font-size: 26px;"> Journal</p> <!-- 增加文字大小 -->


        <el-card
            style="max-width: 1000px; margin: 0 auto; display: flex; align-items: center; justify-content: flex-start;">
          <div style="display: flex; align-items: center; justify-content: flex-start;">
            <!-- 图片左边 -->
            <img
                src="/publications/011.png"
                style="width: 350px; height: 200px; margin-right: 20px;object-fit: contain;"/>
            <!-- 文字右边 -->
            <div style="display: flex; flex-direction: column; justify-content: flex-start; align-items: flex-start;">
              <!-- 使文字稍微往上移 -->
              <p style="margin-top: -10px;font-size: 18px">Baoyao Yang, Hao-wei Yeh, Tatsuya Harada, and Pong C. Yuen*, "Model-induced Generalization Error Bound for Information-theoretic Representation Learning in Source-data-free Unsupervised Domain Adaptation," IEEE Transactions on Image Processing (TIP), 2022.</p>

              <!-- 水平排列的超链接 -->
              <div style="display: flex; gap: 70px; align-items: center;">
                <!--                <div style="display: flex; align-items: center;">-->
                <!--                  <a href="" style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">-->
                <!--                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">-->
                <!--                      <Promotion />-->
                <!--                    </el-icon>-->
                <!--                    <span style="font-size: 18px; font-weight: 500;">Code</span>-->
                <!--                  </a>-->
                <!--                </div>-->

                <!-- 第二个链接（Paper） -->
                <div style="display: flex; align-items: center;">
                  <a href="https://ieeexplore.ieee.org/abstract/document/9640468"
                     style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                      <Tickets/>
                    </el-icon>
                    <span style="font-size: 18px; font-weight: 500;">Paper</span>
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
                src="/publications/012.png"
                style="width: 350px; height: 200px; margin-right: 20px;object-fit: contain;"/>
            <!-- 文字右边 -->
            <div style="display: flex; flex-direction: column; justify-content: flex-start; align-items: flex-start;">
              <!-- 使文字稍微往上移 -->
              <p style="margin-top: -10px;font-size: 18px">Baoyao Yang and Pong C. Yuen*, "Revealing Task-relevant Model Memorization for Source-Protected Unsupervised Domain Adaptation," IEEE Transactions on Information Forensics and Security (TIFS), 2022.</p>

              <!-- 水平排列的超链接 -->
              <div style="display: flex; gap: 70px; align-items: center;">
                <!--                <div style="display: flex; align-items: center;">-->
                <!--                  <a href="" style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">-->
                <!--                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">-->
                <!--                      <Promotion />-->
                <!--                    </el-icon>-->
                <!--                    <span style="font-size: 18px; font-weight: 500;">Code</span>-->
                <!--                  </a>-->
                <!--                </div>-->

                <!-- 第二个链接（Paper） -->
                <div style="display: flex; align-items: center;">
                  <a href="https://ieeexplore.ieee.org/abstract/document/9705526"
                     style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                      <Tickets/>
                    </el-icon>
                    <span style="font-size: 18px; font-weight: 500;">Paper</span>
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
                src="/publications/013.png"
                style="width: 350px; height: 200px; margin-right: 20px;object-fit: contain;"/>
            <!-- 文字右边 -->
            <div style="display: flex; flex-direction: column; justify-content: flex-start; align-items: flex-start;">
              <!-- 使文字稍微往上移 -->
              <p style="margin-top: -10px;font-size: 18px">Baoyao Yang, Mang Ye, Qingxiong Tan and Pong C. Yuen*, "Cross-domain Missingness-aware Time Series Adaptation with Similarity Distillation in Medical Applications," IEEE Transactions on Cybernetics (TCYB), 2022.</p>

              <!-- 水平排列的超链接 -->
              <div style="display: flex; gap: 70px; align-items: center;">
                <!--                <div style="display: flex; align-items: center;">-->
                <!--                  <a href="" style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">-->
                <!--                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">-->
                <!--                      <Promotion />-->
                <!--                    </el-icon>-->
                <!--                    <span style="font-size: 18px; font-weight: 500;">Code</span>-->
                <!--                  </a>-->
                <!--                </div>-->

                <!-- 第二个链接（Paper） -->
                <div style="display: flex; align-items: center;">
                  <a href="https://ieeexplore.ieee.org/abstract/document/9167415"
                     style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                      <Tickets/>
                    </el-icon>
                    <span style="font-size: 18px; font-weight: 500;">Paper</span>
                  </a>
                </div>
              </div>
            </div>
          </div>
        </el-card>

      </div>


    </div>


<!--    2021-->
    <div>
      <div style=" align-items: center;">

        <p style="margin-left: 8px; color: #7d1231; font-size: 26px;">2021</p> <!-- 增加文字大小 -->
        <!--        <p style="margin-left: 8px; color: #7d1231; font-size: 26px;">Journal</p> &lt;!&ndash; 增加文字大小 &ndash;&gt;-->

        <p style="margin-left: 8px; color: #7d1231; font-size: 26px;">Journal</p> <!-- 增加文字大小 -->





        <el-card
            style="max-width: 1000px; margin: 0 auto; display: flex; align-items: center; justify-content: flex-start;">
          <div style="display: flex; align-items: center; justify-content: flex-start;">
            <!-- 图片左边 -->
            <img
                src="/publications/015.png"
                style="width: 350px; height: 200px; margin-right: 20px;object-fit: contain;"/>
            <!-- 文字右边 -->
            <div style="display: flex; flex-direction: column; justify-content: flex-start; align-items: flex-start;">
              <!-- 使文字稍微往上移 -->
              <p style="margin-top: -10px;font-size: 18px">Baoyao Yang and Pong C. Yuen*, "Learning Adaptive Geometry for Unsupervised Domain Adaptation," Pattern Recognition (PR), 2021.</p>

              <!-- 水平排列的超链接 -->
              <div style="display: flex; gap: 70px; align-items: center;">
                <!--                <div style="display: flex; align-items: center;">-->
                <!--                  <a href="" style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">-->
                <!--                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">-->
                <!--                      <Promotion />-->
                <!--                    </el-icon>-->
                <!--                    <span style="font-size: 18px; font-weight: 500;">Code</span>-->
                <!--                  </a>-->
                <!--                </div>-->

                <!-- 第二个链接（Paper） -->
                <div style="display: flex; align-items: center;">
                  <a href="https://www.sciencedirect.com/science/article/abs/pii/S0031320320304416"
                     style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                      <Tickets/>
                    </el-icon>
                    <span style="font-size: 18px; font-weight: 500;">Paper</span>
                  </a>
                </div>
              </div>
            </div>
          </div>
        </el-card>


        <p style="margin-left: 8px; color: #7d1231; font-size: 26px;"> Conference</p> <!-- 增加文字大小 -->


        <el-card
            style="max-width: 1000px; margin: 0 auto; display: flex; align-items: center; justify-content: flex-start;">
          <div style="display: flex; align-items: center; justify-content: flex-start;">
            <!-- 图片左边 -->
            <img
                src="/publications/016.png"
                style="width: 350px; height: 200px; margin-right: 20px;object-fit: contain;"/>
            <!-- 文字右边 -->
            <div style="display: flex; flex-direction: column; justify-content: flex-start; align-items: flex-start;">
              <!-- 使文字稍微往上移 -->
              <p style="margin-top: -10px;font-size: 18px">Fei Lyu, Baoyao Yang, Andy J. Ma and Pong C. Yuen*, "A Segmentation-Assisted Model for Universal Lesion Detection with Partial Labels," International Conference on Medical Image Computing & Computer Assisted Intervention (MICCAI), 2021.</p>

              <!-- 水平排列的超链接 -->
              <div style="display: flex; gap: 70px; align-items: center;">
                <!--                <div style="display: flex; align-items: center;">-->
                <!--                  <a href="" style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">-->
                <!--                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">-->
                <!--                      <Promotion />-->
                <!--                    </el-icon>-->
                <!--                    <span style="font-size: 18px; font-weight: 500;">Code</span>-->
                <!--                  </a>-->
                <!--                </div>-->

                <!-- 第二个链接（Paper） -->
                <div style="display: flex; align-items: center;">
                  <a href="https://link.springer.com/chapter/10.1007/978-3-030-87240-3_12"
                     style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                      <Tickets/>
                    </el-icon>
                    <span style="font-size: 18px; font-weight: 500;">Paper</span>
                  </a>
                </div>
              </div>
            </div>
          </div>
        </el-card>




      </div>


    </div>

<!--    2019-->
    <div>
      <div style=" align-items: center;">

        <p style="margin-left: 8px; color: #7d1231; font-size: 26px;">2019</p> <!-- 增加文字大小 -->
        <!--        <p style="margin-left: 8px; color: #7d1231; font-size: 26px;">Journal</p> &lt;!&ndash; 增加文字大小 &ndash;&gt;-->

        <p style="margin-left: 8px; color: #7d1231; font-size: 26px;"> Journal</p> <!-- 增加文字大小 -->



        <el-card
            style="max-width: 1000px; margin: 0 auto; display: flex; align-items: center; justify-content: flex-start;">
          <div style="display: flex; align-items: center; justify-content: flex-start;">
            <!-- 图片左边 -->
            <img
                src="/publications/017.png"
                style="width: 350px; height: 200px; margin-right: 20px;object-fit: contain;"/>
            <!-- 文字右边 -->
            <div style="display: flex; flex-direction: column; justify-content: flex-start; align-items: flex-start;">
              <!-- 使文字稍微往上移 -->
              <p style="margin-top: -10px;font-size: 18px">Baoyao Yang, Andy J. Ma and Pong C. Yuen*, "Body Parts Synthesis for Cross-Quality Pose Estimation." IEEE Transactions on Circuits and Systems for Video Technology (TCSVT), 2019. </p>

              <!-- 水平排列的超链接 -->
              <div style="display: flex; gap: 70px; align-items: center;">
                <!--                <div style="display: flex; align-items: center;">-->
                <!--                  <a href="" style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">-->
                <!--                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">-->
                <!--                      <Promotion />-->
                <!--                    </el-icon>-->
                <!--                    <span style="font-size: 18px; font-weight: 500;">Code</span>-->
                <!--                  </a>-->
                <!--                </div>-->

                <!-- 第二个链接（Paper） -->
                <div style="display: flex; align-items: center;">
                  <a href="https://ieeexplore.ieee.org/document/8245864"
                     style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                      <Tickets/>
                    </el-icon>
                    <span style="font-size: 18px; font-weight: 500;">Paper</span>
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
                src="/publications/018.png"
                style="width: 350px; height: 200px; margin-right: 20px;object-fit: contain;"/>
            <!-- 文字右边 -->
            <div style="display: flex; flex-direction: column; justify-content: flex-start; align-items: flex-start;">
              <!-- 使文字稍微往上移 -->
              <p style="margin-top: -10px;font-size: 18px">Baoyao Yang and Pong C. Yuen*, "Cross-Domain Visual Representations via Unsupervised Graph Alignment," the 33rd AAAI Conference on Artificial Intelligence (AAAI), 2019.</p>

              <!-- 水平排列的超链接 -->
              <div style="display: flex; gap: 70px; align-items: center;">
                <!--                <div style="display: flex; align-items: center;">-->
                <!--                  <a href="" style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">-->
                <!--                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">-->
                <!--                      <Promotion />-->
                <!--                    </el-icon>-->
                <!--                    <span style="font-size: 18px; font-weight: 500;">Code</span>-->
                <!--                  </a>-->
                <!--                </div>-->

                <!-- 第二个链接（Paper） -->
                <div style="display: flex; align-items: center;">
                  <a href="https://ojs.aaai.org/index.php/AAAI/article/view/4504"
                     style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                      <Tickets/>
                    </el-icon>
                    <span style="font-size: 18px; font-weight: 500;">Paper</span>
                  </a>
                </div>
              </div>
            </div>
          </div>
        </el-card>


      </div>


    </div>

<!--2018-->
    <div>
      <div style=" align-items: center;">

        <p style="margin-left: 8px; color: #7d1231; font-size: 26px;">2018</p> <!-- 增加文字大小 -->
        <!--        <p style="margin-left: 8px; color: #7d1231; font-size: 26px;">Journal</p> &lt;!&ndash; 增加文字大小 &ndash;&gt;-->

        <p style="margin-left: 8px; color: #7d1231; font-size: 26px;"> Journal</p> <!-- 增加文字大小 -->




        <el-card
            style="max-width: 1000px; margin: 0 auto; display: flex; align-items: center; justify-content: flex-start;">
          <div style="display: flex; align-items: center; justify-content: flex-start;">
            <!-- 图片左边 -->
            <img
                src="/publications/019.png"
                style="width: 350px; height: 200px; margin-right: 20px;object-fit: contain;"/>
            <!-- 文字右边 -->
            <div style="display: flex; flex-direction: column; justify-content: flex-start; align-items: flex-start;">
              <!-- 使文字稍微往上移 -->
              <p style="margin-top: -10px;font-size: 18px">Baoyao Yang, Andy J. Ma and Pong C. Yuen*, "Learning Domain-Shared Group-Sparse Representation for Unsupervised Domain Adaptation," Pattern Recognition (PR), 2018. </p>

              <!-- 水平排列的超链接 -->
              <div style="display: flex; gap: 70px; align-items: center;">
                <!--                <div style="display: flex; align-items: center;">-->
                <!--                  <a href="" style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">-->
                <!--                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">-->
                <!--                      <Promotion />-->
                <!--                    </el-icon>-->
                <!--                    <span style="font-size: 18px; font-weight: 500;">Code</span>-->
                <!--                  </a>-->
                <!--                </div>-->

                <!-- 第二个链接（Paper） -->
                <div style="display: flex; align-items: center;">
                  <a href="https://www.sciencedirect.com/science/article/pii/S0031320318301614"
                     style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                      <Tickets/>
                    </el-icon>
                    <span style="font-size: 18px; font-weight: 500;">Paper</span>
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
                src="/publications/020.png"
                style="width: 350px; height: 200px; margin-right: 20px;object-fit: contain;"/>
            <!-- 文字右边 -->
            <div style="display: flex; flex-direction: column; justify-content: flex-start; align-items: flex-start;">
              <!-- 使文字稍微往上移 -->
              <p style="margin-top: -10px;font-size: 18px">Baoyao Yang, Andy J. Ma and Pong C. Yuen*, "Domain-shared Group-sparse Dictionary Learning for Unsupervised Domain Adaptation," the 32nd AAAI Conference on Artificial Intelligence (AAAI), 2018.</p>

              <!-- 水平排列的超链接 -->
              <div style="display: flex; gap: 70px; align-items: center;">
                <!--                <div style="display: flex; align-items: center;">-->
                <!--                  <a href="" style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">-->
                <!--                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">-->
                <!--                      <Promotion />-->
                <!--                    </el-icon>-->
                <!--                    <span style="font-size: 18px; font-weight: 500;">Code</span>-->
                <!--                  </a>-->
                <!--                </div>-->

                <!-- 第二个链接（Paper） -->
                <div style="display: flex; align-items: center;">
                  <a href="https://ojs.aaai.org/index.php/AAAI/article/view/12227"
                     style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                      <Tickets/>
                    </el-icon>
                    <span style="font-size: 18px; font-weight: 500;">Paper</span>
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

  <el-backtop class="mobile-backtop" :right="100" :bottom="100"/>

</template>

<style scoped>

@media (max-width: 768px) {

  .mobile-backtop {
    right: 20px !important;
    bottom: 80px !important;
  }

  /* 添加这个规则 */
  .el-button span {
    font-size: 16px !important;
    font-weight: 500 !important;

  }
  /* 卡片布局调整为垂直 */
  .el-card > div {
    display: flex;
    flex-direction: column; /* 移动端垂直排列 */
    align-items: center; /* 居中对齐 */
  }

  /* 保持PC端的悬停效果在移动端也有效 */
  .custom-button:hover span {
    color: #13393e !important;
  }

  .el-card img {
    width: 100% !important; /* 图片宽度100% */
    max-width: 350px !important; /* 保留最大宽度与PC端一致 */
    height: auto !important;
    margin: 0 auto 15px !important; /* 居中图片 */
    object-fit: contain;
  }

  /* 文字部分样式 */
  .el-card > div > div:last-child {
    width: 100% !important; /* 文字区域占满宽度 */
    padding: 0  10px;
    display: flex;
    flex-direction: column;
    align-items: center; /* 文字和链接居中 */
    text-align: center; /* 文字居中 */
  }

  /* 链接区域居中 */
  .el-card div[style*="display: flex; gap: 70px; align-items: center;"] {
    justify-content: center; /* 链接水平居中 */
    flex-wrap: wrap; /* 允许换行 */
    gap: 20px !important; /* 调整间距 */
    row-gap: 10px !important;
  }

  /* 链接布局调整 */
  .el-space {
    flex-wrap: wrap;
    gap: 15px !important;
  }


  /* 调整按钮组垂直间距 */
  .el-card .el-button {
    margin: 5px 0;
  }

  /* 保持图标和文字对齐 */
  .el-button p {
    margin: 0;
    line-height: 1.2;
  }

  /* 文字大小调整 */
  p {
    font-size: 14px !important;
  }

  /* 统一所有链接的移动端样式 */
  .el-button span,
  a[target="_blank"] span {
    font-size: 16px !important;
    font-weight: 500 !important;
    letter-spacing: 0.5px; /* 增加字母间距提升可读性 */
  }

  /* 保持图标与文字对齐 */
  .el-button,
  a[target="_blank"] {
    display: inline-flex;
    align-items: center;
  }

  /* 调整年份标题 */
  p[style*="font-size: 26px;"] {
    font-size: 30px !important;
    margin-left: 15px !important;
  }

  /* 调整Journal/Conference子标题 */
  p[style*="font-size: 26px;"] + p[style*="font-size: 26px;"] {
    font-size: 28px !important;
    margin-left: 15px !important;
    margin-top: 10px !important;
  }


  /* 强制卡片内容垂直布局 */
  .el-card > div[style*="display: flex"] {
    flex-direction: column !important;
    padding: 10px !important;
  }

  /* 加强链接容器控制 */
  div[style*="gap: 70px"] {
    gap: 8px !important;
    margin-top: 10px !important;
  }

  /* 加强Safari适配 */
  @supports (-webkit-touch-callout: none) {
    .el-card > div {
      row-gap: 0 !important;
      column-gap: 0 !important;
      padding: 0 5px !important;
    }
    .el-card p {
      word-break: break-word; /* 强制换行 */
    }
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

/* 添加按钮悬停效果 */
.custom-button:hover ::v-deep(.el-icon svg) {
  color: #13393e !important;
}
.custom-button:hover span {
  color: #13393e !important;
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
