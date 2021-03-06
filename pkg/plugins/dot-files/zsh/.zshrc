if [[ -r "${XDG_CACHE_HOME:-$HOME/.cache}/p10k-instant-prompt-${(%):-%n}.zsh" ]]; then
  source "${XDG_CACHE_HOME:-$HOME/.cache}/p10k-instant-prompt-${(%):-%n}.zsh"
fi
export ZSH="/home/$USER/.oh-my-zsh"
export PATH=$PATH:/usr/local/go/bin/
export PATH=$PATH:/home/$USER/diff-so-fancy
export GOPATH=$HOME/code/go
export PATH=$PATH:$GOPATH/bin

ZSH_THEME="darkblood"
ZSH_THEME="powerlevel10k/powerlevel10k"

plugins=(git golang alias-finder z wd)

source $ZSH/oh-my-zsh.sh

if [ -f $HOME/.bash_aliases ]; then
  source $HOME/.bash_aliases
fi

[[ ! -f ~/.p10k.zsh ]] || source ~/.p10k.zsh
