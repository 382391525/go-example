package lreptile

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	reQQEmail = `(\d+)@qq.com`
)

// HandleError 处理异常
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

func GetEmail() {
	resp, err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	HandleError(err, "http.get url")
	defer resp.Body.Close()
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	pageStr := string(pageBytes)
	re := regexp.MustCompile(reQQEmail)
	// -1代表取全部
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println("email:", result[0])
		fmt.Println("qq:", result[1])
	}
}

func DownloadFile(url string, filename string) (ok bool) {
	resp, err := http.Get(url)
	HandleError(err, "http.Get(url)")
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll(resp.Body)")
	pwd, _ := os.Getwd()
	filename = pwd + "/lreptile/" + filename
	err = ioutil.WriteFile(filename, bytes, 0666)
	if err != nil {
		return false
	}
	return true
}

func DownloadImg() {
	for imageUrl := range chanImageUrls {
		filename := GetFilenameFromUrl(imageUrl)
		ok := DownloadFile(imageUrl, filename)
		if ok {
			fmt.Printf("%s 下载成功\n", filename)
		} else {
			fmt.Printf("%s 下载失败\n", filename)
		}
	}
	wg.Done()
}

// GetFilenameFromUrl 截取url名字
func GetFilenameFromUrl(url string) (filename string) {
	// 返回最后一个/的位置
	lastIndex := strings.LastIndex(url, "/")
	// 切出来
	filename = url[lastIndex+1:]
	// 时间戳解决重名
	timePrefix := strconv.Itoa(int(time.Now().UnixNano()))
	filename = timePrefix + "_" + filename
	return
}

var (
	chanImageUrls chan string
	wg            sync.WaitGroup
	chanTask      chan string
	reImg         = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

func Concurrent() {
	chanImageUrls = make(chan string, 100000)
	chanTask = make(chan string, 26)

	for i := 1; i < 27; i++ {
		wg.Add(1)
		go getImgUrls("https://www.bizhizu.cn/shouji/tag-%E5%8F%AF%E7%88%B1/" + strconv.Itoa(i) + ".html")
	}
	//检测26个任务是否完成
	wg.Add(1)
	go checkTask()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go DownloadImg()
	}
	wg.Wait()
}

func checkTask() {
	var count int
	for {
		url := <-chanTask
		fmt.Printf("%s 完成了爬取任务\n", url)
		count++
		if count == 26 {
			close(chanImageUrls)
			break
		}
	}
	wg.Done()
}

func getImgUrls(url string) {
	urls := getImages(url)
	for _, v := range urls {
		chanImageUrls <- v
	}
	chanTask <- url
	wg.Done()
}

//获取images
func getImages(url string) (urls []string) {
	pageStr := GetPageStr(url)
	reg := regexp.MustCompile(reImg)
	results := reg.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d条结果\n", len(results))
	for _, result := range results {
		url := result[0]
		urls = append(urls, url)
	}
	return urls
}

// GetPageStr 抽取根据url获取内容
func GetPageStr(url string) (pageStr string) {
	resp, err := http.Get(url)
	HandleError(err, "http.Get url")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	// 字节转字符串
	pageStr = string(pageBytes)
	return pageStr
}
