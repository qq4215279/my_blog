/**
 * markdown
 * @author liuzhen
 * @Description
 * @version 1.0.0 2021/2/1 15:19
 */
package utils

import (
	"github.com/88250/lute"
	"sync"
)

var (
	engine *lute.Lute
	once   sync.Once
)

// 初始化引擎
func InitEngine() *lute.Lute {
	once.Do(func() {
		engine = lute.New(func(lute *lute.Lute) {
			lute.SetToC(true)
			lute.SetGFMTaskListItem(true)
		})
	})
	return engine
}

// MarkDown字符串转Html字符串
func MD2HTML(markdownStr string) string {
	if IsBlank(markdownStr) {
		return ""
	}
	return InitEngine().MarkdownStr("", markdownStr)
}

// Html字符串转MarkDown字符串
func Html2MD(htmlStr string) string {
	if IsBlank(htmlStr) {
		return ""
	}
	return InitEngine().HTML2Md(htmlStr)
}
