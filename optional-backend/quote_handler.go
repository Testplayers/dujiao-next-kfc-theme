package publicconfighttp

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// 语录文件路径：默认 data/kfc_quotes.json（相对进程工作目录），可用环境变量 KFC_QUOTE_FILE 覆盖。
// 文件格式：[{"text":"..."},{"text":"..."}]
const defaultQuotePath = "data/kfc_quotes.json"

type kfcQuoteItem struct {
	Text string `json:"text"`
}

var (
	quoteCacheMu   sync.Mutex
	quoteCacheData []string
	quoteCacheAt   time.Time
)

func quoteFilePath() string {
	if p := os.Getenv("KFC_QUOTE_FILE"); p != "" {
		return p
	}
	return defaultQuotePath
}

func loadQuotes() []string {
	quoteCacheMu.Lock()
	defer quoteCacheMu.Unlock()
	if quoteCacheData != nil && time.Since(quoteCacheAt) < 10*time.Minute {
		return quoteCacheData
	}
	data, err := os.ReadFile(quoteFilePath())
	if err != nil {
		return quoteCacheData
	}
	var items []kfcQuoteItem
	if err := json.Unmarshal(data, &items); err != nil {
		return quoteCacheData
	}
	quotes := make([]string, 0, len(items))
	for _, it := range items {
		if it.Text != "" {
			quotes = append(quotes, it.Text)
		}
	}
	if len(quotes) > 0 {
		quoteCacheData = quotes
		quoteCacheAt = time.Now()
	}
	return quoteCacheData
}

// GetDailyQuote 返回随机一条疯狂星期四语录；语录文件缺失时返回内置兜底文案。
func (h *Handler) GetDailyQuote(c *gin.Context) {
	text := "不疯狂不星期四 · 精神状态遥遥领先"
	if quotes := loadQuotes(); len(quotes) > 0 {
		text = quotes[rand.Intn(len(quotes))]
	}
	c.JSON(http.StatusOK, gin.H{"quote": text})
}
