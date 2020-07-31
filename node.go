package main

import (
	"time"
)

const AGENT_CACHE_KEY string = "agent"

type AiAgentCacheMap map[string]AiAgentDetail

type AiAgentDetail struct {
	// 终端地址
	Uri string `json:"uri"`
	// 名称
	Name string `json:"name"`
	// 地理位置
	Address string `json:"address"`
	// 连接状态
	Connection bool `json:"connection,default='2'"`
	// 心跳时间
	HeartBeatTime time.Time `json:"heartBeatTime"`
	// 内外网标志
	InternalNetwork bool `json:"internalNetwork,default='2'"`
	// 注册设备信息
	DeviceIDs []uint64 `json:"deviceIDs"`
	// 注册分析类型信息
	AnalyzeTypes []uint8 `json:"analyzeTypes"`
	// 分析端GPU服务器节点信息
	NodeDetails []NodeDetail    `json:"nodeDetails,omitempty"`
	BaseInfo    ClusterBaseInfo `json:"baseInfo,omitempty"`
}

// 集群基本信息
type ClusterBaseInfo struct {
	// 机房机器数量
	ServerNum int `json:"serverNum,omitempty"`
	// 机房机平均负载(百分比)
	AvgLoad int `json:"avgLoad,omitempty"`
	// 总内存(bytes)
	TotalRam int `json:"totalRam,omitempty"`
	// 总磁盘(bytes)
	TotalDisk int `json:"totalDisk,omitempty"`
}

// 服务器节点信息
type NodeDetail struct {
	Os string `json:"os,omitempty"` // 操作系统
	IP string `json:"ip,omitempty"` // IP地址

	CPUUsedPercent int    `json:"cpuUsedPercent,omitempty"` // CPU利用率(百分比)
	CPUFramework   string `json:"cpuFramework,omitempty"`   // CPU架构

	MemUsedPercent int   `json:"memUsedPercent,omitempty"` // 内存利用率(百分比)
	MemTotalMem    int64 `json:"memTotalMem,omitempty"`    // 内存总内存(byte)
	MemUsedMem     int64 `json:"memUsedMem,omitempty"`     // 内存已使用(byte)

	DiskUsedPercent int   `json:"diskUsedPercent,omitempty"` // 磁盘利用率(百分比)
	DiskTotalMem    int64 `json:"diskTotalMem,omitempty"`    // 磁盘总内存(byte)
	DiskUsedMem     int64 `json:"diskUsedMem,omitempty"`     // 磁盘已使用(byte)

	LoadInfo int `json:"loadInfo,omitempty"` // 负载情况(百分比)
}
