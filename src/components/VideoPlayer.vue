<template>
  <div class="video-container">
    <!-- 播放器容器 -->
    <video ref="videoElement" class="video-js vjs-big-play-centered"></video>

    <!-- Element Plus自定义控制栏 -->
    <div v-if="showCustomControls" class="custom-controls">
      <el-button @click="togglePlay">
        {{ isPlaying ? '⏸️ 暂停' : '▶️ 播放' }}
      </el-button>
      <el-slider
          v-model="currentTime"
          :max="duration"
          :format-tooltip="formatTime"
          @change="handleSeek"
      />
      <el-button @click="toggleFullscreen">📺 全屏</el-button>
    </div>
  </div>
</template>

<script>
import videojs from 'video.js';
import 'video.js/dist/video-js.css';
import '@videojs/themes/dist/forest/theme.css'; // 使用森林主题[3](@ref)

export default {
  props: {
    src: {
      type: String,
      required: true
    },
    poster: String,
    autoplay: Boolean
  },
  data() {
    return {
      player: null,
      isPlaying: false,
      currentTime: 0,
      duration: 0
    };
  },
  computed: {
    showCustomControls() {
      return this.player && !this.player.isFullscreen();
    }
  },
  mounted() {
    this.initPlayer();
  },
  beforeUnmount() {
    if (this.player) {
      this.player.dispose();
    }
  },
  methods: {
    initPlayer() {
      // 播放器配置[1,3](@ref)
      this.player = videojs(this.$refs.videoElement, {
        controls: false, // 禁用原生控件
        autoplay: this.autoplay,
        poster: this.poster,
        sources: [{
          src: this.src,
          type: this.getVideoType(this.src)
        }],
        theme: 'forest', // 应用主题
        fluid: true      // 自适应容器
      });

      // 事件监听[1,2](@ref)
      this.player.on('play', () => this.isPlaying = true);
      this.player.on('pause', () => this.isPlaying = false);
      this.player.on('timeupdate', () =>
          this.currentTime = this.player.currentTime()
      );
      this.player.on('loadedmetadata', () =>
          this.duration = this.player.duration()
      );
    },

    // 时间格式转换
    formatTime(seconds) {
      const date = new Date(seconds * 1000);
      return date.toISOString().substr(11, 8);
    },

    // 播放/暂停切换
    togglePlay() {
      this.isPlaying ? this.player.pause() : this.player.play();
    },

    // 进度条跳转
    handleSeek(time) {
      this.player.currentTime(time);
    },

    // 全屏切换
    toggleFullscreen() {
      if (this.player.isFullscreen()) {
        this.player.exitFullscreen();
      } else {
        this.player.requestFullscreen();
      }
    },

    // 判断视频类型[4](@ref)
    getVideoType(src) {
      const ext = src.split('.').pop();
      return `video/${ext === 'mp4' ? 'mp4' : 'webm'}`;
    }
  },
  watch: {
    // 动态更新视频源[2](@ref)
    src(newVal) {
      this.player.src({
        src: newVal,
        type: this.getVideoType(newVal)
      });
      this.player.load();
    }
  }
};
</script>

<style scoped>
.video-container {
  position: relative;
}
.custom-controls {
  position: absolute;
  bottom: 10px;
  left: 0;
  right: 0;
  padding: 10px;
  background: rgba(0,0,0,0.7);
  display: flex;
  gap: 10px;
  align-items: center;
}

/* 覆盖Video.js默认样式 */
:deep(.video-js) {
  height: 500px;
}
:deep(.vjs-big-play-button) {
  background-color: var(--el-color-primary) !important;
}
</style>