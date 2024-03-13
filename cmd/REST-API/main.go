package main

import (
  "log"
  "fmt"

	"github.com/andyfilya/restapi/config"
  "github.com/andyfilya/restapi/internal/server"
)

func main() {
  cfg, err := config.InitGlobalConfig()
  if err != nil {
    log.Fatal(err)
  }
  server := server.InitServer(&cfg.ServCfg)
  fmt.Println(server)
  server.Init()
}
