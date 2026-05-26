<template>
  <section class="research-spotlight" aria-label="Featured publications">
    <div v-if="activePublication" class="spotlight-layout">
      <div class="paper-visual">
        <img :src="activePublication.image" :alt="activePublication.title" @error="handleImageError" />
      </div>

      <div class="paper-info">
        <div class="paper-meta">
          <span>{{ activePublication.kind }}</span>
          <span>{{ activePublication.year }}</span>
        </div>

        <h3>{{ activePublication.title }}</h3>
        <p class="authors">{{ activePublication.authors }}</p>
        <p class="venue">{{ activePublication.venue }}</p>

        <div class="paper-links" v-if="activePublication.links.length">
          <button
              v-for="link in sortedLinks(activePublication)"
              :key="`${link.type}-${link.url}-${link.routeName}`"
              type="button"
              class="paper-link"
              @click="handleLinkClick(link)"
          >
            {{ link.label || link.type }}
          </button>
        </div>
      </div>
    </div>

    <div class="carousel-controls" v-if="publications.length > 1">
      <button class="nav-button" type="button" aria-label="Previous publication" @click="prevSlide">
        <el-icon><ArrowLeft /></el-icon>
      </button>

      <div class="paper-tabs" role="tablist" aria-label="Featured publication selector">
        <button
            v-for="(pub, index) in publications"
            :key="pub.id || `${pub.title}-${index}`"
            type="button"
            class="paper-tab"
            :class="{ active: index === currentIndex }"
            :aria-selected="index === currentIndex"
            @click="goToSlide(index)"
        >
          <span class="tab-index">{{ String(index + 1).padStart(2, '0') }}</span>
          <span class="tab-title">{{ pub.title }}</span>
        </button>
      </div>

      <button class="nav-button" type="button" aria-label="Next publication" @click="nextSlide">
        <el-icon><ArrowRight /></el-icon>
      </button>
    </div>
  </section>
</template>

<script setup lang="ts">
import { computed, onMounted, onUnmounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { ArrowLeft, ArrowRight } from '@element-plus/icons-vue'
import type { Publication, PublicationLink } from '../api/client'
import { useFeaturedPublications } from '../composables/useFeaturedPublications'
import { openExternalLink } from '../utils/links'
import { resolveExternalVideoURL, resolveVideoPagePath } from '../utils/videoLinks'

const router = useRouter()
const currentIndex = ref(0)
const publications = useFeaturedPublications()

const activePublication = computed(() => publications.value[currentIndex.value])

const handleImageError = (e: Event) => {
  (e.target as HTMLImageElement).src = '/publications/online.png'
}

const sortedLinks = (pub: Publication) => [...pub.links].sort((a, b) => a.sortOrder - b.sortOrder)

const openLink = (url: string) => openExternalLink(url)

const handleLinkClick = (link: PublicationLink) => {
  if (link.type === 'video') {
    const pagePath = resolveVideoPagePath(link)
    if (pagePath) {
      router.push(pagePath)
      return
    }
    const externalURL = resolveExternalVideoURL(link)
    if (externalURL) {
      openLink(externalURL)
    }
    return
  }

  if (link.url) {
    openLink(link.url)
  }
}

const nextSlide = () => {
  if (!publications.value.length) return
  currentIndex.value = (currentIndex.value + 1) % publications.value.length
}

const prevSlide = () => {
  if (!publications.value.length) return
  currentIndex.value = currentIndex.value === 0 ? publications.value.length - 1 : currentIndex.value - 1
}

const goToSlide = (index: number) => {
  currentIndex.value = index
}

let interval: ReturnType<typeof setInterval>

onMounted(() => {
  interval = setInterval(nextSlide, 7000)
})

onUnmounted(() => {
  clearInterval(interval)
})
</script>

<style scoped>
.research-spotlight {
  border-top: 1px solid #e4cfd7;
}

.spotlight-layout {
  display: grid;
  grid-template-columns: minmax(280px, 0.9fr) minmax(0, 1.1fr);
  min-height: 390px;
  border-bottom: 1px solid #e8edf1;
}

.paper-visual {
  align-items: center;
  background: #faf7f8;
  border-right: 1px solid #e8edf1;
  display: flex;
  justify-content: center;
  padding: 30px;
}

.paper-visual img {
  background: #fff;
  display: block;
  max-height: 300px;
  max-width: 100%;
  object-fit: contain;
}

.paper-info {
  display: flex;
  flex-direction: column;
  justify-content: center;
  padding: 34px 40px;
}

.paper-meta {
  color: #7d1231;
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
  font-size: 0.88rem;
  font-weight: 700;
  margin-bottom: 18px;
}

.paper-meta span {
  border: 1px solid #e4cfd7;
  padding: 5px 9px;
}

.paper-info h3 {
  color: #273445;
  font-size: 1.45rem;
  font-weight: 700;
  line-height: 1.45;
  margin: 0;
}

.authors {
  color: #7d1231;
  font-size: 0.98rem;
  font-weight: 650;
  line-height: 1.65;
  margin: 18px 0 0;
}

.venue {
  color: #566273;
  font-size: 0.98rem;
  line-height: 1.65;
  margin: 8px 0 0;
}

.paper-links {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
  margin-top: 24px;
}

.paper-link {
  background: #fff;
  border: 1px solid #7d1231;
  color: #7d1231;
  cursor: pointer;
  font-size: 0.92rem;
  font-weight: 650;
  padding: 8px 14px;
}

.paper-link:hover,
.paper-link:focus {
  background: #7d1231;
  color: #fff;
}

.carousel-controls {
  align-items: stretch;
  display: grid;
  grid-template-columns: 46px minmax(0, 1fr) 46px;
  border-bottom: 1px solid #e8edf1;
}

.nav-button {
  align-items: center;
  background: #fff;
  border: 0;
  color: #7d1231;
  cursor: pointer;
  display: flex;
  justify-content: center;
}

.nav-button:first-child {
  border-right: 1px solid #e8edf1;
}

.nav-button:last-child {
  border-left: 1px solid #e8edf1;
}

.nav-button:hover,
.nav-button:focus {
  background: #faf7f8;
}

.paper-tabs {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
}

.paper-tab {
  background: #fff;
  border: 0;
  border-right: 1px solid #e8edf1;
  cursor: pointer;
  min-height: 96px;
  padding: 14px;
  text-align: left;
}

.paper-tab:last-child {
  border-right: 0;
}

.paper-tab.active {
  background: #faf7f8;
  box-shadow: inset 0 4px 0 #7d1231;
}

.tab-index {
  color: #7d1231;
  display: block;
  font-size: 0.82rem;
  font-weight: 700;
  margin-bottom: 8px;
}

.tab-title {
  color: #566273;
  display: -webkit-box;
  font-size: 0.88rem;
  line-height: 1.45;
  overflow: hidden;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
}

.paper-tab.active .tab-title {
  color: #273445;
  font-weight: 650;
}

@media (max-width: 900px) {
  .spotlight-layout {
    grid-template-columns: 1fr;
  }

  .paper-visual {
    border-right: 0;
    border-bottom: 1px solid #e8edf1;
  }

  .paper-tabs {
    grid-template-columns: 1fr;
  }

  .paper-tab {
    min-height: auto;
    border-right: 0;
    border-bottom: 1px solid #e8edf1;
  }

  .paper-tab:last-child {
    border-bottom: 0;
  }
}

@media (max-width: 640px) {
  .paper-info {
    padding: 24px 18px;
  }

  .paper-info h3 {
    font-size: 1.2rem;
  }

  .carousel-controls {
    grid-template-columns: 40px minmax(0, 1fr) 40px;
  }
}
</style>
