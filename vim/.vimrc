syntax on
set number
set nobackup
set noswapfile
set relativenumber
set showmode showcmd
set tabstop=4
set modelines=0
set path+=**
set encoding=utf-8
set scrolloff=2
set shiftwidth=4
set softtabstop=4
set autoindent
set showmode showcmd
set scrolloff=10
set nowrap
set incsearch
set ignorecase
set smartcase
set hlsearch
set history=1000
set wildmenu
set wildmode=list:longest
set wildignore=*.jpg,*.gif,*.png,*.img,*.pdf,*.exe,*.xlsx,*.docx

let mapleader=" "

inoremap jj <esc>
nnoremap <leader>\ :nohlsearch<CR>
nnoremap o o<esc>
nnoremap O O<esc>
nnoremap n nzz
nnoremap N Nzz
nnoremap Y y$
nnoremap <leader>l dd

set statusline=
set statusline+=\ %F\ %M\ %Y\ %R
set statusline+=%=
set statusline+=\ ascii:\ %\b\ hex:\ 0x%B\ row:\ %l\ col:\ %c\ percent:\ %p%%
set statusline=2
