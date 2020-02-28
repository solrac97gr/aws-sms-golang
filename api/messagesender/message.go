package message

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func SendSMS_ABTEST(public []string, AWS_ACCESS_KEY_ID string, AWS_SECRET_ACCESS_KEY string) {
	fmt.Println("creating session")
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, ""),
		Region:      aws.String("us-east-1"),
	}))
	fmt.Println("session created")

	svc := sns.New(sess)
	fmt.Println("service created")
	for i := 0; i < len(public); i++ {
		if i%2 == 0 || i == 0 {
			params := &sns.PublishInput{
				Message:     aws.String("mensaje 1"),
				PhoneNumber: aws.String("+51" + public[i]),
				MessageAttributes: map[string]*sns.MessageAttributeValue{
					"DefaultSMSType": {
						DataType:    aws.String("String"),
						StringValue: aws.String("Transactional")},
				},
			}
			_, err := svc.Publish(params)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

		} else {

			params := &sns.PublishInput{
				Message:     aws.String("mensaje 2"),
				PhoneNumber: aws.String("+51" + public[i]),
				MessageAttributes: map[string]*sns.MessageAttributeValue{
					"DefaultSMSType": {
						DataType:    aws.String("String"),
						StringValue: aws.String("Transactional")},
				},
			}
			_, err := svc.Publish(params)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

		}
	}
	fmt.Println("Job Finish")
}

func SendSMS(public []string, AWS_ACCESS_KEY_ID string, AWS_SECRET_ACCESS_KEY string) {
	// Public is a array of numbers, this numbers comes from other module for create A/B Test Publics
	fmt.Println("creating session")
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, ""),
		Region:      aws.String("us-east-1"),
	}))
	fmt.Println("session created")

	svc := sns.New(sess)
	fmt.Println("service created")
	for i := 0; i < len(public); i++ {
		params := &sns.PublishInput{
			Message:     aws.String("Tu increíble Mensajito"),
			PhoneNumber: aws.String("+51" + public[i]),
			MessageAttributes: map[string]*sns.MessageAttributeValue{
				"DefaultSMSType": {
					DataType:    aws.String("String"),
					StringValue: aws.String("Transactional")},
			},
		}
		_, err := svc.Publish(params)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

	}
	fmt.Println("Job Finish")
}
