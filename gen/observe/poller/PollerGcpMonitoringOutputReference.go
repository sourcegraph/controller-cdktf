package poller

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/observe/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/observe/poller/internal"
)

type PollerGcpMonitoringOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExcludeMetricTypePrefixes() *[]*string
	SetExcludeMetricTypePrefixes(val *[]*string)
	ExcludeMetricTypePrefixesInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludeMetricTypePrefixes() *[]*string
	SetIncludeMetricTypePrefixes(val *[]*string)
	IncludeMetricTypePrefixesInput() *[]*string
	InternalValue() *PollerGcpMonitoring
	SetInternalValue(val *PollerGcpMonitoring)
	JsonKey() *string
	SetJsonKey(val *string)
	JsonKeyInput() *string
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
	RateLimit() *float64
	SetRateLimit(val *float64)
	RateLimitInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TotalLimit() *float64
	SetTotalLimit(val *float64)
	TotalLimitInput() *float64
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetExcludeMetricTypePrefixes()
	ResetIncludeMetricTypePrefixes()
	ResetRateLimit()
	ResetTotalLimit()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PollerGcpMonitoringOutputReference
type jsiiProxy_PollerGcpMonitoringOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PollerGcpMonitoringOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PollerGcpMonitoringOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PollerGcpMonitoringOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PollerGcpMonitoringOutputReference) ExcludeMetricTypePrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeMetricTypePrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PollerGcpMonitoringOutputReference) ExcludeMetricTypePrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeMetricTypePrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PollerGcpMonitoringOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PollerGcpMonitoringOutputReference) IncludeMetricTypePrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeMetricTypePrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PollerGcpMonitoringOutputReference) IncludeMetricTypePrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeMetricTypePrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PollerGcpMonitoringOutputReference) InternalValue() *PollerGcpMonitoring {
	var returns *PollerGcpMonitoring
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PollerGcpMonitoringOutputReference) JsonKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsonKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PollerGcpMonitoringOutputReference) JsonKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsonKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PollerGcpMonitoringOutputReference) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PollerGcpMonitoringOutputReference) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PollerGcpMonitoringOutputReference) RateLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rateLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PollerGcpMonitoringOutputReference) RateLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rateLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PollerGcpMonitoringOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PollerGcpMonitoringOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PollerGcpMonitoringOutputReference) TotalLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PollerGcpMonitoringOutputReference) TotalLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalLimitInput",
		&returns,
	)
	return returns
}


func NewPollerGcpMonitoringOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PollerGcpMonitoringOutputReference {
	_init_.Initialize()

	if err := validateNewPollerGcpMonitoringOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PollerGcpMonitoringOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-observe.poller.PollerGcpMonitoringOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPollerGcpMonitoringOutputReference_Override(p PollerGcpMonitoringOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-observe.poller.PollerGcpMonitoringOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PollerGcpMonitoringOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PollerGcpMonitoringOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PollerGcpMonitoringOutputReference)SetExcludeMetricTypePrefixes(val *[]*string) {
	if err := j.validateSetExcludeMetricTypePrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeMetricTypePrefixes",
		val,
	)
}

func (j *jsiiProxy_PollerGcpMonitoringOutputReference)SetIncludeMetricTypePrefixes(val *[]*string) {
	if err := j.validateSetIncludeMetricTypePrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeMetricTypePrefixes",
		val,
	)
}

func (j *jsiiProxy_PollerGcpMonitoringOutputReference)SetInternalValue(val *PollerGcpMonitoring) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PollerGcpMonitoringOutputReference)SetJsonKey(val *string) {
	if err := j.validateSetJsonKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jsonKey",
		val,
	)
}

func (j *jsiiProxy_PollerGcpMonitoringOutputReference)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_PollerGcpMonitoringOutputReference)SetRateLimit(val *float64) {
	if err := j.validateSetRateLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rateLimit",
		val,
	)
}

func (j *jsiiProxy_PollerGcpMonitoringOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PollerGcpMonitoringOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PollerGcpMonitoringOutputReference)SetTotalLimit(val *float64) {
	if err := j.validateSetTotalLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalLimit",
		val,
	)
}

func (p *jsiiProxy_PollerGcpMonitoringOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PollerGcpMonitoringOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PollerGcpMonitoringOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PollerGcpMonitoringOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PollerGcpMonitoringOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PollerGcpMonitoringOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PollerGcpMonitoringOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PollerGcpMonitoringOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PollerGcpMonitoringOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PollerGcpMonitoringOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PollerGcpMonitoringOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PollerGcpMonitoringOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PollerGcpMonitoringOutputReference) ResetExcludeMetricTypePrefixes() {
	_jsii_.InvokeVoid(
		p,
		"resetExcludeMetricTypePrefixes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PollerGcpMonitoringOutputReference) ResetIncludeMetricTypePrefixes() {
	_jsii_.InvokeVoid(
		p,
		"resetIncludeMetricTypePrefixes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PollerGcpMonitoringOutputReference) ResetRateLimit() {
	_jsii_.InvokeVoid(
		p,
		"resetRateLimit",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PollerGcpMonitoringOutputReference) ResetTotalLimit() {
	_jsii_.InvokeVoid(
		p,
		"resetTotalLimit",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PollerGcpMonitoringOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PollerGcpMonitoringOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

