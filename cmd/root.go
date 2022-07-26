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
	"net"
	"time"

	"github.com/NETWAYS/go-check"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	version  string = "0.1.0"
	ip       net.IP
	hostname string
	port     int
	username string
	password string
	ssl      bool
	timeout  time.Duration
	insecure bool
	rootCmd  = &cobra.Command{
		Use:     "check_traefik",
		Short:   "Check Traefik",
		Version: version,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	var err error

	err = rootCmd.Execute()
	if err != nil {
		check.ExitError(err)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().IPVarP(&ip, "IP-address", "I", nil, "IP of the Traefik host to check (required)")
	rootCmd.MarkPersistentFlagRequired("IP-address")
	rootCmd.PersistentFlags().StringVarP(&hostname, "hostname", "H", "", "Hostname of the Traefik host to check")
	rootCmd.PersistentFlags().IntVarP(&port, "port", "P", 80, "Port of the Traefik API")
	rootCmd.PersistentFlags().StringVar(&username, "username", "", "User to access the Traefik health-check endpoint")
	rootCmd.PersistentFlags().StringVar(&password, "password", "", "Password to access the Traefik health-check endpoint")
	rootCmd.PersistentFlags().BoolVarP(&ssl, "ssl", "S", false, "Connect via SSL. Port defaults to 443.")
	rootCmd.PersistentFlags().DurationVarP(&timeout, "timeout", "T", 2, "Timeout in secounds")
	rootCmd.PersistentFlags().BoolVar(&insecure, "insecure", false, "If true accepts any certificate presented by the server and any host name in that certificate")
}
