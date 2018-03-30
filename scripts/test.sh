go test -v ../... \
| sed ''/RUN/s//$(printf "\033[36mRUN\033[0m")/'' \
| sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' \
| sed ''/FAIL/s//$(printf "\033[1;31mFAIL\033[0m")/''
