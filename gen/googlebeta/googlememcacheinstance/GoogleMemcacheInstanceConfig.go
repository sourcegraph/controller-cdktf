package googlememcacheinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleMemcacheInstanceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// The resource name of the instance.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_memcache_instance#name GoogleMemcacheInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// node_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_memcache_instance#node_config GoogleMemcacheInstance#node_config}
	NodeConfig *GoogleMemcacheInstanceNodeConfig `field:"required" json:"nodeConfig" yaml:"nodeConfig"`
	// Number of nodes in the memcache instance.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_memcache_instance#node_count GoogleMemcacheInstance#node_count}
	NodeCount *float64 `field:"required" json:"nodeCount" yaml:"nodeCount"`
	// The full name of the GCE network to connect the instance to.  If not provided, 'default' will be used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_memcache_instance#authorized_network GoogleMemcacheInstance#authorized_network}
	AuthorizedNetwork *string `field:"optional" json:"authorizedNetwork" yaml:"authorizedNetwork"`
	// A user-visible name for the instance.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_memcache_instance#display_name GoogleMemcacheInstance#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_memcache_instance#id GoogleMemcacheInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Resource labels to represent user-provided metadata.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_memcache_instance#labels GoogleMemcacheInstance#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// maintenance_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_memcache_instance#maintenance_policy GoogleMemcacheInstance#maintenance_policy}
	MaintenancePolicy *GoogleMemcacheInstanceMaintenancePolicy `field:"optional" json:"maintenancePolicy" yaml:"maintenancePolicy"`
	// memcache_parameters block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_memcache_instance#memcache_parameters GoogleMemcacheInstance#memcache_parameters}
	MemcacheParameters *GoogleMemcacheInstanceMemcacheParameters `field:"optional" json:"memcacheParameters" yaml:"memcacheParameters"`
	// The major version of Memcached software.
	//
	// If not provided, latest supported version will be used.
	// Currently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically
	// determined by our system based on the latest supported minor version. Default value: "MEMCACHE_1_5" Possible values: ["MEMCACHE_1_5"]
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_memcache_instance#memcache_version GoogleMemcacheInstance#memcache_version}
	MemcacheVersion *string `field:"optional" json:"memcacheVersion" yaml:"memcacheVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_memcache_instance#project GoogleMemcacheInstance#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region of the Memcache instance. If it is not provided, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_memcache_instance#region GoogleMemcacheInstance#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_memcache_instance#timeouts GoogleMemcacheInstance#timeouts}
	Timeouts *GoogleMemcacheInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Zones where memcache nodes should be provisioned.  If not provided, all zones will be used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_memcache_instance#zones GoogleMemcacheInstance#zones}
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

