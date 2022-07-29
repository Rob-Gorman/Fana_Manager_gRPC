const { validationResult } = require('express-validator');

const createFlagset = (req, res, next) => {
  const errors = validationResult(req);
  // checks: flatset is array, sdkkey and flags array are provided
  if (errors.isEmpty()) {
    res.status(201).send('201: Flagset created');
  } else {
    return res.status(404).send('Input field error.');
  }
};

module.exports = {
  createFlagset,
};