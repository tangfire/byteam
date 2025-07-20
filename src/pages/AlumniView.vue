<script setup lang="ts">
// 组件逻辑（如有需要可在此添加）
</script>

<template>



  <el-space direction="vertical"  style="width: 100%">
    <p style="margin-left: -550px; color: #7d1231; font-size: 26px;">Graduate Students</p>
  <el-row :gutter="30">



    <el-col :span="8">
      <div class="block">
        <el-avatar shape="square" :size="250" src="/avatar/YuehuiFan.jpg" />
        <p class="name">Yuehui Fan</p>
        <p class="work">(~2025.07)</p>
      </div>
    </el-col>

    <el-col :span="8">
      <div class="block">
        <el-avatar class="jianxuanHuang" shape="square" :size="250" src="/avatar/JianxuanHuang.jpg" />
        <p class="name">Jianxuan Huang</p>
        <p class="work">(~2025.07)</p>
        <!--          <p class="work">Short video understanding</p>-->
      </div>
    </el-col>

    <el-col :span="8">
      <div class="block">
        <el-avatar shape="square" :size="250" src="/avatar/YuebinXie.jpg" />
        <p class="name">Yuebin Xie</p>
        <p class="work">(~2025.07)</p>
        <!--          <p class="work">Short video understanding</p>-->
      </div>
    </el-col>
  </el-row>





    <div style="height: 50px"></div>

    <p style="margin-left: -550px; color: #7d1231; font-size: 26px;">Research-based undergraduates</p>
    <el-row :gutter="30">
      <el-col :span="6">
        <div class="block">
          <el-avatar class="JingchaoWang" style="background-color: white" shape="square" :size="250" src="/avatar/JingchaoWang.jpg" />
          <p class="name">Jingchao Wang</p>
          <p class="work">(~2025.07)</p>
<!--          <p class="work">Industrial software</p>-->
        </div>
      </el-col>
      <el-col :span="6">
        <div class="block">
          <el-avatar class="DongzheLi" style="background-color: white" shape="square" :size="250" src="/avatar/DongzheLi.jpg" />
          <p class="name">Dongzhe Li</p>
          <p class="work">(~2025.07)</p>
<!--          <p class="work">Short video understanding</p>-->
        </div>
      </el-col>
      <el-col :span="6">
        <div class="block">
          <el-avatar class="XiaochenHe" style="background-color: white" shape="square" :size="250" src="/avatar/XiaochenHe.jpg" />
          <p class="name">Xiaochen He</p>
          <p class="work">(~2025.07)</p>
<!--          <p class="work">Short video understanding</p>-->
        </div>
      </el-col>

      <el-col :span="6">
        <div class="block">
          <el-avatar class="WeideZhan" style="background-color: white" shape="square" :size="250" src="/avatar/WeideZhan.jpg" />
          <p class="name">Weide Zhan</p>
          <p class="work">(~2025.07)</p>
<!--          <p class="work">Short video understanding</p>-->
        </div>
      </el-col>
    </el-row>

    <!--  底部间隔-->
    <div style="height: 50px"></div>

    <el-row :gutter="30">
      <el-col :span="24">
        <div class="block">
          <el-avatar class="ZhixiangFang" style="background-color: white" shape="square" :size="250" src="/avatar/ZhixiangFang.jpg" />
          <p class="name">Zhixiang Fang</p>
          <p class="work">(~2025.07)</p>
          <!--          <p class="work">Industrial software</p>-->
        </div>
      </el-col>
      <!-- 补全剩余空间 -->
      <el-col :span="18"></el-col> <!-- 6 + 18 = 24 -->

    </el-row>

    <!--  底部间隔-->
    <div style="height: 50px"></div>


  </el-space>

  <el-backtop class="mobile-backtop" :right="100" :bottom="100"/>




</template>

<style scoped>



@media (max-width: 768px) {

  .mobile-backtop {
    right: 20px !important;
    bottom: 80px !important;
  }

  .el-row {
    justify-content: center !important;
  }

  .el-col {
    width: 50% !important;
    max-width: 200px !important;
    margin-bottom: 30px;
  }

  p[style*="margin-left"] {
    text-align: center !important;
    margin-left: 0 !important;
    width: 100%;
  }

  /* 调整头像大小 */
  .el-avatar {
    width: 200px !important;
    height: 200px !important;
  }

  /* 调整文字大小 */
  .name {
    font-size: 1.4rem;
  }

  .work {
    font-size: 1.1rem;
  }

  /* 移除左边距 */
  p[style*="margin-left"] {
    margin-left: 0 !important;
    text-align: center;
  }


}

.block {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px; /* 控制整体间距 */
}




.name{
  color: #2c2f3b;
  font-size: 1.6rem;
  margin-top: -5px;  /* 上移文字 */
  margin-bottom: 2px;
}

.work{
  font-size: 1.2rem;
  color: #7f8c8d;
  margin-top: -3px;
}

::v-deep(.el-icon svg) {
  color: #7d1231 !important;
}






::v-deep .jianxuanHuang img{
  object-position: center 5%; /* 垂直方向向下偏移 */
  object-fit: cover; /* 保持图片比例 */
}





::v-deep .DongzheLi img{
  /* 新增缩放相关属性 */
  width: 70%; /* 按比例缩小到85%（可根据需要调整） */
  max-width: 100%; /* 防止溢出容器 */
  height: auto; /* 保持高度自适应 */
  margin-top: 10px;
  /* 保持原有设置 */
  object-position: center 100%;
  object-fit: cover;
}

::v-deep .XiaochenHe img{
  /* 新增缩放相关属性 */
  width: 70%; /* 按比例缩小到85%（可根据需要调整） */
  max-width: 100%; /* 防止溢出容器 */
  height: auto; /* 保持高度自适应 */
  margin-top: 10px;
  /* 保持原有设置 */
  object-position: center 100%;
  object-fit: cover;
}

::v-deep .JingchaoWang img{
  /* 新增缩放相关属性 */
  width: 70%; /* 按比例缩小到85%（可根据需要调整） */
  max-width: 100%; /* 防止溢出容器 */
  height: auto; /* 保持高度自适应 */
  margin-top: 10px;
  /* 保持原有设置 */
  object-position: center 100%;
  object-fit: cover;
}

::v-deep .WeideZhan img{
  /* 新增缩放相关属性 */
  width: 70%; /* 按比例缩小到85%（可根据需要调整） */
  max-width: 100%; /* 防止溢出容器 */
  height: auto; /* 保持高度自适应 */
  margin-top: 10px;
  /* 保持原有设置 */
  object-position: center 100%;
  object-fit: cover;
}

::v-deep .ZhixiangFang img{
  object-position: center 5%; /* 垂直方向向下偏移 */
  object-fit: cover; /* 保持图片比例 */
}


</style>