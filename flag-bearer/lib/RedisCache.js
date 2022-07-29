const redis = require('redis');
const { getRuleset } = require('../utils/apiClient');

class RedisCache {
  constructor(port, host, password) {
    this.redis = redis.createClient({
      name: 'flag-bearer-cache',
      socket: {
        host,
        port,
      },
      password
    });
    this.init();
    this.sdkKeys = {};
    this.flags = {};

    this.redis.on('connect', () => console.log(`Redis cache on port ${port}`));
  }

  async init() {
    this.redis.on('error', (err) => console.error('Error: ' + err));

    try {
      await this.redis.connect();
      console.log('refreshing cache...\n');
      await this.refreshData();


    } catch (err) {
      console.error('--- Error connecting to Redis server: ' + err);
    }
  }
  // fetch full flag ruleset from manager
  async refreshData() {
    try {
      let { sdkKeys, flags } = await getRuleset();
      this.sdkKeys = sdkKeys;
      this.flags = flags;

    } catch (err) {
      console.error('--- Could not fetch data from manager...cache still empty');
    }
  }
  // try reading from cache first
  async getData() {
    try {
      const response = await this.redis.get('data');

      if (response) {
        const { sdkKeys, flags } = JSON.parse(response);
        this.sdkKeys = sdkKeys;
        this.flags = flags;
        console.log('--- Got data from redis successfully: ', this.flags);
      } else {
        // cache is empty or network connection error
        console.log('--- Cache is empty. Need to fetch from Manager...');
        await this.refreshData();
      }
      // returning this for serverSDK
      return { sdkKeys: this.sdkKeys, flags: this.flags };
    } catch (err) {
      console.error(err);
      console.error('=== Error connecting to redis === ');
    }
  }

  validSdkKey(sdkKey) {
    return this.sdkKeys[sdkKey];
  }
}

module.exports = RedisCache;
