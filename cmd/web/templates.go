package main

import "github.com/ejacobg/snippetbox/internal/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
