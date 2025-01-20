import sys
# Adding the path of our Automation framework to the known paths where Python would be searching for modules
sys.path.append("/Users/sarthaksrivastava/VisualStudioProjects/sarthaktokointern/Mercari_Assessment/TestRig")

from workflows.PetStore.testcases.check_available_pets_and_buy_workflow import CheckAvailablePetsAndBuyWorkflow

class TestCheckAvailablePetsAndBuyWorkflow(CheckAvailablePetsAndBuyWorkflow):
    """
    testcase to trigger CheckAvailablePetsAndBuyWorkflow
    """
    def __init__(self):
        """
        constructor for TestCheckAvailablePetsAndBuyWorkflow
        Test Steps :-
            1. Check available pets with name pupo and category pajaro
            2. Place order with the desired quantity
        """
        # As of now hard coding the params required for running this test scenario. 
        # Future scope would be to make it config oriented/driven for scalability
        super(TestCheckAvailablePetsAndBuyWorkflow, self).__init__(pet_name="pupo", category_name="pajaro", pet_quantity=22)

test_check_available_pets_and_buy_workflow = TestCheckAvailablePetsAndBuyWorkflow()
