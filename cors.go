package nbdata

import (
	"strings"

	"github.com/rs/cors"
)

var (
	CORS = []string{
		"https://vis.nucleome.org",
		"http://vis.nucleome.org",
		"https://viz.nucleome.org",
		"http://viz.nucleome.org",
		"https://browser.nucleome.org",
		"http://browser.nucleome.org",
		"https://v.nucleome.org",
		"http://v.nucleome.org",
		"https://doc.nucleome.org",
		"http://doc.nucleome.org",
		"http://www.nucleome.org",
		"https://www.nucleome.org",
		"http://nucleome.org",
		"https://nucleome.org",
		"https://nucleome.github.io",
		"https://4dn.github.io",
		"https://genome.compbio.cs.cmu.edu",
		"http://genome.compbio.cs.cmu.edu:8080",
		"chrome-extension://djcdicpaejhpgncicoglfckiappkoeof",
		/* for development */
		"http://x7.andrew.cmu.edu:8080",
		"https://nimezhu.github.io",
		"https://dev.nucleome.org",
		"http://dev.nucleome.org",
		"chrome-extension://gedcoafficobgcagmjpplpnmempkcpfp",
		"https://youdata.studio",
		"https://youdata.github.io",
		"https://vomics.github.io"
	}
)

func GetCors(customCors string) cors.Options {
	if customCors != "" {
		otherCors := strings.Split(customCors, ";")
		for _, s := range otherCors {
			CORS = append(CORS, s)
		}

	}
	corsOptions := cors.Options{
		AllowedOrigins:   CORS,
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization"},
	}
	return corsOptions

}
