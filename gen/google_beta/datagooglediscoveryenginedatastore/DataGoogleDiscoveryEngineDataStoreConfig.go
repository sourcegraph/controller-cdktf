package datagooglediscoveryenginedatastore

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleDiscoveryEngineDataStoreConfig struct {
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
	// The unique id of the data store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/data-sources/google_discovery_engine_data_store#data_store_id DataGoogleDiscoveryEngineDataStore#data_store_id}
	DataStoreId *string `field:"optional" json:"dataStoreId" yaml:"dataStoreId"`
	// The display name of the data store.
	//
	// This field must be a UTF-8 encoded
	// string with a length limit of 128 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/data-sources/google_discovery_engine_data_store#display_name DataGoogleDiscoveryEngineDataStore#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/data-sources/google_discovery_engine_data_store#id DataGoogleDiscoveryEngineDataStore#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The geographic location where the data store should reside. The value can only be one of "global", "us" and "eu".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/data-sources/google_discovery_engine_data_store#location DataGoogleDiscoveryEngineDataStore#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/data-sources/google_discovery_engine_data_store#project DataGoogleDiscoveryEngineDataStore#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

