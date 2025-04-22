package googleredisclusterusercreatedconnections

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleRedisClusterUserCreatedConnectionsConfig struct {
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
	// The name of the Redis cluster these endpoints should be added to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_redis_cluster_user_created_connections#name GoogleRedisClusterUserCreatedConnections#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the region of the Redis cluster these endpoints should be added to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_redis_cluster_user_created_connections#region GoogleRedisClusterUserCreatedConnections#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// cluster_endpoints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_redis_cluster_user_created_connections#cluster_endpoints GoogleRedisClusterUserCreatedConnections#cluster_endpoints}
	ClusterEndpoints interface{} `field:"optional" json:"clusterEndpoints" yaml:"clusterEndpoints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_redis_cluster_user_created_connections#id GoogleRedisClusterUserCreatedConnections#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_redis_cluster_user_created_connections#project GoogleRedisClusterUserCreatedConnections#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_redis_cluster_user_created_connections#timeouts GoogleRedisClusterUserCreatedConnections#timeouts}
	Timeouts *GoogleRedisClusterUserCreatedConnectionsTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

