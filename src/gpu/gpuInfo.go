package gpu

import (
	"github.com/NVIDIA/gpu-monitoring-tools/bindings/go/nvml"
	"github.com/hyc3z/Omaticaya/src/global"
	"go.uber.org/zap"
	"log"
)

func nvmlDeviceToGpu(d *nvml.Device, id uint) global.GPU {
	status, err := d.Status()
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
		OtherValues:    global.CustomedValues{},
	}
	return gpuInfo
}

func GetGpuInfo() error {
	err := nvml.Init()
	if err != nil {
		global.Logger.Error("Error nvml.Init:",
			zap.Error(err),
		)
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

		currentGpu := nvmlDeviceToGpu(device, i)
		gpus := global.ProjectInfo.Node.Gpus
		global.ProjectInfo.Node.Gpus = append(gpus, currentGpu)
		if err != nil {
			log.Panicln("Template error:", err)
		}
	}
	return err
}
