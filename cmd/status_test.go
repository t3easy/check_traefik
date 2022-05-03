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
	"strings"
	"testing"
)

const (
	apiOverviewReturn = `{"http":{"routers":{"total":7,"warnings":0,"errors":0},"services":{"total":9,"warnings":0,"errors":0},"middlewares":{"total":2,"warnings":0,"errors":0}},"tcp":{"routers":{"total":0,"warnings":0,"errors":0},"services":{"total":0,"warnings":0,"errors":0},"middlewares":{"total":0,"warnings":0,"errors":0}},"udp":{"routers":{"total":0,"warnings":0,"errors":0},"services":{"total":0,"warnings":0,"errors":0}},"features":{"tracing":"","metrics":"","accessLog":false},"providers":["Docker","File"]}`
)

func TestOverviewStruct(t *testing.T) {
	var overview Overview
	if err := json.Unmarshal([]byte(apiOverviewReturn), &overview); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if overview.Http.Routers.Total != 7 {
		t.Fatalf("Total = %d, want 7", overview.Http.Routers.Total)
	}
	if !strings.Contains(strings.Join(overview.Providers, ","), "Docker") {
		t.Fatalf(`"%s" not found in "%s"`, "Docker", strings.Join(overview.Providers, ","))
	}
}
