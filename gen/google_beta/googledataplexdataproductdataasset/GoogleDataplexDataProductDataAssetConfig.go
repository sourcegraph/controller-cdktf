package googledataplexdataproductdataasset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDataplexDataProductDataAssetConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The ID of the data asset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_data_product_data_asset#data_asset_id GoogleDataplexDataProductDataAsset#data_asset_id}
	DataAssetId *string `field:"required" json:"dataAssetId" yaml:"dataAssetId"`
	// The ID of the parent data product.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_data_product_data_asset#data_product_id GoogleDataplexDataProductDataAsset#data_product_id}
	DataProductId *string `field:"required" json:"dataProductId" yaml:"dataProductId"`
	// The location for the data asset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_data_product_data_asset#location GoogleDataplexDataProductDataAsset#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Full resource name of the cloud resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_data_product_data_asset#resource GoogleDataplexDataProductDataAsset#resource}
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// access_group_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_data_product_data_asset#access_group_configs GoogleDataplexDataProductDataAsset#access_group_configs}
	AccessGroupConfigs interface{} `field:"optional" json:"accessGroupConfigs" yaml:"accessGroupConfigs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_data_product_data_asset#id GoogleDataplexDataProductDataAsset#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// User-defined labels.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_data_product_data_asset#labels GoogleDataplexDataProductDataAsset#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_data_product_data_asset#project GoogleDataplexDataProductDataAsset#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_data_product_data_asset#timeouts GoogleDataplexDataProductDataAsset#timeouts}
	Timeouts *GoogleDataplexDataProductDataAssetTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

