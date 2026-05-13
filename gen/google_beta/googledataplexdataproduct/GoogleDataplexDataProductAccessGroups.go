package googledataplexdataproduct


type GoogleDataplexDataProductAccessGroups struct {
	// User friendly display name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_data_product#display_name GoogleDataplexDataProduct#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Unique identifier of the access group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_data_product#group_id GoogleDataplexDataProduct#group_id}
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_data_product#id GoogleDataplexDataProduct#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// principal block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_data_product#principal GoogleDataplexDataProduct#principal}
	Principal *GoogleDataplexDataProductAccessGroupsPrincipal `field:"required" json:"principal" yaml:"principal"`
	// Description of the access group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_data_product#description GoogleDataplexDataProduct#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

