package common

type StoreType int

const (
	_ StoreType = iota
	StoreLocal
	StoreOss
	StoreCeph
	// StoreMix : 混合(Ceph及OSS)
	StoreMix
	// StoreAll : 所有类型的存储都存一份数据
	StoreAll
)
