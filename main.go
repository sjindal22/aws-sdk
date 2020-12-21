package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func createSession(r string, prof string) *session.Session {

	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(r),
		},
		Profile: prof,
	},
	)
	if err != nil {
		fmt.Println(err)
	}
	return sess
}

func createSSMParameters(reg string, profile string, paramsList map[string]string) {

	csess := ssm.New(createSession(reg, profile))
	for name, val := range paramsList {
		putParam, err := csess.PutParameter(&ssm.PutParameterInput{
			Name:  aws.String(name),
			Value: aws.String(val),
			Type:  aws.String("String"),
		})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(putParam.Tier)
	}
}

func getSSMParameters(r string, p string) {

	paramsList := []string{"/testaccount/env/dev", "foo"}
	strPtr := aws.StringSlice(paramsList)
	//csess := ssm.New(createSession(), aws.NewConfig.WithRegion("us-east-1"))
	csess := ssm.New(createSession(r, p))
	paramOP, err := csess.GetParameters(&ssm.GetParametersInput{Names: strPtr})

	if err != nil {
		fmt.Println(err)
	}
	for _, val := range paramOP.Parameters {
		fmt.Println(*val.Value)
	}

}

func main() {

	paramL := map[string]string{
		"username": "username123",
		"password": "password123",
	}
	createSSMParameters("us-east-1", "devops", paramL)
	getSSMParameters("us-east-1", "devops")
}

// Reference https://github.com/ProjectThor/get-ssm-params
