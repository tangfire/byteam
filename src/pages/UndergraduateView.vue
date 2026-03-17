<script setup lang="ts">
defineProps<{ msg: string }>()

// 本科生培养数据
const undergraduateStudents = [
  {
    id: 1,
    name: "穆跃鑫",
    grade: "19 级",
    major: "机器人学院",
    direction: "机器学习与数据挖掘",
    achievements: [
      "申请国家发明专利 1 项",
      "获优秀毕业设计",
      "保送至重庆大学"
    ]
  },
  {
    id: 2,
    name: "何晓琛",
    grade: "20 级",
    major: "软件工程（卓越班）",
    direction: "医学图像处理与单域泛化",
    achievements: [
      "发表 CCF-B 类会议 ICASSP2024 论文一篇"
    ]
  },
  {
    id: 3,
    name: "王敬超",
    grade: "21 级",
    major: "人工智能专业",
    direction: "机器学习与数据挖掘",
    achievements: [
      "以第一作者在 CCF-A 类会议 SIGKDD 2024 发表论文一篇",
      "以第一作者在 CCF-A 类会议 ACMMM 2024 发表论文一篇",
      "发表 JBHI 论文一篇",
      "荣获科技创新奖学金",
      "保送至北京大学计算机科学与技术专业（博士）"
    ]
  },
  {
    id: 4,
    name: "詹伟德",
    grade: "21 级",
    major: "计算机科学与技术（伏羲班）",
    direction: "联邦学习",
    achievements: [
      "授权国家发明专利 1 项，实审 2 项",
      "获省级赛事奖项 2 项",
      "参与发表 CCF-B 类会议 MICCAI 2024 论文一篇",
      "连续两年获得校一等奖学金",
      "国家优秀学生奖学金",
      "保送至复旦大学电子信息专业（博士）"
    ]
  },
  {
    id: 5,
    name: "李东哲",
    grade: "21 级",
    major: "计算机科学与技术（伏羲班）",
    direction: "医学图像分割技术",
    achievements: [
      "发表 TMM（CCF-B）论文一篇",
      "发表医学信息处理顶会 MICCAI 2024（CCF-B）论文一篇",
      "发表 IJCAI-HBAI 论文一篇"
    ]
  },
  {
    id: 6,
    name: "方志祥",
    grade: "21 级",
    major: "计算机科学与技术",
    direction: "联邦学习公平性问题",
    achievements: [
      "发表 ECAI 2025（CCF-B）论文一篇"
    ]
  }
]
</script>

<template>
  <div class="undergraduate-container">
    <div class="header-section">
      <h1 class="page-title">Undergraduate Education</h1>
      <p class="page-subtitle">本科生培养成果展示</p>
    </div>

    <div class="students-content">
      <div class="students-grid">
        <el-card
            v-for="student in undergraduateStudents"
            :key="student.id"
            class="student-card"
        >
          <div class="student-header">
            <div class="student-info">
              <h3 class="student-name">{{ student.name }}</h3>
              <div class="student-meta">
                <span class="student-grade">{{ student.grade }}</span>
                <span class="student-major">{{ student.major }}</span>
              </div>
            </div>
            <div class="student-direction">
              <el-icon><Research /></el-icon>
              <span>{{ student.direction }}</span>
            </div>
          </div>

          <div class="achievements-section">
            <h4 class="achievements-title">培养成果</h4>
            <ul class="achievements-list">
              <li v-for="(achievement, index) in student.achievements" :key="index">
                <el-icon><Star /></el-icon>
                <span>{{ achievement }}</span>
              </li>
            </ul>
          </div>
        </el-card>
      </div>
    </div>
  </div>

  <div class="bottom-spacer"></div>
  <el-backtop class="mobile-backtop" :right="100" :bottom="100"/>
</template>

<style scoped>
.undergraduate-container {
  width: 100%;
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
  box-sizing: border-box;
}

.header-section {
  text-align: center;
  margin-bottom: 40px;
  padding-top: 20px;
}

.page-title {
  color: #7d1231;
  font-size: 32px;
  font-weight: 600;
  margin: 0 0 10px 0;
}

.page-subtitle {
  color: #666;
  font-size: 16px;
  margin: 0;
}

.students-content {
  padding: 20px 0;
}

.students-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(380px, 1fr));
  gap: 24px;
}

.student-card {
  border-radius: 12px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
  transition: all 0.3s ease;
  border: 1px solid #eaeaea;
  background: white;
}

.student-card:hover {
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.12);
  transform: translateY(-4px);
  border-color: #7d1231;
}

.student-header {
  padding: 20px 20px 15px;
  border-bottom: 2px solid #f5f5f5;
}

.student-info {
  margin-bottom: 12px;
}

.student-name {
  font-size: 22px;
  font-weight: 600;
  color: #7d1231;
  margin: 0 0 8px 0;
}

.student-meta {
  display: flex;
  gap: 12px;
  align-items: center;
}

.student-grade,
.student-major {
  font-size: 14px;
  color: #666;
  padding: 4px 10px;
  background: #f8f9fa;
  border-radius: 6px;
  font-weight: 500;
}

.student-direction {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 14px;
  color: #13393e;
  font-weight: 500;
}

.student-direction .el-icon {
  color: #7d1231;
  font-size: 16px;
}

.achievements-section {
  padding: 20px;
}

.achievements-title {
  font-size: 16px;
  font-weight: 600;
  color: #7d1231;
  margin: 0 0 15px 0;
  padding-bottom: 10px;
  border-bottom: 1px solid #eee;
}

.achievements-list {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.achievements-list li {
  display: flex;
  align-items: flex-start;
  gap: 10px;
  font-size: 14px;
  color: #333;
  line-height: 1.6;
}

.achievements-list li .el-icon {
  color: #f9ab00;
  font-size: 16px;
  flex-shrink: 0;
  margin-top: 2px;
}

.bottom-spacer {
  height: 50px;
}

.mobile-backtop {
  right: 100px;
  bottom: 100px;
}

/* 移动端适配 */
@media (max-width: 768px) {
  .undergraduate-container {
    padding: 0 15px;
  }

  .page-title {
    font-size: 26px;
  }

  .page-subtitle {
    font-size: 14px;
  }

  .students-grid {
    grid-template-columns: 1fr;
    gap: 16px;
  }

  .student-name {
    font-size: 20px;
  }

  .student-meta {
    flex-direction: column;
    gap: 6px;
    align-items: flex-start;
  }

  .achievements-list li {
    font-size: 13px;
  }

  .mobile-backtop {
    right: 20px !important;
    bottom: 80px !important;
  }
}

/* 小屏幕优化 */
@media (max-width: 480px) {
  .student-header,
  .achievements-section {
    padding: 15px;
  }

  .student-name {
    font-size: 18px;
  }

  .student-grade,
  .student-major,
  .student-direction {
    font-size: 13px;
  }
}

/* 确保图标颜色正确 */
::v-deep(.el-icon svg) {
  color: #7d1231 !important;
}
</style>
