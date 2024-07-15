package googleapphubserviceprojectattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleApphubServiceProjectAttachmentConfig struct {
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
	// The service project attachment identifier must contain the project_id of the service project specified in the service_project_attachment.service_project field. Hint: "projects/{project_id}"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apphub_service_project_attachment#service_project_attachment_id GoogleApphubServiceProjectAttachment#service_project_attachment_id}
	ServiceProjectAttachmentId *string `field:"required" json:"serviceProjectAttachmentId" yaml:"serviceProjectAttachmentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apphub_service_project_attachment#id GoogleApphubServiceProjectAttachment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apphub_service_project_attachment#project GoogleApphubServiceProjectAttachment#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// "Immutable.
	//
	// Service project name in the format: \"projects/abc\"
	// or \"projects/123\". As input, project name with either project id or number
	// are accepted. As output, this field will contain project number. "
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apphub_service_project_attachment#service_project GoogleApphubServiceProjectAttachment#service_project}
	ServiceProject *string `field:"optional" json:"serviceProject" yaml:"serviceProject"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apphub_service_project_attachment#timeouts GoogleApphubServiceProjectAttachment#timeouts}
	Timeouts *GoogleApphubServiceProjectAttachmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

