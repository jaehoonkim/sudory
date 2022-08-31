package service

import (
	"time"

	"github.com/NexClipper/sudory/pkg/client/log"
	servicev3 "github.com/NexClipper/sudory/pkg/server/model/service/v3"
)

type ServiceExecType int32

const (
	ServiceExecTypeImmediate = iota
	// ServiceExecTypePeriodic
)

type ServiceStatus int32

const (
	ServiceStatusPreparing ServiceStatus = iota + 1
	ServiceStatusStart
	ServiceStatusProcessing
	ServiceStatusSuccess
	ServiceStatusFailed
)

func (s ServiceStatus) String() string {
	switch s {
	case ServiceStatusPreparing:
		return "ServiceStatusPreparing"
	case ServiceStatusStart:
		return "ServiceStatusStart"
	case ServiceStatusProcessing:
		return "ServiceStatusProcessing"
	case ServiceStatusSuccess:
		return "ServiceStatusSuccess"
	case ServiceStatusFailed:
		return "ServiceStatusFailed"
	default:
		return "ServiceStatusUnknown"
	}
}

type Service struct {
	Id          string
	Name        string
	ClusterId   string
	Priority    int
	CreatedTime time.Time
	StartTime   time.Time
	UpdateTime  time.Time
	EndTime     time.Time
	Status      ServiceStatus
	Steps       []Step
	Result      Result
}

type StepStatus int32

const (
	StepStatusPreparing = iota + 1
	StepStatusProcessing
	StepStatusSuccess
	StepStatusFail
)

type StepCommand struct {
	Method string
	Args   map[string]interface{}
}

type Result struct {
	Body string
	Err  error
}

type Step struct {
	Id           string
	ParentId     string
	Command      *StepCommand
	StartTime    time.Time
	EndTime      time.Time
	Status       StepStatus
	ResultFilter string
	Result       Result
}

type UpdateServiceStep struct {
	Uuid      string
	StepCount int
	Sequence  int
	Status    StepStatus
	Result    string
	Started   time.Time
	Ended     time.Time
}

func ConvertServiceListServerToClient(server []servicev3.HttpRsp_ClientServicePolling) map[string]*Service {
	client := make(map[string]*Service)
	for _, v := range server {
		serv := &Service{
			Id:          v.Uuid,
			Name:        v.Name,
			ClusterId:   v.ClusterUuid,
			Priority:    int(v.Priority),
			CreatedTime: v.Created,
		}

		if len(v.Steps) <= 0 {
			log.Warnf("service steps is empty: service_uuid: %s\n", v.Uuid)
			continue
		}

		for _, s := range v.Steps {
			serv.Steps = append(serv.Steps, Step{
				Id:           s.Uuid,
				ParentId:     serv.Id,
				Command:      &StepCommand{Method: s.Method, Args: s.Args},
				ResultFilter: s.ResultFilter.String,
			})
		}
		client[v.Uuid] = serv
	}

	return client
}

func ConvertServiceStepUpdateClientToServer(client UpdateServiceStep) *servicev3.HttpReq_ClientServiceUpdate {
	server := &servicev3.HttpReq_ClientServiceUpdate{
		Uuid:     client.Uuid,
		Sequence: client.Sequence,
		// Status:client.Status,
		Result:  client.Result,
		Started: client.Started,
		Ended:   client.Ended,
	}

	switch client.Status {
	case StepStatusPreparing, StepStatusProcessing:
		server.Status = servicev3.StepStatusProcessing
	case StepStatusSuccess:
		server.Status = servicev3.StepStatusSuccess
	case StepStatusFail:
		server.Status = servicev3.StepStatusFail
	}

	return server
}
