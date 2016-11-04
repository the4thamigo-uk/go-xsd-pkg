//	Auto-generated by the "go-xsd" package located at:
//		github.com/miracl/go-xsd
//	Comments on types and fields (if any) are from the XSD file located at:
//		www.w3.org/TR/2002/WD-SVG11-20020108/xlink.xsd
package go_Xlink

import (
	xsdt "github.com/miracl/go-xsd-pkg/xsdt"
)

type TxsdType xsdt.String

//	Returns true if the value of this enumerated TxsdType is "locator".
func (me TxsdType) IsLocator() bool { return me.String() == "locator" }

//	Returns true if the value of this enumerated TxsdType is "arc".
func (me TxsdType) IsArc() bool { return me.String() == "arc" }

//	Since TxsdType is just a simple String type, this merely sets the current value from the specified string.
func (me *TxsdType) Set(s string) { (*xsdt.String)(me).Set(s) }

//	Since TxsdType is just a simple String type, this merely returns the current string value.
func (me TxsdType) String() string { return xsdt.String(me).String() }

//	This convenience method just performs a simple type conversion to TxsdType's alias type xsdt.String.
func (me TxsdType) ToXsdtString() xsdt.String { return xsdt.String(me) }

//	Returns true if the value of this enumerated TxsdType is "simple".
func (me TxsdType) IsSimple() bool { return me.String() == "simple" }

//	Returns true if the value of this enumerated TxsdType is "extended".
func (me TxsdType) IsExtended() bool { return me.String() == "extended" }

type XsdGoPkgHasAttr_Type struct {
	Type TxsdType `xml:"http://www.w3.org/1999/xlink type,attr"`
}

//	Returns the default value for Type -- "simple"
func (me XsdGoPkgHasAttr_Type) TypeDefault() TxsdType { return TxsdType("simple") }

type XsdGoPkgHasAttr_Href struct {
	Href xsdt.AnyURI `xml:"http://www.w3.org/1999/xlink href,attr"`
}

type XsdGoPkgHasAttr_Role struct {
	Role xsdt.String `xml:"http://www.w3.org/1999/xlink role,attr"`
}

type XsdGoPkgHasAttr_Arcrole struct {
	Arcrole xsdt.String `xml:"http://www.w3.org/1999/xlink arcrole,attr"`
}

type XsdGoPkgHasAttr_Title struct {
	Title xsdt.String `xml:"http://www.w3.org/1999/xlink title,attr"`
}

type TxsdShow xsdt.String

//	Returns true if the value of this enumerated TxsdShow is "new".
func (me TxsdShow) IsNew() bool { return me.String() == "new" }

//	Returns true if the value of this enumerated TxsdShow is "replace".
func (me TxsdShow) IsReplace() bool { return me.String() == "replace" }

//	Returns true if the value of this enumerated TxsdShow is "embed".
func (me TxsdShow) IsEmbed() bool { return me.String() == "embed" }

//	Returns true if the value of this enumerated TxsdShow is "other".
func (me TxsdShow) IsOther() bool { return me.String() == "other" }

//	Returns true if the value of this enumerated TxsdShow is "none".
func (me TxsdShow) IsNone() bool { return me.String() == "none" }

//	Since TxsdShow is just a simple String type, this merely sets the current value from the specified string.
func (me *TxsdShow) Set(s string) { (*xsdt.String)(me).Set(s) }

//	Since TxsdShow is just a simple String type, this merely returns the current string value.
func (me TxsdShow) String() string { return xsdt.String(me).String() }

//	This convenience method just performs a simple type conversion to TxsdShow's alias type xsdt.String.
func (me TxsdShow) ToXsdtString() xsdt.String { return xsdt.String(me) }

type XsdGoPkgHasAttr_Show struct {
	Show TxsdShow `xml:"http://www.w3.org/1999/xlink show,attr"`
}

//	Returns the default value for Show -- "embed"
func (me XsdGoPkgHasAttr_Show) ShowDefault() TxsdShow { return TxsdShow("embed") }

type TxsdActuate xsdt.String

//	Returns true if the value of this enumerated TxsdActuate is "none".
func (me TxsdActuate) IsNone() bool { return me.String() == "none" }

//	Since TxsdActuate is just a simple String type, this merely sets the current value from the specified string.
func (me *TxsdActuate) Set(s string) { (*xsdt.String)(me).Set(s) }

//	Since TxsdActuate is just a simple String type, this merely returns the current string value.
func (me TxsdActuate) String() string { return xsdt.String(me).String() }

//	This convenience method just performs a simple type conversion to TxsdActuate's alias type xsdt.String.
func (me TxsdActuate) ToXsdtString() xsdt.String { return xsdt.String(me) }

//	Returns true if the value of this enumerated TxsdActuate is "onLoad".
func (me TxsdActuate) IsOnLoad() bool { return me.String() == "onLoad" }

//	Returns true if the value of this enumerated TxsdActuate is "onRequest".
func (me TxsdActuate) IsOnRequest() bool { return me.String() == "onRequest" }

//	Returns true if the value of this enumerated TxsdActuate is "other".
func (me TxsdActuate) IsOther() bool { return me.String() == "other" }

type XsdGoPkgHasAttr_Actuate struct {
	Actuate TxsdActuate `xml:"http://www.w3.org/1999/xlink actuate,attr"`
}

//	Returns the default value for Actuate -- "onLoad"
func (me XsdGoPkgHasAttr_Actuate) ActuateDefault() TxsdActuate { return TxsdActuate("onLoad") }

type XsdGoPkgHasAttr_From struct {
	From xsdt.String `xml:"http://www.w3.org/1999/xlink from,attr"`
}

type XsdGoPkgHasAttr_To struct {
	To xsdt.String `xml:"http://www.w3.org/1999/xlink to,attr"`
}

type XsdGoPkgHasCdata struct {
	XsdGoPkgCDATA string `xml:",chardata"`
}

//	If the WalkHandlers.XsdGoPkgHasCdata function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasCdata instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasCdata instance.
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
	//	Set this to false to break a Walk() immediately as soon as the first error is returned by a custom handler function.
	//	If true, Walk() proceeds and accumulates all errors in the WalkErrors slice.
	WalkContinueOnError = true
	//	Contains all errors accumulated during Walk()s. If you're using this, you need to reset this yourself as needed prior to a fresh Walk().
	WalkErrors []error
	//	Your custom error-handling function, if required.
	WalkOnError func(error)
	//	Provides 1 strong-typed hooks for your own custom handler functions to be invoked when the Walk() method is called on any instance of any (non-attribute-related) struct type defined in this package.
	//	If your custom handler does get called at all for a given struct instance, then it always gets called twice, first with the 'enter' bool argument set to true, then (after having Walk()ed all subordinate struct instances, if any) once again with it set to false.
	WalkHandlers = &XsdGoPkgWalkHandlers{}
)

//	Provides 1 strong-typed hooks for your own custom handler functions to be invoked when the Walk() method is called on any instance of any (non-attribute-related) struct type defined in this package.
//	If your custom handler does get called at all for a given struct instance, then it always gets called twice, first with the 'enter' bool argument set to true, then (after having Walk()ed all subordinate struct instances, if any) once again with it set to false.
type XsdGoPkgWalkHandlers struct {
	XsdGoPkgHasCdata func(*XsdGoPkgHasCdata, bool) error
}
