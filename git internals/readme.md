Git Internal Architecture :

Introduction : Git is a version control system. Most of the version control systems (vcs) use delta storage system that is storing incremental difference of each file in a commit with every increasing revision.

![](../../../../../../../../var/folders/q3/tl6vmb3d7zv64nchn4lhn2th0000gn/T/TemporaryItems/NSIRD_screencaptureui_yVrBlg/Screenshot 2022-08-10 at 1.46.11 PM.png)

Disadvantage of using delta storage system :

Suppose in the above-mentioned diagram if we want to move from C5 revision to C4 revision, we will just unapply the patch of C5 revision. But if we want to move from nth revision to (n-1000)th revision then we will have to unapply 1000 patches which is a very difficult task.

How to overcome the disadvantage of delta storage system :

In order to overcome the disadvantage possessed by the delta storage system, git uses directed acyclic graph (dag).

What is DAG?

DAG is directed acyclic graph which means that each edge is a directed edge and while traversing the graph each vertex can be visited at most one time i.e. there is no cycle present in the graph.

Objects in git :

Objects are basically immutable files saved on disk which hash key as their name.
There are 4 types of objects in git :

1.) Blob : They are basically large binary objects which represent any file added in the git.
They store data in compressed format. Point to be noted is that each time when a file is added with a **new content**, a new blob is created corresponding to that file.

Hands on : 

a.) create a new folder in your home dir : mkdir test

b.) switch to that folder : cd test

c.)  initialise a git repo : git init 

d.) save a new blank file : vim test.txt

e.) add it to git : git add test.txt

f.) switch to .git : cd .git

g.) list all the directories : ls -la

h.) switch to 'objects' directory : cd objects

you will see some output like this : 

![](../../../../../../../../var/folders/q3/tl6vmb3d7zv64nchn4lhn2th0000gn/T/TemporaryItems/NSIRD_screencaptureui_TKJQVn/Screenshot 2022-08-10 at 2.27.52 PM.png)

Now if we will add some content to the file test.txt and add the changes to git then a new blob will get created corresponding to that.
The file structure will be something like this :

![](../../../../../../../../var/folders/q3/tl6vmb3d7zv64nchn4lhn2th0000gn/T/TemporaryItems/NSIRD_screencaptureui_rQtO6z/Screenshot 2022-08-10 at 2.29.52 PM.png)
