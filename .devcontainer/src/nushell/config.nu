$env.config = {
    show_banner: false
}

def rmswap [] { fd -H .*.sw[p|o]$ $env.HOME | xargs -I{} rm -rf {}}

alias nv = nvim
alias ll = ls -ls
alias la = ls -la
alias glog = git log --all --oneline --decorate --graph
alias gl = git log --all --oneline -n 10

source ~/.cache/zoxide/init.nu
use ~/.cache/starship/init.nu
