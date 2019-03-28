package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net"
	"net/http"
)
func main()  {
	r := gin.Default()
/*    dir ,err := filepath.Abs("./ipip/tops_77_imminent/*")
    if err != nil{
    	log.Fatal(err)
	}*/
	r.GET("/ip", getip)
	r.LoadHTMLGlob("templates/*.html")
    r.Static("/css","static/css")
	r.Static("/fonts","static/fonts")
	r.Static("/js","static/js")
	r.Static("/images","static/images")
// r.LoadHTMLFiles("templates/index.html")


	r.GET("/index", getindex)
    r.Run(":88")
}
func getindex(c *gin.Context){
	ccip := c.ClientIP()
    c.HTML(http.StatusOK, "index.html",gin.H{
   	"title":ccip,
   })
}
type Response struct {
	ip string `json:"ip"`
	Country   string `json:"country"`
	Region  string `json:"region"`
	City      string `json:"city"`
	isp  string  `json:"isp"`
}

type Response1 struct {
   Code int
   Data map[string]interface{}
}
func getip(g * gin.Context) {
	var ip string
	xip := g.Request.Header.Get("X-Forwarded-For")
	fmt.Println(xip)
	if xip != "" {
		ip = xip
	} else {
		ip = g.ClientIP()
	}
	eip := net.ParseIP(ip)
	g.JSON(200,gin.H{
		"ip":eip,
	})
	dirurl := fmt.Sprintf("http://ip.taobao.com/service/getIpInfo.php?ip=%s", eip)
	resp, err := http.Get(dirurl)
	if err != nil {
		fmt.Printf("err")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err")
	}

//	var dat map[string]interface{}
    dat := Response1{}
	// 解码过程，并检测相关可能存在的错误
	if err := json.Unmarshal([]byte(body), &dat); err != nil {
		panic(err)
	}
	fmt.Println("code", dat.Code)
    fmt.Println("ip", dat.Data["ip"].(string))
    if dat.Code == 0 {
		g.JSON(200, gin.H{
			"ip":  dat.Data["ip"].(string),
			"城市":  dat.Data["city"].(string),
			"省份":  dat.Data["region"].(string),
			"国家":  dat.Data["country"].(string),
			"运营商": dat.Data["isp"].(string),
			"other": Response{},
		})

	}else {
		g.JSON(404, gin.Error{})
	}


/*	for x ,y := range dat.Data{
		fmt.Println(x,":",y)
	}*/

//	fmt.Println(dat)

}

