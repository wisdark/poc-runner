/*
 * poc-runner project
 * Copyright (C) 2024 4ra1n
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package reverse

import (
	"fmt"
	"github.com/4ra1n/poc-runner/client"
	"github.com/4ra1n/poc-runner/rawhttp"
	"testing"
)

func TestDNSLOG(t *testing.T) {
	c, err := client.NewHttpClient(rawhttp.DefaultNoProxy, rawhttp.DefaultTimeout, false)
	if err != nil {
		panic(err)
	}
	rev, err := NewReverse(c)
	if err != nil {
		panic(err)
	}
	fmt.Println(rev)
	rev.Wait(20)
}
