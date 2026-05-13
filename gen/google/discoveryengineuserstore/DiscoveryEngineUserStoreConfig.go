package discoveryengineuserstore

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DiscoveryEngineUserStoreConfig struct {
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
	// The geographic location where the data store should reside. The value can only be one of "global", "us" and "eu".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_user_store#location DiscoveryEngineUserStore#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The resource name of the default license config assigned to users created in this user store.
	//
	// Format:
	// 'projects/{project}/locations/{location}/licenseConfigs/{license_config}'.
	// If 'enableLicenseAutoRegister' is true, new users will automatically
	// register under the default subscription.
	// If the default license config doesn't have remaining license seats left,
	// new users will not be assigned with license.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_user_store#default_license_config DiscoveryEngineUserStore#default_license_config}
	DefaultLicenseConfig *string `field:"optional" json:"defaultLicenseConfig" yaml:"defaultLicenseConfig"`
	// Whether to enable automatic license update for users with expired licenses in this user store.
	//
	// If enabled, users with expired licenses will
	// automatically be updated to the default subscription if there are
	// remaining license seats.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_user_store#enable_expired_license_auto_update DiscoveryEngineUserStore#enable_expired_license_auto_update}
	EnableExpiredLicenseAutoUpdate interface{} `field:"optional" json:"enableExpiredLicenseAutoUpdate" yaml:"enableExpiredLicenseAutoUpdate"`
	// Whether to enable automatic license registration for new users created in this user store.
	//
	// If enabled, new users will automatically register under
	// the default subscription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_user_store#enable_license_auto_register DiscoveryEngineUserStore#enable_license_auto_register}
	EnableLicenseAutoRegister interface{} `field:"optional" json:"enableLicenseAutoRegister" yaml:"enableLicenseAutoRegister"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_user_store#id DiscoveryEngineUserStore#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_user_store#project DiscoveryEngineUserStore#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_user_store#timeouts DiscoveryEngineUserStore#timeouts}
	Timeouts *DiscoveryEngineUserStoreTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The ID of the user store. Currently only accepts "default_user_store".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_user_store#user_store_id DiscoveryEngineUserStore#user_store_id}
	UserStoreId *string `field:"optional" json:"userStoreId" yaml:"userStoreId"`
}

