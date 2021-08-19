call plug#begin()

Plug 'scrooloose/nerdtree', { 'on':  'NERDTreeToggle' }
Plug 'tpope/vim-fugitive'
Plug 'scrooloose/syntastic'
Plug 'airblade/vim-gitgutter'
Plug 'vim-airline/vim-airline'
Plug 'vim-airline/vim-airline-themes'
Plug 'kien/rainbow_parentheses.vim'
Plug 'ervandew/supertab'
Plug 'sheerun/vim-polyglot'
Plug 'ayu-theme/ayu-vim'

call plug#end()

set termguicolors     " enable true colors support
set background=dark
colorscheme ayu

"let ayucolor="light"  " for light version of theme
let ayucolor="mirage" " for mirage version of theme
"let ayucolor="dark"   " for dark version of theme

set shiftwidth=4
set tabstop=4
set number

set showmatch
set noswapfile
set noerrorbells
