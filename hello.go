package hello

import (
	"context"

	"github.com/coredns/coredns/plugin"
	clog "github.com/coredns/coredns/plugin/pkg/log"
	"github.com/miekg/dns"
)

// 日志器
var log = clog.NewWithPlugin("hello")

type Hello struct {
	Next plugin.Handler
}

func (h Hello) Name() string {
	return "hello"
}

func (h Hello) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {

	log.Info("Hello ServeDNS")

	// 可选：如果想返回 TXT 记录作为响应，而不是只打印日志
	// m := new(dns.Msg)
	// m.SetReply(r)
	// m.Authoritative = true
	// txt := "Hello World"
	// rr := &dns.TXT{
	//     Hdr: dns.RR_Header{Name: r.Question[0].Name, Rrtype: dns.TypeTXT, Class: dns.ClassINET, Ttl: 60},
	//     Txt: []string{txt},
	// }
	// m.Answer = append(m.Answer, rr)
	// w.WriteMsg(m)
	// return dns.RcodeSuccess, nil

	// ???
	return plugin.NextOrFailure(h.Name(), h.Next, ctx, w, r)
}
