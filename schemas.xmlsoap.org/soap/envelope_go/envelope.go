// Package goEnvelope : autogenerated by go-xsd
package goEnvelope

import (
	xsdt "github.com/miracl/go-xsd-pkg/xsdt"
)

// TxsdMustUnderstand defines TxsdMustUnderstand
type TxsdMustUnderstand xsdt.Boolean

// Set : Since TxsdMustUnderstand is a non-string scalar type (either boolean or numeric), sets the current value obtained from parsing the specified string.
func (me *TxsdMustUnderstand) Set(s string) { (*xsdt.Boolean)(me).Set(s) }

// String : Returns a string representation of this TxsdMustUnderstand's current non-string scalar value.
func (me TxsdMustUnderstand) String() string { return xsdt.Boolean(me).String() }

// ToXsdtBoolean : This convenience method just performs a simple type conversion to TxsdMustUnderstand's alias type xsdt.Boolean.
func (me TxsdMustUnderstand) ToXsdtBoolean() xsdt.Boolean { return xsdt.Boolean(me) }

// XsdGoPkgHasAttrMustUnderstand defines attribute MustUnderstand
type XsdGoPkgHasAttrMustUnderstand struct {
	MustUnderstand TxsdMustUnderstand `xml:"http://schemas.xmlsoap.org/soap/envelope/ mustUnderstand,attr"`
}

// XsdGoPkgHasAttrActor defines attribute Actor
type XsdGoPkgHasAttrActor struct {
	Actor xsdt.AnyURI `xml:"http://schemas.xmlsoap.org/soap/envelope/ actor,attr"`
}

// TencodingStyle 'encodingStyle' indicates any canonicalization conventions followed in the contents of the containing element.  For example, the value 'http://schemas.xmlsoap.org/soap/encoding/' indicates the pattern described in SOAP specification
type TencodingStyle xsdt.String

// Set : Since TencodingStyle is just a simple String type, this merely sets the current value from the specified string.
func (me *TencodingStyle) Set(s string) { (*xsdt.String)(me).Set(s) }

// String : Since TencodingStyle is just a simple String type, this merely returns the current string value.
func (me TencodingStyle) String() string { return xsdt.String(me).String() }

// ToXsdtString : This convenience method just performs a simple type conversion to TencodingStyle's alias type xsdt.String.
func (me TencodingStyle) ToXsdtString() xsdt.String { return xsdt.String(me) }

// Values : TencodingStyle declares a String containing a whitespace-separated list of xsdt.AnyURI values. This Values() method creates and returns a slice of all elements in that list.
func (me TencodingStyle) Values() (list []xsdt.AnyURI) {
	svals := xsdt.ListValues(string(me))
	list = make([]xsdt.AnyURI, len(svals))
	for i, s := range svals {
		list[i].Set(s)
	}
	return
}

// XsdGoPkgHasAttrEncodingStyle defines attribute EncodingStyle
type XsdGoPkgHasAttrEncodingStyle struct {
	EncodingStyle TencodingStyle `xml:"http://schemas.xmlsoap.org/soap/envelope/ encodingStyle,attr"`
}

// XsdGoPkgHasAttsEncodingStyle defines XsdGoPkgHasAttsEncodingStyle
type XsdGoPkgHasAttsEncodingStyle struct {
	XsdGoPkgHasAttrEncodingStyle
}

// THeader defines THeader
type THeader struct {
}

// Walk : if the WalkHandlers.THeader function is not nil (ie. was set by outside code), calls it with this THeader instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/0 field(s) belonging to this THeader instance.
func (me *THeader) Walk() (err error) {
	if fn := WalkHandlers.THeader; me != nil {
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

// XsdGoPkgHasElemHeader defines XsdGoPkgHasElemHeader
type XsdGoPkgHasElemHeader struct {
	Header *THeader `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
}

// Walk : if the WalkHandlers.XsdGoPkgHasElemHeader function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElemHeader instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 1/1 field(s) belonging to this XsdGoPkgHasElemHeader instance.
func (me *XsdGoPkgHasElemHeader) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElemHeader; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.Header.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
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

// TBody defines TBody
type TBody struct {
}

// Walk : if the WalkHandlers.TBody function is not nil (ie. was set by outside code), calls it with this TBody instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/0 field(s) belonging to this TBody instance.
func (me *TBody) Walk() (err error) {
	if fn := WalkHandlers.TBody; me != nil {
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

// XsdGoPkgHasElemBody defines XsdGoPkgHasElemBody
type XsdGoPkgHasElemBody struct {
	Body *TBody `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
}

// Walk : if the WalkHandlers.XsdGoPkgHasElemBody function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElemBody instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 1/1 field(s) belonging to this XsdGoPkgHasElemBody instance.
func (me *XsdGoPkgHasElemBody) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElemBody; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.Body.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
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

// TEnvelope defines TEnvelope
type TEnvelope struct {
	XsdGoPkgHasElemHeader

	XsdGoPkgHasElemBody
}

// Walk : if the WalkHandlers.TEnvelope function is not nil (ie. was set by outside code), calls it with this TEnvelope instance as the single argument. Then calls the Walk() method on 2/2 embed(s) and 0/0 field(s) belonging to this TEnvelope instance.
func (me *TEnvelope) Walk() (err error) {
	if fn := WalkHandlers.TEnvelope; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.XsdGoPkgHasElemHeader.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if err = me.XsdGoPkgHasElemBody.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
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

// XsdGoPkgHasElemEnvelope defines XsdGoPkgHasElemEnvelope
type XsdGoPkgHasElemEnvelope struct {
	Envelope *TEnvelope `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
}

// Walk : if the WalkHandlers.XsdGoPkgHasElemEnvelope function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElemEnvelope instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 1/1 field(s) belonging to this XsdGoPkgHasElemEnvelope instance.
func (me *XsdGoPkgHasElemEnvelope) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElemEnvelope; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.Envelope.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
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

// XsdGoPkgHasElemsEnvelope defines XsdGoPkgHasElemsEnvelope
type XsdGoPkgHasElemsEnvelope struct {
	Envelopes []*TEnvelope `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
}

// Walk : if the WalkHandlers.XsdGoPkgHasElemsEnvelope function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElemsEnvelope instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElemsEnvelope instance.
func (me *XsdGoPkgHasElemsEnvelope) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElemsEnvelope; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		for _, x := range me.Envelopes {
			if err = x.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
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

// XsdGoPkgHasElemsHeader defines XsdGoPkgHasElemsHeader
type XsdGoPkgHasElemsHeader struct {
	Headers []*THeader `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
}

// Walk : if the WalkHandlers.XsdGoPkgHasElemsHeader function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElemsHeader instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElemsHeader instance.
func (me *XsdGoPkgHasElemsHeader) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElemsHeader; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		for _, x := range me.Headers {
			if err = x.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
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

// XsdGoPkgHasElemsBody defines XsdGoPkgHasElemsBody
type XsdGoPkgHasElemsBody struct {
	Bodies []*TBody `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
}

// Walk : if the WalkHandlers.XsdGoPkgHasElemsBody function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElemsBody instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElemsBody instance.
func (me *XsdGoPkgHasElemsBody) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElemsBody; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		for _, x := range me.Bodies {
			if err = x.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
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

// XsdGoPkgHasElemFaultcodesequenceFaultschemaFaultcodeXsdtQName Fault reporting structure
type XsdGoPkgHasElemFaultcodesequenceFaultschemaFaultcodeXsdtQName struct {
	Faultcode xsdt.QName `xml:"http://schemas.xmlsoap.org/soap/envelope/ faultcode"`
}

// Walk : if the WalkHandlers.XsdGoPkgHasElemFaultcodesequenceFaultschemaFaultcodeXsdtQName function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElemFaultcodesequenceFaultschemaFaultcodeXsdtQName instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElemFaultcodesequenceFaultschemaFaultcodeXsdtQName instance.
func (me *XsdGoPkgHasElemFaultcodesequenceFaultschemaFaultcodeXsdtQName) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElemFaultcodesequenceFaultschemaFaultcodeXsdtQName; me != nil {
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

// XsdGoPkgHasElemFaultstringsequenceFaultschemaFaultstringXsdtString defines XsdGoPkgHasElemFaultstringsequenceFaultschemaFaultstringXsdtString
type XsdGoPkgHasElemFaultstringsequenceFaultschemaFaultstringXsdtString struct {
	Faultstring xsdt.String `xml:"http://schemas.xmlsoap.org/soap/envelope/ faultstring"`
}

// Walk : if the WalkHandlers.XsdGoPkgHasElemFaultstringsequenceFaultschemaFaultstringXsdtString function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElemFaultstringsequenceFaultschemaFaultstringXsdtString instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElemFaultstringsequenceFaultschemaFaultstringXsdtString instance.
func (me *XsdGoPkgHasElemFaultstringsequenceFaultschemaFaultstringXsdtString) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElemFaultstringsequenceFaultschemaFaultstringXsdtString; me != nil {
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

// XsdGoPkgHasElemFaultactorsequenceFaultschemaFaultactorXsdtAnyURI defines XsdGoPkgHasElemFaultactorsequenceFaultschemaFaultactorXsdtAnyURI
type XsdGoPkgHasElemFaultactorsequenceFaultschemaFaultactorXsdtAnyURI struct {
	Faultactor xsdt.AnyURI `xml:"http://schemas.xmlsoap.org/soap/envelope/ faultactor"`
}

// Walk : if the WalkHandlers.XsdGoPkgHasElemFaultactorsequenceFaultschemaFaultactorXsdtAnyURI function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElemFaultactorsequenceFaultschemaFaultactorXsdtAnyURI instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElemFaultactorsequenceFaultschemaFaultactorXsdtAnyURI instance.
func (me *XsdGoPkgHasElemFaultactorsequenceFaultschemaFaultactorXsdtAnyURI) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElemFaultactorsequenceFaultschemaFaultactorXsdtAnyURI; me != nil {
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

// Tdetail defines Tdetail
type Tdetail struct {
}

// Walk : if the WalkHandlers.Tdetail function is not nil (ie. was set by outside code), calls it with this Tdetail instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/0 field(s) belonging to this Tdetail instance.
func (me *Tdetail) Walk() (err error) {
	if fn := WalkHandlers.Tdetail; me != nil {
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

// XsdGoPkgHasElemDetailsequenceFaultschemaDetailTdetail defines XsdGoPkgHasElemDetailsequenceFaultschemaDetailTdetail
type XsdGoPkgHasElemDetailsequenceFaultschemaDetailTdetail struct {
	Detail *Tdetail `xml:"http://schemas.xmlsoap.org/soap/envelope/ detail"`
}

// Walk : if the WalkHandlers.XsdGoPkgHasElemDetailsequenceFaultschemaDetailTdetail function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElemDetailsequenceFaultschemaDetailTdetail instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 1/1 field(s) belonging to this XsdGoPkgHasElemDetailsequenceFaultschemaDetailTdetail instance.
func (me *XsdGoPkgHasElemDetailsequenceFaultschemaDetailTdetail) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElemDetailsequenceFaultschemaDetailTdetail; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.Detail.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
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

// TFault defines TFault
type TFault struct {
	XsdGoPkgHasElemFaultcodesequenceFaultschemaFaultcodeXsdtQName

	XsdGoPkgHasElemFaultstringsequenceFaultschemaFaultstringXsdtString

	XsdGoPkgHasElemFaultactorsequenceFaultschemaFaultactorXsdtAnyURI

	XsdGoPkgHasElemDetailsequenceFaultschemaDetailTdetail
}

// Walk : if the WalkHandlers.TFault function is not nil (ie. was set by outside code), calls it with this TFault instance as the single argument. Then calls the Walk() method on 4/4 embed(s) and 0/0 field(s) belonging to this TFault instance.
func (me *TFault) Walk() (err error) {
	if fn := WalkHandlers.TFault; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.XsdGoPkgHasElemFaultcodesequenceFaultschemaFaultcodeXsdtQName.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if err = me.XsdGoPkgHasElemFaultstringsequenceFaultschemaFaultstringXsdtString.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if err = me.XsdGoPkgHasElemFaultactorsequenceFaultschemaFaultactorXsdtAnyURI.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if err = me.XsdGoPkgHasElemDetailsequenceFaultschemaDetailTdetail.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
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

// XsdGoPkgHasElemFault defines XsdGoPkgHasElemFault
type XsdGoPkgHasElemFault struct {
	Fault *TFault `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`
}

// Walk : if the WalkHandlers.XsdGoPkgHasElemFault function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElemFault instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 1/1 field(s) belonging to this XsdGoPkgHasElemFault instance.
func (me *XsdGoPkgHasElemFault) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElemFault; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.Fault.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
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

// XsdGoPkgHasElemsFault defines XsdGoPkgHasElemsFault
type XsdGoPkgHasElemsFault struct {
	Faults []*TFault `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`
}

// Walk : if the WalkHandlers.XsdGoPkgHasElemsFault function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElemsFault instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElemsFault instance.
func (me *XsdGoPkgHasElemsFault) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElemsFault; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		for _, x := range me.Faults {
			if err = x.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
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

// XsdGoPkgHasElemsDetailsequenceFaultschemaDetailTdetail defines XsdGoPkgHasElemsDetailsequenceFaultschemaDetailTdetail
type XsdGoPkgHasElemsDetailsequenceFaultschemaDetailTdetail struct {
	Details []*Tdetail `xml:"http://schemas.xmlsoap.org/soap/envelope/ detail"`
}

// Walk : if the WalkHandlers.XsdGoPkgHasElemsDetailsequenceFaultschemaDetailTdetail function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElemsDetailsequenceFaultschemaDetailTdetail instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElemsDetailsequenceFaultschemaDetailTdetail instance.
func (me *XsdGoPkgHasElemsDetailsequenceFaultschemaDetailTdetail) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElemsDetailsequenceFaultschemaDetailTdetail; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		for _, x := range me.Details {
			if err = x.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
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

// XsdGoPkgHasElemsFaultstringsequenceFaultschemaFaultstringXsdtString defines XsdGoPkgHasElemsFaultstringsequenceFaultschemaFaultstringXsdtString
type XsdGoPkgHasElemsFaultstringsequenceFaultschemaFaultstringXsdtString struct {
	Faultstrings []xsdt.String `xml:"http://schemas.xmlsoap.org/soap/envelope/ faultstring"`
}

// Walk : if the WalkHandlers.XsdGoPkgHasElemsFaultstringsequenceFaultschemaFaultstringXsdtString function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElemsFaultstringsequenceFaultschemaFaultstringXsdtString instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElemsFaultstringsequenceFaultschemaFaultstringXsdtString instance.
func (me *XsdGoPkgHasElemsFaultstringsequenceFaultschemaFaultstringXsdtString) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElemsFaultstringsequenceFaultschemaFaultstringXsdtString; me != nil {
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

// XsdGoPkgHasElemsFaultactorsequenceFaultschemaFaultactorXsdtAnyURI defines XsdGoPkgHasElemsFaultactorsequenceFaultschemaFaultactorXsdtAnyURI
type XsdGoPkgHasElemsFaultactorsequenceFaultschemaFaultactorXsdtAnyURI struct {
	Faultactors []xsdt.AnyURI `xml:"http://schemas.xmlsoap.org/soap/envelope/ faultactor"`
}

// Walk : if the WalkHandlers.XsdGoPkgHasElemsFaultactorsequenceFaultschemaFaultactorXsdtAnyURI function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElemsFaultactorsequenceFaultschemaFaultactorXsdtAnyURI instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElemsFaultactorsequenceFaultschemaFaultactorXsdtAnyURI instance.
func (me *XsdGoPkgHasElemsFaultactorsequenceFaultschemaFaultactorXsdtAnyURI) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElemsFaultactorsequenceFaultschemaFaultactorXsdtAnyURI; me != nil {
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

// XsdGoPkgHasElemsFaultcodesequenceFaultschemaFaultcodeXsdtQName defines XsdGoPkgHasElemsFaultcodesequenceFaultschemaFaultcodeXsdtQName
type XsdGoPkgHasElemsFaultcodesequenceFaultschemaFaultcodeXsdtQName struct {
	Faultcodes []xsdt.QName `xml:"http://schemas.xmlsoap.org/soap/envelope/ faultcode"`
}

// Walk : if the WalkHandlers.XsdGoPkgHasElemsFaultcodesequenceFaultschemaFaultcodeXsdtQName function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElemsFaultcodesequenceFaultschemaFaultcodeXsdtQName instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElemsFaultcodesequenceFaultschemaFaultcodeXsdtQName instance.
func (me *XsdGoPkgHasElemsFaultcodesequenceFaultschemaFaultcodeXsdtQName) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElemsFaultcodesequenceFaultschemaFaultcodeXsdtQName; me != nil {
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
	// WalkHandlers Provides 22 strong-typed hooks for your own custom handler functions to be invoked when the Walk() method is called on any instance of any (non-attribute-related) struct type defined in this package.
	// If your custom handler does get called at all for a given struct instance, then it always gets called twice, first with the 'enter' bool argument set to true, then (after having Walk()ed all subordinate struct instances, if any) once again with it set to false.
	WalkHandlers = &XsdGoPkgWalkHandlers{}
)

// XsdGoPkgWalkHandlers Provides 22 strong-typed hooks for your own custom handler functions to be invoked when the Walk() method is called on any instance of any (non-attribute-related) struct type defined in this package.
// If your custom handler does get called at all for a given struct instance, then it always gets called twice, first with the 'enter' bool argument set to true, then (after having Walk()ed all subordinate struct instances, if any) once again with it set to false.
type XsdGoPkgWalkHandlers struct {
	THeader                                                             func(*THeader, bool) error
	TBody                                                               func(*TBody, bool) error
	XsdGoPkgHasElemFaultcodesequenceFaultschemaFaultcodeXsdtQName       func(*XsdGoPkgHasElemFaultcodesequenceFaultschemaFaultcodeXsdtQName, bool) error
	Tdetail                                                             func(*Tdetail, bool) error
	XsdGoPkgHasElemFault                                                func(*XsdGoPkgHasElemFault, bool) error
	XsdGoPkgHasElemsFault                                               func(*XsdGoPkgHasElemsFault, bool) error
	XsdGoPkgHasElemHeader                                               func(*XsdGoPkgHasElemHeader, bool) error
	XsdGoPkgHasElemBody                                                 func(*XsdGoPkgHasElemBody, bool) error
	XsdGoPkgHasElemsHeader                                              func(*XsdGoPkgHasElemsHeader, bool) error
	XsdGoPkgHasCdata                                                    func(*XsdGoPkgHasCdata, bool) error
	XsdGoPkgHasElemsFaultactorsequenceFaultschemaFaultactorXsdtAnyURI   func(*XsdGoPkgHasElemsFaultactorsequenceFaultschemaFaultactorXsdtAnyURI, bool) error
	XsdGoPkgHasElemsFaultcodesequenceFaultschemaFaultcodeXsdtQName      func(*XsdGoPkgHasElemsFaultcodesequenceFaultschemaFaultcodeXsdtQName, bool) error
	XsdGoPkgHasElemsEnvelope                                            func(*XsdGoPkgHasElemsEnvelope, bool) error
	XsdGoPkgHasElemsBody                                                func(*XsdGoPkgHasElemsBody, bool) error
	XsdGoPkgHasElemDetailsequenceFaultschemaDetailTdetail               func(*XsdGoPkgHasElemDetailsequenceFaultschemaDetailTdetail, bool) error
	TFault                                                              func(*TFault, bool) error
	XsdGoPkgHasElemsDetailsequenceFaultschemaDetailTdetail              func(*XsdGoPkgHasElemsDetailsequenceFaultschemaDetailTdetail, bool) error
	XsdGoPkgHasElemsFaultstringsequenceFaultschemaFaultstringXsdtString func(*XsdGoPkgHasElemsFaultstringsequenceFaultschemaFaultstringXsdtString, bool) error
	TEnvelope                                                           func(*TEnvelope, bool) error
	XsdGoPkgHasElemEnvelope                                             func(*XsdGoPkgHasElemEnvelope, bool) error
	XsdGoPkgHasElemFaultstringsequenceFaultschemaFaultstringXsdtString  func(*XsdGoPkgHasElemFaultstringsequenceFaultschemaFaultstringXsdtString, bool) error
	XsdGoPkgHasElemFaultactorsequenceFaultschemaFaultactorXsdtAnyURI    func(*XsdGoPkgHasElemFaultactorsequenceFaultschemaFaultactorXsdtAnyURI, bool) error
}
