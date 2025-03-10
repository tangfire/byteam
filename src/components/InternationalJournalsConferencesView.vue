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
                  <el-button
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
                      type="primary"
                      @click="download_HBAI_PPT"
                      style="background-color: white; border-color: white;"
                  >

                    <el-icon size="25" style="margin-right: 8px; vertical-align: middle;"><Files /></el-icon>

                    <p style="color: #7d1231;font-size: 18px">PPT</p>
                  </el-button></div>
                <div>
                  <el-button

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

            <ul class="patent-list">
              <li>Yi Li, Baoyao Yang, Dan Pan, An  Zeng, Yang Yang, "Early diagnosis of Alzheimer's disease based on multimodal hypergraph attention network", International Conference on Multimedia and Expo (ICME), 2023.</li>


            </ul>


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
        <!-- 每个卡片 -->
        <el-card
            style="max-width: 1000px; margin: 0 auto; display: flex; align-items: center; justify-content: flex-start;">
          <div style="display: flex; align-items: center; justify-content: flex-start;">

            <ul class="patent-list">
              <li>Baoyao Yang, Hao-wei Yeh, Tatsuya Harada, and Pong C. Yuen*, "Model-induced Generalization Error Bound for Information-theoretic Representation Learning in Source-data-free Unsupervised Domain Adaptation," IEEE Transactions on Image Processing (TIP), 2022.</li>
              <li>Baoyao Yang and Pong C. Yuen*, "Revealing Task-relevant Model Memorization for Source-Protected Unsupervised Domain Adaptation," IEEE Transactions on Information Forensics and Security (TIFS), 2022.</li>
              <li>Baoyao Yang, Mang Ye, Qingxiong Tan and Pong C. Yuen*, "Cross-domain Missingness-aware Time Series Adaptation with Similarity Distillation in Medical Applications," IEEE Transactions on Cybernetics (TCYB), 2022.</li>

            </ul>


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

            <ul class="patent-list">
              <li>Baoyao Yang and Pong C. Yuen*, "Learning Adaptive Geometry for Unsupervised Domain Adaptation," Pattern Recognition (PR), 2021.</li>


            </ul>


          </div>
        </el-card>

        <p style="margin-left: 8px; color: #7d1231; font-size: 26px;"> Conference</p> <!-- 增加文字大小 -->
        <!-- 每个卡片 -->
        <el-card
            style="max-width: 1000px; margin: 0 auto; display: flex; align-items: center; justify-content: flex-start;">
          <div style="display: flex; align-items: center; justify-content: flex-start;">

            <ul class="patent-list">
              <li>Fei Lyu, Baoyao Yang, Andy J. Ma and Pong C. Yuen*, "A Segmentation-Assisted Model for Universal Lesion Detection with Partial Labels," International Conference on Medical Image Computing & Computer Assisted Intervention (MICCAI), 2021. </li>


            </ul>


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

            <ul class="patent-list">
              <li>Baoyao Yang, Andy J. Ma and Pong C. Yuen*, "Body Parts Synthesis for Cross-Quality Pose Estimation." IEEE Transactions on Circuits and Systems for Video Technology (TCSVT), 2019. </li>


            </ul>


          </div>
        </el-card>

        <p style="margin-left: 8px; color: #7d1231; font-size: 26px;"> Conference</p> <!-- 增加文字大小 -->
        <!-- 每个卡片 -->
        <el-card
            style="max-width: 1000px; margin: 0 auto; display: flex; align-items: center; justify-content: flex-start;">
          <div style="display: flex; align-items: center; justify-content: flex-start;">

            <ul class="patent-list">
              <li>Baoyao Yang and Pong C. Yuen*, "Cross-Domain Visual Representations via Unsupervised Graph Alignment," the 33rd AAAI Conference on Artificial Intelligence (AAAI), 2019.</li>


            </ul>


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

            <ul class="patent-list">
              <li>Baoyao Yang, Andy J. Ma and Pong C. Yuen*, "Learning Domain-Shared Group-Sparse Representation for Unsupervised Domain Adaptation," Pattern Recognition (PR), 2018. </li>


            </ul>


          </div>
        </el-card>


        <p style="margin-left: 8px; color: #7d1231; font-size: 26px;"> Conference</p> <!-- 增加文字大小 -->
        <!-- 每个卡片 -->
        <el-card
            style="max-width: 1000px; margin: 0 auto; display: flex; align-items: center; justify-content: flex-start;">
          <div style="display: flex; align-items: center; justify-content: flex-start;">

            <ul class="patent-list">
              <li>Baoyao Yang, Andy J. Ma and Pong C. Yuen*, "Domain-shared Group-sparse Dictionary Learning for Unsupervised Domain Adaptation," the 32nd AAAI Conference on Artificial Intelligence (AAAI), 2018.</li>


            </ul>


          </div>
        </el-card>


      </div>


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
