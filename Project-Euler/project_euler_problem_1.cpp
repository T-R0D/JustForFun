/**
    @file project_euler_problem_1.cpp

    @author Terence Henriod

    Project Euler Problem 1

    @brief Solves the following problem for the general case of 2 multiple
           bases and a given range:

  "If we list all the natural numbers below 10 that are multiples of 3 or 5, we
   get 3, 5, 6 and 9. The sum of these multiples is 23.

   Find the sum of all the multiples of 3 or 5 below 1000."

    @version Original Code 1.00 (12/26/2013) - T. Henriod
*/

/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   HEADER FILES / NAMESPACES
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
#include <cassert>
#include <iostream>
using namespace std;


/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   GLOBAL CONSTANTS
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
const int kFirstNaturalNumber = 1;

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
long long sumSetsOfMultiples(const int first_number,
                             const int second_number,
                             const int range_start,
                             const int range_end);

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
int findLCM(const int first_number, const int second_number);

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
long long sumMultiples(const int number, const int range_start,
                       const int range_end);


/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   MAIN FUNCTION
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
int main() {
  // variables
  int range_start = kFirstNaturalNumber;
  int range_end = 1000;
  int first_number = 0;
  int second_number = 0;
  long long result = 0;

  // get the first number to be used in the problem
  cout << "Enter the first number: ";
  cin >> first_number;
  cout << endl;

  // get the second number to be used in the problem
  cout << "Enter the second number: ";
  cin >> second_number;
  cout << endl;

  // get the start of the range
  cout << "Enter the start of the range (inclusive): ";
  cin >> range_start;
  cout << endl;

  // get the end of the range
  cout << "Enter the end of the range (inclusive): ";
  cin >> range_end;
  cout << endl;

  // compute the result
  result = sumSetsOfMultiples(first_number, second_number,
                              range_start, range_end);

  // state the results for the user
  cout << "The sum of all multiples of " << first_number << " and "
       << second_number << endl
       << "between " << range_start << " and " << range_end << endl
       << "is " << result << endl;

  // return 0 on successful completion
  return 0;
}


/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   FUNCTION IMPLEMENTATIONS
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
long long sumSetsOfMultiples(const int first_number,
                             const int second_number,
                             const int range_start,
                             const int range_end) {
  // variables
  long long result = 0;
  int least_common_multiple = 0;

  // compute the least common multiple of the two numbers
  least_common_multiple = findLCM(first_number, second_number);

  // use inclusion-exclusion to find the sum of all the unique
  // multiples of the two numbers
  result = sumMultiples(first_number, range_start, range_end) +
           sumMultiples(second_number, range_start, range_end) -
           sumMultiples(least_common_multiple, range_start, range_end);

  // return result
  return result;
}

int findLCM(const int first_number, const int second_number) {

  // TODO: use a better algorithm

  // variables
  int lowest_common_multiple = 0;
  int larger = first_number;
  int smaller = second_number;

  // case: the first number is actually smaller
  if (first_number < second_number) {
    // reset the starting values
    larger = second_number;
    smaller = first_number;
  }

  // set the lowest common multiple variable
  lowest_common_multiple = smaller;

  // continually add smaller to the LCM variable until LCM is a
  // multiple of larger
  while ((lowest_common_multiple % larger) != 0) {
    // increase the LCM variable
    lowest_common_multiple += smaller;
  }

  // return the result
  return lowest_common_multiple;
}

long long sumMultiples(const int number, const int range_start,
                       const int range_end) {
  // assert pre-conditions
  assert(number > 0);
  assert(range_start > 0);
  assert(range_end >= range_start);

  // variables
  long long total = 0;
  int factor_sum = 0;
  int low_factor = 0;
  int high_factor = 0;
  int num_factors = 0;

  // find the other factor of the lowest multiple of the number in the range
  low_factor = range_start / number;

  // case: the number doesn't divide evenly into the start of the range
  if ((range_start % number) > 0) {
    // factor needs to be one larger for the first multiple to fall in the range
    low_factor++;
  }

  // find the other factor of the highest number in the range
  high_factor = range_end / number;

  // compute the number of factors
  num_factors = ((high_factor - low_factor) + 1);

  // compute the sum of the factors
  factor_sum = (low_factor + high_factor) * (num_factors / 2);

  // case: there was an odd number of factors to sum
  if ((num_factors % 2) == 1) {
    // add the median value to the sum
    factor_sum += (low_factor + high_factor) / 2;
  }

  // compute the sum of all the multiples
  total = (long long)factor_sum * (long long)number;

  //return resulting sum
  return total;
}
