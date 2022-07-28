package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var dummy string

func LoadCredentials() ([]reddit.Credentials, time.Time) {
	var tempacc reddit.Credentials
	var accounts []reddit.Credentials
	var iter int
	// Use viper library
	viper.SetConfigName("credentials")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("Lütfen credentials.env de kaç adet hesap girildiğini belirtin (40,25,30 etc.): ")
	_, _ = fmt.Scanf("%d", &iter)
	start := time.Now()
	if err != nil {
		panic(err)
	}
	for j := 0; j < iter; j++ {
		iterStr := strconv.Itoa(j)
		if j == 0 {
			tempacc = reddit.Credentials{
				ID:       viper.Get("CLIENT_ID").(string),
				Secret:   viper.Get("CLIENT_SECRET").(string),
				Username: viper.Get("USERNAMEREDDIT").(string),
				Password: viper.Get("PASSWORD").(string),
			}
		} else {
			tempacc = reddit.Credentials{
				ID:       viper.Get("CLIENT_ID" + "_" + iterStr).(string),
				Secret:   viper.Get("CLIENT_SECRET" + "_" + iterStr).(string),
				Username: viper.Get("USERNAMEREDDIT" + "_" + iterStr).(string),
				Password: viper.Get("PASSWORD" + "_" + iterStr).(string),
			}
		}
		// Append tempacc to accounts[j]
		accounts = append(accounts, tempacc)
	}
	return accounts, start
}
func PostID() string {
	viper.SetConfigName("post")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return viper.Get("POST_ID").(string)
}

// SetProxy Sets a random proxy IP address from proxy.env file.
func SetProxy(usedproxy string) string {
	viper.SetConfigName("proxy")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	// Pick a random number between 1-12
	num := rand.Intn(4) + 1
	proxyName := "PROXY_" + strconv.Itoa(num)
	proxyIP := viper.Get(proxyName)
	for {
		if proxyIP.(string) == usedproxy {
			num = rand.Intn(4) + 1
			proxyName = "PROXY_" + strconv.Itoa(num)
			proxyIP = viper.Get(proxyName)
		} else {
			break
		}
	}
	fmt.Println("Proxy şu adresten alındı: ", proxyIP.(string))
	if err != nil {
		panic(err)
	}
	return proxyIP.(string)
}

func ClientsUpvote(clients []*reddit.Client, usedproxy string) {
	postId := PostID()
	for _, client := range clients {
		pc, resp, err := client.Post.Get(context.Background(), postId)
		if err != nil || resp.StatusCode != 200 {
			panic(err)
		}
		resp, err = client.Post.Upvote(context.Background(), pc.Post.FullID)
		if err != nil || resp.StatusCode != 200 {
			panic(err)
		}
		// Wait a random amount of time between 5-20 seconds
		fmt.Println("Şu hesaptan upvote atıldı: ", client.Username)
		fmt.Println("1 ila 10 saniye arasında bir süre bekleniyor...")
		time.Sleep(time.Duration(rand.Intn(20)+5) * time.Second)
	}
}
func main() {
	// Keep a runtime timer
	accounts, start := LoadCredentials()
	var proxyIP string
	proxyIP = SetProxy(dummy)
	// Hesapları yazdırma
	for i := 0; i < len(accounts); i++ {
		println("Hesap " + accounts[i].Username + ":")
		fmt.Println("Hesap API ID> ", accounts[i].ID)
		fmt.Println("Hesap API Secret> ", accounts[i].Secret)
		fmt.Println("Hesap İsmi> ", accounts[i].Username)
		fmt.Println("Hesap Şifresi> ", accounts[i].Password)
		fmt.Println("-----------------")
	}

	// Create a new reddit client for each account
	var clients []*reddit.Client
	for i := 0; i < len(accounts); i++ {
		// EĞER MÜŞTERİ BÜTÜN HESAPLAR İÇİN AYRI PROXY ALIRSA HER PROXY AYRI BİR CLIENTA ATANACAK
		// EĞER MÜŞTERİ BUNU YAPMAZSA, MECBUR RANDOM ATAYACAĞIZ PROXYLERİ.
		parsedProxyIp, err := url.Parse(proxyIP)
		httpClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(parsedProxyIp)}}
		tempclient, err := reddit.NewClient(accounts[i], reddit.WithHTTPClient(httpClient))
		if err != nil {
			panic(err)
		}
		clients = append(clients, tempclient)
		proxyIP = SetProxy(proxyIP)
	}
	ClientsUpvote(clients, proxyIP)
	// Print runtime
	fmt.Println("Çalışma süresi: ", time.Since(start))
}
