package policydeviceassurancechromeos

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/policydeviceassurancechromeos/internal"
)

// Represents a {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_chromeos okta_policy_device_assurance_chromeos}.
type PolicyDeviceAssuranceChromeos interface {
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TpspAllowScreenLock() interface{}
	SetTpspAllowScreenLock(val interface{})
	TpspAllowScreenLockInput() interface{}
	TpspBrowserVersion() *string
	SetTpspBrowserVersion(val *string)
	TpspBrowserVersionInput() *string
	TpspBuiltinDnsClientEnabled() interface{}
	SetTpspBuiltinDnsClientEnabled(val interface{})
	TpspBuiltinDnsClientEnabledInput() interface{}
	TpspChromeRemoteDesktopAppBlocked() interface{}
	SetTpspChromeRemoteDesktopAppBlocked(val interface{})
	TpspChromeRemoteDesktopAppBlockedInput() interface{}
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
	TpspSiteIsolationEnabled() interface{}
	SetTpspSiteIsolationEnabled(val interface{})
	TpspSiteIsolationEnabledInput() interface{}
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
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTpspAllowScreenLock()
	ResetTpspBrowserVersion()
	ResetTpspBuiltinDnsClientEnabled()
	ResetTpspChromeRemoteDesktopAppBlocked()
	ResetTpspDeviceEnrollmentDomain()
	ResetTpspDiskEncrypted()
	ResetTpspKeyTrustLevel()
	ResetTpspOsFirewall()
	ResetTpspOsVersion()
	ResetTpspPasswordProctectionWarningTrigger()
	ResetTpspRealtimeUrlCheckMode()
	ResetTpspSafeBrowsingProtectionLevel()
	ResetTpspScreenLockSecured()
	ResetTpspSiteIsolationEnabled()
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

// The jsii proxy struct for PolicyDeviceAssuranceChromeos
type jsiiProxy_PolicyDeviceAssuranceChromeos struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) CreatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) CreatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) LastUpdate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) LastUpdatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) Platform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) TpspAllowScreenLock() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspAllowScreenLock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) TpspAllowScreenLockInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspAllowScreenLockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) TpspBrowserVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpspBrowserVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) TpspBrowserVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpspBrowserVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) TpspBuiltinDnsClientEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspBuiltinDnsClientEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) TpspBuiltinDnsClientEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspBuiltinDnsClientEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) TpspChromeRemoteDesktopAppBlocked() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspChromeRemoteDesktopAppBlocked",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) TpspChromeRemoteDesktopAppBlockedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspChromeRemoteDesktopAppBlockedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) TpspDeviceEnrollmentDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpspDeviceEnrollmentDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) TpspDeviceEnrollmentDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpspDeviceEnrollmentDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) TpspDiskEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspDiskEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) TpspDiskEncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspDiskEncryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) TpspKeyTrustLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpspKeyTrustLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) TpspKeyTrustLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpspKeyTrustLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) TpspOsFirewall() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspOsFirewall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) TpspOsFirewallInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspOsFirewallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) TpspOsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpspOsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) TpspOsVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpspOsVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) TpspPasswordProctectionWarningTrigger() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpspPasswordProctectionWarningTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) TpspPasswordProctectionWarningTriggerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpspPasswordProctectionWarningTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) TpspRealtimeUrlCheckMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspRealtimeUrlCheckMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) TpspRealtimeUrlCheckModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspRealtimeUrlCheckModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) TpspSafeBrowsingProtectionLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpspSafeBrowsingProtectionLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) TpspSafeBrowsingProtectionLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpspSafeBrowsingProtectionLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) TpspScreenLockSecured() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspScreenLockSecured",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) TpspScreenLockSecuredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspScreenLockSecuredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) TpspSiteIsolationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspSiteIsolationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos) TpspSiteIsolationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpspSiteIsolationEnabledInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_chromeos okta_policy_device_assurance_chromeos} Resource.
func NewPolicyDeviceAssuranceChromeos(scope constructs.Construct, id *string, config *PolicyDeviceAssuranceChromeosConfig) PolicyDeviceAssuranceChromeos {
	_init_.Initialize()

	if err := validateNewPolicyDeviceAssuranceChromeosParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PolicyDeviceAssuranceChromeos{}

	_jsii_.Create(
		"@cdktf/provider-okta.policyDeviceAssuranceChromeos.PolicyDeviceAssuranceChromeos",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_device_assurance_chromeos okta_policy_device_assurance_chromeos} Resource.
func NewPolicyDeviceAssuranceChromeos_Override(p PolicyDeviceAssuranceChromeos, scope constructs.Construct, id *string, config *PolicyDeviceAssuranceChromeosConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.policyDeviceAssuranceChromeos.PolicyDeviceAssuranceChromeos",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos)SetTpspAllowScreenLock(val interface{}) {
	if err := j.validateSetTpspAllowScreenLockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspAllowScreenLock",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos)SetTpspBrowserVersion(val *string) {
	if err := j.validateSetTpspBrowserVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspBrowserVersion",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos)SetTpspBuiltinDnsClientEnabled(val interface{}) {
	if err := j.validateSetTpspBuiltinDnsClientEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspBuiltinDnsClientEnabled",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos)SetTpspChromeRemoteDesktopAppBlocked(val interface{}) {
	if err := j.validateSetTpspChromeRemoteDesktopAppBlockedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspChromeRemoteDesktopAppBlocked",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos)SetTpspDeviceEnrollmentDomain(val *string) {
	if err := j.validateSetTpspDeviceEnrollmentDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspDeviceEnrollmentDomain",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos)SetTpspDiskEncrypted(val interface{}) {
	if err := j.validateSetTpspDiskEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspDiskEncrypted",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos)SetTpspKeyTrustLevel(val *string) {
	if err := j.validateSetTpspKeyTrustLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspKeyTrustLevel",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos)SetTpspOsFirewall(val interface{}) {
	if err := j.validateSetTpspOsFirewallParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspOsFirewall",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos)SetTpspOsVersion(val *string) {
	if err := j.validateSetTpspOsVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspOsVersion",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos)SetTpspPasswordProctectionWarningTrigger(val *string) {
	if err := j.validateSetTpspPasswordProctectionWarningTriggerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspPasswordProctectionWarningTrigger",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos)SetTpspRealtimeUrlCheckMode(val interface{}) {
	if err := j.validateSetTpspRealtimeUrlCheckModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspRealtimeUrlCheckMode",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos)SetTpspSafeBrowsingProtectionLevel(val *string) {
	if err := j.validateSetTpspSafeBrowsingProtectionLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspSafeBrowsingProtectionLevel",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos)SetTpspScreenLockSecured(val interface{}) {
	if err := j.validateSetTpspScreenLockSecuredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspScreenLockSecured",
		val,
	)
}

func (j *jsiiProxy_PolicyDeviceAssuranceChromeos)SetTpspSiteIsolationEnabled(val interface{}) {
	if err := j.validateSetTpspSiteIsolationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpspSiteIsolationEnabled",
		val,
	)
}

// Generates CDKTF code for importing a PolicyDeviceAssuranceChromeos resource upon running "cdktf plan <stack-name>".
func PolicyDeviceAssuranceChromeos_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validatePolicyDeviceAssuranceChromeos_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.policyDeviceAssuranceChromeos.PolicyDeviceAssuranceChromeos",
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
func PolicyDeviceAssuranceChromeos_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePolicyDeviceAssuranceChromeos_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.policyDeviceAssuranceChromeos.PolicyDeviceAssuranceChromeos",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PolicyDeviceAssuranceChromeos_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePolicyDeviceAssuranceChromeos_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.policyDeviceAssuranceChromeos.PolicyDeviceAssuranceChromeos",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PolicyDeviceAssuranceChromeos_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePolicyDeviceAssuranceChromeos_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.policyDeviceAssuranceChromeos.PolicyDeviceAssuranceChromeos",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PolicyDeviceAssuranceChromeos_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-okta.policyDeviceAssuranceChromeos.PolicyDeviceAssuranceChromeos",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) ResetTpspAllowScreenLock() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspAllowScreenLock",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) ResetTpspBrowserVersion() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspBrowserVersion",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) ResetTpspBuiltinDnsClientEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspBuiltinDnsClientEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) ResetTpspChromeRemoteDesktopAppBlocked() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspChromeRemoteDesktopAppBlocked",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) ResetTpspDeviceEnrollmentDomain() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspDeviceEnrollmentDomain",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) ResetTpspDiskEncrypted() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspDiskEncrypted",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) ResetTpspKeyTrustLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspKeyTrustLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) ResetTpspOsFirewall() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspOsFirewall",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) ResetTpspOsVersion() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspOsVersion",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) ResetTpspPasswordProctectionWarningTrigger() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspPasswordProctectionWarningTrigger",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) ResetTpspRealtimeUrlCheckMode() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspRealtimeUrlCheckMode",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) ResetTpspSafeBrowsingProtectionLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspSafeBrowsingProtectionLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) ResetTpspScreenLockSecured() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspScreenLockSecured",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) ResetTpspSiteIsolationEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetTpspSiteIsolationEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyDeviceAssuranceChromeos) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

