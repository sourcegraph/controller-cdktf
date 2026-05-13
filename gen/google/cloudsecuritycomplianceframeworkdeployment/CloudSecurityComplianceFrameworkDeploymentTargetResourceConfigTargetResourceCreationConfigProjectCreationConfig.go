package cloudsecuritycomplianceframeworkdeployment


type CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfig struct {
	// Billing account id to be used for the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework_deployment#billing_account_id CloudSecurityComplianceFrameworkDeployment#billing_account_id}
	BillingAccountId *string `field:"required" json:"billingAccountId" yaml:"billingAccountId"`
	// organizations/{org} or folders/{folder}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework_deployment#parent CloudSecurityComplianceFrameworkDeployment#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// Display name of the project to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework_deployment#project_display_name CloudSecurityComplianceFrameworkDeployment#project_display_name}
	ProjectDisplayName *string `field:"required" json:"projectDisplayName" yaml:"projectDisplayName"`
}

