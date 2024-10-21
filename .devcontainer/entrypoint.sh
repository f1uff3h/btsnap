#!/bin/bash

mount /btrfs.img /mnt

btrfs su cr /mnt/su1
btrfs su cr /mnt/su2
btrfs su cr /mnt/su3

nu
