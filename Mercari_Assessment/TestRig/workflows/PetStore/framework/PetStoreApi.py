import requests
import copy

class PetStoreApiClient:
    """
    Initializes the PetStore API Client
    """
    def __init__(self, **kwargs):
        """
        Constructor for PetStoreApiClient
        Args:
            kwargs(dict):
                comm_proto(string): specifies type of communication protocol i.e. http/https
                api_key(string): for auth
        Returns:
            None
        """
        self.comm_proto = kwargs.get("comm_proto", "https")
        self.api_key = kwargs.get("api_key")
        self.base_url = "petstore.swagger.io/v2"
        self.headers = self._get_headers()
        print("Headers: {}".format(self.headers))

    def _get_headers(self):
        """
        sets the appropriate headers
        Returns:
            headers(dict): headers for the request
        """
        headers = {}
        headers["accept"] = "application/json"
        if self.api_key:
            headers["api_key"] = self.api_key
        return headers

    def fetch_pets_by_status(self, pet_status="available"):
        """
        fetches pets by status
        Args:
            pet_status(string): available/pending/sold
        Returns:
            pet_list(list): (dict): pet info
        """
        endpoint = f"pet/findByStatus"
        url = f"{self.comm_proto}://{self.base_url}/{endpoint}"
        params = {
            "status": pet_status
        }
        pet_list = requests.get(url, headers=self.headers, params=params)
        if pet_list.status_code == 200:
            return pet_list.json()
        else:
            raise Exception("\nError in fetching pets by status: {}. Error code: {}. Error message: {}\n".format(pet_status, pet_list.status_code, pet_list.text))

        return []
    
    def order_pet(self, pet_id, pet_quantity=1):
        """
        places an order of the pet
        Args:
            pet_id(int): id of the pet which needs to be bought
            pet_quantity(int): quantity of the pet
        Returns:
            order_details(dict): order info
        """
        endpoint = f"store/order"
        url = f"{self.comm_proto}://{self.base_url}/{endpoint}"
        params = {
            "petId": pet_id,
            "quantity": pet_quantity
        }
        headers = copy.deepcopy(self.headers)
        headers["Content-Type"] = "application/json"
        order_details = requests.post(url, headers=headers, params=params)
        if order_details.status_code == 200:
            return order_details.json()
        else:
            raise Exception("\nError in placing order for pet: {}. Error code: {}. Error message: {}\n".format(pet_id, order_details.status_code, order_details.text))

        return {}
    
    def add_pet_to_store(self, pet_id=0, category_id=0, category_name="", pet_name="", photo_urls=[], tags=[]):
        """
        places an order of the pet
        Args:
            pet_id(int): id of the pet which needs to be added
            category_id(int): category id
            category_name(str): category name
            pet_name(str): pet name
            photo_urls(list(str)): list of pet photo urls
            tags(list(dict)): tags associated with pet
        Returns:
            pet_details(dict): pet info
        """
        endpoint = f"pet"
        url = f"{self.comm_proto}://{self.base_url}/{endpoint}"
        body = {
            "id": pet_id,
            "category": {
                "id": category_id,
                "name": category_name
            },
            "name": pet_name,
            "photoUrls": photo_urls,
            "tags": tags,
            "status": "available"
        }
        headers = copy.deepcopy(self.headers)
        headers["Content-Type"] = "application/json"
        pet_details = requests.post(url, headers=headers, json=body)
        if pet_details.status_code == 200:
            return pet_details.json()
        else:
            raise Exception("\nError in adding new pet with details: {}. Error code: {}. Error message: {}\n".format(body, pet_details.status_code, pet_details.text))

        return {}
    
    def get_pet_by_id(self, pet_id):
        """
        fetches pet details by id
        Args:
            pet_id(int): pet id
        Returns:
            pet_details(dict): pet info
        """
        endpoint = f"pet/{pet_id}"
        url = f"{self.comm_proto}://{self.base_url}/{endpoint}"
        pet_details = requests.get(url, headers=self.headers)
        if pet_details.status_code == 200:
            return pet_details.json()
        else:
            raise Exception("\nError in fetching pet by id: {}. Error code: {}. Error message: {}\n".format(pet_id, pet_details.status_code, pet_details.text))

        return {}
    
    def update_pet_details(self, body):
        """
        updates pet details
        Args:
            body(dict): updated pet details
        Returns:
            updated_pet_info(dict): updated pet details
        """
        endpoint = f"pet"
        url = f"{self.comm_proto}://{self.base_url}/{endpoint}"
        headers = copy.deepcopy(self.headers)
        headers["Content-Type"] = "application/json"
        updated_pet_details = requests.put(url, headers=headers, json=body)
        if updated_pet_details.status_code == 200:
            return updated_pet_details.json()
        else:
            raise Exception("\nError in updating pet with details: {}. Error code: {}. Error message: {}\n".format(body, updated_pet_details.status_code, updated_pet_details.text))

        return {}
    
    def delete_pet(self, pet_id):
        """
        deletes pet from the store
        Args:
            pet_id(int): id of pet
        Returns:
            None
        """
        endpoint = f"pet/{pet_id}"
        url = f"{self.comm_proto}://{self.base_url}/{endpoint}"
        pet_del_request = requests.delete(url, headers=headers)
        if pet_del_request.status_code == 400 or pet_del_request.status_code == 404:
            raise Exception("\nError in deleting pet with id: {}. Error code: {}. Error message: {}\n".format(pet_id, pet_del_request.status_code, pet_del_request.text))
        