// Code generated by qtc from "modal.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/modal.qtpl:1
package views

//line views/modal.qtpl:1
import "net/http"

//line views/modal.qtpl:2
import "github.com/bouncepaw/mycorrhiza/util"

//line views/modal.qtpl:4
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/modal.qtpl:4
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/modal.qtpl:4
func StreamDeleteAskHTML(qw422016 *qt422016.Writer, rq *http.Request, hyphaName string, isOld bool) {
//line views/modal.qtpl:4
	qw422016.N().S(`
`)
//line views/modal.qtpl:5
	StreamNavHTML(qw422016, rq, hyphaName, "delete-ask")
//line views/modal.qtpl:5
	qw422016.N().S(`
`)
//line views/modal.qtpl:6
	streammodalBegin(qw422016,
		"delete-confirm",
		hyphaName,
		"",
		"Delete "+util.BeautifulName(hyphaName)+"?")
//line views/modal.qtpl:10
	qw422016.N().S(`
`)
//line views/modal.qtpl:11
	streammodalReallyWant(qw422016, hyphaName, "unattach")
//line views/modal.qtpl:11
	qw422016.N().S(`
		<p>In this version of MycorrhizaWiki you cannot undelete a deleted hypha but the history can still be accessed.</p>
`)
//line views/modal.qtpl:13
	streammodalEnd(qw422016, hyphaName, true)
//line views/modal.qtpl:13
	qw422016.N().S(`
`)
//line views/modal.qtpl:14
}

//line views/modal.qtpl:14
func WriteDeleteAskHTML(qq422016 qtio422016.Writer, rq *http.Request, hyphaName string, isOld bool) {
//line views/modal.qtpl:14
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/modal.qtpl:14
	StreamDeleteAskHTML(qw422016, rq, hyphaName, isOld)
//line views/modal.qtpl:14
	qt422016.ReleaseWriter(qw422016)
//line views/modal.qtpl:14
}

//line views/modal.qtpl:14
func DeleteAskHTML(rq *http.Request, hyphaName string, isOld bool) string {
//line views/modal.qtpl:14
	qb422016 := qt422016.AcquireByteBuffer()
//line views/modal.qtpl:14
	WriteDeleteAskHTML(qb422016, rq, hyphaName, isOld)
//line views/modal.qtpl:14
	qs422016 := string(qb422016.B)
//line views/modal.qtpl:14
	qt422016.ReleaseByteBuffer(qb422016)
//line views/modal.qtpl:14
	return qs422016
//line views/modal.qtpl:14
}

//line views/modal.qtpl:16
func StreamUnattachAskHTML(qw422016 *qt422016.Writer, rq *http.Request, hyphaName string, isOld bool) {
//line views/modal.qtpl:16
	qw422016.N().S(`
`)
//line views/modal.qtpl:17
	StreamNavHTML(qw422016, rq, hyphaName, "unattach-ask")
//line views/modal.qtpl:17
	qw422016.N().S(`
`)
//line views/modal.qtpl:18
	streammodalBegin(qw422016,
		"unattach",
		hyphaName,
		"",
		"Unattach "+util.BeautifulName(hyphaName)+"?")
//line views/modal.qtpl:22
	qw422016.N().S(`
`)
//line views/modal.qtpl:23
	streammodalReallyWant(qw422016, hyphaName, "unattach")
//line views/modal.qtpl:23
	qw422016.N().S(`
`)
//line views/modal.qtpl:24
	streammodalEnd(qw422016, hyphaName, true)
//line views/modal.qtpl:24
	qw422016.N().S(`
`)
//line views/modal.qtpl:25
}

//line views/modal.qtpl:25
func WriteUnattachAskHTML(qq422016 qtio422016.Writer, rq *http.Request, hyphaName string, isOld bool) {
//line views/modal.qtpl:25
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/modal.qtpl:25
	StreamUnattachAskHTML(qw422016, rq, hyphaName, isOld)
//line views/modal.qtpl:25
	qt422016.ReleaseWriter(qw422016)
//line views/modal.qtpl:25
}

//line views/modal.qtpl:25
func UnattachAskHTML(rq *http.Request, hyphaName string, isOld bool) string {
//line views/modal.qtpl:25
	qb422016 := qt422016.AcquireByteBuffer()
//line views/modal.qtpl:25
	WriteUnattachAskHTML(qb422016, rq, hyphaName, isOld)
//line views/modal.qtpl:25
	qs422016 := string(qb422016.B)
//line views/modal.qtpl:25
	qt422016.ReleaseByteBuffer(qb422016)
//line views/modal.qtpl:25
	return qs422016
//line views/modal.qtpl:25
}

//line views/modal.qtpl:27
func StreamRenameAskHTML(qw422016 *qt422016.Writer, rq *http.Request, hyphaName string, isOld bool) {
//line views/modal.qtpl:27
	qw422016.N().S(`
`)
//line views/modal.qtpl:28
	StreamNavHTML(qw422016, rq, hyphaName, "rename-ask")
//line views/modal.qtpl:28
	qw422016.N().S(`
`)
//line views/modal.qtpl:29
	streammodalBegin(qw422016,
		"rename-confirm",
		hyphaName,
		` method="post" enctype="multipart/form-data"`,
		"Rename "+util.BeautifulName(hyphaName))
//line views/modal.qtpl:33
	qw422016.N().S(`
			<label for="new-name">New name</label>
			<input type="text" value="`)
//line views/modal.qtpl:35
	qw422016.E().S(hyphaName)
//line views/modal.qtpl:35
	qw422016.N().S(`" required autofocus id="new-name" name="new-name"/>

			<input type="checkbox" id="recursive" name="recursive" value="true" checked/>
			<label for="recursive">Rename subhyphae too</label>

			<p>If you rename this hypha, all incoming links and all relative outcoming links will break. You will also lose all history for the new name. Rename carefully.</p>
`)
//line views/modal.qtpl:41
	streammodalEnd(qw422016, hyphaName, false)
//line views/modal.qtpl:41
	qw422016.N().S(`
`)
//line views/modal.qtpl:42
}

//line views/modal.qtpl:42
func WriteRenameAskHTML(qq422016 qtio422016.Writer, rq *http.Request, hyphaName string, isOld bool) {
//line views/modal.qtpl:42
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/modal.qtpl:42
	StreamRenameAskHTML(qw422016, rq, hyphaName, isOld)
//line views/modal.qtpl:42
	qt422016.ReleaseWriter(qw422016)
//line views/modal.qtpl:42
}

//line views/modal.qtpl:42
func RenameAskHTML(rq *http.Request, hyphaName string, isOld bool) string {
//line views/modal.qtpl:42
	qb422016 := qt422016.AcquireByteBuffer()
//line views/modal.qtpl:42
	WriteRenameAskHTML(qb422016, rq, hyphaName, isOld)
//line views/modal.qtpl:42
	qs422016 := string(qb422016.B)
//line views/modal.qtpl:42
	qt422016.ReleaseByteBuffer(qb422016)
//line views/modal.qtpl:42
	return qs422016
//line views/modal.qtpl:42
}

//line views/modal.qtpl:44
func streammodalReallyWant(qw422016 *qt422016.Writer, hyphaName, verb string) {
//line views/modal.qtpl:44
	qw422016.N().S(`
			<p class="modal__confirmation-msg">Do you really want to `)
//line views/modal.qtpl:45
	qw422016.E().S(verb)
//line views/modal.qtpl:45
	qw422016.N().S(` hypha <em>`)
//line views/modal.qtpl:45
	qw422016.E().S(hyphaName)
//line views/modal.qtpl:45
	qw422016.N().S(`</em>?</p>
`)
//line views/modal.qtpl:46
}

//line views/modal.qtpl:46
func writemodalReallyWant(qq422016 qtio422016.Writer, hyphaName, verb string) {
//line views/modal.qtpl:46
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/modal.qtpl:46
	streammodalReallyWant(qw422016, hyphaName, verb)
//line views/modal.qtpl:46
	qt422016.ReleaseWriter(qw422016)
//line views/modal.qtpl:46
}

//line views/modal.qtpl:46
func modalReallyWant(hyphaName, verb string) string {
//line views/modal.qtpl:46
	qb422016 := qt422016.AcquireByteBuffer()
//line views/modal.qtpl:46
	writemodalReallyWant(qb422016, hyphaName, verb)
//line views/modal.qtpl:46
	qs422016 := string(qb422016.B)
//line views/modal.qtpl:46
	qt422016.ReleaseByteBuffer(qb422016)
//line views/modal.qtpl:46
	return qs422016
//line views/modal.qtpl:46
}

//line views/modal.qtpl:48
func streammodalBegin(qw422016 *qt422016.Writer, path, hyphaName, formAttrs, legend string) {
//line views/modal.qtpl:48
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
	<form class="modal" action="/`)
//line views/modal.qtpl:51
	qw422016.E().S(path)
//line views/modal.qtpl:51
	qw422016.N().S(`/`)
//line views/modal.qtpl:51
	qw422016.E().S(hyphaName)
//line views/modal.qtpl:51
	qw422016.N().S(`"`)
//line views/modal.qtpl:51
	qw422016.N().S(formAttrs)
//line views/modal.qtpl:51
	qw422016.N().S(`>
		<fieldset class="modal__fieldset">
			<legend class="modal__title">`)
//line views/modal.qtpl:53
	qw422016.N().S(legend)
//line views/modal.qtpl:53
	qw422016.N().S(`</legend>
`)
//line views/modal.qtpl:54
}

//line views/modal.qtpl:54
func writemodalBegin(qq422016 qtio422016.Writer, path, hyphaName, formAttrs, legend string) {
//line views/modal.qtpl:54
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/modal.qtpl:54
	streammodalBegin(qw422016, path, hyphaName, formAttrs, legend)
//line views/modal.qtpl:54
	qt422016.ReleaseWriter(qw422016)
//line views/modal.qtpl:54
}

//line views/modal.qtpl:54
func modalBegin(path, hyphaName, formAttrs, legend string) string {
//line views/modal.qtpl:54
	qb422016 := qt422016.AcquireByteBuffer()
//line views/modal.qtpl:54
	writemodalBegin(qb422016, path, hyphaName, formAttrs, legend)
//line views/modal.qtpl:54
	qs422016 := string(qb422016.B)
//line views/modal.qtpl:54
	qt422016.ReleaseByteBuffer(qb422016)
//line views/modal.qtpl:54
	return qs422016
//line views/modal.qtpl:54
}

//line views/modal.qtpl:56
func streammodalEnd(qw422016 *qt422016.Writer, hyphaName string, shouldFocusOnConfirm bool) {
//line views/modal.qtpl:56
	qw422016.N().S(`
			<input type="submit" value="Confirm" class="btn" `)
//line views/modal.qtpl:57
	if shouldFocusOnConfirm {
//line views/modal.qtpl:57
		qw422016.N().S(`autofocus`)
//line views/modal.qtpl:57
	}
//line views/modal.qtpl:57
	qw422016.N().S(`>
			<a href="/hypha/`)
//line views/modal.qtpl:58
	qw422016.E().S(hyphaName)
//line views/modal.qtpl:58
	qw422016.N().S(`" class="btn btn_weak">Cancel</a>
		</fieldset>
	</form>
</main>
</div>
`)
//line views/modal.qtpl:63
}

//line views/modal.qtpl:63
func writemodalEnd(qq422016 qtio422016.Writer, hyphaName string, shouldFocusOnConfirm bool) {
//line views/modal.qtpl:63
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/modal.qtpl:63
	streammodalEnd(qw422016, hyphaName, shouldFocusOnConfirm)
//line views/modal.qtpl:63
	qt422016.ReleaseWriter(qw422016)
//line views/modal.qtpl:63
}

//line views/modal.qtpl:63
func modalEnd(hyphaName string, shouldFocusOnConfirm bool) string {
//line views/modal.qtpl:63
	qb422016 := qt422016.AcquireByteBuffer()
//line views/modal.qtpl:63
	writemodalEnd(qb422016, hyphaName, shouldFocusOnConfirm)
//line views/modal.qtpl:63
	qs422016 := string(qb422016.B)
//line views/modal.qtpl:63
	qt422016.ReleaseByteBuffer(qb422016)
//line views/modal.qtpl:63
	return qs422016
//line views/modal.qtpl:63
}
