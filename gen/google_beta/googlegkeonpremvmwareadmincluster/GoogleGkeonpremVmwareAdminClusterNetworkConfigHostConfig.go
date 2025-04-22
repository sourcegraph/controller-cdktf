package googlegkeonpremvmwareadmincluster


type GoogleGkeonpremVmwareAdminClusterNetworkConfigHostConfig struct {
	// DNS search domains.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#dns_search_domains GoogleGkeonpremVmwareAdminCluster#dns_search_domains}
	DnsSearchDomains *[]*string `field:"optional" json:"dnsSearchDomains" yaml:"dnsSearchDomains"`
	// DNS servers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#dns_servers GoogleGkeonpremVmwareAdminCluster#dns_servers}
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// NTP servers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#ntp_servers GoogleGkeonpremVmwareAdminCluster#ntp_servers}
	NtpServers *[]*string `field:"optional" json:"ntpServers" yaml:"ntpServers"`
}

