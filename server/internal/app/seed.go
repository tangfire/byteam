package app

import "gorm.io/gorm/clause"

func (s *Server) seedContent() error {
	if err := s.seedNews(); err != nil {
		return err
	}
	if err := s.seedPeople(); err != nil {
		return err
	}
	if err := s.seedUndergraduates(); err != nil {
		return err
	}
	if err := s.seedResearchProjects(); err != nil {
		return err
	}
	if err := s.seedPublications(); err != nil {
		return err
	}
	return s.seedPatents()
}

func (s *Server) seedNews() error {
	items := []NewsItem{
		{Title: "AdaGS-Net Paper Accepted by ICME 2026", Content: `Our paper "AdaGS-Net: An Adaptive Sparse Network for Multimodal Fusion in Alzheimer's Disease" has been accepted by ICME 2026. Congrats to all authors!`, Excerpt: `Our paper on adaptive sparse network for multimodal fusion in Alzheimer's Disease has been accepted by ICME 2026.`, Type: "publication", TypeLabel: "Publication", EventDate: "2026-03-17", Color: "#7d1231", Status: StatusPublished, SortOrder: 1},
		{Title: "Multi-modal Brain Tumor Segmentation Paper Accepted", Content: `Our paper "Multi-modal Brain Tumor Segmentation via Multi-category Interaction and Graph Co-reasoning" has been accepted by IEEE Transactions on Multimedia. Congrats, Li, Yang and co-authors!`, Excerpt: "Our paper on brain tumor segmentation has been accepted by IEEE Transactions on Multimedia.", Type: "publication", TypeLabel: "Publication", EventDate: "2025-11-06", Color: "#7d1231", Status: StatusPublished, SortOrder: 2},
		{Title: "Federated Learning Framework Accepted by PRCV", Content: `Our paper "FedCD: A Hybrid Federated Learning Framework for Adaptive Training under Data Heterogeneity" has been accepted by PRCV. Congrats, Zhan, Yang and co-authors!`, Excerpt: "FedCD framework for adaptive training under data heterogeneity accepted by PRCV conference.", Type: "publication", TypeLabel: "Publication", EventDate: "2025-10-22", Color: "#7d1231", Status: StatusPublished, SortOrder: 3},
		{Title: "CAM-interacted Vision GNN Paper Accepted", Content: `Our paper "CAM-interacted Vision GNN for Multi-label Medical Images" to be published in IEEE Journal of Biomedical and Health Informatics. Congrats, Wang, Yang and co-authors!`, Excerpt: "A multi-label medical image paper will be published in IEEE JBHI.", Type: "publication", TypeLabel: "Publication", EventDate: "2025-10-16", Color: "#7d1231", Status: StatusPublished, SortOrder: 4},
		{Title: "Multi-Agent Reinforcement Learning Paper Accepted", Content: `Our paper "Multi-Agent Reinforcement Learning Algorithm Using Dynamic OW-QMIX in Complex Supply Chain Scenarios" has been accepted by IEEE SMC. Congrats, Liu, Zhu and co-authors!`, Excerpt: "A supply chain optimization paper has been accepted by IEEE SMC.", Type: "publication", TypeLabel: "Publication", EventDate: "2025-07-20", Color: "#7d1231", Status: StatusPublished, SortOrder: 5},
		{Title: "Two Papers Accepted by ECAI", Content: `Two of our papers were accepted by ECAI: "FairFed++: Closing the Fairness Gap in Federated Learning through Self-Evolving Clustered Optimization" and "Unlocking the Potential of mLLMs: Enhancing Video-Text Retrieval through Caption Supplementation and Conical Embedding Optimization".`, Excerpt: "Two papers were accepted by ECAI.", Type: "publication", TypeLabel: "Publication", EventDate: "2025-07-11", Color: "#7d1231", Status: StatusPublished, SortOrder: 6},
		{Title: "Sub-Volume Contrastive Learning Paper Accepted", Content: `Our paper "Simple but Effective: Sub-Volume Contrastive Learning for Class-Imbalanced Semi-Supervised 3D Medical Image Segmentation" has been accepted by ACM Multimedia. Congrats, Xu, Yang and co-authors!`, Excerpt: "A 3D medical image segmentation paper has been accepted by ACM Multimedia.", Type: "publication", TypeLabel: "Publication", EventDate: "2025-07-06", Color: "#7d1231", Status: StatusPublished, SortOrder: 7},
		{Title: "Federated Learning with Noisy Paper Accepted", Content: `Our paper "Harnessing Feature Distribution Consistency for Federated Learning with Noisy" has been accepted by IEEE International Conference on Image Processing (ICIP). Congrats, Ma, Yang and co-authors!`, Excerpt: "A federated learning paper has been accepted by ICIP.", Type: "publication", TypeLabel: "Publication", EventDate: "2025-05-20", Color: "#7d1231", Status: StatusPublished, SortOrder: 8},
		{Title: "Welcome New Group Members", Content: "Warm welcome to new group members Aoqi Yan, Guangyang Lin, Haifeng Lin, Jiahao Lian, Xiaojie Chen, Xi Wang & Yuhao Chen", Excerpt: "Warm welcome to 7 new group members joining our research team this semester.", Type: "team", TypeLabel: "Team Update", EventDate: "2025-05-06", Color: "#7d1231", Status: StatusPublished, SortOrder: 9},
		{Title: "Official Website Launched", Content: "Beyond Machine Learning team introduction official website officially launched!", Excerpt: "The BYML website officially launched.", Type: "event", TypeLabel: "Event", EventDate: "2025-04-08", Color: "#7d1231", Status: StatusPublished, SortOrder: 10},
		{Title: "Two Papers Accepted by ICME", Content: `Two of our papers were accepted by ICME: "Unifying Spatio-Temporal Contexts for Advanced Text-Video Retrieval" and "Action Decomposition-based Actor-Critic for Supply Chain Optimization".`, Excerpt: "Two papers were accepted by ICME.", Type: "publication", TypeLabel: "Publication", EventDate: "2025-03-21", Color: "#7d1231", Status: StatusPublished, SortOrder: 11},
		{Title: "Vessel Segmentation Paper Accepted by ICASSP", Content: `Our paper "Image-assisted Label Connective Completion for Vessel Segmentation with Insufficient Annotations" has been accepted by ICASSP. Congrats, Zheng, Yang and co-authors!`, Excerpt: "A vessel segmentation paper has been accepted by ICASSP.", Type: "publication", TypeLabel: "Publication", EventDate: "2025-03-08", Color: "#7d1231", Status: StatusPublished, SortOrder: 12},
		{Title: "Happy Chinese New Year", Content: "Happy Chinese New Year!", Excerpt: "Happy Chinese New Year!", Type: "event", TypeLabel: "Event", EventDate: "2025-01-29", Color: "#7d1231", Status: StatusPublished, SortOrder: 13},
	}
	return s.db.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "title"}}, DoNothing: true}).Create(&items).Error
}

func (s *Server) seedPeople() error {
	names := []struct {
		Name      string
		AvatarURL string
		Category  string
		Graduate  string
	}{
		{"Dixin Chen", "/avatar/DixinChen.jpg", "graduate", ""},
		{"Huahong Deng", "/avatar/HuahongDeng.jpg", "graduate", ""},
		{"Canrong Du", "/avatar/CanrongDu.jpg", "graduate", ""},
		{"Aoqi Yan", "/avatar/AoqiYan.jpg", "graduate", ""},
		{"Haifeng Lin", "/avatar/HaifengLin.jpg", "graduate", ""},
		{"Yanhao Huang", "/avatar/YanhaoHuang.jpg", "graduate", ""},
		{"Jingsong Lin", "/avatar/JingsongLin.jpg", "graduate", ""},
		{"Yali Ma", "/avatar/YaliMa.jpg", "graduate", ""},
		{"Yanchao Tang", "/avatar/YanchaoTang.jpg", "graduate", ""},
		{"Kexin Xie", "/avatar/KexinXie.jpg", "graduate", ""},
		{"Yuhao Chen", "/avatar/YuhaoChen.jpg", "graduate", ""},
		{"Yufei Xu", "/avatar/YufeiXu.jpg", "graduate", ""},
		{"Xianrun Xu", "/avatar/XianrunXu.jpg", "graduate", ""},
		{"Xiaoqi Zheng", "/avatar/XiaoqiZheng.jpg", "graduate", ""},
		{"Sijia Zhou", "/avatar/SijiaZhou.jpg", "graduate", ""},
		{"Xiaojie Chen", "/avatar/XiaojieChen.jpg", "graduate", ""},
		{"Xi Wang", "/avatar/XiWang.jpg", "graduate", ""},
		{"Jiahao Lian", "/avatar/JiahaoLian.jpg", "graduate", ""},
		{"Guangyang Lin", "/avatar/GuangyangLin.jpg", "graduate", ""},
		{"Yuehui Fan", "/avatar/YuehuiFan.jpg", "graduate_alumni", "Graduated: 2025.07"},
		{"Jianxuan Huang", "/avatar/JianxuanHuang.jpg", "graduate_alumni", "Graduated: 2025.07"},
		{"Yuebin Xie", "/avatar/YuebinXie.jpg", "graduate_alumni", "Graduated: 2025.07"},
		{"Jingchao Wang", "/avatar/JingchaoWang.jpg", "undergraduate_alumni", "Graduated: 2025.07"},
		{"Dongzhe Li", "/avatar/DongzheLi.jpg", "undergraduate_alumni", "Graduated: 2025.07"},
		{"Xiaochen He", "/avatar/XiaochenHe.jpg", "undergraduate_alumni", "Graduated: 2025.07"},
		{"Weide Zhan", "/avatar/WeideZhan.jpg", "undergraduate_alumni", "Graduated: 2025.07"},
		{"Zhixiang Fang", "/avatar/ZhixiangFang.jpg", "undergraduate_alumni", "Graduated: 2025.07"},
	}
	items := make([]Person, 0, len(names))
	for i, item := range names {
		items = append(items, Person{Name: item.Name, AvatarURL: item.AvatarURL, Category: item.Category, GraduationDate: item.Graduate, Status: StatusPublished, SortOrder: i + 1})
	}
	return s.db.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "name"}}, DoNothing: true}).Create(&items).Error
}

func (s *Server) seedUndergraduates() error {
	items := []UndergraduateEducation{
		{Name: "穆跃鑫", Grade: "19 级", Major: "机器人学院", Direction: "机器学习与数据挖掘", Achievements: StringList{"申请国家发明专利 1 项", "获优秀毕业设计", "保送至重庆大学"}, Status: StatusPublished, SortOrder: 1},
		{Name: "何晓琛", Grade: "20 级", Major: "软件工程（卓越班）", Direction: "医学图像处理与单域泛化", Achievements: StringList{"发表 CCF-B 类会议 ICASSP2024 论文一篇"}, Status: StatusPublished, SortOrder: 2},
		{Name: "王敬超", Grade: "21 级", Major: "人工智能专业", Direction: "机器学习与数据挖掘", Achievements: StringList{"以第一作者在 CCF-A 类会议 SIGKDD 2024 发表论文一篇", "以第一作者在 CCF-A 类会议 ACMMM 2024 发表论文一篇", "发表 JBHI 论文一篇", "荣获科技创新奖学金", "保送至北京大学计算机科学与技术专业（博士）"}, Status: StatusPublished, SortOrder: 3},
		{Name: "詹伟德", Grade: "21 级", Major: "计算机科学与技术（伏羲班）", Direction: "联邦学习", Achievements: StringList{"授权国家发明专利 1 项，实审 2 项", "获省级赛事奖项 2 项", "参与发表 CCF-B 类会议 MICCAI 2024 论文一篇", "连续两年获得校一等奖学金", "国家优秀学生奖学金", "保送至复旦大学电子信息专业（博士）"}, Status: StatusPublished, SortOrder: 4},
		{Name: "李东哲", Grade: "21 级", Major: "计算机科学与技术（伏羲班）", Direction: "医学图像分割技术", Achievements: StringList{"发表 TMM（CCF-B）论文一篇", "发表医学信息处理顶会 MICCAI 2024（CCF-B）论文一篇", "发表 IJCAI-HBAI 论文一篇"}, Status: StatusPublished, SortOrder: 5},
		{Name: "方志祥", Grade: "21 级", Major: "计算机科学与技术", Direction: "联邦学习公平性问题", Achievements: StringList{"发表 ECAI 2025（CCF-B）论文一篇"}, Status: StatusPublished, SortOrder: 6},
	}
	return s.db.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "name"}}, DoNothing: true}).Create(&items).Error
}

func (s *Server) seedResearchProjects() error {
	items := []ResearchProject{
		{Title: "弹性医学联邦学习与攻击防御关键技术研究", Fund: "国家自然科学基金，面上项目", Number: "62472105", Period: "2025-01 至 2028-12", Amount: "50 万元", ProjectStatus: "在研", Role: "主持", Status: StatusPublished, SortOrder: 1},
		{Title: "基于多模态时间序列表征学习的肝癌早期风险预测算法研究", Fund: "国家自然科学基金，青年科学基金项目", Number: "62102098", Period: "2022-01 至 2024-12", Amount: "30 万元", ProjectStatus: "在研", Role: "主持", Status: StatusPublished, SortOrder: 2},
		{Title: "产业聚集区域业务资源服务工业软件平台", Fund: "国家重点研发计划，工业软件专项", Number: "2023YFB3308700", Period: "2023-12-01 至 2026-11-31", Amount: "3500 万（50 万）", ProjectStatus: "在研", Role: "子课题负责人", Status: StatusPublished, SortOrder: 3},
		{Title: "非侵入式冠状动脉病变的智能定位与评估关键技术研究", Fund: "广东省自然科学基金委，广东省基础与应用基础研究基金面上项目", Number: "2024A1515010186", Period: "2024-01 至 2026-12", Amount: "15 万", ProjectStatus: "在研", Role: "主持", Status: StatusPublished, SortOrder: 4},
		{Title: "多模态医学大模型设计与医学影像自动标注技术研究", Fund: "广东省自然科学基金委，广东省基础与应用基础研究基金面上项目", Number: "2025A1515011385", Period: "2025-01 至 2027-12", Amount: "10 万", ProjectStatus: "在研", Role: "主持", Status: StatusPublished, SortOrder: 5},
		{Title: "基于可解释性 DNN 的影像辅助归因技术探索直肠癌病程发展规律", Fund: "广东省自然科学基金委，粤港澳应用数学中心青年启动项目", Number: "2025A1515060016", Period: "2025-07 至 2027-06", Amount: "5 万", ProjectStatus: "在研", Role: "主持", Status: StatusPublished, SortOrder: 6},
		{Title: "多源异构热处理数据智能联网及其保护的关键技术研究", Fund: "广东省自然科学基金委，广东省基础与应用基础研究基金区域联合基金项目", Number: "2022A1515140096", Period: "2022-10 至 2025-09", Amount: "30 万元", ProjectStatus: "已结题", Role: "校内主持", Status: StatusPublished, SortOrder: 7},
		{Title: "肝活检图像的多类病变细胞弱监督自动检测算法研究", Fund: "广州市科技局，广州市基础与应用基础研究项目", Number: "202201010266", Period: "2022-04 至 2024-03", Amount: "5 万元", ProjectStatus: "已结题", Role: "主持", Status: StatusPublished, SortOrder: 8},
		{Title: "面向工业物联网的数据保护与协同学习关键技术研究", Fund: "广东工业大学，交叉学科培育项目", Period: "2022-07 至 2024-07", Amount: "5 万元", ProjectStatus: "已结题", Role: "主持", Status: StatusPublished, SortOrder: 9},
		{Title: "合成生物基因数据库的建立与查询工具的研发", Fund: "广东省科技厅，重点领域研发计划'绿色生物制造'重点专项", Number: "2024B1111140001", Period: "2024-01 至 2026-12", Amount: "1000 万元", ProjectStatus: "在研", Role: "参与", Status: StatusPublished, SortOrder: 10},
		{Title: "辅助发现阿尔茨海默症大脑神经退化模式的深度学习模型及可解释性研究", Fund: "国家自然科学基金，面上项目", Number: "62572131", Period: "2026-01 至 2029-12", Amount: "50 万元", ProjectStatus: "在研", Role: "参与", Status: StatusPublished, SortOrder: 11},
		{Title: "基于复杂异构特征医学数据挖掘的脓毒症智能预测方法研究", Fund: "广东省自然科学基金委，广东省基础与应用基础研究面上项目", Number: "2022A1515011592", Period: "2022-01 至 2024-12", Amount: "10 万元", ProjectStatus: "在研", Role: "参与", Status: StatusPublished, SortOrder: 12},
	}
	return s.db.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "title"}}, DoNothing: true}).Create(&items).Error
}

func (s *Server) seedPublications() error {
	items := []Publication{
		pub(1, "/publications/028.png", "Multi-modal Brain Tumor Segmentation via Multi-category Interaction and Graph Co-reasoning", "Baoyao Yang*, Dongzhe Li, Chong Yin, Fei Lyu, Xiaochen He", "IEEE Transactions on Multimedia, 2026", 2026, "journal", true, []PublicationLink{link("paper", "Paper", "https://link.springer.com/chapter/10.1007/978-981-96-4001-0_19", "")}),
		pub(2, "/publications/026.png", "CAM-interacted Vision GNN for Multi-label Medical Images", "Jiangchao Wang, Baoyao Yang*, Siqi Liu, Xiaoqi Zheng, Wenbin Yao* and Junxiang Chen", "IEEE Journal of Biomedical and Health Informatics (JBHI), 2025", 2025, "journal", true, []PublicationLink{link("code", "Code", "https://github.com/BaoyaoGroup/JBHI_code", ""), link("paper", "Paper", "https://ieeexplore.ieee.org/abstract/document/11205277", "")}),
		pub(3, "/publications/027.png", "FedCD: A Hybrid Federated Learning Framework for Adaptive Training under Data Heterogeneity", "Weide Zhan, Baoyao Yang*", "Chinese Conference on Pattern Recognition and Computer Vision (PRCV), 2025", 2025, "conference", true, []PublicationLink{link("paper", "Paper", "https://ieeexplore.ieee.org/abstract/document/10443215", "")}),
		pub(4, "/publications/004.png", "Image-assisted Label Connective Completion for Vessel Segmentation with Insufficient Annotations", "Xiaoqi Zheng, Baoyao Yang*, Xiuwen Fang, Xiuwen Fang, Mang Ye", "IEEE International Conference on Acoustics, Speech, and Signal Processing (ICASSP), 2025", 2025, "conference", true, []PublicationLink{link("code", "Code", "https://github.com/BaoyaoGroup/LabelCompletion", ""), link("paper", "Paper", "https://ieeexplore.ieee.org/document/10888997", ""), link("video", "Video", "", "video-player-XiaoqiZheng01")}),
		pub(5, "/publications/030.png", "Unifying Spatio-Temporal Contexts for Advanced Text-Video Retrieval", "Yanhao Huang, Baoyao Yang*, Junxiang Chen, Wenbin Yao, Dixin Chen", "IEEE International Conference on Multimedia and Expo (ICME), 2025", 2025, "conference", false, []PublicationLink{link("paper", "Paper", "https://ieeexplore.ieee.org/abstract/document/11209054/", "")}),
		pub(6, "/publications/024.png", "FairFed++: Closing the Fairness Gap in Federated Learning through Self-Evolving Clustered Optimization", "Zhixiang Fang, Baoyao Yang*, Weide Zhan, Yanchao Tang, Yiqun Zhang", "the 28th European Conference on Artificial Intelligence (ECAI), 2025", 2025, "conference", false, []PublicationLink{link("code", "Code", "https://github.com/BaoyaoGroup/FairFedPlusPlus", ""), link("ppt", "PPT", "/files/FairFed.pdf", "")}),
		pub(7, "/publications/029.png", "Unlocking the Potential of mLLMs: Enhancing Video-Text Retrieval through Caption Supplementation and Conical Embedding Optimization", "Baoyao Yang*, Junxiang Chen, Wenbin Yao", "the 28th European Conference on Artificial Intelligence (ECAI), 2025", 2025, "conference", false, []PublicationLink{link("ppt", "PPT", "/files/UnlockingThePotentialOfmLLMs.pdf", "")}),
		pub(8, "/publications/023.png", "Simple but Effective: Sub-Volume Contrastive Learning for Class-Imbalanced Semi-Supervised 3D Medical Image Segmentation", "Xianrun Xu, Baoyao Yang*, Wanyun Li, Jingsong Lin, Yufei Xu", "the 33rd ACM International Conference on Multimedia (ACM MM), 2025", 2025, "conference", true, []PublicationLink{link("paper", "Paper", "https://dl.acm.org/doi/abs/10.1145/3746027.3755652", ""), link("video", "Video", "", "video-player-XianrunXu01")}),
		pub(9, "/publications/031.png", "Multi-Agent Reinforcement Learning Algorithm Using Dynamic OW-QMIX in Complex Supply Chain Scenarios", "ZhiQi Liu, QingHua Zhu*, An Zeng, YuZhu Ji, BaoYao Yang", "IEEE International Conference on Systems, Man, and Cybernetics (IEEE SMC), 2025", 2025, "conference", false, []PublicationLink{link("paper", "Paper", "https://ieeexplore.ieee.org/abstract/document/11342630", "")}),
		pub(10, "/publications/032.png", "Action Decomposition-based Actor-Critic for Supply Chain Optimization", "Zhengrong Chen, Qinghua Zhu*, An Zeng, Yuzhu Ji, Baoyao Yang, Dan Pan", "IEEE International Conference on Multimedia and Expo (ICME), 2025", 2025, "conference", false, []PublicationLink{link("paper", "Paper", "https://ieeexplore.ieee.org/abstract/document/11210182", "")}),
		pub(11, "/publications/025.png", "Harnessing Feature Distribution Consistency for Federated Learning with Noisy", "Yali Ma, Baoyao Yang*, Yanchao Tang, Weide Zhan, Wenyin Yang", "IEEE International Conference on Image Processing (ICIP), 2025", 2025, "conference", false, []PublicationLink{link("paper", "Paper", "https://ieeexplore.ieee.org/abstract/document/11084722", ""), link("ppt", "PPT", "/files/icip2025.pdf", ""), link("video", "Video", "", "video-player-YaliMa01")}),
		pub(12, "/publications/002.png", "DNA-T: Deformable Neighborhood Attention Transformer for Irregular Medical Time Series", "Jianxuan Huang, Baoyao Yang*, Kejing Yin, Jingwen Xu", "IEEE Journal of Biomedical and Health Informatics (JBHI), 2024", 2024, "journal", false, []PublicationLink{link("code", "Code", "https://github.com/BaoyaoGroup/DNA-T", ""), link("paper", "Paper", "https://ieeexplore.ieee.org/document/10510586", "")}),
		pub(13, "/publications/005.png", "Allosteric Feature Collaboration for Model-Heterogeneous Federated Learning", "Baoyao Yang*, PC Yuen, Yiqun Zhang, An Zeng", "IEEE Transactions on Neural Networks and Learning Systems (TNNLS), 2024", 2024, "journal", false, []PublicationLink{link("paper", "Paper", "https://ieeexplore.ieee.org/abstract/document/10373104", "")}),
		pub(14, "/publications/001.png", "CAM-Guided translation for unpaired weakly-supervised medical image segmentation", "Yuebin Xie, Xiaochen He, Baoyao Yang*, Fei Lyu, Siqi Liu", "IEEE International Conference on Multimedia and Expo (ICME), 2024", 2024, "conference", false, []PublicationLink{link("code", "Code", "https://github.com/BaoyaoGroup/CAM_Guided", ""), link("paper", "Paper", "https://ieeexplore.ieee.org/abstract/document/10687752", "")}),
		pub(15, "/publications/003.png", "Domain Dilation for Single Domain Generalization", "Yuehui Fan, Baoyao Yang*, Meng Shen, Fei Lyu", "IEEE International Conference on Image Processing (ICIP), 2024", 2024, "conference", false, []PublicationLink{link("code", "Code", "https://github.com/BaoyaoGroup/Domain_Dilation_for_Single_Domain_Generalization", ""), link("paper", "Paper", "https://ieeexplore.ieee.org/abstract/document/10648093", "")}),
		pub(16, "/publications/006.png", "MMS: Morphology-mixup Stylized Data Generation for Single Domain Generalization in Medical Image Segmentation", "Xiaochen He, Baoyao Yang*, Fei Lyu", "IEEE International Conference on Acoustics, Speech, and Signal Processing (ICASSP), 2024", 2024, "conference", false, []PublicationLink{link("code", "Code", "https://github.com/BaoyaoGroup/MMS", ""), link("paper", "Paper", "https://ieeexplore.ieee.org/abstract/document/10448305", ""), link("ppt", "PPT", "/files/MMS.pdf", "")}),
		pub(17, "/publications/007.png", "Multi-category Graph Reasoning for Multi-modal Brain Tumor Segmentation", "Dongzhe Li, Baoyao Yang*, Weide Zhan, Xiaochen He", "Medical Image Computing and Computer Assisted Intervention (MICCAI), 2024", 2024, "conference", false, []PublicationLink{link("code", "Code", "https://github.com/BaoyaoGroup/Graph-Co-reasoning", ""), link("paper", "Paper", "https://link.springer.com/chapter/10.1007/978-3-031-72111-3_42", ""), link("poster", "Poster", "/files/ldzposter001.pdf", "")}),
		pub(18, "/publications/022.png", "Multi-category Brain Tumor Segmentation via Multi-scale and Cross-category Relation Modeling", "Dongzhe Li, Baoyao Yang*, Yuebin Xie, Weide Zhan and Jingsong Lin", "HBAI2024: JJCAI Workshop on Human Brain and Artificial Intelligence, 2024", 2024, "conference", false, []PublicationLink{link("paper", "Paper", "https://link.springer.com/chapter/10.1007/978-981-96-4001-0_19", ""), link("ppt", "PPT", "/files/HBAI.pdf", "")}),
		pub(19, "/publications/008.png", "Beyond Direct Relationships: Exploring Multi-Order Label Pair Dependencies for Knowledge Distillation", "Wang J, Deng Z, Lin T, et al.", "Proceedings of the 32nd ACM International Conference on Multimedia, 2024", 2024, "conference", false, []PublicationLink{link("code", "Code", "https://github.com/BaoyaoGroup/Beyond-Direct-Relationships--Exploring-Multi-Order-Distillation-main", ""), link("paper", "Paper", "https://dl.acm.org/doi/abs/10.1145/3664647.3681029", ""), link("poster", "Poster", "/files/wjcposter001.pdf", "")}),
		pub(20, "/publications/009.png", "A Novel Prompt Tuning for Graph Transformers: Tailoring Prompts to Graph Topologies", "Wang J, Deng Z, Lin T, et al.", "Proceedings of the 30th ACM SIGKDD Conference on Knowledge Discovery and Data Mining, 2024", 2024, "conference", false, []PublicationLink{link("code", "Code", "https://github.com/BaoyaoGroup/TGPT_main-main", ""), link("paper", "Paper", "https://dl.acm.org/doi/10.1145/3637528.3671804", ""), link("poster", "Poster", "/files/wjcposter002.pptx", "")}),
		pub(21, "/publications/021.png", "Deep Learning for Brain MRI Confirms Patterned Pathological Progression in Alzheimer's Disease", "Dan Pan, An Zeng*, Baoyao Yang*, Gangyong Lai, Bing Hu, Xiaowei Song, Tianzi Jiang", "Advanced Science, 2023", 2023, "journal", false, []PublicationLink{link("paper", "Paper", "https://advanced.onlinelibrary.wiley.com/doi/full/10.1002/advs.202204717", "")}),
		pub(22, "/publications/010.png", "Early diagnosis of Alzheimer's disease based on multimodal hypergraph attention network", "Yi Li, Baoyao Yang*, Dan Pan, An Zeng, Yang Yang", "International Conference on Multimedia and Expo (ICME), 2023", 2023, "conference", false, []PublicationLink{link("paper", "Paper", "https://ieeexplore.ieee.org/abstract/document/10219893", "")}),
		pub(23, "/publications/011.png", "Model-induced Generalization Error Bound for Information-theoretic Representation Learning in Source-data-free Unsupervised Domain Adaptation", "Baoyao Yang, Hao-wei Yeh, Tatsuya Harada, and Pong C. Yuen*", "IEEE Transactions on Image Processing (TIP), 2022", 2022, "journal", false, []PublicationLink{link("paper", "Paper", "https://ieeexplore.ieee.org/abstract/document/9640468", "")}),
		pub(24, "/publications/012.png", "Revealing Task-relevant Model Memorization for Source-Protected Unsupervised Domain Adaptation", "Baoyao Yang and Pong C. Yuen*", "IEEE Transactions on Information Forensics and Security (TIFS), 2022", 2022, "journal", false, []PublicationLink{link("paper", "Paper", "https://ieeexplore.ieee.org/abstract/document/9705526", "")}),
		pub(25, "/publications/013.png", "Cross-domain Missingness-aware Time Series Adaptation with Similarity Distillation in Medical Applications", "Baoyao Yang, Mang Ye, Qingxiong Tan and Pong C. Yuen*", "IEEE Transactions on Cybernetics (TCYB), 2022", 2022, "journal", false, []PublicationLink{link("paper", "Paper", "https://ieeexplore.ieee.org/abstract/document/9167415", "")}),
		pub(26, "/publications/015.png", "Learning Adaptive Geometry for Unsupervised Domain Adaptation", "Baoyao Yang and Pong C. Yuen*", "Pattern Recognition (PR), 2021", 2021, "journal", false, []PublicationLink{link("paper", "Paper", "https://www.sciencedirect.com/science/article/abs/pii/S0031320320304416", "")}),
		pub(27, "/publications/016.png", "A Segmentation-Assisted Model for Universal Lesion Detection with Partial Labels", "Fei Lyu, Baoyao Yang, Andy J. Ma and Pong C. Yuen*", "International Conference on Medical Image Computing & Computer Assisted Intervention (MICCAI), 2021", 2021, "conference", false, []PublicationLink{link("paper", "Paper", "https://link.springer.com/chapter/10.1007/978-3-030-87240-3_12", "")}),
		pub(28, "/publications/017.png", "Body Parts Synthesis for Cross-Quality Pose Estimation", "Baoyao Yang, Andy J. Ma and Pong C. Yuen*", "IEEE Transactions on Circuits and Systems for Video Technology (TCSVT), 2019", 2019, "journal", false, []PublicationLink{link("paper", "Paper", "https://ieeexplore.ieee.org/document/8245864", "")}),
		pub(29, "/publications/018.png", "Cross-Domain Visual Representations via Unsupervised Graph Alignment", "Baoyao Yang and Pong C. Yuen*", "the 33rd AAAI Conference on Artificial Intelligence (AAAI), 2019", 2019, "conference", false, []PublicationLink{link("paper", "Paper", "https://ojs.aaai.org/index.php/AAAI/article/view/4504", "")}),
		pub(30, "/publications/019.png", "Learning Domain-Shared Group-Sparse Representation for Unsupervised Domain Adaptation", "Baoyao Yang, Andy J. Ma and Pong C. Yuen*", "Pattern Recognition (PR), 2018", 2018, "journal", false, []PublicationLink{link("paper", "Paper", "https://www.sciencedirect.com/science/article/pii/S0031320318301614", "")}),
		pub(31, "/publications/020.png", "Domain-shared Group-sparse Dictionary Learning for Unsupervised Domain Adaptation", "Baoyao Yang, Andy J. Ma and Pong C. Yuen*", "the 32rd AAAI Conference on Artificial Intelligence (AAAI), 2018", 2018, "conference", false, []PublicationLink{link("paper", "Paper", "https://ojs.aaai.org/index.php/AAAI/article/view/12227", "")}),
	}

	for _, item := range items {
		var existing Publication
		if err := s.db.Where("title = ?", item.Title).First(&existing).Error; err == nil {
			continue
		}
		if err := s.db.Create(&item).Error; err != nil {
			return err
		}
	}
	return nil
}

func pub(order int, image string, title string, authors string, venue string, year int, kind string, featured bool, links []PublicationLink) Publication {
	for i := range links {
		links[i].SortOrder = i + 1
	}
	return Publication{
		ImageURL:  image,
		Title:     title,
		Authors:   authors,
		Venue:     venue,
		Year:      year,
		Kind:      kind,
		Status:    StatusPublished,
		Featured:  featured,
		SortOrder: order,
		Links:     links,
	}
}

func link(kind string, label string, url string, routeName string) PublicationLink {
	return PublicationLink{Type: kind, Label: label, URL: url, RouteName: routeName}
}

func (s *Server) seedPatents() error {
	items := []Patent{
		pat(1, "杨宝瑶，黄彦浩，陈涤新", "一种基于时空信息聚合的视频特征提取模型训练方法、系统及特征提取方法", "2025-09-09", "中国", "ZL202510359255.5", "granted"),
		pat(2, "杨宝瑶，麻亚利，詹伟德，唐彦超，卢泽坚", "一种联邦学习场景下的检测噪声标注的方法及系统", "2025-09-05", "中国", "ZL202510375049.3", "granted"),
		pat(3, "杨宝瑶，陈俊祥，黄彦浩，姚文彬", "一种基于多模态大模型的视频 - 文本检索方法", "2025-06-13", "中国", "ZL202411271756.X", "granted"),
		pat(4, "杨宝瑶，郑晓琦", "一种噪声标注的血管图像分割方法及系统", "2025-02-11", "中国", "ZL 202311792631.7", "granted"),
		pat(5, "Dan Pan, An Zeng, Baoyao Yang", "Method for extracting neuroimaging biomarker based on interpretable ensemble 3dcnn", "2024", "美国", "US Patent App. 18/450,338", "granted"),
		pat(6, "杨宝瑶，詹伟德，王健豪", "基于联邦学习的医学图像分割模型构建方法", "", "中国", "ZL202310500954.8", "granted"),
		pat(7, "穆跃鑫，杨宝瑶，宋杨", "一种心脏血管图像分割方法及系统", "", "中国", "ZL202311018267.9", "granted"),
		pat(8, "詹伟德，杨宝瑶，方志祥，李东哲，胡子逸，吴恺盈", "一种基于多级智联的数据异构性联邦学习方法及存储介质", "2024-06-21", "中国", "ZL202410159303.1", "granted"),
		pat(9, "潘丹，曾安，杨宝瑶", "基于可解释集成 3DCNN 的神经影像学生物标志物的提取方法", "2022-11-22", "中国", "ZL202210987102.5", "granted"),
		pat(10, "曾安，谢锐伟，潘丹，杨宝瑶，张逸群", "一种心脏图像分割方法及系统", "2022-04-22", "中国", "ZL202210030012.3", "granted"),
		pat(11, "甘孟坤，潘丹，曾安，杨宝瑶，张逸群", "一种主动脉弓再缩窄的预后方法及系统", "2025-04-01", "中国", "ZL202210655677.3", "granted"),
		pat(12, "潘丹，罗琳，曾安，廖清青，杨宝瑶，张逸群", "一种基于多头两级注意力的三维点云语义分割方法", "2022-11-04", "中国", "ZL202210709918.8", "granted"),
		pat(13, "曾安，许鸿迈，潘丹，朱清华，杨宝瑶，姬玉柱", "基于全局 Critic 多智能体算法的多重供应链调度方法及系统", "2025-07-18", "中国", "ZL202411209654.5", "granted"),
		pat(14, "杨宝瑶，詹伟德", "一种模型异构性联邦学习方法和系统", "", "中国", "ZL202210989290.1", "review"),
		pat(15, "杨宝瑶，范月辉，郑晓琦，麻亚利", "一种基于纹理增强和形态增强的单域泛化方法", "", "中国", "ZL202510410016.8", "review"),
		pat(16, "曾安，赖峻浩，潘丹，杨宝瑶，赵靖亮", "一种病理图像分割网络模型、方法、装置及介质", "", "中国", "ZL202410014806.X", "review"),
		pat(17, "曾安，谢建萍，潘丹，杨宝瑶，吴菊华，陈文戈", "基于融合超图和超图神经网络的阿尔茨海默症分类方法", "", "中国", "ZL202410739087.8", "review"),
		pat(18, "曾安，陈峥嵘，朱清华，姬玉柱，杨宝瑶，潘丹，吴恒", "一种基于深度强化学习的供应链资源控制优化方法", "", "中国", "ZL202411900442.1", "review"),
		pat(19, "曾安，程贤航，潘丹，杨宝瑶", "一种三维多尺度图像分割模型和基于其的图像分割方法、应用、系统", "", "中国", "ZL202510117218.3", "review"),
		pat(20, "曾安，李彦霆，潘丹，杨宝瑶，谢锐伟", "一种多模态医疗信息融合方法", "", "中国", "ZL202510110325.3", "review"),
		pat(21, "刘培松，王丹，潘丹，曾安，杨宝瑶，刘鑫，杨洋，刘军", "一种基于全局和局部注意力机制的肺部 CT 图像分割方法", "", "中国", "ZL202311629043.1", "review"),
		pat(22, "曾安，何诗婷，潘丹，姬玉柱，杨宝瑶，赵靖亮", "一种二阶段的 3D 冠状动脉分割重构方法和系统", "", "中国", "ZL202310639973.9", "review"),
		pat(23, "曾安，王健斌，潘丹，杨宝瑶，杨洋，刘军", "一种基于三维卷积和孪生神经网络的 AD 分类方法和系统", "", "中国", "ZL202310730443.5", "review"),
		pat(24, "曾安，林先扬，赵靖亮，潘丹，杨宝瑶，赵屾", "一种基于深度强化学习的主动脉夹层分割方法和系统", "", "中国", "ZL202310778789.2", "review"),
		pat(25, "曾安，庞耀幸，潘丹，赵靖亮，杨宝瑶，赵屾", "一种基于深度强化学习的左心室内膜图像分割方法和系统", "", "中国", "ZL202310843676.6", "review"),
		pat(26, "曾安，李艺，潘丹，杨宝瑶，张逸群，杨洋，刘军", "基于多模态超图注意网络的阿尔茨海默症分类方法及系统", "", "中国", "ZL202310565619.6", "review"),
		pat(27, "曾安，罗百荣，潘丹，杨洋，刘军，杨宝瑶，赵靖亮", "一种基于超图神经网络的阿尔茨海默症分类方法和系统", "", "中国", "ZL202310482507.4", "review"),
		pat(28, "詹少强，张逸群，潘丹，曾安，姬玉柱，杨宝瑶，蒋艳荣，孙鸿涛", "一种重症监护数据处理方法、应用及设备", "", "中国", "ZL202211297374.5", "review"),
		pat(29, "曾安，黎锦荣，潘丹，杨宝瑶，赵靖亮", "一种基于迭代计算图神经网络的蛋白质序列设计方法", "", "中国", "ZL202410979540.2", "review"),
		{Authors: "李春林，吴恒，曾安，惠恩明，富锐，杨思维，曹洪江，骆有隆，杨宝瑶，刘俊，张勇，江焜", Title: "产业聚集区域内业务资源服务平台规范", Date: "2025-05-01", Country: "广东省工业软件学会团体标准", Number: "T/GISF 002-2024", Category: "standard", Status: StatusPublished, SortOrder: 30},
	}
	return s.db.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "title"}}, DoNothing: true}).Create(&items).Error
}

func pat(order int, authors string, title string, date string, country string, number string, category string) Patent {
	return Patent{Authors: authors, Title: title, Date: date, Country: country, Number: number, Category: category, Status: StatusPublished, SortOrder: order}
}
