package dataproccluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/dataproccluster/internal"
)

type DataprocClusterClusterConfigLifecycleConfigOutputReference interface {
	cdktn.ComplexObject
	AutoDeleteTime() *string
	SetAutoDeleteTime(val *string)
	AutoDeleteTimeInput() *string
	AutoStopTime() *string
	SetAutoStopTime(val *string)
	AutoStopTimeInput() *string
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
	IdleDeleteTtl() *string
	SetIdleDeleteTtl(val *string)
	IdleDeleteTtlInput() *string
	IdleStartTime() *string
	IdleStopTtl() *string
	SetIdleStopTtl(val *string)
	IdleStopTtlInput() *string
	InternalValue() *DataprocClusterClusterConfigLifecycleConfig
	SetInternalValue(val *DataprocClusterClusterConfigLifecycleConfig)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	ResetAutoDeleteTime()
	ResetAutoStopTime()
	ResetIdleDeleteTtl()
	ResetIdleStopTtl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataprocClusterClusterConfigLifecycleConfigOutputReference
type jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) AutoDeleteTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoDeleteTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) AutoDeleteTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoDeleteTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) AutoStopTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoStopTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) AutoStopTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoStopTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) IdleDeleteTtl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idleDeleteTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) IdleDeleteTtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idleDeleteTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) IdleStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idleStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) IdleStopTtl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idleStopTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) IdleStopTtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idleStopTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) InternalValue() *DataprocClusterClusterConfigLifecycleConfig {
	var returns *DataprocClusterClusterConfigLifecycleConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataprocClusterClusterConfigLifecycleConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataprocClusterClusterConfigLifecycleConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataprocClusterClusterConfigLifecycleConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.dataprocCluster.DataprocClusterClusterConfigLifecycleConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataprocClusterClusterConfigLifecycleConfigOutputReference_Override(d DataprocClusterClusterConfigLifecycleConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dataprocCluster.DataprocClusterClusterConfigLifecycleConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference)SetAutoDeleteTime(val *string) {
	if err := j.validateSetAutoDeleteTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoDeleteTime",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference)SetAutoStopTime(val *string) {
	if err := j.validateSetAutoStopTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoStopTime",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference)SetIdleDeleteTtl(val *string) {
	if err := j.validateSetIdleDeleteTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleDeleteTtl",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference)SetIdleStopTtl(val *string) {
	if err := j.validateSetIdleStopTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleStopTtl",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference)SetInternalValue(val *DataprocClusterClusterConfigLifecycleConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) ResetAutoDeleteTime() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoDeleteTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) ResetAutoStopTime() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoStopTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) ResetIdleDeleteTtl() {
	_jsii_.InvokeVoid(
		d,
		"resetIdleDeleteTtl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) ResetIdleStopTtl() {
	_jsii_.InvokeVoid(
		d,
		"resetIdleStopTtl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigLifecycleConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

