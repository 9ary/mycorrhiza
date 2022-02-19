// Code generated by qtc from "readers.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/readers.qtpl:1
package views

//line views/readers.qtpl:1
import "net/http"

//line views/readers.qtpl:2
import "strings"

//line views/readers.qtpl:3
import "path"

//line views/readers.qtpl:4
import "os"

//line views/readers.qtpl:6
import "github.com/bouncepaw/mycorrhiza/cfg"

//line views/readers.qtpl:7
import "github.com/bouncepaw/mycorrhiza/hyphae"

//line views/readers.qtpl:8
import "github.com/bouncepaw/mycorrhiza/l18n"

//line views/readers.qtpl:9
import "github.com/bouncepaw/mycorrhiza/mimetype"

//line views/readers.qtpl:10
import "github.com/bouncepaw/mycorrhiza/tree"

//line views/readers.qtpl:11
import "github.com/bouncepaw/mycorrhiza/user"

//line views/readers.qtpl:12
import "github.com/bouncepaw/mycorrhiza/util"

//line views/readers.qtpl:14
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/readers.qtpl:14
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/readers.qtpl:14
func StreamAttachmentMenuHTML(qw422016 *qt422016.Writer, rq *http.Request, h hyphae.Hypha, u *user.User) {
//line views/readers.qtpl:14
	qw422016.N().S(`
`)
//line views/readers.qtpl:16
	lc := l18n.FromRequest(rq)

//line views/readers.qtpl:17
	qw422016.N().S(`
<div class="layout">
<main class="main-width attachment-tab">
	<h1>`)
//line views/readers.qtpl:20
	qw422016.N().S(lc.Get("ui.attach_title", &l18n.Replacements{"name": beautifulLink(h.CanonicalName())}))
//line views/readers.qtpl:20
	qw422016.N().S(`</h1>
	`)
//line views/readers.qtpl:21
	switch h.(type) {
//line views/readers.qtpl:22
	case *hyphae.MediaHypha:
//line views/readers.qtpl:22
		qw422016.N().S(`
		<p class="explanation">`)
//line views/readers.qtpl:23
		qw422016.E().S(lc.Get("ui.attach_tip"))
//line views/readers.qtpl:23
		qw422016.N().S(` <a href="/help/en/attachment" class="shy-link">`)
//line views/readers.qtpl:23
		qw422016.E().S(lc.Get("ui.attach_link"))
//line views/readers.qtpl:23
		qw422016.N().S(`</a></p>
	`)
//line views/readers.qtpl:24
	default:
//line views/readers.qtpl:24
		qw422016.N().S(`
		<p class="explanation">`)
//line views/readers.qtpl:25
		qw422016.E().S(lc.Get("ui.attach_empty"))
//line views/readers.qtpl:25
		qw422016.N().S(` <a href="/help/en/attachment" class="shy-link">`)
//line views/readers.qtpl:25
		qw422016.E().S(lc.Get("ui.attach_link"))
//line views/readers.qtpl:25
		qw422016.N().S(`</a></p>
	`)
//line views/readers.qtpl:26
	}
//line views/readers.qtpl:26
	qw422016.N().S(`

	<section class="amnt-grid">
	`)
//line views/readers.qtpl:29
	switch h := h.(type) {
//line views/readers.qtpl:30
	case *hyphae.MediaHypha:
//line views/readers.qtpl:30
		qw422016.N().S(`
		`)
//line views/readers.qtpl:32
		mime := mimetype.FromExtension(path.Ext(h.MediaFilePath()))
		fileinfo, err := os.Stat(h.MediaFilePath())

//line views/readers.qtpl:33
		qw422016.N().S(`
		`)
//line views/readers.qtpl:34
		if err == nil {
//line views/readers.qtpl:34
			qw422016.N().S(`
		<fieldset class="amnt-menu-block">
			<legend class="modal__title modal__title_small">`)
//line views/readers.qtpl:36
			qw422016.E().S(lc.Get("ui.attach_stat"))
//line views/readers.qtpl:36
			qw422016.N().S(`</legend>
			<p class="modal__confirmation-msg"><b>`)
//line views/readers.qtpl:37
			qw422016.E().S(lc.Get("ui.attach_stat_size"))
//line views/readers.qtpl:37
			qw422016.N().S(`</b> `)
//line views/readers.qtpl:37
			qw422016.E().S(lc.GetPlural64("ui.attach_size_value", fileinfo.Size()))
//line views/readers.qtpl:37
			qw422016.N().S(`</p>
			<p><b>`)
//line views/readers.qtpl:38
			qw422016.E().S(lc.Get("ui.attach_stat_mime"))
//line views/readers.qtpl:38
			qw422016.N().S(`</b> `)
//line views/readers.qtpl:38
			qw422016.E().S(mime)
//line views/readers.qtpl:38
			qw422016.N().S(`</p>
		</fieldset>
		`)
//line views/readers.qtpl:40
		}
//line views/readers.qtpl:40
		qw422016.N().S(`

		`)
//line views/readers.qtpl:42
		if strings.HasPrefix(mime, "image/") {
//line views/readers.qtpl:42
			qw422016.N().S(`
		<fieldset class="amnt-menu-block">
			<legend class="modal__title modal__title_small">`)
//line views/readers.qtpl:44
			qw422016.E().S(lc.Get("ui.attach_include"))
//line views/readers.qtpl:44
			qw422016.N().S(`</legend>
			<p class="modal__confirmation-msg">`)
//line views/readers.qtpl:45
			qw422016.E().S(lc.Get("ui.attach_include_tip"))
//line views/readers.qtpl:45
			qw422016.N().S(`</p>
			<pre class="codeblock"><code>img { `)
//line views/readers.qtpl:46
			qw422016.E().S(h.CanonicalName())
//line views/readers.qtpl:46
			qw422016.N().S(` }</code></pre>
		</fieldset>
		`)
//line views/readers.qtpl:48
		}
//line views/readers.qtpl:48
		qw422016.N().S(`
	`)
//line views/readers.qtpl:49
	}
//line views/readers.qtpl:49
	qw422016.N().S(`

	`)
//line views/readers.qtpl:51
	if u.CanProceed("upload-binary") {
//line views/readers.qtpl:51
		qw422016.N().S(`
	<form action="/upload-binary/`)
//line views/readers.qtpl:52
		qw422016.E().S(h.CanonicalName())
//line views/readers.qtpl:52
		qw422016.N().S(`"
			method="post" enctype="multipart/form-data"
			class="upload-binary modal amnt-menu-block">
		<fieldset class="modal__fieldset">
			<legend class="modal__title modal__title_small">`)
//line views/readers.qtpl:56
		qw422016.E().S(lc.Get("ui.attach_new"))
//line views/readers.qtpl:56
		qw422016.N().S(`</legend>
			<p class="modal__confirmation-msg">`)
//line views/readers.qtpl:57
		qw422016.E().S(lc.Get("ui.attach_new_tip"))
//line views/readers.qtpl:57
		qw422016.N().S(`</p>
			<label for="upload-binary__input"></label>
			<input type="file" id="upload-binary__input" name="binary">

			<button type="submit" class="btn stick-to-bottom" value="Upload">`)
//line views/readers.qtpl:61
		qw422016.E().S(lc.Get("ui.attach_upload"))
//line views/readers.qtpl:61
		qw422016.N().S(`</button>
		</fieldset>
	</form>
	`)
//line views/readers.qtpl:64
	}
//line views/readers.qtpl:64
	qw422016.N().S(`


	`)
//line views/readers.qtpl:67
	switch h := h.(type) {
//line views/readers.qtpl:68
	case *hyphae.MediaHypha:
//line views/readers.qtpl:68
		qw422016.N().S(`
		`)
//line views/readers.qtpl:69
		if u.CanProceed("unattach-confirm") {
//line views/readers.qtpl:69
			qw422016.N().S(`
		<form action="/unattach-confirm/`)
//line views/readers.qtpl:70
			qw422016.E().S(h.CanonicalName())
//line views/readers.qtpl:70
			qw422016.N().S(`" method="post" class="modal amnt-menu-block">
			<fieldset class="modal__fieldset">
				<legend class="modal__title modal__title_small">`)
//line views/readers.qtpl:72
			qw422016.E().S(lc.Get("ui.attach_remove"))
//line views/readers.qtpl:72
			qw422016.N().S(`</legend>
				<p class="modal__confirmation-msg">`)
//line views/readers.qtpl:73
			qw422016.E().S(lc.Get("ui.attach_remove_tip"))
//line views/readers.qtpl:73
			qw422016.N().S(`</p>
				<button type="submit" class="btn" value="Unattach">`)
//line views/readers.qtpl:74
			qw422016.E().S(lc.Get("ui.attach_remove_button"))
//line views/readers.qtpl:74
			qw422016.N().S(`</button>
			</fieldset>
		</form>
		`)
//line views/readers.qtpl:77
		}
//line views/readers.qtpl:77
		qw422016.N().S(`
	`)
//line views/readers.qtpl:78
	}
//line views/readers.qtpl:78
	qw422016.N().S(`

	</section>
</main>
</div>
`)
//line views/readers.qtpl:83
}

//line views/readers.qtpl:83
func WriteAttachmentMenuHTML(qq422016 qtio422016.Writer, rq *http.Request, h hyphae.Hypha, u *user.User) {
//line views/readers.qtpl:83
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/readers.qtpl:83
	StreamAttachmentMenuHTML(qw422016, rq, h, u)
//line views/readers.qtpl:83
	qt422016.ReleaseWriter(qw422016)
//line views/readers.qtpl:83
}

//line views/readers.qtpl:83
func AttachmentMenuHTML(rq *http.Request, h hyphae.Hypha, u *user.User) string {
//line views/readers.qtpl:83
	qb422016 := qt422016.AcquireByteBuffer()
//line views/readers.qtpl:83
	WriteAttachmentMenuHTML(qb422016, rq, h, u)
//line views/readers.qtpl:83
	qs422016 := string(qb422016.B)
//line views/readers.qtpl:83
	qt422016.ReleaseByteBuffer(qb422016)
//line views/readers.qtpl:83
	return qs422016
//line views/readers.qtpl:83
}

// If `contents` == "", a helpful message is shown instead.
//
// If you rename .prevnext, change the docs too.

//line views/readers.qtpl:88
func StreamHyphaHTML(qw422016 *qt422016.Writer, rq *http.Request, lc *l18n.Localizer, h hyphae.Hypha, contents string) {
//line views/readers.qtpl:88
	qw422016.N().S(`
`)
//line views/readers.qtpl:90
	siblings, subhyphae, prevHyphaName, nextHyphaName := tree.Tree(h.CanonicalName())
	u := user.FromRequest(rq)

//line views/readers.qtpl:92
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
	<section id="hypha">
		`)
//line views/readers.qtpl:96
	if u.CanProceed("edit") {
//line views/readers.qtpl:96
		qw422016.N().S(`
		<div class="btn btn_navititle">
			<a class="btn__link_navititle" href="/edit/`)
//line views/readers.qtpl:98
		qw422016.E().S(h.CanonicalName())
//line views/readers.qtpl:98
		qw422016.N().S(`">`)
//line views/readers.qtpl:98
		qw422016.E().S(lc.Get("ui.edit_link"))
//line views/readers.qtpl:98
		qw422016.N().S(`</a>
		</div>
		`)
//line views/readers.qtpl:100
	}
//line views/readers.qtpl:100
	qw422016.N().S(`

		`)
//line views/readers.qtpl:102
	if cfg.UseAuth && util.IsProfileName(h.CanonicalName()) && u.Name == strings.TrimPrefix(h.CanonicalName(), cfg.UserHypha+"/") {
//line views/readers.qtpl:102
		qw422016.N().S(`
		<div class="btn btn_navititle">
			<a class="btn__link_navititle" href="/logout">`)
//line views/readers.qtpl:104
		qw422016.E().S(lc.Get("ui.logout_link"))
//line views/readers.qtpl:104
		qw422016.N().S(`</a>
		</div>
		`)
//line views/readers.qtpl:106
		if u.Group == "admin" {
//line views/readers.qtpl:106
			qw422016.N().S(`
		<div class="btn btn_navititle">
			<a class="btn__link_navititle" href="/admin">`)
//line views/readers.qtpl:108
			qw422016.E().S(lc.Get("ui.admin_panel"))
//line views/readers.qtpl:108
			qw422016.N().S(`<a>
		</div>
		`)
//line views/readers.qtpl:110
		}
//line views/readers.qtpl:110
		qw422016.N().S(`
		`)
//line views/readers.qtpl:111
	}
//line views/readers.qtpl:111
	qw422016.N().S(`

		`)
//line views/readers.qtpl:113
	qw422016.N().S(NaviTitleHTML(h))
//line views/readers.qtpl:113
	qw422016.N().S(`
		`)
//line views/readers.qtpl:114
	switch h.(type) {
//line views/readers.qtpl:115
	case *hyphae.EmptyHypha:
//line views/readers.qtpl:115
		qw422016.N().S(`
				`)
//line views/readers.qtpl:116
		streamnonExistentHyphaNotice(qw422016, h, u, lc)
//line views/readers.qtpl:116
		qw422016.N().S(`
			`)
//line views/readers.qtpl:117
	default:
//line views/readers.qtpl:117
		qw422016.N().S(`
				`)
//line views/readers.qtpl:118
		qw422016.N().S(contents)
//line views/readers.qtpl:118
		qw422016.N().S(`
		`)
//line views/readers.qtpl:119
	}
//line views/readers.qtpl:119
	qw422016.N().S(`
	</section>
	<section class="prevnext">
		`)
//line views/readers.qtpl:122
	if prevHyphaName != "" {
//line views/readers.qtpl:122
		qw422016.N().S(`
		<a class="prevnext__el prevnext__prev" href="/hypha/`)
//line views/readers.qtpl:123
		qw422016.E().S(prevHyphaName)
//line views/readers.qtpl:123
		qw422016.N().S(`" rel="prev">← `)
//line views/readers.qtpl:123
		qw422016.E().S(util.BeautifulName(path.Base(prevHyphaName)))
//line views/readers.qtpl:123
		qw422016.N().S(`</a>
		`)
//line views/readers.qtpl:124
	}
//line views/readers.qtpl:124
	qw422016.N().S(`
		`)
//line views/readers.qtpl:125
	if nextHyphaName != "" {
//line views/readers.qtpl:125
		qw422016.N().S(`
		<a class="prevnext__el prevnext__next" href="/hypha/`)
//line views/readers.qtpl:126
		qw422016.E().S(nextHyphaName)
//line views/readers.qtpl:126
		qw422016.N().S(`" rel="next">`)
//line views/readers.qtpl:126
		qw422016.E().S(util.BeautifulName(path.Base(nextHyphaName)))
//line views/readers.qtpl:126
		qw422016.N().S(` →</a>
		`)
//line views/readers.qtpl:127
	}
//line views/readers.qtpl:127
	qw422016.N().S(`
	</section>
`)
//line views/readers.qtpl:129
	StreamSubhyphaeHTML(qw422016, subhyphae, lc)
//line views/readers.qtpl:129
	qw422016.N().S(`
	<section id="hypha-bottom">
   		`)
//line views/readers.qtpl:131
	streamhyphaInfo(qw422016, rq, h)
//line views/readers.qtpl:131
	qw422016.N().S(`
	</section>
</main>
`)
//line views/readers.qtpl:134
	streamsiblingHyphaeHTML(qw422016, siblings, lc)
//line views/readers.qtpl:134
	qw422016.N().S(`
</div>
`)
//line views/readers.qtpl:136
	streamviewScripts(qw422016)
//line views/readers.qtpl:136
	qw422016.N().S(`
`)
//line views/readers.qtpl:137
}

//line views/readers.qtpl:137
func WriteHyphaHTML(qq422016 qtio422016.Writer, rq *http.Request, lc *l18n.Localizer, h hyphae.Hypha, contents string) {
//line views/readers.qtpl:137
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/readers.qtpl:137
	StreamHyphaHTML(qw422016, rq, lc, h, contents)
//line views/readers.qtpl:137
	qt422016.ReleaseWriter(qw422016)
//line views/readers.qtpl:137
}

//line views/readers.qtpl:137
func HyphaHTML(rq *http.Request, lc *l18n.Localizer, h hyphae.Hypha, contents string) string {
//line views/readers.qtpl:137
	qb422016 := qt422016.AcquireByteBuffer()
//line views/readers.qtpl:137
	WriteHyphaHTML(qb422016, rq, lc, h, contents)
//line views/readers.qtpl:137
	qs422016 := string(qb422016.B)
//line views/readers.qtpl:137
	qt422016.ReleaseByteBuffer(qb422016)
//line views/readers.qtpl:137
	return qs422016
//line views/readers.qtpl:137
}

//line views/readers.qtpl:139
func StreamRevisionHTML(qw422016 *qt422016.Writer, rq *http.Request, lc *l18n.Localizer, h hyphae.Hypha, contents, revHash string) {
//line views/readers.qtpl:139
	qw422016.N().S(`
`)
//line views/readers.qtpl:141
	siblings, subhyphae, _, _ := tree.Tree(h.CanonicalName())

//line views/readers.qtpl:142
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
	<section>
		<p>`)
//line views/readers.qtpl:146
	qw422016.E().S(lc.Get("ui.revision_warning"))
//line views/readers.qtpl:146
	qw422016.N().S(` <a href="/rev-text/`)
//line views/readers.qtpl:146
	qw422016.E().S(revHash)
//line views/readers.qtpl:146
	qw422016.N().S(`/`)
//line views/readers.qtpl:146
	qw422016.E().S(h.CanonicalName())
//line views/readers.qtpl:146
	qw422016.N().S(`">`)
//line views/readers.qtpl:146
	qw422016.E().S(lc.Get("ui.revision_link"))
//line views/readers.qtpl:146
	qw422016.N().S(`</a></p>
		`)
//line views/readers.qtpl:147
	qw422016.N().S(NaviTitleHTML(h))
//line views/readers.qtpl:147
	qw422016.N().S(`
		`)
//line views/readers.qtpl:148
	qw422016.N().S(contents)
//line views/readers.qtpl:148
	qw422016.N().S(`
	</section>
`)
//line views/readers.qtpl:150
	StreamSubhyphaeHTML(qw422016, subhyphae, lc)
//line views/readers.qtpl:150
	qw422016.N().S(`
</main>
`)
//line views/readers.qtpl:152
	streamsiblingHyphaeHTML(qw422016, siblings, lc)
//line views/readers.qtpl:152
	qw422016.N().S(`
</div>
`)
//line views/readers.qtpl:154
	streamviewScripts(qw422016)
//line views/readers.qtpl:154
	qw422016.N().S(`
`)
//line views/readers.qtpl:155
}

//line views/readers.qtpl:155
func WriteRevisionHTML(qq422016 qtio422016.Writer, rq *http.Request, lc *l18n.Localizer, h hyphae.Hypha, contents, revHash string) {
//line views/readers.qtpl:155
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/readers.qtpl:155
	StreamRevisionHTML(qw422016, rq, lc, h, contents, revHash)
//line views/readers.qtpl:155
	qt422016.ReleaseWriter(qw422016)
//line views/readers.qtpl:155
}

//line views/readers.qtpl:155
func RevisionHTML(rq *http.Request, lc *l18n.Localizer, h hyphae.Hypha, contents, revHash string) string {
//line views/readers.qtpl:155
	qb422016 := qt422016.AcquireByteBuffer()
//line views/readers.qtpl:155
	WriteRevisionHTML(qb422016, rq, lc, h, contents, revHash)
//line views/readers.qtpl:155
	qs422016 := string(qb422016.B)
//line views/readers.qtpl:155
	qt422016.ReleaseByteBuffer(qb422016)
//line views/readers.qtpl:155
	return qs422016
//line views/readers.qtpl:155
}

//line views/readers.qtpl:157
func streamviewScripts(qw422016 *qt422016.Writer) {
//line views/readers.qtpl:157
	qw422016.N().S(`
`)
//line views/readers.qtpl:158
	for _, scriptPath := range cfg.ViewScripts {
//line views/readers.qtpl:158
		qw422016.N().S(`
<script src="`)
//line views/readers.qtpl:159
		qw422016.E().S(scriptPath)
//line views/readers.qtpl:159
		qw422016.N().S(`"></script>
`)
//line views/readers.qtpl:160
	}
//line views/readers.qtpl:160
	qw422016.N().S(`
`)
//line views/readers.qtpl:161
}

//line views/readers.qtpl:161
func writeviewScripts(qq422016 qtio422016.Writer) {
//line views/readers.qtpl:161
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/readers.qtpl:161
	streamviewScripts(qw422016)
//line views/readers.qtpl:161
	qt422016.ReleaseWriter(qw422016)
//line views/readers.qtpl:161
}

//line views/readers.qtpl:161
func viewScripts() string {
//line views/readers.qtpl:161
	qb422016 := qt422016.AcquireByteBuffer()
//line views/readers.qtpl:161
	writeviewScripts(qb422016)
//line views/readers.qtpl:161
	qs422016 := string(qb422016.B)
//line views/readers.qtpl:161
	qt422016.ReleaseByteBuffer(qb422016)
//line views/readers.qtpl:161
	return qs422016
//line views/readers.qtpl:161
}
