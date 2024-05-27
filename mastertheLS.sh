#!/bin/bash


ls -AtmDp -I . -I .. | tr -d '[:space:]' | sed ':a;N;$!ba;s/\n/,/g'
