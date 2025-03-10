<script setup lang="ts">


import {Files, Postcard, Tickets} from "@element-plus/icons-vue";

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

      <p class="title">Welcome to BYML</p>

      <p class="text">Our team focuses on medical time series representation learning and multimodal data fusion,
        privacy-preserving
        federated learning with heterogeneous collaboration, automated lesion annotation with weakly supervised models,
        short video multimodal understanding and contrastive learning, intelligent perception and privacy security in
        industrial IoT, theoretical research on algorithm generalization and interpretability, as well as reinforcement
        learning and intelligent scheduling in industrial software.</p>
    </div>

    <el-divider/>


    <div>


      <el-space direction="vertical" :size="30" style="width: 100%">

        <p class="title_research" style="font-size: 32px;color: #7d1231;">Latest Research</p>
        <div>
          <div style=" align-items: center;">

            <el-card class="card_1"
                     style="max-width: 1000px; margin: 0 auto; display: flex; align-items: center; justify-content: flex-start;">
              <div style="display: flex; align-items: center; justify-content: flex-start;">
                <!-- 图片左边 -->
                <img
                    src="/publications/004.png"
                    style="width: 350px; height: 200px; margin-right: 20px;object-fit: contain;"/>
                <!-- 文字右边 -->
                <div
                    style="display: flex; flex-direction: column; justify-content: flex-start; align-items: flex-start;">
                  <!-- 使文字稍微往上移 -->
                  <p style="margin-top: -10px;font-size: 18px">X. Zheng, B. Yang, X. Fang, W. Yao and M. Ye,
                    "Image-assisted Label Connective Completion for Vessel Segmentation with Insufficient Annotations,"
                    ICASSP 2025 - 2025 IEEE International Conference on Acoustics, Speech and Signal Processing
                    (ICASSP), Hyderabad, India, 2025, pp. 1-5, doi: 10.1109/ICASSP49660.2025.10888997.</p>

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
                      <a href="https://ieeexplore.ieee.org/document/10888997"
                         style="color: #7d1231; text-decoration: none; display: flex; align-items: center;"
                         target="_blank">
                        <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                          <Tickets/>
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
                class="card_2"
                style="max-width: 1000px; margin: 0 auto; display: flex; align-items: center; justify-content: flex-start;">
              <div style="display: flex; align-items: center; justify-content: flex-start;">
                <!-- 图片左边 -->
                <img
                    src="/publications/002.png"
                    style="width: 350px; height: 200px; margin-right: 20px;object-fit: contain;"/>
                <!-- 文字右边 -->
                <div
                    style="display: flex; flex-direction: column; justify-content: flex-start; align-items: flex-start;">
                  <!-- 使文字稍微往上移 -->
                  <p style="margin-top: -10px;font-size: 18px">Huang J, Yang B, Yin K, et al. Dna-t: Deformable
                    neighborhood attention transformer for irregular medical time series[J]. IEEE Journal of Biomedical
                    and Health Informatics, 2024.</p>

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
                      <a href="https://ieeexplore.ieee.org/document/10510586"
                         style="color: #7d1231; text-decoration: none; display: flex; align-items: center;"
                         target="_blank">
                        <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                          <Tickets/>
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
                class="card_3"
                style="max-width: 1000px; margin: 0 auto; display: flex; align-items: center; justify-content: flex-start;">
              <div style="display: flex; align-items: center; justify-content: flex-start;">
                <!-- 图片左边 -->
                <img
                    src="/publications/001.png"
                    style="width: 350px; height: 200px; margin-right: 20px;object-fit: contain;"/>
                <!-- 文字右边 -->
                <div
                    style="display: flex; flex-direction: column; justify-content: flex-start; align-items: flex-start;">
                  <!-- 使文字稍微往上移 -->
                  <p style="margin-top: -10px;font-size: 18px">Xie Y, He X, Yang B, et al. CAM-Guided Translation for
                    Unpaired Weakly-Supervised Medical Image Segmentation[C]//2024 IEEE International Conference on
                    Multimedia and Expo (ICME). IEEE, 2024: 1-6.</p>

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
                      <a href="https://ieeexplore.ieee.org/abstract/document/10687752"
                         style="color: #7d1231; text-decoration: none; display: flex; align-items: center;"
                         target="_blank">
                        <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                          <Tickets/>
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
                class="card_4"
                style="max-width: 1000px; margin: 0 auto; display: flex; align-items: center; justify-content: flex-start;">
              <div style="display: flex; align-items: center; justify-content: flex-start;">
                <!-- 图片左边 -->
                <img
                    src="/publications/003.png"
                    style="width: 350px; height: 200px; margin-right: 20px;object-fit: contain;"/>
                <!-- 文字右边 -->
                <div
                    style="display: flex; flex-direction: column; justify-content: flex-start; align-items: flex-start;">
                  <!-- 使文字稍微往上移 -->
                  <p style="margin-top: -10px;font-size: 18px">Fan Y, Yang B, Shen M, et al. Domain Dilation for Single
                    Domain Generalization[C]//2024 IEEE International Conference on Image Processing (ICIP). IEEE, 2024:
                    3931-3937.</p>

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
                      <a href="https://ieeexplore.ieee.org/abstract/document/10648093"
                         style="color: #7d1231; text-decoration: none; display: flex; align-items: center;"
                         target="_blank">
                        <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                          <Tickets/>
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
                class="card_5"
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
                    <!--                    <span style="font-size: 18px; font-weight: 520;">Code</span>-->
                    <!--                  </a>-->
                    <!--                </div>-->

                    <!-- 第二个链接（Paper） -->
                    <div style="display: flex; align-items: center;">
                      <a href="https://ieeexplore.ieee.org/abstract/document/10373104"
                         style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                        <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                          <Tickets/>
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
                class="card_6"
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
                      <a href="https://ieeexplore.ieee.org/abstract/document/10448305"
                         style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                        <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                          <Tickets/>
                        </el-icon>
                        <span style="font-size: 18px; font-weight: 520;">Paper</span>
                      </a>
                    </div>
                    <div>
                      <el-button class="custom-button"
                                 type="primary"
                                 @click="download_MMS_PPT"
                                 style="background-color: white; border-color: white;"
                      >

                        <el-icon size="25" style="margin-right: 8px; vertical-align: middle;"><Files /></el-icon>

                        <p style="color: #7d1231;font-size: 18px">PPT</p>
                      </el-button></div>
                  </div>
                </div>
              </div>
            </el-card>

            <div style="height: 35px"></div>

            <el-card
                class="card_7"
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
                      <a href="https://link.springer.com/chapter/10.1007/978-3-031-72111-3_42"
                         style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                        <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                          <Tickets/>
                        </el-icon>
                        <span style="font-size: 18px; font-weight: 520;">Paper</span>
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

                        <p style="color: #7d1231;font-size: 18px">PPT</p>
                      </el-button></div>
                    <div>
                      <el-button
                          class="custom-button"
                          type="primary"
                          @click="download_HBAI_Poster"
                          style="background-color: white; border-color: white;"
                      >

                        <el-icon  size="25" style="margin-right: 8px; vertical-align: middle;"><Postcard /></el-icon>

                        <p  style="color: #7d1231;font-size: 18px">Poster</p>
                      </el-button></div>
                  </div>
                </div>
              </div>
            </el-card>

            <div style="height: 35px"></div>

            <el-card
                class="card_8"
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
                      <a href="https://dl.acm.org/doi/abs/10.1145/3664647.3681029"
                         style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                        <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                          <Tickets/>
                        </el-icon>
                        <span style="font-size: 18px; font-weight: 520;">Paper</span>
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

                    <!--                    <p style="color: #7d1231;font-size: 18px">PPT</p>-->
                    <!--                  </el-button></div>-->
                    <div>
                      <el-button
                          class="custom-button"
                          type="primary"
                          @click="download_WJC001_Poster"
                          style="background-color: white; border-color: white;"
                      >

                        <el-icon  size="25" style="margin-right: 8px; vertical-align: middle;"><Postcard /></el-icon>

                        <p  style="color: #7d1231;font-size: 18px">Poster</p>
                      </el-button></div>
                  </div>
                </div>
              </div>
            </el-card>

            <div style="height: 35px"></div>



            <el-card
                class="card_9"
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
                      <a href="https://dl.acm.org/doi/10.1145/3637528.3671804"
                         style="color: #7d1231; text-decoration: none; display: flex; align-items: center;" target="_blank">
                        <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                          <Tickets/>
                        </el-icon>
                        <span style="font-size: 18px; font-weight: 520;">Paper</span>
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

                    <!--                    <p style="color: #7d1231;font-size: 18px">PPT</p>-->
                    <!--                  </el-button></div>-->
                    <div>
                      <el-button
                          class="custom-button"
                          type="primary"
                          @click="download_WJC002_Poster"
                          style="background-color: white; border-color: white;"
                      >

                        <el-icon  size="25" style="margin-right: 8px; vertical-align: middle;"><Postcard /></el-icon>

                        <p  style="color: #7d1231;font-size: 18px">Poster</p>
                      </el-button></div>
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

    <el-backtop :right="100" :bottom="100"/>
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

.title {
  opacity: 0;
  animation: fadeInTitle 1.5s forwards; /* 标题渐显 */
  animation-delay: 0.5s; /* 延迟出现 */
}

.text {
  opacity: 0;
  animation: fadeInText 2.5s forwards; /* 渐显效果 */
  animation-delay: 1s; /* 延迟出现 */
}

.title_research {
  opacity: 0;
  animation: fadeInTitle 1.5s forwards; /* 渐显效果 */
  animation-delay: 1.5s; /* 延迟出现 */
}

.card_1 {
  opacity: 0;
  animation: fadeInCard 2s forwards; /* 渐显效果 */
  animation-delay: 2s; /* 延迟出现 */
}

.card_2 {
  opacity: 0;
  animation: fadeInCard 2s forwards; /* 渐显效果 */
  animation-delay: 2.5s; /* 延迟出现 */
}

.card_3 {
  opacity: 0;
  animation: fadeInCard 2s forwards; /* 渐显效果 */
  animation-delay: 3s; /* 延迟出现 */
}

.card_4 {
  opacity: 0;
  animation: fadeInCard 2s forwards; /* 渐显效果 */
  animation-delay: 3.5s; /* 延迟出现 */
}

.card_5{
  opacity: 0;
  animation: fadeInCard 2s forwards; /* 渐显效果 */
  animation-delay: 4s; /* 延迟出现 */
}

.card_6{
  opacity: 0;
  animation: fadeInCard 2s forwards; /* 渐显效果 */
  animation-delay: 4.5s; /* 延迟出现 */
}

.card_7{
  opacity: 0;
  animation: fadeInCard 2s forwards; /* 渐显效果 */
  animation-delay: 5s; /* 延迟出现 */
}

.card_8{
  opacity: 0;
  animation: fadeInCard 2s forwards; /* 渐显效果 */
  animation-delay: 5.5s; /* 延迟出现 */
}

.card_9{
  opacity: 0;
  animation: fadeInCard 2s forwards; /* 渐显效果 */
  animation-delay: 6s; /* 延迟出现 */
}



/* 添加按钮悬停效果 */
.custom-button:hover ::v-deep(.el-icon svg) {
  color: #13393e !important;
}
.custom-button:hover p {
  color: #13393e !important;
}

/* 鼠标移到超链接或者图标上时变色 */
a:hover {
  color: #13393e !important; /* 设置悬停时的颜色 */
}

a:hover ::v-deep(.el-icon svg) {
  color: #13393e !important; /* 设置悬停时图标的颜色 */
}


/* 渐变出现的动画 */
@keyframes fadeIn {
  to {
    opacity: 1;
  }
}

@keyframes fadeInTitle {
  to {
    opacity: 1;
  }
}

@keyframes fadeInText {
  to {
    opacity: 1;
  }
}

@keyframes fadeInCard {
  to {
    opacity: 1;
  }
}

</style>
