package googleapihubhostprojectregistration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleApihubHostProjectRegistrationConfig struct {
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
	// Required.
	//
	// Immutable. Google cloud project name in the format: "projects/abc" or "projects/123".
	// As input, project name with either project id or number are accepted.
	// As output, this field will contain project number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_apihub_host_project_registration#gcp_project GoogleApihubHostProjectRegistration#gcp_project}
	GcpProject *string `field:"required" json:"gcpProject" yaml:"gcpProject"`
	// Required.
	//
	// The ID to use for the Host Project Registration, which will become the
	// final component of the host project registration's resource name. The ID
	// must be the same as the Google cloud project specified in the
	// host_project_registration.gcp_project field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_apihub_host_project_registration#host_project_registration_id GoogleApihubHostProjectRegistration#host_project_registration_id}
	HostProjectRegistrationId *string `field:"required" json:"hostProjectRegistrationId" yaml:"hostProjectRegistrationId"`
	// Part of 'parent'. See documentation of 'projectsId'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_apihub_host_project_registration#location GoogleApihubHostProjectRegistration#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_apihub_host_project_registration#id GoogleApihubHostProjectRegistration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_apihub_host_project_registration#project GoogleApihubHostProjectRegistration#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_apihub_host_project_registration#timeouts GoogleApihubHostProjectRegistration#timeouts}
	Timeouts *GoogleApihubHostProjectRegistrationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

