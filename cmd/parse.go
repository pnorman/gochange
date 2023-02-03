/*
Copyright © 2022 Paul Norman <osm@paulnorman.ca>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package cmd

import (
	"encoding/xml"
	"fmt"
	"io"
	"time"

	"github.com/shopspring/decimal"
)

type Changeset struct {
	Id         int             `xml:"id,attr"`
	CreatedAt  time.Time       `xml:"created_at,attr"`
	ClosedAt   time.Time       `xml:"closed_at,attr"`
	Open       bool            `xml:"open,attr"`
	User       string          `xml:"user,attr"`
	Uid        int             `xml:"uid,attr"`
	MinLat     decimal.Decimal `xml:"min_lat,attr"`
	MinLon     decimal.Decimal `xml:"min_lon,attr"`
	MaxLat     decimal.Decimal `xml:"max_lat,attr"`
	MaxLon     decimal.Decimal `xml:"max_lon,attr"`
	NumChanges int             `xml:"num_changes"`
	// Skip comments_count
}

func parseOsm(file io.Reader) {
	fmt.Println("Parsing XML")

	decoder := xml.NewDecoder(file)

	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}

		// Based roughly on https://blog.singleton.io/posts/2012-06-19-parsing-huge-xml-files-with-go/
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "osm" {
				fmt.Println("<osm/> element not yet handled")
			} else if se.Name.Local == "changeset" {
				var cs Changeset
				decoder.DecodeElement(&cs, &se)
				fmt.Printf("Read changeset %v\n", cs)
			}
		}
	}
}
