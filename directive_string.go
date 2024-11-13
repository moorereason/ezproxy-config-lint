// Code generated by "stringer -type Directive --linecomment"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Undefined-0]
	_ = x[AddUserHeader-1]
	_ = x[AllowIP-2]
	_ = x[AllowVars-3]
	_ = x[AnonymousURL-4]
	_ = x[Audit-5]
	_ = x[AuditPurge-6]
	_ = x[AutoLoginIP-7]
	_ = x[AutoLoginIPBanner-8]
	_ = x[BinaryTimeout-9]
	_ = x[Books24x7Site-10]
	_ = x[ByteServe-11]
	_ = x[CASServiceURL-12]
	_ = x[Charset-13]
	_ = x[ClientTimeout-14]
	_ = x[ConnectWindow-15]
	_ = x[Cookie-16]
	_ = x[CookieFilter-17]
	_ = x[DbVar-18]
	_ = x[DenyIfRequestHeader-19]
	_ = x[Description-20]
	_ = x[DNS-21]
	_ = x[Domain-22]
	_ = x[DomainJavaScript-23]
	_ = x[EBLSecret-24]
	_ = x[ebrarySite-25]
	_ = x[EncryptVar-26]
	_ = x[ExcludeIP-27]
	_ = x[ExcludeIPBanner-28]
	_ = x[ExtraLoginCookie-29]
	_ = x[Find-30]
	_ = x[FirstPort-31]
	_ = x[FormSelect-32]
	_ = x[FormSubmit-33]
	_ = x[FormVariable-34]
	_ = x[Gartner-35]
	_ = x[Group-36]
	_ = x[HAName-37]
	_ = x[HAPeer-38]
	_ = x[Host-39]
	_ = x[HostJavaScript-40]
	_ = x[HTTPHeader-41]
	_ = x[HTTPMethod-42]
	_ = x[Identifier-43]
	_ = x[IncludeFile-44]
	_ = x[IncludeIP-45]
	_ = x[Interface-46]
	_ = x[IntruderIPAttempts-47]
	_ = x[IntruderLog-48]
	_ = x[IntruderUserAttempts-49]
	_ = x[IntrusionAPI-50]
	_ = x[LBPeer-51]
	_ = x[Location-52]
	_ = x[LogFile-53]
	_ = x[LogFilter-54]
	_ = x[LogFormat-55]
	_ = x[LoginCookieDomain-56]
	_ = x[LoginCookieName-57]
	_ = x[LoginMenu-58]
	_ = x[LoginPort-59]
	_ = x[LoginPortSSL-60]
	_ = x[LogSPU-61]
	_ = x[MaxConcurrentTransfers-62]
	_ = x[MaxLifetime-63]
	_ = x[MaxSessions-64]
	_ = x[MaxVirtualHosts-65]
	_ = x[MessagesFile-66]
	_ = x[MetaFind-67]
	_ = x[MimeFilter-68]
	_ = x[Name-69]
	_ = x[NeverProxy-70]
	_ = x[OptionAcceptXForwardedFor-71]
	_ = x[OptionAllowWebSubdirectories-72]
	_ = x[OptionAnyDNSHostname-73]
	_ = x[OptionBlockCountryChange-74]
	_ = x[OptionCookie-75]
	_ = x[OptionCookiePassThrough-76]
	_ = x[OptionCSRFToken-77]
	_ = x[OptionDisableSSL40bit-78]
	_ = x[OptionDisableSSL56bit-79]
	_ = x[OptionDisableSSLv2-80]
	_ = x[OptionDomainCookieOnly-81]
	_ = x[OptionExcludeIPMenu-82]
	_ = x[OptionForceHTTPSAdmin-83]
	_ = x[OptionForceHTTPSLogin-84]
	_ = x[OptionForceWildcardCertificate-85]
	_ = x[OptionHideEZproxy-86]
	_ = x[OptionHttpsHyphens-87]
	_ = x[OptionIChooseToUseDomainLinesThatThreatenTheSecurityOfMyNetwork-88]
	_ = x[OptionIgnoreWildcardCertificate-89]
	_ = x[OptionIPv6-90]
	_ = x[OptionLoginReplaceGroups-91]
	_ = x[OptionLogReferer-92]
	_ = x[OptionLogSAML-93]
	_ = x[OptionLogSession-94]
	_ = x[OptionLogSPUEdit-95]
	_ = x[OptionLogUser-96]
	_ = x[OptionMenuByGroups-97]
	_ = x[OptionMetaEZproxyRewriting-98]
	_ = x[OptionNoCookie-99]
	_ = x[OptionNoHideEZproxy-100]
	_ = x[OptionNoHttpsHyphens-101]
	_ = x[OptionNoMetaEZproxyRewriting-102]
	_ = x[OptionNoProxyFTP-103]
	_ = x[OptionNoUTF16-104]
	_ = x[OptionNoXForwardedFor-105]
	_ = x[OptionProxyByHostname-106]
	_ = x[OptionProxyFTP-107]
	_ = x[OptionRecordPeaks-108]
	_ = x[OptionRedirectUnknown-109]
	_ = x[OptionReferInHostname-110]
	_ = x[OptionRelaxedRADIUS-111]
	_ = x[OptionRequireAuthenticate-112]
	_ = x[OptionSafariCookiePatch-113]
	_ = x[OptionStatusUser-114]
	_ = x[OptionTicketIgnoreExcludeIP-115]
	_ = x[OptionUnsafeRedirectUnknown-116]
	_ = x[OptionUsernameCaretN-117]
	_ = x[OptionUTF16-118]
	_ = x[OptionXForwardedFor-119]
	_ = x[OverDriveSite-120]
	_ = x[PDFRefresh-121]
	_ = x[PDFRefreshPost-122]
	_ = x[PDFRefreshPre-123]
	_ = x[PidFile-124]
	_ = x[Proxy-125]
	_ = x[ProxyHostnameEdit-126]
	_ = x[ProxySSL-127]
	_ = x[RADIUSRetry-128]
	_ = x[RedirectSafe-129]
	_ = x[Referer-130]
	_ = x[RejectIP-131]
	_ = x[RemoteIPHeader-132]
	_ = x[RemoteIPInternalProxy-133]
	_ = x[RemoteIPTrustedProxy-134]
	_ = x[RemoteTimeout-135]
	_ = x[Replace-136]
	_ = x[RunAs-137]
	_ = x[SPUEdit-138]
	_ = x[SPUEditVar-139]
	_ = x[ShibbolethDisable-140]
	_ = x[ShibbolethMetadata-141]
	_ = x[SkipPort-142]
	_ = x[SQLiteTempDir-143]
	_ = x[SSLCipherSuite-144]
	_ = x[SSLHonorCipherOrder-145]
	_ = x[SSLOpenSSLConfCmd-146]
	_ = x[SSOUsername-147]
	_ = x[Title-148]
	_ = x[TokenKey-149]
	_ = x[TokenSignatureKey-150]
	_ = x[UMask-151]
	_ = x[URL-152]
	_ = x[URLAppendEncoded-153]
	_ = x[URLRedirect-154]
	_ = x[URLRedirectAppend-155]
	_ = x[URLRedirectAppendEncoded-156]
	_ = x[UsageLimit-157]
	_ = x[Validate-158]
	_ = x[XDebug-159]
}

const _Directive_name = "UndefinedAddUserHeaderAllowIPAllowVarsAnonymousURLAuditAuditPurgeAutoLoginIPAutoLoginIPBannerBinaryTimeoutBooks24x7SiteByteServeCASServiceURLCharsetClientTimeoutConnectWindowCookieCookieFilterDbVarDenyIfRequestHeaderDescriptionDNSDomainDomainJavaScriptEBLSecretebrarySiteEncryptVarExcludeIPExcludeIPBannerExtraLoginCookieFindFirstPortFormSelectFormSubmitFormVariableGartnerGroupHANameHAPeerHostHostJavaScriptHTTPHeaderHTTPMethodIdentifierIncludeFileIncludeIPInterfaceIntruderIPAttemptsIntruderLogIntruderUserAttemptsIntrusionAPILBPeerLocationLogFileLogFilterLogFormatLoginCookieDomainLoginCookieNameLoginMenuLoginPortLoginPortSSLLogSPUMaxConcurrentTransfersMaxLifetimeMaxSessionsMaxVirtualHostsMessagesFileMetaFindMimeFilterNameNeverProxyOption AcceptX-Forwarded-For.Option AllowWebSubdirectories.Option AnyDNSHostname.Option BlockCountryChange.Option Cookie.Option CookiePassThrough.Option CSRFToken.Option DisableSSL40bit.Option DisableSSL56bit.Option DisableSSLv2.Option DomainCookieOnly.Option ExcludeIPMenu.Option ForceHTTPSAdmin.Option ForceHTTPSLogin.Option ForceWildcardCertificate.Option HideEZproxy.Option HttpsHyphens.Option I choose to use Domain lines that threaten the security of my network.Option IgnoreWildcardCertificate.Option IPv6.Option LoginReplaceGroups.Option LogReferer.Option LogSAML.Option LogSession.Option LogSPUEdit.Option LogUser.Option MenuByGroups.Option MetaEZproxyRewriting.Option NoCookie.Option NoHideEZproxy.Option NoHttpsHyphens.Option NoMetaEZproxyRewriting.Option NoProxyFTP.Option NoUTF16.Option NoX-Forwarded-For.Option ProxyByHostname.Option ProxyFTP.Option RecordPeaks.Option RedirectUnknown.Option ReferInHostname.Option RelaxedRADIUS.Option RequireAuthenticate.Option SafariCookiePatch.Option StatusUser.Option TicketIgnoreExcludeIP.Option UnsafeRedirectUnknown.Option UsernameCaretN.Option UTF16.Option X-Forwarded-For.OverDriveSitePDFRefreshPDFRefreshPostPDFRefreshPrePidFileProxyProxyHostnameEditProxySSLRADIUSRetryRedirectSafeRefererRejectIPRemoteIPHeaderRemoteIPInternalProxyRemoteIPTrustedProxyRemoteTimeoutReplaceRunAsSPUEditSPUEditVarShibbolethDisableShibbolethMetadataSkipPortSQLiteTempDirSSLCipherSuiteSSLHonorCipherOrderSSLOpenSSLConfCmdSSOUsernameTitleTokenKeyTokenSignatureKeyUMaskURLURLAppendEncodedURLRedirectURLRedirectAppendURLRedirectAppendEncodedUsageLimitValidateXDebug"

var _Directive_index = [...]uint16{0, 9, 22, 29, 38, 50, 55, 65, 76, 93, 106, 119, 128, 141, 148, 161, 174, 180, 192, 197, 216, 227, 230, 236, 252, 261, 271, 281, 290, 305, 321, 325, 334, 344, 354, 366, 373, 378, 384, 390, 394, 408, 418, 428, 438, 449, 458, 467, 485, 496, 516, 528, 534, 542, 549, 558, 567, 584, 599, 608, 617, 629, 635, 657, 668, 679, 694, 706, 714, 724, 728, 738, 767, 797, 819, 845, 859, 884, 901, 924, 947, 967, 991, 1012, 1035, 1058, 1090, 1109, 1129, 1206, 1239, 1251, 1277, 1295, 1310, 1328, 1346, 1361, 1381, 1409, 1425, 1446, 1468, 1498, 1516, 1531, 1556, 1579, 1595, 1614, 1637, 1660, 1681, 1708, 1733, 1751, 1780, 1809, 1831, 1844, 1867, 1880, 1890, 1904, 1917, 1924, 1929, 1946, 1954, 1965, 1977, 1984, 1992, 2006, 2027, 2047, 2060, 2067, 2072, 2079, 2089, 2106, 2124, 2132, 2145, 2159, 2178, 2195, 2206, 2211, 2219, 2236, 2241, 2244, 2260, 2271, 2288, 2312, 2322, 2330, 2336}

func (i Directive) String() string {
	if i < 0 || i >= Directive(len(_Directive_index)-1) {
		return "Directive(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Directive_name[_Directive_index[i]:_Directive_index[i+1]]
}
