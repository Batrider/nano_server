package main

import (
	"log"
	"net/http"
	"os"
	"github.com/urfave/cli"
	"github.com/lonnng/nano"
	"github.com/lonnng/nano/serialize/json"
)

func main() {

	app := cli.NewApp()
	app.Name = "Unity Move"
	app.Author = "nil"
	app.Version = "0.0.1"
	app.Copyright = "nil"
	app.Usage = "nil"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:"addr",
			Value: ":23456",
		},
	}
	app.Action = server

	app.Run(os.Args)
}

func server(ctx *cli.Context) {

	nano.Register(NewManager())
	nano.Register(NewWorld())
	nano.SetSerializer(json.NewSerializer())

	log.SetFlags(log.LstdFlags | log.Llongfile)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	nano.SetCheckOriginFunc(func(_ *http.Request) bool { return true })
	addr := ctx.String("addr")
	nano.Listen(addr)
}