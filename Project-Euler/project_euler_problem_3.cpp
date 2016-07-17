#include <cassert>
#include <iostream>
using namespace std;

unsigned long long getLargestPrimeFactor(unsigned long long the_number);
unsigned long long getFirstFactor(unsigned long long the_number);

int main() {
  // variables
  unsigned long long the_number = 600851475143;

  cout << "Enter the number: ";
  cin >> the_number;
  cout << endl;

  cout << "The largest prime factor is: "
       << getLargestPrimeFactor(the_number)
       << endl;

  return 0;
}

unsigned long long getLargestPrimeFactor(unsigned long long the_number) {
  // variables
  unsigned long long largest_prime_factor = 0;
  unsigned long long next_large_factor = the_number;
  unsigned long long first_factor = 2;

  // perform prime factorization on the number
  while ( first_factor < next_large_factor) {
    // get the first factor of the largest (not yet confirmed prime)
    // factor found
    first_factor = getFirstFactor(next_large_factor);

cout << first_factor << ",\n";

    // the next large factor will be the complement of the first_factor
    next_large_factor = (next_large_factor / first_factor);

    // case: the recently found first factor is a larger prime factor
    if (first_factor > largest_prime_factor) {
      // save the newly found larger prime factor
      largest_prime_factor = first_factor;
    }
  }

  // assert post-condition
  assert(largest_prime_factor > 1);

  // return the largest prime factor found
  return largest_prime_factor;
}

unsigned long long getFirstFactor(unsigned long long the_number) {
  // variables
  unsigned long long candidate = 2;

  // count-up until the first (non-1) factor is found
  while ((the_number % candidate) != 0) {
    // move on to the next possibility
    candidate++;
  }

  // return the found factor
  return candidate;
}
