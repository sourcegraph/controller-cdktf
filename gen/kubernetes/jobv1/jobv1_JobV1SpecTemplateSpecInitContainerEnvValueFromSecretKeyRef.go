package jobv1


type JobV1SpecTemplateSpecInitContainerEnvValueFromSecretKeyRef struct {
	// The key of the secret to select from. Must be a valid secret key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/job_v1#key JobV1#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/job_v1#name JobV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specify whether the Secret or its key must be defined.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/job_v1#optional JobV1#optional}
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

