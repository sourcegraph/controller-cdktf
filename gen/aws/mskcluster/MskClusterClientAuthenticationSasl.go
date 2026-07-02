package mskcluster

type MskClusterClientAuthenticationSasl struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/msk_cluster#iam MskCluster#iam}.
	Iam any `field:"optional" json:"iam" yaml:"iam"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/msk_cluster#scram MskCluster#scram}.
	Scram any `field:"optional" json:"scram" yaml:"scram"`
}
