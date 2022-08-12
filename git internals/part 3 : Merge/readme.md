Git Internals Part 3 : Merge

1.) Fast Forward Merge : by default this merge happens. Here is an example for the same :

a.) Initially master is pointing to C2 :

b.) check out a new branch and add C3 and C4 commits :

c.) Merge master and new branch :

-> This is the simplest form of merge and no new objects are created in this merge, only the branch references are updated. But in case of diverged branches, fast-forward merge creates a new commit known as merge commit. The following diagram will help in understanding this part better :

a.) Initially master is pointing to C2 commit :

b.) Checkout  new branch and add C3 commit :

c.) Checkout back to master and add C4 commit :

after merging master and new branch :



-> if the branches are not diverged, but we want to perform non fast-forward merge we can use the flag '-no-ff'. In this a new merge commit will be created and it will look something like this :

