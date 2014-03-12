#!/usr/bin/env python

def usage(script_name):
  print('Usage: {} <code>'.format(script_name))

prefixes = ['p', 'n', 'u', 'm', '']

import math

def getPrefix(value):
  print(math.log10(value))
  return prefixes[int(math.log10(value)) / 3]

def simplify(value):
  if round(value, 14) == value:
    return round(value, 14)
  return value

def main(script_name, code=None):
  if code is None:
    usage(script_name)
    return 1

  base = int(code[:2]) /10
  exponent = int(code[2:] or 0)+1
  prefix = prefixes[exponent//3]

  value = simplify(base * 10**(exponent % 3))

  print('{} {}F'.format(value, prefix))

  if value < 100:
    return 0
  value /= 1000
  prefix = prefixes[exponent//3 + 1]

  print('{} {}F'.format(value, prefix))

import sys

if __name__ == '__main__':
  sys.exit(main(*sys.argv))
