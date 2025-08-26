package runtimeapi

import (
	"context"
	"exec-runtime/pkg/config"
	"go.uber.org/zap"
	k8sCriApi "k8s.io/cri-api/pkg/apis/runtime/v1"
)

type Service struct {
	config *config.Runtime
	logger *zap.Logger
}

func (s *Service) Version(ctx context.Context, request *k8sCriApi.VersionRequest) (*k8sCriApi.VersionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) RunPodSandbox(ctx context.Context, request *k8sCriApi.RunPodSandboxRequest) (*k8sCriApi.RunPodSandboxResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) StopPodSandbox(ctx context.Context, request *k8sCriApi.StopPodSandboxRequest) (*k8sCriApi.StopPodSandboxResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) RemovePodSandbox(ctx context.Context, request *k8sCriApi.RemovePodSandboxRequest) (*k8sCriApi.RemovePodSandboxResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) PodSandboxStatus(ctx context.Context, request *k8sCriApi.PodSandboxStatusRequest) (*k8sCriApi.PodSandboxStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) ListPodSandbox(ctx context.Context, request *k8sCriApi.ListPodSandboxRequest) (*k8sCriApi.ListPodSandboxResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) CreateContainer(ctx context.Context, request *k8sCriApi.CreateContainerRequest) (*k8sCriApi.CreateContainerResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) StartContainer(ctx context.Context, request *k8sCriApi.StartContainerRequest) (*k8sCriApi.StartContainerResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) StopContainer(ctx context.Context, request *k8sCriApi.StopContainerRequest) (*k8sCriApi.StopContainerResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) RemoveContainer(ctx context.Context, request *k8sCriApi.RemoveContainerRequest) (*k8sCriApi.RemoveContainerResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) ListContainers(ctx context.Context, request *k8sCriApi.ListContainersRequest) (*k8sCriApi.ListContainersResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) ContainerStatus(ctx context.Context, request *k8sCriApi.ContainerStatusRequest) (*k8sCriApi.ContainerStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) UpdateContainerResources(ctx context.Context, request *k8sCriApi.UpdateContainerResourcesRequest) (*k8sCriApi.UpdateContainerResourcesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) ReopenContainerLog(ctx context.Context, request *k8sCriApi.ReopenContainerLogRequest) (*k8sCriApi.ReopenContainerLogResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) ExecSync(ctx context.Context, request *k8sCriApi.ExecSyncRequest) (*k8sCriApi.ExecSyncResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Exec(ctx context.Context, request *k8sCriApi.ExecRequest) (*k8sCriApi.ExecResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Attach(ctx context.Context, request *k8sCriApi.AttachRequest) (*k8sCriApi.AttachResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) PortForward(ctx context.Context, request *k8sCriApi.PortForwardRequest) (*k8sCriApi.PortForwardResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) ContainerStats(ctx context.Context, request *k8sCriApi.ContainerStatsRequest) (*k8sCriApi.ContainerStatsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) ListContainerStats(ctx context.Context, request *k8sCriApi.ListContainerStatsRequest) (*k8sCriApi.ListContainerStatsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) PodSandboxStats(ctx context.Context, request *k8sCriApi.PodSandboxStatsRequest) (*k8sCriApi.PodSandboxStatsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) ListPodSandboxStats(ctx context.Context, request *k8sCriApi.ListPodSandboxStatsRequest) (*k8sCriApi.ListPodSandboxStatsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) UpdateRuntimeConfig(ctx context.Context, request *k8sCriApi.UpdateRuntimeConfigRequest) (*k8sCriApi.UpdateRuntimeConfigResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Status(ctx context.Context, request *k8sCriApi.StatusRequest) (*k8sCriApi.StatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) CheckpointContainer(ctx context.Context, request *k8sCriApi.CheckpointContainerRequest) (*k8sCriApi.CheckpointContainerResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetContainerEvents(request *k8sCriApi.GetEventsRequest, server k8sCriApi.RuntimeService_GetContainerEventsServer) error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) ListMetricDescriptors(ctx context.Context, request *k8sCriApi.ListMetricDescriptorsRequest) (*k8sCriApi.ListMetricDescriptorsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) ListPodSandboxMetrics(ctx context.Context, request *k8sCriApi.ListPodSandboxMetricsRequest) (*k8sCriApi.ListPodSandboxMetricsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) RuntimeConfig(ctx context.Context, request *k8sCriApi.RuntimeConfigRequest) (*k8sCriApi.RuntimeConfigResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) UpdatePodSandboxResources(ctx context.Context, request *k8sCriApi.UpdatePodSandboxResourcesRequest) (*k8sCriApi.UpdatePodSandboxResourcesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewService(config *config.Runtime, logger *zap.Logger) *Service {
	return &Service{
		config: config,
		logger: logger,
	}
}
