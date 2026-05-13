package apigeeapideployment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigeeApiDeploymentConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The Apigee Environment associated with the Apigee API deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/apigee_api_deployment#environment ApigeeApiDeployment#environment}
	Environment *string `field:"required" json:"environment" yaml:"environment"`
	// The Apigee Organization associated with the Apigee API deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/apigee_api_deployment#org_id ApigeeApiDeployment#org_id}
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// The Apigee API associated with the Apigee API deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/apigee_api_deployment#proxy_id ApigeeApiDeployment#proxy_id}
	ProxyId *string `field:"required" json:"proxyId" yaml:"proxyId"`
	// The revision of the API proxy to be deployed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/apigee_api_deployment#revision ApigeeApiDeployment#revision}
	Revision *string `field:"required" json:"revision" yaml:"revision"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/apigee_api_deployment#timeouts ApigeeApiDeployment#timeouts}
	Timeouts *ApigeeApiDeploymentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

