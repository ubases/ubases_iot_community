package iotutil

import (
	"archive/zip"
	"bufio"
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	imgext "github.com/shamsher31/goimgext"
)

// EncodeMD5 md5 encryption
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

func FileMD5(filename string) (string, error) {
	f, err := os.Open(filename) //打开文件
	if nil != err {
		return "", err
	}
	defer f.Close()

	md5Handle := md5.New()         //创建 md5 句柄
	_, err = io.Copy(md5Handle, f) //将文件内容拷贝到 md5 句柄中
	if nil != err {
		return "", err
	}
	md := md5Handle.Sum(nil)
	md5str := fmt.Sprintf("%x", md) //将 []byte 转为 string
	return md5str, nil
}

func FMd5(f *os.File) (string, error) {
	md5Handle := md5.New()          //创建 md5 句柄
	_, err := io.Copy(md5Handle, f) //将文件内容拷贝到 md5 句柄中
	if nil != err {
		return "", err
	}
	md := md5Handle.Sum(nil)
	md5str := fmt.Sprintf("%x", md) //将 []byte 转为 string
	return md5str, nil
}

// GetSize 获取文件大小
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

// GetExt 获取文件后缀
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// CheckExist 检查文件是否存在
func CheckExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

// CheckPermission 检查文件权限
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

// IsNotExistMkDir 检查文件夹是否存在
// 如果不存在则新建文件夹
func IsNotExistMkDir(src string) error {
	if exist := !CheckExist(src); exist == false {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

func CreateFile(src string) error {
	// 创建日志文件
	file, err := os.OpenFile(src, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

// MkDir 新建文件夹
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// Open 打开文件
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// GetImgType 获取Img文件类型
func GetImgType(p string) (string, error) {
	file, err := os.Open(p)

	if err != nil {
		return "", err
	}

	buff := make([]byte, 512)

	_, err = file.Read(buff)

	if err != nil {
		return "", err
	}

	filetype := http.DetectContentType(buff)

	ext := imgext.Get()

	for i := 0; i < len(ext); i++ {
		if strings.Contains(ext[i], filetype[6:len(filetype)]) {
			return filetype, nil
		}
	}

	return "", errors.New("Invalid image type")
}

// GetType 获取文件类型
func GetType(p string) (string, error) {
	file, err := os.Open(p)

	if err != nil {
		return "", err
	}

	buff := make([]byte, 512)

	_, err = file.Read(buff)

	if err != nil {
		return "", err
	}

	filetype := http.DetectContentType(buff)

	//ext := GetExt(p)
	//var list = strings.Split(filetype, "/")
	//filetype = list[0] + "/" + ext
	return filetype, nil
}

func PathCreate(dir string) error {
	return os.MkdirAll(dir, os.ModePerm)
}

// PathExist 判断目录是否存在
func PathExist(addr string) bool {
	s, err := os.Stat(addr)
	if err != nil {
		log.Println(err)
		return false
	}
	return s.IsDir()
}

func FileCreate(content bytes.Buffer, name string) {
	file, err := os.Create(name)
	if err != nil {
		log.Println(err)
	}
	_, err = file.WriteString(content.String())
	if err != nil {
		log.Println(err)

	}
	file.Close()
}

type ReplaceHelper struct {
	Root    string //路径
	OldText string //需要替换的文本
	NewText string //新的文本
}

func (h *ReplaceHelper) DoWrok() error {

	return filepath.Walk(h.Root, h.walkCallback)

}

func (h ReplaceHelper) walkCallback(path string, f os.FileInfo, err error) error {

	if err != nil {
		return err
	}
	if f == nil {
		return nil
	}
	if f.IsDir() {
		log.Println("DIR:", path)
		return nil
	}

	//文件类型需要进行过滤

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		//err
		return err
	}
	content := string(buf)
	log.Printf("h.OldText: %s \n", h.OldText)
	log.Printf("h.NewText: %s \n", h.NewText)

	//替换
	newContent := strings.Replace(content, h.OldText, h.NewText, -1)

	//重新写入
	ioutil.WriteFile(path, []byte(newContent), 0)

	return err
}

func FileMonitoringById(ctx context.Context, filePth string, id string, group string, hookfn func(context.Context, string, string, []byte)) {
	f, err := os.Open(filePth)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	f.Seek(0, 2)
	for {
		if ctx.Err() != nil {
			break
		}
		line, err := rd.ReadBytes('\n')
		// 如果是文件末尾不返回
		if err == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			log.Fatalln(err)
		}
		go hookfn(ctx, id, group, line)
	}
}

// 获取文件大小
func GetFileSize(filename string) int64 {
	var result int64
	filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})
	return result
}

// 获取当前路径，比如：E:/abc/data/test
func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func CopyFile(src, des string) (written int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	//获取源文件的权限
	fi, _ := srcFile.Stat()
	perm := fi.Mode()

	//desFile, err := os.Create(des)  //无法复制源文件的所有权限
	desFile, err := os.OpenFile(des, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm) //复制源文件的所有权限
	if err != nil {
		return 0, err
	}
	defer desFile.Close()

	return io.Copy(desFile, srcFile)
}

// 检查并且创建文件夹
func CheckAndCreateFolder(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
}
func IsFileExist(filename string, filesize int64) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Println(info)
		return false
	}
	if filesize == info.Size() {
		fmt.Println("安装包已存在！", info.Name(), info.Size(), info.ModTime())
		return true
	}
	del := os.Remove(filename)
	if del != nil {
		fmt.Println(del)
	}
	return false
}

func DownloadFile(url string, localFileName string) error {
	var (
		fsize   int64
		buf     = make([]byte, 32*1024)
		written int64
	)
	tmpFilePath := localFileName //+ ".download"
	fmt.Println(tmpFilePath)
	//创建一个http client
	client := new(http.Client)
	//client.Timeout = time.Second * 60 //设置超时时间
	//get方法获取资源
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	//读取服务器返回的文件大小
	fsize, err = strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	if IsFileExist(localFileName, fsize) {
		return err
	}
	fmt.Println("fsize", fsize)
	//创建文件
	file, err := os.Create(tmpFilePath)
	if err != nil {
		return err
	}
	defer file.Close()
	if resp.Body == nil {
		return errors.New("body is null")
	}
	defer resp.Body.Close()
	//下面是 io.copyBuffer() 的简化版本
	for {
		//读取bytes
		nr, er := resp.Body.Read(buf)
		if nr > 0 {
			//写入bytes
			nw, ew := file.Write(buf[0:nr])
			//数据长度大于0
			if nw > 0 {
				written += int64(nw)
			}
			//写入出错
			if ew != nil {
				err = ew
				break
			}
			//读取是数据长度不等于写入的数据长度
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
		//没有错误了快使用 callback
		//fb(fsize, written)
	}
	fmt.Println(err)
	if err == nil {
		file.Close()
		err = os.Rename(tmpFilePath, localFileName)
		fmt.Println(err)
	}
	return err
}

// 把content写入fileName
func WriteFile(fileName string, content string) error {
	path := fileName
	//os.O_WRONLY | os.O_CREATE
	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	//使用缓存方式写入
	writer := bufio.NewWriter(file)
	_, w_err := writer.WriteString(content)
	//将缓存中数据刷新到文本中
	writer.Flush()
	return w_err
}

// 打包成zip文件
func Zip(src_dir string, zip_file_name string) {

	// 预防：旧文件无法覆盖
	os.RemoveAll(zip_file_name)

	// 创建：zip文件
	zipfile, _ := os.Create(zip_file_name)
	defer zipfile.Close()

	// 打开：zip文件
	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	// 遍历路径信息
	filepath.Walk(src_dir, func(path string, info os.FileInfo, _ error) error {

		// 如果是源路径，提前进行下一个遍历
		if path == src_dir {
			return nil
		}

		// 获取：文件头信息
		header, _ := zip.FileInfoHeader(info)
		//header.Name = strings.TrimPrefix(strings.ReplaceAll(path,"\\","/"), strings.ReplaceAll(src_dir,"\\","/")+"/")
		header.Name = strings.TrimPrefix(strings.ReplaceAll(path, string(filepath.Separator), "/"), strings.ReplaceAll(src_dir, string(filepath.Separator), "/")+"/")

		// 判断：文件是不是文件夹
		if info.IsDir() {
			header.Name += "/" //string(filepath.Separator)
		} else {
			// 设置：zip的文件压缩算法
			header.Method = zip.Deflate
		}

		// 创建：压缩包头部信息
		writer, _ := archive.CreateHeader(header)
		if !info.IsDir() {
			file, _ := os.Open(path)
			defer file.Close()
			io.Copy(writer, file)
		}
		return nil
	})
}

// 参数1：压缩文件的名称, 参数2：文件名称的集合
func ZipFiles(filename string, files []string, oldform, newform string) error {

	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer newZipFile.Close()
	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	// 把files添加到zip中
	for _, file := range files {
		zipfile, err := os.Open(file)
		if err != nil {
			return err
		}
		defer zipfile.Close()
		info, err := zipfile.Stat()
		if err != nil {
			return err
		}
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		//header.Name = strings.Replace(file, oldform, newform, -1)
		header.Method = zip.Deflate
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		if _, err = io.Copy(writer, zipfile); err != nil {
			return err
		}
	}
	return nil
}

// 遍历根目录文件的方法
func GetAllFile(pathname string, s []string) ([]string, error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {

		return s, err
	}

	for _, fi := range rd {
		if !fi.IsDir() {
			fullName := pathname + "/" + fi.Name()
			s = append(s, fullName)
		}
	}
	return s, nil
}
