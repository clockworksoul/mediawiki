package mediawiki

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Return general information about the site.
// https://www.mediawiki.org/wiki/Special:MyLanguage/API:Siteinfo
//
// Flags:
// * This module requires read rights.

const (
	SiteinfoPropGeneral            = "general"
	SiteinfoPropNamespaces         = "namespaces"
	SiteinfoPropNamespacealiases   = "namespacealiases"
	SiteinfoPropSpecialpagealiases = "specialpagealiases"
	SiteinfoPropMagicwords         = "magicwords"
	SiteinfoPropInterwikimap       = "interwikimap"
	SiteinfoPropDbrepllag          = "dbrepllag"
	SiteinfoPropStatistics         = "statistics"
	SiteinfoPropUsergroups         = "usergroups"
	SiteinfoPropLibraries          = "libraries"
	SiteinfoPropExtensions         = "extensions"
	SiteinfoPropFileextensions     = "fileextensions"
	SiteinfoPropRightsinfo         = "rightsinfo"
	SiteinfoPropRestrictions       = "restrictions"
	SiteinfoPropLanguages          = "languages"
	SiteinfoPropLanguagevariants   = "languagevariants"
	SiteinfoPropSkins              = "skins"
	SiteinfoPropExtensiontags      = "extensiontags"
	SiteinfoPropFunctionhooks      = "functionhooks"
	SiteinfoPropShowhooks          = "showhooks"
	SiteinfoPropVariables          = "variables"
	SiteinfoPropProtocols          = "protocols"
	SiteinfoPropDefaultoptions     = "defaultoptions"
	SiteinfoPropUploaddialog       = "uploaddialog"
)

type SiteinfoResponse struct {
	QueryResponse
	Query *SiteinfoResponseQuery `json:"query,omitempty"`
}

type SiteinfoResponseQuery struct {
	General           *SiteinfoGeneral                        `json:"general,omitempty"`
	Namespaces        map[string]SiteinfoNamespace            `json:"namespaces,omitempty"`
	NamespacesAliases []SiteinfoNamespace                     `json:"namespacealiases,omitempty"`
	SpecialPageAlises []SiteinfoSpecialPageAlises             `json:"specialpagealiases,omitempty"`
	MagicWords        []SiteinfoMagicWords                    `json:"magicwords,omitempty"`
	InterwikiMap      []SiteinfoInterwikiMap                  `json:"interwikimap,omitempty"`
	Dbrepllag         []SiteinfoDbrepllag                     `json:"dbrepllag,omitempty"`
	Statistics        map[string]int                          `json:"statistics,omitempty"`
	Usergroups        []SiteinfoUsergroups                    `json:"usergroups,omitempty"`
	Libraries         []SiteinfoLibrary                       `json:"libraries,omitempty"`
	Extensions        []SiteinfoExtension                     `json:"extensions,omitempty"`
	Fileextensions    []SiteinfoFileextension                 `json:"fileextensions,omitempty"`
	Rightsinfo        *SiteinfoRightsinfo                     `json:"rightsinfo,omitempty"`
	Restrictions      *SiteinfoRestrictions                   `json:"restrictions,omitempty"`
	Languages         []SiteinfoLanguages                     `json:"languages,omitempty"`
	LanguageVariants  map[string]map[string]SiteinfoFallbacks `json:"languagevariants,omitempty"`
	Skins             []SiteinfoSkin                          `json:"skins,omitempty"`
	ExtensionTags     []string                                `json:"extensiontags,omitempty"`
	FunctionHooks     []string                                `json:"functionhooks,omitempty"`
	ShowHooks         []SiteinfoShowhooks                     `json:"showhooks,omitempty"`
	Variables         []string                                `json:"variables,omitempty"`
	Protocols         []string                                `json:"protocols,omitempty"`
	DefaultOptions    map[string]any                          `json:"defaultoptions,omitempty"`
	UploadDialog      *SiteinfoUploadDialog                   `json:"uploaddialog,omitempty"`
}

type SiteinfoGeneral struct {
	Mainpage                        string                                          `json:"mainpage,omitempty"`
	Base                            string                                          `json:"base,omitempty"`
	Sitename                        string                                          `json:"sitename,omitempty"`
	Logo                            string                                          `json:"logo,omitempty"`
	Generator                       string                                          `json:"generator,omitempty"`
	Phpversion                      string                                          `json:"phpversion,omitempty"`
	Phpsapi                         string                                          `json:"phpsapi,omitempty"`
	Dbtype                          string                                          `json:"dbtype,omitempty"`
	Dbversion                       string                                          `json:"dbversion,omitempty"`
	Langconversion                  string                                          `json:"langconversion"`
	Linkconversion                  string                                          `json:"linkconversion"`
	Titleconversion                 string                                          `json:"titleconversion"`
	Linkprefixcharset               string                                          `json:"linkprefixcharset"`
	Linkprefix                      string                                          `json:"linkprefix"`
	Linktrail                       string                                          `json:"linktrail,omitempty"`
	Legaltitlechars                 string                                          `json:"legaltitlechars,omitempty"`
	Invalidusernamechars            string                                          `json:"invalidusernamechars,omitempty"`
	Fixarabicunicode                string                                          `json:"fixarabicunicode"`
	Fixmalayalamunicode             string                                          `json:"fixmalayalamunicode"`
	GitHash                         string                                          `json:"git-hash,omitempty"`
	GitBranch                       string                                          `json:"git-branch,omitempty"`
	Case                            string                                          `json:"case,omitempty"`
	Lang                            string                                          `json:"lang,omitempty"`
	Fallback                        []interface{}                                   `json:"fallback"`
	Fallback8BitEncoding            string                                          `json:"fallback8bitEncoding,omitempty"`
	Writeapi                        string                                          `json:"writeapi"`
	Maxarticlesize                  int                                             `json:"maxarticlesize,omitempty"`
	Timezone                        string                                          `json:"timezone,omitempty"`
	Timeoffset                      int                                             `json:"timeoffset"`
	Articlepath                     string                                          `json:"articlepath,omitempty"`
	Scriptpath                      string                                          `json:"scriptpath"`
	Script                          string                                          `json:"script,omitempty"`
	Variantarticlepath              bool                                            `json:"variantarticlepath"`
	Server                          string                                          `json:"server,omitempty"`
	Servername                      string                                          `json:"servername,omitempty"`
	Wikiid                          string                                          `json:"wikiid,omitempty"`
	Time                            *time.Time                                      `json:"time,omitempty"`
	Misermode                       string                                          `json:"misermode,omitempty"`
	Uploadsenabled                  string                                          `json:"uploadsenabled"`
	Maxuploadsize                   int64                                           `json:"maxuploadsize,omitempty"`
	Minuploadchunksize              int                                             `json:"minuploadchunksize,omitempty"`
	Galleryoptions                  *SiteinfoGeneralGalleryoptions                  `json:"galleryoptions,omitempty"`
	Thumblimits                     []int                                           `json:"thumblimits,omitempty"`
	Imagelimits                     []SiteinfoGeneralImagelimits                    `json:"imagelimits,omitempty"`
	Favicon                         string                                          `json:"favicon,omitempty"`
	Centralidlookupprovider         string                                          `json:"centralidlookupprovider,omitempty"`
	Allcentralidlookupproviders     []string                                        `json:"allcentralidlookupproviders,omitempty"`
	Interwikimagic                  string                                          `json:"interwikimagic"`
	Magiclinks                      []string                                        `json:"magiclinks"`
	Categorycollation               string                                          `json:"categorycollation,omitempty"`
	Nofollowlinks                   string                                          `json:"nofollowlinks,omitempty"`
	Nofollownsexceptions            []interface{}                                   `json:"nofollownsexceptions,omitempty"`
	Nofollowdomainexceptions        []string                                        `json:"nofollowdomainexceptions,omitempty"`
	WmfConfig                       *SiteinfoGeneralWmfConfig                       `json:"wmf-config,omitempty"`
	Extensiondistributor            *SiteinfoGeneralExtensiondistributor            `json:"extensiondistributor,omitempty"`
	Mobileserver                    string                                          `json:"mobileserver,omitempty"`
	ReadinglistsConfig              *SiteinfoGeneralReadinglistsConfig              `json:"readinglists-config,omitempty"`
	Citeresponsivereferences        string                                          `json:"citeresponsivereferences,omitempty"`
	Linter                          *SiteinfoGeneralLinter                          `json:"linter,omitempty"`
	PageviewserviceSupportedMetrics *SiteinfoGeneralPageviewserviceSupportedMetrics `json:"pageviewservice-supported-metrics,omitempty"`
}

type SiteinfoNamespace struct {
	ID                  int    `json:"id"`
	Case                string `json:"case,omitempty"`
	Subpages            any    `json:"subpages,omitempty"`
	Canonical           string `json:"canonical,omitempty"`
	Content             any    `json:"content,omitempty"`
	NamespaceProtection string `json:"namespaceprotection,omitempty"`
	Comment             string `json:"*"`
}

type SiteinfoSpecialPageAlises struct {
	RealName string   `json:"realname,omitempty"`
	Alises   []string `json:"aliases,omitempty"`
}

type SiteinfoMagicWords struct {
	Name          string   `json:"name,omitempty"`
	Alises        []string `json:"aliases,omitempty"`
	CaseSensitive any      `json:"case-sensitive,omitempty"`
}

type SiteinfoInterwikiMap struct {
	Prefix string `json:"prefix,omitempty"`
	Local  any    `json:"local,omitempty"`
	URL    string `json:"url,omitempty"`
	API    string `json:"api,omitempty"`
}

type SiteinfoUsergroups struct {
	Name   string   `json:"name,omitempty"`
	Rights []string `json:"rights,omitempty"`
}

type SiteinfoUploadDialog struct {
	Fields          map[string]string `json:"fields,omitempty"`
	LicenseMessages map[string]string `json:"licensemessages,omitempty"`
	Comment         map[string]string `json:"comment,omitempty"`
	Format          map[string]string `json:"format,omitempty"`
}

type SiteinfoShowhooks struct {
	Name        string   `json:"name,omitempty"`
	Subscribers []string `json:"subscribers,omitempty"`
}

type SiteinfoSkin struct {
	Code     string `json:"code,omitempty"`
	Default  any    `json:"default,omitempty"`
	Unusable any    `json:"unusable,omitempty"`
	Comment  string `json:"*,omitempty"`
}

type SiteinfoFallbacks struct {
	Fallbacks []string `json:"fallbacks,omitempty"`
}

type SiteinfoLanguages struct {
	Code    string `json:"code,omitempty"`
	Bcp47   string `json:"bcp47,omitempty"`
	Comment string `json:"*,omitempty"`
}

type SiteinfoRestrictions struct {
	Types               []string `json:"types,omitempty"`
	Levels              []string `json:"levels,omitempty"`
	Cascadinglevels     []string `json:"cascadinglevels,omitempty"`
	Semiprotectedlevels []string `json:"semiprotectedlevels,omitempty"`
}

type SiteinfoRightsinfo struct {
	URL  string `json:"url"`
	Text string `json:"text"`
}

type SiteinfoFileextension struct {
	Ext string `json:"ext"`
}

type SiteinfoExtension struct {
	Type           string     `json:"type,omitempty"`
	Name           string     `json:"name,omitempty"`
	Descriptionmsg string     `json:"descriptionmsg,omitempty"`
	Author         string     `json:"author,omitempty"`
	URL            string     `json:"url,omitempty"`
	VcsSystem      string     `json:"vcs-system,omitempty"`
	VcsVersion     string     `json:"vcs-version,omitempty"`
	VcsURL         string     `json:"vcs-url,omitempty"`
	VcsDate        *time.Time `json:"vcs-date,omitempty"`
	LicenseName    string     `json:"license-name,omitempty"`
	License        string     `json:"license,omitempty"`
	Namemsg        string     `json:"namemsg,omitempty"`
	Version        string     `json:"version,omitempty"`
	Credits        string     `json:"credits,omitempty"`
	Description    string     `json:"description,omitempty"`
}

type SiteinfoLibrary struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type SiteinfoDbrepllag struct {
	Host string  `json:"host"`
	Lag  float64 `json:"lag"`
}

type SiteinfoGeneralGalleryoptions struct {
	ImagesPerRow   int    `json:"imagesPerRow"`
	ImageWidth     int    `json:"imageWidth"`
	ImageHeight    int    `json:"imageHeight"`
	CaptionLength  string `json:"captionLength"`
	ShowBytes      string `json:"showBytes"`
	Mode           string `json:"mode"`
	ShowDimensions string `json:"showDimensions"`
}

type SiteinfoGeneralImagelimits struct {
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}

type SiteinfoGeneralWmfConfig struct {
	WmfMasterDatacenter           string `json:"wmfMasterDatacenter,omitempty"`
	WmfEtcdLastModifiedIndex      int    `json:"wmfEtcdLastModifiedIndex,omitempty"`
	WmgCirrusSearchDefaultCluster string `json:"wmgCirrusSearchDefaultCluster,omitempty"`
	WgCirrusSearchDefaultCluster  string `json:"wgCirrusSearchDefaultCluster,omitempty"`
}

type SiteinfoGeneralExtensiondistributor struct {
	Snapshots []string `json:"snapshots,omitempty"`
	List      string   `json:"list,omitempty"`
}

type SiteinfoGeneralReadinglistsConfig struct {
	MaxListsPerUser      int `json:"maxListsPerUser,omitempty"`
	MaxEntriesPerList    int `json:"maxEntriesPerList,omitempty"`
	DeletedRetentionDays int `json:"deletedRetentionDays,omitempty"`
}

type SiteinfoGeneralLinter struct {
	High   []string `json:"high,omitempty"`
	Medium []string `json:"medium,omitempty"`
	Low    []string `json:"low,omitempty"`
}

type SiteinfoGeneralPageviews struct {
	Pageviews string `json:"pageviews,omitempty"`
}

type SiteinfoGeneralSiteviews struct {
	Pageviews string `json:"pageviews,omitempty"`
	Uniques   string `json:"uniques,omitempty"`
}

type SiteinfoGeneralMostviewed struct {
	Pageviews string `json:"pageviews,omitempty"`
}

type SiteinfoGeneralPageviewserviceSupportedMetrics struct {
	Pageviews  SiteinfoGeneralPageviews  `json:"pageviews,omitempty"`
	Siteviews  SiteinfoGeneralSiteviews  `json:"siteviews,omitempty"`
	Mostviewed SiteinfoGeneralMostviewed `json:"mostviewed,omitempty"`
}

type SiteinfoClient struct {
	o []QueryOption
	c *Client
}

// WithQueryProp
// Which properties to get for the queried pages.
func (c *Client) Siteinfo() *SiteinfoClient {
	return &SiteinfoClient{c: c}
}

// prop
// Which information to get.
// One or more of: dbrepllag, defaultoptions, extensions, extensiontags,
// fileextensions, functionhooks, general, interwikimap, languages,
// languagevariants, libraries, magicwords, namespacealiases, namespaces,
// protocols, restrictions, rightsinfo, showhooks, skins, specialpagealiases,
// statistics, uploaddialog, usergroups, variables
// Default: general
func (w *SiteinfoClient) Prop(s ...string) *SiteinfoClient {
	w.o = append(w.o, func(m map[string]string) {
		m["siprop"] = strings.Join(s, "|")
	})
	return w
}

// filteriw
// Return only local or only nonlocal entries of the interwiki map.
// One of the following values: !local, local
func (w *SiteinfoClient) Filteriw(s string) *SiteinfoClient {
	w.o = append(w.o, func(m map[string]string) {
		m["sifilteriw"] = s
	})
	return w
}

// showalldb
// List all database servers, not just the one lagging the most.
func (w *SiteinfoClient) Showalldb(b bool) *SiteinfoClient {
	w.o = append(w.o, func(m map[string]string) {
		m["sishowalldb"] = strconv.FormatBool(b)
	})
	return w
}

// numberingroup
// Lists the number of users in user groups.
func (w *SiteinfoClient) Numberingroup(b bool) *SiteinfoClient {
	w.o = append(w.o, func(m map[string]string) {
		m["sinumberingroup"] = strconv.FormatBool(b)
	})
	return w
}

// inlanguagecode
// Language code for localised language names (best effort) and skin names.
func (w *SiteinfoClient) Inlanguagecode(s string) *SiteinfoClient {
	w.o = append(w.o, func(m map[string]string) {
		m["siinlanguagecode"] = s
	})
	return w
}

func (w *SiteinfoClient) Do(ctx context.Context) (SiteinfoResponse, error) {
	if err := w.c.checkKeepAlive(ctx); err != nil {
		return SiteinfoResponse{}, err
	}

	// Specify parameters to send.
	parameters := Values{
		"action": "query",
		"meta":   "siteinfo",
	}

	for _, o := range w.o {
		o(parameters)
	}

	// Make the request.
	r := SiteinfoResponse{}
	j, err := w.c.GetInto(ctx, parameters, &r)
	r.RawJSON = j
	if err != nil {
		return r, fmt.Errorf("failed to get: %w", err)
	}

	if e := r.Error; e != nil {
		return r, fmt.Errorf("%s: %s", e.Code, e.Info)
	}

	return r, nil
}
