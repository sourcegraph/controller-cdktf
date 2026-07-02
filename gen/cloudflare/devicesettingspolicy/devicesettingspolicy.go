package devicesettingspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/devicesettingspolicy/internal"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_settings_policy cloudflare_device_settings_policy}.
type DeviceSettingsPolicy interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	AllowedToLeave() any
	SetAllowedToLeave(val any)
	AllowedToLeaveInput() any
	AllowModeSwitch() any
	SetAllowModeSwitch(val any)
	AllowModeSwitchInput() any
	AllowUpdates() any
	SetAllowUpdates(val any)
	AllowUpdatesInput() any
	AutoConnect() *float64
	SetAutoConnect(val *float64)
	AutoConnectInput() *float64
	CaptivePortal() *float64
	SetCaptivePortal(val *float64)
	CaptivePortalInput() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() any
	// Experimental.
	SetConnection(val any)
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	// Experimental.
	Count() any
	// Experimental.
	SetCount(val any)
	Default() any
	SetDefault(val any)
	DefaultInput() any
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisableAutoFallback() any
	SetDisableAutoFallback(val any)
	DisableAutoFallbackInput() any
	Enabled() any
	SetEnabled(val any)
	EnabledInput() any
	ExcludeOfficeIps() any
	SetExcludeOfficeIps(val any)
	ExcludeOfficeIpsInput() any
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Match() *string
	SetMatch(val *string)
	MatchInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Precedence() *float64
	SetPrecedence(val *float64)
	PrecedenceInput() *float64
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]any
	// Experimental.
	SetProvisioners(val *[]any)
	// Experimental.
	RawOverrides() any
	ServiceModeV2Mode() *string
	SetServiceModeV2Mode(val *string)
	ServiceModeV2ModeInput() *string
	ServiceModeV2Port() *float64
	SetServiceModeV2Port(val *float64)
	ServiceModeV2PortInput() *float64
	SupportUrl() *string
	SetSupportUrl(val *string)
	SupportUrlInput() *string
	SwitchLocked() any
	SetSwitchLocked(val any)
	SwitchLockedInput() any
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value any)
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	HasResourceMove() any
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
	MoveTo(moveTarget *string, index any)
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAllowedToLeave()
	ResetAllowModeSwitch()
	ResetAllowUpdates()
	ResetAutoConnect()
	ResetCaptivePortal()
	ResetDefault()
	ResetDisableAutoFallback()
	ResetEnabled()
	ResetExcludeOfficeIps()
	ResetId()
	ResetMatch()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrecedence()
	ResetServiceModeV2Mode()
	ResetServiceModeV2Port()
	ResetSupportUrl()
	ResetSwitchLocked()
	SynthesizeAttributes() *map[string]any
	SynthesizeHclAttributes() *map[string]any
	// Experimental.
	ToHclTerraform() any
	// Experimental.
	ToMetadata() any
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() any
}

// The jsii proxy struct for DeviceSettingsPolicy
type jsiiProxy_DeviceSettingsPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DeviceSettingsPolicy) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) AllowedToLeave() any {
	var returns any
	_jsii_.Get(
		j,
		"allowedToLeave",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) AllowedToLeaveInput() any {
	var returns any
	_jsii_.Get(
		j,
		"allowedToLeaveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) AllowModeSwitch() any {
	var returns any
	_jsii_.Get(
		j,
		"allowModeSwitch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) AllowModeSwitchInput() any {
	var returns any
	_jsii_.Get(
		j,
		"allowModeSwitchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) AllowUpdates() any {
	var returns any
	_jsii_.Get(
		j,
		"allowUpdates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) AllowUpdatesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"allowUpdatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) AutoConnect() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoConnect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) AutoConnectInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoConnectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) CaptivePortal() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"captivePortal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) CaptivePortalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"captivePortalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) Default() any {
	var returns any
	_jsii_.Get(
		j,
		"default",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) DefaultInput() any {
	var returns any
	_jsii_.Get(
		j,
		"defaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) DisableAutoFallback() any {
	var returns any
	_jsii_.Get(
		j,
		"disableAutoFallback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) DisableAutoFallbackInput() any {
	var returns any
	_jsii_.Get(
		j,
		"disableAutoFallbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) Enabled() any {
	var returns any
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) EnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) ExcludeOfficeIps() any {
	var returns any
	_jsii_.Get(
		j,
		"excludeOfficeIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) ExcludeOfficeIpsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"excludeOfficeIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) Match() *string {
	var returns *string
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) MatchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) Precedence() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"precedence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) PrecedenceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"precedenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) ServiceModeV2Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceModeV2Mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) ServiceModeV2ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceModeV2ModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) ServiceModeV2Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serviceModeV2Port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) ServiceModeV2PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serviceModeV2PortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) SupportUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) SupportUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) SwitchLocked() any {
	var returns any
	_jsii_.Get(
		j,
		"switchLocked",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) SwitchLockedInput() any {
	var returns any
	_jsii_.Get(
		j,
		"switchLockedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeviceSettingsPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_settings_policy cloudflare_device_settings_policy} Resource.
func NewDeviceSettingsPolicy(scope constructs.Construct, id *string, config *DeviceSettingsPolicyConfig) DeviceSettingsPolicy {
	_init_.Initialize()

	if err := validateNewDeviceSettingsPolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeviceSettingsPolicy{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.deviceSettingsPolicy.DeviceSettingsPolicy",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_settings_policy cloudflare_device_settings_policy} Resource.
func NewDeviceSettingsPolicy_Override(d DeviceSettingsPolicy, scope constructs.Construct, id *string, config *DeviceSettingsPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.deviceSettingsPolicy.DeviceSettingsPolicy",
		[]any{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DeviceSettingsPolicy) SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_DeviceSettingsPolicy) SetAllowedToLeave(val any) {
	if err := j.validateSetAllowedToLeaveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedToLeave",
		val,
	)
}

func (j *jsiiProxy_DeviceSettingsPolicy) SetAllowModeSwitch(val any) {
	if err := j.validateSetAllowModeSwitchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowModeSwitch",
		val,
	)
}

func (j *jsiiProxy_DeviceSettingsPolicy) SetAllowUpdates(val any) {
	if err := j.validateSetAllowUpdatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowUpdates",
		val,
	)
}

func (j *jsiiProxy_DeviceSettingsPolicy) SetAutoConnect(val *float64) {
	if err := j.validateSetAutoConnectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoConnect",
		val,
	)
}

func (j *jsiiProxy_DeviceSettingsPolicy) SetCaptivePortal(val *float64) {
	if err := j.validateSetCaptivePortalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"captivePortal",
		val,
	)
}

func (j *jsiiProxy_DeviceSettingsPolicy) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DeviceSettingsPolicy) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DeviceSettingsPolicy) SetDefault(val any) {
	if err := j.validateSetDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"default",
		val,
	)
}

func (j *jsiiProxy_DeviceSettingsPolicy) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DeviceSettingsPolicy) SetDisableAutoFallback(val any) {
	if err := j.validateSetDisableAutoFallbackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableAutoFallback",
		val,
	)
}

func (j *jsiiProxy_DeviceSettingsPolicy) SetEnabled(val any) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_DeviceSettingsPolicy) SetExcludeOfficeIps(val any) {
	if err := j.validateSetExcludeOfficeIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeOfficeIps",
		val,
	)
}

func (j *jsiiProxy_DeviceSettingsPolicy) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DeviceSettingsPolicy) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DeviceSettingsPolicy) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DeviceSettingsPolicy) SetMatch(val *string) {
	if err := j.validateSetMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"match",
		val,
	)
}

func (j *jsiiProxy_DeviceSettingsPolicy) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DeviceSettingsPolicy) SetPrecedence(val *float64) {
	if err := j.validateSetPrecedenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"precedence",
		val,
	)
}

func (j *jsiiProxy_DeviceSettingsPolicy) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DeviceSettingsPolicy) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DeviceSettingsPolicy) SetServiceModeV2Mode(val *string) {
	if err := j.validateSetServiceModeV2ModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceModeV2Mode",
		val,
	)
}

func (j *jsiiProxy_DeviceSettingsPolicy) SetServiceModeV2Port(val *float64) {
	if err := j.validateSetServiceModeV2PortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceModeV2Port",
		val,
	)
}

func (j *jsiiProxy_DeviceSettingsPolicy) SetSupportUrl(val *string) {
	if err := j.validateSetSupportUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportUrl",
		val,
	)
}

func (j *jsiiProxy_DeviceSettingsPolicy) SetSwitchLocked(val any) {
	if err := j.validateSetSwitchLockedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"switchLocked",
		val,
	)
}

// Generates CDKTF code for importing a DeviceSettingsPolicy resource upon running "cdktf plan <stack-name>".
func DeviceSettingsPolicy_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDeviceSettingsPolicy_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.deviceSettingsPolicy.DeviceSettingsPolicy",
		"generateConfigForImport",
		[]any{scope, importToId, importFromId, provider},
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
func DeviceSettingsPolicy_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateDeviceSettingsPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.deviceSettingsPolicy.DeviceSettingsPolicy",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func DeviceSettingsPolicy_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateDeviceSettingsPolicy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.deviceSettingsPolicy.DeviceSettingsPolicy",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func DeviceSettingsPolicy_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateDeviceSettingsPolicy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.deviceSettingsPolicy.DeviceSettingsPolicy",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func DeviceSettingsPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.deviceSettingsPolicy.DeviceSettingsPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DeviceSettingsPolicy) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (d *jsiiProxy_DeviceSettingsPolicy) AddOverride(path *string, value any) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]any{path, value},
	)
}

func (d *jsiiProxy_DeviceSettingsPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeviceSettingsPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeviceSettingsPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeviceSettingsPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeviceSettingsPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeviceSettingsPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeviceSettingsPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeviceSettingsPolicy) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeviceSettingsPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeviceSettingsPolicy) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeviceSettingsPolicy) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]any{id, provider},
	)
}

func (d *jsiiProxy_DeviceSettingsPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeviceSettingsPolicy) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]any{id},
	)
}

func (d *jsiiProxy_DeviceSettingsPolicy) MoveTo(moveTarget *string, index any) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (d *jsiiProxy_DeviceSettingsPolicy) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]any{id},
	)
}

func (d *jsiiProxy_DeviceSettingsPolicy) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (d *jsiiProxy_DeviceSettingsPolicy) ResetAllowedToLeave() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowedToLeave",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeviceSettingsPolicy) ResetAllowModeSwitch() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowModeSwitch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeviceSettingsPolicy) ResetAllowUpdates() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowUpdates",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeviceSettingsPolicy) ResetAutoConnect() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoConnect",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeviceSettingsPolicy) ResetCaptivePortal() {
	_jsii_.InvokeVoid(
		d,
		"resetCaptivePortal",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeviceSettingsPolicy) ResetDefault() {
	_jsii_.InvokeVoid(
		d,
		"resetDefault",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeviceSettingsPolicy) ResetDisableAutoFallback() {
	_jsii_.InvokeVoid(
		d,
		"resetDisableAutoFallback",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeviceSettingsPolicy) ResetEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeviceSettingsPolicy) ResetExcludeOfficeIps() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludeOfficeIps",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeviceSettingsPolicy) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeviceSettingsPolicy) ResetMatch() {
	_jsii_.InvokeVoid(
		d,
		"resetMatch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeviceSettingsPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeviceSettingsPolicy) ResetPrecedence() {
	_jsii_.InvokeVoid(
		d,
		"resetPrecedence",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeviceSettingsPolicy) ResetServiceModeV2Mode() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceModeV2Mode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeviceSettingsPolicy) ResetServiceModeV2Port() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceModeV2Port",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeviceSettingsPolicy) ResetSupportUrl() {
	_jsii_.InvokeVoid(
		d,
		"resetSupportUrl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeviceSettingsPolicy) ResetSwitchLocked() {
	_jsii_.InvokeVoid(
		d,
		"resetSwitchLocked",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeviceSettingsPolicy) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeviceSettingsPolicy) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeviceSettingsPolicy) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeviceSettingsPolicy) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeviceSettingsPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeviceSettingsPolicy) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
