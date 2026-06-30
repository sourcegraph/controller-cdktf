package datastreamconnectionprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/datastreamconnectionprofile/internal"
)

type DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference interface {
	cdktf.ComplexObject
	CaCertificate() *string
	SetCaCertificate(val *string)
	CaCertificateInput() *string
	ClientCertificate() *string
	SetClientCertificate(val *string)
	ClientCertificateInput() *string
	ClientKey() *string
	SetClientKey(val *string)
	ClientKeyInput() *string
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
	InternalValue() *DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerification
	SetInternalValue(val *DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerification)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference
type jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) CaCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) CaCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) ClientCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) ClientCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) ClientKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) ClientKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) InternalValue() *DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerification {
	var returns *DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerification
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference {
	_init_.Initialize()

	if err := validateNewDatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.datastreamConnectionProfile.DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference_Override(d DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.datastreamConnectionProfile.DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference",
		[]any{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) SetCaCertificate(val *string) {
	if err := j.validateSetCaCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caCertificate",
		val,
	)
}

func (j *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) SetClientCertificate(val *string) {
	if err := j.validateSetClientCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificate",
		val,
	)
}

func (j *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) SetClientKey(val *string) {
	if err := j.validateSetClientKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientKey",
		val,
	)
}

func (j *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) SetInternalValue(val *DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerification) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		d,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerificationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
