local wez = require("wezterm")
local config = wez.config_builder()
config.enable_scroll_bar = false
config.font = wez.font({
    family = "IosevkaCustom",
    weight = "Regular",
    harfbuzz_features = { 'calt=0', 'clig=0', 'liga=0' },
})
config.font_size = 11.0 --  13.5
-- https://github.com/folke/tokyonight.nvim/blob/main/extras/wezterm/tokyonight_moon.toml
config.colors = {
    foreground = "#c0caf5",
    background = "#1a1b26",
    cursor_bg = "#c0caf5",
    cursor_border = "#c0caf5",
    cursor_fg = "#1a1b26",
    selection_bg = "#283457",
    selection_fg = "#c0caf5",
    split = "#7aa2f7",
    compose_cursor = "#ff9e64",
    scrollbar_thumb = "#292e42",
    ansi = { "#15161e", "#f7768e", "#9ece6a", "#e0af68", "#7aa2f7", "#bb9af7", "#7dcfff", "#a9b1d6" },
    brights = { "#414868", "#ff899d", "#9fe044", "#faba4a", "#8db0ff", "#c7a9ff", "#a4daff", "#c0caf5" },
}
config.enable_tab_bar = false
config.audible_bell = "Disabled"
config.enable_scroll_bar = false
config.enable_wayland = false
config.exit_behavior = "Close"
config.scrollback_lines = 1000
config.window_padding = {
    left   = 0,
    right  = 0,
    top    = 0,
    bottom = 0,
}

return config
