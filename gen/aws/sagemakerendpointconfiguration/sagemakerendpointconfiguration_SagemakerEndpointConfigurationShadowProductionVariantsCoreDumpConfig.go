package sagemakerendpointconfiguration


type SagemakerEndpointConfigurationShadowProductionVariantsCoreDumpConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_endpoint_configuration#destination_s3_uri SagemakerEndpointConfiguration#destination_s3_uri}.
	DestinationS3Uri *string `field:"required" json:"destinationS3Uri" yaml:"destinationS3Uri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_endpoint_configuration#kms_key_id SagemakerEndpointConfiguration#kms_key_id}.
	KmsKeyId *string `field:"required" json:"kmsKeyId" yaml:"kmsKeyId"`
}
