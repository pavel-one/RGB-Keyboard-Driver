<script lang="ts" setup>
import {keyboard} from "../../wailsjs/go/models";

const props = defineProps<{
  kKey: keyboard.Key
}>()
</script>

<template>
  <n-button :style="{width: '82px', background: 'rgb('+kKey.red.toString()+','+kKey.green.toString()+','+kKey.blue.toString()+')'}">
    {{ kKey.name }}
  </n-button>
</template>

