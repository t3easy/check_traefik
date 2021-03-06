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
	"testing"

	"github.com/NETWAYS/go-check"
)

func TestCheckResponseOK(t *testing.T) {
	var (
		resp *http.Response
		rc   int
	)

	resp = &http.Response{
		StatusCode: 200,
	}
	rc = checkHealthResponse(resp)
	if rc != check.OK {
		t.Fatalf(`Want: %v --> return value: %v`, check.OK, rc)
	}
}

func TestCheckResponseCritical(t *testing.T) {
	var (
		resp *http.Response
		rc   int
	)

	resp = &http.Response{
		StatusCode: 503,
	}
	rc = checkHealthResponse(resp)
	if rc != check.Critical {
		t.Fatalf(`Want: %v --> return value: %v`, check.OK, rc)
	}
}
