<script setup lang="ts">
// 组件逻辑（如有需要可在此添加）
</script>

<template>






  <el-space direction="vertical"  style="width: 100%">
    <p style="margin-left: -550px; color: #7d1231; font-size: 26px;">Graduate Students</p>
  <el-row :gutter="30">
    <el-col :span="6">
      <div class="block">
        <el-avatar shape="square" :size="250" src="/avatar/001.png" />
        <p class="name">Dixin Chen</p>
        <p class="work">Short video understanding</p>
      </div>
    </el-col>
    <el-col :span="6">
      <div class="block">
        <el-avatar shape="square" :size="250" src="/avatar/001.png" />
        <p class="name">Huahong Deng</p>
        <p class="work">Industrial software</p>
      </div>
    </el-col>
    <el-col :span="6">
      <div class="block">
        <el-avatar shape="square" :size="250" src="/avatar/001.png" />
        <p class="name">Canrong Du</p>
        <p class="work">Short video understanding</p>
      </div>
    </el-col>
    <el-col :span="6">
      <div class="block">
        <el-avatar shape="square" :size="250" src="/avatar/001.png" />
        <p class="name">Yuehui Fan</p>
        <p class="work">Short video understanding</p>
      </div>
    </el-col>
  </el-row>

    <!--  底部间隔-->
    <div style="height: 50px"></div>



    <el-row :gutter="30">
      <el-col :span="6">
        <div class="block">
          <el-avatar shape="square" :size="250" src="/avatar/001.png" />
          <p class="name">Jianxuan Huang</p>
          <p class="work">Short video understanding</p>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="block">
          <el-avatar shape="square" :size="250" src="/avatar/001.png" />
          <p class="name">Yanhao Huang</p>
          <p class="work">Short video understanding</p>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="block">
          <el-avatar shape="square" :size="250" src="/avatar/001.png" />
          <p class="name">Jingsong Lin</p>
          <p class="work">Short video understanding</p>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="block">
          <el-avatar shape="square" :size="250" src="/avatar/001.png" />
          <p class="name">Yali Ma</p>
          <p class="work">Short video understanding</p>
        </div>
      </el-col>
    </el-row>

    <!--  底部间隔-->
    <div style="height: 50px"></div>



    <el-row :gutter="30">
      <el-col :span="6">
        <div class="block">
          <el-avatar shape="square" :size="250" src="/avatar/001.png" />
          <p class="name">Yanchao Tang</p>
          <p class="work">Short video understanding</p>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="block">
          <el-avatar shape="square" :size="250" src="/avatar/001.png" />
          <p class="name">Kexin Xie</p>
          <p class="work">Short video understanding</p>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="block">
          <el-avatar shape="square" :size="250" src="/avatar/001.png" />
          <p class="name">Yuebin Xie</p>
          <p class="work">Short video understanding</p>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="block">
          <el-avatar shape="square" :size="250" src="/avatar/001.png" />
          <p class="name">Yufei Xu</p>
          <p class="work">Short video understanding</p>
        </div>
      </el-col>
    </el-row>

    <div style="height: 50px"></div>

    <el-row :gutter="30">
      <el-col :span="8">
        <div class="block">
          <el-avatar shape="square" :size="250" src="/avatar/001.png" />
          <p class="name">Xianrun Xu</p>
          <p class="work">Short video understanding</p>
        </div>
      </el-col>
      <el-col :span="8">
        <div class="block">
          <el-avatar shape="square" :size="250" src="/avatar/001.png" />
          <p class="name">Xiaoqi Zheng</p>
          <p class="work">Short video understanding</p>
        </div>
      </el-col>
      <el-col :span="8">
        <div class="block">
          <el-avatar shape="square" :size="250" src="/avatar/001.png" />
          <p class="name">Sijia Zhou</p>
          <p class="work">Short video understanding</p>
        </div>
      </el-col>
<!--      <el-col :span="6">-->
<!--        <div class="block">-->
<!--          <el-avatar shape="square" :size="250" src="/avatar/001.png" />-->
<!--          <p>Yufei Xu</p>-->
<!--          <p>Short video understanding</p>-->
<!--        </div>-->
<!--      </el-col>-->
    </el-row>

    <div style="height: 50px"></div>

    <p style="margin-left: -550px; color: #7d1231; font-size: 26px;">Undergraduate Students</p>
    <el-row :gutter="30">
      <el-col :span="8">
        <div class="block">
          <el-avatar shape="square" :size="250" src="/avatar/001.png" />
          <p class="name">Jingchao Wang</p>
          <p class="work">Industrial software</p>
        </div>
      </el-col>
      <el-col :span="8">
        <div class="block">
          <el-avatar shape="square" :size="250" src="/avatar/001.png" />
          <p class="name">Dongzhe Li</p>
          <p class="work">Short video understanding</p>
        </div>
      </el-col>
      <el-col :span="8">
        <div class="block">
          <el-avatar shape="square" :size="250" src="/avatar/001.png" />
          <p class="name">Yuehui Fan</p>
          <p class="work">Short video understanding</p>
        </div>
      </el-col>
    </el-row>

    <!--  底部间隔-->
    <div style="height: 50px"></div>


  </el-space>

  <el-backtop :right="100" :bottom="100"/>




</template>

<style scoped>

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

</style>