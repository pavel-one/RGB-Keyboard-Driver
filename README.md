## Non-Sudo


Create group
```shell
sudo groupadd plugdev
```
Create file `0-usb-main.rules` in `/etc/udev/rules.d`
```shell
SUBSYSTEM=="usb", ATTR{idVendor}=="0416", ATTR{idProduct}=="c345", GROUP="plugdev", TAG+="uaccess"
```