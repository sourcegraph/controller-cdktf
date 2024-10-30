package policydeviceassurancewindows

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PolicyDeviceAssuranceWindowsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_windows#name PolicyDeviceAssuranceWindows#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of disk encryption type, can be `ALL_INTERNAL_VOLUMES`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_windows#disk_encryption_type PolicyDeviceAssuranceWindows#disk_encryption_type}
	DiskEncryptionType *[]*string `field:"optional" json:"diskEncryptionType" yaml:"diskEncryptionType"`
	// Minimum os version of the device in the device assurance policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_windows#os_version PolicyDeviceAssuranceWindows#os_version}
	OsVersion *string `field:"optional" json:"osVersion" yaml:"osVersion"`
	// List of screenlock type, can be `BIOMETRIC` or `BIOMETRIC, PASSCODE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_windows#screenlock_type PolicyDeviceAssuranceWindows#screenlock_type}
	ScreenlockType *[]*string `field:"optional" json:"screenlockType" yaml:"screenlockType"`
	// Is the device secure with hardware in the device assurance policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_windows#secure_hardware_present PolicyDeviceAssuranceWindows#secure_hardware_present}
	SecureHardwarePresent interface{} `field:"optional" json:"secureHardwarePresent" yaml:"secureHardwarePresent"`
	// Check to include third party signal provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_windows#third_party_signal_providers PolicyDeviceAssuranceWindows#third_party_signal_providers}
	ThirdPartySignalProviders interface{} `field:"optional" json:"thirdPartySignalProviders" yaml:"thirdPartySignalProviders"`
	// Third party signal provider minimum browser version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_windows#tpsp_browser_version PolicyDeviceAssuranceWindows#tpsp_browser_version}
	TpspBrowserVersion *string `field:"optional" json:"tpspBrowserVersion" yaml:"tpspBrowserVersion"`
	// Third party signal provider builtin dns client enable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_windows#tpsp_builtin_dns_client_enabled PolicyDeviceAssuranceWindows#tpsp_builtin_dns_client_enabled}
	TpspBuiltinDnsClientEnabled interface{} `field:"optional" json:"tpspBuiltinDnsClientEnabled" yaml:"tpspBuiltinDnsClientEnabled"`
	// Third party signal provider chrome remote desktop app blocked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_windows#tpsp_chrome_remote_desktop_app_blocked PolicyDeviceAssuranceWindows#tpsp_chrome_remote_desktop_app_blocked}
	TpspChromeRemoteDesktopAppBlocked interface{} `field:"optional" json:"tpspChromeRemoteDesktopAppBlocked" yaml:"tpspChromeRemoteDesktopAppBlocked"`
	// Third party signal provider crowdstrike agent id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_windows#tpsp_crowd_strike_agent_id PolicyDeviceAssuranceWindows#tpsp_crowd_strike_agent_id}
	TpspCrowdStrikeAgentId *string `field:"optional" json:"tpspCrowdStrikeAgentId" yaml:"tpspCrowdStrikeAgentId"`
	// Third party signal provider crowdstrike user id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_windows#tpsp_crowd_strike_customer_id PolicyDeviceAssuranceWindows#tpsp_crowd_strike_customer_id}
	TpspCrowdStrikeCustomerId *string `field:"optional" json:"tpspCrowdStrikeCustomerId" yaml:"tpspCrowdStrikeCustomerId"`
	// Third party signal provider device enrollment domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_windows#tpsp_device_enrollment_domain PolicyDeviceAssuranceWindows#tpsp_device_enrollment_domain}
	TpspDeviceEnrollmentDomain *string `field:"optional" json:"tpspDeviceEnrollmentDomain" yaml:"tpspDeviceEnrollmentDomain"`
	// Third party signal provider disk encrypted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_windows#tpsp_disk_encrypted PolicyDeviceAssuranceWindows#tpsp_disk_encrypted}
	TpspDiskEncrypted interface{} `field:"optional" json:"tpspDiskEncrypted" yaml:"tpspDiskEncrypted"`
	// Third party signal provider key trust level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_windows#tpsp_key_trust_level PolicyDeviceAssuranceWindows#tpsp_key_trust_level}
	TpspKeyTrustLevel *string `field:"optional" json:"tpspKeyTrustLevel" yaml:"tpspKeyTrustLevel"`
	// Third party signal provider os firewall.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_windows#tpsp_os_firewall PolicyDeviceAssuranceWindows#tpsp_os_firewall}
	TpspOsFirewall interface{} `field:"optional" json:"tpspOsFirewall" yaml:"tpspOsFirewall"`
	// Third party signal provider minimum os version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_windows#tpsp_os_version PolicyDeviceAssuranceWindows#tpsp_os_version}
	TpspOsVersion *string `field:"optional" json:"tpspOsVersion" yaml:"tpspOsVersion"`
	// Third party signal provider password protection warning trigger.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_windows#tpsp_password_proctection_warning_trigger PolicyDeviceAssuranceWindows#tpsp_password_proctection_warning_trigger}
	TpspPasswordProctectionWarningTrigger *string `field:"optional" json:"tpspPasswordProctectionWarningTrigger" yaml:"tpspPasswordProctectionWarningTrigger"`
	// Third party signal provider realtime url check mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_windows#tpsp_realtime_url_check_mode PolicyDeviceAssuranceWindows#tpsp_realtime_url_check_mode}
	TpspRealtimeUrlCheckMode interface{} `field:"optional" json:"tpspRealtimeUrlCheckMode" yaml:"tpspRealtimeUrlCheckMode"`
	// Third party signal provider safe browsing protection level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_windows#tpsp_safe_browsing_protection_level PolicyDeviceAssuranceWindows#tpsp_safe_browsing_protection_level}
	TpspSafeBrowsingProtectionLevel *string `field:"optional" json:"tpspSafeBrowsingProtectionLevel" yaml:"tpspSafeBrowsingProtectionLevel"`
	// Third party signal provider screen lock secure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_windows#tpsp_screen_lock_secured PolicyDeviceAssuranceWindows#tpsp_screen_lock_secured}
	TpspScreenLockSecured interface{} `field:"optional" json:"tpspScreenLockSecured" yaml:"tpspScreenLockSecured"`
	// Third party signal provider secure boot enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_windows#tpsp_secure_boot_enabled PolicyDeviceAssuranceWindows#tpsp_secure_boot_enabled}
	TpspSecureBootEnabled interface{} `field:"optional" json:"tpspSecureBootEnabled" yaml:"tpspSecureBootEnabled"`
	// Third party signal provider site isolation enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_windows#tpsp_site_isolation_enabled PolicyDeviceAssuranceWindows#tpsp_site_isolation_enabled}
	TpspSiteIsolationEnabled interface{} `field:"optional" json:"tpspSiteIsolationEnabled" yaml:"tpspSiteIsolationEnabled"`
	// Third party signal provider third party blocking enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_windows#tpsp_third_party_blocking_enabled PolicyDeviceAssuranceWindows#tpsp_third_party_blocking_enabled}
	TpspThirdPartyBlockingEnabled interface{} `field:"optional" json:"tpspThirdPartyBlockingEnabled" yaml:"tpspThirdPartyBlockingEnabled"`
	// Third party signal provider windows machine domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_windows#tpsp_windows_machine_domain PolicyDeviceAssuranceWindows#tpsp_windows_machine_domain}
	TpspWindowsMachineDomain *string `field:"optional" json:"tpspWindowsMachineDomain" yaml:"tpspWindowsMachineDomain"`
	// Third party signal provider windows user domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_windows#tpsp_windows_user_domain PolicyDeviceAssuranceWindows#tpsp_windows_user_domain}
	TpspWindowsUserDomain *string `field:"optional" json:"tpspWindowsUserDomain" yaml:"tpspWindowsUserDomain"`
}

