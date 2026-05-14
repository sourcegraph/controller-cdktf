package googlecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecontainercluster/internal"
)

type GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleContainerClusterMaintenancePolicyDisruptionBudget
	SetInternalValue(val *GoogleContainerClusterMaintenancePolicyDisruptionBudget)
	LastDisruptionTime() *string
	LastMinorVersionDisruptionTime() *string
	MinorVersionDisruptionInterval() *string
	SetMinorVersionDisruptionInterval(val *string)
	MinorVersionDisruptionIntervalInput() *string
	PatchVersionDisruptionInterval() *string
	SetPatchVersionDisruptionInterval(val *string)
	PatchVersionDisruptionIntervalInput() *string
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
	ResetMinorVersionDisruptionInterval()
	ResetPatchVersionDisruptionInterval()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference
type jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference) InternalValue() *GoogleContainerClusterMaintenancePolicyDisruptionBudget {
	var returns *GoogleContainerClusterMaintenancePolicyDisruptionBudget
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference) LastDisruptionTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastDisruptionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference) LastMinorVersionDisruptionTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastMinorVersionDisruptionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference) MinorVersionDisruptionInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minorVersionDisruptionInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference) MinorVersionDisruptionIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minorVersionDisruptionIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference) PatchVersionDisruptionInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchVersionDisruptionInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference) PatchVersionDisruptionIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchVersionDisruptionIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleContainerCluster.GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference_Override(g GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleContainerCluster.GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference)SetInternalValue(val *GoogleContainerClusterMaintenancePolicyDisruptionBudget) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference)SetMinorVersionDisruptionInterval(val *string) {
	if err := j.validateSetMinorVersionDisruptionIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minorVersionDisruptionInterval",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference)SetPatchVersionDisruptionInterval(val *string) {
	if err := j.validateSetPatchVersionDisruptionIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"patchVersionDisruptionInterval",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference) ResetMinorVersionDisruptionInterval() {
	_jsii_.InvokeVoid(
		g,
		"resetMinorVersionDisruptionInterval",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference) ResetPatchVersionDisruptionInterval() {
	_jsii_.InvokeVoid(
		g,
		"resetPatchVersionDisruptionInterval",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterMaintenancePolicyDisruptionBudgetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

