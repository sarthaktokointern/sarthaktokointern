Git Internals Part 6 : Stash, Revert and Reset

1.) Stash -> this command saves all the uncommitted changes (both staged and unstaged) into a working directory so that they can be used later and then reverts the changes.
A new stash commit is formed which points to uncommitted changes (Tracked, Untracked and Ignored).

<img width="604" alt="Screenshot 2022-08-29 at 4 17 47 PM" src="https://user-images.githubusercontent.com/99721005/187184650-e1252d02-8c3e-4f61-9478-a405dfc43723.png">



2.) Revert -> this command also forms a new commit with reverted changes of a particular commit.
This prevents Git from losing history, which is important for the integrity of your revision history and for reliable collaboration.

<img width="554" alt="Screenshot 2022-08-29 at 4 18 27 PM" src="https://user-images.githubusercontent.com/99721005/187184765-709006ef-c644-4240-ba52-6cedd8e31d36.png">


3.) Reset -> this command also serves the undo functionality but if we want to undo the changes of a particular commit which lies somewhere in middle of various commits,
then this command will first remove the commits which lie in front of the target commit, then after reaching the target commit it will revert the changes of that commit then will reapply all the commits which were in front of the target commit.

<img width="568" alt="Screenshot 2022-08-29 at 4 18 41 PM" src="https://user-images.githubusercontent.com/99721005/187184807-a38b3626-bd4b-442a-9a6e-bd5042041661.png">


Revert vs Reset -> Revert is better because it is safe to use as the git history is not altered and is maintained.
