package rediscluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RedisClusterConfig struct {
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
	// Required. Number of shards for the Redis cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/redis_cluster#shard_count RedisCluster#shard_count}
	ShardCount *float64 `field:"required" json:"shardCount" yaml:"shardCount"`
	// Optional.
	//
	// The authorization mode of the Redis cluster. If not provided, auth feature is disabled for the cluster. Default value: "AUTH_MODE_DISABLED" Possible values: ["AUTH_MODE_UNSPECIFIED", "AUTH_MODE_IAM_AUTH", "AUTH_MODE_DISABLED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/redis_cluster#authorization_mode RedisCluster#authorization_mode}
	AuthorizationMode *string `field:"optional" json:"authorizationMode" yaml:"authorizationMode"`
	// automated_backup_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/redis_cluster#automated_backup_config RedisCluster#automated_backup_config}
	AutomatedBackupConfig *RedisClusterAutomatedBackupConfig `field:"optional" json:"automatedBackupConfig" yaml:"automatedBackupConfig"`
	// cross_cluster_replication_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/redis_cluster#cross_cluster_replication_config RedisCluster#cross_cluster_replication_config}
	CrossClusterReplicationConfig *RedisClusterCrossClusterReplicationConfig `field:"optional" json:"crossClusterReplicationConfig" yaml:"crossClusterReplicationConfig"`
	// Optional.
	//
	// Indicates if the cluster is deletion protected or not.
	// If the value if set to true, any delete cluster operation will fail.
	// Default value is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/redis_cluster#deletion_protection_enabled RedisCluster#deletion_protection_enabled}
	DeletionProtectionEnabled interface{} `field:"optional" json:"deletionProtectionEnabled" yaml:"deletionProtectionEnabled"`
	// gcs_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/redis_cluster#gcs_source RedisCluster#gcs_source}
	GcsSource *RedisClusterGcsSource `field:"optional" json:"gcsSource" yaml:"gcsSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/redis_cluster#id RedisCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The KMS key used to encrypt the at-rest data of the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/redis_cluster#kms_key RedisCluster#kms_key}
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Resource labels to represent user provided metadata.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/redis_cluster#labels RedisCluster#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// maintenance_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/redis_cluster#maintenance_policy RedisCluster#maintenance_policy}
	MaintenancePolicy *RedisClusterMaintenancePolicy `field:"optional" json:"maintenancePolicy" yaml:"maintenancePolicy"`
	// This field can be used to trigger self service update to indicate the desired maintenance version.
	//
	// The input to this field can be determined by the available_maintenance_versions field.
	// *Note*: This field can only be specified when updating an existing cluster to a newer version. Downgrades are currently not supported!
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/redis_cluster#maintenance_version RedisCluster#maintenance_version}
	MaintenanceVersion *string `field:"optional" json:"maintenanceVersion" yaml:"maintenanceVersion"`
	// managed_backup_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/redis_cluster#managed_backup_source RedisCluster#managed_backup_source}
	ManagedBackupSource *RedisClusterManagedBackupSource `field:"optional" json:"managedBackupSource" yaml:"managedBackupSource"`
	// Unique name of the resource in this scope including project and location using the form: projects/{projectId}/locations/{locationId}/clusters/{clusterId}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/redis_cluster#name RedisCluster#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The nodeType for the Redis cluster.
	//
	// If not provided, REDIS_HIGHMEM_MEDIUM will be used as default Possible values: ["REDIS_SHARED_CORE_NANO", "REDIS_HIGHMEM_MEDIUM", "REDIS_HIGHMEM_XLARGE", "REDIS_STANDARD_SMALL"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/redis_cluster#node_type RedisCluster#node_type}
	NodeType *string `field:"optional" json:"nodeType" yaml:"nodeType"`
	// persistence_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/redis_cluster#persistence_config RedisCluster#persistence_config}
	PersistenceConfig *RedisClusterPersistenceConfig `field:"optional" json:"persistenceConfig" yaml:"persistenceConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/redis_cluster#project RedisCluster#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// psc_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/redis_cluster#psc_configs RedisCluster#psc_configs}
	PscConfigs interface{} `field:"optional" json:"pscConfigs" yaml:"pscConfigs"`
	// Configure Redis Cluster behavior using a subset of native Redis configuration parameters.
	//
	// Please check Memorystore documentation for the list of supported parameters:
	// https://cloud.google.com/memorystore/docs/cluster/supported-instance-configurations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/redis_cluster#redis_configs RedisCluster#redis_configs}
	RedisConfigs *map[string]*string `field:"optional" json:"redisConfigs" yaml:"redisConfigs"`
	// The name of the region of the Redis cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/redis_cluster#region RedisCluster#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Optional. The number of replica nodes per shard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/redis_cluster#replica_count RedisCluster#replica_count}
	ReplicaCount *float64 `field:"optional" json:"replicaCount" yaml:"replicaCount"`
	// The serverCaMode for the TLS enabled Redis cluster.
	//
	// If not provided, SERVER_CA_MODE_GOOGLE_MANAGED_PER_INSTANCE_CA will be used as default Possible values: ["SERVER_CA_MODE_GOOGLE_MANAGED_PER_INSTANCE_CA", "SERVER_CA_MODE_GOOGLE_MANAGED_SHARED_CA", "SERVER_CA_MODE_CUSTOMER_MANAGED_CAS_CA", "SERVER_CA_MODE_UNSPECIFIED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/redis_cluster#server_ca_mode RedisCluster#server_ca_mode}
	ServerCaMode *string `field:"optional" json:"serverCaMode" yaml:"serverCaMode"`
	// The resource name of the server CA pool for an instance with SERVER_CA_MODE_CUSTOMER_MANAGED_CAS_CA as the server_ca_mode. Format: projects/{project}/locations/{region}/caPools/{caPoolId}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/redis_cluster#server_ca_pool RedisCluster#server_ca_pool}
	ServerCaPool *string `field:"optional" json:"serverCaPool" yaml:"serverCaPool"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/redis_cluster#timeouts RedisCluster#timeouts}
	Timeouts *RedisClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Optional.
	//
	// The in-transit encryption for the Redis cluster.
	// If not provided, encryption is disabled for the cluster. Default value: "TRANSIT_ENCRYPTION_MODE_DISABLED" Possible values: ["TRANSIT_ENCRYPTION_MODE_UNSPECIFIED", "TRANSIT_ENCRYPTION_MODE_DISABLED", "TRANSIT_ENCRYPTION_MODE_SERVER_AUTHENTICATION"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/redis_cluster#transit_encryption_mode RedisCluster#transit_encryption_mode}
	TransitEncryptionMode *string `field:"optional" json:"transitEncryptionMode" yaml:"transitEncryptionMode"`
	// zone_distribution_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/redis_cluster#zone_distribution_config RedisCluster#zone_distribution_config}
	ZoneDistributionConfig *RedisClusterZoneDistributionConfig `field:"optional" json:"zoneDistributionConfig" yaml:"zoneDistributionConfig"`
}

