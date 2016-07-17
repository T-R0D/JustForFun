/**
    @file project_euler_problem_8.cpp

    @author Terence Henriod

    Project Euler Problem 8

    @brief Solves the following problem for the general case of any required
           number of primes (withing system limitations):

  "Find the greatest product of five consecutive digits in the 1000-digit number.

     73167176531330624919225119674426574742355349194934
     96983520312774506326239578318016984801869478851843
     85861560789112949495459501737958331952853208805511
     12540698747158523863050715693290963295227443043557
     66896648950445244523161731856403098711121722383113
     62229893423380308135336276614282806444486645238749
     30358907296290491560440772390713810515859307960866
     70172427121883998797908792274921901699720888093776
     65727333001053367881220235421809751254540594752243
     52584907711670556013604839586446706324415722155397
     53697817977846174064955149290862569321978468622482
     83972241375657056057490261407972968652414535100474
     82166370484403199890008895243450658541227588666881
     16427171479924442928230863465674813919123162824586
     17866458359124566529476545682848912883142607690042
     24219022671055626321111109370544217506941658960408
     07198403850962455444362981230987879927244284909188
     84580156166097919133875499200524063689912560717606
     05886116467109405077541002256983155200055935729725
     71636269561882670428252483600823257530420752963450"

    @version Original Code 1.00 (12/29/2013) - T. Henriod
*/

/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   HEADER FILES / NAMESPACES
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
#include <cassert>
#include <cmath>
#include <cstdio>
#include <list>
using namespace std;


/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   GLOBAL CONSTANTS
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
const char kBigNumber[] = "73167176531330624919225119674426574742355349194934"
                          "96983520312774506326239578318016984801869478851843"
                          "85861560789112949495459501737958331952853208805511"
                          "12540698747158523863050715693290963295227443043557"
                          "66896648950445244523161731856403098711121722383113"
                          "62229893423380308135336276614282806444486645238749"
                          "30358907296290491560440772390713810515859307960866"
                          "70172427121883998797908792274921901699720888093776"
                          "65727333001053367881220235421809751254540594752243"
                          "52584907711670556013604839586446706324415722155397"
                          "53697817977846174064955149290862569321978468622482"
                          "83972241375657056057490261407972968652414535100474"
                          "82166370484403199890008895243450658541227588666881"
                          "16427171479924442928230863465674813919123162824586"
                          "17866458359124566529476545682848912883142607690042"
                          "24219022671055626321111109370544217506941658960408"
                          "07198403850962455444362981230987879927244284909188"
                          "84580156166097919133875499200524063689912560717606"
                          "05886116467109405077541002256983155200055935729725"
                          "71636269561882670428252483600823257530420752963450";
const int kSubsequenceSize = 5;

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
int findLargestSubsequenceProduct(const char* number, const int number_length,
                                  const int subsequence_size,
                                  char* subsequence);
int getCharArrayProduct(const char* array, const int num_elements);


/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   MAIN FUNCTION
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
int main() {
  // variables
  char num_in_file = 'n';
  const char* number = kBigNumber;
  char subsequence[kSubsequenceSize + 1]; // +1 for '\0'
  int number_len = strlen(kBigNumber);
  int subsequence_product = 0;

  // get the nth prime the user wants to know
  printf("Is the number in a file?: ");
  scanf("%c", &num_in_file);
  printf("\n");

num_in_file = 'n';

  // case: the user specified the number is in a file
  if (num_in_file == 'y') {
    // prompt for the file name

    // attempt to read in the number

    // case: the file reading failed

      // notify the user of the failure
  }

  // get the product of the largest subsequence
  subsequence_product = findLargestSubsequenceProduct(kBigNumber, number_len,
                                                      kSubsequenceSize,
                                                      subsequence);

  // state the result for the user
  printf("The sought subsequence is: %s\n", subsequence);
  printf("The product of the sequence is: %u\n", subsequence_product);
  // return 0 on successful completion
  return 0;
}


/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   FUNCTION IMPLEMENTATIONS
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
int findLargestSubsequenceProduct(const char* number, const int number_length,
                                  const int subsequence_size,
                                  char* subsequence) {
  // variables
  int number_ndx = 0;
  int stop_ndx = number_length - subsequence_size;
  int subsequence_ndx = 0;
  int product = 0;
  int test_product = 0;
  const char* test_subsequence = number;

  // check all the subsequences for their product
  while (number_ndx < stop_ndx) {
    // get the current subsequence's product
    test_product = getCharArrayProduct(test_subsequence, subsequence_size);

    // case: the test product is larger than the current largest
    if (product < test_product) {
      // store the new high product
      product = test_product;

      // gather the subsequence
      for (subsequence_ndx = 0;
           subsequence_ndx < subsequence_size;
           subsequence_ndx++) {
        // copy the character over
        subsequence[subsequence_ndx] = test_subsequence[subsequence_ndx];
      }

      // append the null terminator
      subsequence[subsequence_ndx] = '\0';
    }

    // move to the next subsequence
    number_ndx++;
    test_subsequence++;
  }

  // return the resulting product
  return product;
}

int getCharArrayProduct(const char* array, const int num_elements) {
  // variables
  int product = 0;
  int ndx = 0;

  // iterate across all elements of the subsequence
  for (ndx = 0, product = 1; ndx < num_elements; ndx++) {
    // include each converted term in the product
    product *= (array[ndx] - '0');
  }

  // return the product
  return product;
}

