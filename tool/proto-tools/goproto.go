package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	args := os.Args
	//参数示例
	//args := strings.Split(`xxxx\tools\goproto\goproto.exe xxxxx\bat\ubases_iot_community\iot_proto\protos\gen\iot_app_oem\oem_app_assist_release_market_service.gen.proto`, " ")

	fmt.Println("----> args: ", args)

	fileDirs := strings.Split(args[1], `\`)
	nameIndex := len(fileDirs)
	fileName := fileDirs[nameIndex-1:][0]
	fileDir := strings.Join(fileDirs[0:nameIndex-1], `\`) // ./
	fileDir2 := ""
	if strings.Index(args[1], `\ext\`) != -1 {
		fileDir2 = strings.Join(fileDirs[0:nameIndex-2], `\`) // ./
	} else {
		fileDir2 = strings.Join(fileDirs[0:nameIndex-3], `\`) // ./
	}
	fmt.Println(fileName, fileDir)

	fileNames := strings.Split(fileName, ".")
	lastIndex := len(fileNames) - 1
	prefix := fileNames[lastIndex:][0]
	prefixName := strings.Join(fileNames[0:lastIndex], ".")

	if prefix != "proto" {
		fmt.Println("---->[error] proto file: ", "params error: not proto file")
		return
	}

	cmd := exec.Command("protoc",
		fmt.Sprintf(`--proto_path=%s`, fileDir),
		fmt.Sprintf(`--micro_out=%s\protosService`, fileDir2),
		fmt.Sprintf(`--go_out=%s\protosService`, fileDir2), fileName)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("---->[error]", "protoc",
			fmt.Sprintf(`--proto_path=%s`, fileDir),
			fmt.Sprintf(`--micro_out=%s\protosService`, fileDir2),
			fmt.Sprintf(`--go_out=%s\protosService`, fileDir2), fileName, "error:", err.Error()+": "+stderr.String())
		return
	}
	fmt.Println("---->", "protoc micro result: ok", out.String())

	cmd = exec.Command("protoc-go-inject-tag", `-input`,
		fmt.Sprintf(`%s\protosService\%s.pb.go`, fileDir2, prefixName))
	var outTag bytes.Buffer
	var stderrTag bytes.Buffer
	cmd.Stdout = &outTag
	cmd.Stderr = &stderrTag
	err = cmd.Run()
	if err != nil {
		fmt.Println("---->[error]", "protoc-go-inject-tag", `-input`,
			fmt.Sprintf(`%s\protosService\%s.pb.go`, fileDir2, prefixName), "error:", err.Error()+": "+stderr.String())
		return
	}
	fmt.Println("---->", "protoc-go-inject-tag result: ok", outTag.String())

	cmd = exec.Command("protoc",
		fmt.Sprintf(`--proto_path=%s`, fileDir),
		fmt.Sprintf(`--go_out=%s\protosService`, fileDir2),
		fmt.Sprintf(`--micro_out=%s\protosService`, fileDir2), fileName)

	var outProtoc bytes.Buffer
	var stderrProtoc bytes.Buffer
	cmd.Stdout = &outProtoc
	cmd.Stderr = &stderrProtoc
	err = cmd.Run()
	if err != nil {
		fmt.Println("---->[error]", "protoc",
			fmt.Sprintf(`--proto_path=%s`, fileDir),
			fmt.Sprintf(`--go_out=%s\protosService`, fileDir2),
			fmt.Sprintf(`--micro_out=%s\protosService`, fileDir2), fileName, "error:", err.Error()+": "+stderr.String())
		return
	}
	fmt.Println("---->", "protoc go result: ok", outProtoc.String())

	//TODO 替换文件内容，将文件中的 替换为空
	replaceString(fmt.Sprintf(`%s\protosService\%s.pb.go`, fileDir2, prefixName), `_ "/api"`, "")
	replaceString(fmt.Sprintf(`%s\protosService\%s.pb.micro.go`, fileDir2, prefixName), `_ "/api"`, "")

	fmt.Println("ok")
}

func replaceString(fileName string, oldStr, newStr string) {
	in, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file fail:", err)
		os.Exit(-1)
	}
	defer in.Close()

	out, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println("Open write file fail:", err)
		os.Exit(-1)
	}
	defer out.Close()

	br := bufio.NewReader(in)
	index := 1
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read err:", err)
			os.Exit(-1)
		}
		newLine := strings.Replace(string(line), oldStr, newStr, -1)
		_, err = out.WriteString(newLine + "\n")
		if err != nil {
			fmt.Println("write to file fail:", err)
			os.Exit(-1)
		}
		index++
	}
}
