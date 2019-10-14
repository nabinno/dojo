package awslambda

import (
	"fmt"
	"os"
)

// DispatchOnLocal ....
func DispatchOnLocal() {
	if out, err := doDispatchOnLocal(); err == nil {
		fmt.Println(string(out))
	} else {
		fmt.Println(err)
	}
}

func doDispatchOnLocal() (out []byte, err error) {
	if err := os.Chdir("aws_lambda"); err == nil {
		if out, err = createDispatchYaml(); err != nil {
			return out, err
		}

		if out, err = outcmd("bash -c exec docker-compose run --rm sam-local local invoke 'HealthFunction'"); err != nil {
			return out, err
		}

		err = os.Remove("docker-compose.yml")
	}
	return
}

// @todo 2019-08-24
func createDispatchYaml() ([]byte, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	return outcmd(fmt.Sprintf(`
cat <<EOF >docker-compose.yml
version: "2"
services:
  sam-local:
	image: akirakoyasu/aws-sam-local
	environment:
	  # mount host dir
	  SAM_DOCKER_VOLUME_BASEDIR: %v
	  AWS_REGION: us-west-2
	volumes:
	  - .:/var/lib/sam-local
	  - /var/run/docker.sock:/var/run/docker.sock
EOF
`, pwd))
}
