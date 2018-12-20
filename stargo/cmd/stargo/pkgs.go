package main

import (
	"go.starlark.net/stargo"
	"go.starlark.net/starlark"

	π1 "bytes"
	π2 "encoding/json"
	π0 "fmt"
	π6 "go/ast"
	π7 "go/parser"
	π5 "go/token"
	π8 "go/types"
	π3 "io/ioutil"
	π4 "net/http"
)

var goPackages = starlark.StringDict{
	"fmt": &starlark.Module{
		Name: "fmt",
		Members: starlark.StringDict{
			"Errorf":     stargo.ValueOf(π0.Errorf),
			"Formatter":  stargo.TypeOf(new(π0.Formatter)),
			"Fprint":     stargo.ValueOf(π0.Fprint),
			"Fprintf":    stargo.ValueOf(π0.Fprintf),
			"Fprintln":   stargo.ValueOf(π0.Fprintln),
			"Fscan":      stargo.ValueOf(π0.Fscan),
			"Fscanf":     stargo.ValueOf(π0.Fscanf),
			"Fscanln":    stargo.ValueOf(π0.Fscanln),
			"GoStringer": stargo.TypeOf(new(π0.GoStringer)),
			"Print":      stargo.ValueOf(π0.Print),
			"Printf":     stargo.ValueOf(π0.Printf),
			"Println":    stargo.ValueOf(π0.Println),
			"Scan":       stargo.ValueOf(π0.Scan),
			"ScanState":  stargo.TypeOf(new(π0.ScanState)),
			"Scanf":      stargo.ValueOf(π0.Scanf),
			"Scanln":     stargo.ValueOf(π0.Scanln),
			"Scanner":    stargo.TypeOf(new(π0.Scanner)),
			"Sprint":     stargo.ValueOf(π0.Sprint),
			"Sprintf":    stargo.ValueOf(π0.Sprintf),
			"Sprintln":   stargo.ValueOf(π0.Sprintln),
			"Sscan":      stargo.ValueOf(π0.Sscan),
			"Sscanf":     stargo.ValueOf(π0.Sscanf),
			"Sscanln":    stargo.ValueOf(π0.Sscanln),
			"State":      stargo.TypeOf(new(π0.State)),
			"Stringer":   stargo.TypeOf(new(π0.Stringer)),
		},
	},
	"bytes": &starlark.Module{
		Name: "bytes",
		Members: starlark.StringDict{
			"Buffer":          stargo.TypeOf(new(π1.Buffer)),
			"Compare":         stargo.ValueOf(π1.Compare),
			"Contains":        stargo.ValueOf(π1.Contains),
			"ContainsAny":     stargo.ValueOf(π1.ContainsAny),
			"ContainsRune":    stargo.ValueOf(π1.ContainsRune),
			"Count":           stargo.ValueOf(π1.Count),
			"Equal":           stargo.ValueOf(π1.Equal),
			"EqualFold":       stargo.ValueOf(π1.EqualFold),
			"ErrTooLarge":     stargo.VarOf(&π1.ErrTooLarge),
			"Fields":          stargo.ValueOf(π1.Fields),
			"FieldsFunc":      stargo.ValueOf(π1.FieldsFunc),
			"HasPrefix":       stargo.ValueOf(π1.HasPrefix),
			"HasSuffix":       stargo.ValueOf(π1.HasSuffix),
			"Index":           stargo.ValueOf(π1.Index),
			"IndexAny":        stargo.ValueOf(π1.IndexAny),
			"IndexByte":       stargo.ValueOf(π1.IndexByte),
			"IndexFunc":       stargo.ValueOf(π1.IndexFunc),
			"IndexRune":       stargo.ValueOf(π1.IndexRune),
			"Join":            stargo.ValueOf(π1.Join),
			"LastIndex":       stargo.ValueOf(π1.LastIndex),
			"LastIndexAny":    stargo.ValueOf(π1.LastIndexAny),
			"LastIndexByte":   stargo.ValueOf(π1.LastIndexByte),
			"LastIndexFunc":   stargo.ValueOf(π1.LastIndexFunc),
			"Map":             stargo.ValueOf(π1.Map),
			"MinRead":         stargo.ValueOf(π1.MinRead),
			"NewBuffer":       stargo.ValueOf(π1.NewBuffer),
			"NewBufferString": stargo.ValueOf(π1.NewBufferString),
			"NewReader":       stargo.ValueOf(π1.NewReader),
			"Reader":          stargo.TypeOf(new(π1.Reader)),
			"Repeat":          stargo.ValueOf(π1.Repeat),
			"Replace":         stargo.ValueOf(π1.Replace),
			"ReplaceAll":      stargo.ValueOf(π1.ReplaceAll),
			"Runes":           stargo.ValueOf(π1.Runes),
			"Split":           stargo.ValueOf(π1.Split),
			"SplitAfter":      stargo.ValueOf(π1.SplitAfter),
			"SplitAfterN":     stargo.ValueOf(π1.SplitAfterN),
			"SplitN":          stargo.ValueOf(π1.SplitN),
			"Title":           stargo.ValueOf(π1.Title),
			"ToLower":         stargo.ValueOf(π1.ToLower),
			"ToLowerSpecial":  stargo.ValueOf(π1.ToLowerSpecial),
			"ToTitle":         stargo.ValueOf(π1.ToTitle),
			"ToTitleSpecial":  stargo.ValueOf(π1.ToTitleSpecial),
			"ToUpper":         stargo.ValueOf(π1.ToUpper),
			"ToUpperSpecial":  stargo.ValueOf(π1.ToUpperSpecial),
			"Trim":            stargo.ValueOf(π1.Trim),
			"TrimFunc":        stargo.ValueOf(π1.TrimFunc),
			"TrimLeft":        stargo.ValueOf(π1.TrimLeft),
			"TrimLeftFunc":    stargo.ValueOf(π1.TrimLeftFunc),
			"TrimPrefix":      stargo.ValueOf(π1.TrimPrefix),
			"TrimRight":       stargo.ValueOf(π1.TrimRight),
			"TrimRightFunc":   stargo.ValueOf(π1.TrimRightFunc),
			"TrimSpace":       stargo.ValueOf(π1.TrimSpace),
			"TrimSuffix":      stargo.ValueOf(π1.TrimSuffix),
		},
	},
	"encoding/json": &starlark.Module{
		Name: "encoding/json",
		Members: starlark.StringDict{
			"Compact":               stargo.ValueOf(π2.Compact),
			"Decoder":               stargo.TypeOf(new(π2.Decoder)),
			"Delim":                 stargo.TypeOf(new(π2.Delim)),
			"Encoder":               stargo.TypeOf(new(π2.Encoder)),
			"HTMLEscape":            stargo.ValueOf(π2.HTMLEscape),
			"Indent":                stargo.ValueOf(π2.Indent),
			"InvalidUTF8Error":      stargo.TypeOf(new(π2.InvalidUTF8Error)),
			"InvalidUnmarshalError": stargo.TypeOf(new(π2.InvalidUnmarshalError)),
			"Marshal":               stargo.ValueOf(π2.Marshal),
			"MarshalIndent":         stargo.ValueOf(π2.MarshalIndent),
			"Marshaler":             stargo.TypeOf(new(π2.Marshaler)),
			"MarshalerError":        stargo.TypeOf(new(π2.MarshalerError)),
			"NewDecoder":            stargo.ValueOf(π2.NewDecoder),
			"NewEncoder":            stargo.ValueOf(π2.NewEncoder),
			"Number":                stargo.TypeOf(new(π2.Number)),
			"RawMessage":            stargo.TypeOf(new(π2.RawMessage)),
			"SyntaxError":           stargo.TypeOf(new(π2.SyntaxError)),
			"Token":                 stargo.TypeOf(new(π2.Token)),
			"Unmarshal":             stargo.ValueOf(π2.Unmarshal),
			"UnmarshalFieldError":   stargo.TypeOf(new(π2.UnmarshalFieldError)),
			"UnmarshalTypeError":    stargo.TypeOf(new(π2.UnmarshalTypeError)),
			"Unmarshaler":           stargo.TypeOf(new(π2.Unmarshaler)),
			"UnsupportedTypeError":  stargo.TypeOf(new(π2.UnsupportedTypeError)),
			"UnsupportedValueError": stargo.TypeOf(new(π2.UnsupportedValueError)),
			"Valid":                 stargo.ValueOf(π2.Valid),
		},
	},
	"io/ioutil": &starlark.Module{
		Name: "io/ioutil",
		Members: starlark.StringDict{
			"Discard":   stargo.VarOf(&π3.Discard),
			"NopCloser": stargo.ValueOf(π3.NopCloser),
			"ReadAll":   stargo.ValueOf(π3.ReadAll),
			"ReadDir":   stargo.ValueOf(π3.ReadDir),
			"ReadFile":  stargo.ValueOf(π3.ReadFile),
			"TempDir":   stargo.ValueOf(π3.TempDir),
			"TempFile":  stargo.ValueOf(π3.TempFile),
			"WriteFile": stargo.ValueOf(π3.WriteFile),
		},
	},
	"net/http": &starlark.Module{
		Name: "net/http",
		Members: starlark.StringDict{
			"CanonicalHeaderKey":                  stargo.ValueOf(π4.CanonicalHeaderKey),
			"Client":                              stargo.TypeOf(new(π4.Client)),
			"CloseNotifier":                       stargo.TypeOf(new(π4.CloseNotifier)),
			"ConnState":                           stargo.TypeOf(new(π4.ConnState)),
			"Cookie":                              stargo.TypeOf(new(π4.Cookie)),
			"CookieJar":                           stargo.TypeOf(new(π4.CookieJar)),
			"DefaultClient":                       stargo.VarOf(&π4.DefaultClient),
			"DefaultMaxHeaderBytes":               stargo.ValueOf(π4.DefaultMaxHeaderBytes),
			"DefaultMaxIdleConnsPerHost":          stargo.ValueOf(π4.DefaultMaxIdleConnsPerHost),
			"DefaultServeMux":                     stargo.VarOf(&π4.DefaultServeMux),
			"DefaultTransport":                    stargo.VarOf(&π4.DefaultTransport),
			"DetectContentType":                   stargo.ValueOf(π4.DetectContentType),
			"Dir":                                 stargo.TypeOf(new(π4.Dir)),
			"ErrAbortHandler":                     stargo.VarOf(&π4.ErrAbortHandler),
			"ErrBodyNotAllowed":                   stargo.VarOf(&π4.ErrBodyNotAllowed),
			"ErrBodyReadAfterClose":               stargo.VarOf(&π4.ErrBodyReadAfterClose),
			"ErrContentLength":                    stargo.VarOf(&π4.ErrContentLength),
			"ErrHandlerTimeout":                   stargo.VarOf(&π4.ErrHandlerTimeout),
			"ErrHeaderTooLong":                    stargo.VarOf(&π4.ErrHeaderTooLong),
			"ErrHijacked":                         stargo.VarOf(&π4.ErrHijacked),
			"ErrLineTooLong":                      stargo.VarOf(&π4.ErrLineTooLong),
			"ErrMissingBoundary":                  stargo.VarOf(&π4.ErrMissingBoundary),
			"ErrMissingContentLength":             stargo.VarOf(&π4.ErrMissingContentLength),
			"ErrMissingFile":                      stargo.VarOf(&π4.ErrMissingFile),
			"ErrNoCookie":                         stargo.VarOf(&π4.ErrNoCookie),
			"ErrNoLocation":                       stargo.VarOf(&π4.ErrNoLocation),
			"ErrNotMultipart":                     stargo.VarOf(&π4.ErrNotMultipart),
			"ErrNotSupported":                     stargo.VarOf(&π4.ErrNotSupported),
			"ErrServerClosed":                     stargo.VarOf(&π4.ErrServerClosed),
			"ErrShortBody":                        stargo.VarOf(&π4.ErrShortBody),
			"ErrSkipAltProtocol":                  stargo.VarOf(&π4.ErrSkipAltProtocol),
			"ErrUnexpectedTrailer":                stargo.VarOf(&π4.ErrUnexpectedTrailer),
			"ErrUseLastResponse":                  stargo.VarOf(&π4.ErrUseLastResponse),
			"ErrWriteAfterFlush":                  stargo.VarOf(&π4.ErrWriteAfterFlush),
			"Error":                               stargo.ValueOf(π4.Error),
			"File":                                stargo.TypeOf(new(π4.File)),
			"FileServer":                          stargo.ValueOf(π4.FileServer),
			"FileSystem":                          stargo.TypeOf(new(π4.FileSystem)),
			"Flusher":                             stargo.TypeOf(new(π4.Flusher)),
			"Get":                                 stargo.ValueOf(π4.Get),
			"Handle":                              stargo.ValueOf(π4.Handle),
			"HandleFunc":                          stargo.ValueOf(π4.HandleFunc),
			"Handler":                             stargo.TypeOf(new(π4.Handler)),
			"HandlerFunc":                         stargo.TypeOf(new(π4.HandlerFunc)),
			"Head":                                stargo.ValueOf(π4.Head),
			"Header":                              stargo.TypeOf(new(π4.Header)),
			"Hijacker":                            stargo.TypeOf(new(π4.Hijacker)),
			"ListenAndServe":                      stargo.ValueOf(π4.ListenAndServe),
			"ListenAndServeTLS":                   stargo.ValueOf(π4.ListenAndServeTLS),
			"LocalAddrContextKey":                 stargo.VarOf(&π4.LocalAddrContextKey),
			"MaxBytesReader":                      stargo.ValueOf(π4.MaxBytesReader),
			"MethodConnect":                       stargo.ValueOf(π4.MethodConnect),
			"MethodDelete":                        stargo.ValueOf(π4.MethodDelete),
			"MethodGet":                           stargo.ValueOf(π4.MethodGet),
			"MethodHead":                          stargo.ValueOf(π4.MethodHead),
			"MethodOptions":                       stargo.ValueOf(π4.MethodOptions),
			"MethodPatch":                         stargo.ValueOf(π4.MethodPatch),
			"MethodPost":                          stargo.ValueOf(π4.MethodPost),
			"MethodPut":                           stargo.ValueOf(π4.MethodPut),
			"MethodTrace":                         stargo.ValueOf(π4.MethodTrace),
			"NewFileTransport":                    stargo.ValueOf(π4.NewFileTransport),
			"NewRequest":                          stargo.ValueOf(π4.NewRequest),
			"NewServeMux":                         stargo.ValueOf(π4.NewServeMux),
			"NoBody":                              stargo.VarOf(&π4.NoBody),
			"NotFound":                            stargo.ValueOf(π4.NotFound),
			"NotFoundHandler":                     stargo.ValueOf(π4.NotFoundHandler),
			"ParseHTTPVersion":                    stargo.ValueOf(π4.ParseHTTPVersion),
			"ParseTime":                           stargo.ValueOf(π4.ParseTime),
			"Post":                                stargo.ValueOf(π4.Post),
			"PostForm":                            stargo.ValueOf(π4.PostForm),
			"ProtocolError":                       stargo.TypeOf(new(π4.ProtocolError)),
			"ProxyFromEnvironment":                stargo.ValueOf(π4.ProxyFromEnvironment),
			"ProxyURL":                            stargo.ValueOf(π4.ProxyURL),
			"PushOptions":                         stargo.TypeOf(new(π4.PushOptions)),
			"Pusher":                              stargo.TypeOf(new(π4.Pusher)),
			"ReadRequest":                         stargo.ValueOf(π4.ReadRequest),
			"ReadResponse":                        stargo.ValueOf(π4.ReadResponse),
			"Redirect":                            stargo.ValueOf(π4.Redirect),
			"RedirectHandler":                     stargo.ValueOf(π4.RedirectHandler),
			"Request":                             stargo.TypeOf(new(π4.Request)),
			"Response":                            stargo.TypeOf(new(π4.Response)),
			"ResponseWriter":                      stargo.TypeOf(new(π4.ResponseWriter)),
			"RoundTripper":                        stargo.TypeOf(new(π4.RoundTripper)),
			"SameSite":                            stargo.TypeOf(new(π4.SameSite)),
			"SameSiteDefaultMode":                 stargo.ValueOf(π4.SameSiteDefaultMode),
			"SameSiteLaxMode":                     stargo.ValueOf(π4.SameSiteLaxMode),
			"SameSiteStrictMode":                  stargo.ValueOf(π4.SameSiteStrictMode),
			"Serve":                               stargo.ValueOf(π4.Serve),
			"ServeContent":                        stargo.ValueOf(π4.ServeContent),
			"ServeFile":                           stargo.ValueOf(π4.ServeFile),
			"ServeMux":                            stargo.TypeOf(new(π4.ServeMux)),
			"ServeTLS":                            stargo.ValueOf(π4.ServeTLS),
			"Server":                              stargo.TypeOf(new(π4.Server)),
			"ServerContextKey":                    stargo.VarOf(&π4.ServerContextKey),
			"SetCookie":                           stargo.ValueOf(π4.SetCookie),
			"StateActive":                         stargo.ValueOf(π4.StateActive),
			"StateClosed":                         stargo.ValueOf(π4.StateClosed),
			"StateHijacked":                       stargo.ValueOf(π4.StateHijacked),
			"StateIdle":                           stargo.ValueOf(π4.StateIdle),
			"StateNew":                            stargo.ValueOf(π4.StateNew),
			"StatusAccepted":                      stargo.ValueOf(π4.StatusAccepted),
			"StatusAlreadyReported":               stargo.ValueOf(π4.StatusAlreadyReported),
			"StatusBadGateway":                    stargo.ValueOf(π4.StatusBadGateway),
			"StatusBadRequest":                    stargo.ValueOf(π4.StatusBadRequest),
			"StatusConflict":                      stargo.ValueOf(π4.StatusConflict),
			"StatusContinue":                      stargo.ValueOf(π4.StatusContinue),
			"StatusCreated":                       stargo.ValueOf(π4.StatusCreated),
			"StatusExpectationFailed":             stargo.ValueOf(π4.StatusExpectationFailed),
			"StatusFailedDependency":              stargo.ValueOf(π4.StatusFailedDependency),
			"StatusForbidden":                     stargo.ValueOf(π4.StatusForbidden),
			"StatusFound":                         stargo.ValueOf(π4.StatusFound),
			"StatusGatewayTimeout":                stargo.ValueOf(π4.StatusGatewayTimeout),
			"StatusGone":                          stargo.ValueOf(π4.StatusGone),
			"StatusHTTPVersionNotSupported":       stargo.ValueOf(π4.StatusHTTPVersionNotSupported),
			"StatusIMUsed":                        stargo.ValueOf(π4.StatusIMUsed),
			"StatusInsufficientStorage":           stargo.ValueOf(π4.StatusInsufficientStorage),
			"StatusInternalServerError":           stargo.ValueOf(π4.StatusInternalServerError),
			"StatusLengthRequired":                stargo.ValueOf(π4.StatusLengthRequired),
			"StatusLocked":                        stargo.ValueOf(π4.StatusLocked),
			"StatusLoopDetected":                  stargo.ValueOf(π4.StatusLoopDetected),
			"StatusMethodNotAllowed":              stargo.ValueOf(π4.StatusMethodNotAllowed),
			"StatusMisdirectedRequest":            stargo.ValueOf(π4.StatusMisdirectedRequest),
			"StatusMovedPermanently":              stargo.ValueOf(π4.StatusMovedPermanently),
			"StatusMultiStatus":                   stargo.ValueOf(π4.StatusMultiStatus),
			"StatusMultipleChoices":               stargo.ValueOf(π4.StatusMultipleChoices),
			"StatusNetworkAuthenticationRequired": stargo.ValueOf(π4.StatusNetworkAuthenticationRequired),
			"StatusNoContent":                     stargo.ValueOf(π4.StatusNoContent),
			"StatusNonAuthoritativeInfo":          stargo.ValueOf(π4.StatusNonAuthoritativeInfo),
			"StatusNotAcceptable":                 stargo.ValueOf(π4.StatusNotAcceptable),
			"StatusNotExtended":                   stargo.ValueOf(π4.StatusNotExtended),
			"StatusNotFound":                      stargo.ValueOf(π4.StatusNotFound),
			"StatusNotImplemented":                stargo.ValueOf(π4.StatusNotImplemented),
			"StatusNotModified":                   stargo.ValueOf(π4.StatusNotModified),
			"StatusOK":                            stargo.ValueOf(π4.StatusOK),
			"StatusPartialContent":                stargo.ValueOf(π4.StatusPartialContent),
			"StatusPaymentRequired":               stargo.ValueOf(π4.StatusPaymentRequired),
			"StatusPermanentRedirect":             stargo.ValueOf(π4.StatusPermanentRedirect),
			"StatusPreconditionFailed":            stargo.ValueOf(π4.StatusPreconditionFailed),
			"StatusPreconditionRequired":          stargo.ValueOf(π4.StatusPreconditionRequired),
			"StatusProcessing":                    stargo.ValueOf(π4.StatusProcessing),
			"StatusProxyAuthRequired":             stargo.ValueOf(π4.StatusProxyAuthRequired),
			"StatusRequestEntityTooLarge":         stargo.ValueOf(π4.StatusRequestEntityTooLarge),
			"StatusRequestHeaderFieldsTooLarge":   stargo.ValueOf(π4.StatusRequestHeaderFieldsTooLarge),
			"StatusRequestTimeout":                stargo.ValueOf(π4.StatusRequestTimeout),
			"StatusRequestURITooLong":             stargo.ValueOf(π4.StatusRequestURITooLong),
			"StatusRequestedRangeNotSatisfiable":  stargo.ValueOf(π4.StatusRequestedRangeNotSatisfiable),
			"StatusResetContent":                  stargo.ValueOf(π4.StatusResetContent),
			"StatusSeeOther":                      stargo.ValueOf(π4.StatusSeeOther),
			"StatusServiceUnavailable":            stargo.ValueOf(π4.StatusServiceUnavailable),
			"StatusSwitchingProtocols":            stargo.ValueOf(π4.StatusSwitchingProtocols),
			"StatusTeapot":                        stargo.ValueOf(π4.StatusTeapot),
			"StatusTemporaryRedirect":             stargo.ValueOf(π4.StatusTemporaryRedirect),
			"StatusText":                          stargo.ValueOf(π4.StatusText),
			"StatusTooEarly":                      stargo.ValueOf(π4.StatusTooEarly),
			"StatusTooManyRequests":               stargo.ValueOf(π4.StatusTooManyRequests),
			"StatusUnauthorized":                  stargo.ValueOf(π4.StatusUnauthorized),
			"StatusUnavailableForLegalReasons":    stargo.ValueOf(π4.StatusUnavailableForLegalReasons),
			"StatusUnprocessableEntity":           stargo.ValueOf(π4.StatusUnprocessableEntity),
			"StatusUnsupportedMediaType":          stargo.ValueOf(π4.StatusUnsupportedMediaType),
			"StatusUpgradeRequired":               stargo.ValueOf(π4.StatusUpgradeRequired),
			"StatusUseProxy":                      stargo.ValueOf(π4.StatusUseProxy),
			"StatusVariantAlsoNegotiates":         stargo.ValueOf(π4.StatusVariantAlsoNegotiates),
			"StripPrefix":                         stargo.ValueOf(π4.StripPrefix),
			"TimeFormat":                          stargo.ValueOf(π4.TimeFormat),
			"TimeoutHandler":                      stargo.ValueOf(π4.TimeoutHandler),
			"TrailerPrefix":                       stargo.ValueOf(π4.TrailerPrefix),
			"Transport":                           stargo.TypeOf(new(π4.Transport)),
		},
	},
	"go/token": &starlark.Module{
		Name: "go/token",
		Members: starlark.StringDict{
			"ADD":            stargo.ValueOf(π5.ADD),
			"ADD_ASSIGN":     stargo.ValueOf(π5.ADD_ASSIGN),
			"AND":            stargo.ValueOf(π5.AND),
			"AND_ASSIGN":     stargo.ValueOf(π5.AND_ASSIGN),
			"AND_NOT":        stargo.ValueOf(π5.AND_NOT),
			"AND_NOT_ASSIGN": stargo.ValueOf(π5.AND_NOT_ASSIGN),
			"ARROW":          stargo.ValueOf(π5.ARROW),
			"ASSIGN":         stargo.ValueOf(π5.ASSIGN),
			"BREAK":          stargo.ValueOf(π5.BREAK),
			"CASE":           stargo.ValueOf(π5.CASE),
			"CHAN":           stargo.ValueOf(π5.CHAN),
			"CHAR":           stargo.ValueOf(π5.CHAR),
			"COLON":          stargo.ValueOf(π5.COLON),
			"COMMA":          stargo.ValueOf(π5.COMMA),
			"COMMENT":        stargo.ValueOf(π5.COMMENT),
			"CONST":          stargo.ValueOf(π5.CONST),
			"CONTINUE":       stargo.ValueOf(π5.CONTINUE),
			"DEC":            stargo.ValueOf(π5.DEC),
			"DEFAULT":        stargo.ValueOf(π5.DEFAULT),
			"DEFER":          stargo.ValueOf(π5.DEFER),
			"DEFINE":         stargo.ValueOf(π5.DEFINE),
			"ELLIPSIS":       stargo.ValueOf(π5.ELLIPSIS),
			"ELSE":           stargo.ValueOf(π5.ELSE),
			"EOF":            stargo.ValueOf(π5.EOF),
			"EQL":            stargo.ValueOf(π5.EQL),
			"FALLTHROUGH":    stargo.ValueOf(π5.FALLTHROUGH),
			"FLOAT":          stargo.ValueOf(π5.FLOAT),
			"FOR":            stargo.ValueOf(π5.FOR),
			"FUNC":           stargo.ValueOf(π5.FUNC),
			"File":           stargo.TypeOf(new(π5.File)),
			"FileSet":        stargo.TypeOf(new(π5.FileSet)),
			"GEQ":            stargo.ValueOf(π5.GEQ),
			"GO":             stargo.ValueOf(π5.GO),
			"GOTO":           stargo.ValueOf(π5.GOTO),
			"GTR":            stargo.ValueOf(π5.GTR),
			"HighestPrec":    stargo.ValueOf(π5.HighestPrec),
			"IDENT":          stargo.ValueOf(π5.IDENT),
			"IF":             stargo.ValueOf(π5.IF),
			"ILLEGAL":        stargo.ValueOf(π5.ILLEGAL),
			"IMAG":           stargo.ValueOf(π5.IMAG),
			"IMPORT":         stargo.ValueOf(π5.IMPORT),
			"INC":            stargo.ValueOf(π5.INC),
			"INT":            stargo.ValueOf(π5.INT),
			"INTERFACE":      stargo.ValueOf(π5.INTERFACE),
			"LAND":           stargo.ValueOf(π5.LAND),
			"LBRACE":         stargo.ValueOf(π5.LBRACE),
			"LBRACK":         stargo.ValueOf(π5.LBRACK),
			"LEQ":            stargo.ValueOf(π5.LEQ),
			"LOR":            stargo.ValueOf(π5.LOR),
			"LPAREN":         stargo.ValueOf(π5.LPAREN),
			"LSS":            stargo.ValueOf(π5.LSS),
			"Lookup":         stargo.ValueOf(π5.Lookup),
			"LowestPrec":     stargo.ValueOf(π5.LowestPrec),
			"MAP":            stargo.ValueOf(π5.MAP),
			"MUL":            stargo.ValueOf(π5.MUL),
			"MUL_ASSIGN":     stargo.ValueOf(π5.MUL_ASSIGN),
			"NEQ":            stargo.ValueOf(π5.NEQ),
			"NOT":            stargo.ValueOf(π5.NOT),
			"NewFileSet":     stargo.ValueOf(π5.NewFileSet),
			"NoPos":          stargo.ValueOf(π5.NoPos),
			"OR":             stargo.ValueOf(π5.OR),
			"OR_ASSIGN":      stargo.ValueOf(π5.OR_ASSIGN),
			"PACKAGE":        stargo.ValueOf(π5.PACKAGE),
			"PERIOD":         stargo.ValueOf(π5.PERIOD),
			"Pos":            stargo.TypeOf(new(π5.Pos)),
			"Position":       stargo.TypeOf(new(π5.Position)),
			"QUO":            stargo.ValueOf(π5.QUO),
			"QUO_ASSIGN":     stargo.ValueOf(π5.QUO_ASSIGN),
			"RANGE":          stargo.ValueOf(π5.RANGE),
			"RBRACE":         stargo.ValueOf(π5.RBRACE),
			"RBRACK":         stargo.ValueOf(π5.RBRACK),
			"REM":            stargo.ValueOf(π5.REM),
			"REM_ASSIGN":     stargo.ValueOf(π5.REM_ASSIGN),
			"RETURN":         stargo.ValueOf(π5.RETURN),
			"RPAREN":         stargo.ValueOf(π5.RPAREN),
			"SELECT":         stargo.ValueOf(π5.SELECT),
			"SEMICOLON":      stargo.ValueOf(π5.SEMICOLON),
			"SHL":            stargo.ValueOf(π5.SHL),
			"SHL_ASSIGN":     stargo.ValueOf(π5.SHL_ASSIGN),
			"SHR":            stargo.ValueOf(π5.SHR),
			"SHR_ASSIGN":     stargo.ValueOf(π5.SHR_ASSIGN),
			"STRING":         stargo.ValueOf(π5.STRING),
			"STRUCT":         stargo.ValueOf(π5.STRUCT),
			"SUB":            stargo.ValueOf(π5.SUB),
			"SUB_ASSIGN":     stargo.ValueOf(π5.SUB_ASSIGN),
			"SWITCH":         stargo.ValueOf(π5.SWITCH),
			"TYPE":           stargo.ValueOf(π5.TYPE),
			"Token":          stargo.TypeOf(new(π5.Token)),
			"UnaryPrec":      stargo.ValueOf(π5.UnaryPrec),
			"VAR":            stargo.ValueOf(π5.VAR),
			"XOR":            stargo.ValueOf(π5.XOR),
			"XOR_ASSIGN":     stargo.ValueOf(π5.XOR_ASSIGN),
		},
	},
	"go/ast": &starlark.Module{
		Name: "go/ast",
		Members: starlark.StringDict{
			"ArrayType":                  stargo.TypeOf(new(π6.ArrayType)),
			"AssignStmt":                 stargo.TypeOf(new(π6.AssignStmt)),
			"Bad":                        stargo.ValueOf(π6.Bad),
			"BadDecl":                    stargo.TypeOf(new(π6.BadDecl)),
			"BadExpr":                    stargo.TypeOf(new(π6.BadExpr)),
			"BadStmt":                    stargo.TypeOf(new(π6.BadStmt)),
			"BasicLit":                   stargo.TypeOf(new(π6.BasicLit)),
			"BinaryExpr":                 stargo.TypeOf(new(π6.BinaryExpr)),
			"BlockStmt":                  stargo.TypeOf(new(π6.BlockStmt)),
			"BranchStmt":                 stargo.TypeOf(new(π6.BranchStmt)),
			"CallExpr":                   stargo.TypeOf(new(π6.CallExpr)),
			"CaseClause":                 stargo.TypeOf(new(π6.CaseClause)),
			"ChanDir":                    stargo.TypeOf(new(π6.ChanDir)),
			"ChanType":                   stargo.TypeOf(new(π6.ChanType)),
			"CommClause":                 stargo.TypeOf(new(π6.CommClause)),
			"Comment":                    stargo.TypeOf(new(π6.Comment)),
			"CommentGroup":               stargo.TypeOf(new(π6.CommentGroup)),
			"CommentMap":                 stargo.TypeOf(new(π6.CommentMap)),
			"CompositeLit":               stargo.TypeOf(new(π6.CompositeLit)),
			"Con":                        stargo.ValueOf(π6.Con),
			"Decl":                       stargo.TypeOf(new(π6.Decl)),
			"DeclStmt":                   stargo.TypeOf(new(π6.DeclStmt)),
			"DeferStmt":                  stargo.TypeOf(new(π6.DeferStmt)),
			"Ellipsis":                   stargo.TypeOf(new(π6.Ellipsis)),
			"EmptyStmt":                  stargo.TypeOf(new(π6.EmptyStmt)),
			"Expr":                       stargo.TypeOf(new(π6.Expr)),
			"ExprStmt":                   stargo.TypeOf(new(π6.ExprStmt)),
			"Field":                      stargo.TypeOf(new(π6.Field)),
			"FieldFilter":                stargo.TypeOf(new(π6.FieldFilter)),
			"FieldList":                  stargo.TypeOf(new(π6.FieldList)),
			"File":                       stargo.TypeOf(new(π6.File)),
			"FileExports":                stargo.ValueOf(π6.FileExports),
			"Filter":                     stargo.TypeOf(new(π6.Filter)),
			"FilterDecl":                 stargo.ValueOf(π6.FilterDecl),
			"FilterFile":                 stargo.ValueOf(π6.FilterFile),
			"FilterFuncDuplicates":       stargo.ValueOf(π6.FilterFuncDuplicates),
			"FilterImportDuplicates":     stargo.ValueOf(π6.FilterImportDuplicates),
			"FilterPackage":              stargo.ValueOf(π6.FilterPackage),
			"FilterUnassociatedComments": stargo.ValueOf(π6.FilterUnassociatedComments),
			"ForStmt":                    stargo.TypeOf(new(π6.ForStmt)),
			"Fprint":                     stargo.ValueOf(π6.Fprint),
			"Fun":                        stargo.ValueOf(π6.Fun),
			"FuncDecl":                   stargo.TypeOf(new(π6.FuncDecl)),
			"FuncLit":                    stargo.TypeOf(new(π6.FuncLit)),
			"FuncType":                   stargo.TypeOf(new(π6.FuncType)),
			"GenDecl":                    stargo.TypeOf(new(π6.GenDecl)),
			"GoStmt":                     stargo.TypeOf(new(π6.GoStmt)),
			"Ident":                      stargo.TypeOf(new(π6.Ident)),
			"IfStmt":                     stargo.TypeOf(new(π6.IfStmt)),
			"ImportSpec":                 stargo.TypeOf(new(π6.ImportSpec)),
			"Importer":                   stargo.TypeOf(new(π6.Importer)),
			"IncDecStmt":                 stargo.TypeOf(new(π6.IncDecStmt)),
			"IndexExpr":                  stargo.TypeOf(new(π6.IndexExpr)),
			"Inspect":                    stargo.ValueOf(π6.Inspect),
			"InterfaceType":              stargo.TypeOf(new(π6.InterfaceType)),
			"IsExported":                 stargo.ValueOf(π6.IsExported),
			"KeyValueExpr":               stargo.TypeOf(new(π6.KeyValueExpr)),
			"LabeledStmt":                stargo.TypeOf(new(π6.LabeledStmt)),
			"Lbl":                        stargo.ValueOf(π6.Lbl),
			"MapType":                    stargo.TypeOf(new(π6.MapType)),
			"MergeMode":                  stargo.TypeOf(new(π6.MergeMode)),
			"MergePackageFiles":          stargo.ValueOf(π6.MergePackageFiles),
			"NewCommentMap":              stargo.ValueOf(π6.NewCommentMap),
			"NewIdent":                   stargo.ValueOf(π6.NewIdent),
			"NewObj":                     stargo.ValueOf(π6.NewObj),
			"NewPackage":                 stargo.ValueOf(π6.NewPackage),
			"NewScope":                   stargo.ValueOf(π6.NewScope),
			"Node":                       stargo.TypeOf(new(π6.Node)),
			"NotNilFilter":               stargo.ValueOf(π6.NotNilFilter),
			"ObjKind":                    stargo.TypeOf(new(π6.ObjKind)),
			"Object":                     stargo.TypeOf(new(π6.Object)),
			"Package":                    stargo.TypeOf(new(π6.Package)),
			"PackageExports":             stargo.ValueOf(π6.PackageExports),
			"ParenExpr":                  stargo.TypeOf(new(π6.ParenExpr)),
			"Pkg":                        stargo.ValueOf(π6.Pkg),
			"Print":                      stargo.ValueOf(π6.Print),
			"RECV":                       stargo.ValueOf(π6.RECV),
			"RangeStmt":                  stargo.TypeOf(new(π6.RangeStmt)),
			"ReturnStmt":                 stargo.TypeOf(new(π6.ReturnStmt)),
			"SEND":                       stargo.ValueOf(π6.SEND),
			"Scope":                      stargo.TypeOf(new(π6.Scope)),
			"SelectStmt":                 stargo.TypeOf(new(π6.SelectStmt)),
			"SelectorExpr":               stargo.TypeOf(new(π6.SelectorExpr)),
			"SendStmt":                   stargo.TypeOf(new(π6.SendStmt)),
			"SliceExpr":                  stargo.TypeOf(new(π6.SliceExpr)),
			"SortImports":                stargo.ValueOf(π6.SortImports),
			"Spec":                       stargo.TypeOf(new(π6.Spec)),
			"StarExpr":                   stargo.TypeOf(new(π6.StarExpr)),
			"Stmt":                       stargo.TypeOf(new(π6.Stmt)),
			"StructType":                 stargo.TypeOf(new(π6.StructType)),
			"SwitchStmt":                 stargo.TypeOf(new(π6.SwitchStmt)),
			"Typ":                        stargo.ValueOf(π6.Typ),
			"TypeAssertExpr":             stargo.TypeOf(new(π6.TypeAssertExpr)),
			"TypeSpec":                   stargo.TypeOf(new(π6.TypeSpec)),
			"TypeSwitchStmt":             stargo.TypeOf(new(π6.TypeSwitchStmt)),
			"UnaryExpr":                  stargo.TypeOf(new(π6.UnaryExpr)),
			"ValueSpec":                  stargo.TypeOf(new(π6.ValueSpec)),
			"Var":                        stargo.ValueOf(π6.Var),
			"Visitor":                    stargo.TypeOf(new(π6.Visitor)),
			"Walk":                       stargo.ValueOf(π6.Walk),
		},
	},
	"go/parser": &starlark.Module{
		Name: "go/parser",
		Members: starlark.StringDict{
			"AllErrors":         stargo.ValueOf(π7.AllErrors),
			"DeclarationErrors": stargo.ValueOf(π7.DeclarationErrors),
			"ImportsOnly":       stargo.ValueOf(π7.ImportsOnly),
			"Mode":              stargo.TypeOf(new(π7.Mode)),
			"PackageClauseOnly": stargo.ValueOf(π7.PackageClauseOnly),
			"ParseComments":     stargo.ValueOf(π7.ParseComments),
			"ParseDir":          stargo.ValueOf(π7.ParseDir),
			"ParseExpr":         stargo.ValueOf(π7.ParseExpr),
			"ParseExprFrom":     stargo.ValueOf(π7.ParseExprFrom),
			"ParseFile":         stargo.ValueOf(π7.ParseFile),
			"SpuriousErrors":    stargo.ValueOf(π7.SpuriousErrors),
			"Trace":             stargo.ValueOf(π7.Trace),
		},
	},
	"go/types": &starlark.Module{
		Name: "go/types",
		Members: starlark.StringDict{
			"Array":                   stargo.TypeOf(new(π8.Array)),
			"AssertableTo":            stargo.ValueOf(π8.AssertableTo),
			"AssignableTo":            stargo.ValueOf(π8.AssignableTo),
			"Basic":                   stargo.TypeOf(new(π8.Basic)),
			"BasicInfo":               stargo.TypeOf(new(π8.BasicInfo)),
			"BasicKind":               stargo.TypeOf(new(π8.BasicKind)),
			"Bool":                    stargo.ValueOf(π8.Bool),
			"Builtin":                 stargo.TypeOf(new(π8.Builtin)),
			"Byte":                    stargo.ValueOf(π8.Byte),
			"Chan":                    stargo.TypeOf(new(π8.Chan)),
			"ChanDir":                 stargo.TypeOf(new(π8.ChanDir)),
			"Checker":                 stargo.TypeOf(new(π8.Checker)),
			"Comparable":              stargo.ValueOf(π8.Comparable),
			"Complex128":              stargo.ValueOf(π8.Complex128),
			"Complex64":               stargo.ValueOf(π8.Complex64),
			"Config":                  stargo.TypeOf(new(π8.Config)),
			"Const":                   stargo.TypeOf(new(π8.Const)),
			"ConvertibleTo":           stargo.ValueOf(π8.ConvertibleTo),
			"DefPredeclaredTestFuncs": stargo.ValueOf(π8.DefPredeclaredTestFuncs),
			"Default":                 stargo.ValueOf(π8.Default),
			"Error":                   stargo.TypeOf(new(π8.Error)),
			"Eval":                    stargo.ValueOf(π8.Eval),
			"ExprString":              stargo.ValueOf(π8.ExprString),
			"FieldVal":                stargo.ValueOf(π8.FieldVal),
			"Float32":                 stargo.ValueOf(π8.Float32),
			"Float64":                 stargo.ValueOf(π8.Float64),
			"Func":                    stargo.TypeOf(new(π8.Func)),
			"Id":                      stargo.ValueOf(π8.Id),
			"Identical":               stargo.ValueOf(π8.Identical),
			"IdenticalIgnoreTags":     stargo.ValueOf(π8.IdenticalIgnoreTags),
			"Implements":              stargo.ValueOf(π8.Implements),
			"ImportMode":              stargo.TypeOf(new(π8.ImportMode)),
			"Importer":                stargo.TypeOf(new(π8.Importer)),
			"ImporterFrom":            stargo.TypeOf(new(π8.ImporterFrom)),
			"Info":                    stargo.TypeOf(new(π8.Info)),
			"Initializer":             stargo.TypeOf(new(π8.Initializer)),
			"Int":                     stargo.ValueOf(π8.Int),
			"Int16":                   stargo.ValueOf(π8.Int16),
			"Int32":                   stargo.ValueOf(π8.Int32),
			"Int64":                   stargo.ValueOf(π8.Int64),
			"Int8":                    stargo.ValueOf(π8.Int8),
			"Interface":               stargo.TypeOf(new(π8.Interface)),
			"Invalid":                 stargo.ValueOf(π8.Invalid),
			"IsBoolean":               stargo.ValueOf(π8.IsBoolean),
			"IsComplex":               stargo.ValueOf(π8.IsComplex),
			"IsConstType":             stargo.ValueOf(π8.IsConstType),
			"IsFloat":                 stargo.ValueOf(π8.IsFloat),
			"IsInteger":               stargo.ValueOf(π8.IsInteger),
			"IsInterface":             stargo.ValueOf(π8.IsInterface),
			"IsNumeric":               stargo.ValueOf(π8.IsNumeric),
			"IsOrdered":               stargo.ValueOf(π8.IsOrdered),
			"IsString":                stargo.ValueOf(π8.IsString),
			"IsUnsigned":              stargo.ValueOf(π8.IsUnsigned),
			"IsUntyped":               stargo.ValueOf(π8.IsUntyped),
			"Label":                   stargo.TypeOf(new(π8.Label)),
			"LookupFieldOrMethod":     stargo.ValueOf(π8.LookupFieldOrMethod),
			"Map":              stargo.TypeOf(new(π8.Map)),
			"MethodExpr":       stargo.ValueOf(π8.MethodExpr),
			"MethodSet":        stargo.TypeOf(new(π8.MethodSet)),
			"MethodVal":        stargo.ValueOf(π8.MethodVal),
			"MissingMethod":    stargo.ValueOf(π8.MissingMethod),
			"Named":            stargo.TypeOf(new(π8.Named)),
			"NewArray":         stargo.ValueOf(π8.NewArray),
			"NewChan":          stargo.ValueOf(π8.NewChan),
			"NewChecker":       stargo.ValueOf(π8.NewChecker),
			"NewConst":         stargo.ValueOf(π8.NewConst),
			"NewField":         stargo.ValueOf(π8.NewField),
			"NewFunc":          stargo.ValueOf(π8.NewFunc),
			"NewInterface":     stargo.ValueOf(π8.NewInterface),
			"NewInterfaceType": stargo.ValueOf(π8.NewInterfaceType),
			"NewLabel":         stargo.ValueOf(π8.NewLabel),
			"NewMap":           stargo.ValueOf(π8.NewMap),
			"NewMethodSet":     stargo.ValueOf(π8.NewMethodSet),
			"NewNamed":         stargo.ValueOf(π8.NewNamed),
			"NewPackage":       stargo.ValueOf(π8.NewPackage),
			"NewParam":         stargo.ValueOf(π8.NewParam),
			"NewPkgName":       stargo.ValueOf(π8.NewPkgName),
			"NewPointer":       stargo.ValueOf(π8.NewPointer),
			"NewScope":         stargo.ValueOf(π8.NewScope),
			"NewSignature":     stargo.ValueOf(π8.NewSignature),
			"NewSlice":         stargo.ValueOf(π8.NewSlice),
			"NewStruct":        stargo.ValueOf(π8.NewStruct),
			"NewTuple":         stargo.ValueOf(π8.NewTuple),
			"NewTypeName":      stargo.ValueOf(π8.NewTypeName),
			"NewVar":           stargo.ValueOf(π8.NewVar),
			"Nil":              stargo.TypeOf(new(π8.Nil)),
			"Object":           stargo.TypeOf(new(π8.Object)),
			"ObjectString":     stargo.ValueOf(π8.ObjectString),
			"Package":          stargo.TypeOf(new(π8.Package)),
			"PkgName":          stargo.TypeOf(new(π8.PkgName)),
			"Pointer":          stargo.TypeOf(new(π8.Pointer)),
			"Qualifier":        stargo.TypeOf(new(π8.Qualifier)),
			"RecvOnly":         stargo.ValueOf(π8.RecvOnly),
			"RelativeTo":       stargo.ValueOf(π8.RelativeTo),
			"Rune":             stargo.ValueOf(π8.Rune),
			"Scope":            stargo.TypeOf(new(π8.Scope)),
			"Selection":        stargo.TypeOf(new(π8.Selection)),
			"SelectionKind":    stargo.TypeOf(new(π8.SelectionKind)),
			"SelectionString":  stargo.ValueOf(π8.SelectionString),
			"SendOnly":         stargo.ValueOf(π8.SendOnly),
			"SendRecv":         stargo.ValueOf(π8.SendRecv),
			"Signature":        stargo.TypeOf(new(π8.Signature)),
			"Sizes":            stargo.TypeOf(new(π8.Sizes)),
			"SizesFor":         stargo.ValueOf(π8.SizesFor),
			"Slice":            stargo.TypeOf(new(π8.Slice)),
			"StdSizes":         stargo.TypeOf(new(π8.StdSizes)),
			"String":           stargo.ValueOf(π8.String),
			"Struct":           stargo.TypeOf(new(π8.Struct)),
			"Tuple":            stargo.TypeOf(new(π8.Tuple)),
			"Typ":              stargo.VarOf(&π8.Typ),
			"Type":             stargo.TypeOf(new(π8.Type)),
			"TypeAndValue":     stargo.TypeOf(new(π8.TypeAndValue)),
			"TypeName":         stargo.TypeOf(new(π8.TypeName)),
			"TypeString":       stargo.ValueOf(π8.TypeString),
			"Uint":             stargo.ValueOf(π8.Uint),
			"Uint16":           stargo.ValueOf(π8.Uint16),
			"Uint32":           stargo.ValueOf(π8.Uint32),
			"Uint64":           stargo.ValueOf(π8.Uint64),
			"Uint8":            stargo.ValueOf(π8.Uint8),
			"Uintptr":          stargo.ValueOf(π8.Uintptr),
			"Universe":         stargo.VarOf(&π8.Universe),
			"Unsafe":           stargo.VarOf(&π8.Unsafe),
			"UnsafePointer":    stargo.ValueOf(π8.UnsafePointer),
			"UntypedBool":      stargo.ValueOf(π8.UntypedBool),
			"UntypedComplex":   stargo.ValueOf(π8.UntypedComplex),
			"UntypedFloat":     stargo.ValueOf(π8.UntypedFloat),
			"UntypedInt":       stargo.ValueOf(π8.UntypedInt),
			"UntypedNil":       stargo.ValueOf(π8.UntypedNil),
			"UntypedRune":      stargo.ValueOf(π8.UntypedRune),
			"UntypedString":    stargo.ValueOf(π8.UntypedString),
			"Var":              stargo.TypeOf(new(π8.Var)),
			"WriteExpr":        stargo.ValueOf(π8.WriteExpr),
			"WriteSignature":   stargo.ValueOf(π8.WriteSignature),
			"WriteType":        stargo.ValueOf(π8.WriteType),
		},
	},
}
