<script lang="ts" setup>
import {reactive} from 'vue'
import {GetKeyboardMatrix} from '../../wailsjs/go/main/App'
import {keyboard} from "../../wailsjs/go/models";


const data = reactive({
  keys: [] as any,
})

setInterval(() => {
  GetKeyboardMatrix().then(result => {
    data.keys = result
  })
}, 1000)
</script>

<template>
  <div>
      <div v-for="item in [0, 1, 2, 3, 4, 5]">
        <n-space>
          <div v-for="k in data.keys[item]">
            <n-button v-if="k != null"  :style="{
            background: 'rgb('+k.red.toString()+','+k.green.toString()+','+k.blue.toString()+')'
          }">{{ k.name }}</n-button>
          </div>

        </n-space>
      </div>
  </div>
</template>