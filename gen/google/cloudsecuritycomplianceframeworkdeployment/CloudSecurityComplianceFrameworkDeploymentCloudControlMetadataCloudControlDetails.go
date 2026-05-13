package cloudsecuritycomplianceframeworkdeployment


type CloudSecurityComplianceFrameworkDeploymentCloudControlMetadataCloudControlDetails struct {
	// Major revision of cloudcontrol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework_deployment#major_revision_id CloudSecurityComplianceFrameworkDeployment#major_revision_id}
	MajorRevisionId *string `field:"required" json:"majorRevisionId" yaml:"majorRevisionId"`
	// The name of the CloudControl in the format: “organizations/{organization}/locations/{location}/ cloudControls/{cloud-control}”.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework_deployment#name CloudSecurityComplianceFrameworkDeployment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework_deployment#parameters CloudSecurityComplianceFrameworkDeployment#parameters}
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

