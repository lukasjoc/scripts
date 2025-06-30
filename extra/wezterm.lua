-- INSTALLDIR: $HOME/.config/wezterm/wezterm.lua

local wez = require("wezterm")
local config = wez.config_builder()
config.enable_scroll_bar = false

-- TODO: fallback font and size for different systems
config.font = wez.font({
    family = "Hack",
    weight = "Regular",
    harfbuzz_features = { 'calt=0', 'clig=0', 'liga=0' },
})
config.font_size = 14.0 --  13.5
-- https://github.com/folke/tokyonight.nvim/blob/main/extras/wezterm/tokyonight_moon.toml
-- config.colors = {
--     foreground = "#c0caf5",
--     background = "#1a1b26",
--     cursor_bg = "#c0caf5",
--     cursor_border = "#c0caf5",
--     cursor_fg = "#1a1b26",
--     selection_bg = "#283457",
--     selection_fg = "#c0caf5",
--     split = "#7aa2f7",
--     compose_cursor = "#ff9e64",
--     scrollbar_thumb = "#292e42",
--     ansi = { "#15161e", "#f7768e", "#9ece6a", "#e0af68", "#7aa2f7", "#bb9af7", "#7dcfff", "#a9b1d6" },
--     brights = { "#414868", "#ff899d", "#9fe044", "#faba4a", "#8db0ff", "#c7a9ff", "#a4daff", "#c0caf5" },
-- }

-- Vibr Theme
config.colors = {
    -- The default text color
    foreground = "#abb2bf",

    -- The default background color
    background = "#1e1e1e",

    -- Overrides the cell background color when the current cell is occupied by the
    -- cursor and the cursor style is set to Block
    cursor_bg = "#ffffff",

    -- Overrides the text color when the current cell is occupied by the cursor
    cursor_fg = "#000000",

    -- Specifies the border color of the cursor when the cursor style is set to Block,
    -- or the color of the vertical or horizontal bar when the cursor style is set to
    -- Bar or Underline.
    cursor_border = "#ffffff",

    -- the foreground color of selected text
    selection_fg = "#ffffff",

    -- the background color of selected text
    selection_bg = "#283457",

    -- The color of the scrollbar "thumb"; the portion that represents the current viewport
    scrollbar_thumb = "#282c34",

    -- The color of the split lines between panes
    split = "#282c34",

    ansi = {
        "#283457",
        "#f44747",
        "#afff5e",
        "#fff200",
        "#73b8f1",
        "#ff66cc",
        "#29d4f2",
        "#abb2bf",
    },

    brights = {
        "#abb2bf",
        "#f44747",
        "#ffff5f",
        "#fff200",
        "#73b8f1",
        "#ff66cc",
        "#29d4f2",
        "#ffffff",
    },
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
