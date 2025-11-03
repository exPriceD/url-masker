package urlmasker

import (
	"testing"
)

func TestMaskURL_Table(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "example_from_task",
			in:   "Hello, its my page: http://localhost123.com See you",
			out:  "Hello, its my page: http://**************** See you",
		},
		{
			name: "multiple_links_and_whitespace",
			in:   "Go http://a b http://xyz\tend",
			out:  "Go http://* b http://***\tend",
		},
		{
			name: "no_http_links",
			in:   "no links here, just text",
			out:  "no links here, just text",
		},
		{
			name: "https_not_masked",
			in:   "secure https://site.com stays",
			out:  "secure https://site.com stays",
		},
		{
			name: "uppercase_not_masked",
			in:   "check HTTP://UPPERCASE should stay",
			out:  "check HTTP://UPPERCASE should stay",
		},
		{
			name: "link_at_end",
			in:   "end http://abc",
			out:  "end http://***",
		},
		{
			name: "mask_until_space_even_punctuation",
			in:   "see http://host, next",
			out:  "see http://***** next",
		},
		{
			name: "empty",
			in:   "",
			out:  "",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := MaskURL(tc.in)
			if got != tc.out {
				t.Fatalf("MaskURL(%q)\n got:  %q\n want: %q", tc.in, got, tc.out)
			}
		})
	}
}
