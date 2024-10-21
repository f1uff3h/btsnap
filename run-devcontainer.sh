#!/bin/bash

source ./.init

sudo podman login --tls-verify=false $containerRegistry
sudo podman container run --privileged --tls-verify=false --rm -it -v go-nu:/root/.config/nushell -v go-nv:/root/.local/share/nvim -v persist-z:/root/.local/share/zoxide  -v "${HOME}"/repos/.init/:/init -v "${HOME}"/.ssh/:/root/.ssh:ro -v "${HOME}"/.gitconfig:/root/.gitconfig:ro -v ./:/workspace "${containerRegistry}/${containerImage}"
