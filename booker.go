package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

const (
	PS = string(os.PathSeparator)
)

type Config struct {
	Repository string `yaml:"repository"` // git仓库地址
	Branch     string `yaml:"branch"`     // 分支
	Secret     string `yaml:"secret"`     // 密钥
}

var (
	repositoryDir = flag.String("r", "", "repository path")
	bookDir       = flag.String("b", "", "book path")
	configDir     = flag.String("c", "", "config path")
	config        Config
)

func init() {
	// 初始化目录，不存在则创建
	flag.Parse()
	fmt.Println(*repositoryDir, *bookDir, *configDir)
	if !dirExist(*repositoryDir) {
		err := os.MkdirAll(*repositoryDir, 0755)
		printError(err)
	}
	if !dirExist(*bookDir) {
		err := os.MkdirAll(*bookDir, 0755)
		printError(err)
	}
	if !dirExist(*configDir) {
		err := os.MkdirAll(*configDir, 0755)
		printError(err)
	}

	// 读取配置文件
	configPath := *configDir + PS + "config.yml"
	yamlFile, err := ioutil.ReadFile(configPath)
	printError(err)
	err = yaml.Unmarshal(yamlFile, &config)
	printError(err)
}

func main() {

	// 拉取代码
	initBookRepository()

	// github webhook监听
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is booker server"))
	})
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		log.Println("build books")
		signatures, ok := r.Header["X-Hub-Signature"]
		if !ok {
			w.Write([]byte("no header X-Hub-Signature"))
			return
		}
		signature := signatures[0]
		body, _ := ioutil.ReadAll(r.Body)
		secret := "sha1=" + HMACSHA1(string(body))
		if signature != secret {
			log.Println("sign verify fail")
			return
		}

		go buildBooks()
		w.Write([]byte("build books"))
	})
	http.ListenAndServe("0.0.0.0:5454", nil)
}

// 第一次pull book仓库
func initBookRepository() {
	lockFile := *repositoryDir + PS + "booker.lock"
	if fileExist(lockFile) {
		return
	}

	cmd := exec.Command("sh", "-c", "git clone "+config.Repository+" "+*repositoryDir)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err == nil {
		log.Println(out.String())
		f, _ := os.Create(lockFile)
		defer f.Close()
		f.Write([]byte("lock"))
		cmd = exec.Command("sh", "-c", "cd "+*repositoryDir+" && git checkout -b "+config.Branch+" origin/"+config.Branch)
		cmd.Run()
		return
	}
	log.Println(stderr.String())
	printError(err)
	os.Exit(1)
}

// 增量构建book
func buildBooks() {
	repositoryDirCmd := "cd " + *repositoryDir + " && "
	cmd := exec.Command("sh", "-c", repositoryDirCmd+"git fetch")
	cmd.Run()
	cmd = exec.Command("sh", "-c", repositoryDirCmd+"git diff "+config.Branch+" origin/"+config.Branch+" --name-only")
	output, err := cmd.CombinedOutput()
	if err != nil {
		printError(err)
		return
	}
	change := string(output)
	re, err := regexp.Compile(`(\w+)/.*?`)
	if err != nil {
		printError(err)
		return
	}
	books := re.FindAllString(change, -1)
	if len(books) == 0 {
		return
	}
	books = booksUnique(books)
	log.Println(books)

	// 拉取更新
	cmd = exec.Command("sh", "-c", repositoryDirCmd+"git pull")
	err = cmd.Run()
	if err != nil {
		printError(err)
		return
	}

	// 构建book
	var bookPath string
	for _, book := range books {
		// 跳过不存在的book
		bookPath = *repositoryDir + PS + book
		if !dirExist(bookPath) {
			continue
		}

		// 构建book
		log.Println("build book:", bookPath)
		bookPathCmd := "cd " + bookPath + " && "
		cmd = exec.Command("sh", "-c", bookPathCmd+"gitbook build")
		err = cmd.Run()
		if err != nil {
			log.Println("faild")
			continue
		}
		cmd = exec.Command("sh", "-c", bookPathCmd+"gitbook install")
		cmd.Run()

		// 移动构建好的book
		dstBookPath := *bookDir + PS + book
		if !dirExist(dstBookPath) {
			err = os.MkdirAll(dstBookPath, 0755)
			if err != nil {
				printError(err)
				continue
			}
		}
		cmd = exec.Command("sh", "-c", bookPathCmd+"cp -auv _book/* "+dstBookPath)
		err = cmd.Run()
		printError(err)
	}
}

// 打印错误日志
func printError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

// 移除重复book
func booksUnique(books []string) []string {
	set := make(map[string]struct{}, len(books))
	i := 0
	for _, book := range books {
		_, ok := set[book]
		if ok {
			continue
		}
		set[book] = struct{}{}
		books[i] = strings.Replace(book, "/", "", -1)
		i++
	}
	return books[:i]
}

// 判断文件或目录是否存在
func fileExist(file string) bool {
	o, err := os.Stat(file)
	if err != nil {
		printError(err)
		return false
	}
	return !o.IsDir()
}

// 判断目录是否存在
func dirExist(dir string) bool {
	o, err := os.Stat(dir)
	if err != nil {
		return false
	}
	return o.IsDir()
}

// hmac sha1
func HMACSHA1(str string) string {
	o := hmac.New(sha1.New, []byte(config.Secret))
	o.Write([]byte(str))
	return hex.EncodeToString(o.Sum(nil))
}
