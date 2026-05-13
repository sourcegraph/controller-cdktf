package googlenetappstoragepool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetappStoragePoolConfig struct {
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
	// Capacity of the storage pool (in GiB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_storage_pool#capacity_gib GoogleNetappStoragePool#capacity_gib}
	CapacityGib *string `field:"required" json:"capacityGib" yaml:"capacityGib"`
	// Name of the location. For zonal Flex pools specify a zone name, in all other cases a region name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_storage_pool#location GoogleNetappStoragePool#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The resource name of the storage pool. Needs to be unique per location/region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_storage_pool#name GoogleNetappStoragePool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// VPC network name with format: 'projects/{{project}}/global/networks/{{network}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_storage_pool#network GoogleNetappStoragePool#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Service level of the storage pool. Possible values: ["PREMIUM", "EXTREME", "STANDARD", "FLEX"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_storage_pool#service_level GoogleNetappStoragePool#service_level}
	ServiceLevel *string `field:"required" json:"serviceLevel" yaml:"serviceLevel"`
	// Specifies the Active Directory policy to be used.
	//
	// Format: 'projects/{{project}}/locations/{{location}}/activeDirectories/{{name}}'.
	// The policy needs to be in the same location as the storage pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_storage_pool#active_directory GoogleNetappStoragePool#active_directory}
	ActiveDirectory *string `field:"optional" json:"activeDirectory" yaml:"activeDirectory"`
	// Optional.
	//
	// True if the storage pool supports Auto Tiering enabled volumes. Default is false.
	// Auto-tiering can be enabled after storage pool creation but it can't be disabled once enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_storage_pool#allow_auto_tiering GoogleNetappStoragePool#allow_auto_tiering}
	AllowAutoTiering interface{} `field:"optional" json:"allowAutoTiering" yaml:"allowAutoTiering"`
	// Optional. True if using Independent Scaling of capacity and performance (Hyperdisk). Default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_storage_pool#custom_performance_enabled GoogleNetappStoragePool#custom_performance_enabled}
	CustomPerformanceEnabled interface{} `field:"optional" json:"customPerformanceEnabled" yaml:"customPerformanceEnabled"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_storage_pool#description GoogleNetappStoragePool#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Flag indicating that the hot-tier threshold will be auto-increased by 10% of the hot-tier when it hits 100%.
	//
	// Default is true.
	// The increment will kick in only if the new size after increment is still less than or equal to storage pool size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_storage_pool#enable_hot_tier_auto_resize GoogleNetappStoragePool#enable_hot_tier_auto_resize}
	EnableHotTierAutoResize interface{} `field:"optional" json:"enableHotTierAutoResize" yaml:"enableHotTierAutoResize"`
	// Total hot tier capacity for the Storage Pool.
	//
	// It is applicable only to Flex service level.
	// It should be less than the minimum storage pool size and cannot be more than the current storage pool size. It cannot be decreased once set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_storage_pool#hot_tier_size_gib GoogleNetappStoragePool#hot_tier_size_gib}
	HotTierSizeGib *string `field:"optional" json:"hotTierSizeGib" yaml:"hotTierSizeGib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_storage_pool#id GoogleNetappStoragePool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the CMEK policy to be used for volume encryption.
	//
	// Format: 'projects/{{project}}/locations/{{location}}/kmsConfigs/{{name}}'.
	// The policy needs to be in the same location as the storage pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_storage_pool#kms_config GoogleNetappStoragePool#kms_config}
	KmsConfig *string `field:"optional" json:"kmsConfig" yaml:"kmsConfig"`
	// Labels as key value pairs. Example: '{ "owner": "Bob", "department": "finance", "purpose": "testing" }'.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_storage_pool#labels GoogleNetappStoragePool#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// When enabled, the volumes uses Active Directory as LDAP name service for UID/GID lookups.
	//
	// Required to enable extended group support for NFSv3,
	// using security identifiers for NFSv4.1 or principal names for kerberized NFSv4.1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_storage_pool#ldap_enabled GoogleNetappStoragePool#ldap_enabled}
	LdapEnabled interface{} `field:"optional" json:"ldapEnabled" yaml:"ldapEnabled"`
	// Mode of the storage pool.
	//
	// The operational mode of the storage pool. ONTAP mode enables operations
	// via ONTAP Mode APIs, while DEFAULT mode enables operations via NetApp Volumes APIs.
	// If not specified during creation, the mode defaults to DEFAULT. Possible values: ["MODE_UNSPECIFIED", "DEFAULT", "ONTAP"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_storage_pool#mode GoogleNetappStoragePool#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_storage_pool#project GoogleNetappStoragePool#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// QoS (Quality of Service) type of the storage pool. Possible values are: AUTO, MANUAL. Possible values: ["QOS_TYPE_UNSPECIFIED", "AUTO", "MANUAL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_storage_pool#qos_type GoogleNetappStoragePool#qos_type}
	QosType *string `field:"optional" json:"qosType" yaml:"qosType"`
	// Specifies the replica zone for regional Flex pools. 'zone' and 'replica_zone' values can be swapped to initiate a [zone switch](https://cloud.google.com/netapp/volumes/docs/configure-and-use/storage-pools/edit-or-delete-storage-pool#switch_active_and_replica_zones).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_storage_pool#replica_zone GoogleNetappStoragePool#replica_zone}
	ReplicaZone *string `field:"optional" json:"replicaZone" yaml:"replicaZone"`
	// The effective scale tier of the storage pool.
	//
	// If 'scale_tier' is not
	// specified during creation, this defaults to 'SCALE_TIER_STANDARD'. Possible values: ["SCALE_TIER_UNSPECIFIED", "SCALE_TIER_STANDARD", "SCALE_TIER_ENTERPRISE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_storage_pool#scale_tier GoogleNetappStoragePool#scale_tier}
	ScaleTier *string `field:"optional" json:"scaleTier" yaml:"scaleTier"`
	// The scale type of the storage pool. Defaults to 'SCALE_TYPE_DEFAULT' if not specified. Possible values: ["SCALE_TYPE_UNSPECIFIED", "SCALE_TYPE_DEFAULT", "SCALE_TYPE_SCALEOUT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_storage_pool#scale_type GoogleNetappStoragePool#scale_type}
	ScaleType *string `field:"optional" json:"scaleType" yaml:"scaleType"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_storage_pool#timeouts GoogleNetappStoragePool#timeouts}
	Timeouts *GoogleNetappStoragePoolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Optional. Custom Performance Total IOPS of the pool If not provided, it will be calculated based on the totalThroughputMibps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_storage_pool#total_iops GoogleNetappStoragePool#total_iops}
	TotalIops *string `field:"optional" json:"totalIops" yaml:"totalIops"`
	// Optional. Custom Performance Total Throughput of the pool (in MiB/s).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_storage_pool#total_throughput_mibps GoogleNetappStoragePool#total_throughput_mibps}
	TotalThroughputMibps *string `field:"optional" json:"totalThroughputMibps" yaml:"totalThroughputMibps"`
	// Type of the storage pool.
	//
	// This field is used to control whether the pool supports FILE based volumes only or UNIFIED (both FILE and BLOCK) volumes.
	// If not specified during creation, it defaults to FILE. Possible values: ["STORAGE_POOL_TYPE_UNSPECIFIED", "FILE", "UNIFIED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_storage_pool#type GoogleNetappStoragePool#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Specifies the active zone for regional Flex pools.
	//
	// 'zone' and 'replica_zone' values can be swapped to initiate a
	// [zone switch](https://cloud.google.com/netapp/volumes/docs/configure-and-use/storage-pools/edit-or-delete-storage-pool#switch_active_and_replica_zones).
	// If you want to create a zonal Flex pool, specify a zone name for 'location' and omit 'zone'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_storage_pool#zone GoogleNetappStoragePool#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

