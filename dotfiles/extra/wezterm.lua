-- INSTALLDIR: $HOME/.config/wezterm/wezterm.lua

local wez = require("wezterm")
local config = wez.config_builder()

--https://larsenwork.com/monoid/
config.font = wez.font({ family = "Monoid HalfLoose"})
config.font_size = 12.0

config.enable_scroll_bar = false
config.enable_tab_bar = false
config.audible_bell = "Disabled"
config.enable_scroll_bar = false
config.enable_wayland = false
config.exit_behavior = "Close"
config.scrollback_lines = 5000
config.window_padding = {
    left   = 0,
    right  = 0,
    top    = 0,
    bottom = 0,
}

return config
