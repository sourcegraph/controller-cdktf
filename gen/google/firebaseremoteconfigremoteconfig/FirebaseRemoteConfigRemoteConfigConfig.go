package firebaseremoteconfigremoteconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FirebaseRemoteConfigRemoteConfigConfig struct {
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
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/firebase_remote_config_remote_config#conditions FirebaseRemoteConfigRemoteConfig#conditions}
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/firebase_remote_config_remote_config#id FirebaseRemoteConfigRemoteConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// parameter_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/firebase_remote_config_remote_config#parameter_groups FirebaseRemoteConfigRemoteConfig#parameter_groups}
	ParameterGroups interface{} `field:"optional" json:"parameterGroups" yaml:"parameterGroups"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/firebase_remote_config_remote_config#parameters FirebaseRemoteConfigRemoteConfig#parameters}
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/firebase_remote_config_remote_config#project FirebaseRemoteConfigRemoteConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/firebase_remote_config_remote_config#timeouts FirebaseRemoteConfigRemoteConfig#timeouts}
	Timeouts *FirebaseRemoteConfigRemoteConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

