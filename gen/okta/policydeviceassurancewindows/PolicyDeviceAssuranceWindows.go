package policydeviceassurancewindows

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/policydeviceassurancewindows/internal"
)

// Represents a {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_windows okta_policy_device_assurance_windows}.
type PolicyDeviceAssuranceWindows interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedBy() *string
	CreatedDate() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DiskEncryptionType() *[]*string
	SetDiskEncryptionType(val *[]*string)
	DiskEncryptionTypeInput() *[]*string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	LastUpdate() *string
	LastUpdatedBy() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OsVersion() *string
	SetOsVersion(val *string)
	OsVersionInput() *string
	Platform() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ScreenlockType() *[]*string
	SetScreenlockType(val *[]*string)
	ScreenlockTypeInput() *[]*string
	SecureHardwarePresent() interface{}
	SetSecureHardwarePresent(val interface{})
	SecureHardwarePresentInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ThirdPartySignalProviders() interface{}
	SetThirdPartySignalProviders(val interface{})
	ThirdPartySignalProvidersInput() interface{}
	TpspBrowserVersion() *string
	SetTpspBrowserVersion(val *string)
	TpspBrowserVersionInput() *string
	TpspBuiltinDnsClientEnabled() interface{}
	SetTpspBuiltinDnsClientEnabled(val interface{})
	TpspBuiltinDnsClientEnabledInput() interface{}
	TpspChromeRemoteDesktopAppBlocked() interface{}
	SetTpspChromeRemoteDesktopAppBlocked(val interface{})
	TpspChromeRemoteDesktopAppBlockedInput() interface{}
	TpspCrowdStrikeAgentId() *string
	SetTpspCrowdStrikeAgentId(val *string)
	TpspCrowdStrikeAgentIdInput() *string
	TpspCrowdStrikeCustomerId() *string
	SetTpspCrowdStrikeCustomerId(val *string)
	TpspCrowdStrikeCustomerIdInput() *string
	TpspDeviceEnrollmentDomain() *string
	SetTpspDeviceEnrollmentDomain(val *string)
	TpspDeviceEnrollmentDomainInput() *string
	TpspDiskEncrypted() interface{}
	SetTpspDiskEncrypted(val interface{})
	TpspDiskEncryptedInput() interface{}
	TpspKeyTrustLevel() *string
	SetTpspKeyTrustLevel(val *string)
	TpspKeyTrustLevelInput() *string
	TpspOsFirewall() interface{}
	SetTpspOsFirewall(val interface{})
	TpspOsFirewallInput() interface{}
	TpspOsVersion() *string
	SetTpspOsVersion(val *string)
	TpspOsVersionInput() *string
	TpspPasswordProctectionWarningTrigger() *string
	SetTpspPasswordProctectionWarningTrigger(val *string)
	TpspPasswordProctectionWarningTriggerInput() *string
	TpspRealtimeUrlCheckMode() interface{}
	SetTpspRealtimeUrlCheckMode(val interface{})
	TpspRealtimeUrlCheckModeInput() interface{}
	TpspSafeBrowsingProtectionLevel() *string
	SetTpspSafeBrowsingProtectionLevel(val *string)
	TpspSafeBrowsingProtectionLevelInput() *string
	TpspScreenLockSecured() interface{}
	SetTpspScreenLockSecured(val interface{})
	TpspScreenLockSecuredInput() interface{}
	TpspSecureBootEnabled() interface{}
	SetTpspSecureBootEnabled(val interface{})
	TpspSecureBootEnabledInput() interface{}
	TpspSiteIsolationEnabled() interface{}
	SetTpspSiteIsolationEnabled(val interface{})
	TpspSiteIsolationEnabledInput() interface{}
	TpspThirdPartyBlockingEnabled() interface{}
	SetTpspThirdPartyBlockingEnabled(val interface{})
	TpspThirdPartyBlockingEnabledInput() interface{}
	TpspWindowsMachineDomain() *string
	SetTpspWindowsMachineDomain(val *string)
	TpspWindowsMachineDomainInput() *string
	TpspWindowsUserDomain() *string
	SetTpspWindowsUserDomain(val *string)
	TpspWindowsUserDomainInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetDiskEncryptionType()
	ResetOsVersion()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetScreenlockType()
	ResetSecureHardwarePresent()
	ResetThirdPartySignalProviders()
	ResetTpspBrowserVersion()
	ResetTpspBuiltinDnsClientEnabled()
	ResetTpspChromeRemoteDesktopAppBlocked()
	ResetTpspCrowdStrikeAgentId()
	ResetTpspCrowdStrikeCustomerId()
	ResetTpspDeviceEnrollmentDomain()
	ResetTpspDiskEncrypted()
	ResetTpspKeyTrustLevel()
	ResetTpspOsFirewall()
	ResetTpspOsVersion()
	ResetTpspPasswordProctectionWarningTrigger()
	ResetTpspRealtimeUrlCheckMode()
	ResetTpspSafeBrowsingProtectionLevel()
	ResetTpspScreenLockSecured()
	ResetTpspSecureBootEnabled()
	ResetTpspSiteIsolationEnabled()
	ResetTpspThirdPartyBlockingEnabled()
	ResetTpspWindowsMachineDomain()
	ResetTpspWindowsUserDomain()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for PolicyDeviceAssuranceWindows
type jsiiProxy_PolicyDeviceAssuranceWindows struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) CreatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) CreatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) DiskEncryptionType() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"diskEncryptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) DiskEncryptionTypeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"diskEncryptionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) LastUpdate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) LastUpdatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) OsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) OsVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) Platform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) ScreenlockType() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"screenlockType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) ScreenlockTypeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"screenlockTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) SecureHardwarePresent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secureHardwarePresent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) SecureHardwarePresentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secureHardwarePresentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) ThirdPartySignalProviders() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"thirdPartySignalProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) ThirdPartySignalProvidersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"thirdPartySignalProvidersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspBrowserVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpspBrowserVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspBrowserVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpspBrowserVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspBuiltinDnsClientEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspBuiltinDnsClientEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspBuiltinDnsClientEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspBuiltinDnsClientEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspChromeRemoteDesktopAppBlocked() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspChromeRemoteDesktopAppBlocked",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspChromeRemoteDesktopAppBlockedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspChromeRemoteDesktopAppBlockedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspCrowdStrikeAgentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpspCrowdStrikeAgentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspCrowdStrikeAgentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpspCrowdStrikeAgentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspCrowdStrikeCustomerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpspCrowdStrikeCustomerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspCrowdStrikeCustomerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpspCrowdStrikeCustomerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspDeviceEnrollmentDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpspDeviceEnrollmentDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspDeviceEnrollmentDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpspDeviceEnrollmentDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspDiskEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspDiskEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspDiskEncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspDiskEncryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspKeyTrustLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpspKeyTrustLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspKeyTrustLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpspKeyTrustLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspOsFirewall() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspOsFirewall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspOsFirewallInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspOsFirewallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspOsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpspOsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspOsVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpspOsVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspPasswordProctectionWarningTrigger() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpspPasswordProctectionWarningTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspPasswordProctectionWarningTriggerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpspPasswordProctectionWarningTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspRealtimeUrlCheckMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspRealtimeUrlCheckMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspRealtimeUrlCheckModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspRealtimeUrlCheckModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspSafeBrowsingProtectionLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpspSafeBrowsingProtectionLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspSafeBrowsingProtectionLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpspSafeBrowsingProtectionLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspScreenLockSecured() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspScreenLockSecured",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspScreenLockSecuredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspScreenLockSecuredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspSecureBootEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspSecureBootEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspSecureBootEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspSecureBootEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspSiteIsolationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspSiteIsolationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspSiteIsolationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspSiteIsolationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspThirdPartyBlockingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspThirdPartyBlockingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspThirdPartyBlockingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspThirdPartyBlockingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspWindowsMachineDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpspWindowsMachineDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspWindowsMachineDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpspWindowsMachineDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspWindowsUserDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpspWindowsUserDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows) TpspWindowsUserDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpspWindowsUserDomainInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_windows okta_policy_device_assurance_windows} Resource.
func NewPolicyDeviceAssuranceWindows(scope constructs.Construct, id *string, config *PolicyDeviceAssuranceWindowsConfig) PolicyDeviceAssuranceWindows {
	_init_.Initialize()

	if err := validateNewPolicyDeviceAssuranceWindowsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PolicyDeviceAssuranceWindows{}

	_jsii_.Create(
		"@cdktf/provider-okta.policyDeviceAssuranceWindows.PolicyDeviceAssuranceWindows",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_windows okta_policy_device_assurance_windows} Resource.
func NewPolicyDeviceAssuranceWindows_Override(p PolicyDeviceAssuranceWindows, scope constructs.Construct, id *string, config *PolicyDeviceAssuranceWindowsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.policyDeviceAssuranceWindows.PolicyDeviceAssuranceWindows",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows)SetDiskEncryptionType(val *[]*string) {
	if err := j.validateSetDiskEncryptionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskEncryptionType",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows)SetOsVersion(val *string) {
	if err := j.validateSetOsVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osVersion",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows)SetScreenlockType(val *[]*string) {
	if err := j.validateSetScreenlockTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"screenlockType",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows)SetSecureHardwarePresent(val interface{}) {
	if err := j.validateSetSecureHardwarePresentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secureHardwarePresent",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows)SetThirdPartySignalProviders(val interface{}) {
	if err := j.validateSetThirdPartySignalProvidersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thirdPartySignalProviders",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows)SetTpspBrowserVersion(val *string) {
	if err := j.validateSetTpspBrowserVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspBrowserVersion",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows)SetTpspBuiltinDnsClientEnabled(val interface{}) {
	if err := j.validateSetTpspBuiltinDnsClientEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspBuiltinDnsClientEnabled",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows)SetTpspChromeRemoteDesktopAppBlocked(val interface{}) {
	if err := j.validateSetTpspChromeRemoteDesktopAppBlockedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspChromeRemoteDesktopAppBlocked",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows)SetTpspCrowdStrikeAgentId(val *string) {
	if err := j.validateSetTpspCrowdStrikeAgentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspCrowdStrikeAgentId",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows)SetTpspCrowdStrikeCustomerId(val *string) {
	if err := j.validateSetTpspCrowdStrikeCustomerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspCrowdStrikeCustomerId",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows)SetTpspDeviceEnrollmentDomain(val *string) {
	if err := j.validateSetTpspDeviceEnrollmentDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspDeviceEnrollmentDomain",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows)SetTpspDiskEncrypted(val interface{}) {
	if err := j.validateSetTpspDiskEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspDiskEncrypted",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows)SetTpspKeyTrustLevel(val *string) {
	if err := j.validateSetTpspKeyTrustLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspKeyTrustLevel",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows)SetTpspOsFirewall(val interface{}) {
	if err := j.validateSetTpspOsFirewallParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspOsFirewall",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows)SetTpspOsVersion(val *string) {
	if err := j.validateSetTpspOsVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspOsVersion",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows)SetTpspPasswordProctectionWarningTrigger(val *string) {
	if err := j.validateSetTpspPasswordProctectionWarningTriggerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspPasswordProctectionWarningTrigger",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows)SetTpspRealtimeUrlCheckMode(val interface{}) {
	if err := j.validateSetTpspRealtimeUrlCheckModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspRealtimeUrlCheckMode",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows)SetTpspSafeBrowsingProtectionLevel(val *string) {
	if err := j.validateSetTpspSafeBrowsingProtectionLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspSafeBrowsingProtectionLevel",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows)SetTpspScreenLockSecured(val interface{}) {
	if err := j.validateSetTpspScreenLockSecuredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspScreenLockSecured",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows)SetTpspSecureBootEnabled(val interface{}) {
	if err := j.validateSetTpspSecureBootEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspSecureBootEnabled",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows)SetTpspSiteIsolationEnabled(val interface{}) {
	if err := j.validateSetTpspSiteIsolationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspSiteIsolationEnabled",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows)SetTpspThirdPartyBlockingEnabled(val interface{}) {
	if err := j.validateSetTpspThirdPartyBlockingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspThirdPartyBlockingEnabled",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows)SetTpspWindowsMachineDomain(val *string) {
	if err := j.validateSetTpspWindowsMachineDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspWindowsMachineDomain",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceWindows)SetTpspWindowsUserDomain(val *string) {
	if err := j.validateSetTpspWindowsUserDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspWindowsUserDomain",
		val,
	)
}

// Generates CDKTF code for importing a PolicyDeviceAssuranceWindows resource upon running "cdktf plan <stack-name>".
func PolicyDeviceAssuranceWindows_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validatePolicyDeviceAssuranceWindows_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.policyDeviceAssuranceWindows.PolicyDeviceAssuranceWindows",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func PolicyDeviceAssuranceWindows_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePolicyDeviceAssuranceWindows_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.policyDeviceAssuranceWindows.PolicyDeviceAssuranceWindows",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PolicyDeviceAssuranceWindows_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePolicyDeviceAssuranceWindows_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.policyDeviceAssuranceWindows.PolicyDeviceAssuranceWindows",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PolicyDeviceAssuranceWindows_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePolicyDeviceAssuranceWindows_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.policyDeviceAssuranceWindows.PolicyDeviceAssuranceWindows",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PolicyDeviceAssuranceWindows_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-okta.policyDeviceAssuranceWindows.PolicyDeviceAssuranceWindows",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) ResetDiskEncryptionType() {
	_jsii_.InvokeVoid(
		p,
		"resetDiskEncryptionType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) ResetOsVersion() {
	_jsii_.InvokeVoid(
		p,
		"resetOsVersion",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) ResetScreenlockType() {
	_jsii_.InvokeVoid(
		p,
		"resetScreenlockType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) ResetSecureHardwarePresent() {
	_jsii_.InvokeVoid(
		p,
		"resetSecureHardwarePresent",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) ResetThirdPartySignalProviders() {
	_jsii_.InvokeVoid(
		p,
		"resetThirdPartySignalProviders",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) ResetTpspBrowserVersion() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspBrowserVersion",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) ResetTpspBuiltinDnsClientEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspBuiltinDnsClientEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) ResetTpspChromeRemoteDesktopAppBlocked() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspChromeRemoteDesktopAppBlocked",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) ResetTpspCrowdStrikeAgentId() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspCrowdStrikeAgentId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) ResetTpspCrowdStrikeCustomerId() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspCrowdStrikeCustomerId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) ResetTpspDeviceEnrollmentDomain() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspDeviceEnrollmentDomain",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) ResetTpspDiskEncrypted() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspDiskEncrypted",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) ResetTpspKeyTrustLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspKeyTrustLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) ResetTpspOsFirewall() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspOsFirewall",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) ResetTpspOsVersion() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspOsVersion",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) ResetTpspPasswordProctectionWarningTrigger() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspPasswordProctectionWarningTrigger",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) ResetTpspRealtimeUrlCheckMode() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspRealtimeUrlCheckMode",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) ResetTpspSafeBrowsingProtectionLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspSafeBrowsingProtectionLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) ResetTpspScreenLockSecured() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspScreenLockSecured",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) ResetTpspSecureBootEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspSecureBootEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) ResetTpspSiteIsolationEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspSiteIsolationEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) ResetTpspThirdPartyBlockingEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspThirdPartyBlockingEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) ResetTpspWindowsMachineDomain() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspWindowsMachineDomain",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) ResetTpspWindowsUserDomain() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspWindowsUserDomain",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyDeviceAssuranceWindows) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

