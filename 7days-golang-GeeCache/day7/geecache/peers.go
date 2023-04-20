package geecache

import pb "geecache/geecachepb"

type PeerGetter interface {
	Get(in *pb.Request, out *pb.Response) error
}

type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}
