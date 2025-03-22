<script>
import VideoPlayer from './components/VideoPlayer.vue';

export default {
  components: { VideoPlayer },
  data() {
    return {
      videoList: [
        { title: '科研介绍视频1', src: '/videos/v1.mp4', duration: '03:20' },
        { title: '实验过程记录', src: '/videos/v2.mp4', duration: '12:45' }
      ]
    }
  },
  methods: {
    selectVideo(row) {
      this.$refs.videoPlayer.src = row.src; // 触发视频源更新[2](@ref)
    }
  }
}
</script>

<template>
  <el-main>
    <el-row :gutter="20">
      <el-col :span="16">
        <VideoPlayer
            src="/videos/demo.mp4"
            poster="/images/poster.jpg"
            :autoplay="false"
        />
      </el-col>
      <el-col :span="8">
        <el-card header="视频列表">
          <el-table :data="videoList" @row-click="selectVideo">
            <el-table-column prop="title" label="标题"/>
            <el-table-column prop="duration" label="时长"/>
          </el-table>
        </el-card>
      </el-col>
    </el-row>
  </el-main>
</template>

