package googleclouddeploytarget


type GoogleClouddeployTargetAssociatedEntities struct {
	// The name for the key in the map for which this object is mapped to in the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_clouddeploy_target#entity_id GoogleClouddeployTarget#entity_id}
	EntityId *string `field:"required" json:"entityId" yaml:"entityId"`
	// anthos_clusters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_clouddeploy_target#anthos_clusters GoogleClouddeployTarget#anthos_clusters}
	AnthosClusters interface{} `field:"optional" json:"anthosClusters" yaml:"anthosClusters"`
	// gke_clusters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_clouddeploy_target#gke_clusters GoogleClouddeployTarget#gke_clusters}
	GkeClusters interface{} `field:"optional" json:"gkeClusters" yaml:"gkeClusters"`
}

