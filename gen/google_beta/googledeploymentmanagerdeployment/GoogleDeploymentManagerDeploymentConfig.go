package googledeploymentmanagerdeployment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDeploymentManagerDeploymentConfig struct {
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
	// Unique name for the deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_deployment_manager_deployment#name GoogleDeploymentManagerDeployment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_deployment_manager_deployment#target GoogleDeploymentManagerDeployment#target}
	Target *GoogleDeploymentManagerDeploymentTarget `field:"required" json:"target" yaml:"target"`
	// Set the policy to use for creating new resources.
	//
	// Only used on
	// create and update. Valid values are 'CREATE_OR_ACQUIRE' (default) or
	// 'ACQUIRE'. If set to 'ACQUIRE' and resources do not already exist,
	// the deployment will fail. Note that updating this field does not
	// actually affect the deployment, just how it is updated. Default value: "CREATE_OR_ACQUIRE" Possible values: ["ACQUIRE", "CREATE_OR_ACQUIRE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_deployment_manager_deployment#create_policy GoogleDeploymentManagerDeployment#create_policy}
	CreatePolicy *string `field:"optional" json:"createPolicy" yaml:"createPolicy"`
	// Set the policy to use for deleting new resources on update/delete.
	//
	// Valid values are 'DELETE' (default) or 'ABANDON'. If 'DELETE',
	// resource is deleted after removal from Deployment Manager. If
	// 'ABANDON', the resource is only removed from Deployment Manager
	// and is not actually deleted. Note that updating this field does not
	// actually change the deployment, just how it is updated. Default value: "DELETE" Possible values: ["ABANDON", "DELETE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_deployment_manager_deployment#delete_policy GoogleDeploymentManagerDeployment#delete_policy}
	DeletePolicy *string `field:"optional" json:"deletePolicy" yaml:"deletePolicy"`
	// Optional user-provided description of deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_deployment_manager_deployment#description GoogleDeploymentManagerDeployment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_deployment_manager_deployment#id GoogleDeploymentManagerDeployment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_deployment_manager_deployment#labels GoogleDeploymentManagerDeployment#labels}
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
	// If set to true, a deployment is created with "shell" resources that are not actually instantiated.
	//
	// This allows you to preview a
	// deployment. It can be updated to false to actually deploy
	// with real resources.
	//  ~>**NOTE:** Deployment Manager does not allow update
	// of a deployment in preview (unless updating to preview=false). Thus,
	// Terraform will force-recreate deployments if either preview is updated
	// to true or if other fields are updated while preview is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_deployment_manager_deployment#preview GoogleDeploymentManagerDeployment#preview}
	Preview interface{} `field:"optional" json:"preview" yaml:"preview"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_deployment_manager_deployment#project GoogleDeploymentManagerDeployment#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_deployment_manager_deployment#timeouts GoogleDeploymentManagerDeployment#timeouts}
	Timeouts *GoogleDeploymentManagerDeploymentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

