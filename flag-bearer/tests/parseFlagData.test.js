// const { evaluateFlags } = require('../utils/parseFlagData');
const { evaluateCondition } = require('../utils/evaluateCondition');
const { data, testUser1, testUser2, testUser3 } = require('./data');
// const {cache} = require('../services/services')
const RedisCache = require('../lib/RedisCache');
// how to mock evaluateFlags without a cache?

const betaCondition = {
  negate: false,
  operator: 'EQ',
  attribute: 'beta',
  vals: ['true'],
};
const noBetaConditon = {
  negate: true,
  operator: 'EQ',
  attribute: 'beta',
  vals: ['true'],
};

// can i use this instead of real cache? 
const cache = {
  flags: data.flags,
  sdkKeys: data.sdkKeys,
};

describe('evaluate beta condition for users', () => {
  console.log('access to cache', cache);

  test('user with beta true returns true', () => {
    expect(evaluateCondition(testUser1.userContext, betaCondition)).toEqual(true);
  });
  test('user without beta attribute returns false', () => {
    expect(evaluateCondition(testUser3.userContext, betaCondition)).toEqual(false);
  });
  test('user with beta false returns false', () => {
    expect(evaluateCondition(testUser2.userContext, betaCondition)).toEqual(false);
  });
  test('negated beta condition for user with beta true', () => {
    expect(evaluateCondition(testUser1.userContext, noBetaConditon)).toEqual(false);
  });
});


// describe('for a flag toggled off', () => {
//   test('all users are not served feature', () => {
//     users.map(user => {
//       expect(evaluateFlags(user.userContext)).toHaveProperty()
//     })
//   })
// })

/*
- beta is true
- Not beta
- off flag for all
- on flag for all,
- cali or student
- cali AND student
- neither cali nor student
- userid is in 
- user id contains "u"
- age is greater than 
- age lt eq 20
*/
