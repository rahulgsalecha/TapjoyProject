Tapjoy Project
=====================

Prerequisites:
==========================
1. Install Go 1.13 on Mac / Linux

Instructions:
==========================
1. Navigate to directory internal/pkgs
2. Run following command: 

    `" go run validator/validator.go data/input.txt "`  
    `" go run validator/validator.go data/sample.txt "`  
    
3. Check the output
4. To run unit test, run below commands to open coverage in html format:
   
   `" go test -v -coverpkg=./... -coverprofile=/tmp/profile.cov ./... "` 
   
   `" go tool cover -html=/tmp/profile.cov -o /tmp/c_report.html "` 
   
   `" open -a "Google Chrome" /tmp/c_report.html "` 

Output:
=======

1. go run validator/validator.go data/sample.txt

==== VALID LINES ====
rttr[mnop]qrst
irttrj[asdfgh]zxcvbns

==== TOTAL VALID LINES: 2 ====


2. go run validator/validator.go data/input.txt

==== VALID LINES ====
Kjhgfdsawrfv[adgnhytsgdavsa]asfjyrdfzbnnbtahjh]ebvsafdbhrfsda
poihkjnaviu[abwhergvbasiuhjgadbghwesfdvxcih[nwjkkjijkniubgqerqwuy]aehgkuyasgavs
uxpvoytxfazjjhi[qogwhtzmwxvjwxreuz]zduoybbzxiwwggwu[lamifchqqwbphhsqnf]qrjdjwtnhsjqftnq

==== TOTAL VALID LINES: 3 ====


Testing:
========

Run following command for testing:

"go test -v -coverpkg=./... -coverprofile=/tmp/profile.cov ./..."

pkgs $go test -v -coverpkg=./... -coverprofile=/tmp/profile.cov ./...
=== RUN   Test_checkExternalSubstringsValidity
=== RUN   Test_checkExternalSubstringsValidity/Test_checkExternalSubstringsValidity_true_case_1
=== RUN   Test_checkExternalSubstringsValidity/Test_checkExternalSubstringsValidity_true_case_2
=== RUN   Test_checkExternalSubstringsValidity/Test_checkExternalSubstringsValidity_false_case_1
--- PASS: Test_checkExternalSubstringsValidity (0.00s)
    --- PASS: Test_checkExternalSubstringsValidity/Test_checkExternalSubstringsValidity_true_case_1 (0.00s)
    --- PASS: Test_checkExternalSubstringsValidity/Test_checkExternalSubstringsValidity_true_case_2 (0.00s)
    --- PASS: Test_checkExternalSubstringsValidity/Test_checkExternalSubstringsValidity_false_case_1 (0.00s)
=== RUN   Test_checkInternalSubstringsValidity
=== RUN   Test_checkInternalSubstringsValidity/Test_checkInternalSubstringsValidity_true_case_1
=== RUN   Test_checkInternalSubstringsValidity/Test_checkInternalSubstringsValidity_false_case_1
=== RUN   Test_checkInternalSubstringsValidity/Test_checkInternalSubstringsValidity_false_case_2
--- PASS: Test_checkInternalSubstringsValidity (0.00s)
    --- PASS: Test_checkInternalSubstringsValidity/Test_checkInternalSubstringsValidity_true_case_1 (0.00s)
    --- PASS: Test_checkInternalSubstringsValidity/Test_checkInternalSubstringsValidity_false_case_1 (0.00s)
    --- PASS: Test_checkInternalSubstringsValidity/Test_checkInternalSubstringsValidity_false_case_2 (0.00s)
=== RUN   Test_isValidSubstring
=== RUN   Test_isValidSubstring/Test_isValidSubstring_case1_true
=== RUN   Test_isValidSubstring/Test_isValidSubstring_case2_false
=== RUN   Test_isValidSubstring/Test_isValidSubstring_case3_false
--- PASS: Test_isValidSubstring (0.00s)
    --- PASS: Test_isValidSubstring/Test_isValidSubstring_case1_true (0.00s)
    --- PASS: Test_isValidSubstring/Test_isValidSubstring_case2_false (0.00s)
    --- PASS: Test_isValidSubstring/Test_isValidSubstring_case3_false (0.00s)
=== RUN   Test_lineValidator
=== RUN   Test_lineValidator/Test_lineValidator_case1_true
=== RUN   Test_lineValidator/Test_lineValidator_case2_false
=== RUN   Test_lineValidator/Test_lineValidator_case3_false
=== RUN   Test_lineValidator/Test_lineValidator_case4_true
--- PASS: Test_lineValidator (0.00s)
    --- PASS: Test_lineValidator/Test_lineValidator_case1_true (0.00s)
    --- PASS: Test_lineValidator/Test_lineValidator_case2_false (0.00s)
    --- PASS: Test_lineValidator/Test_lineValidator_case3_false (0.00s)
    --- PASS: Test_lineValidator/Test_lineValidator_case4_true (0.00s)
=== RUN   Test_splitLineToSubstrings
=== RUN   Test_splitLineToSubstrings/Test_splitLineToSubstrings
--- PASS: Test_splitLineToSubstrings (0.00s)
    --- PASS: Test_splitLineToSubstrings/Test_splitLineToSubstrings (0.00s)
=== RUN   Test_countAndPrintValidLines
=== RUN   Test_countAndPrintValidLines/Test_countAndPrintValidLines_case1_validFile

==== VALID LINES ====
rttr[mnop]qrst
irttrj[asdfgh]zxcvbns

==== TOTAL VALID LINES: 2 ====
=== RUN   Test_countAndPrintValidLines/Test_countAndPrintValidLines_case2_validFile

==== VALID LINES ====
Kjhgfdsawrfv[adgnhytsgdavsa]asfjyrdfzbnnbtahjh]ebvsafdbhrfsda
poihkjnaviu[abwhergvbasiuhjgadbghwesfdvxcih[nwjkkjijkniubgqerqwuy]aehgkuyasgavs
uxpvoytxfazjjhi[qogwhtzmwxvjwxreuz]zduoybbzxiwwggwu[lamifchqqwbphhsqnf]qrjdjwtnhsjqftnq

==== TOTAL VALID LINES: 3 ====
=== RUN   Test_countAndPrintValidLines/Test_countAndPrintValidLines_case3_InvalidFile
--- PASS: Test_countAndPrintValidLines (0.00s)
    --- PASS: Test_countAndPrintValidLines/Test_countAndPrintValidLines_case1_validFile (0.00s)
    --- PASS: Test_countAndPrintValidLines/Test_countAndPrintValidLines_case2_validFile (0.00s)
    --- PASS: Test_countAndPrintValidLines/Test_countAndPrintValidLines_case3_InvalidFile (0.00s)
PASS
coverage: 85.7% of statements in ./...
ok  	TapjoyProject/internal/pkgs/validator	0.008s	coverage: 85.7% of statements in ./...
