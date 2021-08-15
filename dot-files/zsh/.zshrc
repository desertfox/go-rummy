if [[ -r "${XDG_CACHE_HOME:-$HOME/.cache}/p10k-instant-prompt-${(%):-%n}.zsh" ]]; then
  source "${XDG_CACHE_HOME:-$HOME/.cache}/p10k-instant-prompt-${(%):-%n}.zsh"
fi
export ZSH="/home/pi/.oh-my-zsh"
export PATH=$PATH:/usr/local/go/bin/
export PATH=$PATH:/home/pi/diff-so-fancy
export GOPATH=$HOME/code/go
export PATH=$PATH:$GOPATH/bin

ZSH_THEME="darkblood"
ZSH_THEME="powerlevel10k/powerlevel10k"

plugins=(git golang alias-finder z wd)

source $ZSH/oh-my-zsh.sh

[[ ! -f ~/.p10k.zsh ]] || source ~/.p10k.zsh
