package network

type PeerAddress struct {
	ip   string
	port int
}

type BasicPeer struct {
	peerAddress PeerAddress
	peerId      int
}

type peer struct {
	BasicPeer BasicPeer
	isRun     bool
	viewOK    bool
}
