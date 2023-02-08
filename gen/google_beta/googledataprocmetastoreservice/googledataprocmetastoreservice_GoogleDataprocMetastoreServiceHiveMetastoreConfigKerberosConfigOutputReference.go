package googledataprocmetastoreservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataprocmetastoreservice/internal"
)

type GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference interface {
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
	InternalValue() *GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfig
	SetInternalValue(val *GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfig)
	Keytab() GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigKeytabOutputReference
	KeytabInput() *GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigKeytab
	Krb5ConfigGcsUri() *string
	SetKrb5ConfigGcsUri(val *string)
	Krb5ConfigGcsUriInput() *string
	Principal() *string
	SetPrincipal(val *string)
	PrincipalInput() *string
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
	PutKeytab(value *GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigKeytab)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference
type jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) InternalValue() *GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfig {
	var returns *GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) Keytab() GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigKeytabOutputReference {
	var returns GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigKeytabOutputReference
	_jsii_.Get(
		j,
		"keytab",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) KeytabInput() *GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigKeytab {
	var returns *GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigKeytab
	_jsii_.Get(
		j,
		"keytabInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) Krb5ConfigGcsUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"krb5ConfigGcsUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) Krb5ConfigGcsUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"krb5ConfigGcsUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) Principal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) PrincipalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference{}

	_jsii_.Create(
		"google-beta.googleDataprocMetastoreService.GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference_Override(g GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDataprocMetastoreService.GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference)SetInternalValue(val *GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference)SetKrb5ConfigGcsUri(val *string) {
	if err := j.validateSetKrb5ConfigGcsUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"krb5ConfigGcsUri",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference)SetPrincipal(val *string) {
	if err := j.validateSetPrincipalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"principal",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) PutKeytab(value *GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigKeytab) {
	if err := g.validatePutKeytabParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putKeytab",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

