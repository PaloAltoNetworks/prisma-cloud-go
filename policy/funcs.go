package policy

import (
	"fmt"
	"net/url"
	"strings"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"gopkg.in/yaml.v2"
)

// List returns a list of available policies, both system default and custom.
func List(c pc.PrismaCloudClient, query map[string]string) ([]Policy, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	qv := url.Values{}
	for k, v := range query {
		qv.Set(k, v)
	}

	var ans []Policy
	_, err := c.Communicate("GET", []string{"v2", "policy"}, qv, nil, &ans)

	return ans, err
}

// Identify returns the ID for the given policy name.
func Identify(c pc.PrismaCloudClient, name string) (string, error) {
	c.Log(pc.LogAction, "(get) id for %s name:%s", singular, name)

	ans, err := List(c, map[string]string{"policy.name": name})
	if err != nil {
		return "", err
	}

	switch len(ans) {
	case 0:
		return "", pc.ObjectNotFoundError
	case 1:
		return ans[0].PolicyId, nil
	}

	return "", fmt.Errorf("Got %d results back not 1", len(ans))
}

// Get returns the policy that has the specified ID.
func Get(c pc.PrismaCloudClient, id string) (Policy, error) {
	c.Log(pc.LogAction, "(get) %s id:%s", singular, id)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)

	var ans Policy
	_, err := c.Communicate("GET", path, nil, nil, &ans)

	return ans, err
}

// Create adds a new policy.
func Create(c pc.PrismaCloudClient, policy Policy) error {
	return createUpdate(false, c, policy)
}

// Update modifies the existing policy.
func Update(c pc.PrismaCloudClient, policy Policy) error {
	return createUpdate(true, c, policy)
}

// Delete removes a policy using its ID.
func Delete(c pc.PrismaCloudClient, id string) error {
	c.Log(pc.LogAction, "(delete) %s id:%s", singular, id)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)
	_, err := c.Communicate("DELETE", path, nil, nil, nil)
	return err
}

// Validate the metadata.code string for a code security build policy.
func ValidateBuildPolicy(c pc.PrismaCloudClient, policy Policy) error {
	var code map[interface{}]interface{}
	var codeStr string
	c.Log(pc.LogAction, "(validate) %s", singular)
	c.Log(pc.LogAction, "(validate) policy.Rule.Children %s", policy.Rule.Children)

	if len(policy.Rule.Children) > 0 {
		codeStr = policy.Rule.Children[0].Metadata.Code
		if codeStr == "" {
			return nil
		}
	} else {
		return nil
	}

	if err := yaml.Unmarshal([]byte(codeStr), &code); err != nil {
		return err
	}

	codeMapS := convertMapI2MapS(code)
	path := make([]string, 0, len(Suffix)+1)
	path = append(path, BridgecrewSuffix...)
	path = append(path, "definition/none")

	c.Log(pc.LogAction, "(validate) printing path:\n%s", path)
	_, err := c.Communicate("POST", path, nil, codeMapS, nil)
	return err
}

func createUpdate(exists bool, c pc.PrismaCloudClient, policy Policy) error {
	var (
		logMsg strings.Builder
		method string
	)

	logMsg.Grow(30)
	logMsg.WriteString("(")
	if exists {
		logMsg.WriteString("update")
		method = "PUT"
	} else {
		logMsg.WriteString("create")
		method = "POST"
	}
	logMsg.WriteString(") ")

	logMsg.WriteString(singular)
	if exists {
		fmt.Fprintf(&logMsg, ":%s", policy.PolicyId)
	}

	c.Log(pc.LogAction, logMsg.String())

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	if exists {
		path = append(path, policy.PolicyId)
	}

	_, err := c.Communicate(method, path, nil, policy, nil)
	return err
}

// ConvertMapI2MapS walks the given dynamic build policy object recursively, and
// converts maps with interface{} key type to maps with string key type.
func convertMapI2MapS(v interface{}) interface{} {
	switch x := v.(type) {
	case map[interface{}]interface{}:
		m := map[string]interface{}{}
		for k, v2 := range x {
			switch k2 := k.(type) {
			case string: // Fast check if it's already a string
				m[k2] = convertMapI2MapS(v2)
			default:
				m[fmt.Sprint(k)] = convertMapI2MapS(v2)
			}
		}
		v = m

	case []interface{}:
		for i, v2 := range x {
			x[i] = convertMapI2MapS(v2)
		}

	case map[string]interface{}:
		for k, v2 := range x {
			x[k] = convertMapI2MapS(v2)
		}
	}

	return v
}
