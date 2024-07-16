package googlednsmanagedzone


type GoogleDnsManagedZonePrivateVisibilityConfig struct {
	// gke_clusters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_managed_zone#gke_clusters GoogleDnsManagedZone#gke_clusters}
	GkeClusters interface{} `field:"optional" json:"gkeClusters" yaml:"gkeClusters"`
	// networks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_managed_zone#networks GoogleDnsManagedZone#networks}
	Networks interface{} `field:"optional" json:"networks" yaml:"networks"`
}

