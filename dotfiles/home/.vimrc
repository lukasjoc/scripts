set nocompatible modifiable
filetype plugin indent on
set encoding=utf-8
set ttyfast
set title
set nopaste
set noswapfile
set noundofile
set ruler
set list
set backspace=indent,eol,start
set expandtab
set tabstop=4
set shiftwidth=4
set softtabstop=4
set smartindent
set nowrap
set linebreak
set hlsearch
set nonumber
set termguicolors
set wildmenu
set omnifunc=syntaxcomplete#Complete
set completeopt=menu,menuone,noselect
set wildignore+=*.so,*.swp,*.zip
set updatetime=200
set timeout
set timeoutlen=300
set textwidth=92
set colorcolumn=92
set listchars=tab:\ \ ,trail:.
set scrolloff=21
set showmode
set signcolumn=yes
let g:netrw_banner=0
let g:netrw_fastbrowse=1
let g:netrw_list_hide=0
let g:netrw_liststyle=1
syntax on
set background=dark
colorscheme synthwave

augroup vimrc_help
    autocmd!
    autocmd BufEnter *.txt if &buftype == 'help' | wincmd L | endif
augroup END

"" Plugins
" https://github.com/tpope/vim-commentary
" https://github.com/junegunn/fzf
" https://github.com/junegunn/fzf.vim
" https://github.com/dense-analysis/ale
" https://github.com/ashervb/synthwave.vim
" https://github.com/fatih/vim-go

let g:mapleader=","
let g:maplocalleader=","
nnoremap <Leader><Leader>z :nohl<CR><C-L>
nnoremap <Leader><Leader>e :Explore<CR>
nnoremap <Leader><Leader>f :Files!<CR>
nnoremap <Leader><Leader>o :History!<CR>
nnoremap <Leader><Leader>g :Rg!<CR>
nnoremap  <S-J> :ALEGoToDefinition<CR>
nnoremap  <Leader><Leader>re :ALERename<CR>
nnoremap  <Leader><Leader>ca :ALECodeAction<CR>

let g:ale_fixers = {
\   '*': ['remove_trailing_lines', 'trim_whitespace'],
\   'c': ['clangd'],
\   'sh': ['shellcheck'],
\   'go': ['gofmt', 'gobuild'],
\   'javascript': ['eslint'],
\   'typescript': ['eslint', 'tsc'],
\}

let g:ale_sign_column_always = 1
let g:ale_sign_error = 'E'
let g:ale_sign_warning = 'W'
let g:ale_lint_on_text_changed = 'never'
let g:ale_lint_on_insert_leave = 0
let g:ale_lint_on_enter = 0
let g:ale_set_loclist = 0
let g:ale_set_quickfix = 1

let g:go_fmt_fail_silently = 0
let g:go_fmt_command = 'goimports'
let g:go_fmt_autosave = 1
let g:go_gopls_enabled = 1
let g:go_highlight_types = 1
let g:go_highlight_fields = 1
let g:go_highlight_functions = 1
let g:go_highlight_function_calls = 1
let g:go_highlight_operators = 1
let g:go_highlight_extra_types = 1
let g:go_highlight_variable_declarations = 1
let g:go_highlight_variable_assignments = 1
let g:go_highlight_build_constraints = 1
let g:go_highlight_diagnostic_errors = 1
let g:go_highlight_diagnostic_warnings = 1

" -- Markdown preview of current file using glow
command! GlowPreview !glow %:S

