package googlecesapp


type GoogleCesAppVariableDeclarations struct {
	// The description of the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#description GoogleCesApp#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name of the variable.
	//
	// The name must start with a letter or underscore
	// and contain only letters, numbers, or underscores.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#name GoogleCesApp#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#schema GoogleCesApp#schema}
	Schema *GoogleCesAppVariableDeclarationsSchema `field:"required" json:"schema" yaml:"schema"`
}

