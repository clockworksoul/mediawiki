package mediawiki

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSiteinfoAll(t *testing.T) {
	c, err := New(apiUrl, agent)
	require.NoError(t, err)

	_, err = c.BotLogin(context.Background(), username, password)
	require.NoError(t, err)

	ctx := context.Background()
	props := []string{
		"",
		SiteinfoPropGeneral,
		SiteinfoPropNamespaces,
		SiteinfoPropNamespacealiases,
		SiteinfoPropSpecialpagealiases,
		SiteinfoPropMagicwords,
		SiteinfoPropInterwikimap,
		SiteinfoPropDbrepllag,
		SiteinfoPropStatistics,
		SiteinfoPropUsergroups,
		SiteinfoPropLibraries,
		SiteinfoPropExtensions,
		SiteinfoPropFileextensions,
		SiteinfoPropRightsinfo,
		SiteinfoPropRestrictions,
		SiteinfoPropLanguages,
		SiteinfoPropLanguagevariants,
		SiteinfoPropSkins,
		SiteinfoPropExtensiontags,
		SiteinfoPropFunctionhooks,
		SiteinfoPropShowhooks,
		SiteinfoPropVariables,
		SiteinfoPropProtocols,
		SiteinfoPropDefaultoptions,
		SiteinfoPropUploaddialog,
	}

	for _, prop := range props {
		c := c.Siteinfo()
		if prop != "" {
			c = c.Prop(prop)
		}
		r, err := c.Do(ctx)

		require.NoError(t, err)
		assert.Nil(t, r.Error)
		require.NotNil(t, r.Query)

		CompareJSON(t, r.RawJSON, r, false)
	}
}
