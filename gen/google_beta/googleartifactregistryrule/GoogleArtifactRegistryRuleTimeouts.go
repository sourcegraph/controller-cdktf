package googleartifactregistryrule


type GoogleArtifactRegistryRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_artifact_registry_rule#create GoogleArtifactRegistryRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_artifact_registry_rule#delete GoogleArtifactRegistryRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_artifact_registry_rule#update GoogleArtifactRegistryRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

