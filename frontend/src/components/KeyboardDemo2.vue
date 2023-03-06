<!--https://codepen.io/iamsainikhil/pen/gdYmOQ-->
<template>
  <main
      style="--keyboard-bg-color:rgba(0, 0, 0, 0.9); --key-bg-color:#353535; --key-color:#f8f8f8; --lock-on-color:#008996;">
    <div class="keyboard-container">
      <div class="keyboard">
        <!-- 1st row -->
        <div class="row">
          <div class="group">
            <kbd data-key="27" class="key">Esc</kbd>
            <div class="mini-group">
              <kbd data-key="112" class="key"
              >F1<i class="fas fa-home"></i
              ></kbd>
              <kbd data-key="113" class="key"
              >F2<i class="fas fa-envelope"></i
              ></kbd>
              <kbd data-key="114" class="key"
              >F3<i class="fas fa-search"></i
              ></kbd>
              <kbd data-key="115" class="key"
              >F4<i class="fas fa-copy"></i
              ></kbd>
            </div>
            <div class="mini-group">
              <kbd data-key="116" class="key"
              >F5<i class="fas fa-play"><i class="fas fa-pause"></i></i
              ></kbd>
              <kbd data-key="117" class="key"
              >F6<i class="fas fa-stop"></i
              ></kbd>
              <kbd data-key="118" class="key"
              >F7<i class="fas fa-fast-backward"></i
              ></kbd>
              <kbd data-key="119" class="key"
              >F8<i class="fas fa-fast-forward"></i
              ></kbd>
            </div>
            <div class="mini-group">
              <kbd data-key="120" class="key"
              >F9<i class="fas fa-volume-down"></i
              ></kbd>
              <kbd data-key="121" class="key"
              >F10<i class="fas fa-volume-up"></i
              ></kbd>
              <kbd data-key="122" class="key"
              >F11<i class="fas fa-volume-off"></i
              ></kbd>
              <kbd data-key="123" class="key"
              >F12<i class="fas fa-sun"></i
              ></kbd>
            </div>
          </div>
          <div class="group locks">
            <p class="lock"><span class="light-key off" id="nlk">N</span></p>
            <p class="lock"><span class="light-key off" id="clk">C</span></p>
            <p class="lock"><span class="light-key off" id="slk">S</span></p>
          </div>
          <div class="right">
            <div class="group">
              <kbd data-key="-1" class="key">Print Screen</kbd>
              <kbd data-key="145" class="key">Scroll Lock</kbd>
              <kbd data-key="19" class="key">Pause Break</kbd>
            </div>
          </div>
        </div>

        <!-- 2nd row -->
        <div class="row">
          <div class="group">
            <kbd data-key="192" class="key large">~<i>`</i></kbd>
            <kbd data-key="49" class="key">!<span>1</span></kbd>
            <kbd data-key="50" class="key">@<span>2</span></kbd>
            <kbd data-key="51" class="key">#<span>3</span></kbd>
            <kbd data-key="52" class="key">$<span>4</span></kbd>
            <kbd data-key="53" class="key">%<span>5</span></kbd>
            <kbd data-key="54" class="key">^<span>6</span></kbd>
            <kbd data-key="55" class="key">&<span>7</span></kbd>
            <kbd data-key="56" class="key">*<span>8</span></kbd>
            <kbd data-key="57" class="key">(<span>9</span></kbd>
            <kbd data-key="48" class="key">)<span>0</span></kbd>
            <kbd data-key="189" class="key">-<i class="underscore">_</i></kbd>
            <kbd data-key="187" class="key">+<span>=</span></kbd>
            <kbd data-key="8" class="key backspace">Backspace</kbd>
          </div>
          <div class="group">
            <kbd data-key="45" class="key">Insert</kbd>
            <kbd data-key="36" class="key">Home</kbd>
            <kbd data-key="33" class="key">Page Up</kbd>
          </div>
          <div class="group">
            <kbd data-key="144" class="key">Num Lock</kbd>
            <kbd data-key="111" class="key">/</kbd>
            <kbd data-key="106" class="key">*</kbd>
            <kbd data-key="109" class="key">-</kbd>
          </div>
        </div>

        <!-- 3rd row -->
        <div class="row">
          <div class="group">
            <kbd data-key="9" class="key tab">Tab⭾</kbd>
            <kbd data-key="81" class="key">Q</kbd>
            <kbd data-key="87" class="key">W</kbd>
            <kbd data-key="69" class="key">E</kbd>
            <kbd data-key="82" class="key">R</kbd>
            <kbd data-key="84" class="key">T</kbd>
            <kbd data-key="89" class="key">Y</kbd>
            <kbd data-key="85" class="key">U</kbd>
            <kbd data-key="73" class="key">I</kbd>
            <kbd data-key="79" class="key">O</kbd>
            <kbd data-key="80" class="key">P</kbd>
            <kbd data-key="219" class="key">{<span>[</span></kbd>
            <kbd data-key="221" class="key">}<span>]</span></kbd>
            <kbd data-key="220" class="key">|<span>\</span></kbd>
          </div>
          <div class="group">
            <kbd data-key="46" class="key">Delete</kbd>
            <kbd data-key="35" class="key">End</kbd>
            <kbd data-key="34" class="key">Page Down</kbd>
          </div>
          <div class="group">
            <kbd data-key="103" class="key">7<span>Home</span></kbd>
            <kbd data-key="104" class="key num-special"
            >8<i class="fas fa-long-arrow-alt-up"></i
            ></kbd>
            <kbd data-key="105" class="key">9<span>PgUp</span></kbd>
            <kbd data-key="107" class="key num-large">+</kbd>
          </div>
        </div>

        <!-- 4th row -->
        <div class="row">
          <div class="group">
            <kbd data-key="20" class="key tab">CapsLock</kbd>
            <kbd data-key="65" class="key">A</kbd>
            <kbd data-key="83" class="key">S</kbd>
            <kbd data-key="68" class="key">D</kbd>
            <kbd data-key="70" class="key">F</kbd>
            <kbd data-key="71" class="key">G</kbd>
            <kbd data-key="72" class="key">H</kbd>
            <kbd data-key="74" class="key">J</kbd>
            <kbd data-key="75" class="key">K</kbd>
            <kbd data-key="76" class="key">L</kbd>
            <kbd data-key="186" class="key">:<span>;</span></kbd>
            <kbd data-key="222" class="key">"<span>'</span></kbd>
            <kbd data-key="13" class="key enter">Enter ↩</kbd>
          </div>
          <div class="empty-group"></div>
          <div class="group">
            <kbd data-key="100" class="key"
            >4<i class="fas fa-long-arrow-alt-left"></i
            ></kbd>
            <kbd data-key="101" class="key num-special">5<i>_</i></kbd>
            <kbd data-key="102" class="key"
            >6<i class="fas fa-long-arrow-alt-right"></i
            ></kbd>
            <p class="key key-empty"></p>
          </div>
        </div>

        <!-- 5th row -->
        <div class="row">
          <div class="group">
            <kbd data-key="16" class="key shift l-s"
            ><span>⇧</span>&nbsp;Shift</kbd
            >
            <kbd data-key="90" class="key">Z</kbd>
            <kbd data-key="88" class="key">X</kbd>
            <kbd data-key="67" class="key">C</kbd>
            <kbd data-key="86" class="key">V</kbd>
            <kbd data-key="66" class="key">B</kbd>
            <kbd data-key="78" class="key">N</kbd>
            <kbd data-key="77" class="key">M</kbd>
            <kbd data-key="188" class="key">&lt;<span>,</span></kbd>
            <kbd data-key="190" class="key">&gt;<span>.</span></kbd>
            <kbd data-key="191" class="key">?<span>/</span></kbd>
            <kbd data-key="16" class="key shift r-s"
            ><span>⇧</span>&nbsp;Shift</kbd
            >
          </div>
          <div class="group">
            <p class="key key-empty"></p>
            <kbd data-key="38" class="key"
            ><i class="fas fa-arrow-up"></i
            ></kbd>
            <p class="key key-empty"></p>
          </div>
          <div class="group">
            <kbd data-key="97" class="key">1<span>End</span></kbd>
            <kbd data-key="98" class="key"
            >2<i class="fas fa-long-arrow-alt-down"></i
            ></kbd>
            <kbd data-key="99" class="key">3<span>PdDn</span></kbd>
            <kbd data-key="13" class="key num-large">Enter</kbd>
          </div>
        </div>

        <!-- 6th row -->
        <div class="row">
          <div class="group">
            <kbd data-key="17" class="key">Ctrl</kbd>
            <kbd data-key="91" class="key"
            ><i class="fab fa-windows"></i
            ></kbd>
            <kbd data-key="18" class="key">Alt</kbd>
            <kbd data-key="32" class="key spacebar">TechGeek</kbd>
            <kbd data-key="18" class="key">Alt</kbd>
            <kbd data-key="-2" class="key">Fn</kbd>
            <kbd data-key="93" class="key c-menu"
            ><i class="fas fa-bars"></i
            ></kbd>
            <kbd data-key="17" class="key">Ctrl</kbd>
          </div>
          <div class="group">
            <kbd data-key="37" class="key"
            ><i class="fas fa-arrow-left"></i
            ></kbd>
            <kbd data-key="40" class="key"
            ><i class="fas fa-arrow-down"></i
            ></kbd>
            <kbd data-key="39" class="key"
            ><i class="fas fa-arrow-right"></i
            ></kbd>
          </div>
          <div class="group">
            <kbd data-key="96" class="key num-ins">0<span>Ins</span></kbd>
            <kbd data-key="110" class="key">.<span>Del</span></kbd>
            <p class="key key-empty"></p>
          </div>
        </div>
      </div>
    </div>

    <!-- audio -->
    <audio id="keyAudio">
      <source src="https://raw.githubusercontent.com/iamsainikhil/keyboard/master/assets/button-click.mp3"
              type="audio/mpeg"/>
      Your browser does not support the audio element.
    </audio>

    <div class="color-themes">
      <h2 class="title">Keyboard Theme</h2>
      <p class="sub-title">
        Select any default theme below or generate a custom theme to switch
        colors
      </p>
      <div class="button-wrapper">
        <button class="button" id="t-0">Default</button>
        <button class="button" id="t-1">Light</button>
        <button class="button" id="t-2">Dark</button>
        <button class="button" id="t-3">Rose Gold</button>
        <button class="button" id="t-4">Custom</button>
      </div>
      <div class="custom-theme" id="c-t">
        <h3 class="title">
          Set below 4 properties to develop your custom theme
        </h3>
        <div class="cp-wrapper">
          <button class="button" id="kbc">Keyboard Background Color</button>
          <button class="button" id="keybc">Key Background Color</button>
          <button class="button" id="kc">Key Color</button>
          <button class="button" id="lkc">NCS Key Color</button>
        </div>
      </div>
    </div>

    <div class="notes-wrapper">
      <h2 class="title">Information</h2>
      <button class="button" id="info">Learn More</button>
    </div>

    <div class="notes">
      <p>
        <b>Fun Fact:</b> Press any key on this virtual keyboard or on your
        keyboard to hear keypress sound. Try playing with few keys like Caps
        Lock, Num Lock, Scroll Lock, and F12 and moreover there won't be any
        keypress sound for these keys due to let users know about their unique
        functionality.
      </p>
      <p>
        <b>Note:</b> You will experience that whenever a key is pressed, there
        is animation that makes key is pressed and when you press any other
        key or the same key, then that key is back to normal state. This is
        normal behavior in this virtual keyboard context. However, in reality
        key will popped in-out very quickly.
      </p>
    </div>
  </main>

  <footer class="footer">
    <p class="tech-love">Made with 🖤 using JavaScript</p>
    <p>
      &copy;2018 <span id="year"></span>
      <a
          href="https://iamsainikhil.github.io"
          target="_blank"
          id="footer-name"
      >Sai&nbsp;Nikhil</a
      >
    </p>
  </footer>
</template>

<style lang="scss" scoped>
.key {
  background-color: var(--key-bg-color);
  color: var(--key-color);
  font-family: var(--font-m);
  font-size: 12px;
  width: 26px;
  height: 20px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  text-align: center;
  margin: 0 4px;
  padding: 0 2px;
  border-radius: 10px;
  box-shadow: 0px 5px 0px rgba(0, 0, 0, 0.2);
}

.key:hover,
.key-press {
  cursor: pointer;
  box-shadow: 0px 2px 0px rgba(0, 0, 0, 0.2);
  transform: translateY(3px);
  -webkit-transform: translateY(2px);
  transition: all 0.1s ease-in-out;
  -webkit-transition: all 0.1s ease-in-out;
  -moz-transition: all 0.1s ease-in-out;
  -o-transition: all 0.1s ease-in-out;
}

.key-empty {
  background-color: transparent;
  box-shadow: none;
}

.key i {
  margin-top: 1px;
}

.key .fas {
  font-size: 11px;
}

.key.large {
  padding-top: 5px;
  font-size: 18px;
}

.key.large i {
  font-size: 18px;
  margin-top: 0px;
}

i.underscore {
  margin-top: -8px;
}

/* special keys */

.backspace,
.tab {
  width: 80px;
}

.enter {
  width: 100px;
}

.spacebar {
  width: 405px;
  font-size: 18px;
}

.shift {
  display: flex;
  flex-flow: row nowrap;
}

.l-s {
  width: 95px;
}

.r-s {
  width: 140px;
}

.c-menu i {
  border: 1px solid #f5f5f5;
  padding: 3px;
}

.num-special {
  justify-content: flex-start;
  padding-top: 5px;
}

.num-ins {
  width: 100px;
}

.num-large {
  margin-bottom: -58px;
  height: 100px;
}

/* caps lock, num lock, and scroll lock indicators */

.locks {
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  width: 200px;
  margin-top: 5px;
  margin-right: 10px;
  background-color: var(--key-bg-color);
  border-radius: 5px;
}

.lock {
  margin: 5px 22px;
  width: 135px;
}

.light-key {
  padding: 3px 5px;
  border-radius: 5px;
}

.off {
  color: var(--key-color);
  border: 1px solid var(--key-color);
}

.on {
  color: var(--lock-on-color);
  border: 1px solid var(--lock-on-color);
}

/* keyboard layout */
.keyboard-container {
  width: 100%;
  overflow-x: scroll;
}

.keyboard {
  background-color: var(--keyboard-bg-color);
  width: 1250px;
  margin: 75px auto;
  padding: 20px 10px;
  border-radius: 15px;
  display: grid;
  grid-template-columns: 1fr;
  grid-template-rows: 1fr;
  justify-content: center;
  box-shadow: 0 27px 24px 0 rgba(0, 0, 0, 0.2),
  0 40px 77px 0 rgba(0, 0, 0, 0.22);
}

.row {
  display: grid;
  grid-template-columns: auto 0.2fr 0.2fr;
  grid-gap: 25px;
  padding: 5px;
  margin-top: 5px;
}

.group,
.mini-group {
  display: flex;
  flex-flow: row nowrap;
  justify-content: flex-start;
  align-items: center;
}

.mini-group {
  margin-left: 28px;
}

.right {
  display: flex;
  justify-content: flex-end;
}

/* color themes */

.color-themes {
  text-align: center;
  font-family: var(--font-m);
}

.button-wrapper .button {
  width: 150px;
}

.button {
  height: 45px;
  background-color: transparent;
  border: 1px solid #5b5b5b;
  border-radius: 25px;
  color: #333;
  margin: 5px 15px;
  font-family: var(--font-b);
  font-size: 16px;
}

.button:hover {
  cursor: pointer;
  background-color: #5b5b5b;
  color: #fff;
  transition: all 0.3s ease-in-out;
  -webkit-transition: all 0.3s ease-in-out;
  -moz-transition: all 0.3s ease-in-out;
  box-shadow: 0 12px 15px 0 rgba(0, 0, 0, 0.24),
  0 17px 50px 0 rgba(0, 0, 0, 0.19);
}

/* custom theme */

.custom-theme {
  padding-top: 50px;
  display: none;
}

.cp-wrapper {
  display: flex;
  flex-flow: row wrap;
  justify-content: center;
  align-items: center;
}

.cp-wrapper .button {
  margin: 20px 15px;
  padding: 5px 25px;
}

.title {
  font-family: var(--font-b);
  text-align: center;
}

.sub-title {
  font-family: var(--font-m);
  letter-spacing: 1px;
}

/* notes */

.notes-wrapper {
  text-align: center;
  padding-top: 40px;
}

.notes-wrapper .title {
  line-height: 0;
}

.notes-wrapper b {
  font-family: var(--font-b);
}

.notes-wrapper .button {
  padding: 5px 25px;
  width: 150px;
}

.notes {
  max-width: 960px;
  margin: 0 auto;
  font-family: var(--font-m);
  letter-spacing: 1px;
  display: none;
}

.notes p {
  background-color: #eeeeee;
  border-radius: 20px;
  padding: 20px 40px;
  line-height: 1.5;
  margin: 20px 0;
}

.show-notes {
  display: block;
}

/* footer */

.footer {
  text-align: center;
  background: linear-gradient(270deg, #f36b38, #f2354d);
  background-size: 100% 100%;
  -webkit-animation: gradientFill 8s ease infinite;
  -moz-animation: gradientFill 8s ease infinite;
  animation: gradientFill 8s ease infinite;
  margin: 2rem 0 0 0;
  padding: 0;
  font-family: var(--font-b);
}

.tech-love {
  text-align: center;
  margin-bottom: 0;
}

.footer a {
  position: relative;
  text-decoration: none;
  color: #fff;
}

.footer a::before {
  content: '';
  position: absolute;
  width: 100%;
  height: 2px;
  bottom: -2px;
  left: 0;
  background: #000;
  visibility: hidden;
  border-radius: 10px;
  -webkit-transform: scaleX(0);
  transform: scaleX(0);
  -webkit-transition: all 0.3s cubic-bezier(0.25, 0.46, 0.45, 0.94);
  transition: all 0.3s cubic-bezier(0.25, 0.46, 0.45, 0.94);
}

.footer a:hover {
  color: #000;
}

.footer a:hover::before {
  visibility: visible;
  -webkit-transform: scaleX(1);
  transform: scaleX(1);
}

@-webkit-keyframes gradientFill {
  0% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
  100% {
    background-position: 0% 50%;
  }
}

@-moz-keyframes gradientFill {
  0% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
  100% {
    background-position: 0% 50%;
  }
}

@keyframes gradientFill {
  0% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
  100% {
    background-position: 0% 50%;
  }
}

</style>