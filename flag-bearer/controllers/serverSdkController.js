const { validationResult } = require('express-validator');
const {cache} = require('../services/services');

const initializeServerSDK = async (req, res) => {
  const errors = validationResult(req);

  if (errors.isEmpty()) {
    const sdkKey = req.header('Authorization');
    const { flags } = await cache.getData();

    if (!cache.validSdkKey(sdkKey)) {
      return res.status(400).send({ error: 'Invalid SDK key.' });
    }

    res.json(flags);
  } else {
    res.status(401).send({ error: 'SDK key not authorized.' });
  }
};

module.exports = {
  initializeServerSDK,
};
