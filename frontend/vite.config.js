import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  define: {
    'api': JSON.stringify('https://5101-idx-fhxparsergit-1729582237031.cluster-23wp6v3w4jhzmwncf7crloq3kw.cloudworkstations.dev'),
  }
})
