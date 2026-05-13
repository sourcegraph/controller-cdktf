package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigOtherCloudStartingLocationAwsLocation struct {
	// The AWS account ID that this discovery config applies to.
	//
	// Within an organization, you can find the AWS account ID inside an AWS account ARN. Example: arn:<partition>:organizations::<management-account-id>:account/<organization-id>/<account-id>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/data_loss_prevention_discovery_config#account_id DataLossPreventionDiscoveryConfig#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// All AWS assets stored in Asset Inventory that didn't match other AWS discovery configs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/data_loss_prevention_discovery_config#all_asset_inventory_assets DataLossPreventionDiscoveryConfig#all_asset_inventory_assets}
	AllAssetInventoryAssets interface{} `field:"optional" json:"allAssetInventoryAssets" yaml:"allAssetInventoryAssets"`
}

