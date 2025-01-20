import sys
# Adding the path of our Automation framework to the known paths where Python would be searching for modules
sys.path.append("/Users/sarthaksrivastava/VisualStudioProjects/sarthaktokointern/Mercari_Assessment/TestRig")

from workflows.PetStore.testcases.add_new_pet_workflow import AddNewPetWorkflow

class TestAddNewPetWorkflow(AddNewPetWorkflow):
    """
    testcase to trigger AddNewPetWorkflow
    """
    def __init__(self):
        """
        constructor for TestAddNewPetWorkflow
        Test Steps :-
            1. Adds new pet to the store 
            2. Does sanity check by retreiving its details
        """
        # As of now hard coding the params required for running this test scenario. 
        # Future scope would be to make it config oriented for scalability
        super(TestAddNewPetWorkflow, self).__init__(pet_name="churro", pet_id=216378)

test_add_new_pet_workflow = TestAddNewPetWorkflow()
