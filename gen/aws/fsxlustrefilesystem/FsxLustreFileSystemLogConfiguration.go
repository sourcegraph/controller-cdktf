package fsxlustrefilesystem


type FsxLustreFileSystemLogConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/fsx_lustre_file_system#destination FsxLustreFileSystem#destination}.
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/fsx_lustre_file_system#level FsxLustreFileSystem#level}.
	Level *string `field:"optional" json:"level" yaml:"level"`
}

