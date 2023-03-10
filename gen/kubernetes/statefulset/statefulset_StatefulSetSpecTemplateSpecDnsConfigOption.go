package statefulset


type StatefulSetSpecTemplateSpecDnsConfigOption struct {
	// Name of the option.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#name StatefulSet#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of the option. Optional: Defaults to empty.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#value StatefulSet#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

