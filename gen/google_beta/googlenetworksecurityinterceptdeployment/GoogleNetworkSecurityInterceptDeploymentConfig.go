package googlenetworksecurityinterceptdeployment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetworkSecurityInterceptDeploymentConfig struct {
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
	// The regional forwarding rule that fronts the interceptors, for example: 'projects/123456789/regions/us-central1/forwardingRules/my-rule'. See https://google.aip.dev/124.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_network_security_intercept_deployment#forwarding_rule GoogleNetworkSecurityInterceptDeployment#forwarding_rule}
	ForwardingRule *string `field:"required" json:"forwardingRule" yaml:"forwardingRule"`
	// The deployment group that this deployment is a part of, for example: 'projects/123456789/locations/global/interceptDeploymentGroups/my-dg'. See https://google.aip.dev/124.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_network_security_intercept_deployment#intercept_deployment_group GoogleNetworkSecurityInterceptDeployment#intercept_deployment_group}
	InterceptDeploymentGroup *string `field:"required" json:"interceptDeploymentGroup" yaml:"interceptDeploymentGroup"`
	// The ID to use for the new deployment, which will become the final component of the deployment's resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_network_security_intercept_deployment#intercept_deployment_id GoogleNetworkSecurityInterceptDeployment#intercept_deployment_id}
	InterceptDeploymentId *string `field:"required" json:"interceptDeploymentId" yaml:"interceptDeploymentId"`
	// The cloud location of the deployment, e.g. 'us-central1-a' or 'asia-south1-b'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_network_security_intercept_deployment#location GoogleNetworkSecurityInterceptDeployment#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// User-provided description of the deployment. Used as additional context for the deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_network_security_intercept_deployment#description GoogleNetworkSecurityInterceptDeployment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_network_security_intercept_deployment#id GoogleNetworkSecurityInterceptDeployment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels are key/value pairs that help to organize and filter resources.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_network_security_intercept_deployment#labels GoogleNetworkSecurityInterceptDeployment#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_network_security_intercept_deployment#project GoogleNetworkSecurityInterceptDeployment#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_network_security_intercept_deployment#timeouts GoogleNetworkSecurityInterceptDeployment#timeouts}
	Timeouts *GoogleNetworkSecurityInterceptDeploymentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

