package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func main() {

	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String("us-east-1")}))
	ssmSess := ssm.New(sess)
	op, err := ssmSess.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String("shivika"),
		WithDecryption: aws.Bool(false),
	})
	if err != nil {
		fmt.Println(err)
	}
	value := *op.Parameter.Value
	fmt.Println(value)
}
