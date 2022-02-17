package client

import (
	"encoding/json"
	"testing"

	"opentext.com/axcelerate/adp/task"
)

func TestListEntitiesResponse(t *testing.T) {
	var taskResp task.Response

	listEntitiesResponse := `{"executionId":"016a8089-41e9-4446-9784-29b0133d0064","taskType":"List Entities","loggingEnabled":"true","executionMetaData":{"adp_entities_output_file_name":"S:\\Projects\\adp.adp\\adpRootDir\\output.json","adp_entities_json_output":"[{\"id\":\"singleMindServer.G01610\",\"displayName\":\"G01610\",\"hostId\":\"dev1-eca0.axcelerate.local\"},{\"id\":\"singleMindServer.G01610_003\",\"displayName\":\"G01610_003\",\"hostId\":\"dev1-eca0.axcelerate.local\"},{\"id\":\"singleMindServer.G01610_002\",\"displayName\":\"G01610_002\",\"hostId\":\"dev1-eca1.axcelerate.local\"}]"},"progressMax":1,"executionStatus":"success","executionRootDir":"S:\\Projects\\adp.adp\\adpRootDir","contextId":"f40e08ec-a899-4f7e-836d-51c7a942b56f","executionPersistent":"true","progressCurrent":1,"progressPercentage":1.0,"taskDisplayName":"List Applications"}`

	taskResp.ExecutionMetaData = &task.ListEntitiesExecutionMetaData{}

	json.Unmarshal([]byte(listEntitiesResponse), &taskResp)
	out := taskResp.ExecutionMetaData.Output()

	expected := `[{"id":"singleMindServer.G01610","displayName":"G01610","hostId":"dev1-eca0.axcelerate.local"},{"id":"singleMindServer.G01610_003","displayName":"G01610_003","hostId":"dev1-eca0.axcelerate.local"},{"id":"singleMindServer.G01610_002","displayName":"G01610_002","hostId":"dev1-eca1.axcelerate.local"}]`
	if out != expected {
		t.Errorf("\nexpected: %s\n vs\n output: %s\n", expected, out)
	}
}
