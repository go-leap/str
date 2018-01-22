package ustr

import (
	"strings"
)

// Pats is a slice of `Pat`s.
type Pats []Pat

// Add `append`s all the specified `pats` to `me`.
func (me *Pats) Add(pats ...Pat) {
	*me = append(*me, pats...)
}

// FirstMatch returns the first `Pat` in `me` to `Match(s)`, or `""`.
func (me Pats) FirstMatch(s string) Pat {
	for _, pat := range me {
		if pat.Match(s) {
			return pat
		}
	}
	return ""
}

// NoMatch returns whether not a single `Pat` in `me` does `Match(s)`.
func (me Pats) NoMatch(s string) bool {
	return me.FirstMatch(s) == ""
}

// Pat is a most-simplistic, overly-rudimentary, simplest-of-simpletons string pattern matcher.
// It allows a single asterisk `*` wild-card at its beginning, its end, or both, as described in `Pat.Match`.
// This covers a bafflingly substantial amount of real-world use-cases — if more is needed, Go's `path.Match`, reg-exps etc. will deliver instead.
type Pat string

// AllMatch returns whether all the specified `strs` satisfy `me.Match`.
func (me Pat) AllMatch(strs ...string) bool {
	if me != "" && me != "*" {
		for _, s := range strs {
			if !me.Match(s) {
				return false
			}
		}
	}
	return true
}

// FirstMatch returns the first in `strs` that `me.Match`es, or `""`.
func (me Pat) FirstMatch(strs ...string) string {
	for _, s := range strs {
		if s != "" && me.Match(s) {
			return s
		}
	}
	return ""
}

// Matches returns whether `s` matches `me`, which could:
//
// - begin and end with an asterisk `*`  wildcard: "contains" semantics
//
// - only begin with an asterisk `*` wildcard: "endsWith" semantics
//
// - only end with an asterisk `*` wildcard: "beginsWith" semantics
//
// - only consist of an asterisk `*` wildcard: always matches any `s`
//
// - otherwise: matches if `s == me`.
func (me Pat) Match(s string) bool {
	l := len(me)
	if l == 0 || me == "*" {
		return true
	}
	prefix, suffix := me[0] == '*', me[l-1] == '*'
	if prefix && suffix {
		return strings.Contains(s, string(me)[1:l-1])
	} else if prefix {
		return strings.HasSuffix(s, string(me)[1:])
	} else if suffix {
		return strings.HasPrefix(s, string(me)[:l-1])
	}
	return s == string(me)
}
