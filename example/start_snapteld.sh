#! /bin/bash

snapteld --config ./snap-config.json &

sleep 5

snaptel plugin load ../build/linux/x86_64/snap-plugin-collector-pciessd
snaptel plugin load ../../snap-plugin-publisher-file/build/linux/x86_64/snap-plugin-publisher-file

snaptel task create -t ./pciesssd-file.yaml
