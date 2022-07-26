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
	"testing"
)

const (
	apiVersionReturn = `{"Version":"2.6.3","Codename":"rocamadour","startDate":"2022-04-21T15:33:01.344178802+02:00","pilotEnabled":false}`
)

func TestVersionStruct(t *testing.T) {
	var (
		version Version
		err     error
	)

	if err = json.Unmarshal([]byte(apiVersionReturn), &version); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if version.Version != "2.6.3" {
		t.Fatalf("Version = %s, want 2.6.3", version.Version)
	}
}

func TestNormalizeVersion(t *testing.T) {
	var (
		versionWithV    string = "v2.6.4"
		versionWithoutV string = "2.6.5"
	)
	if normalizeVersion(versionWithV) != versionWithV {
		t.Fatalf("Error normalizing version %s.", versionWithV)
	}
	if normalizeVersion(versionWithoutV) != "v"+versionWithoutV {
		t.Fatalf("Error normalizing version %s.", versionWithoutV)
	}
}
