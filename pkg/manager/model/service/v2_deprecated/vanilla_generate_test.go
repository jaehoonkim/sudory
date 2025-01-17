package v2_test

import (
	"os"
	"testing"

	"github.com/NexClipper/sudory/pkg/manager/database/vanilla/ice_cream_maker"
	v2 "github.com/NexClipper/sudory/pkg/manager/model/service/v2_deprecated"
)

var objs = []interface{}{
	v2.Service_essential{},
	v2.Service{},
	v2.ServiceStatus_essential{},
	v2.ServiceStatus{},
	v2.ServiceResults_essential{},
	v2.ServiceResult{},
	v2.Service_tangled{},
	v2.Service_status{},

	v2.ServiceStep_essential{},
	v2.ServiceStep{},
	v2.ServiceStepStatus_essential{},
	v2.ServiceStepStatus{},
	v2.ServiceStep_tangled{},
}

func TestNoXormColumns(t *testing.T) {

	s, err := ice_cream_maker.GenerateParts(objs, ice_cream_maker.Ingredients)
	if err != nil {
		t.Fatal(err)
	}

	println(s)

	if true {
		filename := "vanilla_generated.go"
		fd, err := os.Create(filename)
		if err != nil {
			t.Fatal(err)
		}

		if _, err = fd.WriteString(s); err != nil {
			t.Fatal(err)
		}
	}
}
