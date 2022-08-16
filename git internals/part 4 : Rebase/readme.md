Git Internals Part 4 : Rebase

1.) How Rebase works : Rebase means to establish a new base. The following steps are performed while rebasing a target branch w.r.t source branch -> 

a.) Switch/Checkout to the source branch

b.) Find the fork point or the point from where branches are getting diverged

c.) After finding the fork point it reapplies each commit present in the target branch after fork commit create new commits. For reapplying commits, it uses commit patch.

The below diagram shows how rebase is done : 



We can see that before rebase 2 branches master and new_branch diverged from C2 commit. So C2 commit is the base commit or the common ancestor point between these 2 branches. When we do rebase; the process checkout to target branch that is master and starts applying commit which are present in new_branch after fork point i.e. C2.

In process of applying commit patch it creates new commit C3' which is same patch as C3 (in case of file merge conflict C3' will also contain resolved files).
The fork point between two branches is changed to commit C4; means a new base is established between master and new_branch. That is why this process is called rebase.

Above we can see that now both branches are on same commit linked list rather than being diverged. So whenever we rebase the target branch a new fork commit will establish to commit on which source branch currently points. The commit C3 becomes dangling commit as it is not being referenced by any branch or tag and will be destroyed during git garbage collection process.

Note : Golden Rule For Rebase states that we should not rebase public/shared branches like master as the target branch because rebase changes its history from that of remote and you need to perform force push to remote after rebase.

Conflicts while rebasing : We might get conflicts while rebasing target branch with the source branch while applying one of the commit patches. In that case we are provided with 3 options ->

a.) Resolve the conflict and do git rebase --continue , it will continue applying same commit which failed with your resolved conflict changes.

b.) Skip the applied commit using git rebase --skip , it will skip to add current commit in rebase process and changes from current commit are removed forever after rebase.

c.) Abort the complete rebase process using git rebase --abort , it will abort the complete process and return to the older state of target branch.

Working files for rebase : 

We might also wonder how rebase process tracks all the process and help us to return back to previous state on performing abort.

When we start a rebase process; a directory .git/rebase-apply gets created. Incase a merge conflict is encountered during rebase we can go and check out this directory.

It consists of commit patches in form of file stating blob objects file moved from one to another in that particular commit.