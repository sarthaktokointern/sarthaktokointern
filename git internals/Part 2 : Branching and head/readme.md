Git Internals Part 2 : Branching and Head

1.) Branching : These are just references to commit. They are just like tags, but they are not considered as objects in git because objects are immutable and branches are mutable.

One can find branches in the following directory : .git/refs/heads

This directory will contain name of all the branches. 

Run the following command : cat .git/refs/heads/<branch_name> 

After running the above-mentioned command you will see the commit id to which that particular branch is pointing.

Now let's take a hypothetical scenario for branching :

1.) Create a git repo and create 2 commits C1 and C2 on branch master

<img width="351" alt="Screenshot 2022-08-12 at 11 17 27 AM" src="https://user-images.githubusercontent.com/99721005/184292443-b82c84c0-ca61-4351-9a3d-ace527929c13.png">



2.) Checkout a new branch test and create 2 commits C3 and C4

<img width="707" alt="Screenshot 2022-08-12 at 11 18 05 AM" src="https://user-images.githubusercontent.com/99721005/184292498-f2e55ccf-b82f-4443-98cd-7880dccc5e1e.png">



3.) Checkout master branch and create commit C5

<img width="738" alt="Screenshot 2022-08-12 at 11 18 21 AM" src="https://user-images.githubusercontent.com/99721005/184292524-7a4aa048-2cdb-48a5-8565-016954455f40.png">


4.) Create a new branch test1 and create commit C6

<img width="661" alt="Screenshot 2022-08-12 at 11 18 47 AM" src="https://user-images.githubusercontent.com/99721005/184292566-0aa5da00-c902-4fc9-b096-48749ddc204b.png">

2.) HEAD : is the current version of the code or pointer to the commit recently checked out.

When we do checkout to a new branch the HEAD revision changes.

run the following command :

git checkout master
cat .git/HEAD

the following output will appear :


the above filepath corresponds to the branch name to which HEAD is currently pointing and that branch further contains the commit id of the commit to which it is pointing.

Note : In git, we can check out to a commit as well. In that case HEAD is called as detached HEAD. One can make changes and commit them as well but those changes won't reflect on any branch because they are not referenced. In order to make them referenced, one can check out a branch from that commit or simply tag that commit.

Below is diagram to explain how HEAD changes with checkout and what is detached HEAD :

1.) HEAD points to the master branch :

2.) Add a new commit C3 on master branch :

3.) checkout to commit C2 :

4.) add a new commit C4 :

5.) checkout a new branch from recent commit :

