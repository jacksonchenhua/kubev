package model

type K8sNodes struct {
	JoinString  string
	MasterNode  *K8sNode
	WorkerNodes []*K8sNode
}

type K8sNode struct {
	VMName         string
	IP             string
	Mo             string
	FolderPath     string
	DatacenterName string
	DatastoreName  string
	MasterNode     bool
	Ready          bool
}
