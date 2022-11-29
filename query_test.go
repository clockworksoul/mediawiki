package mediawiki

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponsePageFullPageName(t *testing.T) {
	cc := []struct {
		Name     string
		Expected string
	}{
		{"Main page", "Main page"},
		{"Main_page", "Main page"},
		{"Talk:Main page", "Talk:Main page"},
		{"YP:FOO", "YP:FOO"},
		{"User:Mtitmus", "User:Mtitmus"},
		{"User talk:Mtitmus", "User talk:Mtitmus"},
		{"User_talk:Mtitmus", "User talk:Mtitmus"},
		{"User_talk:Foo_bar", "User talk:Foo bar"},
	}

	for _, c := range cc {
		p := QueryResponseQueryPage{Title: c.Name}
		assert.Equal(t, c.Expected, p.FullPageName())
	}
}

func TestResponsePagePageName(t *testing.T) {
	cc := []struct {
		Name      string
		Namespace int
		Expected  string
	}{
		{"Main page", NamespaceMain, "Main page"},
		{"Main_page", NamespaceMain, "Main page"},
		{"Talk:Main page", NamespaceTalk, "Main page"},
		{"YP:FOO", NamespaceMain, "YP:FOO"}, // In main namespace
		{"User:Mtitmus", NamespaceUser, "Mtitmus"},
		{"User talk:Mtitmus", NamespaceUserTalk, "Mtitmus"},
		{"User_talk:Mtitmus", NamespaceUserTalk, "Mtitmus"},
		{"User_talk:Foo_bar", NamespaceUserTalk, "Foo bar"},
		{"User:Mtitmus/sandbox", NamespaceUser, "Mtitmus/sandbox"},
		{"Help:Title/Foo/Bar", NamespaceHelp, "Title/Foo/Bar"},
	}

	for _, c := range cc {
		p := QueryResponseQueryPage{Title: c.Name, Namespace: c.Namespace}
		assert.Equal(t, c.Expected, p.PageName())
	}
}

func TestResponsePageBasePageName(t *testing.T) {
	cc := []struct {
		Name      string
		Namespace int
		Expected  string
	}{
		{"Main page", NamespaceMain, "Main page"},
		{"Main_page", NamespaceMain, "Main page"},
		{"Talk:Main page", NamespaceTalk, "Main page"},
		{"YP:FOO", NamespaceMain, "YP:FOO"}, // In main namespace
		{"User:Mtitmus", NamespaceUser, "Mtitmus"},
		{"User talk:Mtitmus", NamespaceUserTalk, "Mtitmus"},
		{"User_talk:Mtitmus", NamespaceUserTalk, "Mtitmus"},
		{"User_talk:Foo_bar", NamespaceUserTalk, "Foo bar"},
		{"User:Mtitmus/sandbox", NamespaceUser, "Mtitmus"},
		{"Help:Title/Foo/Bar", NamespaceHelp, "Title/Foo"},
	}

	for _, c := range cc {
		p := QueryResponseQueryPage{Title: c.Name, Namespace: c.Namespace}
		assert.Equal(t, c.Expected, p.BasePageName())
	}
}

func TestResponsePageRootPageName(t *testing.T) {
	cc := []struct {
		Name      string
		Namespace int
		Expected  string
	}{
		{"Main page", NamespaceMain, "Main page"},
		{"Main_page", NamespaceMain, "Main page"},
		{"Talk:Main page", NamespaceTalk, "Main page"},
		{"YP:FOO", NamespaceMain, "YP:FOO"}, // In main namespace
		{"User:Mtitmus", NamespaceUser, "Mtitmus"},
		{"User talk:Mtitmus", NamespaceUserTalk, "Mtitmus"},
		{"User_talk:Mtitmus", NamespaceUserTalk, "Mtitmus"},
		{"User_talk:Foo_bar", NamespaceUserTalk, "Foo bar"},
		{"User:Mtitmus/sandbox", NamespaceUser, "Mtitmus"},
		{"Help:Title/Foo/Bar", NamespaceHelp, "Title"},
	}

	for _, c := range cc {
		p := QueryResponseQueryPage{Title: c.Name, Namespace: c.Namespace}
		assert.Equal(t, c.Expected, p.RootPageName())
	}
}

func TestResponsePageArticlePageName(t *testing.T) {
	cc := []struct {
		Name      string
		Namespace int
		Expected  string
	}{
		{"Main page", NamespaceMain, "Main page"},
		{"Main_page", NamespaceMain, "Main page"},
		{"Talk:Main page", NamespaceTalk, "Main page"},
		{"YP:FOO", NamespaceMain, "YP:FOO"}, // In main namespace
		{"User:Mtitmus", NamespaceUser, "User:Mtitmus"},
		{"User talk:Mtitmus", NamespaceUserTalk, "User:Mtitmus"},
		{"User_talk:Mtitmus", NamespaceUserTalk, "User:Mtitmus"},
		{"User_talk:Foo_bar", NamespaceUserTalk, "User:Foo bar"},
		{"User:Mtitmus/sandbox", NamespaceUser, "User:Mtitmus/sandbox"},
		{"Help:Title/Foo/Bar", NamespaceHelp, "Help:Title/Foo/Bar"},
		{"Help talk:Title/Foo/Bar", NamespaceHelpTalk, "Help:Title/Foo/Bar"},
	}

	for _, c := range cc {
		p := QueryResponseQueryPage{Title: c.Name, Namespace: c.Namespace}
		assert.Equal(t, c.Expected, p.ArticlePageName())
	}
}

func TestResponsePageTalkPageName(t *testing.T) {
	cc := []struct {
		Name      string
		Namespace int
		Expected  string
	}{
		{"Main page", NamespaceMain, "Talk:Main page"},
		{"Main_page", NamespaceMain, "Talk:Main page"},
		{"Talk:Main page", NamespaceTalk, "Talk:Main page"},
		{"YP:FOO", NamespaceMain, "Talk:YP:FOO"}, // In main namespace
		{"User:Mtitmus", NamespaceUser, "User talk:Mtitmus"},
		{"User talk:Mtitmus", NamespaceUserTalk, "User talk:Mtitmus"},
		{"User_talk:Mtitmus", NamespaceUserTalk, "User talk:Mtitmus"},
		{"User_talk:Foo_bar", NamespaceUserTalk, "User talk:Foo bar"},
		{"User:Mtitmus/sandbox", NamespaceUser, "User talk:Mtitmus/sandbox"},
		{"Help:Title/Foo/Bar", NamespaceHelp, "Help talk:Title/Foo/Bar"},
		{"Help talk:Title/Foo/Bar", NamespaceHelpTalk, "Help talk:Title/Foo/Bar"},
	}

	for _, c := range cc {
		p := QueryResponseQueryPage{Title: c.Name, Namespace: c.Namespace}
		assert.Equal(t, c.Expected, p.TalkPageName())
	}
}
