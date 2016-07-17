#include <cassert>
#include <iostream>
using namespace std;

unsigned long sumEvenFibonacciNumbers(const unsigned long upper_bound);

int main() {
  // variables
  unsigned long upper_bound = 4000000;

  cout << sumEvenFibonacciNumbers(upper_bound)
       << endl;

  return 0;
}

unsigned long sumEvenFibonacciNumbers(const unsigned long upper_bound) {
  // variables
  unsigned long total = 0;
  unsigned long n_minus_2_term = 0;
  unsigned long n_minus_1_term = 1;
  unsigned long temp = 0;

  // continue generating terms until the upper bound is reached
  while (n_minus_1_term <= upper_bound) {
    // case: the newest term is an even one
    if ((n_minus_1_term % 2) == 0) {
      // add the term to the total
      total += n_minus_1_term;
    }

    // generate the next Fibonacci term
    temp = n_minus_1_term;
    n_minus_1_term = n_minus_1_term + n_minus_2_term;
    n_minus_2_term = temp;
  }

  // return the resulting sum
  return total;
}

