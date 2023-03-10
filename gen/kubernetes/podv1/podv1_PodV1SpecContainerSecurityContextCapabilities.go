package podv1


type PodV1SpecContainerSecurityContextCapabilities struct {
	// Added capabilities.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_v1#add PodV1#add}
	Add *[]*string `field:"optional" json:"add" yaml:"add"`
	// Removed capabilities.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_v1#drop PodV1#drop}
	Drop *[]*string `field:"optional" json:"drop" yaml:"drop"`
}

