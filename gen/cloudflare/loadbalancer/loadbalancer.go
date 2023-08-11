package loadbalancer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/loadbalancer/internal"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer cloudflare_load_balancer}.
type LoadBalancer interface {
	cdktf.TerraformResource
	AdaptiveRouting() LoadBalancerAdaptiveRoutingList
	AdaptiveRoutingInput() interface{}
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
	CountryPools() LoadBalancerCountryPoolsList
	CountryPoolsInput() interface{}
	CreatedOn() *string
	DefaultPoolIds() *[]*string
	SetDefaultPoolIds(val *[]*string)
	DefaultPoolIdsInput() *[]*string
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
	FallbackPoolId() *string
	SetFallbackPoolId(val *string)
	FallbackPoolIdInput() *string
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
	LocationStrategy() LoadBalancerLocationStrategyList
	LocationStrategyInput() interface{}
	ModifiedOn() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PopPools() LoadBalancerPopPoolsList
	PopPoolsInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	Proxied() interface{}
	SetProxied(val interface{})
	ProxiedInput() interface{}
	RandomSteering() LoadBalancerRandomSteeringList
	RandomSteeringInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	RegionPools() LoadBalancerRegionPoolsList
	RegionPoolsInput() interface{}
	Rules() LoadBalancerRulesList
	RulesInput() interface{}
	SessionAffinity() *string
	SetSessionAffinity(val *string)
	SessionAffinityAttributes() LoadBalancerSessionAffinityAttributesList
	SessionAffinityAttributesInput() interface{}
	SessionAffinityInput() *string
	SessionAffinityTtl() *float64
	SetSessionAffinityTtl(val *float64)
	SessionAffinityTtlInput() *float64
	SteeringPolicy() *string
	SetSteeringPolicy(val *string)
	SteeringPolicyInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Ttl() *float64
	SetTtl(val *float64)
	TtlInput() *float64
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
	PutAdaptiveRouting(value interface{})
	PutCountryPools(value interface{})
	PutLocationStrategy(value interface{})
	PutPopPools(value interface{})
	PutRandomSteering(value interface{})
	PutRegionPools(value interface{})
	PutRules(value interface{})
	PutSessionAffinityAttributes(value interface{})
	ResetAdaptiveRouting()
	ResetCountryPools()
	ResetDescription()
	ResetEnabled()
	ResetId()
	ResetLocationStrategy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPopPools()
	ResetProxied()
	ResetRandomSteering()
	ResetRegionPools()
	ResetRules()
	ResetSessionAffinity()
	ResetSessionAffinityAttributes()
	ResetSessionAffinityTtl()
	ResetSteeringPolicy()
	ResetTtl()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for LoadBalancer
type jsiiProxy_LoadBalancer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LoadBalancer) AdaptiveRouting() LoadBalancerAdaptiveRoutingList {
	var returns LoadBalancerAdaptiveRoutingList
	_jsii_.Get(
		j,
		"adaptiveRouting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) AdaptiveRoutingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adaptiveRoutingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) CountryPools() LoadBalancerCountryPoolsList {
	var returns LoadBalancerCountryPoolsList
	_jsii_.Get(
		j,
		"countryPools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) CountryPoolsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"countryPoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) CreatedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) DefaultPoolIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultPoolIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) DefaultPoolIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultPoolIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) FallbackPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fallbackPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) FallbackPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fallbackPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) LocationStrategy() LoadBalancerLocationStrategyList {
	var returns LoadBalancerLocationStrategyList
	_jsii_.Get(
		j,
		"locationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) LocationStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"locationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) ModifiedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) PopPools() LoadBalancerPopPoolsList {
	var returns LoadBalancerPopPoolsList
	_jsii_.Get(
		j,
		"popPools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) PopPoolsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"popPoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Proxied() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proxied",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) ProxiedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proxiedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) RandomSteering() LoadBalancerRandomSteeringList {
	var returns LoadBalancerRandomSteeringList
	_jsii_.Get(
		j,
		"randomSteering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) RandomSteeringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"randomSteeringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) RegionPools() LoadBalancerRegionPoolsList {
	var returns LoadBalancerRegionPoolsList
	_jsii_.Get(
		j,
		"regionPools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) RegionPoolsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regionPoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Rules() LoadBalancerRulesList {
	var returns LoadBalancerRulesList
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) RulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) SessionAffinity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) SessionAffinityAttributes() LoadBalancerSessionAffinityAttributesList {
	var returns LoadBalancerSessionAffinityAttributesList
	_jsii_.Get(
		j,
		"sessionAffinityAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) SessionAffinityAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sessionAffinityAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) SessionAffinityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) SessionAffinityTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionAffinityTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) SessionAffinityTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionAffinityTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) SteeringPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"steeringPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) SteeringPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"steeringPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Ttl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) TtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) ZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer cloudflare_load_balancer} Resource.
func NewLoadBalancer(scope constructs.Construct, id *string, config *LoadBalancerConfig) LoadBalancer {
	_init_.Initialize()

	if err := validateNewLoadBalancerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LoadBalancer{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.loadBalancer.LoadBalancer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer cloudflare_load_balancer} Resource.
func NewLoadBalancer_Override(l LoadBalancer, scope constructs.Construct, id *string, config *LoadBalancerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.loadBalancer.LoadBalancer",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LoadBalancer)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer)SetDefaultPoolIds(val *[]*string) {
	if err := j.validateSetDefaultPoolIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultPoolIds",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer)SetFallbackPoolId(val *string) {
	if err := j.validateSetFallbackPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fallbackPoolId",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer)SetProxied(val interface{}) {
	if err := j.validateSetProxiedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxied",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer)SetSessionAffinity(val *string) {
	if err := j.validateSetSessionAffinityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionAffinity",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer)SetSessionAffinityTtl(val *float64) {
	if err := j.validateSetSessionAffinityTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionAffinityTtl",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer)SetSteeringPolicy(val *string) {
	if err := j.validateSetSteeringPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"steeringPolicy",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer)SetTtl(val *float64) {
	if err := j.validateSetTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer)SetZoneId(val *string) {
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
func LoadBalancer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLoadBalancer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.loadBalancer.LoadBalancer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LoadBalancer_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLoadBalancer_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.loadBalancer.LoadBalancer",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LoadBalancer_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLoadBalancer_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.loadBalancer.LoadBalancer",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LoadBalancer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.loadBalancer.LoadBalancer",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LoadBalancer) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LoadBalancer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LoadBalancer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LoadBalancer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LoadBalancer) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LoadBalancer) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LoadBalancer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LoadBalancer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LoadBalancer) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LoadBalancer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LoadBalancer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LoadBalancer) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LoadBalancer) PutAdaptiveRouting(value interface{}) {
	if err := l.validatePutAdaptiveRoutingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAdaptiveRouting",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LoadBalancer) PutCountryPools(value interface{}) {
	if err := l.validatePutCountryPoolsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putCountryPools",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LoadBalancer) PutLocationStrategy(value interface{}) {
	if err := l.validatePutLocationStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putLocationStrategy",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LoadBalancer) PutPopPools(value interface{}) {
	if err := l.validatePutPopPoolsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putPopPools",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LoadBalancer) PutRandomSteering(value interface{}) {
	if err := l.validatePutRandomSteeringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putRandomSteering",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LoadBalancer) PutRegionPools(value interface{}) {
	if err := l.validatePutRegionPoolsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putRegionPools",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LoadBalancer) PutRules(value interface{}) {
	if err := l.validatePutRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putRules",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LoadBalancer) PutSessionAffinityAttributes(value interface{}) {
	if err := l.validatePutSessionAffinityAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSessionAffinityAttributes",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LoadBalancer) ResetAdaptiveRouting() {
	_jsii_.InvokeVoid(
		l,
		"resetAdaptiveRouting",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) ResetCountryPools() {
	_jsii_.InvokeVoid(
		l,
		"resetCountryPools",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) ResetDescription() {
	_jsii_.InvokeVoid(
		l,
		"resetDescription",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) ResetEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) ResetLocationStrategy() {
	_jsii_.InvokeVoid(
		l,
		"resetLocationStrategy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) ResetPopPools() {
	_jsii_.InvokeVoid(
		l,
		"resetPopPools",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) ResetProxied() {
	_jsii_.InvokeVoid(
		l,
		"resetProxied",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) ResetRandomSteering() {
	_jsii_.InvokeVoid(
		l,
		"resetRandomSteering",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) ResetRegionPools() {
	_jsii_.InvokeVoid(
		l,
		"resetRegionPools",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) ResetRules() {
	_jsii_.InvokeVoid(
		l,
		"resetRules",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) ResetSessionAffinity() {
	_jsii_.InvokeVoid(
		l,
		"resetSessionAffinity",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) ResetSessionAffinityAttributes() {
	_jsii_.InvokeVoid(
		l,
		"resetSessionAffinityAttributes",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) ResetSessionAffinityTtl() {
	_jsii_.InvokeVoid(
		l,
		"resetSessionAffinityTtl",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) ResetSteeringPolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetSteeringPolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) ResetTtl() {
	_jsii_.InvokeVoid(
		l,
		"resetTtl",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

