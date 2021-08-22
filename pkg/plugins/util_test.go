package plugins

import (
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
