package resourcequota


type ResourceQuotaTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/resource_quota#create ResourceQuota#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/resource_quota#update ResourceQuota#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

