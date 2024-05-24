package botmanagement

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/botmanagement/internal"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/bot_management cloudflare_bot_management}.
type BotManagement interface {
	cdktf.TerraformResource
	AutoUpdateModel() interface{}
	SetAutoUpdateModel(val interface{})
	AutoUpdateModelInput() interface{}
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnableJs() interface{}
	SetEnableJs(val interface{})
	EnableJsInput() interface{}
	FightMode() interface{}
	SetFightMode(val interface{})
	FightModeInput() interface{}
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
	// The tree node.
	Node() constructs.Node
	OptimizeWordpress() interface{}
	SetOptimizeWordpress(val interface{})
	OptimizeWordpressInput() interface{}
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
	SbfmDefinitelyAutomated() *string
	SetSbfmDefinitelyAutomated(val *string)
	SbfmDefinitelyAutomatedInput() *string
	SbfmLikelyAutomated() *string
	SetSbfmLikelyAutomated(val *string)
	SbfmLikelyAutomatedInput() *string
	SbfmStaticResourceProtection() interface{}
	SetSbfmStaticResourceProtection(val interface{})
	SbfmStaticResourceProtectionInput() interface{}
	SbfmVerifiedBots() *string
	SetSbfmVerifiedBots(val *string)
	SbfmVerifiedBotsInput() *string
	SuppressSessionScore() interface{}
	SetSuppressSessionScore(val interface{})
	SuppressSessionScoreInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UsingLatestModel() cdktf.IResolvable
	ZoneId() *string
	SetZoneId(val *string)
	ZoneIdInput() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAutoUpdateModel()
	ResetEnableJs()
	ResetFightMode()
	ResetId()
	ResetOptimizeWordpress()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSbfmDefinitelyAutomated()
	ResetSbfmLikelyAutomated()
	ResetSbfmStaticResourceProtection()
	ResetSbfmVerifiedBots()
	ResetSuppressSessionScore()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for BotManagement
type jsiiProxy_BotManagement struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BotManagement) AutoUpdateModel() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoUpdateModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) AutoUpdateModelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoUpdateModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) EnableJs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableJs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) EnableJsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableJsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) FightMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fightMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) FightModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fightModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) OptimizeWordpress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"optimizeWordpress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) OptimizeWordpressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"optimizeWordpressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) SbfmDefinitelyAutomated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sbfmDefinitelyAutomated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) SbfmDefinitelyAutomatedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sbfmDefinitelyAutomatedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) SbfmLikelyAutomated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sbfmLikelyAutomated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) SbfmLikelyAutomatedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sbfmLikelyAutomatedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) SbfmStaticResourceProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sbfmStaticResourceProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) SbfmStaticResourceProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sbfmStaticResourceProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) SbfmVerifiedBots() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sbfmVerifiedBots",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) SbfmVerifiedBotsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sbfmVerifiedBotsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) SuppressSessionScore() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suppressSessionScore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) SuppressSessionScoreInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suppressSessionScoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) UsingLatestModel() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"usingLatestModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotManagement) ZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/bot_management cloudflare_bot_management} Resource.
func NewBotManagement(scope constructs.Construct, id *string, config *BotManagementConfig) BotManagement {
	_init_.Initialize()

	if err := validateNewBotManagementParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BotManagement{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.botManagement.BotManagement",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/bot_management cloudflare_bot_management} Resource.
func NewBotManagement_Override(b BotManagement, scope constructs.Construct, id *string, config *BotManagementConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.botManagement.BotManagement",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BotManagement)SetAutoUpdateModel(val interface{}) {
	if err := j.validateSetAutoUpdateModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoUpdateModel",
		val,
	)
}

func (j *jsiiProxy_BotManagement)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BotManagement)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BotManagement)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BotManagement)SetEnableJs(val interface{}) {
	if err := j.validateSetEnableJsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableJs",
		val,
	)
}

func (j *jsiiProxy_BotManagement)SetFightMode(val interface{}) {
	if err := j.validateSetFightModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fightMode",
		val,
	)
}

func (j *jsiiProxy_BotManagement)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BotManagement)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_BotManagement)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BotManagement)SetOptimizeWordpress(val interface{}) {
	if err := j.validateSetOptimizeWordpressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optimizeWordpress",
		val,
	)
}

func (j *jsiiProxy_BotManagement)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BotManagement)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BotManagement)SetSbfmDefinitelyAutomated(val *string) {
	if err := j.validateSetSbfmDefinitelyAutomatedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sbfmDefinitelyAutomated",
		val,
	)
}

func (j *jsiiProxy_BotManagement)SetSbfmLikelyAutomated(val *string) {
	if err := j.validateSetSbfmLikelyAutomatedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sbfmLikelyAutomated",
		val,
	)
}

func (j *jsiiProxy_BotManagement)SetSbfmStaticResourceProtection(val interface{}) {
	if err := j.validateSetSbfmStaticResourceProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sbfmStaticResourceProtection",
		val,
	)
}

func (j *jsiiProxy_BotManagement)SetSbfmVerifiedBots(val *string) {
	if err := j.validateSetSbfmVerifiedBotsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sbfmVerifiedBots",
		val,
	)
}

func (j *jsiiProxy_BotManagement)SetSuppressSessionScore(val interface{}) {
	if err := j.validateSetSuppressSessionScoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suppressSessionScore",
		val,
	)
}

func (j *jsiiProxy_BotManagement)SetZoneId(val *string) {
	if err := j.validateSetZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneId",
		val,
	)
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
func BotManagement_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBotManagement_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.botManagement.BotManagement",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BotManagement_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBotManagement_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.botManagement.BotManagement",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BotManagement_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBotManagement_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.botManagement.BotManagement",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BotManagement_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.botManagement.BotManagement",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BotManagement) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BotManagement) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotManagement) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotManagement) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotManagement) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotManagement) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotManagement) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotManagement) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotManagement) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotManagement) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotManagement) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotManagement) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BotManagement) ResetAutoUpdateModel() {
	_jsii_.InvokeVoid(
		b,
		"resetAutoUpdateModel",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotManagement) ResetEnableJs() {
	_jsii_.InvokeVoid(
		b,
		"resetEnableJs",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotManagement) ResetFightMode() {
	_jsii_.InvokeVoid(
		b,
		"resetFightMode",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotManagement) ResetId() {
	_jsii_.InvokeVoid(
		b,
		"resetId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotManagement) ResetOptimizeWordpress() {
	_jsii_.InvokeVoid(
		b,
		"resetOptimizeWordpress",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotManagement) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotManagement) ResetSbfmDefinitelyAutomated() {
	_jsii_.InvokeVoid(
		b,
		"resetSbfmDefinitelyAutomated",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotManagement) ResetSbfmLikelyAutomated() {
	_jsii_.InvokeVoid(
		b,
		"resetSbfmLikelyAutomated",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotManagement) ResetSbfmStaticResourceProtection() {
	_jsii_.InvokeVoid(
		b,
		"resetSbfmStaticResourceProtection",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotManagement) ResetSbfmVerifiedBots() {
	_jsii_.InvokeVoid(
		b,
		"resetSbfmVerifiedBots",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotManagement) ResetSuppressSessionScore() {
	_jsii_.InvokeVoid(
		b,
		"resetSuppressSessionScore",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotManagement) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotManagement) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotManagement) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotManagement) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

