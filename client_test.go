package microk8s

import (
	"fmt"
	// "github.com/eneoti/microk8s"
	"testing"
)

type TestStruct struct {
	Name string `json:"name"`
}

func TestMicroCli(t *testing.T) {
	var result TestStruct
	err := CallMicro("dsdsd", "dsdsd", "dsdsds", &result)
	fmt.Println(err)
}
