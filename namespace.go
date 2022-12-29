package mediawiki

type Namespace int64

const (
	// Namespace all is a special value used to indicate "all namespaces"
	// for functions like LinkshereClient.Namespace.
	NamespaceAll Namespace = -1024

	NamespaceMedia         Namespace = -2
	NamespaceSpecial       Namespace = -1
	NamespaceMain          Namespace = 0
	NamespaceTalk          Namespace = 1
	NamespaceUser          Namespace = 2
	NamespaceUserTalk      Namespace = 3
	NamespaceProject       Namespace = 4
	NamespaceProjectTalk   Namespace = 5
	NamespaceFile          Namespace = 6
	NamespaceFileTalk      Namespace = 7
	NamespaceMediaWiki     Namespace = 8
	NamespaceMediaWikiTalk Namespace = 9
	NamespaceTemplate      Namespace = 10
	NamespaceTemplateTalk  Namespace = 11
	NamespaceHelp          Namespace = 12
	NamespaceHelpTalk      Namespace = 13
	NamespaceCategory      Namespace = 14
	NamespaceCategoryTalk  Namespace = 15
)

var namespaceNames = map[Namespace]string{
	NamespaceAll:           "*",
	NamespaceMedia:         "Media",
	NamespaceSpecial:       "Special",
	NamespaceMain:          "Main",
	NamespaceTalk:          "Talk",
	NamespaceUser:          "User",
	NamespaceUserTalk:      "User_talk",
	NamespaceProject:       "Project",
	NamespaceProjectTalk:   "Project_talk",
	NamespaceFile:          "File",
	NamespaceFileTalk:      "File_talk",
	NamespaceMediaWiki:     "MediaWiki",
	NamespaceMediaWikiTalk: "MediaWiki_talk",
	NamespaceTemplate:      "Template",
	NamespaceTemplateTalk:  "Template_talk",
	NamespaceHelp:          "Help",
	NamespaceHelpTalk:      "Help_talk",
	NamespaceCategory:      "Category",
	NamespaceCategoryTalk:  "Category_talk",
}

// NewNamespace is a convenience function that simplifies the definition of
// named namespaces. Using NewNamespace will add the code and name to the
// internal lookup table, so that future calls to String() will return the
// assigned name. Note that this could be used to (accidentally or otherwise)
// redefine the standard namespace names.
func NewNamespace(name string, code int) Namespace {
	ns := Namespace(code)
	namespaceNames[ns] = name
	return ns
}

func (n Namespace) String() string {
	if name, ok := namespaceNames[n]; ok {
		return name
	} else if n%2 == 0 {
		return "Custom"
	} else {
		return "Custom_talk"
	}
}
