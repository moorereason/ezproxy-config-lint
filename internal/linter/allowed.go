// Copyright Carleton University Library All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.
package linter

// optionPairs maps each opening Option directive to the directive that closes it.
var optionPairs = map[Directive]Directive{ //nolint:gochecknoglobals
	OptionCookiePassThrough:    OptionCookie,
	OptionDomainCookieOnly:     OptionCookie,
	OptionHideEZproxy:          OptionNoHideEZproxy,
	OptionMetaEZproxyRewriting: OptionNoMetaEZproxyRewriting,
	OptionNoCookie:             OptionCookie,
	OptionNoHttpsHyphens:       OptionHttpsHyphens,
	OptionProxyFTP:             OptionNoProxyFTP,
	OptionUTF16:                OptionNoUTF16,
	OptionXForwardedFor:        OptionNoXForwardedFor,
}

// titlePrevious is the set of directives that may appear immediately before Title.
// It is built once. Openers are added from optionPairs in init.
var titlePrevious = map[Directive]struct{}{ //nolint:gochecknoglobals
	AddUserHeader:               {},
	AnonymousURL:                {},
	Cookie:                      {},
	DbVar:                       {},
	DbVar0:                      {},
	DbVar1:                      {},
	DbVar2:                      {},
	DbVar3:                      {},
	DbVar4:                      {},
	DbVar5:                      {},
	DbVar6:                      {},
	DbVar7:                      {},
	DbVar8:                      {},
	DbVar9:                      {},
	Group:                       {},
	HTTPMethod:                  {},
	OptionCookie:                {},
	OptionEbraryUnencodedTokens: {},
	ProxyHostnameEdit:           {},
	Referer:                     {},
	Undefined:                   {},
}

// descriptionPrevious is the set of directives that may appear immediately before Description.
var descriptionPrevious = map[Directive]struct{}{ //nolint:gochecknoglobals
	Description: {},
	Title:       {},
}

// proxyHostnameEditPrevious is the set of directives that may appear immediately before ProxyHostnameEdit.
// Openers are added from optionPairs in init.
var proxyHostnameEditPrevious = map[Directive]struct{}{ //nolint:gochecknoglobals
	AddUserHeader:     {},
	AnonymousURL:      {},
	Cookie:            {},
	DbVar:             {},
	DbVar0:            {},
	DbVar1:            {},
	DbVar2:            {},
	DbVar3:            {},
	DbVar4:            {},
	DbVar5:            {},
	DbVar6:            {},
	DbVar7:            {},
	DbVar8:            {},
	DbVar9:            {},
	Group:             {},
	HTTPMethod:        {},
	OptionCookie:      {},
	ProxyHostnameEdit: {},
	Undefined:         {},
}

// urlPrevious is the set of directives that may appear immediately before URL.
var urlPrevious = map[Directive]struct{}{ //nolint:gochecknoglobals
	AllowVars:   {},
	ByteServe:   {},
	Description: {},
	EBLSecret:   {},
	EbrarySite:  {},
	EncryptVar:  {},
	HTTPHeader:  {},
	MimeFilter:  {},
	Title:       {},
}

// optionOpenerPrevious is the set of directives that may appear immediately before an opening Option.
// Openers are added from optionPairs in init.
var optionOpenerPrevious = map[Directive]struct{}{ //nolint:gochecknoglobals
	AddUserHeader:     {},
	AnonymousURL:      {},
	DbVar:             {},
	DbVar0:            {},
	DbVar1:            {},
	DbVar2:            {},
	DbVar3:            {},
	DbVar4:            {},
	DbVar5:            {},
	DbVar6:            {},
	DbVar7:            {},
	DbVar8:            {},
	DbVar9:            {},
	Group:             {},
	HTTPMethod:        {},
	OptionCookie:      {},
	ProxyHostnameEdit: {},
	Undefined:         {},
}

// addUserHeaderPrevious is the set of directives that may appear immediately before AddUserHeader with qualifiers.
// Openers are added from optionPairs in init.
var addUserHeaderPrevious = map[Directive]struct{}{ //nolint:gochecknoglobals
	AddUserHeader:     {},
	AnonymousURL:      {},
	Cookie:            {},
	DbVar:             {},
	DbVar0:            {},
	DbVar1:            {},
	DbVar2:            {},
	DbVar3:            {},
	DbVar4:            {},
	DbVar5:            {},
	DbVar6:            {},
	DbVar7:            {},
	DbVar8:            {},
	DbVar9:            {},
	Group:             {},
	HTTPMethod:        {},
	OptionCookie:      {},
	ProxyHostnameEdit: {},
	Undefined:         {},
}

// addUserHeaderBarePrevious is the set of directives that may appear immediately before AddUserHeader with no qualifiers.
// Closers are added from optionPairs in init.
var addUserHeaderBarePrevious = map[Directive]struct{}{ //nolint:gochecknoglobals
	AnonymousURL:     {},
	Domain:           {},
	DomainJavaScript: {},
	Host:             {},
	HostJavaScript:   {},
	NeverProxy:       {},
	Replace:          {},
	URL:              {},
}

// anonymousURLPrevious is the set of directives that may appear immediately before AnonymousURL.
// Openers are added from optionPairs in init.
var anonymousURLPrevious = map[Directive]struct{}{ //nolint:gochecknoglobals
	AddUserHeader:     {},
	AnonymousURL:      {},
	Cookie:            {},
	DbVar:             {},
	DbVar0:            {},
	DbVar1:            {},
	DbVar2:            {},
	DbVar3:            {},
	DbVar4:            {},
	DbVar5:            {},
	DbVar6:            {},
	DbVar7:            {},
	DbVar8:            {},
	DbVar9:            {},
	Group:             {},
	HTTPMethod:        {},
	ProxyHostnameEdit: {},
	Undefined:         {},
}

// anonymousURLStarPrevious is the set of directives that may appear immediately before "AnonymousURL -*".
// Closers are added from optionPairs in init.
var anonymousURLStarPrevious = map[Directive]struct{}{ //nolint:gochecknoglobals
	AddUserHeader:    {},
	Domain:           {},
	DomainJavaScript: {},
	Host:             {},
	HostJavaScript:   {},
	NeverProxy:       {},
	Replace:          {},
	URL:              {},
}

// optionCloserPrevious is the set of directives that may appear immediately before a closing Option.
// Closers are added from optionPairs in init.
var optionCloserPrevious = map[Directive]struct{}{ //nolint:gochecknoglobals
	AddUserHeader:    {},
	AnonymousURL:     {},
	DbVar:            {},
	Domain:           {},
	DomainJavaScript: {},
	Host:             {},
	HostJavaScript:   {},
	NeverProxy:       {},
	Replace:          {},
	URL:              {},
}

// openerDirectives is the set of Option directives that open a pair.
// It is filled from optionPairs in init.
var openerDirectives = map[Directive]struct{}{} //nolint:gochecknoglobals

// closerDirectives is the set of Option directives that close a pair.
// It is filled from optionPairs in init.
var closerDirectives = map[Directive]struct{}{} //nolint:gochecknoglobals

func init() {
	for _, opener := range OpenerOptions() {
		addUserHeaderPrevious[opener] = struct{}{}
		anonymousURLPrevious[opener] = struct{}{}
		openerDirectives[opener] = struct{}{}
		optionOpenerPrevious[opener] = struct{}{}
		proxyHostnameEditPrevious[opener] = struct{}{}
		titlePrevious[opener] = struct{}{}
	}
	for _, closer := range CloserOptions() {
		addUserHeaderBarePrevious[closer] = struct{}{}
		anonymousURLStarPrevious[closer] = struct{}{}
		closerDirectives[closer] = struct{}{}
		optionCloserPrevious[closer] = struct{}{}
	}
}
