package main

type fgfOptions interface {
	apply(*fgfInstance)
}

type withFontsDir struct {
	value string
}

func (dir withFontsDir) apply(instance *fgfInstance) {
	instance.fontsDir = dir.value
}
