package sitemap

import (
	"encoding/xml"
)

type Urlset struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	Urls    []*Url   `xml:"url"`
}

type Url struct {
	XMLName xml.Name `xml:"url"`
	Loc     string   `xml:"loc"`
}

func SiteMap(links []*Url) ([]byte, error) {
	u := &Urlset{Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9"}
	u.Urls = links

	out, err := xml.MarshalIndent(u, "", "  ")
	if err != nil {
		return nil, err
	}

	headerByte := []byte(xml.Header)

	return append(headerByte, out...), nil
}
