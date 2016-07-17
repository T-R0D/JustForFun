/**
    @file project_euler_problem_7.cpp

    @author Terence Henriod

    Project Euler Problem 7

    @brief Solves the following problem for the general case of any required
           number of primes (withing system limitations):

  "By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see
   that the 6th prime is 13.

   What is the 10 001st prime number?"

    @version Original Code 1.00 (12/29/2013) - T. Henriod
*/

/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   HEADER FILES / NAMESPACES
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
#include <cassert>
#include <cmath>
#include <cstdio>
#include <fstream>
#include <list>
using namespace std;


/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   GLOBAL CONSTANTS
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
const int kFirstPrime = 2;
const int kMoreNumbers = 10;

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
long long findNthPrime(const unsigned int nth);
list<long long> generateListOfPrimes(const int num_primes);
void removeNonPrimes(list<long long>::iterator& old_list_end,
                     list<long long>& primes_list);

/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   MAIN FUNCTION
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
int main() {
  // variables
  unsigned int nth = 10001;
  long long nth_prime = 0;
  int counter = 1;
  list<long long> list_of_primes;
  list<long long>::iterator ith_prime;

  // get the nth prime the user wants to know
  printf("Enter the nth prime to be found: ");
  scanf("%u", &nth);
  printf("\n");

  // find the nth prime
  nth_prime = findNthPrime(nth);

/*
  // get a suitable list of primes
  list_of_primes = generateListOfPrimes(nth);

  // retrieve the nth prime from the list
  for (counter = 1, ith_prime = list_of_primes.begin(); counter < nth_prime;
       counter++, ++ith_prime) {
  }
*/


  // state the result for the user
  printf("The %uth prime number is: %d \n", nth, nth_prime);
  // return 0 on successful completion
  return 0;
}


/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   FUNCTION IMPLEMENTATIONS
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
long long findNthPrime(const unsigned int nth) {
  // variables
  long long nth_prime = 0;
  long long temp = 3;
  long long temp_root = ceil(sqrt(temp));
  long long* primes = NULL;
  int found_prime = 0;
  int last_found_prime = 0;

  // allocate memory for the list
  primes = new long long [nth];

  // prime the search with the first prime
  primes[0] = 2;
  last_found_prime = 1;

  // successively check numbers as they are generated for prime-ness
  while (last_found_prime < nth) {
    // check the next attempted number against the previously found primes
    while (((temp % primes[found_prime]) != 0) &&
           ((primes[found_prime] <= temp_root) ||
           (found_prime < last_found_prime))) {
      // move to the next found prime
      found_prime++;
    }

    // case: the temporary number passed the test
    if (found_prime >= last_found_prime) {
      // store it
      primes[last_found_prime] = temp;
      last_found_prime++;

if ((last_found_prime % 1000) == 0) {
printf("Another 1000th has appeared: %i!\n", temp);
}

    }

    // either way, move on to the next number
    temp++;
    temp_root = ceil(sqrt(temp));

    // start over at the beginning
    found_prime = 0;
  }

  // save the nth prime
  nth_prime = primes[nth - 1];

  // create a prime number file
  ofstream fout;
  fout.clear();
  fout.open("z_primes.h");
  fout << "#ifndef __Z_PRIMES_H__" << endl
       << "#define __Z_PRIMES_H__" << endl << endl
       << "const unsigned int kNumPrimes = " << nth << ';' << endl << endl
       << "const unsigned long long kPrimes[] = {" << endl;
  for (int i = 0; i < nth; ) {
    for (int j = 0; j < 5; i++, j++) {
      fout << primes[i] << ',';
    }
    fout << '\n';
  }
  fout << "\b\b\n};" << endl << endl;
  fout << "#endif" << endl;
  fout.close();

  // return the dynamic memory
  delete [] primes;

  // return the nth prime that was found
  return nth_prime;
}



list<long long> generateListOfPrimes(const int num_primes) {
  // variables
  list<long long> primes_list;
  list<long long>::iterator old_list_end;
  long long last_number_added = 2;
  long long new_last_number = 0;
  long long multiple_of_more_numbers = 1;

  // prime the list with the first prime number
  primes_list.push_back(kFirstPrime);

  // continue adding numbers to the list until the list is of requisite size
  while (primes_list.size() <= num_primes) {
    // track the end of the last set of progress
    old_list_end = primes_list.end();
    --old_list_end;

    // add more numbers to the list
    for(last_number_added = last_number_added + 1,
        new_last_number = multiple_of_more_numbers * kMoreNumbers;
        last_number_added <= new_last_number;
        last_number_added++) {
      // add the number to the end of the list
      primes_list.push_back(last_number_added);
    }

    // sift out the non-primes using a sieve method
    removeNonPrimes(old_list_end, primes_list);

    // do housekeeping for the next loop
    last_number_added = new_last_number;
    multiple_of_more_numbers++;
  }

  // return the generated list
  return primes_list;
}

void removeNonPrimes(list<long long>::iterator& old_list_end,
                     list<long long>& primes_list) {
  // variables
  list<long long>::iterator current_prime = primes_list.begin();
  list<long long>::iterator last_prime = old_list_end;

  // check all the newly added numbers against the previously found primes
  while (last_prime != primes_list.end()) {
    // start at the beginning of the list
    current_prime = primes_list.begin();

    // move on to the next possibile prime in the list
    ++last_prime;

    // check to see if each newly added element is prime
    while ((last_prime != primes_list.end()) &&
           (current_prime != last_prime)) {
      // case: the possibility is a multiple of one of the prime numbers
      if ((*last_prime % *current_prime) == 0) {
        // the number is not prime, throw it out
        last_prime = primes_list.erase(last_prime);

        // start at the beginning of the list
        current_prime = primes_list.begin();
      }
      // case: the possibility was not a multiple of a particular prime
      else {
         // try again with the next established prime
         current_prime++;
      }
    }
  }  

  // no return - void
}

