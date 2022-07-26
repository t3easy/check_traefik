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
package internal

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/NETWAYS/go-check"
)

func NewRequest(method string, ip net.IP, hostname string, ssl bool, port int, path string, username string, password string) *http.Request {
	var (
		hostPort  string = ""
		schema    string = "http"
		healthUrl *url.URL
		req       *http.Request
		err       error
	)

	if ssl {
		schema = "https"
		if port == 80 {
			port = 443
		}
	}
	if (ssl && port != 443) || (!ssl && port != 80) {
		hostPort = ":" + strconv.Itoa(port)
	}
	healthUrl = &url.URL{
		Scheme: schema,
		Host:   ip.String() + hostPort,
		Path:   path,
	}

	req, err = http.NewRequest(method, healthUrl.String(), nil)
	if err != nil {
		check.ExitError(err)
	}
	if hostname != "" {
		req.Host = hostname
	}
	if username != "" && password != "" {
		req.SetBasicAuth(username, password)
	}

	return req
}

func GetResp(req *http.Request, timeout time.Duration, insecure bool) *http.Response {
	var (
		tr     *http.Transport
		client *http.Client
		resp   *http.Response
		err    error
	)

	tr = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: insecure,
			ServerName:         req.Host,
		},
	}

	client = &http.Client{
		Timeout:   timeout * time.Second,
		Transport: tr,
	}

	resp, err = client.Do(req)
	if err != nil {
		check.ExitError(err)
	}

	return resp
}
