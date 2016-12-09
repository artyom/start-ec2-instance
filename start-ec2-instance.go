// Command start-ec2-instance starts AWS EC2 instance with given instance id
package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ec2-instance-id\n", os.Args[0])
		os.Exit(1)
	}
	if err := startInstance(os.Args[1]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func startInstance(name string) error {
	sess, err := session.NewSession()
	if err != nil {
		return err
	}
	svc := ec2.New(sess)
	params := &ec2.StartInstancesInput{
		InstanceIds: []*string{
			aws.String(name),
		},
	}
	_, err = svc.StartInstances(params)
	return err
}
