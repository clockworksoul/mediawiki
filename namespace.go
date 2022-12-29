package mediawiki

type Namespace int

const (
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

func (n Namespace) String() string {
	switch n {
	case NamespaceMedia:
		return "Media"
	case NamespaceSpecial:
		return "Special"
	case NamespaceMain:
		return "Main"
	case NamespaceTalk:
		return "Talk"
	case NamespaceUser:
		return "User"
	case NamespaceUserTalk:
		return "User_talk"
	case NamespaceProject:
		return "Project"
	case NamespaceProjectTalk:
		return "Project_talk"
	case NamespaceFile:
		return "File"
	case NamespaceFileTalk:
		return "File_talk"
	case NamespaceMediaWiki:
		return "MediaWiki"
	case NamespaceMediaWikiTalk:
		return "MediaWiki_talk"
	case NamespaceTemplate:
		return "Template"
	case NamespaceTemplateTalk:
		return "Template_talk"
	case NamespaceHelp:
		return "Help"
	case NamespaceHelpTalk:
		return "Help_talk"
	case NamespaceCategory:
		return "Category"
	case NamespaceCategoryTalk:
		return "Category_talk"
	default:
		if n%2 == 0 {
			return "Custom"
		} else {
			return "Custom talk"
		}
	}
}
