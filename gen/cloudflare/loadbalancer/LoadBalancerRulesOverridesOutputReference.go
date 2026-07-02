package loadbalancer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/loadbalancer/internal"
)

type LoadBalancerRulesOverridesOutputReference interface {
	cdktf.ComplexObject
	AdaptiveRouting() LoadBalancerRulesOverridesAdaptiveRoutingList
	AdaptiveRoutingInput() any
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	CountryPools() LoadBalancerRulesOverridesCountryPoolsList
	CountryPoolsInput() any
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultPools() *[]*string
	SetDefaultPools(val *[]*string)
	DefaultPoolsInput() *[]*string
	FallbackPool() *string
	SetFallbackPool(val *string)
	FallbackPoolInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() any
	SetInternalValue(val any)
	LocationStrategy() LoadBalancerRulesOverridesLocationStrategyList
	LocationStrategyInput() any
	PopPools() LoadBalancerRulesOverridesPopPoolsList
	PopPoolsInput() any
	RandomSteering() LoadBalancerRulesOverridesRandomSteeringList
	RandomSteeringInput() any
	RegionPools() LoadBalancerRulesOverridesRegionPoolsList
	RegionPoolsInput() any
	SessionAffinity() *string
	SetSessionAffinity(val *string)
	SessionAffinityAttributes() LoadBalancerRulesOverridesSessionAffinityAttributesList
	SessionAffinityAttributesInput() any
	SessionAffinityInput() *string
	SessionAffinityTtl() *float64
	SetSessionAffinityTtl(val *float64)
	SessionAffinityTtlInput() *float64
	SteeringPolicy() *string
	SetSteeringPolicy(val *string)
	SteeringPolicyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Ttl() *float64
	SetTtl(val *float64)
	TtlInput() *float64
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutAdaptiveRouting(value any)
	PutCountryPools(value any)
	PutLocationStrategy(value any)
	PutPopPools(value any)
	PutRandomSteering(value any)
	PutRegionPools(value any)
	PutSessionAffinityAttributes(value any)
	ResetAdaptiveRouting()
	ResetCountryPools()
	ResetDefaultPools()
	ResetFallbackPool()
	ResetLocationStrategy()
	ResetPopPools()
	ResetRandomSteering()
	ResetRegionPools()
	ResetSessionAffinity()
	ResetSessionAffinityAttributes()
	ResetSessionAffinityTtl()
	ResetSteeringPolicy()
	ResetTtl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LoadBalancerRulesOverridesOutputReference
type jsiiProxy_LoadBalancerRulesOverridesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) AdaptiveRouting() LoadBalancerRulesOverridesAdaptiveRoutingList {
	var returns LoadBalancerRulesOverridesAdaptiveRoutingList
	_jsii_.Get(
		j,
		"adaptiveRouting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) AdaptiveRoutingInput() any {
	var returns any
	_jsii_.Get(
		j,
		"adaptiveRoutingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) CountryPools() LoadBalancerRulesOverridesCountryPoolsList {
	var returns LoadBalancerRulesOverridesCountryPoolsList
	_jsii_.Get(
		j,
		"countryPools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) CountryPoolsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"countryPoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) DefaultPools() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultPools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) DefaultPoolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultPoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) FallbackPool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fallbackPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) FallbackPoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fallbackPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) LocationStrategy() LoadBalancerRulesOverridesLocationStrategyList {
	var returns LoadBalancerRulesOverridesLocationStrategyList
	_jsii_.Get(
		j,
		"locationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) LocationStrategyInput() any {
	var returns any
	_jsii_.Get(
		j,
		"locationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) PopPools() LoadBalancerRulesOverridesPopPoolsList {
	var returns LoadBalancerRulesOverridesPopPoolsList
	_jsii_.Get(
		j,
		"popPools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) PopPoolsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"popPoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) RandomSteering() LoadBalancerRulesOverridesRandomSteeringList {
	var returns LoadBalancerRulesOverridesRandomSteeringList
	_jsii_.Get(
		j,
		"randomSteering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) RandomSteeringInput() any {
	var returns any
	_jsii_.Get(
		j,
		"randomSteeringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) RegionPools() LoadBalancerRulesOverridesRegionPoolsList {
	var returns LoadBalancerRulesOverridesRegionPoolsList
	_jsii_.Get(
		j,
		"regionPools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) RegionPoolsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"regionPoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) SessionAffinity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) SessionAffinityAttributes() LoadBalancerRulesOverridesSessionAffinityAttributesList {
	var returns LoadBalancerRulesOverridesSessionAffinityAttributesList
	_jsii_.Get(
		j,
		"sessionAffinityAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) SessionAffinityAttributesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"sessionAffinityAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) SessionAffinityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) SessionAffinityTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionAffinityTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) SessionAffinityTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionAffinityTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) SteeringPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"steeringPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) SteeringPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"steeringPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) Ttl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) TtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}

func NewLoadBalancerRulesOverridesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LoadBalancerRulesOverridesOutputReference {
	_init_.Initialize()

	if err := validateNewLoadBalancerRulesOverridesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LoadBalancerRulesOverridesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.loadBalancer.LoadBalancerRulesOverridesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLoadBalancerRulesOverridesOutputReference_Override(l LoadBalancerRulesOverridesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.loadBalancer.LoadBalancerRulesOverridesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) SetDefaultPools(val *[]*string) {
	if err := j.validateSetDefaultPoolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultPools",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) SetFallbackPool(val *string) {
	if err := j.validateSetFallbackPoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fallbackPool",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) SetSessionAffinity(val *string) {
	if err := j.validateSetSessionAffinityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionAffinity",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) SetSessionAffinityTtl(val *float64) {
	if err := j.validateSetSessionAffinityTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionAffinityTtl",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) SetSteeringPolicy(val *string) {
	if err := j.validateSetSteeringPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"steeringPolicy",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerRulesOverridesOutputReference) SetTtl(val *float64) {
	if err := j.validateSetTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) PutAdaptiveRouting(value any) {
	if err := l.validatePutAdaptiveRoutingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAdaptiveRouting",
		[]any{value},
	)
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) PutCountryPools(value any) {
	if err := l.validatePutCountryPoolsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putCountryPools",
		[]any{value},
	)
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) PutLocationStrategy(value any) {
	if err := l.validatePutLocationStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putLocationStrategy",
		[]any{value},
	)
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) PutPopPools(value any) {
	if err := l.validatePutPopPoolsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putPopPools",
		[]any{value},
	)
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) PutRandomSteering(value any) {
	if err := l.validatePutRandomSteeringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putRandomSteering",
		[]any{value},
	)
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) PutRegionPools(value any) {
	if err := l.validatePutRegionPoolsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putRegionPools",
		[]any{value},
	)
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) PutSessionAffinityAttributes(value any) {
	if err := l.validatePutSessionAffinityAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSessionAffinityAttributes",
		[]any{value},
	)
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) ResetAdaptiveRouting() {
	_jsii_.InvokeVoid(
		l,
		"resetAdaptiveRouting",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) ResetCountryPools() {
	_jsii_.InvokeVoid(
		l,
		"resetCountryPools",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) ResetDefaultPools() {
	_jsii_.InvokeVoid(
		l,
		"resetDefaultPools",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) ResetFallbackPool() {
	_jsii_.InvokeVoid(
		l,
		"resetFallbackPool",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) ResetLocationStrategy() {
	_jsii_.InvokeVoid(
		l,
		"resetLocationStrategy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) ResetPopPools() {
	_jsii_.InvokeVoid(
		l,
		"resetPopPools",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) ResetRandomSteering() {
	_jsii_.InvokeVoid(
		l,
		"resetRandomSteering",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) ResetRegionPools() {
	_jsii_.InvokeVoid(
		l,
		"resetRegionPools",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) ResetSessionAffinity() {
	_jsii_.InvokeVoid(
		l,
		"resetSessionAffinity",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) ResetSessionAffinityAttributes() {
	_jsii_.InvokeVoid(
		l,
		"resetSessionAffinityAttributes",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) ResetSessionAffinityTtl() {
	_jsii_.InvokeVoid(
		l,
		"resetSessionAffinityTtl",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) ResetSteeringPolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetSteeringPolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) ResetTtl() {
	_jsii_.InvokeVoid(
		l,
		"resetTtl",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		l,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerRulesOverridesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
