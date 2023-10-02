package websocket_example

type Message struct {
	UserID  int
	Payload string
}

type Connection struct {
	UserID   int
	DeviceID string
}

func (c *Connection) Write(p []byte) (n int, err error) {
	// Pretend it is implemented
	return 0, nil
}

type WSServer struct {
	connectedClientsCount uint64
}

func (w *WSServer) handleConnect(c Connection) {
	// TODO
}

func (w *WSServer) handleDisconnect(c Connection) {
	// TODO
}

func (w *WSServer) totalConnectedClients() uint64 {
	return w.connectedClientsCount
}

func (w *WSServer) handleQueueMessages(messages []Message) (int, error) {
	for i, m := range messages {
		err := w.sendToConnectedDevices(m)
		if err != nil {
			return i, err
		}
	}
	return len(messages), nil
}

func (w *WSServer) sendToConnectedDevices(m Message) error {
	// TODO
	return nil
}
