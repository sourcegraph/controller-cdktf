package cesapp


type CesAppVariableDeclarations struct {
	// The description of the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#description CesApp#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name of the variable.
	//
	// The name must start with a letter or underscore
	// and contain only letters, numbers, or underscores.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#name CesApp#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#schema CesApp#schema}
	Schema *CesAppVariableDeclarationsSchema `field:"required" json:"schema" yaml:"schema"`
}

