package fb_user_agent

import (
	"testing"
)

func TestParseFBUserAgent(t *testing.T) {

	type testCase struct {
		ua      string
		name    string
		version string
	}

	tests := []testCase{
		{
			ua:      "Mozilla/5.0 (iPhone; CPU iPhone OS 9_2 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Mobile/13C75 [FBAN/FBIOS;FBAV/53.0.0.36.140;FBBV/27547874;FBRV/0;FBDV/iPhone5,2;FBMD/iPhone;FBSN/iPhone OS;FBSV/9.2;FBSS/2;FBCR/VodafoneAL;FBID/phone;FBLC/en_GB;FBOP/5]",
			name:    "FBIOS",
			version: "53.0.0.36.140",
		},
	}

	for _, tc := range tests {
		f, err := ParseFBUserAgent(tc.ua)
		if tc.name == "" {
			if err == nil {
				t.Fatalf("expected err parsing %q", tc.ua)
			} else {
				continue
			}
		}
		if err != nil {
			t.Fatalf("got unexpected err %s parsing %q", err, tc.ua)
		}
		if tc.name != f.ApplicationName {
			t.Errorf("got name %q expected %q for %q", f.ApplicationName, tc.name, tc.ua)
		}
		if tc.version != f.ApplicationVersion {
			t.Errorf("got version %q expected %q for %q", f.ApplicationVersion, tc.version, tc.ua)
		}
	}

}
