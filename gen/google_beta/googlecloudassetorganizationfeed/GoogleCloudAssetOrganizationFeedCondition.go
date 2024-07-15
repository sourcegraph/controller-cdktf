package googlecloudassetorganizationfeed


type GoogleCloudAssetOrganizationFeedCondition struct {
	// Textual representation of an expression in Common Expression Language syntax.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_asset_organization_feed#expression GoogleCloudAssetOrganizationFeed#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Description of the expression.
	//
	// This is a longer text which describes the expression,
	// e.g. when hovered over it in a UI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_asset_organization_feed#description GoogleCloudAssetOrganizationFeed#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_asset_organization_feed#location GoogleCloudAssetOrganizationFeed#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_asset_organization_feed#title GoogleCloudAssetOrganizationFeed#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

