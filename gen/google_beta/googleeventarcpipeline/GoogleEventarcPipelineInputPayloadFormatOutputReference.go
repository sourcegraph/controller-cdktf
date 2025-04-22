package googleeventarcpipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleeventarcpipeline/internal"
)

type GoogleEventarcPipelineInputPayloadFormatOutputReference interface {
	cdktf.ComplexObject
	Avro() GoogleEventarcPipelineInputPayloadFormatAvroOutputReference
	AvroInput() *GoogleEventarcPipelineInputPayloadFormatAvro
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
	InternalValue() *GoogleEventarcPipelineInputPayloadFormat
	SetInternalValue(val *GoogleEventarcPipelineInputPayloadFormat)
	Json() GoogleEventarcPipelineInputPayloadFormatJsonOutputReference
	JsonInput() *GoogleEventarcPipelineInputPayloadFormatJson
	Protobuf() GoogleEventarcPipelineInputPayloadFormatProtobufOutputReference
	ProtobufInput() *GoogleEventarcPipelineInputPayloadFormatProtobuf
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
	PutAvro(value *GoogleEventarcPipelineInputPayloadFormatAvro)
	PutJson(value *GoogleEventarcPipelineInputPayloadFormatJson)
	PutProtobuf(value *GoogleEventarcPipelineInputPayloadFormatProtobuf)
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

// The jsii proxy struct for GoogleEventarcPipelineInputPayloadFormatOutputReference
type jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) Avro() GoogleEventarcPipelineInputPayloadFormatAvroOutputReference {
	var returns GoogleEventarcPipelineInputPayloadFormatAvroOutputReference
	_jsii_.Get(
		j,
		"avro",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) AvroInput() *GoogleEventarcPipelineInputPayloadFormatAvro {
	var returns *GoogleEventarcPipelineInputPayloadFormatAvro
	_jsii_.Get(
		j,
		"avroInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) InternalValue() *GoogleEventarcPipelineInputPayloadFormat {
	var returns *GoogleEventarcPipelineInputPayloadFormat
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) Json() GoogleEventarcPipelineInputPayloadFormatJsonOutputReference {
	var returns GoogleEventarcPipelineInputPayloadFormatJsonOutputReference
	_jsii_.Get(
		j,
		"json",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) JsonInput() *GoogleEventarcPipelineInputPayloadFormatJson {
	var returns *GoogleEventarcPipelineInputPayloadFormatJson
	_jsii_.Get(
		j,
		"jsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) Protobuf() GoogleEventarcPipelineInputPayloadFormatProtobufOutputReference {
	var returns GoogleEventarcPipelineInputPayloadFormatProtobufOutputReference
	_jsii_.Get(
		j,
		"protobuf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) ProtobufInput() *GoogleEventarcPipelineInputPayloadFormatProtobuf {
	var returns *GoogleEventarcPipelineInputPayloadFormatProtobuf
	_jsii_.Get(
		j,
		"protobufInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleEventarcPipelineInputPayloadFormatOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleEventarcPipelineInputPayloadFormatOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleEventarcPipelineInputPayloadFormatOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleEventarcPipeline.GoogleEventarcPipelineInputPayloadFormatOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleEventarcPipelineInputPayloadFormatOutputReference_Override(g GoogleEventarcPipelineInputPayloadFormatOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleEventarcPipeline.GoogleEventarcPipelineInputPayloadFormatOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference)SetInternalValue(val *GoogleEventarcPipelineInputPayloadFormat) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) PutAvro(value *GoogleEventarcPipelineInputPayloadFormatAvro) {
	if err := g.validatePutAvroParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAvro",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) PutJson(value *GoogleEventarcPipelineInputPayloadFormatJson) {
	if err := g.validatePutJsonParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putJson",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) PutProtobuf(value *GoogleEventarcPipelineInputPayloadFormatProtobuf) {
	if err := g.validatePutProtobufParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putProtobuf",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) ResetAvro() {
	_jsii_.InvokeVoid(
		g,
		"resetAvro",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) ResetJson() {
	_jsii_.InvokeVoid(
		g,
		"resetJson",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) ResetProtobuf() {
	_jsii_.InvokeVoid(
		g,
		"resetProtobuf",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleEventarcPipelineInputPayloadFormatOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

