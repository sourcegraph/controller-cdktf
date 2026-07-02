package preferredpath

type PreferredPathStep struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/preferred_path#link PreferredPath#link}.
	Link *string `field:"optional" json:"link" yaml:"link"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/preferred_path#link_label PreferredPath#link_label}.
	LinkLabel *string `field:"optional" json:"linkLabel" yaml:"linkLabel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/preferred_path#reverse PreferredPath#reverse}.
	Reverse any `field:"optional" json:"reverse" yaml:"reverse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/preferred_path#reverse_from PreferredPath#reverse_from}.
	ReverseFrom *string `field:"optional" json:"reverseFrom" yaml:"reverseFrom"`
}
