#include <cassert>
#include <iostream>
#include <stack>
#include <cmath>
using namespace std;



bool isPalindrome(const int number);
int reverseNumber(const int number);
bool has3DigitFactors(const int number);
bool is3Digit(const int number);


int main() {
  // variables
  int number = 999 * 999; // largest number that is a product of two 3-digit #
  bool keep_searching = true;

  // check for the first product to be a palindrome that 
  while (keep_searching) {
    // case: the number is a palindrome
    if (isPalindrome(number)) {

cout << number << endl;

      // case: the number has 3 digit factors
      if (has3DigitFactors(number)) {
        // halt the search
        keep_searching = false;
      }
      // case: the number does not have 3 digit factors
      else {
        // continue the search
        number--;
      }
    }
    // case: the number did not meet the criteria
    else {
      // decrease the number and try again
      number--;
    }
  }

  if (number < (100 * 100)) {
    cout << "FAIL!" << endl;
  }
  else {
    cout << number << endl;
  }

  // return 0 for successful completion
  return 0;
}

bool isPalindrome(const int number) {
  // variables
  bool result = false;
  int reversed_number = 0;

  // get the reversed version of the number
  reversed_number = reverseNumber(number);

  // case: the number and its reverse are the same
  if (number == reversed_number) {
    // indicate that the number is a palindrome
    result = true;
  }

  // return the result
  return result;
}

int reverseNumber(const int number) {
  // variables
  int reversed_number = 0;
  int temp = number;
  int base_10_place = 1;
  stack<int> digits;

  // start by stripping the digits and placing them in a stack
  while (temp > 0) {
    // grab the lowest order digit and add it to the stack
    digits.push(temp % 10);

    // discard the digit from the temporary value
    temp /= 10;
  }

  // concatenate the digits in reverse to create the reversed number
  while (!digits.empty()) {
    // grab the digit off the top of the stack, and add it to the number in
    // it's appropriate power of 10 place
    reversed_number += (digits.top() * base_10_place);

    // move on to the next digit and the next place
    digits.pop();
    base_10_place *= 10;
  }

  // return the reversed number
  return reversed_number;
}

bool has3DigitFactors(const int number) {
  // variables
  bool result = false;
  bool keep_checking = true;
  int factor = ceil(sqrt(number));

cout << "Checking " << number << " for 3 digit factors" << endl;   

  // keep looking until it is determined that the number either has no
  // 3 digit factors or a pair has been found
  while (keep_checking) {
    // case: the number being checked is actually a factor
    if ((number % factor) == 0) {
      // case: the number is a 3 digit number
      if (is3Digit(factor)) {
        // case: the counterpart factor is 3 digits also
        if (is3Digit(number / factor)) {
          // halt the search, signal that the number has 3 digit factors
          keep_checking = false;
          result = true;
        }
        // case: both numbers aren't factors
        else {
          // continue the search
          factor--;
        }
      }
      // case: the number is not a 3 digit
      else {
        // we have searched all 3 digit possibilities, halt the search
        keep_checking = false;
      }
    }
    // case: the number was not a factor
    else {
      // move on to the next possibility
      factor--;

      // case: factor is no longer 3 digit
      if (!is3Digit(factor)) {
        // halt the search
        keep_checking = false;
      }
    }
  }

cout << "Finished checking " << number << " for 3 digit factors" << endl;  

  // return the result
  return result;
}

bool is3Digit(const int number) {
  // return the truth of the number being 3 digit
  return ((100 <= number) && (number <= 999));
}
