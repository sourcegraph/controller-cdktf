package comprehenddocumentclassifier


type ComprehendDocumentClassifierOutputDataConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/comprehend_document_classifier#s3_uri ComprehendDocumentClassifier#s3_uri}.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/comprehend_document_classifier#kms_key_id ComprehendDocumentClassifier#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

