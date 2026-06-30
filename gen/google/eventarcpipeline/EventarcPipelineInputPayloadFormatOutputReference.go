package eventarcpipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/eventarcpipeline/internal"
)

type EventarcPipelineInputPayloadFormatOutputReference interface {
	cdktf.ComplexObject
	Avro() EventarcPipelineInputPayloadFormatAvroOutputReference
	AvroInput() *EventarcPipelineInputPayloadFormatAvro
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *EventarcPipelineInputPayloadFormat
	SetInternalValue(val *EventarcPipelineInputPayloadFormat)
	Json() EventarcPipelineInputPayloadFormatJsonOutputReference
	JsonInput() *EventarcPipelineInputPayloadFormatJson
	Protobuf() EventarcPipelineInputPayloadFormatProtobufOutputReference
	ProtobufInput() *EventarcPipelineInputPayloadFormatProtobuf
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
	PutAvro(value *EventarcPipelineInputPayloadFormatAvro)
	PutJson(value *EventarcPipelineInputPayloadFormatJson)
	PutProtobuf(value *EventarcPipelineInputPayloadFormatProtobuf)
	ResetAvro()
	ResetJson()
	ResetProtobuf()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventarcPipelineInputPayloadFormatOutputReference
type jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) Avro() EventarcPipelineInputPayloadFormatAvroOutputReference {
	var returns EventarcPipelineInputPayloadFormatAvroOutputReference
	_jsii_.Get(
		j,
		"avro",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) AvroInput() *EventarcPipelineInputPayloadFormatAvro {
	var returns *EventarcPipelineInputPayloadFormatAvro
	_jsii_.Get(
		j,
		"avroInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) InternalValue() *EventarcPipelineInputPayloadFormat {
	var returns *EventarcPipelineInputPayloadFormat
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) Json() EventarcPipelineInputPayloadFormatJsonOutputReference {
	var returns EventarcPipelineInputPayloadFormatJsonOutputReference
	_jsii_.Get(
		j,
		"json",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) JsonInput() *EventarcPipelineInputPayloadFormatJson {
	var returns *EventarcPipelineInputPayloadFormatJson
	_jsii_.Get(
		j,
		"jsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) Protobuf() EventarcPipelineInputPayloadFormatProtobufOutputReference {
	var returns EventarcPipelineInputPayloadFormatProtobufOutputReference
	_jsii_.Get(
		j,
		"protobuf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) ProtobufInput() *EventarcPipelineInputPayloadFormatProtobuf {
	var returns *EventarcPipelineInputPayloadFormatProtobuf
	_jsii_.Get(
		j,
		"protobufInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewEventarcPipelineInputPayloadFormatOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventarcPipelineInputPayloadFormatOutputReference {
	_init_.Initialize()

	if err := validateNewEventarcPipelineInputPayloadFormatOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.eventarcPipeline.EventarcPipelineInputPayloadFormatOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventarcPipelineInputPayloadFormatOutputReference_Override(e EventarcPipelineInputPayloadFormatOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.eventarcPipeline.EventarcPipelineInputPayloadFormatOutputReference",
		[]any{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) SetInternalValue(val *EventarcPipelineInputPayloadFormat) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) PutAvro(value *EventarcPipelineInputPayloadFormatAvro) {
	if err := e.validatePutAvroParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAvro",
		[]any{value},
	)
}

func (e *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) PutJson(value *EventarcPipelineInputPayloadFormatJson) {
	if err := e.validatePutJsonParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putJson",
		[]any{value},
	)
}

func (e *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) PutProtobuf(value *EventarcPipelineInputPayloadFormatProtobuf) {
	if err := e.validatePutProtobufParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putProtobuf",
		[]any{value},
	)
}

func (e *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) ResetAvro() {
	_jsii_.InvokeVoid(
		e,
		"resetAvro",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) ResetJson() {
	_jsii_.InvokeVoid(
		e,
		"resetJson",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) ResetProtobuf() {
	_jsii_.InvokeVoid(
		e,
		"resetProtobuf",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		e,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventarcPipelineInputPayloadFormatOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
