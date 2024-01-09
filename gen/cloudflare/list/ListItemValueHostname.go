package list


type ListItemValueHostname struct {
	// The FQDN to match on. Wildcard sub-domain matching is allowed. Eg. *.abc.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/list#url_hostname List#url_hostname}
	UrlHostname *string `field:"required" json:"urlHostname" yaml:"urlHostname"`
}

