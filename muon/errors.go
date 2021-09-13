package muon

import (
	"errors"
	"net/http"
)

var PageNotFound = errors.New("Page not found")

type ErrorPageParams struct {
	Page *PageParams
}

func RenderNotFoundPage(srv *Server, res http.ResponseWriter, req *http.Request) error {
	res.WriteHeader(http.StatusNotFound)
	language := req.FormValue("language")
	if language == "" {
		language = srv.Configuration.ContentConfig.DefaultLanguage
	}
	return srv.Templates.ExecuteTemplate(res, "404.html", ErrorPageParams{
		Page: NewPageParams(srv.Configuration.ContentConfig, language, ""),
	})
}

func RenderInternalErrorPage(srv *Server, res http.ResponseWriter, req *http.Request) error {
	res.WriteHeader(http.StatusInternalServerError)
	language := req.FormValue("language")
	if language == "" {
		language = srv.Configuration.ContentConfig.DefaultLanguage
	}
	return srv.Templates.ExecuteTemplate(res, "500.html", ErrorPageParams{
		Page: NewPageParams(srv.Configuration.ContentConfig, language, ""),
	})
}
