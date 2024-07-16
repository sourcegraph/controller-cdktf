package googlegkeonprembaremetalcluster


type GoogleGkeonpremBareMetalClusterProxy struct {
	// Specifies the address of your proxy server.
	//
	// Examples: http://domain
	// WARNING: Do not provide credentials in the format
	// http://(username:password@)domain these will be rejected by the server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#uri GoogleGkeonpremBareMetalCluster#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// A list of IPs, hostnames, and domains that should skip the proxy. Examples: ["127.0.0.1", "example.com", ".corp", "localhost"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#no_proxy GoogleGkeonpremBareMetalCluster#no_proxy}
	NoProxy *[]*string `field:"optional" json:"noProxy" yaml:"noProxy"`
}

