package sagemakerfeaturegroup


type SagemakerFeatureGroupOnlineStoreConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/sagemaker_feature_group#enable_online_store SagemakerFeatureGroup#enable_online_store}.
	EnableOnlineStore interface{} `field:"optional" json:"enableOnlineStore" yaml:"enableOnlineStore"`
	// security_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/sagemaker_feature_group#security_config SagemakerFeatureGroup#security_config}
	SecurityConfig *SagemakerFeatureGroupOnlineStoreConfigSecurityConfig `field:"optional" json:"securityConfig" yaml:"securityConfig"`
}

