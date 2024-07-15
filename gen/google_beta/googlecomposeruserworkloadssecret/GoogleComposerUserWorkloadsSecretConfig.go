package googlecomposeruserworkloadssecret

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComposerUserWorkloadsSecretConfig struct {
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
	// Name of the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_composer_user_workloads_secret#environment GoogleComposerUserWorkloadsSecret#environment}
	Environment *string `field:"required" json:"environment" yaml:"environment"`
	// Name of the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_composer_user_workloads_secret#name GoogleComposerUserWorkloadsSecret#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A map of the secret data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_composer_user_workloads_secret#data GoogleComposerUserWorkloadsSecret#data}
	Data *map[string]*string `field:"optional" json:"data" yaml:"data"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_composer_user_workloads_secret#id GoogleComposerUserWorkloadsSecret#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ID of the project in which the resource belongs.
	//
	// If it is not provided, the provider project is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_composer_user_workloads_secret#project GoogleComposerUserWorkloadsSecret#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The location or Compute Engine region for the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_composer_user_workloads_secret#region GoogleComposerUserWorkloadsSecret#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_composer_user_workloads_secret#timeouts GoogleComposerUserWorkloadsSecret#timeouts}
	Timeouts *GoogleComposerUserWorkloadsSecretTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

