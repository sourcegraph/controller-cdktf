package googlesaasruntimerelease


type GoogleSaasRuntimeReleaseReleaseRequirements struct {
	// A list of releases from which a unit can be upgraded to this one (optional).
	//
	// If left empty no constraints will be applied. When provided,
	// unit upgrade requests to this release will check and enforce this
	// constraint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_release#upgradeable_from_releases GoogleSaasRuntimeRelease#upgradeable_from_releases}
	UpgradeableFromReleases *[]*string `field:"optional" json:"upgradeableFromReleases" yaml:"upgradeableFromReleases"`
}

