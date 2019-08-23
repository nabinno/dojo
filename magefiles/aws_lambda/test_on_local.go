package awslambda

import (
	"fmt"
	"os"
)

// TestOnLocal ....
func TestOnLocal() {
	if out, err := doTestOnLocal(); err == nil {
		fmt.Println(string(out))
	} else {
		fmt.Println(err)
	}
}

func doTestOnLocal() (out []byte, err error) {
	if err := os.Chdir("aws_lambda"); err == nil {
		if out, err = createTestYaml(); err != nil {
			return out, err
		}

		if out, err = outcmd("bash -c exec docker-compose run --rm lambda-python -m unittest discover"); err != nil {
			return out, err
		}

		err = os.Remove("docker-compose.yml")
	}
	return
}

// @todo 2019-08-24
func createTestYaml() (out []byte, err error) {
	return outcmd(`
cat <<EOF >docker-compose.yml
version: "2"
services:
  lambda-python:
	image: lambci/lambda:python3.6
	environment:
	  PYTHONPATH: /var/task/src/spec:/var/task/src/app:/var/task/src/app/site-packages:/var/runtime
	volumes:
	  - .:/var/task
	working_dir: /var/task/src/spec
	user: root
	entrypoint: []
EOF
`)
}
