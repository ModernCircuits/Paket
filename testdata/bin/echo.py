#!/usr/bin/python

import sys

args = ""
for arg in sys.argv[1:]:
    args += f"{arg} "

print(args)
