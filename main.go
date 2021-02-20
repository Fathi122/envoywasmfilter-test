package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

type httpHeadersHttpContext struct {
	proxywasm.DefaultHttpContext
}

func main() {
	// set http context
	proxywasm.SetNewHttpContext(newHttpContext)
}

func newHttpContext(uint32, uint32) proxywasm.HttpContext {
	return &httpHeadersHttpContext{}
}

// on Http Request Handler
func (ctx *httpHeadersHttpContext) OnHttpRequestHeaders(numHeaders int, _ bool) types.Action {

	if numHeaders > 0 {
		headers, err := proxywasm.GetHttpRequestHeaders()
		if err != nil {
			proxywasm.LogErrorf("failed to get request headers with '%v'", err)
			return types.ActionContinue
		}
		proxywasm.LogInfof("tutut request headers: '%+v'", headers)
	}

	return types.ActionContinue
}

// on Log Handler
func (ctx *httpHeadersHttpContext) OnLog() {
	hdr, err := proxywasm.GetHttpRequestHeader(":path")
	if err != nil {
		proxywasm.LogCritical(err.Error())
		return
	}
	proxywasm.LogInfof("tatat OnLog: :path = %s", hdr)
}
