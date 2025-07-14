#!/bin/bash

# Source Prelude if possible ( This includes scripts etc)
[[ -d "$HOME/.local/scripts" ]] && source "$HOME/.local/scripts/prelude";

export HISTCONTROL=ignoreboth
export MANWIDTH="92"
export PAGER="less"
export VISUAL="vim"
export PS1="$ "
