/**
    @file project_euler_problem_17.cpp

    @author Terence Henriod

    Project Euler Problem 17

    @brief Solves the following problem for the general case of any range of
           numbers ([1, 999,999]):

  "If the numbers 1 to 5 are written out in words: one, two, three, four, five,
   then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

   If all the numbers from 1 to 1000 (one thousand) inclusive were written out
   in words, how many letters would be used?

   NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and
         forty-two) contains 23 letters and 115 (one hundred and fifteen)
         contains 20 letters. The use of "and" when writing out numbers is in
         compliance with British usage."

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
const char kDecimalDigitStrings[20][10] =
    {"zero", "one", "two", "three", "four", "five", "six", "seven", "eight",
     "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen",
     "sixteen", "seventeen", "eighteen", "nineteen"};

const char kTenMultipleStrings[10][10] =
    {"", "ten", "twenty", "thirty", "forty", "fifty", "sixty",
     "seventy", "eighty", "ninety"};

const char kPowerOfTenStrings[4][15] =
    {"", "thousand", "million", "hundred"};

const int kDecimalDigitStringLengths [] =
    {strlen(kDecimalDigitStrings[0]), strlen(kDecimalDigitStrings[1]),
     strlen(kDecimalDigitStrings[2]), strlen(kDecimalDigitStrings[3]),
     strlen(kDecimalDigitStrings[4]), strlen(kDecimalDigitStrings[5]),
     strlen(kDecimalDigitStrings[6]), strlen(kDecimalDigitStrings[7]),
     strlen(kDecimalDigitStrings[8]), strlen(kDecimalDigitStrings[9]),
     strlen(kDecimalDigitStrings[10]), strlen(kDecimalDigitStrings[11]),
     strlen(kDecimalDigitStrings[12]), strlen(kDecimalDigitStrings[13]),
     strlen(kDecimalDigitStrings[14]), strlen(kDecimalDigitStrings[15]),
     strlen(kDecimalDigitStrings[16]), strlen(kDecimalDigitStrings[17]),
     strlen(kDecimalDigitStrings[18]), strlen(kDecimalDigitStrings[19])};

const int kTenMultipleStringLengths [] =
    {strlen(kTenMultipleStrings[0]), strlen(kTenMultipleStrings[1]),
     strlen(kTenMultipleStrings[2]), strlen(kTenMultipleStrings[3]),
     strlen(kTenMultipleStrings[4]), strlen(kTenMultipleStrings[5]),
     strlen(kTenMultipleStrings[6]), strlen(kTenMultipleStrings[7]),
     strlen(kTenMultipleStrings[8]), strlen(kTenMultipleStrings[9])};

const int kPowerOfTenStringLengths [] =
    {strlen(kPowerOfTenStrings[0]), strlen(kPowerOfTenStrings[1]),
     strlen(kPowerOfTenStrings[2])};

const int kAnd = 3;


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
unsigned int sumNumberLengths(const int range_start, const int range_end);
unsigned int getNumberLengthInLetters(const int number);


/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   MAIN FUNCTION
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
int main(/*int argc, char** argv*/) {
  // variables
  unsigned int number_string_length_total = 0;
  unsigned int range_start = 115;
  unsigned int range_end = 115;

  // process the command line arguments
  

  // compute the sum of all the number lengths 
  number_string_length_total = sumNumberLengths(range_start, range_end);

  // report the result
  cout << number_string_length_total << endl;

  // return 0 on successful completion
  return 0;
}


/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   FUNCTION IMPLEMENTATIONS
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
unsigned int sumNumberLengths(const int range_start, const int range_end) {
  // variables
  unsigned int length_sum = 0;
  int temp;

  // process every number in the range
  for (temp = range_start; temp <= range_end; temp++) {
    // add the length of the number as a string (not counting spaces or
    // hyphens) to the sum
    length_sum += getNumberLengthInLetters(temp);
  }

  // return the sum of the lengths of the numbers in the range
  return length_sum;
}

unsigned int getNumberLengthInLetters(const int number) {
    int numberLength = 0;
    int workingNumber = number;

    while (workingNumber > 0) {
        int numberSegment = workingNumber %1000;
        workingNumber %= 1000;

        int doubleDigits = numberSegment %100;
        if (doubleDigits > 0) {
            if (doubleDigits > 9 && doubleDigits < 20) {

            }
        }
    }



  //// variables
  //unsigned int number_string_length = 0;
  //int remaining_digits = number;
  //int segment = 0;
  //int ten_to_3rd_count = 0;

  //// strip off 3 digit segments at a time until there is nothing left
  //while (remaining_digits > 0) {
  //  // strip the least significant 3 digits
  //  segment = remaining_digits % 1000;
  //  remaining_digits /= 1000;

  //  // case: there is a mulitple of 100 in the digits
  //  if (segment > 99) {
  //    // add the number of hundreds
  //    number_string_length += kDecimalDigitStringLengths[segment / 100];

  //    // add "hundred"
  //    number_string_length += kPowerOfTenStringLengths[3];

  //    // keep the lower two digits
  //    segment %= 100;

  //    // case: the segment was not a multiple of 100
  //    if (segment > 0) {
  //      // include the "and"
  //      number_string_length += kAnd; // length of "and"
  //    }
  //  }

  //  // case: the number segment is still greater than zero
  //  if (segment > 0) {
  //    // case: the number is a 'teen
  //    if (segment < 20) {
  //      // add the length of that number to the sum
  //      number_string_length += kDecimalDigitStringLengths[segment];
  //    }
  //    // case: the number is not a 'teen
  //    else {
  //      // add the length of the 10s place digit
  //      number_string_length += kTenMultipleStringLengths[segment / 10];

  //      // add the length of the 1s place digit
  //      number_string_length += kDecimalDigitStringLengths[segment % 10];
  //    }
  //  }

  //  // add the proper "power of ten"
  //  number_string_length += kPowerOfTenStringLengths[ten_to_3rd_count];

  //  // move up to the next 10^3
  //  ten_to_3rd_count++;
  //}

  //// return the length of the number in letters
  //return number_string_length;
}

