package googlegkeonpremvmwarecluster


type GoogleGkeonpremVmwareClusterNetworkConfigHostConfig struct {
	// DNS search domains.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_vmware_cluster#dns_search_domains GoogleGkeonpremVmwareCluster#dns_search_domains}
	DnsSearchDomains *[]*string `field:"optional" json:"dnsSearchDomains" yaml:"dnsSearchDomains"`
	// DNS servers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_vmware_cluster#dns_servers GoogleGkeonpremVmwareCluster#dns_servers}
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// NTP servers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_vmware_cluster#ntp_servers GoogleGkeonpremVmwareCluster#ntp_servers}
	NtpServers *[]*string `field:"optional" json:"ntpServers" yaml:"ntpServers"`
}

