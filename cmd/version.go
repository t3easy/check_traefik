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
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/NETWAYS/go-check"
	"github.com/spf13/cobra"
	"github.com/t3easy/check_traefik/internal"
	"golang.org/x/mod/semver"
)

type Version struct {
	Version      string    `json:"Version"`
	Codename     string    `json:"Codename"`
	StartDate    time.Time `json:"startDate"`
	PilotEnabled bool      `json:"pilotEnabled"`
}

// versionCmd represents the version command
var (
	minVersion         string
	versionInformation Version
	versionCmd         = &cobra.Command{
		Use:     "version",
		Short:   "Check the version of your Traefik instance",
		Version: version,
		Run: func(cmd *cobra.Command, args []string) {
			req := internal.NewRequest(http.MethodGet, ip, hostname, ssl, port, path, username, password)
			req.Header.Set("Accept", "application/json")
			resp := internal.GetResp(req, timeout, insecure)
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				check.Exitf(
					check.Unknown,
					"%s returned %s",
					req.URL.String(),
					http.StatusText(resp.StatusCode),
				)
			}
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				check.ExitError(err)
			}

			if err := json.Unmarshal(b, &versionInformation); err != nil {
				check.ExitError(err)
			}

			checkVersion(versionInformation.Version, minVersion)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.PersistentFlags().StringVarP(&path, "url", "u", "/api/version", "Path of the Traefik API version endpoint")
	versionCmd.PersistentFlags().StringVar(&minVersion, "minVersion", "", "Minimum Traefik version")
}

func normalizeVersion(v string) string {
	if !strings.HasPrefix(v, "v") {
		v = "v" + v
	}
	if !semver.IsValid(v) {
		check.Exitf(
			check.Unknown,
			"%s is not a valid samver version.",
			v,
		)
	}
	return v
}

func checkVersion(v string, minV string) {
	v = normalizeVersion(v)
	minV = normalizeVersion(minV)

	minCompare := semver.Compare(v, minV)

	switch {
	case minCompare == -1:
		check.Exitf(
			check.Critical,
			"Traefik version is %s, which is lower than minimum version %s.",
			v,
			minV,
		)
	case minCompare == 0:
		check.Exitf(
			check.OK,
			"Traefik version is %s, which equals the minimum version %s.",
			v,
			minV,
		)
	case minCompare == 1:
		check.Exitf(
			check.OK,
			"Traefik version is %s, which is higher than the minimum version %s.",
			v,
			minV,
		)
	}
}
