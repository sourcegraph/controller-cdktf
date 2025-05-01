package googleeventarcpipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleeventarcpipeline/internal"
)

type GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference interface {
	cdktf.ComplexObject
	Avro() GoogleEventarcPipelineDestinationsOutputPayloadFormatAvroOutputReference
	AvroInput() *GoogleEventarcPipelineDestinationsOutputPayloadFormatAvro
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
	InternalValue() *GoogleEventarcPipelineDestinationsOutputPayloadFormat
	SetInternalValue(val *GoogleEventarcPipelineDestinationsOutputPayloadFormat)
	Json() GoogleEventarcPipelineDestinationsOutputPayloadFormatJsonOutputReference
	JsonInput() *GoogleEventarcPipelineDestinationsOutputPayloadFormatJson
	Protobuf() GoogleEventarcPipelineDestinationsOutputPayloadFormatProtobufOutputReference
	ProtobufInput() *GoogleEventarcPipelineDestinationsOutputPayloadFormatProtobuf
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
	PutAvro(value *GoogleEventarcPipelineDestinationsOutputPayloadFormatAvro)
	PutJson(value *GoogleEventarcPipelineDestinationsOutputPayloadFormatJson)
	PutProtobuf(value *GoogleEventarcPipelineDestinationsOutputPayloadFormatProtobuf)
	ResetAvro()
	ResetJson()
	ResetProtobuf()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference
type jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) Avro() GoogleEventarcPipelineDestinationsOutputPayloadFormatAvroOutputReference {
	var returns GoogleEventarcPipelineDestinationsOutputPayloadFormatAvroOutputReference
	_jsii_.Get(
		j,
		"avro",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) AvroInput() *GoogleEventarcPipelineDestinationsOutputPayloadFormatAvro {
	var returns *GoogleEventarcPipelineDestinationsOutputPayloadFormatAvro
	_jsii_.Get(
		j,
		"avroInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) InternalValue() *GoogleEventarcPipelineDestinationsOutputPayloadFormat {
	var returns *GoogleEventarcPipelineDestinationsOutputPayloadFormat
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) Json() GoogleEventarcPipelineDestinationsOutputPayloadFormatJsonOutputReference {
	var returns GoogleEventarcPipelineDestinationsOutputPayloadFormatJsonOutputReference
	_jsii_.Get(
		j,
		"json",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) JsonInput() *GoogleEventarcPipelineDestinationsOutputPayloadFormatJson {
	var returns *GoogleEventarcPipelineDestinationsOutputPayloadFormatJson
	_jsii_.Get(
		j,
		"jsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) Protobuf() GoogleEventarcPipelineDestinationsOutputPayloadFormatProtobufOutputReference {
	var returns GoogleEventarcPipelineDestinationsOutputPayloadFormatProtobufOutputReference
	_jsii_.Get(
		j,
		"protobuf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) ProtobufInput() *GoogleEventarcPipelineDestinationsOutputPayloadFormatProtobuf {
	var returns *GoogleEventarcPipelineDestinationsOutputPayloadFormatProtobuf
	_jsii_.Get(
		j,
		"protobufInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleEventarcPipeline.GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference_Override(g GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleEventarcPipeline.GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference)SetInternalValue(val *GoogleEventarcPipelineDestinationsOutputPayloadFormat) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) PutAvro(value *GoogleEventarcPipelineDestinationsOutputPayloadFormatAvro) {
	if err := g.validatePutAvroParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAvro",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) PutJson(value *GoogleEventarcPipelineDestinationsOutputPayloadFormatJson) {
	if err := g.validatePutJsonParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putJson",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) PutProtobuf(value *GoogleEventarcPipelineDestinationsOutputPayloadFormatProtobuf) {
	if err := g.validatePutProtobufParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putProtobuf",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) ResetAvro() {
	_jsii_.InvokeVoid(
		g,
		"resetAvro",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) ResetJson() {
	_jsii_.InvokeVoid(
		g,
		"resetJson",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) ResetProtobuf() {
	_jsii_.InvokeVoid(
		g,
		"resetProtobuf",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

