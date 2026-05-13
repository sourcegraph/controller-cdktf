package googlebackupdrrestoreworkload

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebackupdrrestoreworkload/internal"
)

type GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference interface {
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
	EnableDisplay() interface{}
	SetEnableDisplay(val interface{})
	EnableDisplayInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDevice
	SetInternalValue(val *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDevice)
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
	ResetEnableDisplay()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference
type jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference) EnableDisplay() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDisplay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference) EnableDisplayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDisplayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference) InternalValue() *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDevice {
	var returns *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDevice
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBackupDrRestoreWorkload.GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference_Override(g GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBackupDrRestoreWorkload.GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference)SetEnableDisplay(val interface{}) {
	if err := j.validateSetEnableDisplayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDisplay",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference)SetInternalValue(val *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDevice) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference) ResetEnableDisplay() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableDisplay",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

