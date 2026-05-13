package apigeeenvironmentapirevisiondeployment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigeeEnvironmentApiRevisionDeploymentConfig struct {
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
	// Apigee API proxy name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/apigee_environment_api_revision_deployment#api ApigeeEnvironmentApiRevisionDeployment#api}
	Api *string `field:"required" json:"api" yaml:"api"`
	// Apigee environment name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/apigee_environment_api_revision_deployment#environment ApigeeEnvironmentApiRevisionDeployment#environment}
	Environment *string `field:"required" json:"environment" yaml:"environment"`
	// Apigee organization ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/apigee_environment_api_revision_deployment#org_id ApigeeEnvironmentApiRevisionDeployment#org_id}
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// API proxy revision number to deploy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/apigee_environment_api_revision_deployment#revision ApigeeEnvironmentApiRevisionDeployment#revision}
	Revision *float64 `field:"required" json:"revision" yaml:"revision"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/apigee_environment_api_revision_deployment#id ApigeeEnvironmentApiRevisionDeployment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If true, replaces other deployed revisions of this proxy in the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/apigee_environment_api_revision_deployment#override ApigeeEnvironmentApiRevisionDeployment#override}
	Override interface{} `field:"optional" json:"override" yaml:"override"`
	// If true, enables sequenced rollout for safe traffic switching.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/apigee_environment_api_revision_deployment#sequenced_rollout ApigeeEnvironmentApiRevisionDeployment#sequenced_rollout}
	SequencedRollout interface{} `field:"optional" json:"sequencedRollout" yaml:"sequencedRollout"`
	// Optional service account the deployed proxy runs as.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/apigee_environment_api_revision_deployment#service_account ApigeeEnvironmentApiRevisionDeployment#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/apigee_environment_api_revision_deployment#timeouts ApigeeEnvironmentApiRevisionDeployment#timeouts}
	Timeouts *ApigeeEnvironmentApiRevisionDeploymentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

