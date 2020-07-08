package gpu

import (
	"github.com/NVIDIA/gpu-monitoring-tools/bindings/go/nvml"
	"github.com/hyc3z/Omaticaya/src/global"
	"go.uber.org/zap"
	"log"
	"time"
)

func nvmlDeviceToGpu(d *nvml.Device, id uint) global.GPU {
	t0 := time.Now()
	status, err := d.Status()
	t1t0 := time.Since(t0)
	global.Logger.Info("GetGPUStatus Spent",
		zap.Int64("GetGPUStatus Duration", t1t0.Milliseconds()),
	)
	if err != nil {
		global.Logger.Panic("getDeviceStatus Fail",
			zap.Error(err),
		)
	}
	gpuInfo := global.GPU{
		Device:         d,
		CountID:        id,
		UUID:           d.UUID,
		CudaComputeCap: d.CudaComputeCapability,
		Model:          *d.Model,
		Power:          *status.Power,
		Memory:         *d.Memory,
		MemoryClock:    *d.Clocks.Memory,
		FreeMemory:     *status.Memory.Global.Free,
		CoreClock:      *d.Clocks.Cores,
		Bandwidth:      *d.PCI.Bandwidth,
		Processes:      &status.Processes,
	}
	return gpuInfo
}

func GetGpuInfo() error {
	err := nvml.Init()
	if err != nil {
		global.Logger.Error("Error nvml.Init:",
			zap.Error(err),
		)
		return err
	}
	defer func() {
		err = nvml.Shutdown()
		if err != nil {
			global.Logger.Error("Error nvml.Shutdown:",
				zap.Error(err),
			)
		}
	}()

	count, err := nvml.GetDeviceCount()
	if err != nil {
		global.Logger.Panic("Error getting device count:",
			zap.Error(err))
	}

	driverVersion, err := nvml.GetDriverVersion()
	if err != nil {
		global.Logger.Panic("Error getting driver version:",
			zap.Error(err))
	}

	global.Logger.Info("getGpuData",
		zap.String("Node name", global.ProjectInfo.Node.NodeName),
		zap.Uint("Gpu Count", count),
		zap.String("driver version", driverVersion),
	)
	if count > 0 {
		global.ProjectInfo.Node.HasGpu = true
	}
	for i := uint(0); i < count; i++ {
		device, err := nvml.NewDevice(i)
		if err != nil {
			global.Logger.Panic("Error getting device : ",
				zap.Uint("Device count", i),
				zap.Error(err),
			)
		}
		t0 := time.Now()
		currentGpu := nvmlDeviceToGpu(device, i)
		t1t0 := time.Since(t0)
		global.Logger.Info("nvmlDeviceToGpu Spent",
			zap.Int64("nvmlDeviceToGpu Duration", t1t0.Milliseconds()),
		)

		gpus := global.ProjectInfo.Node.Gpus
		flagAppend := true
		for i := range gpus {
			if gpus[i].UUID == currentGpu.UUID {
				flagAppend = false
				gpus[i] = currentGpu
				break
			}
		}
		if flagAppend {
			global.ProjectInfo.Node.Gpus = append(gpus, currentGpu)
		}
		if err != nil {
			log.Panicln("Template error:", err)
		}
	}
	return err
}
