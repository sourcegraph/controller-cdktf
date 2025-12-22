package poller


type PollerMongodbatlas struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#private_key Poller#private_key}.
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#public_key Poller#public_key}.
	PublicKey *string `field:"required" json:"publicKey" yaml:"publicKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#exclude_groups Poller#exclude_groups}.
	ExcludeGroups *[]*string `field:"optional" json:"excludeGroups" yaml:"excludeGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#include_groups Poller#include_groups}.
	IncludeGroups *[]*string `field:"optional" json:"includeGroups" yaml:"includeGroups"`
}

