package googlestoragemanagedfolder


type GoogleStorageManagedFolderTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_storage_managed_folder#create GoogleStorageManagedFolder#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_storage_managed_folder#delete GoogleStorageManagedFolder#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_storage_managed_folder#update GoogleStorageManagedFolder#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

