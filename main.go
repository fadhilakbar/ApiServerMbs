package main

import (
	"ApiServerMbs/functions"
	"ApiServerMbs/routers"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kpango/glg"
	"github.com/vjeantet/jodaTime"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	_ = glg.Log("Application name : ", os.Getenv("APP_NAME"))
	var clog []string
	clog = append(clog, "Version App : "+os.Getenv("APP_VERSION")+"\n")
	_ = glg.Log("Connecting Database ...")
	err := functions.MysqlConnect(os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"))
	if err != nil {
		_ = glg.Error("Connection databases error : ", err.Error())
		os.Exit(1)
	}
	_ = glg.Log("Starting Services ...")
	err = Start()
	if err != nil {
		_ = glg.Log(err.Error())
	}
}

func Start() error {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(RequestLogger())
	routers.SetupRouter(router)
	portHttp := fmt.Sprint(":", os.Getenv("APP_PORT"))
	_ = glg.Log("[HTTP] Listening at ", portHttp)
	_ = router.Run(portHttp)

	return nil
}

func init() {
	_ = glg.Log("Loading configuration file....")
	err := godotenv.Load(".env")
	if err != nil {
		_ = glg.Error("Error loading configuration file!")
		os.Exit(1)
	}

	sCurrDate := jodaTime.Format("yyyyMMdd", time.Now())
	log := glg.FileWriter("log/apigombs_"+sCurrDate+".log", 0666)
	glg.Get().
		SetMode(glg.BOTH).
		AddLevelWriter(glg.LOG, log).
		AddLevelWriter(glg.ERR, log).
		AddLevelWriter(glg.WARN, log).
		AddLevelWriter(glg.DEBG, log).
		AddLevelWriter(glg.INFO, log)
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf, _ := ioutil.ReadAll(c.Request.Body)
		req, _ := functions.PrettyPrint(buf)
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(req))
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(req))
		fmt.Println("--------------------------REQUEST----------------------")
		fmt.Println(readBody(rdr1))
		c.Request.Body = rdr2
		c.Next()
	}
}

func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(reader)
	s := buf.String()
	return s
}
