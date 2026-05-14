package googlelustreinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlelustreinstance/internal"
)

type GoogleLustreInstanceMaintenancePolicyOutputReference interface {
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
	InternalValue() *GoogleLustreInstanceMaintenancePolicy
	SetInternalValue(val *GoogleLustreInstanceMaintenancePolicy)
	MaintenanceExclusionWindow() GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference
	MaintenanceExclusionWindowInput() *GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindow
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WeeklyMaintenanceWindows() GoogleLustreInstanceMaintenancePolicyWeeklyMaintenanceWindowsOutputReference
	WeeklyMaintenanceWindowsInput() *GoogleLustreInstanceMaintenancePolicyWeeklyMaintenanceWindows
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
	PutMaintenanceExclusionWindow(value *GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindow)
	PutWeeklyMaintenanceWindows(value *GoogleLustreInstanceMaintenancePolicyWeeklyMaintenanceWindows)
	ResetMaintenanceExclusionWindow()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleLustreInstanceMaintenancePolicyOutputReference
type jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference) InternalValue() *GoogleLustreInstanceMaintenancePolicy {
	var returns *GoogleLustreInstanceMaintenancePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference) MaintenanceExclusionWindow() GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference {
	var returns GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference
	_jsii_.Get(
		j,
		"maintenanceExclusionWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference) MaintenanceExclusionWindowInput() *GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindow {
	var returns *GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindow
	_jsii_.Get(
		j,
		"maintenanceExclusionWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference) WeeklyMaintenanceWindows() GoogleLustreInstanceMaintenancePolicyWeeklyMaintenanceWindowsOutputReference {
	var returns GoogleLustreInstanceMaintenancePolicyWeeklyMaintenanceWindowsOutputReference
	_jsii_.Get(
		j,
		"weeklyMaintenanceWindows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference) WeeklyMaintenanceWindowsInput() *GoogleLustreInstanceMaintenancePolicyWeeklyMaintenanceWindows {
	var returns *GoogleLustreInstanceMaintenancePolicyWeeklyMaintenanceWindows
	_jsii_.Get(
		j,
		"weeklyMaintenanceWindowsInput",
		&returns,
	)
	return returns
}


func NewGoogleLustreInstanceMaintenancePolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleLustreInstanceMaintenancePolicyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleLustreInstanceMaintenancePolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleLustreInstance.GoogleLustreInstanceMaintenancePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleLustreInstanceMaintenancePolicyOutputReference_Override(g GoogleLustreInstanceMaintenancePolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleLustreInstance.GoogleLustreInstanceMaintenancePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference)SetInternalValue(val *GoogleLustreInstanceMaintenancePolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference) PutMaintenanceExclusionWindow(value *GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindow) {
	if err := g.validatePutMaintenanceExclusionWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMaintenanceExclusionWindow",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference) PutWeeklyMaintenanceWindows(value *GoogleLustreInstanceMaintenancePolicyWeeklyMaintenanceWindows) {
	if err := g.validatePutWeeklyMaintenanceWindowsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWeeklyMaintenanceWindows",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference) ResetMaintenanceExclusionWindow() {
	_jsii_.InvokeVoid(
		g,
		"resetMaintenanceExclusionWindow",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleLustreInstanceMaintenancePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

