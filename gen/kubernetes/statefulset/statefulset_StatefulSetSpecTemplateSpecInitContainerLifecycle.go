package statefulset


type StatefulSetSpecTemplateSpecInitContainerLifecycle struct {
	// post_start block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#post_start StatefulSet#post_start}
	PostStart interface{} `field:"optional" json:"postStart" yaml:"postStart"`
	// pre_stop block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#pre_stop StatefulSet#pre_stop}
	PreStop interface{} `field:"optional" json:"preStop" yaml:"preStop"`
}

