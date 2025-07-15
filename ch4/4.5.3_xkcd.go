package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const xkcdURL string = "https://xkcd.com/571/info.0.json"

type xkcd struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

func xkcd1() {
	resp, err := http.Get(xkcdURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "get failure:%v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, "http response failure:%s", resp.Status)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read response failure:%v", err)
	}
	var result xkcd
	if err := json.Unmarshal(b, &result); err != nil {
		fmt.Fprintf(os.Stderr, "unmarshal json failure:%v", err)
	}

	fmt.Printf("Title: %s\n", result.Title)
	fmt.Printf("Number: %d\n", result.Num)
	fmt.Printf("Date: %s-%s-%s\n", result.Year, result.Month, result.Day)
	fmt.Printf("Image URL: %s\n", result.Img)
	fmt.Printf("Alt Text: %s\n", result.Alt)
	fmt.Printf("Safe Title: %s\n", result.SafeTitle)
	// fmt.Printf("Transcript: %s\n", result.Transcript)
	fmt.Printf("Link: %s\n", result.Link)
	fmt.Printf("News: %s\n", result.News)
}

// 需求分析:
// 1. 下载所有 XKCD 漫画的 JSON 数据：
// (1) 从 https://xkcd.com/info.0.json 获取最新漫画编号。
// (2) 遍历从 1 到最新编号，下载每个漫画的 JSON（如 https://xkcd.com/{num}/info.0.json）。
// (3) 确保每个链接只下载一次，避免重复请求。
// 2. 创建离线索引：
// (1) 将下载的 JSON 数据存储到本地文件（例如 xkcd_data/{num}.json）。
// (2) 创建一个索引文件，存储漫画编号与标题、描述、转录等字段的映射，便于快速搜索。
// (3) 例如 index.json），包含所有漫画的元数据（编号、标题、转录、描述等）
// 3. 实现搜索功能：
// (1) 接受命令行输入的检索词。
// (2) 在离线索引中搜索匹配的漫画（例如，标题、转录或描述中包含检索词）。
// (3) 接受命令行参数（如 search terms），在索引中搜索 title, safe_title, transcript, alt 字段。
// (4) 返回匹配漫画的 URL（img 字段）。
// 4. 优化与错误处理：
// (1) XKCD 编号从 1 开始，但某些编号（如 404）不存在，需处理 404 错误。
// (2) 下载可能耗时，需添加延时或并发控制以避免触发速率限制。
// (3) 确保只下载一次，使用本地文件检查避免重复请求。

// 实现步骤
// (1) 定义数据结构：
// * 复用 xkcd 结构体，字段需导出以支持 JSON 解码。
// * 定义索引结构，存储漫画元数据。
// (2) 下载漫画数据：
// * 获取最新漫画编号。
// * 遍历所有编号，下载 JSON 文件，跳过已存在的文件。
// * 处理 404 错误和其他异常。
// (3) 创建索引：
// * 解析每个 JSON 文件，提取关键字段（num, title, safe_title, transcript, alt, img）。
// * 保存到索引文件（index.json）。
// (4) 实现搜索功能：
// * 解析命令行参数（检索词）。
// * 加载索引，搜索匹配的漫画，打印 img URL。
// (5) 命令行接口：
// * 支持子命令（如 download 和 search）。
// * 使用 flag 包解析参数。

const latestURL = "https://xkcd.com/info.0.json"

// 第一步: 根据http地址, 解析最新的网址json字符串, 解析到结构体中
func getLatestComic() (*xkcd, error) {
	resp, err := http.Get(latestURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest comic: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get latest comic: %s,body: %s", resp.Status, body)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	var latestComic xkcd
	if err := json.Unmarshal(body, &latestComic); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}
	return &latestComic, nil
}

// 第二步: 根据文件路径, 给定一个具体的文件路径, 解析文件json, 返回这个具体文件的结构体实例
func readComic(filePath string) (*xkcd, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	var comic xkcd
	if err := json.Unmarshal(data, &comic); err != nil {
		return nil, fmt.Errorf("failed to unmarshal file: %w", err)
	}
	return &comic, nil
}

// 这个结构体用于记录所有 num 的, 相关索引信息
type IndexEntry struct {
	Num        int    `json:"num"`
	Title      string `json:"title"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
}

const dataDir1 = "xkcd_data"
const comicURL = "https://xkcd.com/%d/info.0.json"
const indexFile = "index.json"

func downloadComics() error {
	if err := os.MkdirAll(dataDir2, 0755); err != nil {
		return fmt.Errorf("failed to create data directory: %w", err)
	}
	//
	latest, err := getLatestComic()
	if err != nil {
		return fmt.Errorf("failed to get latest comic: %w", err)
	}
	fmt.Printf("Latest comic number:%d\n", (*latest).Num)

	var comics []IndexEntry
	// 从num=1 一直遍历到最新的 num, 循环有两个目的
	// (1) 遍历每一个 num 的json串写到文件中, (2) 把每一个json 串的索引信息写到 []IndexEntry 切片中
	for num := 3101; num <= (*latest).Num; num++ {
		// 分支流程1: num=404
		if num == 404 {
			continue
		}
		// 分支流程2: 如果文件存在
		filePath := filepath.Join(dataDir2, fmt.Sprintf("%d.json", num))

		if _, err := os.Stat(filePath); err == nil {
			fmt.Printf("Comic #%d has been already downloaded\n", num)
			comic, err := readComic(filePath)
			if err != nil {
				fmt.Printf("Warning: failed to read comic #%d: %v\n", num, err)
				continue
			}
			// 当前内存中的变量 comics 是空的, 即使文件存在, 当前内存中也是没有内容的
			// 对于已经存在的文件, 也要读取起内容, 追加到 comics 中, 否则 comics 不会包含这些文件的元数据
			// comics []IndexEntry 的目的是在当前运行中构建一个完整的漫画元数据索引，包含所有漫画（现有和新下载的）
			comics = append(comics, IndexEntry{
				Num:        comic.Num,
				Title:      comic.Title,
				SafeTitle:  comic.SafeTitle,
				Transcript: comic.Transcript,
				Alt:        comic.Alt,
				Img:        comic.Img,
			})
			continue
		}

		// 一般的流程, 遍历到每个 num, 请求每一个 num 的json字符串, 把响应封装到一个xkcd 结构体实例里面
		url := fmt.Sprintf(comicURL, num)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("faild to get num #%d comics: %v\n", num, err)
			continue
		}
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusNotFound {
			fmt.Printf("Comic #%d not found(404)\n", num)
			continue
		}
		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("Warning: failed to get comic #%d: %s, body: %s\n", num, resp.Status, body)
			continue
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Warning: failed to get comic #%d: %s, body: %s\n", num, resp.Status, body)
			continue
		}
		var comic xkcd
		if err := json.Unmarshal(body, &comic); err != nil {
			fmt.Printf("Warning: failed to unmarshal comic #%d: %v\n", num, err)
			continue
		}
		// 得到一个 xkcd 结构体实例, 对它做保存到文件
		if err := os.WriteFile(filePath, body, 0644); err != nil {
			fmt.Printf("Warning: failed to save comic #%d: %v\n", num, err)
			continue
		}
		fmt.Printf("Download comic #%d: %s\n", comic.Num, comic.Title)

		// 添加到[]IndexEntry, 这个切片 里面记录了所有的 num 的索引信息
		comics = append(comics, IndexEntry{
			Num:        comic.Num,
			Title:      comic.Title,
			SafeTitle:  comic.SafeTitle,
			Transcript: comic.Transcript,
			Alt:        comic.Alt,
			Img:        comic.Img,
		})
		// Avoid rate limiting
		time.Sleep(500 * time.Millisecond)
	}
	// 把索引切片编码成 json 字符串, 写入文件
	indexData, err := json.MarshalIndent(comics, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal index: %w", err)
	}
	if err := os.WriteFile(indexFile, indexData, 0644); err != nil {
		return fmt.Errorf("faild to save index: %w", err)
	}

	fmt.Printf("Index saved to %s\n", indexFile)
	return nil
}

func searchComics(terms []string) error {
	data, err := os.ReadFile(indexFile)
	if err != nil {
		return fmt.Errorf("failed to read index: %w", err)
	}

	var index []IndexEntry
	if err := json.Unmarshal(data, &index); err != nil {
		return fmt.Errorf("failed to unmarshal json: %w", err)
	}
	matchedCount := 0
	for _, entry := range index {
		matched := false
		for _, term := range terms {
			term = strings.TrimSpace(term)
			if term == "" {
				continue
			}
			if strings.Contains(strings.ToLower(entry.Title), term) ||
				strings.Contains(strings.ToLower(entry.SafeTitle), term) ||
				strings.Contains(strings.ToLower(entry.Transcript), term) ||
				strings.Contains(strings.ToLower(entry.Alt), term) {
				matched = true
				break
			}
		}
		if matched {
			fmt.Printf("Comic #%d: %s\nURL: %s\n\n", entry.Num, entry.Title, entry.Img)
			matchedCount++
		}
	}
	if matchedCount == 0 {
		fmt.Printf("Comic Not Found\n")
	}
	return nil
}
