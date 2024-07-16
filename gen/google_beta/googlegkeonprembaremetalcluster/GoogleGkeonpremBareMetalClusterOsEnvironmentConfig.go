package googlegkeonprembaremetalcluster


type GoogleGkeonpremBareMetalClusterOsEnvironmentConfig struct {
	// Whether the package repo should not be included when initializing bare metal machines.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#package_repo_excluded GoogleGkeonpremBareMetalCluster#package_repo_excluded}
	PackageRepoExcluded interface{} `field:"required" json:"packageRepoExcluded" yaml:"packageRepoExcluded"`
}

