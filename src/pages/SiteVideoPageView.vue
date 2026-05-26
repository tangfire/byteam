<template>
  <SiteVideoPage
    :key="slug"
    :slug="slug"
    fallback-title="Publication Video"
    fallback-video=""
  />
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import SiteVideoPage from '../components/SiteVideoPage.vue'

const route = useRoute()
const slug = computed(() => String(route.params.slug || ''))
</script>
