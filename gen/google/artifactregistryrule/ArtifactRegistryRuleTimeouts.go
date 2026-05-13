package artifactregistryrule


type ArtifactRegistryRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/artifact_registry_rule#create ArtifactRegistryRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/artifact_registry_rule#delete ArtifactRegistryRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/artifact_registry_rule#update ArtifactRegistryRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

