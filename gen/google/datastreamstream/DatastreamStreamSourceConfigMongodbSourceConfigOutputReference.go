package datastreamstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/datastreamstream/internal"
)

type DatastreamStreamSourceConfigMongodbSourceConfigOutputReference interface {
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
	ExcludeObjects() DatastreamStreamSourceConfigMongodbSourceConfigExcludeObjectsOutputReference
	ExcludeObjectsInput() *DatastreamStreamSourceConfigMongodbSourceConfigExcludeObjects
	// Experimental.
	Fqn() *string
	IncludeObjects() DatastreamStreamSourceConfigMongodbSourceConfigIncludeObjectsOutputReference
	IncludeObjectsInput() *DatastreamStreamSourceConfigMongodbSourceConfigIncludeObjects
	InternalValue() *DatastreamStreamSourceConfigMongodbSourceConfig
	SetInternalValue(val *DatastreamStreamSourceConfigMongodbSourceConfig)
	MaxConcurrentBackfillTasks() *float64
	SetMaxConcurrentBackfillTasks(val *float64)
	MaxConcurrentBackfillTasksInput() *float64
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
	PutExcludeObjects(value *DatastreamStreamSourceConfigMongodbSourceConfigExcludeObjects)
	PutIncludeObjects(value *DatastreamStreamSourceConfigMongodbSourceConfigIncludeObjects)
	ResetExcludeObjects()
	ResetIncludeObjects()
	ResetMaxConcurrentBackfillTasks()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatastreamStreamSourceConfigMongodbSourceConfigOutputReference
type jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference) ExcludeObjects() DatastreamStreamSourceConfigMongodbSourceConfigExcludeObjectsOutputReference {
	var returns DatastreamStreamSourceConfigMongodbSourceConfigExcludeObjectsOutputReference
	_jsii_.Get(
		j,
		"excludeObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference) ExcludeObjectsInput() *DatastreamStreamSourceConfigMongodbSourceConfigExcludeObjects {
	var returns *DatastreamStreamSourceConfigMongodbSourceConfigExcludeObjects
	_jsii_.Get(
		j,
		"excludeObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference) IncludeObjects() DatastreamStreamSourceConfigMongodbSourceConfigIncludeObjectsOutputReference {
	var returns DatastreamStreamSourceConfigMongodbSourceConfigIncludeObjectsOutputReference
	_jsii_.Get(
		j,
		"includeObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference) IncludeObjectsInput() *DatastreamStreamSourceConfigMongodbSourceConfigIncludeObjects {
	var returns *DatastreamStreamSourceConfigMongodbSourceConfigIncludeObjects
	_jsii_.Get(
		j,
		"includeObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference) InternalValue() *DatastreamStreamSourceConfigMongodbSourceConfig {
	var returns *DatastreamStreamSourceConfigMongodbSourceConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference) MaxConcurrentBackfillTasks() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentBackfillTasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference) MaxConcurrentBackfillTasksInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentBackfillTasksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatastreamStreamSourceConfigMongodbSourceConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatastreamStreamSourceConfigMongodbSourceConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDatastreamStreamSourceConfigMongodbSourceConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.datastreamStream.DatastreamStreamSourceConfigMongodbSourceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatastreamStreamSourceConfigMongodbSourceConfigOutputReference_Override(d DatastreamStreamSourceConfigMongodbSourceConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.datastreamStream.DatastreamStreamSourceConfigMongodbSourceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference)SetInternalValue(val *DatastreamStreamSourceConfigMongodbSourceConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference)SetMaxConcurrentBackfillTasks(val *float64) {
	if err := j.validateSetMaxConcurrentBackfillTasksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentBackfillTasks",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference) PutExcludeObjects(value *DatastreamStreamSourceConfigMongodbSourceConfigExcludeObjects) {
	if err := d.validatePutExcludeObjectsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putExcludeObjects",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference) PutIncludeObjects(value *DatastreamStreamSourceConfigMongodbSourceConfigIncludeObjects) {
	if err := d.validatePutIncludeObjectsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIncludeObjects",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference) ResetExcludeObjects() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludeObjects",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference) ResetIncludeObjects() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeObjects",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference) ResetMaxConcurrentBackfillTasks() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxConcurrentBackfillTasks",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigMongodbSourceConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

