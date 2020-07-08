package scheduler

import (
	"context"
	_ "fmt"
	v1 "k8s.io/api/core/v1"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
	"k8s.io/kubernetes/pkg/scheduler/nodeinfo"
)

var (
	_ framework.QueueSortPlugin = &CustomedScheduler{}
	_ framework.FilterPlugin    = &CustomedScheduler{}
	_ framework.ScorePlugin     = &CustomedScheduler{}
	_ framework.ScoreExtensions = &CustomedScheduler{}
)

type Args struct {
	KubeConfig string `json:"kubeconfig,omitempty"`
	Master     string `json:"master,omitempty"`
}

type CustomedScheduler struct {
	args   *Args
	handle framework.FrameworkHandle
}

func (y *CustomedScheduler) NormalizeScore(ctx context.Context, state *framework.CycleState, p *v1.Pod, scores framework.NodeScoreList) *framework.Status {
	panic("implement me")
}

func (y *CustomedScheduler) Score(ctx context.Context, state *framework.CycleState, p *v1.Pod, nodeName string) (int64, *framework.Status) {
	panic("implement me")
}

func (y *CustomedScheduler) ScoreExtensions() framework.ScoreExtensions {
	panic("implement me")
}

func (y *CustomedScheduler) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeInfo *nodeinfo.NodeInfo) *framework.Status {
	panic("implement me")
}

func (y *CustomedScheduler) Less(info *framework.PodInfo, info2 *framework.PodInfo) bool {
	panic("implement me")
}

func (y *CustomedScheduler) Name() string {
	return "CustomedScheduler"
}
