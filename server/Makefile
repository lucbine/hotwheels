#!/bin/bash

PROJECT=hotwheels-server
CURRENT_DIR=$(shell pwd)
CONFIGPATH=$(shell echo ${CURRENT_DIR}/config)
OUTPUT_DIR=_publish_dir
UNAME=$(shell uname)
LG_OUT_PATH="/data/logs/app/${PROJECT}"

.PHONY:common

common:
	rm -rf $(OUTPUT_DIR)
	mkdir $(OUTPUT_DIR)
	cd _publish_dir
	mkdir -p _publish_dir/$(PROJECT)/bin
	mkdir -p _publish_dir/$(PROJECT)/config
	cp -f run.sh _publish_dir/$(PROJECT)/run.sh
	export GO111MODULE=on && go build -o _publish_dir/$(PROJECT)/bin/$(PROJECT) -mod=vendor main.go

	cp -af $(CONFIGPATH)/* _publish_dir/$(PROJECT)/config/

dev: common
	$(call func_env_load,$@)

qa: common
	$(call func_env_load,$@)

prd: common
	$(call func_env_load,$@)

pre: common
	$(call func_env_load,$@)

clean:
	rm -rf $(OUTPUT_DIR)

run: common
	_publish_dir/$(PROJECT)/bin/$(PROJECT) --env=dev

install: dev
	bash $(OUTPUT_DIR)/$(PROJECT)/run.sh startOrReload


define func_env_load
	#替换run.sh 的环境变量
    if [ $(UNAME) == "Darwin" ];then  \
    		sed -i "" "s/BIN_FILE_CONST/$(PROJECT)/g" ${OUTPUT_DIR}/$(PROJECT)/run.sh ;\
    		sed -i "" "s/ENV_CONST/$(1)/g" ${OUTPUT_DIR}/$(PROJECT)/run.sh ;\
    		sed -i "" "s?LG_OUT_CONST?$(LG_OUT_PATH)?" ${OUTPUT_DIR}/$(PROJECT)/run.sh ;\
    else \
    		sed -i "s/BIN_FILE_CONST/$(PROJECT)/g" ${OUTPUT_DIR}/$(PROJECT)/run.sh ;\
    		sed -i "s/ENV_CONST/$(1)/g" ${OUTPUT_DIR}/$(PROJECT)/run.sh ;\
    		sed -i "" "s?LG_OUT_CONST?$(LG_OUT_PATH)?" ${OUTPUT_DIR}/$(PROJECT)/run.sh ;\
    fi
endef

