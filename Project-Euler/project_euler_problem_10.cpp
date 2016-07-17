/**
    @file project_euler_problem_10.cpp

    @author Terence Henriod

    Project Euler Problem 10

    @brief Solves the following problem for the general case of any required
           number of primes (withing system limitations):

  "The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

   Find the sum of all the primes below two million."

    @version Original Code 1.00 (12/31/2013) - T. Henriod
*/

/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   HEADER FILES / NAMESPACES
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
#include <cassert>
#include <cmath>
#include <cstdio>
#include <iostream>
#include <list>
using namespace std;


/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   GLOBAL CONSTANTS
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
const int kFirstPrime = 2;

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
long double findPrimeSumUnder(const unsigned int nth);
char isPrime(unsigned int number);

/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   MAIN FUNCTION
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
int main() {
  // variables
  unsigned int upper_bound = 2000000;
  long double prime_sum = 0;

  // get the nth prime the user wants to know
  printf("Enter the upper bound of primes to be summed: ");
  scanf("%u", &upper_bound);
  printf("\n");

  // find the sum of the first n primes
  prime_sum = findPrimeSumUnder(upper_bound);

  // state the result for the user
  printf("The sum of all primes below %u is: %12f\n", upper_bound, prime_sum);

  // return 0 on successful completion
  return 0;
}


/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   FUNCTION IMPLEMENTATIONS
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
long double findPrimeSumUnder(const unsigned int upper_bound) {
  // variables
  double prime_sum = 0;
  unsigned int temp = 3;
  unsigned int temp_root = ceil(sqrt(temp));
  unsigned int new_prime_ndx = 0;
  unsigned int found_prime_ndx = 0;
  unsigned int* primes = NULL;

  // allocate memory for the list
  primes = new unsigned int [upper_bound / 2];

  // prime the search with the first prime
  primes[0] = kFirstPrime;
  new_prime_ndx = 1;
  prime_sum += kFirstPrime;

  // successively check all numbers under the bound for prime-ness
  while (temp < upper_bound) {
    // check the next attempted number against the previously found primes
    while (((temp % primes[found_prime_ndx]) != 0) &&
           ((primes[found_prime_ndx] <= temp_root) &&
           (found_prime_ndx < new_prime_ndx))) {
      // move to the next found prime
      found_prime_ndx++;
    }

    // case: the temporary number passed the test
    if ((primes[found_prime_ndx] > temp_root) ||
        (found_prime_ndx >= new_prime_ndx)) {
      // store it
      primes[new_prime_ndx] = temp;
      new_prime_ndx++;

      // add it to the sum
      prime_sum += (long double)temp;

/*
char y;
if ((new_prime_ndx % 1000) == 0) {
printf("Just found prime %u and its test result is: %c\n", temp, isPrime(temp));
printf("The current sum is: %u\n", prime_sum);
scanf("%c", &y);
}
*/
    }

    // either way, move on to the next number
    temp++;
    temp_root = ceil(sqrt(temp));

    // start over at the beginning of the found primes list for testing
    found_prime_ndx = 0;
  }

  // return the dynamic memory
  delete [] primes;

  // return the nth prime that was found
  return prime_sum;
}

char isPrime(unsigned int number) {
  // variables
  char result = ' ';
  unsigned int i = kFirstPrime;
  unsigned int root = ceil(sqrt(number));

  // test all possible factors
  for (i = kFirstPrime; i <= root; i++) {
    if ((number % i) == 0) {
      result = '#';
    }
  }

  return result;
}