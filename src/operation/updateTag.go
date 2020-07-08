package operation

import (
	"context"
	"github.com/NVIDIA/gpu-monitoring-tools/bindings/go/nvml"
	"github.com/hyc3z/Omaticaya/src/global"
	"go.uber.org/zap"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"reflect"
	"strconv"
	"strings"
)

func InfoToMap() map[string]string {
	m := make(map[string]string)
	for gpuId := range global.ProjectInfo.Node.Gpus {
		elem := reflect.ValueOf(global.ProjectInfo.Node.Gpus[gpuId])
		relType := elem.Type()
		for i := 0; i < relType.NumField(); i++ {
			if elem.Field(i).Type() == reflect.TypeOf("") {
				m[global.ProjectInfo.ProjectName+"-Gpu"+strconv.Itoa(gpuId)+"-"+relType.Field(i).Name] = strings.Replace(elem.Field(i).String(), " ", "-", -1)
			} else if elem.Field(i).Type() == reflect.TypeOf(true) {
				m[global.ProjectInfo.ProjectName+"-Gpu"+strconv.Itoa(gpuId)+"-"+relType.Field(i).Name] = strconv.FormatBool(elem.Field(i).Bool())
			} else if elem.Field(i).Type() == reflect.TypeOf(uint(0)) || elem.Field(i).Type() == reflect.TypeOf(uint64(0)) {
				m[global.ProjectInfo.ProjectName+"-Gpu"+strconv.Itoa(gpuId)+"-"+relType.Field(i).Name] = strconv.Itoa(int(elem.Field(i).Uint()))
			} else if elem.Field(i).Type() == reflect.TypeOf(nvml.CudaComputeCapabilityInfo{}) {
				versionString := ""
				cudaElem := reflect.ValueOf(nvml.CudaComputeCapabilityInfo{}).Type()
				for t := 0; t < cudaElem.NumField(); t++ {
					if t > 0 {
						versionString += "."
					}
					val := elem.Field(i).Field(t).Elem()
					versionString += strconv.Itoa(int(val.Int()))
				}
				m[global.ProjectInfo.ProjectName+"-Gpu"+strconv.Itoa(gpuId)+"-"+relType.Field(i).Name] = versionString
			} else if elem.Field(i).Type() == reflect.TypeOf([]nvml.ProcessInfo{}) {
				//prefixString := global.ProjectInfo.ProjectName + "-Gpu" + strconv.Itoa(gpuId) + "-PID-"
				arr := elem.Field(i)
				for t := 0; t < arr.Len(); t++ {
					//curProc := arr.Index(t)
					//procMap := make(map[string]string)
					//refType := reflect.TypeOf(nvml.ProcessInfo{})
					//for s := 0; s < refType.NumField(); s++ {
					//	typeStr := refType.Field(s).Type
					//	switch typeStr {
					//	case reflect.TypeOf(""):
					//		procMap[refType.Field(s).Name] = curProc.Field(s).String()
					//	case reflect.TypeOf(uint(0)):
					//		procMap[refType.Field(s).Name] = strconv.Itoa(int(curProc.Field(s).Uint()))
					//	case reflect.TypeOf(uint64(0)):
					//		procMap[refType.Field(s).Name] = strconv.Itoa(int(curProc.Field(s).Uint()))
					//	case reflect.TypeOf(nvml.ProcessType(0)):
					//		procMap[refType.Field(s).Name] = strconv.Itoa(int(curProc.Field(s).Uint()))
					//	}
					//}
					//byteval, _ := json.Marshal(procMap)
					//m[prefixString+procMap["PID"]] = string(byteval)
				}

			}
		}
	}
	return m
}

func UpdateTagForNode() {
	client, err := kubernetes.NewForConfig(global.ProjectInfo.Node.Config)
	if err != nil {
		global.Logger.Panic("Get Config Error UpdateTag",
			zap.Error(err),
		)
	}
	node, err := client.CoreV1().Nodes().Get(context.TODO(), global.ProjectInfo.Node.NodeName, v1.GetOptions{})
	if err != nil {
		global.Logger.Error(
			"updateTagForNodeError",
			zap.Error(err),
		)
	} else {
		nodeLabel := node.GetLabels()
		newLabel := InfoToMap()
		for k, v := range newLabel {
			nodeLabel[k] = v
		}
		node.SetLabels(nodeLabel)
		if _, err := client.CoreV1().Nodes().Update(context.TODO(), node, v1.UpdateOptions{}); err != nil {
			global.Logger.Error("Update Tag failed.",
				zap.Error(err),
			)
		}
	}
}

func CleanTag() {
	client, err := kubernetes.NewForConfig(global.ProjectInfo.Node.Config)
	if err != nil {
		global.Logger.Panic("Get Config Error CleanTag",
			zap.Error(err),
		)
	}
	node, err := client.CoreV1().Nodes().Get(context.TODO(), global.ProjectInfo.Node.NodeName, v1.GetOptions{})
	if err != nil {
		global.Logger.Error(
			"cleanTagError",
			zap.Error(err),
		)
	} else {
		nodeLabel := node.GetLabels()
		for k := range nodeLabel {
			if strings.Contains(k, global.ProjectInfo.ProjectName+"-") {
				delete(nodeLabel, k)
			}
		}
		node.SetLabels(nodeLabel)
		if _, err := client.CoreV1().Nodes().Update(context.TODO(), node, v1.UpdateOptions{}); err != nil {
			global.Logger.Error("Clean Tag failed.",
				zap.Error(err),
			)
		}
	}
}
