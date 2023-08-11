package googlecloudrunv2service

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudrunv2service/internal"
)

type GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference interface {
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
	CpuIdle() interface{}
	SetCpuIdle(val interface{})
	CpuIdleInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleCloudRunV2ServiceTemplateContainersResources
	SetInternalValue(val *GoogleCloudRunV2ServiceTemplateContainersResources)
	Limits() *map[string]*string
	SetLimits(val *map[string]*string)
	LimitsInput() *map[string]*string
	StartupCpuBoost() interface{}
	SetStartupCpuBoost(val interface{})
	StartupCpuBoostInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetCpuIdle()
	ResetLimits()
	ResetStartupCpuBoost()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference
type jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference) CpuIdle() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuIdle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference) CpuIdleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuIdleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference) InternalValue() *GoogleCloudRunV2ServiceTemplateContainersResources {
	var returns *GoogleCloudRunV2ServiceTemplateContainersResources
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference) Limits() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"limits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference) LimitsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"limitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference) StartupCpuBoost() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startupCpuBoost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference) StartupCpuBoostInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startupCpuBoostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCloudRunV2ServiceTemplateContainersResourcesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudRunV2Service.GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference_Override(g GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudRunV2Service.GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference)SetCpuIdle(val interface{}) {
	if err := j.validateSetCpuIdleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuIdle",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference)SetInternalValue(val *GoogleCloudRunV2ServiceTemplateContainersResources) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference)SetLimits(val *map[string]*string) {
	if err := j.validateSetLimitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limits",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference)SetStartupCpuBoost(val interface{}) {
	if err := j.validateSetStartupCpuBoostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startupCpuBoost",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference) ResetCpuIdle() {
	_jsii_.InvokeVoid(
		g,
		"resetCpuIdle",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference) ResetLimits() {
	_jsii_.InvokeVoid(
		g,
		"resetLimits",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference) ResetStartupCpuBoost() {
	_jsii_.InvokeVoid(
		g,
		"resetStartupCpuBoost",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersResourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

