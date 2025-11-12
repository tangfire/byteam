<script setup lang="ts">
import { Files, Postcard, Promotion, Tickets, VideoPlay } from "@element-plus/icons-vue";
import { ref, computed } from 'vue';

defineProps<{ msg: string }>()

// 统一的下载函数
const downloadFile = (filePath: string, fileName: string) => {
  const link = document.createElement('a')
  link.href = filePath
  link.download = fileName
  document.body.appendChild(link)
  link.click()
  document.body.removeChild(link)
}

// 简化的下载函数
const download_MMS_PPT = () => downloadFile('/files/MMS.pdf', 'Team_Presentation.pdf')
const download_ICIP2025_PPT = () => downloadFile('/files/icip2025.pdf', 'Team_Presentation.pdf')
const download_HBAI_PPT = () => downloadFile('/files/HBAI.pdf', 'Team_Presentation.pdf')
const download_FairFed_PPT = () => downloadFile('/files/FairFed.pdf', 'Team_Presentation.pdf')
const download_UnlockingThePotentialOfmLLMs_PPT = () => downloadFile('/files/UnlockingThePotentialOfmLLMs.pdf', 'Team_Presentation.pdf')
const download_HBAI_Poster = () => downloadFile('/files/ldzposter001.pdf', 'Team_Poster.pdf')
const download_WJC001_Poster = () => downloadFile('/files/wjcposter001.pdf', 'Team_Poster.pdf')
const download_WJC002_Poster = () => downloadFile('/files/wjcposter002.pptx', 'Team_Poster.pptx')

// 图标映射
const iconMap = {
  code: Promotion,
  paper: Tickets,
  video: VideoPlay,
  ppt: Files,
  poster: Postcard
}

const getIcon = (type: string) => iconMap[type as keyof typeof iconMap] || Promotion
const getLinkText = (type: string) => {
  if (type === 'ppt') return 'PPT'
  return type.charAt(0).toUpperCase() + type.slice(1)
}

// 出版物数据
const publications = {
  2025: {
    journal: [
      {
        id: 1,
        image: "/publications/028.png",
        title: "Multi-modal Brain Tumor Segmentation via Multi-category Interaction and Graph Co-reasoning",
        authors: "Baoyao Yang*, Dongzhe Li, Chong Yin, Fei Lyu, Xiaochen He ",
        venue: "IEEE Transactions on Multimedia, 2025",
        links: {
          // code: "https://github.com/BaoyaoGroup/JBHI_code",
          // paper: "https://ieeexplore.ieee.org/abstract/document/11205277"
        }
      },
      {
        id: 2,
        image: "/publications/026.png",
        title: "CAM-interacted Vision GNN for Multi-label Medical Images",
        authors: "Jiangchao Wang, Baoyao Yang*, Siqi Liu, Xiaoqi Zheng, Wenbin Yao* and Junxiang Chen",
        venue: "IEEE Journal of Biomedical and Health Informatics (JBHI), 2025",
        links: {
          code: "https://github.com/BaoyaoGroup/JBHI_code",
          paper: "https://ieeexplore.ieee.org/abstract/document/11205277"
        }
      }
    ],
    conference: [
      {
        id: 3,
        image: "/publications/027.png",
        title: "FedCD: A Hybrid Federated Learning Framework for Adaptive Training under Data Heterogeneity",
        authors: "Weide Zhan, Baoyao Yang*",
        venue: "Chinese Conference on Pattern Recognition and Computer Vision (PRCV), 2025",
        links: {}
      },
      {
        id: 4,
        image: "/publications/004.png",
        title: "Image-assisted Label Connective Completion for Vessel Segmentation with Insufficient Annotations",
        authors: "Xiaoqi Zheng, Baoyao Yang*, Xiuwen Fang, Xiuwen Fang, Mang Ye",
        venue: "IEEE International Conference on Acoustics, Speech, and Signal Processing (ICASSP), 2025",
        links: {
          code: "https://github.com/BaoyaoGroup/LabelCompletion",
          paper: "https://ieeexplore.ieee.org/document/10888997",
          video: { name: 'video-player-XiaoqiZheng01' }
        }
      },
      {
        id: 5,
        image: "/publications/030.png",
        title: "Unifying Spatio-Temporal Contexts for Advanced Text-Video Retrieval",
        authors: "Yanhao Huang, Baoyao Yang*, Junxiang Chen, Wenbin Yao, Dixin Chen",
        venue: "IEEE International Conference on Multimedia and Expo (ICME), 2025",
        links: {}
      },
      {
        id: 6,
        image: "/publications/024.png",
        title: "FairFed++: Closing the Fairness Gap in Federated Learning through Self-Evolving Clustered Optimization",
        authors: "Zhixiang Fang, Baoyao Yang*, Weide Zhan, Yanchao Tang, Yiqun Zhang",
        venue: "the 28th European Conference on Artificial Intelligence (ECAI), 2025",
        links: {
          code: "https://github.com/BaoyaoGroup/FairFedPlusPlus",
          ppt: { handler: download_FairFed_PPT }
        }
      },
      {
        id: 7,
        image: "/publications/029.png",
        title: "Unlocking the Potential of mLLMs: Enhancing Video-Text Retrieval through Caption Supplementation and Conical Embedding Optimization",
        authors: "Baoyao Yang*, Junxiang Chen, Wenbin Yao",
        venue: "the 28th European Conference on Artificial Intelligence (ECAI), 2025",
        links: {
          ppt: { handler: download_UnlockingThePotentialOfmLLMs_PPT }
        }
      },
      {
        id: 8,
        image: "/publications/023.png",
        title: "Simple but Effective: Sub-Volume Contrastive Learning for Class-Imbalanced Semi-Supervised 3D Medical Image Segmentation",
        authors: "Xianrun Xu, Baoyao Yang*, Wanyun Li, Jingsong Lin, Yufei Xu",
        venue: "the 33rd ACM International Conference on Multimedia (ACM MM), 2025",
        links: {
          paper: "https://dl.acm.org/doi/abs/10.1145/3746027.3755652",
          video: { name: 'video-player-XianrunXu01' }
        }
      },
      {
        id: 9,
        image: "/publications/031.png",
        title: "Multi-Agent Reinforcement Learning Algorithm Using Dynamic OW-QMIX in Complex Supply Chain Scenarios",
        authors: "ZhiQi Liu, QingHua Zhu*, An Zeng, YuZhu Ji, BaoYao Yang",
        venue: "IEEE International Conference on Systems, Man, and Cybernetics (IEEE SMC), 2025",
        links: {}
      },
      {
        id: 10,
        image: "/publications/online.png",
        title: "Action Decomposition-based Actor-Critic for Supply Chain Optimization",
        authors: "Zhengrong Chen, Qinghua Zhu*, An Zeng, Yuzhu Ji, Baoyao Yang, Dan Pan",
        venue: "IEEE International Conference on Multimedia and Expo (ICME), 2025",
        links: {}
      },
      {
        id: 11,
        image: "/publications/025.png",
        title: "Harnessing Feature Distribution Consistency for Federated Learning with Noisy",
        authors: "Yali Ma, Baoyao Yang*, Yanchao Tang, Weide Zhan, Wenyin Yang",
        venue: "IEEE International Conference on Image Processing (ICIP), 2025",
        links: {
          paper: "https://ieeexplore.ieee.org/abstract/document/11084722",
          ppt: { handler: download_ICIP2025_PPT },
          video: { name: 'video-player-YaliMa01' }
        }
      }
    ]
  },
  2024: {
    journal: [
      {
        id: 12,
        image: "/publications/002.png",
        title: "DNA-T: Deformable Neighborhood Attention Transformer for Irregular Medical Time Series",
        authors: "Jianxuan Huang, Baoyao Yang*, Kejing Yin, Jingwen Xu",
        venue: "IEEE Journal of Biomedical and Health Informatics (JBHI), 2024",
        links: {
          code: "https://github.com/BaoyaoGroup/DNA-T",
          paper: "https://ieeexplore.ieee.org/document/10510586"
        }
      },
      {
        id: 13,
        image: "/publications/005.png",
        title: "Allosteric Feature Collaboration for Model-Heterogeneous Federated Learning",
        authors: "Baoyao Yang*, PC Yuen, Yiqun Zhang, An Zeng",
        venue: "IEEE Transactions on Neural Networks and Learning Systems (TNNLS), 2024",
        links: {
          paper: "https://ieeexplore.ieee.org/abstract/document/10373104"
        }
      }
    ],
    conference: [
      {
        id: 14,
        image: "/publications/001.png",
        title: "CAM-Guided translation for unpaired weakly-supervised medical image segmentation",
        authors: "Yuebin Xie, Xiaochen He, Baoyao Yang*, Fei Lyu, Siqi Liu",
        venue: "IEEE International Conference on Multimedia and Expo (ICME), 2024",
        links: {
          code: "https://github.com/BaoyaoGroup/CAM_Guided",
          paper: "https://ieeexplore.ieee.org/abstract/document/10687752"
        }
      },
      {
        id: 15,
        image: "/publications/003.png",
        title: "Domain Dilation for Single Domain Generalization",
        authors: "Yuehui Fan, Baoyao Yang*, Meng Shen, Fei Lyu",
        venue: "IEEE International Conference on Image Processing (ICIP), 2024",
        links: {
          code: "https://github.com/BaoyaoGroup/Domain_Dilation_for_Single_Domain_Generalization",
          paper: "https://ieeexplore.ieee.org/abstract/document/10648093"
        }
      },
      {
        id: 16,
        image: "/publications/006.png",
        title: "MMS: Morphology-mixup Stylized Data Generation for Single Domain Generalization in Medical Image Segmentation",
        authors: "Xiaochen He, Baoyao Yang*, Fei Lyu",
        venue: "IEEE International Conference on Acoustics, Speech, and Signal Processing (ICASSP), 2024",
        links: {
          code: "https://github.com/BaoyaoGroup/MMS",
          paper: "https://ieeexplore.ieee.org/abstract/document/10448305",
          ppt: { handler: download_MMS_PPT }
        }
      },
      {
        id: 17,
        image: "/publications/007.png",
        title: "Multi-category Graph Reasoning for Multi-modal Brain Tumor Segmentation",
        authors: "Dongzhe Li, Baoyao Yang*, Weide Zhan, Xiaochen He",
        venue: "Medical Image Computing and Computer Assisted Intervention (MICCAI), 2024",
        links: {
          code: "https://github.com/BaoyaoGroup/Graph-Co-reasoning",
          paper: "https://link.springer.com/chapter/10.1007/978-3-031-72111-3_42",
          poster: { handler: download_HBAI_Poster }
        }
      },
      {
        id: 18,
        image: "/publications/022.png",
        title: "Multi-category Brain Tumor Segmentation via Multi-scale and Cross-category Relation Modeling",
        authors: "Dongzhe Li, Baoyao Yang*, Yuebin Xie, Weide Zhan and Jingsong Lin",
        venue: "HBAI2024: JJCAI Workshop on Human Brain and Artificial Intelligence, 2024",
        links: {
          paper: "https://link.springer.com/chapter/10.1007/978-981-96-4001-0_19",
          ppt: { handler: download_HBAI_PPT }
        }
      },
      {
        id: 19,
        image: "/publications/008.png",
        title: "Beyond Direct Relationships: Exploring Multi-Order Label Pair Dependencies for Knowledge Distillation",
        authors: "Wang J, Deng Z, Lin T, et al.",
        venue: "Proceedings of the 32nd ACM International Conference on Multimedia, 2024",
        links: {
          code: "https://github.com/BaoyaoGroup/Beyond-Direct-Relationships--Exploring-Multi-Order-Distillation-main",
          paper: "https://dl.acm.org/doi/abs/10.1145/3664647.3681029",
          poster: { handler: download_WJC001_Poster }
        }
      },
      {
        id: 20,
        image: "/publications/009.png",
        title: "A Novel Prompt Tuning for Graph Transformers: Tailoring Prompts to Graph Topologies",
        authors: "Wang J, Deng Z, Lin T, et al.",
        venue: "Proceedings of the 30th ACM SIGKDD Conference on Knowledge Discovery and Data Mining, 2024",
        links: {
          code: "https://github.com/BaoyaoGroup/TGPT_main-main",
          paper: "https://dl.acm.org/doi/10.1145/3637528.3671804",
          poster: { handler: download_WJC002_Poster }
        }
      }
    ]
  },
  2023: {
    journal: [
      {
        id: 21,
        image: "/publications/021.png",
        title: "Deep Learning for Brain MRI Confirms Patterned Pathological Progression in Alzheimer's Disease",
        authors: "Dan Pan, An Zeng*, Baoyao Yang*, Gangyong Lai, Bing Hu, Xiaowei Song, Tianzi Jiang",
        venue: "Advanced Science, 2023",
        links: {
          paper: "https://advanced.onlinelibrary.wiley.com/doi/full/10.1002/advs.202204717"
        }
      }
    ],
    conference: [
      {
        id: 22,
        image: "/publications/010.png",
        title: "Early diagnosis of Alzheimer's disease based on multimodal hypergraph attention network",
        authors: "Yi Li, Baoyao Yang*, Dan Pan, An Zeng, Yang Yang",
        venue: "International Conference on Multimedia and Expo (ICME), 2023",
        links: {
          paper: "https://ieeexplore.ieee.org/abstract/document/10219893"
        }
      }
    ]
  },
  2022: {
    journal: [
      {
        id: 23,
        image: "/publications/011.png",
        title: "Model-induced Generalization Error Bound for Information-theoretic Representation Learning in Source-data-free Unsupervised Domain Adaptation",
        authors: "Baoyao Yang, Hao-wei Yeh, Tatsuya Harada, and Pong C. Yuen*",
        venue: "IEEE Transactions on Image Processing (TIP), 2022",
        links: {
          paper: "https://ieeexplore.ieee.org/abstract/document/9640468"
        }
      },
      {
        id: 24,
        image: "/publications/012.png",
        title: "Revealing Task-relevant Model Memorization for Source-Protected Unsupervised Domain Adaptation",
        authors: "Baoyao Yang and Pong C. Yuen*",
        venue: "IEEE Transactions on Information Forensics and Security (TIFS), 2022",
        links: {
          paper: "https://ieeexplore.ieee.org/abstract/document/9705526"
        }
      },
      {
        id: 25,
        image: "/publications/013.png",
        title: "Cross-domain Missingness-aware Time Series Adaptation with Similarity Distillation in Medical Applications",
        authors: "Baoyao Yang, Mang Ye, Qingxiong Tan and Pong C. Yuen*",
        venue: "IEEE Transactions on Cybernetics (TCYB), 2022",
        links: {
          paper: "https://ieeexplore.ieee.org/abstract/document/9167415"
        }
      }
    ],
    conference: []
  },
  2021: {
    journal: [
      {
        id: 26,
        image: "/publications/015.png",
        title: "Learning Adaptive Geometry for Unsupervised Domain Adaptation",
        authors: "Baoyao Yang and Pong C. Yuen*",
        venue: "Pattern Recognition (PR), 2021",
        links: {
          paper: "https://www.sciencedirect.com/science/article/abs/pii/S0031320320304416"
        }
      }
    ],
    conference: [
      {
        id: 27,
        image: "/publications/016.png",
        title: "A Segmentation-Assisted Model for Universal Lesion Detection with Partial Labels",
        authors: "Fei Lyu, Baoyao Yang, Andy J. Ma and Pong C. Yuen*",
        venue: "International Conference on Medical Image Computing & Computer Assisted Intervention (MICCAI), 2021",
        links: {
          paper: "https://link.springer.com/chapter/10.1007/978-3-030-87240-3_12"
        }
      }
    ]
  },
  2019: {
    journal: [
      {
        id: 28,
        image: "/publications/017.png",
        title: "Body Parts Synthesis for Cross-Quality Pose Estimation",
        authors: "Baoyao Yang, Andy J. Ma and Pong C. Yuen*",
        venue: "IEEE Transactions on Circuits and Systems for Video Technology (TCSVT), 2019",
        links: {
          paper: "https://ieeexplore.ieee.org/document/8245864"
        }
      }
    ],
    conference: [
      {
        id: 29,
        image: "/publications/018.png",
        title: "Cross-Domain Visual Representations via Unsupervised Graph Alignment",
        authors: "Baoyao Yang and Pong C. Yuen*",
        venue: "the 33rd AAAI Conference on Artificial Intelligence (AAAI), 2019",
        links: {
          paper: "https://ojs.aaai.org/index.php/AAAI/article/view/4504"
        }
      }
    ]
  },
  2018: {
    journal: [
      {
        id: 30,
        image: "/publications/019.png",
        title: "Learning Domain-Shared Group-Sparse Representation for Unsupervised Domain Adaptation",
        authors: "Baoyao Yang, Andy J. Ma and Pong C. Yuen*",
        venue: "Pattern Recognition (PR), 2018",
        links: {
          paper: "https://www.sciencedirect.com/science/article/pii/S0031320318301614"
        }
      }
    ],
    conference: [
      {
        id: 31,
        image: "/publications/020.png",
        title: "Domain-shared Group-sparse Dictionary Learning for Unsupervised Domain Adaptation",
        authors: "Baoyao Yang, Andy J. Ma and Pong C. Yuen*",
        venue: "the 32nd AAAI Conference on Artificial Intelligence (AAAI), 2028",
        links: {
          paper: "https://ojs.aaai.org/index.php/AAAI/article/view/12227"
        }
      }
    ]
  }
}

// 年份数组，按降序排列
const years = computed(() => Object.keys(publications).sort((a, b) => Number(b) - Number(a)))
</script>

<template>
  <div class="publications-container">
    <!-- 动态生成年份 -->
    <div v-for="year in years" :key="year" class="year-section">
      <p class="year-title">{{ year }}</p>

      <!-- 期刊部分 -->
      <template v-if="publications[year].journal && publications[year].journal.length">
        <p class="section-title">Journal</p>
        <div v-for="pub in publications[year].journal" :key="pub.id" class="publication-card-wrapper">
          <el-card class="publication-card">
            <div class="card-content">
              <img :src="pub.image" :alt="pub.title" class="publication-image"/>
              <div class="publication-info">
                <p class="publication-title">{{ pub.authors }}, "{{ pub.title }}," {{ pub.venue }}</p>
                <div class="publication-links" v-if="Object.keys(pub.links).length > 0">
                  <template v-for="(link, type) in pub.links" :key="type">
                    <div class="link-item">
                      <!-- 代码和论文链接 -->
                      <a
                          v-if="type === 'code' || type === 'paper'"
                          :href="link"
                          class="publication-link"
                          target="_blank"
                      >
                        <el-icon size="25">
                          <component :is="getIcon(type)" />
                        </el-icon>
                        <span>{{ getLinkText(type) }}</span>
                      </a>

                      <!-- 视频链接 -->
                      <router-link
                          v-else-if="type === 'video'"
                          :to="link"
                          class="publication-link"
                      >
                        <el-icon size="25">
                          <VideoPlay />
                        </el-icon>
                        <span>Video</span>
                      </router-link>

                      <!-- PPT和海报下载按钮 -->
                      <el-button
                          v-else-if="type === 'ppt' || type === 'poster'"
                          class="custom-button"
                          type="primary"
                          @click="link.handler"
                      >
                        <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                          <component :is="getIcon(type)" />
                        </el-icon>
                        <span class="button-text">{{ getLinkText(type) }}</span>
                      </el-button>
                    </div>
                  </template>
                </div>
              </div>
            </div>
          </el-card>
        </div>
      </template>

      <!-- 会议部分 -->
      <template v-if="publications[year].conference && publications[year].conference.length">
        <p class="section-title">Conference</p>
        <div v-for="pub in publications[year].conference" :key="pub.id" class="publication-card-wrapper">
          <el-card class="publication-card">
            <div class="card-content">
              <img :src="pub.image" :alt="pub.title" class="publication-image"/>
              <div class="publication-info">
                <p class="publication-title">{{ pub.authors }}, "{{ pub.title }}," {{ pub.venue }}</p>
                <div class="publication-links" v-if="Object.keys(pub.links).length > 0">
                  <template v-for="(link, type) in pub.links" :key="type">
                    <div class="link-item">
                      <!-- 代码和论文链接 -->
                      <a
                          v-if="type === 'code' || type === 'paper'"
                          :href="link"
                          class="publication-link"
                          target="_blank"
                      >
                        <el-icon size="25">
                          <component :is="getIcon(type)" />
                        </el-icon>
                        <span>{{ getLinkText(type) }}</span>
                      </a>

                      <!-- 视频链接 -->
                      <router-link
                          v-else-if="type === 'video'"
                          :to="link"
                          class="publication-link"
                      >
                        <el-icon size="25">
                          <VideoPlay />
                        </el-icon>
                        <span>Video</span>
                      </router-link>

                      <!-- PPT和海报下载按钮 -->
                      <el-button
                          v-else-if="type === 'ppt' || type === 'poster'"
                          class="custom-button"
                          type="primary"
                          @click="link.handler"
                      >
                        <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                          <component :is="getIcon(type)" />
                        </el-icon>
                        <span class="button-text">{{ getLinkText(type) }}</span>
                      </el-button>
                    </div>
                  </template>
                </div>
              </div>
            </div>
          </el-card>
        </div>
      </template>
    </div>
  </div>

  <div class="bottom-spacer"></div>
  <el-backtop class="mobile-backtop" :right="100" :bottom="100"/>
</template>

<style scoped>
/* 统一变量管理 */
:root {
  --primary-color: #7d1231;
  --hover-color: #13393e;
  --card-max-width: 1000px;
  --mobile-breakpoint: 768px;
}

/* 基础容器 */
.publications-container {
  width: 100%;
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
  box-sizing: border-box;
}

/* 年份和标题样式 */
.year-title {
  margin-left: 8px;
  color: #7d1231;
  font-size: 26px;
  font-weight: 600;
  margin-bottom: 10px;
}

.section-title {
  margin-left: 8px;
  color: #7d1231;
  font-size: 26px;
  margin-top: 20px;
  margin-bottom: 15px;
}

/* 卡片布局 */
.publication-card-wrapper {
  margin-bottom: 35px;
}

.publication-card {
  max-width: var(--card-max-width);
  margin: 0 auto;
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
  border: 1px solid #eaeaea;
}

.publication-card:hover {
  box-shadow: 0 4px 20px 0 rgba(0, 0, 0, 0.15);
  transform: translateY(-2px);
}

.card-content {
  display: flex;
  align-items: flex-start;
  padding: 20px;
  gap: 20px;
}

.publication-image {
  width: 350px;
  height: 200px;
  object-fit: contain;
  flex-shrink: 0;
  border-radius: 4px;
  transition: transform 0.3s ease;
}

.publication-image:hover {
  transform: scale(1.02);
}

.publication-info {
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: flex-start;
  flex: 1;
  min-height: 200px;
}

.publication-title {
  font-size: 18px;
  line-height: 1.5;
  color: #333;
  margin-bottom: 15px;
}

.publication-links {
  display: flex;
  gap: 70px;
  align-items: center;
  margin-top: auto;
  flex-wrap: wrap;
}

.link-item {
  display: flex;
  align-items: center;
}

.publication-link {
  color: var(--primary-color);
  text-decoration: none;
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 18px;
  font-weight: 500;
  transition: all 0.3s ease;
  padding: 6px 12px;
  border-radius: 4px;
}

.publication-link:hover {
  color: var(--hover-color);
  background-color: rgba(125, 18, 49, 0.05);
}

.publication-link:hover ::v-deep(.el-icon svg) {
  color: var(--hover-color) !important;
}

.bottom-spacer {
  height: 50px;
}

/* 按钮样式 */
.custom-button {
  background-color: white !important;
  border-color: #eaeaea !important;
  padding: 8px 16px;
  box-shadow: none !important;
  transition: all 0.3s ease !important;
  border-radius: 4px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.custom-button:hover {
  background-color: rgba(125, 18, 49, 0.05) !important;
  border-color: var(--primary-color) !important;
}

.custom-button:hover .button-text {
  color: var(--hover-color) !important;
}

.custom-button:hover ::v-deep(.el-icon svg) {
  color: var(--hover-color) !important;
}

.button-text {
  font-size: 18px;
  font-weight: 500;
  color: #7d1231;
  transition: color 0.3s ease;
}

.mobile-backtop {
  right: 100px;
  bottom: 100px;
}

/* 保证 el-icon 的图标颜色在初始时是正确的 */
::v-deep(.el-icon svg) {
  color: #7d1231 !important;
}

/* 移动端适配 */
@media (max-width: 768px) {
  .publications-container {
    padding: 0 15px;
  }

  .card-content {
    flex-direction: column;
    text-align: center;
    gap: 15px;
    padding: 15px;
  }

  .publication-image {
    width: 100%;
    max-width: 100%;
    height: auto;
    margin: 0;
  }

  .publication-info {
    align-items: center;
    width: 100%;
    min-height: auto;
  }

  .publication-title {
    font-size: 16px;
    text-align: left;
    width: 100%;
  }

  .publication-links {
    gap: 20px;
    flex-wrap: wrap;
    justify-content: center;
    width: 100%;
  }

  .link-item {
    margin: 5px;
  }

  .publication-link {
    font-size: 16px;
    padding: 5px 10px;
  }

  .year-title,
  .section-title {
    font-size: 22px;
    margin-left: 15px;
  }

  .publication-card {
    margin: 0 0 20px;
  }

  .mobile-backtop {
    right: 20px !important;
    bottom: 80px !important;
  }

  .publication-card-wrapper {
    margin-bottom: 20px;
  }

  .button-text {
    font-size: 16px;
  }

  .custom-button {
    padding: 6px 12px;
  }

  .publication-links {
    gap: 15px;
  }
}

/* 小屏幕手机优化 */
@media (max-width: 480px) {
  .publication-links {
    gap: 10px;
  }

  .publication-title {
    font-size: 14px;
  }

  .publication-link span,
  .button-text {
    font-size: 14px !important;
  }

  .publications-container {
    padding: 0 10px;
  }

  .link-item {
    margin: 3px;
  }
}
</style>