package publicconfighttp

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	presenceMu  sync.Mutex
	presenceMap = make(map[string]int64)
)

// PostPresence 记录访问者心跳并返回当前实时在线人数。
// 请求体：{"visitor_id":"<随机 UUID>"}；90 秒无心跳视为离线。
func (h *Handler) PostPresence(c *gin.Context) {
	var req struct {
		VisitorID string `json:"visitor_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.VisitorID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "visitor_id required"})
		return
	}
	now := time.Now().Unix()
	presenceMu.Lock()
	presenceMap[req.VisitorID] = now
	for k, ts := range presenceMap {
		if now-ts > 90 {
			delete(presenceMap, k)
		}
	}
	count := len(presenceMap)
	presenceMu.Unlock()
	c.JSON(http.StatusOK, gin.H{"count": count})
}
