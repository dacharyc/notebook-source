#! /bin/bash

PROJECT=$(git rev-parse --show-toplevel)
JS_EXAMPLES=$PROJECT/javascript/examples
GENERATED_EXAMPLES=$PROJECT/generated/javascrpt

echo "Bluehawking JavaScript examples"
npx bluehawk snip $JS_EXAMPLES -o $GENERATED_EXAMPLES
