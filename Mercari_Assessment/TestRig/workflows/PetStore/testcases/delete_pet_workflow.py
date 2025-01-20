from workflows.PetStore.framework.PetStoreApi import PetStoreApiClient

class DeletePetWorkflow(PetStoreApiClient):
    """
    workflow for deleting pet 
    """
    def __init__(self, **kwargs):
        """
        constructor for DeletePetWorkflow
        """
        super(DeletePetWorkflow, self).__init__(**kwargs)
        self.kwargs = kwargs
        # prepares the test steps for the aforementioned test scenario
        self.prepare_test_steps()
    def prepare_test_steps(self):
        """
        performs the test scenario
        Returns:
            None
        """
        pet_id = None
        # fetch all the list of available pets
        pet_list = self.fetch_pets_by_status()
        if pet_list:
            for pet_info in pet_list:
                self.delete_pet(pet_id=pet_info["id"])
                pet_id = pet_info["id"]
                break
            pet_list = self.fetch_pets_by_status()
            for pet_info in pet_list:
                if pet_id == pet_info["id"]:
                    raise Exception("\nPet with id: {} still present\n".format(pet_id))
        else:
            raise Exception(f"\nNo pets available to delete\n")
