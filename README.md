# Conky_CZ_BTC_Spotify

Personal conky inspired by:</br>
* https://github.com/Calosis/Graffias-Custom</br>
* https://github.com/rayzr522/now-clocking </br>
* https://github.com/matteocasonato/conky-cryptoWidget</br>
  
## Final look
  ![Snímek obrazovky z 2022-06-21 11-57-27](https://user-images.githubusercontent.com/63755464/174772940-b865f77f-8dd7-4bb8-a304-cf8017aff76e.png)

## Dependencies
This script require the follwing modules/libraries: 
* Conky module; which can be installed using:
```
sudo apt install conky-all
```
### Packages
* Conky module; which can be installed using:
- [conky](https://github.com/brndnmtthws/conky/)
```
sudo apt install conky-all
```
- [ffmpeg](https://www.ffmpeg.org/)
- [playerctl](https://github.com/altdesktop/playerctl)
```
sudo apt install playerctl
```

```
# Ubuntu
$ sudo apt install conky ffmpeg playerctl
# Arch
$ sudo pacman -S conky ffmpeg playerctl
```

## Installation

1. Install all required [packages](#packages):
```bash
# Ubuntu
$ sudo apt install conky ffmpeg playerctl
# Arch
$ sudo pacman -S conky ffmpeg playerctl
```
2. Clone the repo:
```bash
$ git clone git@github.com:Nevoral/Conky_CZ_BTC_Spotify.git
```
3. Run the `start.sh` script to start the widget (forks to background):
```bash
$ path/to/Conky_CZ_BTC_Spotify/sh start.sh
```
