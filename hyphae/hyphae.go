// The `hyphae` package is for the Hypha type, hypha storage and stuff like that. It shall not depend on mycorrhiza modules other than util.
package hyphae

import (
	"log"
	"regexp"
	"sync"
)

// HyphaPattern is a pattern which all hyphae must match.
var HyphaPattern = regexp.MustCompile(`[^?!:#@><*|"\'&%{}]+`)

type Hypha struct {
	sync.RWMutex

	Name       string
	Exists     bool
	TextPath   string
	BinaryPath string
	OutLinks   []*Hypha
	BackLinks  []*Hypha
}

var byNames = make(map[string]*Hypha)
var byNamesMutex = sync.Mutex{}

// EmptyHypha returns an empty hypha struct with given name.
func EmptyHypha(hyphaName string) *Hypha {
	return &Hypha{
		Name:       hyphaName,
		Exists:     false,
		TextPath:   "",
		BinaryPath: "",
		OutLinks:   make([]*Hypha, 0),
		BackLinks:  make([]*Hypha, 0),
	}
}

// ByName returns a hypha by name. If h.Exists, the returned hypha pointer is known to be part of the hypha index (byNames map).
func ByName(hyphaName string) (h *Hypha) {
	h, exists := byNames[hyphaName]
	if exists {
		return h
	}
	return EmptyHypha(hyphaName)
}

// Insert inserts the hypha into the storage. It overwrites the previous record, if there was any, and returns false. If the was no previous record, return true.
func (h *Hypha) Insert() (justCreated bool) {
	hp := ByName(h.Name)

	byNamesMutex.Lock()
	defer byNamesMutex.Unlock()
	if hp.Exists {
		hp = h
	} else {
		h.Exists = true
		byNames[h.Name] = h
		IncrementCount()
	}

	return !hp.Exists
}

func (h *Hypha) InsertIfNew() (justCreated bool) {
	if !h.Exists {
		return h.Insert()
	}
	return false
}

func (h *Hypha) InsertIfNewKeepExistence() {
	hp := ByName(h.Name)

	byNamesMutex.Lock()
	defer byNamesMutex.Unlock()
	if hp.Exists {
		hp = h
	} else {
		byNames[h.Name] = h
	}
}

func (h *Hypha) Delete() {
	byNamesMutex.Lock()
	h.Lock()
	delete(byNames, h.Name)
	DecrementCount()
	byNamesMutex.Unlock()
	h.Unlock()
}

func (h *Hypha) RenameTo(newName string) {
	byNamesMutex.Lock()
	h.Lock()
	delete(byNames, h.Name)
	h.Name = newName
	byNames[h.Name] = h
	byNamesMutex.Unlock()
	h.Unlock()
}

// MergeIn merges in content file paths from a different hypha object. Prints warnings sometimes.
func (h *Hypha) MergeIn(oh *Hypha) {
	if h.TextPath == "" && oh.TextPath != "" {
		h.TextPath = oh.TextPath
	}
	if oh.BinaryPath != "" {
		if h.BinaryPath != "" {
			log.Println("There is a file collision for binary part of a hypha:", h.BinaryPath, "and", oh.BinaryPath, "-- going on with the latter")
		}
		h.BinaryPath = oh.BinaryPath
	}
}

// Link related stuff:

func (h *Hypha) AddOutLink(oh *Hypha) (added bool) {
	h.Lock()
	defer h.Unlock()

	for _, outlink := range h.OutLinks {
		if outlink == oh {
			return false
		}
	}
	h.OutLinks = append(h.OutLinks, oh)
	return true
}

func (h *Hypha) AddBackLink(bh *Hypha) (added bool) {
	h.Lock()
	defer h.Unlock()

	for _, backlink := range h.BackLinks {
		if backlink == h {
			return false
		}
	}
	h.BackLinks = append(h.BackLinks, bh)
	return true
}
