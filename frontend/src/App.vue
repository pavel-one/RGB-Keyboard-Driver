<script lang="ts" setup>
import Keyboard from "./components/Keyboard.vue";
import {darkTheme} from 'naive-ui'
import {reactive} from "vue";
import ThemeSwitcher from "./components/ThemeSwitcher.vue";
import RandomButton from "./components/RandomButton.vue";
import {ReloadCircle} from '@vicons/ionicons5'
import {GetConnectedStatus} from "../wailsjs/go/main/App";

let data = reactive({
  theme: darkTheme as any,
  background: '#151515',
  connected: false
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

function updateStatus() {
  GetConnectedStatus().then(result => {
    data.connected = result
  })
}

updateStatus()
setInterval(() => {
  updateStatus()
}, 2000)

</script>

<template>
  <n-config-provider :style="{background: data.background, position: 'relative'}" :theme="data.theme">
    <n-message-provider v-if="!data.connected">
      <n-notification-provider>
        <n-dialog-provider>
          <n-result class="error-layout" status="error" title="Error" description="Keyboard not connected">
            <template #footer>
              <n-button @click="updateStatus" size="medium">
                <n-icon style="margin-right: 10px">
                  <reload-circle />
                </n-icon>
                Check it!
              </n-button>
            </template>
          </n-result>
        </n-dialog-provider>
      </n-notification-provider>
    </n-message-provider>
    <n-layout v-else>
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

<style lang="scss" scoped>
.n-layout {
  height: 100%;
}
.error-layout {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
</style>

