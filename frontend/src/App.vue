<script lang="ts" setup>
import Keyboard from "./components/Keyboard.vue";
import {darkTheme} from 'naive-ui'
import {reactive} from "vue";
import ThemeSwitcher from "./components/ThemeSwitcher.vue";
import RandomButton from "./components/RandomButton.vue";

let data = reactive({
  theme: darkTheme as any,
  background: '#151515'
})

function changeTheme(dark: boolean) {
  if (dark) {
    data.theme = darkTheme
    data.background = '#151515'
    return
  }

  data.theme = null
  data.background = '#f5f5f5'
}

</script>

<template>
  <n-config-provider :style="{background: data.background, position: 'relative'}" :theme="data.theme">
    <n-layout>
      <n-layout-header>Dark Project KD87a</n-layout-header>
      <n-layout has-sider>
        <n-layout-sider content-style="padding: 24px;">
          <n-button>Test</n-button>
        </n-layout-sider>
        <n-layout-content content-style="padding: 24px;">
          <keyboard></keyboard>
        </n-layout-content>
      </n-layout>
      <n-layout-footer>footer</n-layout-footer>
    </n-layout>

    <theme-switcher @change="changeTheme"></theme-switcher>
    <random-button></random-button>


  </n-config-provider>
</template>

<style lang="scss">
.n-layout {
  height: 100%;
}
</style>

