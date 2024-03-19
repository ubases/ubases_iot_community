@echo off
rem 执行本脚本之前，请将本目录下的google文件夹拷贝到GOPATH下的include目录下，若原来有google文件夹，请直接覆盖

call ./ext/aconfig.bat
call ./gen/iot-app/aconfig.bat

pause
