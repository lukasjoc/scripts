# Basic functions

#Check for system (mac specific stuff)
switch (uname -s)
  case "Darwin"
    function locate_update -d "Update locate db"
      sudo "/usr/libexec/locate.updatedb"
    end
    function runbrewstuff -d "update and upgrade brew plugins"
      echo "update/upgrade brew formulars..."
      ;and brew update
      ;and brew upgrade --greedy
      ;and rm -rf (brew --cache)
      ;and brew cleanup
      ;and brew doctor
      ;and echo "done ;)"
    end
    function bc -d "install brew caskroom"
      brew "$argv[1]" "$argv[2]" --cask
    end
end

function ns -d "Start new default shell"
  exec "$SHELL" -l
end

function npmls -d "List all globally installed npm packages"
  npm list -g --depth=0;
end

function sc -d "List all available scripts for package.json and composer.json"
  if test -e "$PWD/package.json"
    jq .scripts "package.json"
  end
  if test -e "$PWD/composer.json"
    jq .scripts "composer.json"
  end
end

function fun -d "Go to directory $HOME/fun"
  cd "$HOME/fun"
end

function builds -d "Go to directory $HOME/builds"
  cd "$HOME/builds"
end

function gocmd -d "Go Language Helper"
  switch "$argv[1]"
    case -b -bench
      go test -v --bench . --benchmem -race
    case -v -clean
      go clean -x -r -cache -modcache
    case -mod
      go mod init "github.com/lukasjoc/$argv[2]"
    case "" -h
      echo "Usage: gocmd <option>"
      echo "Options:"
      echo "  -c or -clean    Clean cache and modcache"
      echo "  -b or -bench    Test and benchmark go code"
      echo "  -mod            Create mod init files"
      echo "  -h              Print this help message"
  end
end

# Taken from -> https://github.com/decors/fish-colored-man/blob/master/functions/man.fish
function man --wraps man --description 'Format and display manual pages'
  set -q man_blink; and set -l blink (set_color $man_blink); or set -l blink (set_color -o red)
  set -q man_bold; and set -l bold (set_color $man_bold); or set -l bold (set_color -o 5fafd7)
  set -q man_standout; and set -l standout (set_color $man_standout); or set -l standout (set_color 949494)
  set -q man_underline; and set -l underline (set_color $man_underline); or set -l underline (set_color -u afafd7)

  set -l end (printf "\e[0m")
  set -lx LESS_TERMCAP_mb "$blink"
  set -lx LESS_TERMCAP_md "$bold"
  set -lx LESS_TERMCAP_me "$end"
  set -lx LESS_TERMCAP_so "$standout"
  set -lx LESS_TERMCAP_se "$end"
  set -lx LESS_TERMCAP_us "$underline"
  set -lx LESS_TERMCAP_ue "$end"
  set -lx LESS '-R -s'

  set -lx GROFF_NO_SGR yes # fedora
  set -lx MANPATH (string join : $MANPATH)
  if test -z "$MANPATH"
      type -q manpath
      and set MANPATH (command manpath)
  end

  # Check data dir for Fish 2.x compatibility
  set -l fish_data_dir
  if set -q __fish_data_dir
    set fish_data_dir $__fish_data_dir
  else
    set fish_data_dir $__fish_datadir
  end

  set -l fish_manpath (dirname $fish_data_dir)/fish/man
  if test -d "$fish_manpath" -a -n "$MANPATH"
    set MANPATH "$fish_manpath":$MANPATH
  end
  command man "$argv"
end

