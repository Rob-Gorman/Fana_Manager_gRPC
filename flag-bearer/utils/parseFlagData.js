const { evaluateCondition } = require('./evaluateCondition')
const {cache} = require('../services/services')

// evaluate audience conditions based on user attributes
function evaluateAudience(audienceContext, userContext) {
  let evaluation = false; // default to false
  for (const condition of audienceContext.conditions) {
    if (evaluateCondition(userContext, condition)) {
      evaluation = true;
      if (audienceContext.combine === 'ANY') {
        // if 'ANY' set, then return as soon as one condition is satisfied
        break;
      } else if (audienceContext.combine === 'ALL') {
        // keep going until all conditions are evaluated
        continue;
      }
    } else { // condition was false
      // if condition is not met and ANY, keep going until theres no more conditions
      if (audienceContext.combine === 'ALL') {
        // exit early because at least one condition was not met.
        evaluation = false;
        break;
      }
      // 'ANY' and false, go to next condition
    }
  }
  return evaluation;
}

// Returns overall flag evaluation based on set of audience logic evaluations for a given user
function evaluateFlags(userContext) {
  const flagEvals = {};
  const audienceEvals = {};
  // const userAudienceEvals = evaluateAudiences(sdkInstance, userContext);
  // Evaluate each flag
  for (const flag in cache.flags) {
    let evaluation = false;
    const { status, ...audiences} = cache.flags[flag]
    const audienceKeys = Object.keys(audiences)
    if (status && audienceKeys.length == 0)  {
      // flags without any audience targeting apply to everyone
      evaluation = true
    } else if (status) {
      // loop through all audiences to determine final flag evaluation
      // whenever audience eval is true, break. if it's false, go to next;
      for (const audience of audienceKeys) {
        if (audienceEvals[audience]) {
          evaluation = true;
          break; // return early as soon as one audience satisfied
        } else if (!audienceEvals.hasOwnProperty(audience)) {
          // if audience eval is undefine, evaluate;
          const audienceContext = cache.flags[flag][audience];
          audienceEvals[audience] = evaluateAudience(audienceContext, userContext)
          // console.log('audience:', audience, 'eval:', evaluateAudience(audienceContext, userContext))
          if (audienceEvals[audience]) {
            evaluation = true;
            break;
          }
        }
      }
    } 
    flagEvals[flag] = evaluation; 
  };
  console.log('evaluated flags: ', flagEvals)
  return flagEvals;
}

// // TEST
// let flags = {
//   sdkKeys: {
//     "beta_sdk_0":true,
//     "de9-6bf1a0c-3":true,
//     "fa4-d731f0e-4":true
//   },

//   "beta-header": {
//       "status": true,
//       "beta-testers": {
//         combine: "ANY",
//         conditions: [
//           {
//             "attribute": "beta",
//             "operator": "EQ",
//             "vals": ["true"],
//             "negate": false
//           }
//         ]
//       }
//   },
//   "CA-header": {
//     "status": true,
//     "na-testers": {
//       combine: "ALL",
//       conditions: [
//         {
//           "attribute": "country",
//           "operator": "IN",
//           "vals": ["canada", "usa"],
//           "negate": false
//         },
//         {
//           "attribute": "age",
//           "operator": "GT",
//           "vals": ["18"],
//           "negate": false
//         }
//       ]
//     }
//   }
// }
// const testUserContext1 = {
//   userId: 'JJUY',
//   beta: false,
//   age: 16,
//   country: 'CanadA',
// }
// const testUserContext2 = {
//   userId: 'JJUY',
//   beta: true,
//   age: 24,
//   country: 'CanadA',
// }
// console.log('NONE', evaluateFlags(testUserContext1))
// console.log('=====')
// console.log('beta & canada', evaluateFlags(testUserContext2))

module.exports = { evaluateFlags }