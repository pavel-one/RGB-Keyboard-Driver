import {createApp} from 'vue'
import App from './App.vue'
import Naive from 'naive-ui'
import vue from "@vitejs/plugin-vue"

const app = createApp(App)
app.use(Naive)

app.mount('#app')
