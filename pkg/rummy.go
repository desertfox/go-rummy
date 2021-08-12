package rummy

type Installer interface {
	Install()
}

func Go(plugins []Installer) {
	for _, plugin := range plugins {
		plugin.Install()
	}
}
