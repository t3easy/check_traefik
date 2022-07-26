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
	"net"
	"net/http"
	"testing"
)

func TestNewRequest(t *testing.T) {
	var (
		hostname string = "localhost"
		url      string = "http://127.0.0.1:8080/ping"
		req      *http.Request
	)

	req = NewRequest(http.MethodHead, net.ParseIP("127.0.0.1"), hostname, false, 8080, "ping", "user", "password")
	if req.Host != hostname {
		t.Fatalf(`Want: %v --> return value: %v`, hostname, req.Host)
	}
	if req.URL.String() != url {
		t.Fatalf(`Want: %v --> return value: %v`, url, req.RequestURI)
	}
}
