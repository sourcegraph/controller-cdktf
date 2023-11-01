package googlerediscluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleRedisClusterConfig struct {
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
	// psc_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.4.0/docs/resources/google_redis_cluster#psc_configs GoogleRedisCluster#psc_configs}
	PscConfigs interface{} `field:"required" json:"pscConfigs" yaml:"pscConfigs"`
	// Required. Number of shards for the Redis cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.4.0/docs/resources/google_redis_cluster#shard_count GoogleRedisCluster#shard_count}
	ShardCount *float64 `field:"required" json:"shardCount" yaml:"shardCount"`
	// Optional.
	//
	// The authorization mode of the Redis cluster. If not provided, auth feature is disabled for the cluster. Default value: "AUTH_MODE_DISABLED" Possible values: ["AUTH_MODE_UNSPECIFIED", "AUTH_MODE_IAM_AUTH", "AUTH_MODE_DISABLED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.4.0/docs/resources/google_redis_cluster#authorization_mode GoogleRedisCluster#authorization_mode}
	AuthorizationMode *string `field:"optional" json:"authorizationMode" yaml:"authorizationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.4.0/docs/resources/google_redis_cluster#id GoogleRedisCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Unique name of the resource in this scope including project and location using the form: projects/{projectId}/locations/{locationId}/clusters/{clusterId}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.4.0/docs/resources/google_redis_cluster#name GoogleRedisCluster#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.4.0/docs/resources/google_redis_cluster#project GoogleRedisCluster#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The name of the region of the Redis cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.4.0/docs/resources/google_redis_cluster#region GoogleRedisCluster#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Optional. The number of replica nodes per shard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.4.0/docs/resources/google_redis_cluster#replica_count GoogleRedisCluster#replica_count}
	ReplicaCount *float64 `field:"optional" json:"replicaCount" yaml:"replicaCount"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.4.0/docs/resources/google_redis_cluster#timeouts GoogleRedisCluster#timeouts}
	Timeouts *GoogleRedisClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Optional.
	//
	// The in-transit encryption for the Redis cluster.
	// If not provided, encryption is disabled for the cluster. Default value: "TRANSIT_ENCRYPTION_MODE_DISABLED" Possible values: ["TRANSIT_ENCRYPTION_MODE_UNSPECIFIED", "TRANSIT_ENCRYPTION_MODE_DISABLED", "TRANSIT_ENCRYPTION_MODE_SERVER_AUTHENTICATION"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.4.0/docs/resources/google_redis_cluster#transit_encryption_mode GoogleRedisCluster#transit_encryption_mode}
	TransitEncryptionMode *string `field:"optional" json:"transitEncryptionMode" yaml:"transitEncryptionMode"`
}

