package dataplexdatascan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/dataplexdatascan/internal"
)

type DataplexDatascanExecutionSpecTriggerOutputReference interface {
	cdktn.ComplexObject
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
	InternalValue() *DataplexDatascanExecutionSpecTrigger
	SetInternalValue(val *DataplexDatascanExecutionSpecTrigger)
	OnDemand() DataplexDatascanExecutionSpecTriggerOnDemandOutputReference
	OnDemandInput() *DataplexDatascanExecutionSpecTriggerOnDemand
	OneTime() DataplexDatascanExecutionSpecTriggerOneTimeOutputReference
	OneTimeInput() *DataplexDatascanExecutionSpecTriggerOneTime
	Schedule() DataplexDatascanExecutionSpecTriggerScheduleOutputReference
	ScheduleInput() *DataplexDatascanExecutionSpecTriggerSchedule
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
	PutOnDemand(value *DataplexDatascanExecutionSpecTriggerOnDemand)
	PutOneTime(value *DataplexDatascanExecutionSpecTriggerOneTime)
	PutSchedule(value *DataplexDatascanExecutionSpecTriggerSchedule)
	ResetOnDemand()
	ResetOneTime()
	ResetSchedule()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataplexDatascanExecutionSpecTriggerOutputReference
type jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) InternalValue() *DataplexDatascanExecutionSpecTrigger {
	var returns *DataplexDatascanExecutionSpecTrigger
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) OnDemand() DataplexDatascanExecutionSpecTriggerOnDemandOutputReference {
	var returns DataplexDatascanExecutionSpecTriggerOnDemandOutputReference
	_jsii_.Get(
		j,
		"onDemand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) OnDemandInput() *DataplexDatascanExecutionSpecTriggerOnDemand {
	var returns *DataplexDatascanExecutionSpecTriggerOnDemand
	_jsii_.Get(
		j,
		"onDemandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) OneTime() DataplexDatascanExecutionSpecTriggerOneTimeOutputReference {
	var returns DataplexDatascanExecutionSpecTriggerOneTimeOutputReference
	_jsii_.Get(
		j,
		"oneTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) OneTimeInput() *DataplexDatascanExecutionSpecTriggerOneTime {
	var returns *DataplexDatascanExecutionSpecTriggerOneTime
	_jsii_.Get(
		j,
		"oneTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) Schedule() DataplexDatascanExecutionSpecTriggerScheduleOutputReference {
	var returns DataplexDatascanExecutionSpecTriggerScheduleOutputReference
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) ScheduleInput() *DataplexDatascanExecutionSpecTriggerSchedule {
	var returns *DataplexDatascanExecutionSpecTriggerSchedule
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataplexDatascanExecutionSpecTriggerOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataplexDatascanExecutionSpecTriggerOutputReference {
	_init_.Initialize()

	if err := validateNewDataplexDatascanExecutionSpecTriggerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.dataplexDatascan.DataplexDatascanExecutionSpecTriggerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataplexDatascanExecutionSpecTriggerOutputReference_Override(d DataplexDatascanExecutionSpecTriggerOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dataplexDatascan.DataplexDatascanExecutionSpecTriggerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference)SetInternalValue(val *DataplexDatascanExecutionSpecTrigger) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) PutOnDemand(value *DataplexDatascanExecutionSpecTriggerOnDemand) {
	if err := d.validatePutOnDemandParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOnDemand",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) PutOneTime(value *DataplexDatascanExecutionSpecTriggerOneTime) {
	if err := d.validatePutOneTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOneTime",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) PutSchedule(value *DataplexDatascanExecutionSpecTriggerSchedule) {
	if err := d.validatePutScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSchedule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) ResetOnDemand() {
	_jsii_.InvokeVoid(
		d,
		"resetOnDemand",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) ResetOneTime() {
	_jsii_.InvokeVoid(
		d,
		"resetOneTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) ResetSchedule() {
	_jsii_.InvokeVoid(
		d,
		"resetSchedule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataplexDatascanExecutionSpecTriggerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

