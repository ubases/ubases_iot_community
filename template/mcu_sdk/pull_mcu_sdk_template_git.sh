#!/bin/bash

mcu_sdk_dir=$1
mcu_sdk_code_dir=$1/$2
echo "template dir: $mcu_sdk_dir"
echo "code dir: $mcu_sdk_code_dir"

if [ ! -f "$mcu_sdk_dir/conf" ];then
    echo "conf 文件不存在"
    exit 1
fi
# 引入参数
. $mcu_sdk_dir/conf

# 拉取mcu sdk template到指定目录
# 拉取源码 https://e.coding.net/axk/BAT_AIoT_PaaS/iot-mcu-sdk-template.git
mcu_sdk_template_dir=$mcu_sdk_dir/iot-mcu-sdk-template
if [ ! -d $mcu_sdk_template_dir ]; then
        git clone -b $mcu_sdk_template_branch $mcu_sdk_template_git $mcu_sdk_template_dir
        # 判断是否拉取成功
        ret=$?
        if [ $ret != 0 ];then
        echo "git clone mcu sdk模板代码失败"
        exit 2
        fi
else
        cd $mcu_sdk_template_dir
        git pull origin $mcu_sdk_template_branch
        ret=$?
        if [ $ret != 0 ];then
                echo "git pull mcu sdk模板代码失败"
                exit 3
        fi
        cd -
fi

if [ ! -d $mcu_sdk_code_dir ];then
        mkdir -p $mcu_sdk_code_dir
        ret=$?
        if [ $ret != 0 ];then
                echo "mkdir -p $mcu_sdk_code_dir 失败"
                exit 4
        fi
        cp -r $mcu_sdk_template_dir/mcu_sdk $mcu_sdk_code_dir
        ret=$?
        if [ $ret != 0 ];then
                echo "cp -r $mcu_sdk_template_dir/mcu_sdk $mcu_sdk_code_dir 失败"
                exit 5
        fi
else
    cp -r $mcu_sdk_template_dir/mcu_sdk $mcu_sdk_code_dir
    ret=$?
    if [ $ret != 0 ];then
        echo "cp -r $mcu_sdk_template_dir/mcu_sdk $mcu_sdk_code_dir 失败"
        exit 5
    fi
fi

