package googlesccv2projectmuteconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleSccV2ProjectMuteConfigConfig struct {
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
	// An expression that defines the filter to apply across create/update events of findings.
	//
	// While creating a filter string, be mindful of
	// the scope in which the mute configuration is being created. E.g.,
	// If a filter contains project = X but is created under the
	// project = Y scope, it might not match any findings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_scc_v2_project_mute_config#filter GoogleSccV2ProjectMuteConfig#filter}
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// Unique identifier provided by the client within the parent scope.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_scc_v2_project_mute_config#mute_config_id GoogleSccV2ProjectMuteConfig#mute_config_id}
	MuteConfigId *string `field:"required" json:"muteConfigId" yaml:"muteConfigId"`
	// The type of the mute config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_scc_v2_project_mute_config#type GoogleSccV2ProjectMuteConfig#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// A description of the mute config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_scc_v2_project_mute_config#description GoogleSccV2ProjectMuteConfig#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_scc_v2_project_mute_config#id GoogleSccV2ProjectMuteConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// location Id is provided by project. If not provided, Use global as default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_scc_v2_project_mute_config#location GoogleSccV2ProjectMuteConfig#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_scc_v2_project_mute_config#project GoogleSccV2ProjectMuteConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_scc_v2_project_mute_config#timeouts GoogleSccV2ProjectMuteConfig#timeouts}
	Timeouts *GoogleSccV2ProjectMuteConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

