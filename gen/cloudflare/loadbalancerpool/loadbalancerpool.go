package loadbalancerpool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/loadbalancerpool/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_pool cloudflare_load_balancer_pool}.
type LoadBalancerPool interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CheckRegions() *[]*string
	SetCheckRegions(val *[]*string)
	CheckRegionsInput() *[]*string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	CreatedOn() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
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
	Latitude() *float64
	SetLatitude(val *float64)
	LatitudeInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadShedding() LoadBalancerPoolLoadSheddingList
	LoadSheddingInput() interface{}
	Longitude() *float64
	SetLongitude(val *float64)
	LongitudeInput() *float64
	MinimumOrigins() *float64
	SetMinimumOrigins(val *float64)
	MinimumOriginsInput() *float64
	ModifiedOn() *string
	Monitor() *string
	SetMonitor(val *string)
	MonitorInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NotificationEmail() *string
	SetNotificationEmail(val *string)
	NotificationEmailInput() *string
	Origins() LoadBalancerPoolOriginsList
	OriginsInput() interface{}
	OriginSteering() LoadBalancerPoolOriginSteeringList
	OriginSteeringInput() interface{}
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
	PutLoadShedding(value interface{})
	PutOrigins(value interface{})
	PutOriginSteering(value interface{})
	ResetCheckRegions()
	ResetDescription()
	ResetEnabled()
	ResetId()
	ResetLatitude()
	ResetLoadShedding()
	ResetLongitude()
	ResetMinimumOrigins()
	ResetMonitor()
	ResetNotificationEmail()
	ResetOriginSteering()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for LoadBalancerPool
type jsiiProxy_LoadBalancerPool struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LoadBalancerPool) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) CheckRegions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"checkRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) CheckRegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"checkRegionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) CreatedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) Latitude() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"latitude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) LatitudeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"latitudeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) LoadShedding() LoadBalancerPoolLoadSheddingList {
	var returns LoadBalancerPoolLoadSheddingList
	_jsii_.Get(
		j,
		"loadShedding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) LoadSheddingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loadSheddingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) Longitude() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longitude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) LongitudeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longitudeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) MinimumOrigins() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumOrigins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) MinimumOriginsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumOriginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) ModifiedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) Monitor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) MonitorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) NotificationEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) NotificationEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) Origins() LoadBalancerPoolOriginsList {
	var returns LoadBalancerPoolOriginsList
	_jsii_.Get(
		j,
		"origins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) OriginsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) OriginSteering() LoadBalancerPoolOriginSteeringList {
	var returns LoadBalancerPoolOriginSteeringList
	_jsii_.Get(
		j,
		"originSteering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) OriginSteeringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originSteeringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPool) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_pool cloudflare_load_balancer_pool} Resource.
func NewLoadBalancerPool(scope constructs.Construct, id *string, config *LoadBalancerPoolConfig) LoadBalancerPool {
	_init_.Initialize()

	if err := validateNewLoadBalancerPoolParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LoadBalancerPool{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.loadBalancerPool.LoadBalancerPool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_pool cloudflare_load_balancer_pool} Resource.
func NewLoadBalancerPool_Override(l LoadBalancerPool, scope constructs.Construct, id *string, config *LoadBalancerPoolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.loadBalancerPool.LoadBalancerPool",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LoadBalancerPool)SetCheckRegions(val *[]*string) {
	if err := j.validateSetCheckRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checkRegions",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPool)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPool)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPool)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPool)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPool)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPool)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPool)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPool)SetLatitude(val *float64) {
	if err := j.validateSetLatitudeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"latitude",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPool)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPool)SetLongitude(val *float64) {
	if err := j.validateSetLongitudeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"longitude",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPool)SetMinimumOrigins(val *float64) {
	if err := j.validateSetMinimumOriginsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumOrigins",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPool)SetMonitor(val *string) {
	if err := j.validateSetMonitorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitor",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPool)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPool)SetNotificationEmail(val *string) {
	if err := j.validateSetNotificationEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationEmail",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPool)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPool)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
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
func LoadBalancerPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLoadBalancerPool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.loadBalancerPool.LoadBalancerPool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LoadBalancerPool_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLoadBalancerPool_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.loadBalancerPool.LoadBalancerPool",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LoadBalancerPool_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLoadBalancerPool_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.loadBalancerPool.LoadBalancerPool",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LoadBalancerPool_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.loadBalancerPool.LoadBalancerPool",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LoadBalancerPool) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LoadBalancerPool) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerPool) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerPool) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerPool) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerPool) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerPool) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerPool) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerPool) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerPool) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerPool) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerPool) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LoadBalancerPool) PutLoadShedding(value interface{}) {
	if err := l.validatePutLoadSheddingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putLoadShedding",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LoadBalancerPool) PutOrigins(value interface{}) {
	if err := l.validatePutOriginsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putOrigins",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LoadBalancerPool) PutOriginSteering(value interface{}) {
	if err := l.validatePutOriginSteeringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putOriginSteering",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LoadBalancerPool) ResetCheckRegions() {
	_jsii_.InvokeVoid(
		l,
		"resetCheckRegions",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerPool) ResetDescription() {
	_jsii_.InvokeVoid(
		l,
		"resetDescription",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerPool) ResetEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerPool) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerPool) ResetLatitude() {
	_jsii_.InvokeVoid(
		l,
		"resetLatitude",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerPool) ResetLoadShedding() {
	_jsii_.InvokeVoid(
		l,
		"resetLoadShedding",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerPool) ResetLongitude() {
	_jsii_.InvokeVoid(
		l,
		"resetLongitude",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerPool) ResetMinimumOrigins() {
	_jsii_.InvokeVoid(
		l,
		"resetMinimumOrigins",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerPool) ResetMonitor() {
	_jsii_.InvokeVoid(
		l,
		"resetMonitor",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerPool) ResetNotificationEmail() {
	_jsii_.InvokeVoid(
		l,
		"resetNotificationEmail",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerPool) ResetOriginSteering() {
	_jsii_.InvokeVoid(
		l,
		"resetOriginSteering",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerPool) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerPool) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerPool) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerPool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerPool) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

