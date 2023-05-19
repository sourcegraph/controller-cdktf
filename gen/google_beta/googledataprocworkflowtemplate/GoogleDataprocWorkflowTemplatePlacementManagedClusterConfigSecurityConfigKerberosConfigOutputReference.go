package googledataprocworkflowtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataprocworkflowtemplate/internal"
)

type GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference interface {
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
	CrossRealmTrustAdminServer() *string
	SetCrossRealmTrustAdminServer(val *string)
	CrossRealmTrustAdminServerInput() *string
	CrossRealmTrustKdc() *string
	SetCrossRealmTrustKdc(val *string)
	CrossRealmTrustKdcInput() *string
	CrossRealmTrustRealm() *string
	SetCrossRealmTrustRealm(val *string)
	CrossRealmTrustRealmInput() *string
	CrossRealmTrustSharedPassword() *string
	SetCrossRealmTrustSharedPassword(val *string)
	CrossRealmTrustSharedPasswordInput() *string
	EnableKerberos() interface{}
	SetEnableKerberos(val interface{})
	EnableKerberosInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig
	SetInternalValue(val *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig)
	KdcDbKey() *string
	SetKdcDbKey(val *string)
	KdcDbKeyInput() *string
	KeyPassword() *string
	SetKeyPassword(val *string)
	KeyPasswordInput() *string
	Keystore() *string
	SetKeystore(val *string)
	KeystoreInput() *string
	KeystorePassword() *string
	SetKeystorePassword(val *string)
	KeystorePasswordInput() *string
	KmsKey() *string
	SetKmsKey(val *string)
	KmsKeyInput() *string
	Realm() *string
	SetRealm(val *string)
	RealmInput() *string
	RootPrincipalPassword() *string
	SetRootPrincipalPassword(val *string)
	RootPrincipalPasswordInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TgtLifetimeHours() *float64
	SetTgtLifetimeHours(val *float64)
	TgtLifetimeHoursInput() *float64
	Truststore() *string
	SetTruststore(val *string)
	TruststoreInput() *string
	TruststorePassword() *string
	SetTruststorePassword(val *string)
	TruststorePasswordInput() *string
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
	ResetCrossRealmTrustAdminServer()
	ResetCrossRealmTrustKdc()
	ResetCrossRealmTrustRealm()
	ResetCrossRealmTrustSharedPassword()
	ResetEnableKerberos()
	ResetKdcDbKey()
	ResetKeyPassword()
	ResetKeystore()
	ResetKeystorePassword()
	ResetKmsKey()
	ResetRealm()
	ResetRootPrincipalPassword()
	ResetTgtLifetimeHours()
	ResetTruststore()
	ResetTruststorePassword()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference
type jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) CrossRealmTrustAdminServer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustAdminServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) CrossRealmTrustAdminServerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustAdminServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) CrossRealmTrustKdc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustKdc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) CrossRealmTrustKdcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustKdcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) CrossRealmTrustRealm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustRealm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) CrossRealmTrustRealmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustRealmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) CrossRealmTrustSharedPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustSharedPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) CrossRealmTrustSharedPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustSharedPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) EnableKerberos() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableKerberos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) EnableKerberosInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableKerberosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) InternalValue() *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig {
	var returns *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) KdcDbKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kdcDbKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) KdcDbKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kdcDbKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) KeyPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) KeyPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) Keystore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keystore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) KeystoreInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keystoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) KeystorePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keystorePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) KeystorePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keystorePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) KmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) KmsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) Realm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) RealmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) RootPrincipalPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootPrincipalPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) RootPrincipalPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootPrincipalPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) TgtLifetimeHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tgtLifetimeHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) TgtLifetimeHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tgtLifetimeHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) Truststore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"truststore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) TruststoreInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"truststoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) TruststorePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"truststorePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) TruststorePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"truststorePasswordInput",
		&returns,
	)
	return returns
}


func NewGoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference{}

	_jsii_.Create(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference_Override(g GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetCrossRealmTrustAdminServer(val *string) {
	if err := j.validateSetCrossRealmTrustAdminServerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossRealmTrustAdminServer",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetCrossRealmTrustKdc(val *string) {
	if err := j.validateSetCrossRealmTrustKdcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossRealmTrustKdc",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetCrossRealmTrustRealm(val *string) {
	if err := j.validateSetCrossRealmTrustRealmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossRealmTrustRealm",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetCrossRealmTrustSharedPassword(val *string) {
	if err := j.validateSetCrossRealmTrustSharedPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossRealmTrustSharedPassword",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetEnableKerberos(val interface{}) {
	if err := j.validateSetEnableKerberosParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableKerberos",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetInternalValue(val *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetKdcDbKey(val *string) {
	if err := j.validateSetKdcDbKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kdcDbKey",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetKeyPassword(val *string) {
	if err := j.validateSetKeyPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyPassword",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetKeystore(val *string) {
	if err := j.validateSetKeystoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keystore",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetKeystorePassword(val *string) {
	if err := j.validateSetKeystorePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keystorePassword",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetKmsKey(val *string) {
	if err := j.validateSetKmsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKey",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetRealm(val *string) {
	if err := j.validateSetRealmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"realm",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetRootPrincipalPassword(val *string) {
	if err := j.validateSetRootPrincipalPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootPrincipalPassword",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetTgtLifetimeHours(val *float64) {
	if err := j.validateSetTgtLifetimeHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tgtLifetimeHours",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetTruststore(val *string) {
	if err := j.validateSetTruststoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"truststore",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)SetTruststorePassword(val *string) {
	if err := j.validateSetTruststorePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"truststorePassword",
		val,
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ResetCrossRealmTrustAdminServer() {
	_jsii_.InvokeVoid(
		g,
		"resetCrossRealmTrustAdminServer",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ResetCrossRealmTrustKdc() {
	_jsii_.InvokeVoid(
		g,
		"resetCrossRealmTrustKdc",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ResetCrossRealmTrustRealm() {
	_jsii_.InvokeVoid(
		g,
		"resetCrossRealmTrustRealm",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ResetCrossRealmTrustSharedPassword() {
	_jsii_.InvokeVoid(
		g,
		"resetCrossRealmTrustSharedPassword",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ResetEnableKerberos() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableKerberos",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ResetKdcDbKey() {
	_jsii_.InvokeVoid(
		g,
		"resetKdcDbKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ResetKeyPassword() {
	_jsii_.InvokeVoid(
		g,
		"resetKeyPassword",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ResetKeystore() {
	_jsii_.InvokeVoid(
		g,
		"resetKeystore",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ResetKeystorePassword() {
	_jsii_.InvokeVoid(
		g,
		"resetKeystorePassword",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ResetKmsKey() {
	_jsii_.InvokeVoid(
		g,
		"resetKmsKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ResetRealm() {
	_jsii_.InvokeVoid(
		g,
		"resetRealm",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ResetRootPrincipalPassword() {
	_jsii_.InvokeVoid(
		g,
		"resetRootPrincipalPassword",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ResetTgtLifetimeHours() {
	_jsii_.InvokeVoid(
		g,
		"resetTgtLifetimeHours",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ResetTruststore() {
	_jsii_.InvokeVoid(
		g,
		"resetTruststore",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ResetTruststorePassword() {
	_jsii_.InvokeVoid(
		g,
		"resetTruststorePassword",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

