package googlecomputeinstancefromtemplate


type GoogleComputeInstanceFromTemplateBootDisk struct {
	// Whether the disk will be auto-deleted when the instance is deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_instance_from_template#auto_delete GoogleComputeInstanceFromTemplate#auto_delete}
	AutoDelete interface{} `field:"optional" json:"autoDelete" yaml:"autoDelete"`
	// Name with which attached disk will be accessible under /dev/disk/by-id/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_instance_from_template#device_name GoogleComputeInstanceFromTemplate#device_name}
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// A 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to encrypt this disk.
	//
	// Only one of kms_key_self_link, disk_encryption_key_raw and disk_encryption_key_rsa may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_instance_from_template#disk_encryption_key_raw GoogleComputeInstanceFromTemplate#disk_encryption_key_raw}
	DiskEncryptionKeyRaw *string `field:"optional" json:"diskEncryptionKeyRaw" yaml:"diskEncryptionKeyRaw"`
	// Specifies an RFC 4648 base64 encoded, RSA-wrapped 2048-bit customer-supplied encryption key to either encrypt or decrypt this resource.
	//
	// Only one of kms_key_self_link, disk_encryption_key_raw and disk_encryption_key_rsa may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_instance_from_template#disk_encryption_key_rsa GoogleComputeInstanceFromTemplate#disk_encryption_key_rsa}
	DiskEncryptionKeyRsa *string `field:"optional" json:"diskEncryptionKeyRsa" yaml:"diskEncryptionKeyRsa"`
	// The service account being used for the encryption request for the given KMS key.
	//
	// If absent, the Compute Engine default service account is used
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_instance_from_template#disk_encryption_service_account GoogleComputeInstanceFromTemplate#disk_encryption_service_account}
	DiskEncryptionServiceAccount *string `field:"optional" json:"diskEncryptionServiceAccount" yaml:"diskEncryptionServiceAccount"`
	// A list of features to enable on the guest operating system. Applicable only for bootable images.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_instance_from_template#guest_os_features GoogleComputeInstanceFromTemplate#guest_os_features}
	GuestOsFeatures *[]*string `field:"optional" json:"guestOsFeatures" yaml:"guestOsFeatures"`
	// initialize_params block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_instance_from_template#initialize_params GoogleComputeInstanceFromTemplate#initialize_params}
	InitializeParams *GoogleComputeInstanceFromTemplateBootDiskInitializeParams `field:"optional" json:"initializeParams" yaml:"initializeParams"`
	// The disk interface used for attaching this disk.
	//
	// One of SCSI or NVME. (This field is shared with attached_disk and only used for specific cases, please don't specify this field without advice from Google.)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_instance_from_template#interface GoogleComputeInstanceFromTemplate#interface}
	Interface *string `field:"optional" json:"interface" yaml:"interface"`
	// The self_link of the encryption key that is stored in Google Cloud KMS to encrypt this disk.
	//
	// Only one of kms_key_self_link, disk_encryption_key_raw and disk_encryption_key_rsa may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_instance_from_template#kms_key_self_link GoogleComputeInstanceFromTemplate#kms_key_self_link}
	KmsKeySelfLink *string `field:"optional" json:"kmsKeySelfLink" yaml:"kmsKeySelfLink"`
	// Read/write mode for the disk. One of "READ_ONLY" or "READ_WRITE".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_instance_from_template#mode GoogleComputeInstanceFromTemplate#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The name or self_link of the disk attached to this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_instance_from_template#source GoogleComputeInstanceFromTemplate#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
}

