#!/usr/bin/env bash
#! Goland 命令行可用
pushd `dirname $0` > /dev/null
base_dir=$(pwd -P);

sequence="sequence"
protoc ./$sequence/$sequence.proto --micro_out=. --go_out=.
