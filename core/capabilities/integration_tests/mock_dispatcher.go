package integration_tests

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/services"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	remotetypes "github.com/smartcontractkit/chainlink/v2/core/capabilities/remote/types"
	p2ptypes "github.com/smartcontractkit/chainlink/v2/core/services/p2p/types"
)

type receiverKey struct {
	peerID       p2ptypes.PeerID
	capabilityId string
	donId        string
}

type testAsyncMessageBroker struct {
	services.StateMachine
	t *testing.T

	nodes map[receiverKey]remotetypes.Receiver

	sendCh chan *remotetypes.MessageBody

	stopCh services.StopChan
	wg     sync.WaitGroup
}

func (a *testAsyncMessageBroker) HealthReport() map[string]error {
	return nil
}

func (a *testAsyncMessageBroker) Name() string {
	return "testAsyncMessageBroker"
}

func newTestAsyncMessageBroker(t *testing.T, sendChBufferSize int) *testAsyncMessageBroker {
	return &testAsyncMessageBroker{
		t:      t,
		nodes:  make(map[receiverKey]remotetypes.Receiver),
		stopCh: make(services.StopChan),
		sendCh: make(chan *remotetypes.MessageBody, sendChBufferSize),
	}
}

func (a *testAsyncMessageBroker) Start(ctx context.Context) error {
	return a.StartOnce("testAsyncMessageBroker", func() error {
		a.wg.Add(1)
		go func() {
			defer a.wg.Done()

			for {
				select {
				case <-a.stopCh:
					return
				case msg := <-a.sendCh:
					receiverId := toPeerID(msg.Receiver)

					key := receiverKey{
						peerID:       receiverId,
						capabilityId: msg.CapabilityId,
						donId:        msg.CapabilityDonId,
					}

					receiver, ok := a.nodes[key]
					if !ok {
						panic("server not found for peer id")
					}

					receiver.Receive(tests.Context(a.t), msg)
				}
			}
		}()
		return nil
	})
}

func (a *testAsyncMessageBroker) Close() error {
	return a.StopOnce("testAsyncMessageBroker", func() error {
		close(a.stopCh)

		a.wg.Wait()
		return nil
	})
}

func (a *testAsyncMessageBroker) NewDispatcherForNode(nodePeerID p2ptypes.PeerID) remotetypes.Dispatcher {
	return &nodeDispatcher{
		callerPeerID: nodePeerID,
		broker:       a,
	}
}

func (a *testAsyncMessageBroker) RegisterReceiverNode(nodePeerID p2ptypes.PeerID, capabilityId string, capabilityDonID string, node remotetypes.Receiver) {
	key := receiverKey{
		peerID:       nodePeerID,
		capabilityId: capabilityId,
		donId:        capabilityDonID,
	}

	if _, ok := a.nodes[key]; ok {
		panic("capability already registered for peer id")
	}

	a.nodes[key] = node
}

func (a *testAsyncMessageBroker) Send(msg *remotetypes.MessageBody) {
	a.sendCh <- msg
}

func toPeerID(id []byte) p2ptypes.PeerID {
	return [32]byte(id)
}

type broker interface {
	Send(msg *remotetypes.MessageBody)
}

type nodeDispatcher struct {
	callerPeerID p2ptypes.PeerID
	broker       broker
}

func (t *nodeDispatcher) Send(peerID p2ptypes.PeerID, msgBody *remotetypes.MessageBody) error {
	msgBody.Version = 1
	msgBody.Sender = t.callerPeerID[:]
	msgBody.Receiver = peerID[:]
	msgBody.Timestamp = time.Now().UnixMilli()
	t.broker.Send(msgBody)
	return nil
}

func (t *nodeDispatcher) SetReceiver(capabilityId string, donId string, receiver remotetypes.Receiver) error {
	return nil
}
func (t *nodeDispatcher) RemoveReceiver(capabilityId string, donId string) {}
