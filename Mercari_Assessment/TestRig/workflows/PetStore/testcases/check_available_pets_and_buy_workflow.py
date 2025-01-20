from workflows.PetStore.framework.PetStoreApi import PetStoreApiClient

class CheckAvailablePetsAndBuyWorkflow(PetStoreApiClient):
    """
    workflow for checking available pets and buying 
    """
    def __init__(self, **kwargs):
        """
        constructor for CheckAvailablePetsAndBuyWorkflow
        """
        super(CheckAvailablePetsAndBuyWorkflow, self).__init__(**kwargs)
        self.pet_name = kwargs.get("pet_name")
        self.category_name = kwargs.get("category_name")
        self.pet_quant = kwargs.get("pet_quantity")
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
                # if the desired pet is found, append its id to pet_id_list
                if petname == self.pet_name and petcategory == self.category_name:
                    pet_id_list.append(pet_info["id"])
        if not len(pet_id_list):
            raise Exception("\nDesired pets with name: {} and category: {} not available\n".format(self.pet_name, self.category_name))
        else:
            if self.pet_quant > len(pet_id_list):
                raise Exception("\nDesired pet doesn't have the required pet quant: {}. Pet count available: {}\n".format(self.pet_quant, len(pet_id_list)))
            else:
                pet_id_list = pet_id_list[:self.pet_quant]
                print("Pet ID List: {}".format(pet_id_list))
                for pet_id in pet_id_list:
                    order_details = self.order_pet(pet_id=pet_id)
                    if order_details:
                        print("\nOrder placed successfully!! Order Details: {}\n".format(order_details))
