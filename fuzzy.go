package main

import "unicode/utf8"

var noop = func(r rune) rune { return r }

func match(source, target string, fn func(rune) rune) bool {
	lenDiff := len(target) - len(source)

	if lenDiff < 0 {
		return false
	}

	if lenDiff == 0 && source == target {
		return true
	}

outer:
	for _, r1 := range source {
		for i, r2 := range target {
			if fn(r1) == fn(r2) {
				target = target[i+utf8.RuneLen(r2):]
				continue outer
			}
		}
		return false
	}

	return true
}

func matchNoop(source, target string) bool {
	return match(source, target, noop)
}
