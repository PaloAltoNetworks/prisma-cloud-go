package azureTemplate

import (
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	//	"net/http"
	//	"io/ioutil"
	//	"fmt"
)

func GetAzureTemplate(c pc.PrismaCloudClient, req AzureTemplateReq) (AzureTemplate, error) {
	c.Log(pc.LogAction, "(post) performing azure template")

	var resp AzureTemplate

	_, err := c.Communicate("POST", Suffix, nil, req, &resp)
	return resp, err
}

//func GetAzureTemplate(c pc.PrismaCloudClient, req AzureTemplateReq) (AzureTemplate, error) {
//	c.Log(pc.LogAction, "(post) performing azure template")
//
//	var resp AzureTemplate
//
//	_, err := c.Communicate("POST", Suffix, nil, req, &resp)
//	resp1, err := http.Post("https://api-stage.prismacloud.io/cas/v1/azure_template","AzureTemplate",resp)
//	fmt.Println("ssssssssss",resp1)
//	check(err)
//	defer resp1.Body.Close()
//	body, err := ioutil.ReadAll(resp1.Body)
//	check(err)
//	err = ioutil.WriteFile("./data.txt", body, 0666)
//	check(err)
//	return resp, err
//}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
