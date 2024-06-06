package utils

const TmplDirURL = "./internal/web/templates"

const (
	BaseTmplURL   = TmplDirURL + "/base.html"
	NavbarTmplURL = TmplDirURL + "/navbar.html"
	FooterTmplURL = TmplDirURL + "/footer.html"
)

type PageData struct {
	Title      string
	Navbar     interface{}
	Footer     interface{}
	Content    interface{}
	FormData   interface{}
	FormErrors map[string]string
	ExtraCSS   []string
	ExtraJS    []string
	BaseURL    string
}

type NavbarData struct {
	Links  []Link
	Active string
}

type FooterData struct {
	CopyrightText string
	ExtraInfo     string
}

type Link struct {
	Href  string
	Label string
}
