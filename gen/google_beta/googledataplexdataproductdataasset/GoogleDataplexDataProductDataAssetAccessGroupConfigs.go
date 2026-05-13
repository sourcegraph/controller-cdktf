package googledataplexdataproductdataasset


type GoogleDataplexDataProductDataAssetAccessGroupConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_data_product_data_asset#access_group GoogleDataplexDataProductDataAsset#access_group}.
	AccessGroup *string `field:"required" json:"accessGroup" yaml:"accessGroup"`
	// IAM roles granted on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_data_product_data_asset#iam_roles GoogleDataplexDataProductDataAsset#iam_roles}
	IamRoles *[]*string `field:"optional" json:"iamRoles" yaml:"iamRoles"`
}

