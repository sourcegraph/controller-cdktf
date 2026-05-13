package googleartifactregistryrule


type GoogleArtifactRegistryRuleCondition struct {
	// Textual representation of an expression in Common Expression Language syntax.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_artifact_registry_rule#expression GoogleArtifactRegistryRule#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Optional.
	//
	// Description of the expression. This is a longer text which
	// describes the expression, e.g. when hovered over it in a UI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_artifact_registry_rule#description GoogleArtifactRegistryRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Optional.
	//
	// String indicating the location of the expression for error
	// reporting, e.g. a file name and a position in the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_artifact_registry_rule#location GoogleArtifactRegistryRule#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Optional. Title for the expression, i.e. a short string describing its purpose.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_artifact_registry_rule#title GoogleArtifactRegistryRule#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

