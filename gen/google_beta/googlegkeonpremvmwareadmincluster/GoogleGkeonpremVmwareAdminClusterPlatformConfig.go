package googlegkeonpremvmwareadmincluster


type GoogleGkeonpremVmwareAdminClusterPlatformConfig struct {
	// The required platform version e.g. 1.13.1. If the current platform version is lower than the target version, the platform version will be updated to the target version. If the target version is not installed in the platform (bundle versions), download the target version bundle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#required_platform_version GoogleGkeonpremVmwareAdminCluster#required_platform_version}
	RequiredPlatformVersion *string `field:"optional" json:"requiredPlatformVersion" yaml:"requiredPlatformVersion"`
}

