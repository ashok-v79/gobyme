**_ Assert vs Require _**

assert, is used to validate conditions without stopping the test execution, allowing multiple failures to be captured in a single run.

Using require ensures that tests stop immediately if fundamental conditions are not met, preventing the execution of meaningless test logic.

**\* Notes **

In the "Valid Data" test, require assertions are used to ensure the data is valid. If these fail, subsequent steps should not execute, as indicated by the print statements. assert is then used to check processing and saving, with the expectation that the test will continue even if these assertions fail.

In the "Invalid Input" test, require is used to confirm that validation fails for empty input. If validation does not fail as expected, the test stops immediately after printing the relevant message.

In the "Invalid Processing" test, require is used to check the initial validation, and assert is used to check that processing fails as expected. The print statements help trace the flow of execution.
