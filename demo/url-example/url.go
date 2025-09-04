package urlexample
package url

import (
	"net/url"
	"strings"
	"path"
)

func parseResourceUrl(urlRaw, sourceUrl, baseTagHref string) (*urlpkg.URL, error) {
	urlU, err := urlpkg.QueryUnescape(urlRaw)
	if err != nil {
		return nil, err
	}
	url, err := urlpkg.Parse(urlU)
	if err != nil {
		return nil, err
	}
	sourceUrlP, err := urlpkg.Parse(sourceUrl)
	if err != nil {
		return nil, err
	}
	// /site/x.jpg
	if url.Scheme == "" && url.Host == "" && baseTagHref != "" {
		url.Path = baseTagHref + strings.TrimPrefix(url.Path, baseTagHref)
	}
	if url.Scheme == "" && url.Host == "" && baseTagHref == "" {
		// x/y
		// ./x/y
		if !strings.HasPrefix(urlRaw, "/") {
			url.Path = strings.TrimSuffix(path.Dir(sourceUrlP.Path), "/") + PathSeparator + url.Path
		}
	}
	// /x/y
	// //www.test.com/x.y
	if url.Scheme == "" {
		url.Scheme = sourceUrlP.Scheme
	}
	if url.Host == "" {
		url.Host = sourceUrlP.Host
	}
	return url, nil
}
