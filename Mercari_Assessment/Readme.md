TestRig is an automation framework. Its file structure is as follows:

testcases - for triggering various test scenarios via workflows

workflows - consists of workflow files having business logic for the test scenarios and their corresponding framework files
for instance, 

workflows -> PetStore -> framework : houses PetStore API

workflows -> PetStore -> testcases : for performing test scenarios which consume the aforementioned framework

General flow -> 

TestRig/testcases/PetStore/<test_case> -> TestRig/workflows/PetStore/testcases/<workflow> -> TestRig/workflows/PetStore/framework/PetStoreApi.py
