const MINIMUM_ID_LENGTH = 30;
// recommended ms to wait in-between failed SSE connection attempts
const SEC_TO_RETRY = 7000;
const SUBSCRIPTION_TYPES = ['server', 'client'];

class ClientsManager {
  constructor() {
    this.subscriptions = { servers: [], clients: [] }; // tracks open SSE connections & their response objects
    this.responseHeaders = {
      'Content-Type': 'text/event-stream',
      Connection: 'keep-alive',
      'Cache-Control': 'no-cache',
      "Access-Control-Allow-Origin": "*",
      "Access-Control-Allow-Headers": "Origin, X-Requested-With, Content-Type, Accept"
    };

    this.subscriptionTypes = SUBSCRIPTION_TYPES;
    this.retryTimeout = SEC_TO_RETRY;
  }

  stream(req, res) {
    const { sdkType, id } = req.params;
    const sdkKey = req.query.sdkKey
    
    // sdkType is either 'client' or 'server'
    if (!this.subscriptionTypes.includes(sdkType)) {
      return res.status(400).send('Invalid subscription topic.');
    }

    const client = this.addClient(sdkType, id, sdkKey);
    this.connectClient(client, req, res);
    this.reportOnAllConnections();
  }

  // SSE connections are organized by sdk type
  addClient(sdkType, id, sdkKey) {
    if (!id || typeof id !== 'string')
      id = this.generateRandomId(MINIMUM_ID_LENGTH);

    const newClient = { id, sdkKey };

    if (sdkType === 'client') {
      this.subscriptions.clients.push(newClient);
    } else {
      this.subscriptions.servers.push(newClient);
    }
    return newClient;
  }

  // create initial SSE streaming connection
  connectClient(client, req, res) {
    res.writeHead(200, this.responseHeaders);
    res.write(`retry: ${this.retryTimeout}.\n`);
    res.write(`data: Success: subscribed to messages.`);
    res.write('\n\n');
    console.log(`--- client ${client.id} connected successfully. ---`);

    // store response obj to be written to later
    client.stream = res;

    // remove SSE client from appropriate subscriptions list when conn is closed
    req.on('close', () => {
      this.subscriptions.clients = this.subscriptions.clients.filter(
        (c) => c.id !== client.id
      );
      console.log(`client ${client.id} disconnected`);
      this.reportOnAllConnections();
    });
  }

  reportOnAllConnections() {
    let totalConns = 0;
    for (let sdkType in this.subscriptions) {
      totalConns += this.subscriptions[sdkType].length;
    }
    console.log(`--- ${totalConns} total connections ---`);
  }

  generateRandomId(size) {
    return [...Array(size)]
      .map(() => {
        return Math.floor(Math.random() * 36).toString(36);
      })
      .join('');
  }
}

module.exports = ClientsManager;
