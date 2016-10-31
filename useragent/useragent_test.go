package useragent

import (
	"testing"

	"github.com/JREAMLU/study/useragent"
	. "github.com/smartystreets/goconvey/convey"
)

func BenchmarkUseragent(b *testing.B) {
	var str = `Mozilla/5.0 (Windows NT 5.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.101 Safari/537.36 QIHU 360EE`
	Convey("bench Useragent()", b, func() {
		for i := 0; i < b.N; i++ {
			ParseByString(str)
		}
	})
}

func ParseUserAgent(ual string) map[string]interface{} {
	res := make(map[string]interface{})
	agent := useragent.ParseByString(ual)
	bot := false
	mobile := false

	if agent.Type == "robot" {
		bot = true
	}

	if agent.Device.Type == "mobile" {
		mobile = true
	}

	mu.Lock()
	res["browser_name"] = agent.Client["name"]
	res["browser_version"] = agent.Client["version"]
	res["platform"] = agent.OS.Name
	res["os"] = agent.OS.Version
	res["bot"] = bot
	res["mobile"] = mobile
	mu.Unlock()

	return res
}
