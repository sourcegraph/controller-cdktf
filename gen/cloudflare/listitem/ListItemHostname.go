package listitem


type ListItemHostname struct {
	// The FQDN to match on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/list_item#url_hostname ListItemA#url_hostname}
	UrlHostname *string `field:"required" json:"urlHostname" yaml:"urlHostname"`
}

