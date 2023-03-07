## Non-Sudo

Create file `21-gaming-keyboard.rules` in `/etc/udev/rules.d`  
Update rules: `sudo udevadm control --reload-rules && sudo udevadm trigger`
```shell
SUBSYSTEM=="usb", ATTR{idVendor}=="0416", ATTR{idProduct}=="c345", TAG+="uaccess"
```