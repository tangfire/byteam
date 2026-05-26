import type { NewsItem, Patent, Person, ResearchProject, UndergraduateEducation } from '../../api/client'

export type NewsActivity = {
  content: string
  timestamp?: string
  color?: string
  type?: 'publication' | 'team' | 'award' | 'event' | 'general'
}

export const fallbackHomeNews: NewsItem[] = [
  {
    type: 'publication',
    typeLabel: 'Publication',
    date: '2026-03-17',
    title: 'AdaGS-Net Paper Accepted by ICME 2026',
    content: 'Our paper on adaptive sparse network for multimodal fusion in Alzheimer\'s Disease has been accepted by ICME 2026.',
    excerpt: 'Our paper on adaptive sparse network for multimodal fusion in Alzheimer\'s Disease has been accepted by ICME 2026.',
    color: '#7d1231',
    status: 'published',
    sortOrder: 1,
  },
  {
    type: 'publication',
    typeLabel: 'Publication',
    date: '2025-11-06',
    title: 'Multi-modal Brain Tumor Segmentation Paper Accepted',
    content: 'Our paper on brain tumor segmentation has been accepted by IEEE Transactions on Multimedia.',
    excerpt: 'Our paper on brain tumor segmentation has been accepted by IEEE Transactions on Multimedia.',
    color: '#7d1231',
    status: 'published',
    sortOrder: 2,
  },
  {
    type: 'publication',
    typeLabel: 'Publication',
    date: '2025-10-22',
    title: 'Federated Learning Framework Accepted by PRCV',
    content: 'FedCD framework for adaptive training under data heterogeneity accepted by PRCV conference.',
    excerpt: 'FedCD framework for adaptive training under data heterogeneity accepted by PRCV conference.',
    color: '#7d1231',
    status: 'published',
    sortOrder: 3,
  },
  {
    type: 'team',
    typeLabel: 'Team Update',
    date: '2025-05-06',
    title: 'Welcome New Group Members',
    content: 'Warm welcome to 7 new group members joining our research team this semester.',
    excerpt: 'Warm welcome to 7 new group members joining our research team this semester.',
    color: '#7d1231',
    status: 'published',
    sortOrder: 4,
  },
]

export const fallbackNewsActivities: NewsActivity[] = [
  {
    content: 'Our paper "AdaGS-Net: An Adaptive Sparse Network for Multimodal Fusion in Alzheimer\'s Disease" has been accepted by ICME 2026. Congrats to all authors!',
    timestamp: '2026-03-17',
    color: '#7d1231',
    type: 'publication',
  },
  {
    content: 'Our paper "Multi-modal Brain Tumor Segmentation via Multi-category Interaction and Graph Co-reasoning" has been accepted by IEEE Transactions on Multimedia. Congrats, Li, Yang and co-authors!',
    timestamp: '2025-11-06',
    color: '#7d1231',
    type: 'publication',
  },
  {
    content: 'Our paper "FedCD: A Hybrid Federated Learning Framework for Adaptive Training under Data Heterogeneity" has been accepted by PRCV. Congrats, Zhan, Yang and co-authors!',
    timestamp: '2025-10-22',
    color: '#7d1231',
    type: 'publication',
  },
  {
    content: 'Our paper "CAM-interacted Vision GNN for Multi-label Medical Images" to be published in IEEE Journal of Biomedical and Health Informatics. Congrats, Wang, Yang and co-authors!',
    timestamp: '2025-10-16',
    color: '#7d1231',
    type: 'publication',
  },
  {
    content: 'Our paper "Multi-Agent Reinforcement Learning Algorithm Using Dynamic OW-QMIX in Complex Supply Chain Scenarios" has been accepted by IEEE SMC. Congrats, Liu, Zhu and co-authors!',
    timestamp: '2025-07-20',
    color: '#7d1231',
    type: 'publication',
  },
  {
    content: 'Two of our papers were accepted by ECAI: "FairFed++: Closing the Fairness Gap in Federated Learning through Self-Evolving Clustered Optimization" and "Unlocking the Potential of mLLMs: Enhancing Video-Text Retrieval through Caption Supplementation and Conical Embedding Optimization".',
    timestamp: '2025-07-11',
    color: '#7d1231',
    type: 'publication',
  },
  {
    content: 'Our paper "Simple but Effective: Sub-Volume Contrastive Learning for Class-Imbalanced Semi-Supervised 3D Medical Image Segmentation" has been accepted by ACM Multimedia. Congrats, Xu, Yang and co-authors!',
    timestamp: '2025-07-06',
    color: '#7d1231',
    type: 'publication',
  },
  {
    content: 'Our paper "Harnessing Feature Distribution Consistency for Federated Learning with Noisy" has been accepted by IEEE International Conference on Image Processing (ICIP). Congrats, Ma, Yang and co-authors!',
    timestamp: '2025-05-20',
    color: '#7d1231',
    type: 'publication',
  },
  {
    content: 'Warm welcome to new group members Aoqi Yan, Guangyang Lin, Haifeng Lin, Jiahao Lian, Xiaojie Chen, Xi Wang & Yuhao Chen',
    timestamp: '2025-05-06',
    color: '#7d1231',
    type: 'team',
  },
  {
    content: 'Beyond Machine Learning team introduction official website officially launched!',
    timestamp: '2025-04-08',
    color: '#7d1231',
    type: 'event',
  },
  {
    content: 'Two of our papers were accepted by ICME: "Unifying Spatio-Temporal Contexts for Advanced Text-Video Retrieval" and "Action Decomposition-based Actor-Critic for Supply Chain Optimization".',
    timestamp: '2025-03-21',
    color: '#7d1231',
    type: 'publication',
  },
  {
    content: 'Our paper "Image-assisted Label Connective Completion for Vessel Segmentation with Insufficient Annotations" has been accepted by ICASSP. Congrats, Zheng, Yang and co-authors!',
    timestamp: '2025-03-08',
    color: '#7d1231',
    type: 'publication',
  },
  {
    content: 'Happy Chinese New Year!',
    timestamp: '2025-01-29',
    color: '#7d1231',
    type: 'event',
  },
]

export const fallbackGraduatePeople: Person[] = [
  { name: 'Dixin Chen', avatarUrl: '/avatar/DixinChen.jpg', category: 'graduate', research: '', graduationDate: '', status: 'published', sortOrder: 1 },
  { name: 'Huahong Deng', avatarUrl: '/avatar/HuahongDeng.jpg', category: 'graduate', research: '', graduationDate: '', status: 'published', sortOrder: 2 },
  { name: 'Canrong Du', avatarUrl: '/avatar/CanrongDu.jpg', category: 'graduate', research: '', graduationDate: '', status: 'published', sortOrder: 3 },
  { name: 'Aoqi Yan', avatarUrl: '/avatar/AoqiYan.jpg', category: 'graduate', research: '', graduationDate: '', status: 'published', sortOrder: 4 },
  { name: 'Haifeng Lin', avatarUrl: '/avatar/HaifengLin.jpg', category: 'graduate', research: '', graduationDate: '', status: 'published', sortOrder: 5 },
  { name: 'Yanhao Huang', avatarUrl: '/avatar/YanhaoHuang.jpg', category: 'graduate', research: '', graduationDate: '', status: 'published', sortOrder: 6 },
  { name: 'Jingsong Lin', avatarUrl: '/avatar/JingsongLin.jpg', category: 'graduate', research: '', graduationDate: '', status: 'published', sortOrder: 7 },
  { name: 'Yali Ma', avatarUrl: '/avatar/YaliMa.jpg', category: 'graduate', research: '', graduationDate: '', status: 'published', sortOrder: 8 },
  { name: 'Yanchao Tang', avatarUrl: '/avatar/YanchaoTang.jpg', category: 'graduate', research: '', graduationDate: '', status: 'published', sortOrder: 9 },
  { name: 'Kexin Xie', avatarUrl: '/avatar/KexinXie.jpg', category: 'graduate', research: '', graduationDate: '', status: 'published', sortOrder: 10 },
  { name: 'Yuhao Chen', avatarUrl: '/avatar/YuhaoChen.jpg', category: 'graduate', research: '', graduationDate: '', status: 'published', sortOrder: 11 },
  { name: 'Yufei Xu', avatarUrl: '/avatar/YufeiXu.jpg', category: 'graduate', research: '', graduationDate: '', status: 'published', sortOrder: 12 },
  { name: 'Xianrun Xu', avatarUrl: '/avatar/XianrunXu.jpg', category: 'graduate', research: '', graduationDate: '', status: 'published', sortOrder: 13 },
  { name: 'Xiaoqi Zheng', avatarUrl: '/avatar/XiaoqiZheng.jpg', category: 'graduate', research: '', graduationDate: '', status: 'published', sortOrder: 14 },
  { name: 'Sijia Zhou', avatarUrl: '/avatar/SijiaZhou.jpg', category: 'graduate', research: '', graduationDate: '', status: 'published', sortOrder: 15 },
  { name: 'Xiaojie Chen', avatarUrl: '/avatar/XiaojieChen.jpg', category: 'graduate', research: '', graduationDate: '', status: 'published', sortOrder: 16 },
  { name: 'Xi Wang', avatarUrl: '/avatar/XiWang.jpg', category: 'graduate', research: '', graduationDate: '', status: 'published', sortOrder: 17 },
  { name: 'Jiahao Lian', avatarUrl: '/avatar/JiahaoLian.jpg', category: 'graduate', research: '', graduationDate: '', status: 'published', sortOrder: 18 },
  { name: 'Guangyang Lin', avatarUrl: '/avatar/GuangyangLin.jpg', category: 'graduate', research: '', graduationDate: '', status: 'published', sortOrder: 19 },
]

export const fallbackAlumni: Person[] = [
  { name: 'Yuehui Fan', avatarUrl: '/avatar/YuehuiFan.jpg', category: 'graduate_alumni', research: '', graduationDate: 'Graduated: 2025.07', status: 'published', sortOrder: 1 },
  { name: 'Jianxuan Huang', avatarUrl: '/avatar/JianxuanHuang.jpg', category: 'graduate_alumni', research: '', graduationDate: 'Graduated: 2025.07', status: 'published', sortOrder: 2 },
  { name: 'Yuebin Xie', avatarUrl: '/avatar/YuebinXie.jpg', category: 'graduate_alumni', research: '', graduationDate: 'Graduated: 2025.07', status: 'published', sortOrder: 3 },
  { name: 'Jingchao Wang', avatarUrl: '/avatar/JingchaoWang.jpg', category: 'undergraduate_alumni', research: '', graduationDate: 'Graduated: 2025.07', status: 'published', sortOrder: 4 },
  { name: 'Dongzhe Li', avatarUrl: '/avatar/DongzheLi.jpg', category: 'undergraduate_alumni', research: '', graduationDate: 'Graduated: 2025.07', status: 'published', sortOrder: 5 },
  { name: 'Xiaochen He', avatarUrl: '/avatar/XiaochenHe.jpg', category: 'undergraduate_alumni', research: '', graduationDate: 'Graduated: 2025.07', status: 'published', sortOrder: 6 },
  { name: 'Weide Zhan', avatarUrl: '/avatar/WeideZhan.jpg', category: 'undergraduate_alumni', research: '', graduationDate: 'Graduated: 2025.07', status: 'published', sortOrder: 7 },
  { name: 'Zhixiang Fang', avatarUrl: '/avatar/ZhixiangFang.jpg', category: 'undergraduate_alumni', research: '', graduationDate: 'Graduated: 2025.07', status: 'published', sortOrder: 8 },
]

export const fallbackUndergraduates: UndergraduateEducation[] = [
  {
    id: 1,
    name: '穆跃鑫',
    grade: '19 级',
    major: '机器人学院',
    direction: '机器学习与数据挖掘',
    achievements: ['申请国家发明专利 1 项', '获优秀毕业设计', '保送至重庆大学'],
    status: 'published',
    sortOrder: 1,
  },
  {
    id: 2,
    name: '何晓琛',
    grade: '20 级',
    major: '软件工程（卓越班）',
    direction: '医学图像处理与单域泛化',
    achievements: ['发表 CCF-B 类会议 ICASSP2024 论文一篇'],
    status: 'published',
    sortOrder: 2,
  },
  {
    id: 3,
    name: '王敬超',
    grade: '21 级',
    major: '人工智能专业',
    direction: '机器学习与数据挖掘',
    achievements: [
      '以第一作者在 CCF-A 类会议 SIGKDD 2024 发表论文一篇',
      '以第一作者在 CCF-A 类会议 ACMMM 2024 发表论文一篇',
      '发表 JBHI 论文一篇',
      '荣获科技创新奖学金',
      '保送至北京大学计算机科学与技术专业（博士）',
    ],
    status: 'published',
    sortOrder: 3,
  },
  {
    id: 4,
    name: '詹伟德',
    grade: '21 级',
    major: '计算机科学与技术（伏羲班）',
    direction: '联邦学习',
    achievements: [
      '授权国家发明专利 1 项，实审 2 项',
      '获省级赛事奖项 2 项',
      '参与发表 CCF-B 类会议 MICCAI 2024 论文一篇',
      '连续两年获得校一等奖学金',
      '国家优秀学生奖学金',
      '保送至复旦大学电子信息专业（博士）',
    ],
    status: 'published',
    sortOrder: 4,
  },
  {
    id: 5,
    name: '李东哲',
    grade: '21 级',
    major: '计算机科学与技术（伏羲班）',
    direction: '医学图像分割技术',
    achievements: ['发表 TMM（CCF-B）论文一篇', '发表医学信息处理顶会 MICCAI 2024（CCF-B）论文一篇', '发表 IJCAI-HBAI 论文一篇'],
    status: 'published',
    sortOrder: 5,
  },
  {
    id: 6,
    name: '方志祥',
    grade: '21 级',
    major: '计算机科学与技术',
    direction: '联邦学习公平性问题',
    achievements: ['发表 ECAI 2025（CCF-B）论文一篇'],
    status: 'published',
    sortOrder: 6,
  },
]

export const fallbackPatents: Patent[] = [
  { authors: '杨宝瑶，黄彦浩，陈涤新', title: '一种基于时空信息聚合的视频特征提取模型训练方法、系统及特征提取方法', date: '2025-09-09', country: '中国', number: 'ZL202510359255.5', category: 'granted', status: 'published', sortOrder: 1 },
  { authors: '杨宝瑶，麻亚利，詹伟德，唐彦超，卢泽坚', title: '一种联邦学习场景下的检测噪声标注的方法及系统', date: '2025-09-05', country: '中国', number: 'ZL202510375049.3', category: 'granted', status: 'published', sortOrder: 2 },
  { authors: '杨宝瑶，陈俊祥，黄彦浩，姚文彬', title: '一种基于多模态大模型的视频 - 文本检索方法', date: '2025-06-13', country: '中国', number: 'ZL202411271756.X', category: 'granted', status: 'published', sortOrder: 3 },
  { authors: '杨宝瑶，郑晓琦', title: '一种噪声标注的血管图像分割方法及系统', date: '2025-02-11', country: '中国', number: 'ZL 202311792631.7', category: 'granted', status: 'published', sortOrder: 4 },
  { authors: '杨宝瑶，詹伟德', title: '一种模型异构性联邦学习方法和系统', date: '', country: '中国', number: 'ZL202210989290.1', category: 'review', status: 'published', sortOrder: 14 },
  { authors: '李春林，吴恒，曾安，惠恩明，富锐，杨思维，曹洪江，骆有隆，杨宝瑶，刘俊，张勇，江焜', title: '产业聚集区域内业务资源服务平台规范', date: '2025-05-01', country: '广东省工业软件学会团体标准', number: 'T/GISF 002-2024', category: 'standard', status: 'published', sortOrder: 30 },
]

export const fallbackResearchProjects: ResearchProject[] = [
  { id: 1, title: '弹性医学联邦学习与攻击防御关键技术研究', fund: '国家自然科学基金，面上项目', number: '62472105', period: '2025-01 至 2028-12', amount: '50 万元', projectStatus: '在研', role: '主持', status: 'published', sortOrder: 1 },
  { id: 2, title: '基于多模态时间序列表征学习的肝癌早期风险预测算法研究', fund: '国家自然科学基金，青年科学基金项目', number: '62102098', period: '2022-01 至 2024-12', amount: '30 万元', projectStatus: '在研', role: '主持', status: 'published', sortOrder: 2 },
  { id: 3, title: '产业聚集区域业务资源服务工业软件平台', fund: '国家重点研发计划，工业软件专项', number: '2023YFB3308700', period: '2023-12-01 至 2026-11-31', amount: '3500 万（50 万）', projectStatus: '在研', role: '子课题负责人', status: 'published', sortOrder: 3 },
  { id: 4, title: '非侵入式冠状动脉病变的智能定位与评估关键技术研究', fund: '广东省自然科学基金委，广东省基础与应用基础研究基金面上项目', number: '2024A1515010186', period: '2024-01 至 2026-12', amount: '15 万', projectStatus: '在研', role: '主持', status: 'published', sortOrder: 4 },
  { id: 5, title: '多模态医学大模型设计与医学影像自动标注技术研究', fund: '广东省自然科学基金委，广东省基础与应用基础研究基金面上项目', number: '2025A1515011385', period: '2025-01 至 2027-12', amount: '10 万', projectStatus: '在研', role: '主持', status: 'published', sortOrder: 5 },
  { id: 6, title: '基于可解释性 DNN 的影像辅助归因技术探索直肠癌病程发展规律', fund: '广东省自然科学基金委，粤港澳应用数学中心青年启动项目', number: '2025A1515060016', period: '2025-07 至 2027-06', amount: '5 万', projectStatus: '在研', role: '主持', status: 'published', sortOrder: 6 },
  { id: 7, title: '多源异构热处理数据智能联网及其保护的关键技术研究', fund: '广东省自然科学基金委，广东省基础与应用基础研究基金区域联合基金项目', number: '2022A1515140096', period: '2022-10 至 2025-09', amount: '30 万元', projectStatus: '已结题', role: '校内主持', status: 'published', sortOrder: 7 },
  { id: 8, title: '肝活检图像的多类病变细胞弱监督自动检测算法研究', fund: '广州市科技局，广州市基础与应用基础研究项目', number: '202201010266', period: '2022-04 至 2024-03', amount: '5 万元', projectStatus: '已结题', role: '主持', status: 'published', sortOrder: 8 },
  { id: 9, title: '面向工业物联网的数据保护与协同学习关键技术研究', fund: '广东工业大学，交叉学科培育项目', number: '', period: '2022-07 至 2024-07', amount: '5 万元', projectStatus: '已结题', role: '主持', status: 'published', sortOrder: 9 },
  { id: 10, title: "合成生物基因数据库的建立与查询工具的研发", fund: "广东省科技厅，重点领域研发计划'绿色生物制造'重点专项", number: '2024B1111140001', period: '2024-01 至 2026-12', amount: '1000 万元', projectStatus: '在研', role: '参与', status: 'published', sortOrder: 10 },
  { id: 11, title: '辅助发现阿尔茨海默症大脑神经退化模式的深度学习模型及可解释性研究', fund: '国家自然科学基金，面上项目', number: '62572131', period: '2026-01 至 2029-12', amount: '50 万元', projectStatus: '在研', role: '参与', status: 'published', sortOrder: 11 },
  { id: 12, title: '基于复杂异构特征医学数据挖掘的脓毒症智能预测方法研究', fund: '广东省自然科学基金委，广东省基础与应用基础研究面上项目', number: '2022A1515011592', period: '2022-01 至 2024-12', amount: '10 万元', projectStatus: '在研', role: '参与', status: 'published', sortOrder: 12 },
]
