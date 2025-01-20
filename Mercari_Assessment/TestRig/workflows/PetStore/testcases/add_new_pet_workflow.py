from workflows.PetStore.framework.PetStoreApi import PetStoreApiClient

class AddNewPetWorkflow(PetStoreApiClient):
    """
    workflow for adding new pet and fetching details 
    """
    def __init__(self, **kwargs):
        """
        constructor for AddNewPetWorkflow
        """
        super(AddNewPetWorkflow, self).__init__(**kwargs)
        self.kwargs = kwargs
        # prepares the test steps for the aforementioned test scenario
        self.prepare_test_steps()
    def prepare_test_steps(self):
        """
        performs the test scenario
        Returns:
            None
        """
        pet_details = self.add_pet_to_store(pet_id=self.kwargs.get("pet_id"), pet_name=self.kwargs.get("pet_name"))
        # fetch all the list of available pets
        pet_list = self.fetch_pets_by_status()
        if pet_list:
            for pet_info in pet_list:
                print("Pet Info: {}".format(pet_info))
                petname = None
                petid = None
                if "name" in pet_info:
                    petname = pet_info["name"]
                if "id" in pet_info:
                    petid = pet_info["id"]
                # looking for the newly added pet
                if petname == self.kwargs.get("pet_name") and petid == self.kwargs.get("pet_id"):
                    print("Successfully Added Pet with following details: {}".format(pet_details))
                    return
    
        raise Exception(f"\nPet with id: {self.kwargs.get("pet_id")} has not been added!!!\n")
