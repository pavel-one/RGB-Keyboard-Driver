<script lang="ts" setup>
import {reactive} from 'vue'
import {GetKeyboardKeys} from '../../wailsjs/go/main/App'
import {keyboard} from "../../wailsjs/go/models";
import Key from "./Key.vue";
import KeyboardDemo from "./KeyboardDemo.vue";
import KeyboardDemo2 from "./KeyboardDemo2.vue";

interface KeyboardData {
  keys: Array<keyboard.Key> | null;
}

const data: KeyboardData = reactive({
  keys: null
})

setInterval(() => {
  GetKeyboardKeys().then(result => {
    data.keys = result
  })
}, 1000)
</script>

<template>
  <div>
    <div v-if="data.keys !== null">
      <n-space>
        <key :k-key="data.keys[0]"></key>
        <key style="margin-left: 20px" :k-key="data.keys[1]"></key>
        <key :k-key="data.keys[2]"></key>
        <key :k-key="data.keys[3]"></key>
        <key :k-key="data.keys[4]"></key>

        <key style="margin-left: 20px" :k-key="data.keys[5]"></key>
        <key :k-key="data.keys[6]"></key>
        <key :k-key="data.keys[7]"></key>
        <key :k-key="data.keys[8]"></key>

        <key style="margin-left: 20px" :k-key="data.keys[9]"></key>
        <key :k-key="data.keys[10]"></key>
        <key :k-key="data.keys[11]"></key>
        <key :k-key="data.keys[12]"></key>

        <key style="margin-left: 20px" :k-key="data.keys[13]"></key>
        <key :k-key="data.keys[14]"></key>
        <key :k-key="data.keys[15]"></key>
      </n-space>
      <n-space style="margin-top: 10px">
        <key :k-key="data.keys[16]"></key>
        <key :k-key="data.keys[17]"></key>
        <key :k-key="data.keys[18]"></key>
        <key :k-key="data.keys[19]"></key>
        <key :k-key="data.keys[20]"></key>
        <key :k-key="data.keys[21]"></key>
        <key :k-key="data.keys[22]"></key>
        <key :k-key="data.keys[23]"></key>
        <key :k-key="data.keys[24]"></key>
        <key :k-key="data.keys[25]"></key>
        <key :k-key="data.keys[26]"></key>
        <key :k-key="data.keys[27]"></key>
        <key :k-key="data.keys[28]"></key>
        <key :k-key="data.keys[29]"></key>

        <key style="margin-left: 20px" :k-key="data.keys[30]"></key>
        <key :k-key="data.keys[31]"></key>
        <key :k-key="data.keys[32]"></key>
      </n-space>
    </div>
    <keyboard-demo></keyboard-demo>
    <div>
      <pre>
        {{ data.keys }}
      </pre>
    </div>
  </div>
</template>