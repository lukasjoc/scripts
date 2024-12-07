#!/bin/bash
set -o vi

# https://www.gnu.org/software/bash/manual/html_node/The-Shopt-Builtin.html
shopt -s cdspell dirspell
shopt -s checkhash
shopt -s failglob
shopt -u hostcomplete
shopt -s cmdhist
shopt -s histappend

export HISTCONTROL=ignoreboth
export MANWIDTH="92"
export PAGER="less"
export VISUAL="vim"

export PS1="$ "
# export PROMPT_COMMAND='PS1="\w $($HOME/.local/scripts/gitstatus-bin)# "'

alias ..="cd ../"
alias ls="ls --author -bcFA --human-readable -N -S --color=auto"
alias ll="ls -ls"
alias grep="grep --color=auto"
alias watch="watch --color"
alias clear="printf '\e[1;1H\e[2J'"
alias npm="pnpm"
alias bye="sudo shutdown -h now"
bin2dec() { echo "obase=10;ibase=2;$1" | bc; }
bin2hex() { echo "obase=10000;ibase=2;$1" | bc; }
dec2bin() { echo "obase=2;$1" | bc; }

# overwrites path to $PATH checking if path is in $PATH before doing so.
pathadd() {
    [[ ! "${PATH//:/ }" =~ "$1" && -d  $1 ]] \
    && export PATH="$1:$PATH"
}

export XDG_CONFIG_HOME="$HOME/.config"
export XDG_CACHE_HOME="$HOME/.cache"
export XDG_DATA_HOME="$HOME/.local/share"
export XDG_DATA_BIN="$HOME/.local/bin"
pathadd $XDG_DATA_BIN

export SCRIPTSPATH="$HOME/.local/scripts"
export SCRIPTSINST="$HOME/.local/scripts/debian"
. "$HOME/.local/scripts/env-completion"
pathadd $SCRIPTSPATH
pathadd $SCRIPTSINST

export GOPATH="$HOME/.local/go/path"
export GOBIN="$HOME/.local/go/bin"
export CGO_ENABLED=0
pathadd $GOBIN

source ~/perl5/perlbrew/etc/bashrc
. "$HOME/.cargo/env"


export NVM_DIR="$HOME/.config/nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion

# pnpm
export PNPM_HOME="/home/lukas/.local/share/pnpm"
case ":$PATH:" in
  *":$PNPM_HOME:"*) ;;
  *) export PATH="$PNPM_HOME:$PATH" ;;
esac
# pnpm end

# Add RVM to PATH for scripting. Make sure this is the last PATH variable change.
export PATH="$PATH:$HOME/.rvm/bin"
