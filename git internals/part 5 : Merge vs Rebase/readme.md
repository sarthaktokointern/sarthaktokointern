Git Internals Part 5 : Merge vs Rebase

Simple and Familiar Process : If someone is new to git, then git merge is the best option because it is simple to use and while merging if you get conflicts, then you will get them all at once and can be resolved in the merge commit.

So if you had erroneous merge while resolving conflict, and you get to know about it in some time future, you can always checkout to commit before merge and check the code part missed.

With rebase original commits are over-ridden with rebased commit and original commits are left dangling. So you tend to lose code information.

In above diagram [Diagram 1] the merge process creates a new merge commit C4 with branch start pointing to it. All merge conflict specific changes are part of C4 and your original commit C2 remains untouched.
While in rebase the C2 commit is left dangling and a new commit C2' is created. If there is a merge conflict when rebasing new_branch C2' commit will hold them and C2 being dangling commit will be destroyed when garbage collection runs.

Preserving History in Original form :

From [Diagram 1] you can see that git merge has tendency of always moving forward without mutating older history, i.e. whenever you do a git merge, a new merge commit is created.
While in rebase a new commit is created replacing older commit i.e. history is mutated.
So if priority is non mutated history and keep it in original form merge is a way to go forward, as you will have commits in original state without any mutation that happened during merge conflict in rebase, so you can always rewind to commits before to check for implementation.

Streamlined and Clean Commit History :

If one wants to achieve a streamlined and clean git history, then rebase is the way to go.



The above diagram [Diagram 2] show 2 developers (Dev1 and Dev2) checking out their branch from C1 commit and adding C2 and C5 commit respectively. Commit C3 is added in master branch.

Now when a Dev1 merge his branch to master the commit C4 is created in master branch and later when Dev2 merge his commit in C6.
Later Dev1 pull the master and checkout to new branch and added C7 commit and similarly Dev2 checkout to new branch and added C8. Here Dev1 now merged first, merge commit not created as branch pointer fast-forwarded to C7 (i.e. why it is represented with master color gradient).

Now when Dev2 merge his branch and a new commit C9 is created as commit history is diverged.

You can easily point out with only 3 merges the history looks too much diverged and converged. Also commit history forms DAG (Directed Acyclic Graph); so if you have to perform any operation on commit history you have to use graph algorithms.

Lets us now see how it will work in rebase


As the name says rebase always re-establishes the base when there is diversion in commit history. You can see that a new commit is created for C2, C5 and C8 and these commits are discarded and destroyed during garbage collection process. For C7 since diversion was not there during rebase so branch just fast-forwarded to C7.

Your history is very clean with commits forming a linked list [Diagram 3]. To operate on this history you can easily go on with more efficient list based algorithms.
Moreover unwanted commits are discarded and you prevent commit explosion which was happening in merge.

But while rebasing we should always take care of golden rule to rebase i.e. not to rebase your shared branch (like master). Always rebase your developer branch.

Note : If we need to squash multiple commits into fewer commits then rebase is the way to go.

**Golden rule : Always rebase individually owned developer from shared branch and then merge individually owned developer branch to shared branch.**