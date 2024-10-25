import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  define: {
    'api': JSON.stringify('https://5101-idx-fhxparser-1729854320246.cluster-qtqwjj3wgzff6uxtk26wj7fzq6.cloudworkstations.dev'),    
  }
})
