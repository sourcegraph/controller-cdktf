package fsxopenzfsvolume


type FsxOpenzfsVolumeTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/fsx_openzfs_volume#create FsxOpenzfsVolume#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/fsx_openzfs_volume#delete FsxOpenzfsVolume#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/fsx_openzfs_volume#update FsxOpenzfsVolume#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

