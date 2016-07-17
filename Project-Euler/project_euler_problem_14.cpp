/**
  TODO: Improve by utilizing remembering previous sequences/lengths.

    @file project_euler_problem_14.cpp

    @author Terence Henriod

    Project Euler Problem 14

    @brief Solves the following problem for the general case of any number of
           large numbers:

  "The following iterative sequence is defined for the set of positive integers:

   n → n/2 (n is even)
   n → 3n + 1 (n is odd)

   Using the rule above and starting with 13, we generate the following
   sequence:

   13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1

   It can be seen that this sequence (starting at 13 and finishing at 1)
   contains 10 terms. Although it has not been proved yet (Collatz Problem), it
   is thought that all starting numbers finish at 1.

   Which starting number, under one million, produces the longest chain?

   NOTE: Once the chain starts the terms are allowed to go above one million."

    @version Original Code 1.00 (1/4/2014) - T. Henriod
*/

/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   HEADER FILES / NAMESPACES
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
#include <cassert>
#include <cmath>
#include <cstdio>
#include <iostream>
#include <iomanip>
#include <fstream>
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
unsigned int longestCollatzSequenceNumber(const unsigned int upper_bound);
unsigned int getCollatzSequenceLength(const unsigned int start_val);


/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   MAIN FUNCTION
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
int main() {
  // variables
  unsigned int upper_bound = 999999;
  unsigned int number = 1;
  unsigned int collatz_sequence_length = 0;

/*
  // get the highest term to be used
  printf("Enter the upper bound (exclusive) of terms to be considered: ");
  scanf("%u", &upper_bound);

*/

  // get the collatz number with the longest length that's below the bound
  number = longestCollatzSequenceNumber(upper_bound);

  // get the collatz sequence length of that number
  collatz_sequence_length = getCollatzSequenceLength(number);

  // report the result to the user
  printf("\nThe number %u has a Collatz sequence length of: %u\n", number,
         collatz_sequence_length);

  // return 0 on successful completion
  return 0;
}


/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   FUNCTION IMPLEMENTATIONS
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
unsigned int longestCollatzSequenceNumber(const unsigned int upper_bound) {
  // variables
  unsigned int great_collatz_number = 0;
  unsigned int test_number = 0;
  unsigned int great_length = 0;
  unsigned int test_length = 0;

  // test all numbers below the boundary
  for (test_number = 1, test_length = 0; test_number < upper_bound;
       test_number++) {
    // get the collatz sequence length
    test_length = getCollatzSequenceLength(test_number);

    // case: the test sequence length is longer than the current longest
    if (test_length > great_length) {
      // store the new number with the longest length and its corresponding
      // sequence length
      great_collatz_number = test_number;
      great_length = test_length;
    }
  }

  // return the number
  return great_collatz_number;
}

unsigned int getCollatzSequenceLength(const unsigned int start_val) {
  // ensure the number is greater than 1
  assert(start_val > 0);

  // variables
  unsigned int sequence_length = 1;
  unsigned int intermediate = start_val;

  // work out the sequence until it reaches 1
  while (intermediate > 1) {
    // count the sequence term
    sequence_length++;

    // case: the number is even
    if ((intermediate & 0x01) == 0) {
      // apply the transformation
      intermediate /= 2;
    }
    // case: the number is odd
    else {
      // apply the transformation
      intermediate *= 3;
      intermediate++;
    }
  }

  // return the sequence length
  return sequence_length;
}

