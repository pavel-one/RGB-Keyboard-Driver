<script lang="ts" setup>
import {keyboard} from "../../wailsjs/go/models";

interface KeyboardData {
  keys: Array<keyboard.Key> | null
}

const reactive = (target: KeyboardData) => console.log(target.keys);

reactive({
  keys: [
      new (keyboard.Key)
  ]
});
</script>

<template>
<!--  thanks https://codepen.io/irajsuhail/pen/mYMZVm-->
  <div class="keyboard-base">
    <div class="key">~</div>
    <div class="key">1</div>
    <div class="key">2</div>
    <div class="key">3</div>
    <div class="key">4</div>
    <div class="key">5</div>
    <div class="key">6</div>
    <div class="key">7</div>
    <div class="key">8</div>
    <div class="key">9</div>
    <div class="key">0</div>
    <div class="key">-</div>
    <div class="key">+</div>
    <div class="key delete">Delete</div>
    <div class="key tab">Tab</div>
    <div class="key">Q</div>
    <div class="key">w</div>
    <div class="key">E</div>
    <div class="key">R</div>
    <div class="key">T</div>
    <div class="key">Y</div>
    <div class="key">U</div>
    <div class="key">I</div>
    <div class="key">O</div>
    <div class="key">P</div>
    <div class="key">[</div>
    <div class="key">]</div>
    <div class="key backslash">\</div>
    <div class="key capslock">CapsLock</div>
    <div class="key">A</div>
    <div class="key">S</div>
    <div class="key">D</div>
    <div class="key">F</div>
    <div class="key">G</div>
    <div class="key">H</div>
    <div class="key">J</div>
    <div class="key">K</div>
    <div class="key">L</div>
    <div class="key">;</div>
    <div class="key">'</div>
    <div class="key return">Return</div>
    <div class="key leftshift">Shift</div>
    <div class="key">Z</div>
    <div class="key">X</div>
    <div class="key">C</div>
    <div class="key">V</div>
    <div class="key">B</div>
    <div class="key">N</div>
    <div class="key">M</div>
    <div class="key">,</div>
    <div class="key">.</div>
    <div class="key">/</div>
    <div class="key rightshift">Shift</div>
    <div class="key leftctrl">Ctrl</div>
    <div class="key">Alt</div>
    <div class="key command">Command</div>
    <div class="key space">Space</div>
    <div class="key command">command</div>
    <div class="key">Alt</div>
    <div class="key">Ctrl</div>
    <div class="key">Fn</div>
  </div>
</template>

<style lang="scss">

.keyboard-base {
  max-width: 1085px;
  padding: 20px;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  border-radius: 10px;
  display: grid;
  grid-template-columns: repeat(30, 30px);
  grid-template-rows: repeat(5, 60px);
  grid-gap: 5px;
}


.key {
  border: 2px solid black;
  border-radius: 5px;
  grid-column: span 2;
  font-size: 10px;
  text-align: center;
  padding-top: 17px;
}

.key:hover {
  border: 1px solid #eeeeee;

}

.delete {
  grid-column: span 4;
}

.tab {
  grid-column: span 3;
}

.backslash {
  grid-column: span 3;
}

.capslock {
  grid-column: span 4;
}

.return {
  grid-column: span 4;
}

.leftshift {
  grid-column: span 5;
}

.rightshift {
  grid-column: span 5;
}

.leftctrl {
  grid-column: span 3;
}

.command {
  grid-column: span 3;
  font-size: 14px;
}

.space {
  grid-column: span 13;
}
</style>

<script>

</script>