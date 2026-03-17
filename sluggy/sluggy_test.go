package sluggy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlug(t *testing.T) {
	cases := []struct {
		testName, input, want string
	}{
		{"Basic case", "Convert to URL", "convert-to-url"},
		{"Basic case with numbers", "1-2-3 4 5 6", "1-2-3-4-5-6"},
		{"Spaces and punctuations case", "Dude,  this  is  new  line  for  you!", "dude-this-is-new-line-for-you"},
		{"Repeating separators case", "Kubernetes,, Helm;; EKS-- for||you", "kubernetes-helm-eks-for-you"},
		{"Unicode symbols case", "Разработка и поддержка ПО", "разработка-и-поддержка-по"},
		{"Empty string case", "", ""},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			got := Slug(c.input)
			assert.Equal(t, c.want, got)
		})
	}
}
