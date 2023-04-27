package numeronym

import (
	"testing"
)

func TestNumeronym(t *testing.T) {
	tests := []struct {
		word string
		want string
	}{
		{
			word: "number1", // All chars must be an alphabet.
			want: "number1",
		},
		{
			word: "age", // Too short.
			want: "age",
		},
		{
			word: "word",
			want: "w2d",
		},
		{
			word: "shorten",
			want: "s5n",
		},
		{
			word: "Kubernetes",
			want: "K8s",
		},
		{
			word: "accessibility",
			want: "a11y",
		},
		{
			word: "internationalization",
			want: "i18n",
		},
	}
	for _, test := range tests {
		name := test.word + ":" + test.want
		t.Run(name, func(t *testing.T) {
			actual := Numeronize(test.word)
			if actual != test.want {
				t.Errorf("Not match. actual:%s, want:%s", actual, test.want)
			}
		})
	}
}
