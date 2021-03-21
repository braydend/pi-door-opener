# Pi Door Opener

Raspberry Pi powered garage door opener.

## Install instructions

1. Flash an SD card with Raspbian lite
1a. Configure wifi and ssh for use in headless mode
2. Boot Pi and SSH in. Run `sudo apt-get update`.
3. Install GO. I usually use [these steps](https://www.e-tinkers.com/2019/06/better-way-to-install-golang-go-on-raspberry-pi/)
4. Install git with `sudo apt install git`
5. Clone this repo to `/home/pi/git/pi-door-opener`
6. Register the service with systemd. `sudo cp /home/pi/git/pi-door-opener/garageDoor.service /lib/systemd/system/` then `sudo systemctl enable garageDoor.service`