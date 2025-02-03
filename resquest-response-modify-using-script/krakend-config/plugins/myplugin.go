

package main

import (
	"github.com/luraproject/lura/v2/proxy"
)

func ModifyRequest(req *proxy.Request) {
	req.Headers["X-Go-Header"] = []string{"ModifiedByGo"}
}

func ModifyResponse(res *proxy.Response) {
	res.Data["message_by_go"] = "Modified by Go Plugin"
}

var Plugin = struct {
	Pre  func(*proxy.Request)
	Post func(*proxy.Response)
}{
	Pre:  ModifyRequest,
	Post: ModifyResponse,
}

