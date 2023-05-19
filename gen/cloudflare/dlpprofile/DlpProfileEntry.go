package dlpprofile


type DlpProfileEntry struct {
	// Name of the entry to deploy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/dlp_profile#name DlpProfile#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether the entry is active. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/dlp_profile#enabled DlpProfile#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Unique entry identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/dlp_profile#id DlpProfile#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/dlp_profile#pattern DlpProfile#pattern}
	Pattern *DlpProfileEntryPattern `field:"optional" json:"pattern" yaml:"pattern"`
}

