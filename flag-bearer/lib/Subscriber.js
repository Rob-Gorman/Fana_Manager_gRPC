require('dotenv').config();
const redis = require('redis');

const CHANNELS = process.env.CHANNELS || [
  'flag-update-channel',
  'flag-toggle-channel',
];

class Subscriber {
  constructor(port, host, manager, password) {
    this.redis = redis.createClient({
      name: 'flag-bearer',
      socket: {
        host,
        port,
      },
      password
    });
    this.init();
    this.manager = manager;

    this.redis.on('connect', () =>
      console.log(`Redis subscriber on port ${port}`)
    );
  }

  async init() {
    this.redis.on('error', (err) => console.error('Error: ' + err));

    try {
      await this.redis.connect();
      await Promise.all([CHANNELS.map((channel) => this.subscribeTo(channel))]);
    } catch (err) {
      console.error('Error: ' + err);
    }
  }

  async subscribeTo(channelName) {
    try {
      await this.redis.subscribe(channelName, (message, channel) => {
        console.log(`${channel}: ${message}\n`);

        this.publish(channel, JSON.parse(message));
      });
    } catch (err) {
      console.error(err);
    }
  }

  publish(channel, data) {
    let status;
    for (let key in data) {
      status = data[key].status;
    }

    if (channel === 'flag-toggle-channel' && status) {
      this.manager.subscriptions.servers.forEach((client) => {
        client.stream.write(`event: ${client.sdkKey}\n`);
        client.stream.write(`channel: ${channel}\n`);
        client.stream.write(`data: ${JSON.stringify(data)}\n`);
        client.stream.write(`\n\n`);
      });
    } else if (channel === 'flag-toggle-channel' && !status) {
      // all sdk streams get updated when a flag is toggle off
      for (let sdkType in this.manager.subscriptions) {
        if (!this.manager.subscriptions[sdkType].length) {
          continue
        }

        this.manager.subscriptions[sdkType].forEach((client) => {
          // console.log('=== check client stream: ', client.stream);
          client.stream.write(`event: ${client.sdkKey}\n`);
          client.stream.write(`channel: ${channel}\n`);
          client.stream.write(`data: ${JSON.stringify(data)}\n`);
          client.stream.write(`\n\n`);
        });
      }
    } else {
      // only server streams get updates to individual flags
      this.manager.subscriptions.servers.forEach((client) => {
        client.stream.write(`event: ${client.sdkKey}\n`);
        client.stream.write(`channel: ${channel}\n`);
        client.stream.write(`data: ${JSON.stringify(data)}\n`);
        client.stream.write(`\n\n`);
      });
    }
  }
}

module.exports = Subscriber;
