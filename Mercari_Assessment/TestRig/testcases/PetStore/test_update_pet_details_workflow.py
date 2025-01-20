import sys
# Adding the path of our Automation framework to the known paths where Python would be searching for modules
sys.path.append("/Users/sarthaksrivastava/VisualStudioProjects/sarthaktokointern/Mercari_Assessment/TestRig")

from workflows.PetStore.testcases.update_pet_details_workflow import UpdatePetDetailsWorkflow

class TestUpdatePetDetailsWorkflow(UpdatePetDetailsWorkflow):
    """
    testcase to trigger UpdatePetDetailsWorkflow
    """
    def __init__(self):
        """
        constructor for TestUpdatePetDetailsWorkflow
        Test Steps :-
            1. Updates pet information for pets named “kurikuri” under category “Pomeranian” to add the tag “Super Cute”
            2. Does sanity check by retreiving its details
        """
        # As of now hard coding the params required for running this test scenario. 
        # Future scope would be to make it config oriented for scalability
        super(TestUpdatePetDetailsWorkflow, self).__init__(pet_name="“kurikuri”", category_name="Pomeranian", tag={"id": 3176, "name": "Super Cute"})

test_update_pet_details_wf = TestUpdatePetDetailsWorkflow()
