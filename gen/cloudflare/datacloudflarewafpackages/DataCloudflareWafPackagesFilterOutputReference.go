package datacloudflarewafpackages

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/datacloudflarewafpackages/internal"
)

type DataCloudflareWafPackagesFilterOutputReference interface {
	cdktf.ComplexObject
	ActionMode() *string
	SetActionMode(val *string)
	ActionModeInput() *string
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
	DetectionMode() *string
	SetDetectionMode(val *string)
	DetectionModeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataCloudflareWafPackagesFilter
	SetInternalValue(val *DataCloudflareWafPackagesFilter)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Sensitivity() *string
	SetSensitivity(val *string)
	SensitivityInput() *string
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
	ResetActionMode()
	ResetDetectionMode()
	ResetName()
	ResetSensitivity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataCloudflareWafPackagesFilterOutputReference
type jsiiProxy_DataCloudflareWafPackagesFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) ActionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) ActionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) DetectionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"detectionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) DetectionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"detectionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) InternalValue() *DataCloudflareWafPackagesFilter {
	var returns *DataCloudflareWafPackagesFilter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) Sensitivity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sensitivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) SensitivityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sensitivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataCloudflareWafPackagesFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataCloudflareWafPackagesFilterOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareWafPackagesFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareWafPackagesFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareWafPackages.DataCloudflareWafPackagesFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflareWafPackagesFilterOutputReference_Override(d DataCloudflareWafPackagesFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareWafPackages.DataCloudflareWafPackagesFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference)SetActionMode(val *string) {
	if err := j.validateSetActionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionMode",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference)SetDetectionMode(val *string) {
	if err := j.validateSetDetectionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"detectionMode",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference)SetInternalValue(val *DataCloudflareWafPackagesFilter) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference)SetSensitivity(val *string) {
	if err := j.validateSetSensitivityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sensitivity",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) ResetActionMode() {
	_jsii_.InvokeVoid(
		d,
		"resetActionMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) ResetDetectionMode() {
	_jsii_.InvokeVoid(
		d,
		"resetDetectionMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) ResetSensitivity() {
	_jsii_.InvokeVoid(
		d,
		"resetSensitivity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareWafPackagesFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

