package hello

import (
	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
)

// init 函数注册插件
func init() {
	plugin.Register("hello", setup)
}

// setup 函数解析配置
// setup is the function that gets called when the config parser see the token "hello". Setup is responsible
// for parsing any extra options the hello plugin may have. The first token this function sees is "hello".
func setup(c *caddy.Controller) error {
	c.Next() //  Ignore "hello" and give us the next token. By MZTfan 20251030 12:06

	if c.NextArg() {
		// If there was another token, return an error, because we don't have any configuration.
		// Any errors returned from this setup function should be wrapped with plugin.Error, so we
		// can present a slightly nicer error message to the user.
		return plugin.Error("hello", c.ArgErr())
	}

	// Add the Plugin to CoreDNS, so Servers can use it in their plugin chain.
	// By MZTfan 20251030 func(下一个插件) 返回(当前插件)
	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return Hello{Next: next}
	})

	// All OK, return a nil error.
	return nil
}
