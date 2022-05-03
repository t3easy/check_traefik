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

type States struct {
	Total    int `json:"total"`
	Warnings int `json:"warnings"`
	Errors   int `json:"errors"`
}

type Service struct {
	Routers     States `json:"routers"`
	Services    States `json:"services"`
	Middlewares States `json:"middlewares"`
}

type Overview struct {
	Http Service `json:"http"`
	Tcp  Service `json:"tcp"`
	UDP  struct {
		Routers  States `json:"routers"`
		Services States `json:"services"`
	} `json:"udp"`
	Features struct {
		Tracing   string `json:"tracing"`
		Metrics   string `json:"metrics"`
		AccessLog bool   `json:"accessLog"`
	} `json:"features"`
	Providers []string `json:"providers"`
}
