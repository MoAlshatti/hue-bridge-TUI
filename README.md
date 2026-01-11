![Philips Hue](https://a11ybadges.com/badge?logo=philipshue)  
![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![License](https://img.shields.io/badge/MIT-green?style=for-the-badge)
![TUI](https://img.shields.io/badge/UI-TUI-blueviolet?style=for-the-badge)

# Huecli 💡🖥️
A terminal UI for controlling philips hue lights via the terminal

![DEMO](img/tuidemo.gif)

## Table of contents
- [Installation](#installation) 
- [Usage](#usage)
- [Features](#features)
- [License](#license)

<h2 id="installation">Installation</h2>
<h3>Install via Go</h3>

  ```
  go install github.com/MoAlshatti/hue-bridge-TUI/cmd/huecli@latest
```
<h3>Install via homebrew (works for mac and linux)</h3>

  ```
  brew tap MoAlshatti/homebrew-tap
  brew install --cask huecli
```

<h2 id="usage">Usage</h2>


After installation, run:

```bash
huecli
```

It will guide you through the setup.  


<p align="center">
  <b>Quick walkthrough on YT</b><br><br>
  <a href="https://youtu.be/j0Z38CyYGIs">
    <img src="https://img.youtube.com/vi/j0Z38CyYGIs/maxresdefault.jpg" alt="Video guide" width="600">
  </a>
</p>

<h2 id="features">Features</h2>

#### Filtering 
 * Filtering lights and scenes based on groups
#### server-sent events
 * Syncs with changes made outside the app (from the main hue app for example)
#### multiple control options 
 * Control individual lights, as well as groups
#### vim-like keybinds
 * Supports vim keybinds, and arrows
#### Hue API V2 support
 * Uses the most recent version of the hue API


<h2 id="license">License ⚖️</h2>

[MIT](https://choosealicense.com/licenses/mit/)


