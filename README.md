# poecampain

Terminal based campaign guide for Path of Exile.

<p>
    <img src="img/poecampain.gif" alt="poecampain showcase">
</p>

## Features
* Move to the next step automatically when entering a new zone.
* Remember progress when exiting.

### Notes
* You will need [Go](https://go.dev/) to build the app.
* A [Nerd Font](https://www.nerdfonts.com) is required to render all text.
* This guide was made for my own use. There may be important information, directions, or optional content not included. Feel free to modify the guide for your own purpose.
* Mainly tested on Linux, but should work on all platforms.
* No other features planned at this time. (ex. POB import for gems)

## Config
### ~/.config/poecampain/config.yaml
`client: /path/to/Path of Exile/logs/Client.txt`

Default: `~/.steam/steam/steamapps/common/Path of Exile/logs/Client.txt`

The Client.txt file is used to read the current zone information.

## Keybinds
Manual navigation: `↑` `↓` `←` `→`

Quit: `q` or `ctrl+c`

Reset: `r` (same as `↑`)

## Screenshot
This is how I use the guide pinned on top of the game. Each step has a maximum of 5 lines.

<p>
    <img src="img/example.png" width="1000" alt="Show poecampaign pinned on top of the game">
</p>

* Terminal: Ghostty
* Font: JetBrainsMono Nerd Font
* Size: 14
* Theme: Catppuccin Mocha

For Hyprland: `windowrule = float on, pin on, size 620 140, opacity 0.7, match:title poecampain`

## Shout-out
Thanks to [Exile-UI](https://github.com/Lailloken/Exile-UI) and [Exile Leveling](https://heartofphos.github.io/exile-leveling/#/).

This guide uses information taken from both leveling guides.
