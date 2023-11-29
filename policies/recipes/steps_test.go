package recipes

import (
	"fmt"
	"testing"

	agentendpointpb "google.golang.org/genproto/googleapis/cloud/osconfig/agentendpoint/v1beta"
)

func TestStepCopyFile(t *testing.T) {
	step := &agentendpointpb.SoftwareRecipe_Step_CopyFile{ArtifactId: "copy-test", Destination: "c:\\osconfig-copy-test"}
	var artifacts map[string]string
	var runEnvs []string
	var stepDir string
	err := stepCopyFile(step, artifacts, runEnvs, stepDir)
	if err != nil {
		fmt.Println(err)
		t.Errorf("q")
	}
}
