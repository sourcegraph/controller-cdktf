package googlecontainerazurecluster


type GoogleContainerAzureClusterControlPlaneSshConfig struct {
	// The SSH public key data for VMs managed by Anthos.
	//
	// This accepts the authorized_keys file format used in OpenSSH according to the sshd(8) manual page.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_azure_cluster#authorized_key GoogleContainerAzureCluster#authorized_key}
	AuthorizedKey *string `field:"required" json:"authorizedKey" yaml:"authorizedKey"`
}

