package ustr

import (
	"strings"
)

// Pats is a slice of `Pat`s.
type Pats []Pat

// Add `append`s all the specified `pats` to `this`.
func (this *Pats) Add(pats ...Pat) {
	*this = append(*this, pats...)
}

// FirstMatch returns the first `Pat` in `this` to `Match(s)`, or `""`.
func (this Pats) FirstMatch(s string) Pat {
	for _, pat := range this {
		if pat.Match(s) {
			return pat
		}
	}
	return ""
}

// NoMatch returns whether not a single `Pat` in `this` does `Match(s)`.
func (this Pats) NoMatch(s string) bool {
	return this.FirstMatch(s) == ""
}

// Pat is a most-simplistic, overly-rudimentary, simplest-of-simpletons string pattern matcher.
// It allows a single asterisk `*` wild-card at its beginning, its end, or both, as described in `Pat.Match`.
// This covers a bafflingly substantial amount of real-world use-cases — if more is needed, Go's `path.Match`, reg-exps etc. will deliver instead.
type Pat string

// AllMatch returns whether all the specified `strs` satisfy `this.Match`.
func (this Pat) AllMatch(strs ...string) bool {
	if this != "" && this != "*" {
		for _, s := range strs {
			if !this.Match(s) {
				return false
			}
		}
	}
	return true
}

// FirstMatch returns the first in `strs` that `this.Match`es, or `""`.
func (this Pat) FirstMatch(strs ...string) string {
	for _, s := range strs {
		if s != "" && this.Match(s) {
			return s
		}
	}
	return ""
}

// Match returns whether `s` matches `this`, which could:
//
// - begin and end with an asterisk `*`  wildcard: "contains" semantics
//
// - only begin with an asterisk `*` wildcard: "endsWith" semantics
//
// - only end with an asterisk `*` wildcard: "beginsWith" semantics
//
// - only consist of an asterisk `*` wildcard: always matches any `s`
//
// - otherwise: matches if `s == this`.
func (this Pat) Match(s string) bool {
	l := len(this)
	if l == 0 || this == "*" {
		return true
	}
	prefix, suffix := this[0] == '*', this[l-1] == '*'
	if prefix && suffix {
		return strings.Contains(s, string(this)[1:l-1])
	} else if prefix {
		return strings.HasSuffix(s, string(this)[1:])
	} else if suffix {
		return strings.HasPrefix(s, string(this)[:l-1])
	}
	return s == string(this)
}
