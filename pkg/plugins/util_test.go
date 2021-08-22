package plugins

import (
	"path/filepath"
	"testing"
)

func TestBuildDestWithFile(t *testing.T) {
	plugin := &PluginData{DestFilesDir: "testDir"}
	file := "testFile"

	want := "testDir/testFile"
	got := plugin.BuildDestWithFile(file)

	if got != want {
		t.Errorf("Did not build file path correclty got:%v want:%v", got, want)
	}

}

func TestAddConfigToCreate(t *testing.T) {
	source := "text"
	to := "file.txt"
	overwrite := false
	destDir := "testDir"

	want := ConfigToCreate{&source, filepath.Join(destDir, to), overwrite}

	got := &PluginData{DestFilesDir: destDir}
	got.AddConfigToCreate(&source, to, overwrite)

	if got.Configs[0] != want {
		t.Errorf("Did not add config to create correclty got:%v want:%v", got, want)
	}

}
