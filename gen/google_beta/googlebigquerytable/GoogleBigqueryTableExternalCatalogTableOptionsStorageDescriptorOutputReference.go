package googlebigquerytable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebigquerytable/internal"
)

type GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference interface {
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
	InputFormat() *string
	SetInputFormat(val *string)
	InputFormatInput() *string
	InternalValue() *GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptor
	SetInternalValue(val *GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptor)
	LocationUri() *string
	SetLocationUri(val *string)
	LocationUriInput() *string
	OutputFormat() *string
	SetOutputFormat(val *string)
	OutputFormatInput() *string
	SerdeInfo() GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference
	SerdeInfoInput() *GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfo
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
	PutSerdeInfo(value *GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfo)
	ResetInputFormat()
	ResetLocationUri()
	ResetOutputFormat()
	ResetSerdeInfo()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference
type jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) InputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) InputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) InternalValue() *GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptor {
	var returns *GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptor
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) LocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) LocationUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) OutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) OutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) SerdeInfo() GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference {
	var returns GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference
	_jsii_.Get(
		j,
		"serdeInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) SerdeInfoInput() *GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfo {
	var returns *GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfo
	_jsii_.Get(
		j,
		"serdeInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBigqueryTable.GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference_Override(g GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBigqueryTable.GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference)SetInputFormat(val *string) {
	if err := j.validateSetInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputFormat",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference)SetInternalValue(val *GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptor) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference)SetLocationUri(val *string) {
	if err := j.validateSetLocationUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locationUri",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference)SetOutputFormat(val *string) {
	if err := j.validateSetOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputFormat",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) PutSerdeInfo(value *GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfo) {
	if err := g.validatePutSerdeInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSerdeInfo",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) ResetInputFormat() {
	_jsii_.InvokeVoid(
		g,
		"resetInputFormat",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) ResetLocationUri() {
	_jsii_.InvokeVoid(
		g,
		"resetLocationUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) ResetOutputFormat() {
	_jsii_.InvokeVoid(
		g,
		"resetOutputFormat",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) ResetSerdeInfo() {
	_jsii_.InvokeVoid(
		g,
		"resetSerdeInfo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

