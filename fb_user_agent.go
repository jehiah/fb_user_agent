package fb_user_agent

import (
	"errors"
	"strings"
	"fmt"
)

// Based on http://mpulp.mobi/2012/01/funky-user-agent-on-facebook-iphone-app/
type FBUserAgent struct {
	ApplicationName    string // FBAN
	ApplicationVersion string // FBAV
	BuildVersion       string // FBBV
	Device             string // FBDV
	FBMD               string
	FBSN               string
	FBSV               string // OS Version?
	FBSS               string
	Carrier            string // FBCR
	FBID               string
	Language           string // FBLC
	FBSF               string
	FBOP               string
	FBCA               string
	FBDM               string // {density=3.0,width=1080,height=1920}
	FBPN               string // com.facebook.katana
}

func ParseFBUserAgent(ua string) (FBUserAgent, error) {
	f := FBUserAgent{}
	if !strings.Contains(ua, "[FBAN/") {
		return f, errors.New("Not a valid Facebook User-Agent")
	}
	u := ua[strings.Index(ua, "[FBAN/"):]
	u = strings.Trim(u, "[]")
	chunks := strings.Split(u, ";")
	if len(chunks) < 3 {
		return f, errors.New("Error parsing User-Agent")
	}

	for _, c := range chunks {
		if c == "" {
			continue
		}
		data := strings.SplitN(c, "/", 2)
		if len(data) != 2 {
			if c == "FBOP" {
				f.FBOP = c
				continue
			}
			return f, fmt.Errorf("found %d expected 2 in %#v", len(data), data)
		}
		key := strings.Trim(data[0], " ")
		switch key {
		case "":
			continue
		case "FBAN":
			f.ApplicationName = data[1]
		case "FBAV":
			f.ApplicationVersion = data[1]
		case "FBBV":
			f.BuildVersion = data[1]
		case "FBDV":
			f.Device = data[1]
		case "FBMD":
			f.FBMD = data[1]
		case "FBSN":
			f.FBSN = data[1]
		case "FBSV":
			f.FBSV = data[1]
		case "FBSS":
			f.FBSS = data[1]
		case "FBCR":
			f.Carrier = data[1]
		case "FBID":
			f.FBID = data[1]
		case "FBLC":
			f.Language = data[1]
		case "FBSF":
			f.FBSF = data[1]
		case "FBOP":
			f.FBOP = data[1]
		case "FBCA":
			f.FBCA = data[1]
		case "FBDM":
			f.FBDM = data[1]
		case "FBPN":
			f.FBPN = data[1]
		default:
			return f, fmt.Errorf("unknown %v:%q in %q", data[0], data[1], ua)
		}
	}
	return f, nil
}
