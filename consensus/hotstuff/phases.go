package consensus

// QC 证书结构体
type QC struct {
	View       int
	NodeID     int
	Signatures []string
}
