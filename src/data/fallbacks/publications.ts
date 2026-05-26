export type PublicationLinkValue = string | { path?: string; url?: string; handler?: () => void }
export type PublicationLinkMap = Record<string, PublicationLinkValue>
export type PublicationViewItem = {
  id?: number
  image: string
  title: string
  authors: string
  venue: string
  links: PublicationLinkMap
}
export type PublicationGroups = Record<string, { journal: PublicationViewItem[]; conference: PublicationViewItem[] }>

type DownloadFile = (filePath: string, fileName: string) => void

export const createFallbackPublicationGroups = (downloadFile: DownloadFile): PublicationGroups => {
  const downloadMMS = () => downloadFile('/files/MMS.pdf', 'Team_Presentation.pdf')
  const downloadICIP2025 = () => downloadFile('/files/icip2025.pdf', 'Team_Presentation.pdf')
  const downloadHBAI = () => downloadFile('/files/HBAI.pdf', 'Team_Presentation.pdf')
  const downloadFairFed = () => downloadFile('/files/FairFed.pdf', 'Team_Presentation.pdf')
  const downloadMLLMs = () => downloadFile('/files/UnlockingThePotentialOfmLLMs.pdf', 'Team_Presentation.pdf')
  const downloadHBAIPoster = () => downloadFile('/files/ldzposter001.pdf', 'Team_Poster.pdf')
  const downloadWJC001Poster = () => downloadFile('/files/wjcposter001.pdf', 'Team_Poster.pdf')
  const downloadWJC002Poster = () => downloadFile('/files/wjcposter002.pptx', 'Team_Poster.pptx')

  return {
    2026: {
      journal: [
        {
          id: 1,
          image: '/publications/028.png',
          title: 'Multi-modal Brain Tumor Segmentation via Multi-category Interaction and Graph Co-reasoning',
          authors: 'Baoyao Yang*, Dongzhe Li, Chong Yin, Fei Lyu, Xiaochen He ',
          venue: 'IEEE Transactions on Multimedia, 2026',
          links: {
            paper: 'https://link.springer.com/chapter/10.1007/978-981-96-4001-0_19',
          },
        },
      ],
      conference: [],
    },
    2025: {
      journal: [
        {
          id: 2,
          image: '/publications/026.png',
          title: 'CAM-interacted Vision GNN for Multi-label Medical Images',
          authors: 'Jiangchao Wang, Baoyao Yang*, Siqi Liu, Xiaoqi Zheng, Wenbin Yao* and Junxiang Chen',
          venue: 'IEEE Journal of Biomedical and Health Informatics (JBHI), 2025',
          links: {
            code: 'https://github.com/BaoyaoGroup/JBHI_code',
            paper: 'https://ieeexplore.ieee.org/abstract/document/11205277',
          },
        },
      ],
      conference: [
        {
          id: 3,
          image: '/publications/027.png',
          title: 'FedCD: A Hybrid Federated Learning Framework for Adaptive Training under Data Heterogeneity',
          authors: 'Weide Zhan, Baoyao Yang*',
          venue: 'Chinese Conference on Pattern Recognition and Computer Vision (PRCV), 2025',
          links: {
            paper: 'https://ieeexplore.ieee.org/abstract/document/10443215',
          },
        },
        {
          id: 4,
          image: '/publications/004.png',
          title: 'Image-assisted Label Connective Completion for Vessel Segmentation with Insufficient Annotations',
          authors: 'Xiaoqi Zheng, Baoyao Yang*, Xiuwen Fang, Xiuwen Fang, Mang Ye',
          venue: 'IEEE International Conference on Acoustics, Speech, and Signal Processing (ICASSP), 2025',
          links: {
            code: 'https://github.com/BaoyaoGroup/LabelCompletion',
            paper: 'https://ieeexplore.ieee.org/document/10888997',
            video: { path: '/video/video-xiaoqi-zheng-01' },
          },
        },
        {
          id: 5,
          image: '/publications/030.png',
          title: 'Unifying Spatio-Temporal Contexts for Advanced Text-Video Retrieval',
          authors: 'Yanhao Huang, Baoyao Yang*, Junxiang Chen, Wenbin Yao, Dixin Chen',
          venue: 'IEEE International Conference on Multimedia and Expo (ICME), 2025',
          links: {
            paper: 'https://ieeexplore.ieee.org/abstract/document/11209054/',
          },
        },
        {
          id: 6,
          image: '/publications/024.png',
          title: 'FairFed++: Closing the Fairness Gap in Federated Learning through Self-Evolving Clustered Optimization',
          authors: 'Zhixiang Fang, Baoyao Yang*, Weide Zhan, Yanchao Tang, Yiqun Zhang',
          venue: 'the 28th European Conference on Artificial Intelligence (ECAI), 2025',
          links: {
            code: 'https://github.com/BaoyaoGroup/FairFedPlusPlus',
            ppt: { handler: downloadFairFed },
          },
        },
        {
          id: 7,
          image: '/publications/029.png',
          title: 'Unlocking the Potential of mLLMs: Enhancing Video-Text Retrieval through Caption Supplementation and Conical Embedding Optimization',
          authors: 'Baoyao Yang*, Junxiang Chen, Wenbin Yao',
          venue: 'the 28th European Conference on Artificial Intelligence (ECAI), 2025',
          links: {
            ppt: { handler: downloadMLLMs },
          },
        },
        {
          id: 8,
          image: '/publications/023.png',
          title: 'Simple but Effective: Sub-Volume Contrastive Learning for Class-Imbalanced Semi-Supervised 3D Medical Image Segmentation',
          authors: 'Xianrun Xu, Baoyao Yang*, Wanyun Li, Jingsong Lin, Yufei Xu',
          venue: 'the 33rd ACM International Conference on Multimedia (ACM MM), 2025',
          links: {
            paper: 'https://dl.acm.org/doi/abs/10.1145/3746027.3755652',
            video: { path: '/video/video-xianrun-xu-01' },
          },
        },
        {
          id: 9,
          image: '/publications/031.png',
          title: 'Multi-Agent Reinforcement Learning Algorithm Using Dynamic OW-QMIX in Complex Supply Chain Scenarios',
          authors: 'ZhiQi Liu, QingHua Zhu*, An Zeng, YuZhu Ji, BaoYao Yang',
          venue: 'IEEE International Conference on Systems, Man, and Cybernetics (IEEE SMC), 2025',
          links: {
            paper: 'https://ieeexplore.ieee.org/abstract/document/11342630',
          },
        },
        {
          id: 10,
          image: '/publications/032.png',
          title: 'Action Decomposition-based Actor-Critic for Supply Chain Optimization',
          authors: 'Zhengrong Chen, Qinghua Zhu*, An Zeng, Yuzhu Ji, Baoyao Yang, Dan Pan',
          venue: 'IEEE International Conference on Multimedia and Expo (ICME), 2025',
          links: {
            paper: 'https://ieeexplore.ieee.org/abstract/document/11210182',
          },
        },
        {
          id: 11,
          image: '/publications/025.png',
          title: 'Harnessing Feature Distribution Consistency for Federated Learning with Noisy',
          authors: 'Yali Ma, Baoyao Yang*, Yanchao Tang, Weide Zhan, Wenyin Yang',
          venue: 'IEEE International Conference on Image Processing (ICIP), 2025',
          links: {
            paper: 'https://ieeexplore.ieee.org/abstract/document/11084722',
            ppt: { handler: downloadICIP2025 },
            video: { path: '/video/video-yali-ma-01' },
          },
        },
      ],
    },
    2024: {
      journal: [
        {
          id: 12,
          image: '/publications/002.png',
          title: 'DNA-T: Deformable Neighborhood Attention Transformer for Irregular Medical Time Series',
          authors: 'Jianxuan Huang, Baoyao Yang*, Kejing Yin, Jingwen Xu',
          venue: 'IEEE Journal of Biomedical and Health Informatics (JBHI), 2024',
          links: {
            code: 'https://github.com/BaoyaoGroup/DNA-T',
            paper: 'https://ieeexplore.ieee.org/document/10510586',
          },
        },
        {
          id: 13,
          image: '/publications/005.png',
          title: 'Allosteric Feature Collaboration for Model-Heterogeneous Federated Learning',
          authors: 'Baoyao Yang*, PC Yuen, Yiqun Zhang, An Zeng',
          venue: 'IEEE Transactions on Neural Networks and Learning Systems (TNNLS), 2024',
          links: {
            paper: 'https://ieeexplore.ieee.org/abstract/document/10373104',
          },
        },
      ],
      conference: [
        {
          id: 14,
          image: '/publications/001.png',
          title: 'CAM-Guided translation for unpaired weakly-supervised medical image segmentation',
          authors: 'Yuebin Xie, Xiaochen He, Baoyao Yang*, Fei Lyu, Siqi Liu',
          venue: 'IEEE International Conference on Multimedia and Expo (ICME), 2024',
          links: {
            code: 'https://github.com/BaoyaoGroup/CAM_Guided',
            paper: 'https://ieeexplore.ieee.org/abstract/document/10687752',
          },
        },
        {
          id: 15,
          image: '/publications/003.png',
          title: 'Domain Dilation for Single Domain Generalization',
          authors: 'Yuehui Fan, Baoyao Yang*, Meng Shen, Fei Lyu',
          venue: 'IEEE International Conference on Image Processing (ICIP), 2024',
          links: {
            code: 'https://github.com/BaoyaoGroup/Domain_Dilation_for_Single_Domain_Generalization',
            paper: 'https://ieeexplore.ieee.org/abstract/document/10648093',
          },
        },
        {
          id: 16,
          image: '/publications/006.png',
          title: 'MMS: Morphology-mixup Stylized Data Generation for Single Domain Generalization in Medical Image Segmentation',
          authors: 'Xiaochen He, Baoyao Yang*, Fei Lyu',
          venue: 'IEEE International Conference on Acoustics, Speech, and Signal Processing (ICASSP), 2024',
          links: {
            code: 'https://github.com/BaoyaoGroup/MMS',
            paper: 'https://ieeexplore.ieee.org/abstract/document/10448305',
            ppt: { handler: downloadMMS },
          },
        },
        {
          id: 17,
          image: '/publications/007.png',
          title: 'Multi-category Graph Reasoning for Multi-modal Brain Tumor Segmentation',
          authors: 'Dongzhe Li, Baoyao Yang*, Weide Zhan, Xiaochen He',
          venue: 'Medical Image Computing and Computer Assisted Intervention (MICCAI), 2024',
          links: {
            code: 'https://github.com/BaoyaoGroup/Graph-Co-reasoning',
            paper: 'https://link.springer.com/chapter/10.1007/978-3-031-72111-3_42',
            poster: { handler: downloadHBAIPoster },
          },
        },
        {
          id: 18,
          image: '/publications/022.png',
          title: 'Multi-category Brain Tumor Segmentation via Multi-scale and Cross-category Relation Modeling',
          authors: 'Dongzhe Li, Baoyao Yang*, Yuebin Xie, Weide Zhan and Jingsong Lin',
          venue: 'HBAI2024: JJCAI Workshop on Human Brain and Artificial Intelligence, 2024',
          links: {
            paper: 'https://link.springer.com/chapter/10.1007/978-981-96-4001-0_19',
            ppt: { handler: downloadHBAI },
          },
        },
        {
          id: 19,
          image: '/publications/008.png',
          title: 'Beyond Direct Relationships: Exploring Multi-Order Label Pair Dependencies for Knowledge Distillation',
          authors: 'Wang J, Deng Z, Lin T, et al.',
          venue: 'Proceedings of the 32nd ACM International Conference on Multimedia, 2024',
          links: {
            code: 'https://github.com/BaoyaoGroup/Beyond-Direct-Relationships--Exploring-Multi-Order-Distillation-main',
            paper: 'https://dl.acm.org/doi/abs/10.1145/3664647.3681029',
            poster: { handler: downloadWJC001Poster },
          },
        },
        {
          id: 20,
          image: '/publications/009.png',
          title: 'A Novel Prompt Tuning for Graph Transformers: Tailoring Prompts to Graph Topologies',
          authors: 'Wang J, Deng Z, Lin T, et al.',
          venue: 'Proceedings of the 30th ACM SIGKDD Conference on Knowledge Discovery and Data Mining, 2024',
          links: {
            code: 'https://github.com/BaoyaoGroup/TGPT_main-main',
            paper: 'https://dl.acm.org/doi/10.1145/3637528.3671804',
            poster: { handler: downloadWJC002Poster },
          },
        },
      ],
    },
    2023: {
      journal: [
        {
          id: 21,
          image: '/publications/021.png',
          title: "Deep Learning for Brain MRI Confirms Patterned Pathological Progression in Alzheimer's Disease",
          authors: 'Dan Pan, An Zeng*, Baoyao Yang*, Gangyong Lai, Bing Hu, Xiaowei Song, Tianzi Jiang',
          venue: 'Advanced Science, 2023',
          links: {
            paper: 'https://advanced.onlinelibrary.wiley.com/doi/full/10.1002/advs.202204717',
          },
        },
      ],
      conference: [
        {
          id: 22,
          image: '/publications/010.png',
          title: "Early diagnosis of Alzheimer's disease based on multimodal hypergraph attention network",
          authors: 'Yi Li, Baoyao Yang*, Dan Pan, An Zeng, Yang Yang',
          venue: 'International Conference on Multimedia and Expo (ICME), 2023',
          links: {
            paper: 'https://ieeexplore.ieee.org/abstract/document/10219893',
          },
        },
      ],
    },
    2022: {
      journal: [
        {
          id: 23,
          image: '/publications/011.png',
          title: 'Model-induced Generalization Error Bound for Information-theoretic Representation Learning in Source-data-free Unsupervised Domain Adaptation',
          authors: 'Baoyao Yang, Hao-wei Yeh, Tatsuya Harada, and Pong C. Yuen*',
          venue: 'IEEE Transactions on Image Processing (TIP), 2022',
          links: {
            paper: 'https://ieeexplore.ieee.org/abstract/document/9640468',
          },
        },
        {
          id: 24,
          image: '/publications/012.png',
          title: 'Revealing Task-relevant Model Memorization for Source-Protected Unsupervised Domain Adaptation',
          authors: 'Baoyao Yang and Pong C. Yuen*',
          venue: 'IEEE Transactions on Information Forensics and Security (TIFS), 2022',
          links: {
            paper: 'https://ieeexplore.ieee.org/abstract/document/9705526',
          },
        },
        {
          id: 25,
          image: '/publications/013.png',
          title: 'Cross-domain Missingness-aware Time Series Adaptation with Similarity Distillation in Medical Applications',
          authors: 'Baoyao Yang, Mang Ye, Qingxiong Tan and Pong C. Yuen*',
          venue: 'IEEE Transactions on Cybernetics (TCYB), 2022',
          links: {
            paper: 'https://ieeexplore.ieee.org/abstract/document/9167415',
          },
        },
      ],
      conference: [],
    },
    2021: {
      journal: [
        {
          id: 26,
          image: '/publications/015.png',
          title: 'Learning Adaptive Geometry for Unsupervised Domain Adaptation',
          authors: 'Baoyao Yang and Pong C. Yuen*',
          venue: 'Pattern Recognition (PR), 2021',
          links: {
            paper: 'https://www.sciencedirect.com/science/article/abs/pii/S0031320320304416',
          },
        },
      ],
      conference: [
        {
          id: 27,
          image: '/publications/016.png',
          title: 'A Segmentation-Assisted Model for Universal Lesion Detection with Partial Labels',
          authors: 'Fei Lyu, Baoyao Yang, Andy J. Ma and Pong C. Yuen*',
          venue: 'International Conference on Medical Image Computing & Computer Assisted Intervention (MICCAI), 2021',
          links: {
            paper: 'https://link.springer.com/chapter/10.1007/978-3-030-87240-3_12',
          },
        },
      ],
    },
    2019: {
      journal: [
        {
          id: 28,
          image: '/publications/017.png',
          title: 'Body Parts Synthesis for Cross-Quality Pose Estimation',
          authors: 'Baoyao Yang, Andy J. Ma and Pong C. Yuen*',
          venue: 'IEEE Transactions on Circuits and Systems for Video Technology (TCSVT), 2019',
          links: {
            paper: 'https://ieeexplore.ieee.org/document/8245864',
          },
        },
      ],
      conference: [
        {
          id: 29,
          image: '/publications/018.png',
          title: 'Cross-Domain Visual Representations via Unsupervised Graph Alignment',
          authors: 'Baoyao Yang and Pong C. Yuen*',
          venue: 'the 33rd AAAI Conference on Artificial Intelligence (AAAI), 2019',
          links: {
            paper: 'https://ojs.aaai.org/index.php/AAAI/article/view/4504',
          },
        },
      ],
    },
    2018: {
      journal: [
        {
          id: 30,
          image: '/publications/019.png',
          title: 'Learning Domain-Shared Group-Sparse Representation for Unsupervised Domain Adaptation',
          authors: 'Baoyao Yang, Andy J. Ma and Pong C. Yuen*',
          venue: 'Pattern Recognition (PR), 2018',
          links: {
            paper: 'https://www.sciencedirect.com/science/article/pii/S0031320318301614',
          },
        },
      ],
      conference: [
        {
          id: 31,
          image: '/publications/020.png',
          title: 'Domain-shared Group-sparse Dictionary Learning for Unsupervised Domain Adaptation',
          authors: 'Baoyao Yang, Andy J. Ma and Pong C. Yuen*',
          venue: 'the 32nd AAAI Conference on Artificial Intelligence (AAAI), 2028',
          links: {
            paper: 'https://ojs.aaai.org/index.php/AAAI/article/view/12227',
          },
        },
      ],
    },
  }
}
