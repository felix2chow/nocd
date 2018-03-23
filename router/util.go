/*
 * Copyright (c) 2018, 奶爸<1@5.nu>
 * All rights reserved.
 */

package router

import (
	"github.com/gin-gonic/gin"
	"git.cm/naiba/gocd"
	"html/template"
	"github.com/utrack/gin-csrf"
)

func setCookie(c *gin.Context, key string, val string) {
	c.SetCookie(key, val, 60*60*24*365*1.5, "/", "", false, false)
}

func commonData(c *gin.Context, csrfToken bool, data gin.H) gin.H {
	data["domain"] = gocd.Conf.Section("gocd").Key("domain").String()
	isLogin := c.GetBool(CtxIsLogin)
	data["isLogin"] = isLogin
	if isLogin {
		data["user"] = c.MustGet(CtxUser)
	}
	if csrfToken {
		data["csrf_token"] = template.HTML(`<input type="hidden" name="_csrf" value="` + csrf.GetToken(c) + `">`)
	}
	return data
}

func jsAlertAndRedirect(msg, url string, c *gin.Context) {
	c.Writer.WriteString(`
<script>
alert('` + msg + `');window.location.href='` + url + `'
</script>
`)
	c.Abort()
}
