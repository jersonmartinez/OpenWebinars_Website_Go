package models

import "html/template"

type PageData struct {
	Title         string
	Author        string
	Welcome       string
	ErrorCode     int
	ErrorMessage  string
	HeadContent   template.HTML
	NavbarContent template.HTML
}
