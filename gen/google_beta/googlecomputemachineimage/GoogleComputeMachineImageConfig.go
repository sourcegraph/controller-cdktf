package googlecomputemachineimage

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeMachineImageConfig struct {
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
	// Name of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_machine_image#name GoogleComputeMachineImage#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The source instance used to create the machine image.
	//
	// You can provide this as a partial or full URL to the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_machine_image#source_instance GoogleComputeMachineImage#source_instance}
	SourceInstance *string `field:"required" json:"sourceInstance" yaml:"sourceInstance"`
	// A text description of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_machine_image#description GoogleComputeMachineImage#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specify this to create an application consistent machine image by informing the OS to prepare for the snapshot process.
	//
	// Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_machine_image#guest_flush GoogleComputeMachineImage#guest_flush}
	GuestFlush interface{} `field:"optional" json:"guestFlush" yaml:"guestFlush"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_machine_image#id GoogleComputeMachineImage#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// machine_image_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_machine_image#machine_image_encryption_key GoogleComputeMachineImage#machine_image_encryption_key}
	MachineImageEncryptionKey *GoogleComputeMachineImageMachineImageEncryptionKey `field:"optional" json:"machineImageEncryptionKey" yaml:"machineImageEncryptionKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_machine_image#project GoogleComputeMachineImage#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_machine_image#timeouts GoogleComputeMachineImage#timeouts}
	Timeouts *GoogleComputeMachineImageTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

