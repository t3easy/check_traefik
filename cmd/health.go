/*
Copyright © 2022 Jan Kiesewetter <jan@t3easy.de>

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/NETWAYS/go-check"
	"github.com/spf13/cobra"
)

// healthCmd represents the health command
var (
	path      string
	healthCmd = &cobra.Command{
		Use:     "health",
		Short:   "Checking the health of your Traefik instance",
		Version: version,
		Run: func(cmd *cobra.Command, args []string) {
			var hostPort string
			schema := "http"
			if ssl {
				schema = "https"
				if port == 80 {
					port = 443
				}
			}
			if (ssl && port == 443) || (!ssl && port == 80) {
				hostPort = ""
			} else {
				hostPort = ":" + strconv.Itoa(port)
			}

			healthUrl := &url.URL{
				Scheme: schema,
				Host:   ip.String() + hostPort,
				Path:   path,
			}

			tr := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
			}

			client := &http.Client{
				Timeout:   time.Duration(timeout) * time.Second,
				Transport: tr,
			}
			req, err := http.NewRequest(http.MethodHead, healthUrl.String(), nil)
			if err != nil {
				check.ExitError(err)
			}

			if hostname != "" {
				req.Host = hostname
			}

			if username != "" && password != "" {
				req.SetBasicAuth(username, password)
			}

			resp, err := client.Do(req)
			if resp != nil {
				resp.Body.Close()
			}
			if err != nil {
				check.ExitError(err)
			}

			rc := check.Critical
			if resp.StatusCode == http.StatusOK {
				rc = check.OK
			}

			check.Exitf(
				rc,
				"%s returned %s",
				healthUrl.String(),
				resp.Status,
			)
		},
	}
)

func init() {
	rootCmd.AddCommand(healthCmd)
	healthCmd.PersistentFlags().StringVarP(&path, "url", "u", "/ping", "URL of the Traefik health-check endpoint")
}
