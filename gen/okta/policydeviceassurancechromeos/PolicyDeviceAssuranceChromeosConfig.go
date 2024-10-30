package policydeviceassurancechromeos

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PolicyDeviceAssuranceChromeosConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Name of the device assurance policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_chromeos#name PolicyDeviceAssuranceChromeos#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Third party signal provider allow screen lock.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_chromeos#tpsp_allow_screen_lock PolicyDeviceAssuranceChromeos#tpsp_allow_screen_lock}
	TpspAllowScreenLock interface{} `field:"optional" json:"tpspAllowScreenLock" yaml:"tpspAllowScreenLock"`
	// Third party signal provider minimum browser version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_chromeos#tpsp_browser_version PolicyDeviceAssuranceChromeos#tpsp_browser_version}
	TpspBrowserVersion *string `field:"optional" json:"tpspBrowserVersion" yaml:"tpspBrowserVersion"`
	// Third party signal provider builtin dns client enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_chromeos#tpsp_builtin_dns_client_enabled PolicyDeviceAssuranceChromeos#tpsp_builtin_dns_client_enabled}
	TpspBuiltinDnsClientEnabled interface{} `field:"optional" json:"tpspBuiltinDnsClientEnabled" yaml:"tpspBuiltinDnsClientEnabled"`
	// Third party signal provider chrome remote desktop app blocked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_chromeos#tpsp_chrome_remote_desktop_app_blocked PolicyDeviceAssuranceChromeos#tpsp_chrome_remote_desktop_app_blocked}
	TpspChromeRemoteDesktopAppBlocked interface{} `field:"optional" json:"tpspChromeRemoteDesktopAppBlocked" yaml:"tpspChromeRemoteDesktopAppBlocked"`
	// Third party signal provider device enrollment domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_chromeos#tpsp_device_enrollment_domain PolicyDeviceAssuranceChromeos#tpsp_device_enrollment_domain}
	TpspDeviceEnrollmentDomain *string `field:"optional" json:"tpspDeviceEnrollmentDomain" yaml:"tpspDeviceEnrollmentDomain"`
	// Third party signal provider disk encrypted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_chromeos#tpsp_disk_encrypted PolicyDeviceAssuranceChromeos#tpsp_disk_encrypted}
	TpspDiskEncrypted interface{} `field:"optional" json:"tpspDiskEncrypted" yaml:"tpspDiskEncrypted"`
	// Third party signal provider key trust level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_chromeos#tpsp_key_trust_level PolicyDeviceAssuranceChromeos#tpsp_key_trust_level}
	TpspKeyTrustLevel *string `field:"optional" json:"tpspKeyTrustLevel" yaml:"tpspKeyTrustLevel"`
	// Third party signal provider os firewall.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_chromeos#tpsp_os_firewall PolicyDeviceAssuranceChromeos#tpsp_os_firewall}
	TpspOsFirewall interface{} `field:"optional" json:"tpspOsFirewall" yaml:"tpspOsFirewall"`
	// Third party signal provider minimum os version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_chromeos#tpsp_os_version PolicyDeviceAssuranceChromeos#tpsp_os_version}
	TpspOsVersion *string `field:"optional" json:"tpspOsVersion" yaml:"tpspOsVersion"`
	// Third party signal provider password protection warning trigger.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_chromeos#tpsp_password_proctection_warning_trigger PolicyDeviceAssuranceChromeos#tpsp_password_proctection_warning_trigger}
	TpspPasswordProctectionWarningTrigger *string `field:"optional" json:"tpspPasswordProctectionWarningTrigger" yaml:"tpspPasswordProctectionWarningTrigger"`
	// Third party signal provider realtime url check mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_chromeos#tpsp_realtime_url_check_mode PolicyDeviceAssuranceChromeos#tpsp_realtime_url_check_mode}
	TpspRealtimeUrlCheckMode interface{} `field:"optional" json:"tpspRealtimeUrlCheckMode" yaml:"tpspRealtimeUrlCheckMode"`
	// Third party signal provider safe browsing protection level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_chromeos#tpsp_safe_browsing_protection_level PolicyDeviceAssuranceChromeos#tpsp_safe_browsing_protection_level}
	TpspSafeBrowsingProtectionLevel *string `field:"optional" json:"tpspSafeBrowsingProtectionLevel" yaml:"tpspSafeBrowsingProtectionLevel"`
	// Third party signal provider screen lock secure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_chromeos#tpsp_screen_lock_secured PolicyDeviceAssuranceChromeos#tpsp_screen_lock_secured}
	TpspScreenLockSecured interface{} `field:"optional" json:"tpspScreenLockSecured" yaml:"tpspScreenLockSecured"`
	// Third party signal provider site isolation enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_chromeos#tpsp_site_isolation_enabled PolicyDeviceAssuranceChromeos#tpsp_site_isolation_enabled}
	TpspSiteIsolationEnabled interface{} `field:"optional" json:"tpspSiteIsolationEnabled" yaml:"tpspSiteIsolationEnabled"`
}

