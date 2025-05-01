package googlememorystoreinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleMemorystoreInstanceConfig struct {
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
	// Required. The ID to use for the instance, which will become the final component of the instance's resource name.
	//
	// This value is subject to the following restrictions:
	//
	// * Must be 4-63 characters in length
	// * Must begin with a letter or digit
	// * Must contain only lowercase letters, digits, and hyphens
	// * Must not end with a hyphen
	// * Must be unique within a location
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_memorystore_instance#instance_id GoogleMemorystoreInstance#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// Resource ID segment making up resource 'name'.
	//
	// It identifies the resource within its parent collection as described in https://google.aip.dev/122. See documentation for resource type 'memorystore.googleapis.com/CertificateAuthority'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_memorystore_instance#location GoogleMemorystoreInstance#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Required. Number of shards for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_memorystore_instance#shard_count GoogleMemorystoreInstance#shard_count}
	ShardCount *float64 `field:"required" json:"shardCount" yaml:"shardCount"`
	// Optional. Immutable. Authorization mode of the instance. Possible values:  AUTH_DISABLED IAM_AUTH.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_memorystore_instance#authorization_mode GoogleMemorystoreInstance#authorization_mode}
	AuthorizationMode *string `field:"optional" json:"authorizationMode" yaml:"authorizationMode"`
	// automated_backup_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_memorystore_instance#automated_backup_config GoogleMemorystoreInstance#automated_backup_config}
	AutomatedBackupConfig *GoogleMemorystoreInstanceAutomatedBackupConfig `field:"optional" json:"automatedBackupConfig" yaml:"automatedBackupConfig"`
	// cross_instance_replication_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_memorystore_instance#cross_instance_replication_config GoogleMemorystoreInstance#cross_instance_replication_config}
	CrossInstanceReplicationConfig *GoogleMemorystoreInstanceCrossInstanceReplicationConfig `field:"optional" json:"crossInstanceReplicationConfig" yaml:"crossInstanceReplicationConfig"`
	// Optional. If set to true deletion of the instance will fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_memorystore_instance#deletion_protection_enabled GoogleMemorystoreInstance#deletion_protection_enabled}
	DeletionProtectionEnabled interface{} `field:"optional" json:"deletionProtectionEnabled" yaml:"deletionProtectionEnabled"`
	// desired_psc_auto_connections block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_memorystore_instance#desired_psc_auto_connections GoogleMemorystoreInstance#desired_psc_auto_connections}
	DesiredPscAutoConnections interface{} `field:"optional" json:"desiredPscAutoConnections" yaml:"desiredPscAutoConnections"`
	// Optional. User-provided engine configurations for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_memorystore_instance#engine_configs GoogleMemorystoreInstance#engine_configs}
	EngineConfigs *map[string]*string `field:"optional" json:"engineConfigs" yaml:"engineConfigs"`
	// Optional. Engine version of the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_memorystore_instance#engine_version GoogleMemorystoreInstance#engine_version}
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// gcs_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_memorystore_instance#gcs_source GoogleMemorystoreInstance#gcs_source}
	GcsSource *GoogleMemorystoreInstanceGcsSource `field:"optional" json:"gcsSource" yaml:"gcsSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_memorystore_instance#id GoogleMemorystoreInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional. Labels to represent user-provided metadata.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_memorystore_instance#labels GoogleMemorystoreInstance#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// maintenance_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_memorystore_instance#maintenance_policy GoogleMemorystoreInstance#maintenance_policy}
	MaintenancePolicy *GoogleMemorystoreInstanceMaintenancePolicy `field:"optional" json:"maintenancePolicy" yaml:"maintenancePolicy"`
	// managed_backup_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_memorystore_instance#managed_backup_source GoogleMemorystoreInstance#managed_backup_source}
	ManagedBackupSource *GoogleMemorystoreInstanceManagedBackupSource `field:"optional" json:"managedBackupSource" yaml:"managedBackupSource"`
	// Optional. cluster or cluster-disabled.   Possible values:  CLUSTER  CLUSTER_DISABLED Possible values: ["CLUSTER", "CLUSTER_DISABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_memorystore_instance#mode GoogleMemorystoreInstance#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Optional. Machine type for individual nodes of the instance.   Possible values:  SHARED_CORE_NANO HIGHMEM_MEDIUM HIGHMEM_XLARGE STANDARD_SMALL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_memorystore_instance#node_type GoogleMemorystoreInstance#node_type}
	NodeType *string `field:"optional" json:"nodeType" yaml:"nodeType"`
	// persistence_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_memorystore_instance#persistence_config GoogleMemorystoreInstance#persistence_config}
	PersistenceConfig *GoogleMemorystoreInstancePersistenceConfig `field:"optional" json:"persistenceConfig" yaml:"persistenceConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_memorystore_instance#project GoogleMemorystoreInstance#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Optional. Number of replica nodes per shard. If omitted the default is 0 replicas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_memorystore_instance#replica_count GoogleMemorystoreInstance#replica_count}
	ReplicaCount *float64 `field:"optional" json:"replicaCount" yaml:"replicaCount"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_memorystore_instance#timeouts GoogleMemorystoreInstance#timeouts}
	Timeouts *GoogleMemorystoreInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Optional. Immutable. In-transit encryption mode of the instance.   Possible values:  TRANSIT_ENCRYPTION_DISABLED SERVER_AUTHENTICATION.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_memorystore_instance#transit_encryption_mode GoogleMemorystoreInstance#transit_encryption_mode}
	TransitEncryptionMode *string `field:"optional" json:"transitEncryptionMode" yaml:"transitEncryptionMode"`
	// zone_distribution_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_memorystore_instance#zone_distribution_config GoogleMemorystoreInstance#zone_distribution_config}
	ZoneDistributionConfig *GoogleMemorystoreInstanceZoneDistributionConfig `field:"optional" json:"zoneDistributionConfig" yaml:"zoneDistributionConfig"`
}

