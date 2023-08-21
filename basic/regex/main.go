package regex

import (
	"fmt"
	"log"
	"regexp"
)

// <base id="headbase" href="/b/" />
// <base id="headbase" href="http://www.a.com/b/" />
func HTMLBase(html string) string {
	expr := `<base[^>]*?href="([^>]*)?"[^>]*?/>`
	mached :=FindStringSubmatch(html, expr)
	if len(mached) > 1{
		fmt.Println(mached[1])
		return mached[1]
	}
	return ""
}

func Title(html string) string {
	expr := `<title.*>\s*?(.*)\s*?<\/title>`
	mached := FindStringSubmatch(html, expr)
	if len(mached) > 1{
		return mached[1]
	}
	return ""
}

func FindStringSubmatch(html, expr string) []string {
	re, err := regexp.Compile(expr)
	if err != nil {
		log.Println(err)
	}
	return re.FindStringSubmatch(html)
}


func HtmlImage(html string) [][]string {
	re, err := regexp.Compile(`<img.*?src="(.*?)".*?>`)
	if err != nil {
		fmt.Println(err)
	}
	return re.FindAllStringSubmatch(html, -1)
}

func HtmlJs(html string) [][]string {
	re, err := regexp.Compile(`<script[^>]*?src="([^>]*\.js)(\?[^>"]*)?"[^>]?>.*</script>`)
	if err != nil {
		fmt.Println(err)
	}
	return re.FindAllStringSubmatch(html, -1)
}

func HtmlCss(html string) [][]string {
	// re, err := regexp.Compile(` <link[^>]*?href=\"([^>]*?)\"[^>]*?>`)
	//                               <link[^>]*?href=\"([^>]*?)\"[^>]*?>
	re, err := regexp.Compile(`<link[^>]*?href="([^>]*\.css)(\?[^>"]*)?"[^>]*?/>`)

	if err != nil {
		fmt.Println(err)
	}
	return re.FindAllStringSubmatch(html, -1)
}
