package policydeviceassurancemacos

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PolicyDeviceAssuranceMacosConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_macos#name PolicyDeviceAssuranceMacos#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of disk encryption type, can be `ALL_INTERNAL_VOLUMES`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_macos#disk_encryption_type PolicyDeviceAssuranceMacos#disk_encryption_type}
	DiskEncryptionType *[]*string `field:"optional" json:"diskEncryptionType" yaml:"diskEncryptionType"`
	// Minimum os version of the device in the device assurance policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_macos#os_version PolicyDeviceAssuranceMacos#os_version}
	OsVersion *string `field:"optional" json:"osVersion" yaml:"osVersion"`
	// List of screenlock type, can be `BIOMETRIC` or `BIOMETRIC, PASSCODE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_macos#screenlock_type PolicyDeviceAssuranceMacos#screenlock_type}
	ScreenlockType *[]*string `field:"optional" json:"screenlockType" yaml:"screenlockType"`
	// Is the device secure with hardware in the device assurance policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_macos#secure_hardware_present PolicyDeviceAssuranceMacos#secure_hardware_present}
	SecureHardwarePresent interface{} `field:"optional" json:"secureHardwarePresent" yaml:"secureHardwarePresent"`
	// Check to include third party signal provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_macos#third_party_signal_providers PolicyDeviceAssuranceMacos#third_party_signal_providers}
	ThirdPartySignalProviders interface{} `field:"optional" json:"thirdPartySignalProviders" yaml:"thirdPartySignalProviders"`
	// Third party signal provider minimum browser version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_macos#tpsp_browser_version PolicyDeviceAssuranceMacos#tpsp_browser_version}
	TpspBrowserVersion *string `field:"optional" json:"tpspBrowserVersion" yaml:"tpspBrowserVersion"`
	// Third party signal provider builtin dns client enable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_macos#tpsp_builtin_dns_client_enabled PolicyDeviceAssuranceMacos#tpsp_builtin_dns_client_enabled}
	TpspBuiltinDnsClientEnabled interface{} `field:"optional" json:"tpspBuiltinDnsClientEnabled" yaml:"tpspBuiltinDnsClientEnabled"`
	// Third party signal provider chrome remote desktop app blocked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_macos#tpsp_chrome_remote_desktop_app_blocked PolicyDeviceAssuranceMacos#tpsp_chrome_remote_desktop_app_blocked}
	TpspChromeRemoteDesktopAppBlocked interface{} `field:"optional" json:"tpspChromeRemoteDesktopAppBlocked" yaml:"tpspChromeRemoteDesktopAppBlocked"`
	// Third party signal provider device enrollment domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_macos#tpsp_device_enrollment_domain PolicyDeviceAssuranceMacos#tpsp_device_enrollment_domain}
	TpspDeviceEnrollmentDomain *string `field:"optional" json:"tpspDeviceEnrollmentDomain" yaml:"tpspDeviceEnrollmentDomain"`
	// Third party signal provider disk encrypted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_macos#tpsp_disk_encrypted PolicyDeviceAssuranceMacos#tpsp_disk_encrypted}
	TpspDiskEncrypted interface{} `field:"optional" json:"tpspDiskEncrypted" yaml:"tpspDiskEncrypted"`
	// Third party signal provider key trust level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_macos#tpsp_key_trust_level PolicyDeviceAssuranceMacos#tpsp_key_trust_level}
	TpspKeyTrustLevel *string `field:"optional" json:"tpspKeyTrustLevel" yaml:"tpspKeyTrustLevel"`
	// Third party signal provider os firewall.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_macos#tpsp_os_firewall PolicyDeviceAssuranceMacos#tpsp_os_firewall}
	TpspOsFirewall interface{} `field:"optional" json:"tpspOsFirewall" yaml:"tpspOsFirewall"`
	// Third party signal provider minimum os version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_macos#tpsp_os_version PolicyDeviceAssuranceMacos#tpsp_os_version}
	TpspOsVersion *string `field:"optional" json:"tpspOsVersion" yaml:"tpspOsVersion"`
	// Third party signal provider password protection warning trigger.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_macos#tpsp_password_proctection_warning_trigger PolicyDeviceAssuranceMacos#tpsp_password_proctection_warning_trigger}
	TpspPasswordProctectionWarningTrigger *string `field:"optional" json:"tpspPasswordProctectionWarningTrigger" yaml:"tpspPasswordProctectionWarningTrigger"`
	// Third party signal provider realtime url check mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_macos#tpsp_realtime_url_check_mode PolicyDeviceAssuranceMacos#tpsp_realtime_url_check_mode}
	TpspRealtimeUrlCheckMode interface{} `field:"optional" json:"tpspRealtimeUrlCheckMode" yaml:"tpspRealtimeUrlCheckMode"`
	// Third party signal provider safe browsing protection level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_macos#tpsp_safe_browsing_protection_level PolicyDeviceAssuranceMacos#tpsp_safe_browsing_protection_level}
	TpspSafeBrowsingProtectionLevel *string `field:"optional" json:"tpspSafeBrowsingProtectionLevel" yaml:"tpspSafeBrowsingProtectionLevel"`
	// Third party signal provider screen lock secure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_macos#tpsp_screen_lock_secured PolicyDeviceAssuranceMacos#tpsp_screen_lock_secured}
	TpspScreenLockSecured interface{} `field:"optional" json:"tpspScreenLockSecured" yaml:"tpspScreenLockSecured"`
	// Third party signal provider site isolation enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_macos#tpsp_site_isolation_enabled PolicyDeviceAssuranceMacos#tpsp_site_isolation_enabled}
	TpspSiteIsolationEnabled interface{} `field:"optional" json:"tpspSiteIsolationEnabled" yaml:"tpspSiteIsolationEnabled"`
}

