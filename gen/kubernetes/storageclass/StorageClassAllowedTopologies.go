package storageclass

type StorageClassAllowedTopologies struct {
	// match_label_expressions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/storage_class#match_label_expressions StorageClass#match_label_expressions}
	MatchLabelExpressions any `field:"optional" json:"matchLabelExpressions" yaml:"matchLabelExpressions"`
}
