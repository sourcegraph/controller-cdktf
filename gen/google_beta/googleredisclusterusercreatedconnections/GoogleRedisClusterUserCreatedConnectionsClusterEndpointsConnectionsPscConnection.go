package googleredisclusterusercreatedconnections


type GoogleRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnection struct {
	// The IP allocated on the consumer network for the PSC forwarding rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_redis_cluster_user_created_connections#address GoogleRedisClusterUserCreatedConnections#address}
	Address *string `field:"required" json:"address" yaml:"address"`
	// The URI of the consumer side forwarding rule. Format: projects/{project}/regions/{region}/forwardingRules/{forwarding_rule}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_redis_cluster_user_created_connections#forwarding_rule GoogleRedisClusterUserCreatedConnections#forwarding_rule}
	ForwardingRule *string `field:"required" json:"forwardingRule" yaml:"forwardingRule"`
	// The consumer network where the IP address resides, in the form of projects/{project_id}/global/networks/{network_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_redis_cluster_user_created_connections#network GoogleRedisClusterUserCreatedConnections#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// The PSC connection id of the forwarding rule connected to the service attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_redis_cluster_user_created_connections#psc_connection_id GoogleRedisClusterUserCreatedConnections#psc_connection_id}
	PscConnectionId *string `field:"required" json:"pscConnectionId" yaml:"pscConnectionId"`
	// The service attachment which is the target of the PSC connection, in the form of projects/{project-id}/regions/{region}/serviceAttachments/{service-attachment-id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_redis_cluster_user_created_connections#service_attachment GoogleRedisClusterUserCreatedConnections#service_attachment}
	ServiceAttachment *string `field:"required" json:"serviceAttachment" yaml:"serviceAttachment"`
	// The consumer project_id where the forwarding rule is created from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_redis_cluster_user_created_connections#project_id GoogleRedisClusterUserCreatedConnections#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

