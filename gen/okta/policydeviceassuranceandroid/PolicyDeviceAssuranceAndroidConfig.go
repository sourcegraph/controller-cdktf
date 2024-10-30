package policydeviceassuranceandroid

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PolicyDeviceAssuranceAndroidConfig struct {
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
	// Policy device assurance name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_android#name PolicyDeviceAssuranceAndroid#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of disk encryption type, can be `FULL`, `USER`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_android#disk_encryption_type PolicyDeviceAssuranceAndroid#disk_encryption_type}
	DiskEncryptionType *[]*string `field:"optional" json:"diskEncryptionType" yaml:"diskEncryptionType"`
	// Is the device jailbroken in the device assurance policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_android#jailbreak PolicyDeviceAssuranceAndroid#jailbreak}
	Jailbreak interface{} `field:"optional" json:"jailbreak" yaml:"jailbreak"`
	// Minimum os version of the device in the device assurance policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_android#os_version PolicyDeviceAssuranceAndroid#os_version}
	OsVersion *string `field:"optional" json:"osVersion" yaml:"osVersion"`
	// List of screenlock type, can be `BIOMETRIC` or `BIOMETRIC, PASSCODE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_android#screenlock_type PolicyDeviceAssuranceAndroid#screenlock_type}
	ScreenlockType *[]*string `field:"optional" json:"screenlockType" yaml:"screenlockType"`
	// Indicates if the device contains a secure hardware functionality.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_android#secure_hardware_present PolicyDeviceAssuranceAndroid#secure_hardware_present}
	SecureHardwarePresent interface{} `field:"optional" json:"secureHardwarePresent" yaml:"secureHardwarePresent"`
}

