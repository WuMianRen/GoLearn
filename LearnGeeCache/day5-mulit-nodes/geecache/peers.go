/*
@Time : 2020/9/11 16:30
@Author : wumian
@File : peers
@Software: GoLand
*/
package geecache

type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

type PeerGetter interface {
	Get(group string, key string) ([]byte, error)
}
