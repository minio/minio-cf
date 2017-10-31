/*
 * Minio tool to be used with Pivotal "cf" tool to generate the
 * config for create-service command.
 */

package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/minio/cli"
)

func main() {
	app := cli.NewApp()
	app.Usage = "Minio CF Tool"
	app.Author = "https://minio.io"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "access-key",
		},
		cli.StringFlag{
			Name: "secret-key",
		},
		cli.StringFlag{
			Name: "gcs",
		},
		cli.BoolFlag{
			Name: "azure",
		},
		cli.BoolFlag{
			Name: "s3",
		},
		cli.StringFlag{
			Name: "subdomain",
		},
	}
	app.CustomAppHelpTemplate = `NAME:
  {{.HelpName}} - {{.Usage}}

USAGE:
  {{.HelpName}} [FLAGS]

FLAGS:
  {{range .VisibleFlags}}{{.}}
  {{end}}

EXAMPLES:
  1. Generate config file for Minio Server instance
     $ {{.HelpName}} --access-key minio --secret-key minio123 > minio-server-pcf.conf

  2. Generate config file for Minio GCS gateway instance
     $ {{.HelpName}} --access-key minio --secret-key minio123 --gcs /path/to/credentials.json > minio-gcs-pcf.conf

  3. Generate config file for Minio Azure gateway instance
     $ {{.HelpName}} --access-key azureaccountname --secret-key azureaccountkey > minio-azure-pcf.conf

`

	app.Action = func(ctx *cli.Context) {
		if len(ctx.Args()) == 0 {
			cli.ShowAppHelpAndExit(ctx, 1)
		}
		m := make(map[string]string)
		m["accesskey"] = ctx.GlobalString("access-key")
		m["secretkey"] = ctx.GlobalString("secret-key")
		subdomain := ctx.GlobalString("subdomain")
		if subdomain != "" {
			m["subdomain"] = subdomain
		}
		if ctx.GlobalBool("azure") {
			m["gateway"] = "azure"
		}
		if ctx.GlobalBool("s3") {
			m["gateway"] = "s3"
		}
		if ctx.GlobalString("gcs") != "" {
			m["gateway"] = "gcs"
			b, err := ioutil.ReadFile(ctx.GlobalString("gcs"))
			if err != nil {
				log.Fatal(err)
			}
			m["googlecredentials"] = base64.StdEncoding.EncodeToString(b)
		}
		b, err := json.Marshal(m)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(`%s`, string(b))
	}
	app.Run(os.Args)
}
