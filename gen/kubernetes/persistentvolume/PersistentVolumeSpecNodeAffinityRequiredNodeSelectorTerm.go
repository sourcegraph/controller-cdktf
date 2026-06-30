package persistentvolume

type PersistentVolumeSpecNodeAffinityRequiredNodeSelectorTerm struct {
	// match_expressions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/persistent_volume#match_expressions PersistentVolume#match_expressions}
	MatchExpressions any `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	// match_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/persistent_volume#match_fields PersistentVolume#match_fields}
	MatchFields any `field:"optional" json:"matchFields" yaml:"matchFields"`
}
