// Package goXml : autogenerated by go-xsd
package goXml

//	See http://www.w3.org/XML/1998/namespace.html and http://www.w3.org/TR/REC-xml for information about this namespace. This schema document describes the XML namespace, in a form suitable for import by other schema documents. Note that local names in this namespace are intended to be defined only by the World Wide Web Consortium or its subgroups.  The following names are currently defined in this namespace and should not be used with conflicting semantics by any Working Group, specification, or document instance: base (as an attribute name): denotes an attribute whose value provides a URI to be used as the base for interpreting any relative URIs in the scope of the element on which it appears; its value is inherited.  This name is reserved by virtue of its definition in the XML Base specification. lang (as an attribute name): denotes an attribute whose value is a language code for the natural language of the content of any element; its value is inherited.  This name is reserved by virtue of its definition in the XML specification. space (as an attribute name): denotes an attribute whose value is a keyword indicating what whitespace processing discipline is intended for the content of the element; its value is inherited.  This name is reserved by virtue of its definition in the XML specification. Father (in any context at all): denotes Jon Bosak, the chair of the original XML Working Group.  This name is reserved by the following decision of the W3C XML Plenary and XML Coordination groups: In appreciation for his vision, leadership and dedication the W3C XML Plenary on this 10th day of February, 2000 reserves for Jon Bosak in perpetuity the XML name xml:Father This schema defines attributes and an attribute group suitable for use by schemas wishing to allow xml:base, xml:lang or xml:space attributes on elements they define. To enable this, such a schema must import this schema for the XML namespace, e.g. as follows: <schema . . .> . . . <import namespace="http://www.w3.org/XML/1998/namespace" schemaLocation="http://www.w3.org/2001/03/xml.xsd"/> Subsequently, qualified reference to any of the attributes or the group defined below will have the desired effect, e.g. <type . . .> . . . <attributeGroup ref="xml:specialAttrs"/> will define a type which will schema-validate an instance element with any of those attributes In keeping with the XML Schema WG's standard versioning policy, this schema document will persist at http://www.w3.org/2001/03/xml.xsd. At the date of issue it can also be found at http://www.w3.org/2001/xml.xsd. The schema document at that URI may however change in the future, in order to remain compatible with the latest version of XML Schema itself.  In other words, if the XML Schema namespace changes, the version of this document at http://www.w3.org/2001/xml.xsd will change accordingly; the version at http://www.w3.org/2001/03/xml.xsd will not change.

import (
	xsdt "github.com/miracl/go-xsd-pkg/xsdt"
)

// XsdGoPkgHasAttrLang In due course, we should install the relevant ISO 2- and 3-letter
// codes as the enumerated possible values . . .
type XsdGoPkgHasAttrLang struct {
	//	In due course, we should install the relevant ISO 2- and 3-letter
	//	codes as the enumerated possible values . . .
	Lang xsdt.Language `xml:"http://www.w3.org/XML/1998/namespace lang,attr"`
}

// TxsdSpace defines TxsdSpace
type TxsdSpace xsdt.NCName

// IsPreserve : Returns true if the value of this enumerated TxsdSpace is "preserve".
func (me TxsdSpace) IsPreserve() bool { return me.String() == "preserve" }

// Set : Since TxsdSpace is just a simple String type, this merely sets the current value from the specified string.
func (me *TxsdSpace) Set(s string) { (*xsdt.NCName)(me).Set(s) }

// String : Since TxsdSpace is just a simple String type, this merely returns the current string value.
func (me TxsdSpace) String() string { return xsdt.NCName(me).String() }

// ToXsdtNCName : This convenience method just performs a simple type conversion to TxsdSpace's alias type xsdt.NCName.
func (me TxsdSpace) ToXsdtNCName() xsdt.NCName { return xsdt.NCName(me) }

// IsDefault : Returns true if the value of this enumerated TxsdSpace is "default".
func (me TxsdSpace) IsDefault() bool { return me.String() == "default" }

// XsdGoPkgHasAttrSpace defines attribute Space
type XsdGoPkgHasAttrSpace struct {
	Space TxsdSpace `xml:"http://www.w3.org/XML/1998/namespace space,attr"`
}

// SpaceDefault : Returns the default value for Space -- "preserve"
func (me XsdGoPkgHasAttrSpace) SpaceDefault() TxsdSpace { return TxsdSpace("preserve") }

// XsdGoPkgHasAttrBase See http://www.w3.org/TR/xmlbase/ for
// information about this attribute.
type XsdGoPkgHasAttrBase struct {
	//	See http://www.w3.org/TR/xmlbase/ for
	//	information about this attribute.
	Base xsdt.AnyURI `xml:"http://www.w3.org/XML/1998/namespace base,attr"`
}

// XsdGoPkgHasAttsSpecialAttrs defines XsdGoPkgHasAttsSpecialAttrs
type XsdGoPkgHasAttsSpecialAttrs struct {
	XsdGoPkgHasAttrBase

	XsdGoPkgHasAttrLang

	XsdGoPkgHasAttrSpace
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
	// WalkHandlers Provides 1 strong-typed hooks for your own custom handler functions to be invoked when the Walk() method is called on any instance of any (non-attribute-related) struct type defined in this package.
	// If your custom handler does get called at all for a given struct instance, then it always gets called twice, first with the 'enter' bool argument set to true, then (after having Walk()ed all subordinate struct instances, if any) once again with it set to false.
	WalkHandlers = &XsdGoPkgWalkHandlers{}
)

// XsdGoPkgWalkHandlers Provides 1 strong-typed hooks for your own custom handler functions to be invoked when the Walk() method is called on any instance of any (non-attribute-related) struct type defined in this package.
// If your custom handler does get called at all for a given struct instance, then it always gets called twice, first with the 'enter' bool argument set to true, then (after having Walk()ed all subordinate struct instances, if any) once again with it set to false.
type XsdGoPkgWalkHandlers struct {
	XsdGoPkgHasCdata func(*XsdGoPkgHasCdata, bool) error
}
