package computeinstancefromtemplate

type ComputeInstanceFromTemplateSchedulingOnInstanceStopAction struct {
	// If true, the contents of any attached Local SSD disks will be discarded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_instance_from_template#discard_local_ssd ComputeInstanceFromTemplate#discard_local_ssd}
	DiscardLocalSsd any `field:"optional" json:"discardLocalSsd" yaml:"discardLocalSsd"`
}
