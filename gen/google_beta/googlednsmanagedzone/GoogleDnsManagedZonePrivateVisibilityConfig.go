package googlednsmanagedzone


type GoogleDnsManagedZonePrivateVisibilityConfig struct {
	// networks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_dns_managed_zone#networks GoogleDnsManagedZone#networks}
	Networks interface{} `field:"required" json:"networks" yaml:"networks"`
	// gke_clusters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_dns_managed_zone#gke_clusters GoogleDnsManagedZone#gke_clusters}
	GkeClusters interface{} `field:"optional" json:"gkeClusters" yaml:"gkeClusters"`
}

