#!/bin/sh

mcu_sdk_dir=$1
mcu_sdk_code_dir=$1/$2
echo "template dir: $mcu_sdk_dir"
echo "code dir: $mcu_sdk_code_dir"

mcu_sdk_template_dir=$mcu_sdk_dir/iot-mcu-sdk-template
cp $mcu_sdk_template_dir $mcu_sdk_code_dir/iot-mcu-sdk-template -r

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