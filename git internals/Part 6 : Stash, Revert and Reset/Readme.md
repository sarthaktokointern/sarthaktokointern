Git Internals Part 6 : Stash, Revert and Reset

1.) Stash -> this command saves all the uncommitted changes (both staged and unstaged) into a working directory so that they can be used later and then reverts the changes.
A new stash commit is formed which points to uncommitted changes (Tracked, Untracked and Ignored).

2.) Revert -> this command also forms a new commit with reverted changes of a particular commit.
This prevents Git from losing history, which is important for the integrity of your revision history and for reliable collaboration.

3.) Reset -> this command also serves the undo functionality but if we want to undo the changes of a particular commit which lies somewhere in middle of various commits,
then this command will first remove the commits which lie in front of the target commit, then after reaching the target commit it will revert the changes of that commit then will reapply all the commits which were in front of the target commit.

Revert vs Reset -> Revert is better because it is safe to use as the git history is not altered and is maintained.