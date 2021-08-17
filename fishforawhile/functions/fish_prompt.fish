set -g __fish_git_prompt_showdirtystate 'no'
set -g __fish_git_prompt_showstashstate 'no'
set -g __fish_git_prompt_showuntrackedfiles 'no'
set -g __fish_git_prompt_showupstream 'no'

set -g __fish_git_prompt_color_branch cyan
set -g __fish_git_prompt_color cyan
set -g __fish_git_prompt_color_branch cyan
set -g __fish_git_prompt_color_upstream_ahead cyan
set -g __fish_git_prompt_color_upstream_behind cyan

set -g __fish_git_prompt_char_dirtystate '*-'
set -g __fish_git_prompt_char_stagedstate '*+'
set -g __fish_git_prompt_char_untrackedfiles ''
set -g __fish_git_prompt_char_stashstate ''
set -g __fish_git_prompt_char_upstream_equal ''
set -g __fish_git_prompt_char_upstream_ahead '^'
set -g __fish_git_prompt_char_upstream_behind 'd'

function fish_user_keybidnings; fish_vi_keybindings; end
function fish_mode_prompt; end
function fish_prompt -d "Write out the prompt"
  printf '%s%s%s%s%s'\
    (set_color green) (whoami) \
    (set_color normal) ' :: ' \
    (set_color red) (printf '%s' $PWD | sed -e "s|^$HOME|~|" -e 's|^/private||') \
    (__fish_vcs_prompt) \
    (set_color normal) ' $ '
end
