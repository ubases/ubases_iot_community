package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var (
	version string = "1.0.1"
)

func main() {
	args := os.Args
	//args := strings.Split(`xxxxxxx\tools\goproto\goproto.exe xxxxxxx\cloud-platform_v2\iot_proto\protos\gen\iot_config\config_oss_service.gen.proto`, " ")
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
	replaceString(fmt.Sprintf(`%s\protosService\%s.pb.go`, fileDir2, prefixName), `_ "/api"`, `// _ "/api"`)
	replaceString(fmt.Sprintf(`%s\protosService\%s.pb.micro.go`, fileDir2, prefixName), `_ "/api"`, `// _ "/api"`)

	fmt.Println("ok")
}

func replaceString(fileName string, oldStr, newStr string) {
	// 读取文件内容
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("读取文件失败:", err)
		return
	}

	// 将文件内容转换为字符串
	text := string(content)

	// 替换字符串
	newText := strings.ReplaceAll(text, oldStr, newStr)

	// 写入替换后的内容到文件
	err = ioutil.WriteFile(fileName, []byte(newText), os.ModePerm)
	if err != nil {
		fmt.Println("写入文件失败:", err)
		return
	}
	fmt.Println("FINISH!")
}
