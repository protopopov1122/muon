package muon

import (
	"errors"
	"html/template"
	"net/http"
	"net/url"
	"path"
	"path/filepath"

	"github.com/gorilla/mux"
)

type Server struct {
	Logger        *Logger
	Configuration *ServiceConfiguration
	Templates     *template.Template
	Localization  *Localization
	router        *mux.Router
}

func buildArticleUrl(language string, articlePath string) string {
	articleUrl := url.URL{
		Path: path.Join("/article", articlePath),
	}
	if language != "" {
		query := url.Values{}
		query.Add("language", language)
		articleUrl.RawQuery = query.Encode()
	}
	return articleUrl.String()
}

func NewServer(logger *Logger, config *ServiceConfiguration) (*Server, error) {
	srv := Server{
		Logger:        logger,
		Configuration: config,
		Localization:  NewLocalization(config.ContentConfig.DefaultLanguage),
		Templates:     template.New(""),
		router:        mux.NewRouter(),
	}

	if err := srv.Localization.Load(logger, config.LocalizationRoot); err != nil {
		return nil, err
	}

	if err := srv.loadTemplates(); err != nil {
		return nil, err
	}

	srv.setupRoutes()
	logger.Info.Println("Server initial set up is done")
	return &srv, nil
}

func (srv *Server) setupRoutes() {
	srv.Logger.Info.Println("Setting up routing")
	srv.router.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir(srv.Configuration.StaticRoot))))

	articlePrefix := srv.router.PathPrefix("/article")
	articleHandler := http.StripPrefix("/article", http.HandlerFunc(srv.serveArticlePage))
	articlePrefix.Handler(articleHandler)

	srv.router.HandleFunc("/", srv.serveIndexPage)
	srv.router.NotFoundHandler = http.HandlerFunc(srv.serveNotFoundPage)
}

func (srv *Server) loadTemplates() error {
	srv.Logger.Info.Println("Loading templates from", srv.Configuration.TemplateRoot)
	tplPath := filepath.Join(srv.Configuration.TemplateRoot, "*.html")
	tplFuncMap := template.FuncMap{
		"localize":   srv.Localization.Localize,
		"articleUrl": buildArticleUrl,
	}

	srv.Templates.Funcs(tplFuncMap)
	_, err := srv.Templates.ParseGlob(tplPath)
	return err
}

func (srv *Server) serveIndexPage(res http.ResponseWriter, req *http.Request) {
	contentConfig := srv.Configuration.ContentConfig
	url := buildArticleUrl(contentConfig.DefaultLanguage, contentConfig.Index)
	http.Redirect(res, req, url, http.StatusPermanentRedirect)
}

func (srv *Server) serveArticlePage(res http.ResponseWriter, req *http.Request) {
	language := req.FormValue("language")
	if language == "" {
		language = srv.Configuration.ContentConfig.DefaultLanguage
	}

	err := RenderArticlePage(srv, res, language, req.URL.Path)
	if err == nil {
		srv.Logger.Info.Printf("%v - Success\n", req.URL.Path)
	} else if errors.Is(err, PageNotFound) {
		srv.serveNotFoundPage(res, req)
	} else {
		srv.serveInternalErrorPage(res, req, err)
	}
}

func (srv *Server) serveNotFoundPage(res http.ResponseWriter, req *http.Request) {
	srv.Logger.Warning.Printf("%v - Not found\n", req.URL.Path)
	err := RenderNotFoundPage(srv, res, req)
	if err != nil {
		srv.Logger.Error.Println("Failed to render 404 page:", err)
	}
}

func (srv *Server) serveInternalErrorPage(res http.ResponseWriter, req *http.Request, internalError error) {
	srv.Logger.Error.Printf("%v - Failure: %v\n", req.URL.Path, internalError)
	err := RenderInternalErrorPage(srv, res, req)
	if err != nil {
		srv.Logger.Error.Println("Failed to render 500 page:", err)
	}
}

func (srv *Server) Start() error {
	srv.Logger.Info.Println("Starting service on", srv.Configuration.BindTo)
	http.Handle("/", srv.router)
	return http.ListenAndServe(srv.Configuration.BindTo, nil)
}
