package datacloudflarewafpackages


type DataCloudflareWafPackagesFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/d/waf_packages#action_mode DataCloudflareWafPackages#action_mode}.
	ActionMode *string `field:"optional" json:"actionMode" yaml:"actionMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/d/waf_packages#detection_mode DataCloudflareWafPackages#detection_mode}.
	DetectionMode *string `field:"optional" json:"detectionMode" yaml:"detectionMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/d/waf_packages#name DataCloudflareWafPackages#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/d/waf_packages#sensitivity DataCloudflareWafPackages#sensitivity}.
	Sensitivity *string `field:"optional" json:"sensitivity" yaml:"sensitivity"`
}

