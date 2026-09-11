package googledataprocsessiontemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataprocsessiontemplate/internal"
)

type GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference interface {
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
	InternalValue() *GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfig
	SetInternalValue(val *GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfig)
	MetastoreService() *string
	SetMetastoreService(val *string)
	MetastoreServiceInput() *string
	SparkHistoryServerConfig() GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference
	SparkHistoryServerConfigInput() *GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfig
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
	PutSparkHistoryServerConfig(value *GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfig)
	ResetMetastoreService()
	ResetSparkHistoryServerConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference
type jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) InternalValue() *GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfig {
	var returns *GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) MetastoreService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metastoreService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) MetastoreServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metastoreServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) SparkHistoryServerConfig() GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference {
	var returns GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference
	_jsii_.Get(
		j,
		"sparkHistoryServerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) SparkHistoryServerConfigInput() *GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfig {
	var returns *GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfig
	_jsii_.Get(
		j,
		"sparkHistoryServerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataprocSessionTemplate.GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference_Override(g GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataprocSessionTemplate.GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference)SetInternalValue(val *GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference)SetMetastoreService(val *string) {
	if err := j.validateSetMetastoreServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metastoreService",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) PutSparkHistoryServerConfig(value *GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfig) {
	if err := g.validatePutSparkHistoryServerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSparkHistoryServerConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) ResetMetastoreService() {
	_jsii_.InvokeVoid(
		g,
		"resetMetastoreService",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) ResetSparkHistoryServerConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSparkHistoryServerConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

