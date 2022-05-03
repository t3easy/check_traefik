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
	"net/http"

	"github.com/NETWAYS/go-check"
	"github.com/spf13/cobra"
	"github.com/t3easy/check_traefik/internal"
)

// healthCmd represents the health command
var (
	path      string
	healthCmd = &cobra.Command{
		Use:     "health",
		Short:   "Checking the health of your Traefik instance",
		Version: version,
		Example: `check_traefik health -I 192.0.2.101 -H traefik.domain.tld --user="monitoring" --password="password"`,
		Run: func(cmd *cobra.Command, args []string) {
			req := internal.NewRequest(http.MethodHead, ip, hostname, ssl, port, path, username, password)

			resp := internal.GetResp(req, timeout, insecure)
			defer resp.Body.Close()

			rc := checkResponse(resp)

			check.Exitf(
				rc,
				"%s returned %s",
				req.URL.String(),
				resp.Status,
			)
		},
	}
)

func init() {
	rootCmd.AddCommand(healthCmd)
	healthCmd.PersistentFlags().StringVarP(&path, "url", "u", "/ping", "URL of the Traefik health-check endpoint")
}

func checkResponse(resp *http.Response) int {
	rc := check.Critical
	if resp.StatusCode == http.StatusOK {
		rc = check.OK
	}

	return rc
}
