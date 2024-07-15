package googlecomputeinstancefromtemplate


type GoogleComputeInstanceFromTemplateAttachedDisk struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_template#device_name GoogleComputeInstanceFromTemplate#device_name}.
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_template#disk_encryption_key_raw GoogleComputeInstanceFromTemplate#disk_encryption_key_raw}.
	DiskEncryptionKeyRaw *string `field:"optional" json:"diskEncryptionKeyRaw" yaml:"diskEncryptionKeyRaw"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_template#disk_encryption_key_sha256 GoogleComputeInstanceFromTemplate#disk_encryption_key_sha256}.
	DiskEncryptionKeySha256 *string `field:"optional" json:"diskEncryptionKeySha256" yaml:"diskEncryptionKeySha256"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_template#kms_key_self_link GoogleComputeInstanceFromTemplate#kms_key_self_link}.
	KmsKeySelfLink *string `field:"optional" json:"kmsKeySelfLink" yaml:"kmsKeySelfLink"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_template#mode GoogleComputeInstanceFromTemplate#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_template#source GoogleComputeInstanceFromTemplate#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
}

