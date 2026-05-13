package googlecestool


type GoogleCesToolClientFunction struct {
	// The function name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_tool#name GoogleCesTool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The function description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_tool#description GoogleCesTool#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_tool#parameters GoogleCesTool#parameters}
	Parameters *GoogleCesToolClientFunctionParameters `field:"optional" json:"parameters" yaml:"parameters"`
	// response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_tool#response GoogleCesTool#response}
	Response *GoogleCesToolClientFunctionResponse `field:"optional" json:"response" yaml:"response"`
}

