package googlecloudsecuritycomplianceframework


type GoogleCloudSecurityComplianceFrameworkCloudControlDetails struct {
	// Major revision of cloudcontrol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_framework#major_revision_id GoogleCloudSecurityComplianceFramework#major_revision_id}
	MajorRevisionId *string `field:"required" json:"majorRevisionId" yaml:"majorRevisionId"`
	// The name of the CloudControl in the format: “organizations/{organization}/locations/{location}/cloudControls/{cloud-control}”.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_framework#name GoogleCloudSecurityComplianceFramework#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_framework#parameters GoogleCloudSecurityComplianceFramework#parameters}
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

