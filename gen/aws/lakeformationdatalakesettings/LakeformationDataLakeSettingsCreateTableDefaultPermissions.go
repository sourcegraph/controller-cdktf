package lakeformationdatalakesettings


type LakeformationDataLakeSettingsCreateTableDefaultPermissions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/lakeformation_data_lake_settings#permissions LakeformationDataLakeSettings#permissions}.
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/lakeformation_data_lake_settings#principal LakeformationDataLakeSettings#principal}.
	Principal *string `field:"optional" json:"principal" yaml:"principal"`
}

