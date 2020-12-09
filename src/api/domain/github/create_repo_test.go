package github

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestCreateRepoRequestAsJSON(t *testing.T) {
	request := CreateRepoRequest{
		Name: "Hello-World",
		Description: "This your first repo!",
		Homepage: "https://github.com",
		Private: true,
		HasIssues: true,
		HasProjects: true,
		HasWiki: true,
	}

	// Marshal takes an input interface and attempts to create a valid json string
	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	fmt.Println(string(bytes))
	assert.EqualValues(t, `{"name":"Hello-World","description":"This your first repo!","homepage":"https://github.com","private":true,"has_issues":true,"has_projects":true,"has_wiki":true}`, string(bytes))

	var target CreateRepoRequest
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)
	assert.EqualValues(t, target.Name, request.Name)
	assert.EqualValues(t, target.HasIssues, request.HasIssues)

}