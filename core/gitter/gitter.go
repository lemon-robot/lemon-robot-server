package gitter

var ins Standard

func ConfigIns(gitType string, gitConfig map[string]string) {
	ins = SupportedTypes()[gitType]
	ins.Init(gitConfig)
}

func Ins() Standard {
	return ins
}
