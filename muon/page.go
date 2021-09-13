package muon

type PageHeaderParams struct {
	AvailableLanguages []string
	NavigationLinks    []ContentLink
}

type PageFooterParams struct {
	ContactUri string
}

type PageParams struct {
	Header             PageHeaderParams
	Footer             PageFooterParams
	CurrentArticlePath string
	Language           string
}

func NewPageParams(config *ContentConfiguration, language string, currentArticlePath string) *PageParams {
	return &PageParams{
		Header: PageHeaderParams{
			AvailableLanguages: config.Languages,
			NavigationLinks:    config.Links,
		},
		Footer: PageFooterParams{
			ContactUri: config.ContactUri,
		},
		CurrentArticlePath: currentArticlePath,
		Language:           language,
	}
}
