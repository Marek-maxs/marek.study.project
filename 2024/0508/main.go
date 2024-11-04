package main

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/en-trak/sbos.energy/v2/models"
	pb "github.com/en-trak/protobuf/v2/sbos/energy"
	"time"
	"unsafe"
)

type DatapointDataConsumption struct {
	TmpVirtualItems      map[string]map[string]float64
	Co2Value             *Co2
	VirtualDatapointList map[uuid.UUID]*models.EnergyVirtualDatapoint
	TimeLoc              *time.Location
	Relationship         map[string][]NodeValue
	NewEnergyUnitPrice   map[string]map[string][]float64
	EndTime              int64
	StartTime            int64
	Area                 float64
	VirtualCondIndex     int
	VirtualDatapointNum  int
	TimezoneHour         int32
	Unit                 pb.MetricType
	Interval             string
	VirtualCond          []string
	DatapointID          uuid.UUID
}

type Co2 struct {
	Values       map[string]float64
	Datapoints   []interface{}
	AlreadyQuery bool
}

type NodeValue struct {
	ConstantOperator   string
	Operator           string
	Constant           float64
	IsChildNode        int
	Sort               int
	DataType           int
	VirtualDatapointID uuid.UUID
	DependID           uuid.UUID
}

type LiveData struct {
	Ctx                    context.Context
	DatapointRefList       map[string]int64
	NewEnergyUnitPrice     map[string]map[string][]float64
	DatapointMeterMap      map[string]string
	TimeLoc                *time.Location
	TmpVirtualItems        map[string]map[string]float64
	VirtualDatapointList   map[uuid.UUID]*models.EnergyVirtualDatapoint
	Relationship           map[string][]NodeValue
	NewCo2ConversionFactor map[string]float64
	Area                   float64
	St                     int64
	Et                     int64
	VirtualCondIndex       int
	VirtualDatapointNum    int
	Unit                   pb.MetricType
	TimezoneHour           int32
	CondArr                []string
	VirtualCond            []string
}

func main() {
	fmt.Printf("LiveData   大小 %d 对齐系数 %d\n", unsafe.Sizeof(LiveData{}), unsafe.Alignof(LiveData{}))
	fmt.Printf("LiveData.Ctx 大小 %d 对齐系数 %d 偏移量 %d\n", unsafe.Sizeof(LiveData{}.Ctx), unsafe.Alignof(LiveData{}.Ctx), unsafe.Offsetof(LiveData{}.Ctx))
	fmt.Printf("LiveData.DatapointRefList 大小 %d 对齐系数 %d 偏移量 %d\n", unsafe.Sizeof(LiveData{}.DatapointRefList), unsafe.Alignof(LiveData{}.DatapointRefList), unsafe.Offsetof(LiveData{}.DatapointRefList))
	fmt.Printf("LiveData.NewEnergyUnitPrice 大小 %d 对齐系数 %d 偏移量 %d\n", unsafe.Sizeof(LiveData{}.NewEnergyUnitPrice), unsafe.Alignof(LiveData{}.NewEnergyUnitPrice), unsafe.Offsetof(LiveData{}.NewEnergyUnitPrice))
	fmt.Printf("LiveData.DatapointMeterMap 大小 %d 对齐系数 %d 偏移量 %d\n", unsafe.Sizeof(LiveData{}.DatapointMeterMap), unsafe.Alignof(LiveData{}.DatapointMeterMap), unsafe.Offsetof(LiveData{}.DatapointMeterMap))
	fmt.Printf("LiveData.TimeLoc 大小 %d 对齐系数 %d 偏移量 %d\n", unsafe.Sizeof(LiveData{}.TimeLoc), unsafe.Alignof(LiveData{}.TimeLoc), unsafe.Offsetof(LiveData{}.TimeLoc))
	fmt.Printf("LiveData.TmpVirtualItems 大小 %d 对齐系数 %d 偏移量 %d\n", unsafe.Sizeof(LiveData{}.TmpVirtualItems), unsafe.Alignof(LiveData{}.TmpVirtualItems), unsafe.Offsetof(LiveData{}.TmpVirtualItems))
	fmt.Printf("LiveData.VirtualDatapointList 大小 %d 对齐系数 %d 偏移量 %d\n", unsafe.Sizeof(LiveData{}.VirtualDatapointList), unsafe.Alignof(LiveData{}.VirtualDatapointList), unsafe.Offsetof(LiveData{}.VirtualDatapointList))
	fmt.Printf("LiveData.Relationship 大小 %d 对齐系数 %d 偏移量 %d\n", unsafe.Sizeof(LiveData{}.Relationship), unsafe.Alignof(LiveData{}.Relationship), unsafe.Offsetof(LiveData{}.Relationship))
	fmt.Printf("LiveData.NewCo2ConversionFactor 大小 %d 对齐系数 %d 偏移量 %d\n", unsafe.Sizeof(LiveData{}.NewCo2ConversionFactor), unsafe.Alignof(LiveData{}.NewCo2ConversionFactor), unsafe.Offsetof(LiveData{}.NewCo2ConversionFactor))
	fmt.Printf("LiveData.Area 大小 %d 对齐系数 %d 偏移量 %d\n", unsafe.Sizeof(LiveData{}.Area), unsafe.Alignof(LiveData{}.Area), unsafe.Offsetof(LiveData{}.Area))
	fmt.Printf("LiveData.St 大小 %d 对齐系数 %d 偏移量 %d\n", unsafe.Sizeof(LiveData{}.St), unsafe.Alignof(LiveData{}.St), unsafe.Offsetof(LiveData{}.St))
	fmt.Printf("LiveData.Et 大小 %d 对齐系数 %d 偏移量 %d\n", unsafe.Sizeof(LiveData{}.Et), unsafe.Alignof(LiveData{}.Et), unsafe.Offsetof(LiveData{}.Et))
	fmt.Printf("LiveData.VirtualCondIndex 大小 %d 对齐系数 %d 偏移量 %d\n", unsafe.Sizeof(LiveData{}.VirtualCondIndex), unsafe.Alignof(LiveData{}.VirtualCondIndex), unsafe.Offsetof(LiveData{}.VirtualCondIndex))
	fmt.Printf("LiveData.VirtualDatapointNum 大小 %d 对齐系数 %d 偏移量 %d\n", unsafe.Sizeof(LiveData{}.VirtualDatapointNum), unsafe.Alignof(LiveData{}.VirtualDatapointNum), unsafe.Offsetof(LiveData{}.VirtualDatapointNum))
	fmt.Printf("LiveData.Unit 大小 %d 对齐系数 %d 偏移量 %d\n", unsafe.Sizeof(LiveData{}.Unit), unsafe.Alignof(LiveData{}.Unit), unsafe.Offsetof(LiveData{}.Unit))
	fmt.Printf("LiveData.TimezoneHour 大小 %d 对齐系数 %d 偏移量 %d\n", unsafe.Sizeof(LiveData{}.TimezoneHour), unsafe.Alignof(LiveData{}.TimezoneHour), unsafe.Offsetof(LiveData{}.TimezoneHour))
	fmt.Printf("LiveData.CondArr 大小 %d 对齐系数 %d 偏移量 %d\n", unsafe.Sizeof(LiveData{}.CondArr), unsafe.Alignof(LiveData{}.CondArr), unsafe.Offsetof(LiveData{}.CondArr))
	fmt.Printf("LiveData.VirtualCond 大小 %d 对齐系数 %d 偏移量 %d\n", unsafe.Sizeof(LiveData{}.VirtualCond), unsafe.Alignof(LiveData{}.VirtualCond), unsafe.Offsetof(LiveData{}.VirtualCond))

}