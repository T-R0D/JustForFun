/**
    @file project_euler_problem_5.cpp

    @author Terence Henriod

    Project Euler Problem 5

    @brief Solves the following problem for the general case of any range of 
           required factors:

  "2520 is the smallest number that can be divided by each of the numbers from
   1 to 10 without any remainder.

   What is the smallest positive number that is evenly divisible by all of the
   numbers from 1 to 20?"

    @version Original Code 1.00 (12/28/2013) - T. Henriod


RE_WORK, USE PRIME FACTORIZATION METHOD

*/

/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   HEADER FILES / NAMESPACES
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
#include <cassert>
#include <cmath>
#include <iostream>
#include <cstdio>
using namespace std;


/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   GLOBAL CONSTANTS
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/


/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   FUNCTION PROTOTYPES
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
/**
FunctionName

A short description

@param

@return

@pre
-# 

@post
-# 

@detail @bAlgorithm
-# 

@exception

@code
@endcode
*/
unsigned long long findLeastCommonMultipleOfRange(const unsigned long long range_start,
                                         const unsigned long long range_end);

/**
FunctionName

A short description

@param

@return

@pre
-# 

@post
-# 

@detail @bAlgorithm
-# 

@exception

@code
@endcode
*/
bool isDivisibleByRange(const unsigned long long number, const unsigned long long range_start,
                        const unsigned long long range_end);
      
/**
FunctionName

A short description

@param

@return

@pre
-# 

@post
-# 

@detail @bAlgorithm
-# 

@exception

@code
@endcode
*/
//unsigned long long listProduct(const list<unsigned long long> numbers);

/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   MAIN FUNCTION
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
int main() {
  // variables
  unsigned long long range_start = 0;
  unsigned long long range_end = 0;
  unsigned long long result = 0;

  // get the start of the range of required factors
  printf("Enter the start of the range (inclusive): ");
  scanf("%d", &range_start);

  // case: range_start was 1
  if (range_start <= 1) {
    // to prevent trivialities, bump it to 2
    range_start = 2;
  }

  // get the end of the range of required factors
  printf("Enter the end of the range (inclusive): ");
  scanf("%d", &range_end);

  // compute the result
  result = findLeastCommonMultipleOfRange(range_start, range_end);

  // state the result for the user
  printf("\nThe lowest common multiple of the set is: %u \n", result);

  // return 0 on successful completion
  return 0;
}


/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   FUNCTION IMPLEMENTATIONS
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
unsigned long long findLeastCommonMultipleOfRange(const unsigned long long range_start,
                                         const unsigned long long range_end) {
  // variables
  unsigned long long product = 1;
  unsigned long long factor = (range_end / 2) + 1;
  unsigned long long temp = 0;

  // case: the start of the range is greater than half of the upper bound
  if (factor < range_start) {
    // simply start the range at the low end of the range
    factor = range_start;
  } 

  // generate a product of the greater half of the range
  // (the lower half would be redundant)
  while (factor <= range_end) {
    // include the current number in the product
    product *= factor;

    // move up to the next factor
    factor++;
  }

  // attempt to reduce the number by successively dividing by each factor in the
  // range
  factor = range_start;
  while (factor <= range_end) {
    // perform the reduction attempt
    temp = product / factor;

    // case: the current result can be reduced by the currently considered
    //       factor and still be divisible by the range
    if (isDivisibleByRange(temp, range_start, range_end)) {
      // keep the reduced number
      product = temp;

      // start the process over
      factor = range_start;
    }
    // case: result could not be reduced by the current factor and still work
    else {
      // try the next factor down
      factor++;
    }
  }

  // return the resulting product
  return product;
}

bool isDivisibleByRange(const unsigned long long number, const unsigned long long range_start,
                        const unsigned long long range_end) {
  // variables
  bool result = false;
  unsigned long long temp = range_end;

  // check for even divisibility by all numbers in range
  while ((temp >= range_start) && ((number % temp) == 0)) {
    // move down to the next number in the range
    temp--;
  }

  // case: all numbers down to the range_start were factors
  if (temp < range_start) {
    // the number is divisible by the range
    result = true;
  }

  // return the result
  return result;
}

/*
unsigned long long listProduct(const list<unsigned long long> numbers) {
  // variables
  unsigned long long product = 1;
  list<unsigned long long>::iterator ndx = numbers.begin();

  // use all numbers in the list
  while (ndx != numbers.end()) {
    // inlude the number in the list to the product
    product *= *ndx;

    // mve on to the next list element
    ++ndx;
  }

  // return the resulting product
  return product;
}
*/
