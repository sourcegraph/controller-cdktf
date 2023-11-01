package googlestorageinsightsreportconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlestorageinsightsreportconfig/internal"
)

type GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference interface {
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
	InternalValue() *GoogleStorageInsightsReportConfigObjectMetadataReportOptions
	SetInternalValue(val *GoogleStorageInsightsReportConfigObjectMetadataReportOptions)
	MetadataFields() *[]*string
	SetMetadataFields(val *[]*string)
	MetadataFieldsInput() *[]*string
	StorageDestinationOptions() GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference
	StorageDestinationOptionsInput() *GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptions
	StorageFilters() GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageFiltersOutputReference
	StorageFiltersInput() *GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageFilters
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
	PutStorageDestinationOptions(value *GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptions)
	PutStorageFilters(value *GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageFilters)
	ResetStorageFilters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference
type jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) InternalValue() *GoogleStorageInsightsReportConfigObjectMetadataReportOptions {
	var returns *GoogleStorageInsightsReportConfigObjectMetadataReportOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) MetadataFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metadataFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) MetadataFieldsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metadataFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) StorageDestinationOptions() GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference {
	var returns GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference
	_jsii_.Get(
		j,
		"storageDestinationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) StorageDestinationOptionsInput() *GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptions {
	var returns *GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptions
	_jsii_.Get(
		j,
		"storageDestinationOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) StorageFilters() GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageFiltersOutputReference {
	var returns GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageFiltersOutputReference
	_jsii_.Get(
		j,
		"storageFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) StorageFiltersInput() *GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageFilters {
	var returns *GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageFilters
	_jsii_.Get(
		j,
		"storageFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleStorageInsightsReportConfig.GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference_Override(g GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleStorageInsightsReportConfig.GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference)SetInternalValue(val *GoogleStorageInsightsReportConfigObjectMetadataReportOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference)SetMetadataFields(val *[]*string) {
	if err := j.validateSetMetadataFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadataFields",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) PutStorageDestinationOptions(value *GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptions) {
	if err := g.validatePutStorageDestinationOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStorageDestinationOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) PutStorageFilters(value *GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageFilters) {
	if err := g.validatePutStorageFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStorageFilters",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) ResetStorageFilters() {
	_jsii_.InvokeVoid(
		g,
		"resetStorageFilters",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

