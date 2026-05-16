<template>
  <div class="min-h-screen bg-gradient-to-b from-sky-50 via-white to-stone-50">
    <Navbar />

    <section class="max-w-5xl mx-auto px-4 sm:px-6 lg:px-8 pt-24 sm:pt-28 pb-12 lg:pb-16">
      <div class="text-center mb-12">
        <div class="flex items-center justify-center gap-3 text-brand-600 text-xs font-semibold tracking-widest uppercase mb-4">
          <span class="w-8 h-px bg-brand-400"></span>
          {{ t.nav_news }}
          <span class="w-8 h-px bg-brand-400"></span>
        </div>
        <h1 class="text-3xl md:text-4xl font-bold text-stone-900 mb-3">{{ t.news_title }}</h1>
        <p class="text-stone-500 text-lg max-w-xl mx-auto">{{ t.news_subtitle }}</p>
      </div>

      <div v-if="loading" class="text-center py-20">
        <div class="inline-block w-10 h-10 border-[3px] border-brand-200 border-t-brand-600 rounded-full animate-spin"></div>
        <p class="mt-4 text-stone-400 font-medium">{{ t.news_loading }}</p>
      </div>

      <div v-else-if="posts.length === 0" class="text-center py-20">
        <div class="w-20 h-20 bg-stone-100 rounded-2xl flex items-center justify-center mx-auto mb-4">
          <svg class="w-10 h-10 text-stone-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
            <path stroke-linecap="round" stroke-linejoin="round" d="M12 7.5h1.5m-1.5 3h1.5m-7.5 3h7.5m-7.5 3h7.5m3-9h3.375c.621 0 1.125.504 1.125 1.125V18a2.25 2.25 0 01-2.25 2.25M16.5 7.5V18a2.25 2.25 0 002.25 2.25M16.5 7.5V4.875c0-.621-.504-1.125-1.125-1.125H4.125C3.504 3.75 3 4.254 3 4.875V18a2.25 2.25 0 002.25 2.25h13.5M6 7.5h3v3H6v-3z" />
          </svg>
        </div>
        <p class="text-stone-400 text-lg font-medium">{{ t.news_empty }}</p>
      </div>

      <div v-else class="grid sm:grid-cols-2 gap-6">
        <article
          v-for="post in posts"
          :key="post.id"
          class="bg-white rounded-3xl shadow-sm border border-stone-100 overflow-hidden hover:shadow-md hover:-translate-y-0.5 transition-all duration-300"
        >
          <!-- Multiple images carousel (if images exist) -->
          <div v-if="post.images && post.images.length > 0" class="h-52 overflow-hidden relative">
            <img :src="post.images[0].image_path" :alt="post.title" class="w-full h-full object-cover" />
            <span v-if="post.images.length > 1" class="absolute bottom-2 right-2 bg-black/50 text-white text-[10px] font-bold px-2 py-0.5 rounded-full">
              +{{ post.images.length - 1 }}
            </span>
          </div>
          <!-- Fallback single image -->
          <div v-else-if="post.image_path" class="h-52 overflow-hidden">
            <img :src="post.image_path" :alt="post.title" class="w-full h-full object-cover" />
          </div>
          <!-- Video -->
          <div v-else-if="post.video_url" class="aspect-video">
            <iframe
              :src="toEmbedUrl(post.video_url)"
              class="w-full h-full"
              frameborder="0"
              allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
              allowfullscreen
            ></iframe>
          </div>

          <div class="p-5 sm:p-6">
            <p class="text-xs text-stone-400 mb-2">{{ formatDate(post.created_at) }}</p>
            <h2 class="text-base sm:text-lg font-bold text-stone-900 mb-2 leading-snug">{{ post.title }}</h2>
            <p v-if="post.description" class="text-stone-500 text-sm leading-relaxed line-clamp-4">{{ post.description }}</p>

            <!-- Additional images grid -->
            <div v-if="post.images && post.images.length > 1" class="grid grid-cols-3 gap-1.5 mt-3">
              <div
                v-for="img in post.images.slice(1, 4)"
                :key="img.id"
                class="aspect-square rounded-lg overflow-hidden"
              >
                <img :src="img.image_path" class="w-full h-full object-cover" />
              </div>
            </div>
          </div>
        </article>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import axios from 'axios'
import Navbar from '../components/Navbar.vue'
import { useLangStore } from '../stores/lang'

const langStore = useLangStore()
const t = computed(() => langStore.t)

const posts = ref([])
const loading = ref(true)

function formatDate(val) {
  return new Date(val).toLocaleDateString(langStore.current === 'uz' ? 'uz-UZ' : 'ru-RU', {
    year: 'numeric', month: 'long', day: 'numeric'
  })
}

function toEmbedUrl(url) {
  if (!url) return ''
  const match = url.match(/(?:v=|youtu\.be\/)([A-Za-z0-9_-]{11})/)
  if (match) return `https://www.youtube.com/embed/${match[1]}`
  return url
}

onMounted(async () => {
  try {
    const res = await axios.get('/api/news')
    posts.value = res.data || []
  } finally {
    loading.value = false
  }
})
</script>
