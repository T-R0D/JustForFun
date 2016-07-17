/**
    @file project_euler_problem_6.cpp

    @author Terence Henriod

    Project Euler Problem 6

    @brief Solves the following problem for the general case of any range of 
           required numbers (almost):

  "The sum of the squares of the first ten natural numbers is,
   1^2 + 2^2 + ... + 10^2 = 385
   The square of the sum of the first ten natural numbers is,
   (1 + 2 + ... + 10)^2 = 552 = 3025
   Hence the difference between the sum of the squares of the first ten
   natural numbers and the square of the sum is 3025 ? 385 = 2640.

   Find the difference between the sum of the squares of the first one hundred
   natural numbers and the square of the sum."

    @version Original Code 1.00 (12/29/2013) - T. Henriod
*/

/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   HEADER FILES / NAMESPACES
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
#include <cassert>
#include <cmath>
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
unsigned int sumRangeOfNumbers(const unsigned int range_start,
                               const unsigned int range_end);

unsigned int sumRangeOfSquares(const unsigned int range_start,
                               const unsigned int range_end);


/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   MAIN FUNCTION
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
int main() {
  // variables
  unsigned int range_start = 0;
  unsigned int range_end = 0;
  unsigned int sum_of_range = 0;
  unsigned int sum_squared = 0;
  unsigned int sum_of_squares = 0;
  unsigned int difference = 0;

  // get the start of the range of the numbers
  printf("Enter the start of the range (inclusive): ");
  scanf("%u", &range_start);


  // get the end of the range of the numbers
  printf("Enter the end of the range (inclusive): ");
  scanf("%u", &range_end);

  // compute the sum of the numbers, squared
  sum_of_range = sumRangeOfNumbers(range_start, range_end);
  sum_squared = sum_of_range * sum_of_range;

printf("\nThe sum squared is: %u\n",sum_squared);

  // compute the sum of the squared numbers
  sum_of_squares = sumRangeOfSquares(range_start, range_end);

printf("\nThe sum of the squares is: %u\n",sum_of_squares);

  // find the difference
  difference = sum_squared - sum_of_squares;

  // state the result for the user
  printf("\nThe difference of the sum of the numbers squared and the");
  printf("\nsum of the squared numbers is: %u \n", difference);

  // return 0 on successful completion
  return 0;
}


/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   FUNCTION IMPLEMENTATIONS
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
unsigned int sumRangeOfNumbers(const unsigned int range_start,
                               const unsigned int range_end) {
  // variables
  unsigned int sum = 0;
  unsigned int num_pairs = (range_end - range_start + 1) / 2;
  unsigned int pair_sum  = range_start + range_end;

  // compute the sum of all the pairs
  sum = num_pairs * pair_sum;

  // case: there was an odd number of numbers in the range
  if (((range_end - range_start) % 2) == 0) {
    // be sure to add the median value to the sum as well
    sum += (range_end / 2) + range_start;
  }

  // return the resulting sum
  return sum;
}

unsigned int sumRangeOfSquares(const unsigned int range_start,
                               const unsigned int range_end) {

//TODO: generalize this to accept ranges x to n, not just 1 to n

  // variables
  unsigned int sum = 0;

  // apply the formula that results from summing all integers i^2
  // for i = 1 to n: (2n^3 + 3n^2 + n) / 6
  sum = ((2 * pow(range_end, 3)) + (3 * pow(range_end, 2)) + range_end) / 6;

  // return the resulting sum
  return sum;
}


