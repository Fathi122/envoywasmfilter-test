package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

// pluginContext
type pluginContext struct {
	// Embed the default plugin context here,
	types.DefaultPluginContext
}

// Override types.DefaultPluginContext.
func (*pluginContext) NewHttpContext(contextID uint32) types.HttpContext {
	return &httpHeadersHttpContext{contextID: contextID}
}

// vmContext
type vmContext struct {
	// Embed the default VM context here,
	types.DefaultVMContext
}

// Override types.DefaultVMContext.
func (*vmContext) NewPluginContext(contextID uint32) types.PluginContext {
	return &pluginContext{}
}

// httpHeadersHttpContext
type httpHeadersHttpContext struct {
	// Embed the default http context here,
	types.DefaultHttpContext
	contextID uint32
}

// on Http Request Handler
func (ctx *httpHeadersHttpContext) OnHttpRequestHeaders(numHeaders int, _ bool) types.Action {
	if numHeaders > 0 {
		headers, err := proxywasm.GetHttpRequestHeaders()
		if err != nil {
			proxywasm.LogErrorf("failed to get request headers with '%v'", err)
			return types.ActionContinue
		}
		proxywasm.LogInfof("On request headers: '%+v'", headers)
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
	proxywasm.LogInfof("OnLog: :path = %s", hdr)
}

func main() {
	proxywasm.SetVMContext(&vmContext{})
}
