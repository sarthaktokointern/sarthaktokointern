Git Internals Part 3 : Merge

1.) Fast Forward Merge : by default this merge happens. Here is an example for the same :

a.) Initially master is pointing to C2 :

<img width="363" alt="Screenshot 2022-08-12 at 4 25 55 PM" src="https://user-images.githubusercontent.com/99721005/184340558-8df1c90a-cfeb-4332-988a-40e16576d6b8.png">

b.) check out a new branch and add C3 and C4 commits :

<img width="649" alt="Screenshot 2022-08-12 at 4 26 12 PM" src="https://user-images.githubusercontent.com/99721005/184340605-32ececce-a665-4992-97f4-b1b70ae5fbf1.png">


c.) Merge master and new branch :

<img width="683" alt="Screenshot 2022-08-12 at 4 26 30 PM" src="https://user-images.githubusercontent.com/99721005/184340649-044281fb-9b74-4ffc-91d0-e6d64ef6743e.png">



-> This is the simplest form of merge and no new objects are created in this merge, only the branch references are updated. But in case of diverged branches, fast-forward merge creates a new commit known as merge commit. The following diagram will help in understanding this part better :

a.) Initially master is pointing to C2 commit :

<img width="349" alt="Screenshot 2022-08-12 at 4 26 59 PM" src="https://user-images.githubusercontent.com/99721005/184340716-e847f7d0-59c5-40eb-80ff-cfa8893245e1.png">


b.) Checkout  new branch and add C3 commit :

<img width="485" alt="Screenshot 2022-08-12 at 4 27 12 PM" src="https://user-images.githubusercontent.com/99721005/184340750-8e3d5fe8-e030-4adc-8b67-5c771604f395.png">


c.) Checkout back to master and add C4 commit :


<img width="505" alt="Screenshot 2022-08-12 at 4 27 29 PM" src="https://user-images.githubusercontent.com/99721005/184340791-954e521d-2633-45da-b997-1fa5675bd43c.png">

after merging master and new branch :

<img width="732" alt="Screenshot 2022-08-12 at 4 28 12 PM" src="https://user-images.githubusercontent.com/99721005/184340896-19bb01e8-fd22-4ae2-8f3c-637891fe92f6.png">


2.) Non fast forward merge : if the branches are not diverged, but we want to perform non fast-forward merge we can use the flag '-no-ff'. In this a new merge commit will be created and it will look something like this :

<img width="807" alt="Screenshot 2022-08-12 at 4 28 29 PM" src="https://user-images.githubusercontent.com/99721005/184340943-bb821896-c706-4b94-a3df-9ae325939b4e.png">


