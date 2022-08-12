Git Internals Part 2 : Branching and Head

1.) Branching : These are just references to commit. They are just like tags, but they are not considered as objects in git because objects are immutable and branches are mutable.

One can find branches in the following directory : .git/refs/heads

This directory will contain name of all the branches. 

Run the following command : cat .git/refs/heads/<branch_name> 

After running the above-mentioned command you will see the commit id to which that particular branch is pointing.

Now let's take a hypothetical scenario for branching :

1.) Create a git repo and create 2 commits C1 and C2 on branch master


2.) Checkout a new branch test and create 2 commits C3 and C4


3.) Checkout master branch and create commit C5


4.) Create a new branch test1 and create commit C6