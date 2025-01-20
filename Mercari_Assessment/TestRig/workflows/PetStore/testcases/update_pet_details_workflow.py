from workflows.PetStore.framework.PetStoreApi import PetStoreApiClient

class UpdatePetDetailsWorkflow(PetStoreApiClient):
    """
    workflow for updating pet details 
    """
    def __init__(self, **kwargs):
        """
        constructor for UpdatePetDetailsWorkflow
        """
        super(UpdatePetDetailsWorkflow, self).__init__(**kwargs)
        self.kwargs = kwargs
        # prepares the test steps for the aforementioned test scenario
        self.prepare_test_steps()
    def prepare_test_steps(self):
        """
        performs the test scenario
        Returns:
            None
        """
        pet_id_list = []
        # fetch all the list of available pets
        pet_list = self.fetch_pets_by_status()
        if pet_list:
            for pet_info in pet_list:
                print("Pet Info: {}".format(pet_info))
                petname = None
                if "name" in pet_info:
                    petname = pet_info["name"]
                petcategory = None
                if "category" in pet_info and "name" in pet_info["category"]:
                    petcategory = pet_info["category"]["name"]
                # filter out desired pets
                if petname == self.kwargs.get("pet_name") and petcategory == self.kwargs.get("category_name"):
                    pet_id_list.append(pet_info["id"])
        pet_id_to_details_map = {}
        if not len(pet_id_list):
             raise Exception("\nDesired pets with name: {} and category: {} not available\n".format(self.kwargs.get("pet_name"), self.kwargs.get("category_name")))
        else:
            for pet_id in pet_id_list:
                # fetch the pet details
                pet_info = self.get_pet_by_id(pet_id=pet_id)
                # add the tag
                pet_info["tags"].append(self.kwargs.get("tag"))
                pet_id_to_details_map[pet_id] = pet_info
                # update the details
                self.update_pet_details(body=pet_info)
        
        pet_list = self.fetch_pets_by_status()
        for pet_info in pet_list:
            print("Pet Info: {}".format(pet_info))
            pet_id = pet_info["id"]
            if pet_id in pet_id_to_details_map and pet_id_to_details_map[pet_id] != pet_info:
                raise Exception("\nUpdate for pet with id: {} didn't happen. Expected pet info: {}. Current pet info: {}\n".format(pet_id, pet_id_to_details_map[pet_id], pet_info))
