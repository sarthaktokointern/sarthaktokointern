import sys
sys.path.append("/Users/sarthaksrivastava/VisualStudioProjects/sarthaktokointern/Mercari_Assessment/TestRig")

from workflows.PetStore.testcases.delete_pet_workflow import DeletePetWorkflow

class TestDeletePetWorkflow(DeletePetWorkflow):
    """
    testcase to trigger DeletePetWorkflow
    """
    def __init__(self):
        """
        constructor for TestDeletePetWorkflow
        Test Steps :-
            1. deletes pet from the store 
            2. Does sanity check by retreiving its details
        """
        # As of now hard coding the params required for running this test scenario. 
        # Future scope would be to make it config oriented for scalability
        super(TestDeletePetWorkflow, self).__init__(api_key="test_sarthak")

test_delete_pet_workflow = TestDeletePetWorkflow()
