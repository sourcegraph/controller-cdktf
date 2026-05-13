package googlelustreinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleLustreInstanceConfig struct {
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
	// The storage capacity of the instance in gibibytes (GiB).
	//
	// Allowed values
	// are from '9000' to '7632000', depending on the 'perUnitStorageThroughput'.
	// See [Performance tiers and maximum storage
	// capacities](https://cloud.google.com/managed-lustre/docs/create-instance#performance-tiers)
	// for specific minimums, maximums, and step sizes for each performance tier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_lustre_instance#capacity_gib GoogleLustreInstance#capacity_gib}
	CapacityGib *string `field:"required" json:"capacityGib" yaml:"capacityGib"`
	// The filesystem name for this instance.
	//
	// This name is used by client-side
	// tools, including when mounting the instance. Must be eight characters or
	// less and can only contain letters and numbers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_lustre_instance#filesystem GoogleLustreInstance#filesystem}
	Filesystem *string `field:"required" json:"filesystem" yaml:"filesystem"`
	// The name of the Managed Lustre instance.
	//
	// * Must contain only lowercase letters, numbers, and hyphens.
	// * Must start with a letter.
	// * Must be between 1-63 characters.
	// * Must end with a number or a letter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_lustre_instance#instance_id GoogleLustreInstance#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_lustre_instance#location GoogleLustreInstance#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The full name of the VPC network to which the instance is connected. Must be in the format 'projects/{project_id}/global/networks/{network_name}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_lustre_instance#network GoogleLustreInstance#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// access_rules_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_lustre_instance#access_rules_options GoogleLustreInstance#access_rules_options}
	AccessRulesOptions *GoogleLustreInstanceAccessRulesOptions `field:"optional" json:"accessRulesOptions" yaml:"accessRulesOptions"`
	// A user-readable description of the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_lustre_instance#description GoogleLustreInstance#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// dynamic_tier_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_lustre_instance#dynamic_tier_options GoogleLustreInstance#dynamic_tier_options}
	DynamicTierOptions *GoogleLustreInstanceDynamicTierOptions `field:"optional" json:"dynamicTierOptions" yaml:"dynamicTierOptions"`
	// Indicates whether you want to enable support for GKE clients. By default, GKE clients are not supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_lustre_instance#gke_support_enabled GoogleLustreInstance#gke_support_enabled}
	GkeSupportEnabled interface{} `field:"optional" json:"gkeSupportEnabled" yaml:"gkeSupportEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_lustre_instance#id GoogleLustreInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The Cloud KMS key name to use for data encryption.
	//
	// If not set, the instance will use Google-managed encryption keys.
	// If set, the instance will use customer-managed encryption keys.
	// The key must be in the same region as the instance.
	// The key format is:
	// projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{key}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_lustre_instance#kms_key GoogleLustreInstance#kms_key}
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Labels as key value pairs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_lustre_instance#labels GoogleLustreInstance#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// maintenance_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_lustre_instance#maintenance_policy GoogleLustreInstance#maintenance_policy}
	MaintenancePolicy *GoogleLustreInstanceMaintenancePolicy `field:"optional" json:"maintenancePolicy" yaml:"maintenancePolicy"`
	// The throughput of the instance in MBps per TiB.
	//
	// Valid values are 125, 250,
	// 500, 1000.
	// See [Performance tiers and maximum storage
	// capacities](https://cloud.google.com/managed-lustre/docs/create-instance#performance-tiers)
	// for more information.
	//
	// If the instance is using the Dynamic tier, this field must not be set or
	// must be set to zero.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_lustre_instance#per_unit_storage_throughput GoogleLustreInstance#per_unit_storage_throughput}
	PerUnitStorageThroughput *string `field:"optional" json:"perUnitStorageThroughput" yaml:"perUnitStorageThroughput"`
	// The placement policy name for the instance in the format of projects/{project}/locations/{location}/resourcePolicies/{resource_policy}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_lustre_instance#placement_policy GoogleLustreInstance#placement_policy}
	PlacementPolicy *string `field:"optional" json:"placementPolicy" yaml:"placementPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_lustre_instance#project GoogleLustreInstance#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_lustre_instance#timeouts GoogleLustreInstance#timeouts}
	Timeouts *GoogleLustreInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

