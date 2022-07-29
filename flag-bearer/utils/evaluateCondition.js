// for both strings and booleans
// targetVal is from flag rules
// candidateVal is from userContext
// if condition is negate, flip
const isEq = (targetVal, candidateVal) => {
  if (typeof candidateVal === 'string') {
    candidateVal = candidateVal.toLowerCase();
  }

  return candidateVal === targetVal;
};

const isIn = (targetVals, candidateVal) =>
  targetVals.includes(candidateVal.toLowerCase());
const isNotIn = (targetVals, candidateVal) =>
  !targetVals.includes(candidateVal.toLowerCase());
const strContains = (targetVal, candidateVal) =>
  candidateVal.toLowerCase().includes(targetVal);
const strEndsWith = (targetVal, candidateVal) =>
  candidateVal.toLowerCase().endsWith(targetVal);
const strStartsWith = (targetVal, candidateVal) =>
  candidateVal.toLowerCase().startsWith(targetVal);

const isGreaterThan = (targetVal, candidateVal) => candidateVal > targetVal;
const isGreaterEqThan = (targetVal, candidateVal) => candidateVal >= targetVal;
const isLessThan = (targetVal, candidateVal) => candidateVal < targetVal;
const isLessEqThan = (targetVal, candidateVal) => candidateVal <= targetVal;

let operandMapper = {
  EQ: isEq,
  IN: isIn,
  NOT_IN: isNotIn,
  STR_CONTAINS: strContains,
  STR_ENDS_WITH: strEndsWith,
  STR_STARTS_WITH: strStartsWith,
  GT: isGreaterThan,
  LT: isLessThan,
  LT_EQ: isLessEqThan,
  GT_EQ: isGreaterEqThan,
};

// convert targetAttribute (from flag rule)
const convertAttributeType = (attributeValue, op) => {
  switch (op) {
    // EQ will always compare 1:1
    case 'EQ':
      attributeValue = attributeValue[0];
      // string, do nothing
      // boolean, number convert
      if (attributeValue === 'true' || attributeValue === 'false') {
        return Boolean(attributeValue);
      } else if (!isNaN(Number(attributeValue))) {
        return Number(attributeValue);
      }
      break;
    // following cases will always be just a single number
    case 'GT':
    case 'LT':
    case 'LT_EQ':
    case 'GT_EQ':
      attributeValue = attributeValue[0];
      return Number(attributeValue);
    // following cases won't be multiple values
    case 'STR_CONTAINS':
    case 'STR_ENDS_WITH':
    case 'STR_STARTS_WITH':
      attributeValue = attributeValue[0];
      break;
  }
  return attributeValue;
};

// for one attribute
const evaluateCondition = (userContext, condition) => {
  const op = condition.operator;
  const attribute = condition.attribute;
  // cast targetValue type based on operator and candidateValue type
  const targetValue = convertAttributeType(condition.vals, op);
  const candidateValue = userContext[attribute];
  let result;
  // if attribute not provided in userContext, always false regardless of negate
  if (!userContext.hasOwnProperty(attribute)) return false;
  result = operandMapper[op](targetValue, candidateValue);
  return condition.negate ? !result : result;
};

module.exports = {
  evaluateCondition,
};
