package accesskey

import (
	"strconv"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

// List returns a list of access keys
func List(c pc.PrismaCloudClient) ([]AccessKey, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var ans []AccessKey
	path := Suffix[:]

	_, err := c.Communicate("GET", path, nil, nil, &ans)
	return ans, err
}

// Get returns an access key by ID
func Get(c pc.PrismaCloudClient, id string) (AccessKey, error) {
	c.Log(pc.LogAction, "(get) %s", singular)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)

	var ans AccessKey
	_, err := c.Communicate("GET", path, nil, nil, &ans)
	return ans, err
}

// Create adds a new access key
func Create(c pc.PrismaCloudClient, accKey AccessKey) (AccessKey, error) {
	c.Log(pc.LogAction, "(create) %s", singular)

	path := Suffix[:]
	var ans AccessKey
	_, err := c.Communicate("POST", path, nil, accKey, &ans)
	return ans, err
}

// Update changes a status of an existing access key
func Update(c pc.PrismaCloudClient, accKey AccessKey) error {
	c.Log(pc.LogAction, "(udpate) %s", singular)

	path := make([]string, 0, len(Suffix)+3)
	path = append(path, Suffix...)
	path = append(path, accKey.Id)
	path = append(path, "status")
	path = append(path, strconv.FormatBool(accKey.Status))

	_, err := c.Communicate("PATCH", path, nil, nil, nil)
	return err
}

// Delete removes an access key by ID
func Delete(c pc.PrismaCloudClient, id string) error {
	c.Log(pc.LogAction, "(delete) %s", singular)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)

	_, err := c.Communicate("DELETE", path, nil, nil, nil)
	return err
}
