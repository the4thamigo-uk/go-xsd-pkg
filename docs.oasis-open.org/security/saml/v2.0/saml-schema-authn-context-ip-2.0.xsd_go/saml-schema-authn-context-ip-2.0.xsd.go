// Package goSamlAuthnContextIp20 : autogenerated by go-xsd
package goSamlAuthnContextIp20

import (
	sac "github.com/miracl/go-xsd-pkg/docs.oasis-open.org/security/saml/v2.0/saml-schema-authn-context-2.0.xsd_go"
	sact "github.com/miracl/go-xsd-pkg/docs.oasis-open.org/security/saml/v2.0/saml-schema-authn-context-types-2.0.xsd_go"
	xsdt "github.com/miracl/go-xsd-pkg/xsdt"
)

// XsdGoPkgHasAttrIDXsdtID defines attribute IdXsdtId
type XsdGoPkgHasAttrIDXsdtID struct {
	ID xsdt.ID `xml:"urn:oasis:names:tc:SAML:2.0:ac:classes:InternetProtocol ID,attr"`
}

// TAuthnContextDeclarationBaseType defines TAuthnContextDeclarationBaseType
type TAuthnContextDeclarationBaseType struct {
	sact.XsdGoPkgHasElemTechnicalProtection

	sac.XsdGoPkgHasElemOperationalProtection

	sac.XsdGoPkgHasElemAuthnMethod

	sac.XsdGoPkgHasElemGoverningAgreements

	sac.XsdGoPkgHasElemsExtension

	XsdGoPkgHasAttrIDXsdtID

	*TAuthnContextDeclarationBaseType

	sac.XsdGoPkgHasElemIdentification
}

// Walk : if the WalkHandlers.TAuthnContextDeclarationBaseType function is not nil (ie. was set by outside code), calls it with this TAuthnContextDeclarationBaseType instance as the single argument. Then calls the Walk() method on 1/8 embed(s) and 0/0 field(s) belonging to this TAuthnContextDeclarationBaseType instance.
func (me *TAuthnContextDeclarationBaseType) Walk() (err error) {
	if fn := WalkHandlers.TAuthnContextDeclarationBaseType; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.TAuthnContextDeclarationBaseType.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

// TAuthnMethodBaseType defines TAuthnMethodBaseType
type TAuthnMethodBaseType struct {
	sac.XsdGoPkgHasElemsExtension

	*TAuthnMethodBaseType

	sac.XsdGoPkgHasElemPrincipalAuthenticationMechanism

	sac.XsdGoPkgHasElemAuthenticator

	sac.XsdGoPkgHasElemAuthenticatorTransportProtocol
}

// Walk : if the WalkHandlers.TAuthnMethodBaseType function is not nil (ie. was set by outside code), calls it with this TAuthnMethodBaseType instance as the single argument. Then calls the Walk() method on 1/5 embed(s) and 0/0 field(s) belonging to this TAuthnMethodBaseType instance.
func (me *TAuthnMethodBaseType) Walk() (err error) {
	if fn := WalkHandlers.TAuthnMethodBaseType; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.TAuthnMethodBaseType.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

// TAuthenticatorBaseType defines TAuthenticatorBaseType
type TAuthenticatorBaseType struct {
	*TAuthenticatorBaseType

	sac.XsdGoPkgHasElemIPAddress
}

// Walk : if the WalkHandlers.TAuthenticatorBaseType function is not nil (ie. was set by outside code), calls it with this TAuthenticatorBaseType instance as the single argument. Then calls the Walk() method on 1/2 embed(s) and 0/0 field(s) belonging to this TAuthenticatorBaseType instance.
func (me *TAuthenticatorBaseType) Walk() (err error) {
	if fn := WalkHandlers.TAuthenticatorBaseType; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.TAuthenticatorBaseType.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

// XsdGoPkgHasCdata defines type CDATA
type XsdGoPkgHasCdata struct {
	XsdGoPkgCDATA string `xml:",chardata"`
}

// Walk : if the WalkHandlers.XsdGoPkgHasCdata function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasCdata instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasCdata instance.
func (me *XsdGoPkgHasCdata) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasCdata; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

var (
	// WalkContinueOnError can be set to false to break a Walk() immediately as soon as the first error is returned by a custom handler function.
	// If true, Walk() proceeds and accumulates all errors in the WalkErrors slice.
	WalkContinueOnError = true
	// WalkErrors contains all errors accumulated during Walk()s. If you're using this, you need to reset this yourself as needed prior to a fresh Walk().
	WalkErrors []error
	// WalkOnError is your custom error-handling function, if required.
	WalkOnError func(error)
	// WalkHandlers Provides 4 strong-typed hooks for your own custom handler functions to be invoked when the Walk() method is called on any instance of any (non-attribute-related) struct type defined in this package.
	// If your custom handler does get called at all for a given struct instance, then it always gets called twice, first with the 'enter' bool argument set to true, then (after having Walk()ed all subordinate struct instances, if any) once again with it set to false.
	WalkHandlers = &XsdGoPkgWalkHandlers{}
)

// XsdGoPkgWalkHandlers Provides 4 strong-typed hooks for your own custom handler functions to be invoked when the Walk() method is called on any instance of any (non-attribute-related) struct type defined in this package.
// If your custom handler does get called at all for a given struct instance, then it always gets called twice, first with the 'enter' bool argument set to true, then (after having Walk()ed all subordinate struct instances, if any) once again with it set to false.
type XsdGoPkgWalkHandlers struct {
	TAuthnMethodBaseType             func(*TAuthnMethodBaseType, bool) error
	TAuthenticatorBaseType           func(*TAuthenticatorBaseType, bool) error
	XsdGoPkgHasCdata                 func(*XsdGoPkgHasCdata, bool) error
	TAuthnContextDeclarationBaseType func(*TAuthnContextDeclarationBaseType, bool) error
}
