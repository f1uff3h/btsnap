#!/bin/bash

source ./.init

# get home folder of user with UID 1000
user_home=$(grep 1000 /etc/passwd | cut -d: -f6)

sudo podman container run --rm --privileged -it \
  -v go-nu:/root/.config/nushell \
  -v go-nv:/root/.local/share/nvim \
  -v persist-z:/root/.local/share/zoxide \
  -v "${user_home}"/.cache/nvim/codeium:/root/.cache/nvim/codeium \
  -v "${user_home}"/repos/.init/:/init \
  -v "${user_home}"/.ssh/:/root/.ssh:ro \
  -v "${user_home}"/.gitconfig:/root/.gitconfig:ro \
  -v ./:/workspace \
  "${containerRegistry}/${containerImage}"
