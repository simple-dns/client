package daemon

import (
	"bytes"
	"context"
	"dns-client/command/path"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

type ResponseInfo struct {
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
}

var lastIp net.IP

func Run(ctx context.Context) error {

	config, err := getConfig()
	if err != nil {
		return err
	}

	client := &http.Client{}
	tick := time.Tick(5 * time.Second)
	for {
		select {
		case <-tick:
			doReport(ctx, config, client)
		case <-ctx.Done():
			return nil
		}
	}
}

func doReport(ctx context.Context, config *Config, client *http.Client) {

	if checkServerHealth(config.Server.Address) {

		var ip net.IP
		for ip = getLocalIp(config.Server.Host); ip == nil; {
		}

		byteSlice, err := ioutil.ReadFile(filepath.Join(path.ConfigDirPath(), config.Domain.DomainFile))
		if err != nil {
			log.Println(err.Error())
		}
		domains := strings.Fields(strings.TrimSpace(string(byteSlice)))
		domains = append(domains, "")
		for _, domain := range domains {
			doPost(ctx, config, client, domain, ip)
		}
	}
}

func doPost(ctx context.Context, config *Config, client *http.Client, domain string, ip net.IP) {
	name := buildDomainName(domain, config)
	if name == "" {
	}
	params := map[string]string{
		"name": name,
		"ip":   ip.String(),
	}
	jsonByte, err := json.Marshal(params)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	request, err := http.NewRequestWithContext(ctx, "POST", config.Server.Address+"/rpc/record", bytes.NewBuffer(jsonByte))
	if err != nil {
		log.Printf(err.Error())
		return
	}
	request.Header.Set("X-Dns-Token", config.Server.Token)
	resp, err := client.Do(request)
	if err != nil {
		log.Printf("domain: %v add fail，error info : %v\n", domain, err.Error())
		return
	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Printf("domain: %v add fail，error info : %v\n", domain, err.Error())
		return
	}
	log.Printf("resp body :%v", string(body))
	resp.Body.Close()
}

func buildDomainName(domain string, config *Config) string {
	if domain != "" {
		return domain + config.Domain.Suffix
	} else {
		if config.Domain.Suffix != "" {
			return config.Domain.Suffix[1:]
		} else {
			return ""
		}
	}
}

func checkServerHealth(serverAddr string) bool {
	serverUrl := serverAddr + "/healthz"
	res, err := http.Get(serverUrl)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	if res.StatusCode != 200 {
		log.Println("server error")
		return false
	}
	return true
}

func getLocalIp(serverHost string) net.IP {

	serverIp := net.ParseIP(serverHost)
	var addr *net.IPAddr
	if serverIp == nil {
		var err error
		for addr, err = net.ResolveIPAddr("ip", serverHost); err != nil; {
			log.Printf(err.Error())
		}
		serverIp = addr.IP
	}

	adds, err := net.InterfaceAddrs()
	if err != nil {
		log.Println(err.Error())
	}

	for _, addr := range adds {
		if ip, ok := addr.(*net.IPNet); ok && !ip.IP.IsLoopback() {
			localIp := ip.IP.To4()
			if localIp != nil && ip.Contains(serverIp) {
				if localIp.String() != lastIp.String() {
					lastIp = localIp
					log.Printf("get local ip success %v", ip)
				}
				return localIp
			}
		}
	}
	log.Printf("get local ip fail")
	return nil
}
