package muon

import (
	"errors"
	"html/template"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/gomarkdown/markdown"
)

type ArticlePageParams struct {
	Page    PageParams
	Content template.HTML
}

func renderArticleFile(filePath string) (string, error) {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	html := markdown.ToHTML(fileContent, nil, nil)
	return string(html), nil
}

func renderArticle(root string, language string, path string) (string, error) {
	filePath := filepath.Join(root, language, filepath.FromSlash(path)+".md")
	stat, err := os.Stat(filePath)
	if errors.Is(err, fs.ErrNotExist) {
		return "", PageNotFound
	}
	if err != nil {
		return "", err
	}

	if stat.IsDir() {
		return "", PageNotFound
	} else {
		return renderArticleFile(filePath)
	}
}

func RenderArticlePage(srv *Server, out io.Writer, language string, articlePath string) error {
	articlePath = path.Clean(articlePath)
	if articlePath == "/" || articlePath == "" {
		articlePath = srv.Configuration.ContentConfig.Index
	}
	articleContent, err := renderArticle(srv.Configuration.ArticlesRoot, language, articlePath)
	if err != nil {
		return err
	}

	params := ArticlePageParams{
		Page:    *NewPageParams(srv.Configuration.ContentConfig, language, articlePath),
		Content: template.HTML(articleContent),
	}
	err = srv.Templates.ExecuteTemplate(out, "article.html", params)
	return err
}
