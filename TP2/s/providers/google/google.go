package google

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("google", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for Google.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.google.com/search?q=%s", url.QueryEscape(q))
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	return []string{"search"}
}
