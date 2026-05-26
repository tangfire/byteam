<script setup lang="ts">
import { Files, Postcard, Promotion, Tickets, VideoPlay } from '@element-plus/icons-vue'
import { computed, onMounted, ref } from 'vue'
import { getPublicPublications } from '../api/public'
import type { Publication, PublicationLink } from '../api/client'
import { createFallbackPublicationGroups, type PublicationGroups, type PublicationLinkMap, type PublicationLinkValue } from '../data/fallbacks/publications'
import { openExternalLink } from '../utils/links'
import { resolveExternalVideoURL, resolveVideoPagePath } from '../utils/videoLinks'

const downloadFile = (filePath: string, fileName: string) => {
  const link = document.createElement('a')
  link.href = filePath
  link.download = fileName
  document.body.appendChild(link)
  link.click()
  document.body.removeChild(link)
}

const iconMap = {
  code: Promotion,
  paper: Tickets,
  video: VideoPlay,
  ppt: Files,
  poster: Postcard,
}

const getIcon = (type: string) => iconMap[type as keyof typeof iconMap] || Promotion
const getLinkText = (type: string) => {
  if (type === 'ppt') return 'PPT'
  return type.charAt(0).toUpperCase() + type.slice(1)
}

const mapLinks = (links: PublicationLink[]) => links.reduce((acc: PublicationLinkMap, link) => {
  if (link.type === 'video') {
    acc.video = { path: resolveVideoPagePath(link), url: resolveExternalVideoURL(link) }
  } else if (link.type === 'ppt' || link.type === 'poster') {
    acc[link.type] = { handler: () => downloadFile(link.url, link.label || getLinkText(link.type)) }
  } else {
    acc[link.type] = link.url
  }
  return acc
}, {})

const linkType = (type: string | number) => String(type)
const linkHref = (link: PublicationLinkValue) => (typeof link === 'string' ? link : undefined)
const linkRoute = (link: PublicationLinkValue) => (typeof link === 'object' && link.path ? link.path : '/')
const videoHasRoute = (link: PublicationLinkValue) => typeof link === 'object' && Boolean(link.path)
const videoHasURL = (link: PublicationLinkValue) => typeof link === 'object' && Boolean(link.url)
const openVideoLink = (link: PublicationLinkValue) => {
  if (typeof link === 'object' && link.url) {
    openExternalLink(link.url)
  }
}
const runLinkHandler = (link: PublicationLinkValue) => {
  if (typeof link === 'object') {
    link.handler?.()
  }
}

const mapPublication = (pub: Publication) => ({
  id: pub.id,
  image: pub.image,
  title: pub.title,
  authors: pub.authors,
  venue: pub.venue,
  links: mapLinks(pub.links || []),
})

const publications = ref<PublicationGroups>(createFallbackPublicationGroups(downloadFile))

// 年份数组，按降序排列
const years = computed(() => Object.keys(publications.value).sort((a, b) => Number(b) - Number(a)))

onMounted(async () => {
  try {
    const result = await getPublicPublications()
    const next: PublicationGroups = {}
    result.items.forEach((pub) => {
      const year = String(pub.year)
      next[year] = next[year] || { journal: [], conference: [] }
      next[year][pub.kind as 'journal' | 'conference'].push(mapPublication(pub))
    })
    publications.value = next
  } catch (error) {
    console.warn('Using local publication fallback data', error)
  }
})
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
                          v-if="linkType(type) === 'code' || linkType(type) === 'paper'"
                          :href="linkHref(link)"
                          class="publication-link"
                          target="_blank"
                          rel="noopener noreferrer"
                      >
                        <el-icon size="25">
                          <component :is="getIcon(linkType(type))" />
                        </el-icon>
                        <span>{{ getLinkText(linkType(type)) }}</span>
                      </a>

                      <!-- 视频链接 -->
                      <router-link
                          v-else-if="linkType(type) === 'video' && videoHasRoute(link)"
                          :to="linkRoute(link)"
                          class="publication-link"
                      >
                        <el-icon size="25">
                          <VideoPlay />
                        </el-icon>
                        <span>Video</span>
                      </router-link>
                      <button
                          v-else-if="linkType(type) === 'video' && videoHasURL(link)"
                          type="button"
                          class="publication-link link-button"
                          @click="openVideoLink(link)"
                      >
                        <el-icon size="25">
                          <VideoPlay />
                        </el-icon>
                        <span>Video</span>
                      </button>

                      <!-- PPT 和海报下载按钮 -->
                      <el-button
                          v-else-if="linkType(type) === 'ppt' || linkType(type) === 'poster'"
                          class="custom-button"
                          @click="runLinkHandler(link)"
                      >
                        <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                          <component :is="getIcon(linkType(type))" />
                        </el-icon>
                        <span class="button-text">{{ getLinkText(linkType(type)) }}</span>
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
                          v-if="linkType(type) === 'code' || linkType(type) === 'paper'"
                          :href="linkHref(link)"
                          class="publication-link"
                          target="_blank"
                          rel="noopener noreferrer"
                      >
                        <el-icon size="25">
                          <component :is="getIcon(linkType(type))" />
                        </el-icon>
                        <span>{{ getLinkText(linkType(type)) }}</span>
                      </a>

                      <!-- 视频链接 -->
                      <router-link
                          v-else-if="linkType(type) === 'video' && videoHasRoute(link)"
                          :to="linkRoute(link)"
                          class="publication-link"
                      >
                        <el-icon size="25">
                          <VideoPlay />
                        </el-icon>
                        <span>Video</span>
                      </router-link>
                      <button
                          v-else-if="linkType(type) === 'video' && videoHasURL(link)"
                          type="button"
                          class="publication-link link-button"
                          @click="openVideoLink(link)"
                      >
                        <el-icon size="25">
                          <VideoPlay />
                        </el-icon>
                        <span>Video</span>
                      </button>

                      <!-- PPT 和海报下载按钮 -->
                      <el-button
                          v-else-if="linkType(type) === 'ppt' || linkType(type) === 'poster'"
                          class="custom-button"
                          @click="runLinkHandler(link)"
                      >
                        <el-icon size="25" style="margin-right: 8px; vertical-align: middle;">
                          <component :is="getIcon(linkType(type))" />
                        </el-icon>
                        <span class="button-text">{{ getLinkText(linkType(type)) }}</span>
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
  color: #7d1231; /* 强制使用主题色深红色，与 PPT 按钮一致 */
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

.link-button {
  border: 0;
  background: transparent;
  cursor: pointer;
  font-family: inherit;
}

.publication-link:hover {
  color: #13393e; /* 悬停时使用深青色，与 PPT 按钮一致 */
  background-color: rgba(125, 18, 49, 0.05);
}

.publication-link:hover ::v-deep(.el-icon svg) {
  color: #13393e !important; /* 悬停时图标也使用深青色 */
}

.bottom-spacer {
  height: 50px;
}

/* 按钮样式 */
.custom-button {
  background: transparent !important;
  border: none !important;
  padding: 0.6rem 1.2rem !important;
  box-shadow: none !important;
  transition: all 0.3s ease !important;
  border-radius: 12px !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  cursor: pointer !important;
  min-width: auto !important;
}

.custom-button:hover {
  background: rgba(125, 18, 49, 0.05) !important;
}

.custom-button .button-text {
  font-size: 1.05rem;
  font-weight: 600;
  color: #7d1231 !important; /* 强制使用主题色深红色 */
  transition: color 0.3s ease;
}

.custom-button:hover .button-text {
  color: #13393e !important; /* 悬停时使用深青色 */
}

.custom-button :deep(.el-icon svg) {
  color: #7d1231 !important; /* 强制使用主题色深红色，与 Paper/Code 链接一致 */
  transition: color 0.3s ease;
}

.custom-button:hover :deep(.el-icon svg) {
  color: #13393e !important; /* 悬停时使用深青色，与 Paper/Code 链接一致 */
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
