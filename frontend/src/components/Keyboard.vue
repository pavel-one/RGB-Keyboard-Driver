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
  <div style="height: 100%; display: flex; flex-direction: column; justify-content: center; align-items: center">
    <keyboard-demo></keyboard-demo>
    <div>
      <pre>
<!--        {{ data.keys }}-->
      </pre>
    </div>
  </div>
</template>