const express = require('express');
const router = express.Router();
const { validateClientInit } = require('../validators/validators');
const { initializeServerSDK } = require('../controllers/serverSdkController');
const { initializeClientSDK } = require('../controllers/clientSdkController');
const {clientManager, cache } = require('../services/services')

// receives client SDK initialization requests
router.post(
  `/connect/clientInit`,
  validateClientInit,
  initializeClientSDK
);

// SSE connection endpoint
router.get('/stream/:sdkType', (req, res, next) => {
  const sdkKey = req.query.sdkKey
  
  if (!cache.sdkKeys) {
    return res.status(400).send({ error: 'Flag bearer cannot accept stream connections at this time...try again later.'})
  }
  if (!cache.validSdkKey(sdkKey)) {
    console.log('Cant find sdkKey in cache:', cache.sdkKeys);
    return res.status(400).send({ error: 'Invalid SDK key.' });
  } else {
    clientManager.stream(req, res, next)
  }
});

// receives server SDK initialization requests
router.get(`/connect/serverInit`, initializeServerSDK);

router.get('/', (req, res) => {
  return res
      .status(200)
      .send({ message: 'Flag bearer healthy!' });
  
})
module.exports = router;
