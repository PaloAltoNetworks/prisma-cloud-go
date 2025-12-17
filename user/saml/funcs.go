package saml

import (
	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

// GetBypassUsers retrieves the list of users that are allowed to bypass SAML authentication.
func GetBypassUsers(c pc.PrismaCloudClient) ([]string, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var ans BypassUsersResponse
	if _, err := c.Communicate("GET", Suffix, nil, nil, &ans); err != nil {
		return nil, err
	}

	return ans, nil
}

// UpdateBypassUsers updates the list of users that are allowed to bypass SAML authentication.
func UpdateBypassUsers(c pc.PrismaCloudClient, usernames []string) error {
	c.Log(pc.LogAction, "(update) %s", plural)

	// Convert the input array to the request type
	req := UpdateBypassUsersRequest(usernames)

	_, err := c.Communicate("PUT", Suffix, nil, req, nil)
	return err
}
