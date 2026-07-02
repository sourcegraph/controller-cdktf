package dnsmanagedzone

type DnsManagedZonePrivateVisibilityConfig struct {
	// gke_clusters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/dns_managed_zone#gke_clusters DnsManagedZone#gke_clusters}
	GkeClusters any `field:"optional" json:"gkeClusters" yaml:"gkeClusters"`
	// networks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/dns_managed_zone#networks DnsManagedZone#networks}
	Networks any `field:"optional" json:"networks" yaml:"networks"`
}
