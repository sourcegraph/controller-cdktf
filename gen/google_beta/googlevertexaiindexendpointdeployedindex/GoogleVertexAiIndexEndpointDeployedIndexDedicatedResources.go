package googlevertexaiindexendpointdeployedindex


type GoogleVertexAiIndexEndpointDeployedIndexDedicatedResources struct {
	// machine_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_vertex_ai_index_endpoint_deployed_index#machine_spec GoogleVertexAiIndexEndpointDeployedIndex#machine_spec}
	MachineSpec *GoogleVertexAiIndexEndpointDeployedIndexDedicatedResourcesMachineSpec `field:"required" json:"machineSpec" yaml:"machineSpec"`
	// The minimum number of machine replicas this DeployedModel will be always deployed on.
	//
	// This value must be greater than or equal to 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_vertex_ai_index_endpoint_deployed_index#min_replica_count GoogleVertexAiIndexEndpointDeployedIndex#min_replica_count}
	MinReplicaCount *float64 `field:"required" json:"minReplicaCount" yaml:"minReplicaCount"`
	// The maximum number of replicas this DeployedModel may be deployed on when the traffic against it increases.
	//
	// If maxReplicaCount is not set, the default value is minReplicaCount
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_vertex_ai_index_endpoint_deployed_index#max_replica_count GoogleVertexAiIndexEndpointDeployedIndex#max_replica_count}
	MaxReplicaCount *float64 `field:"optional" json:"maxReplicaCount" yaml:"maxReplicaCount"`
}

