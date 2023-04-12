package googleactivedirectorypeering


type GoogleActiveDirectoryPeeringTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_active_directory_peering#create GoogleActiveDirectoryPeering#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_active_directory_peering#delete GoogleActiveDirectoryPeering#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_active_directory_peering#update GoogleActiveDirectoryPeering#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

